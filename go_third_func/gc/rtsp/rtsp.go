package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/bluenviron/gortsplib/v4"
	"github.com/bluenviron/gortsplib/v4/pkg/base"
	"github.com/bluenviron/gortsplib/v4/pkg/format"
	"github.com/bluenviron/gortsplib/v4/pkg/format/rtph265"
	"github.com/bluenviron/mediacommon/v2/pkg/codecs/h265"
	"github.com/bluenviron/mediacommon/v2/pkg/formats/mpegts"
	"github.com/pion/rtp"
	onvif2 "github.com/quocson95/go-onvif"
)

func main() {
	camHost := "192.168.1.25"
	username := "admin"
	password := "123456"
	dev1 := onvif2.Device{
		XAddr:    "http://" + camHost + "/onvif/device_service",
		User:     username,
		Password: password,
	}
	caps, err := dev1.GetCapabilities()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(caps)
	pfs, err := dev1.GetProfiles()
	if err != nil {
		log.Fatal(err)
	}
	mediaUri, err := dev1.GetStreamURI(pfs[0].Token, "rtsp")
	if err != nil {
		log.Fatal(err)
	}
	ps, err := dev1.CreatePullPointSubscription()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ps)
	unsubUrl, err := dev1.Subscribe(caps.EventsCap.XAddr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("unsubUrl:", unsubUrl)
	defer dev1.UnSubscribe(unsubUrl)
	go func() {
		for {
			msg, err := dev1.PullMessages(unsubUrl)
			if err != nil {
				return
			}
			if len(msg) > 0 {
				fmt.Println(msg)
			}
			time.Sleep(time.Second * 2)
		}
	}()
	c := gortsplib.Client{}
	// parse URL
	u, err := base.ParseURL(mediaUri.URI)
	if err != nil {
		panic(err)
	}
	// connect to the server
	err = c.Start(u.Scheme, u.Host)
	if err != nil {
		panic(err)
	}
	defer c.Close()
	// find available medias
	desc, _, err := c.Describe(u)
	if err != nil {
		panic(err)
	}
	// find the H265 media and format
	var forma *format.H265
	medi := desc.FindFormat(&forma)
	if medi == nil {
		panic("media not found")
	}
	// setup RTP -> H265 decoder
	rtpDec, err := forma.CreateDecoder()
	if err != nil {
		panic(err)
	}
	muxer := &MpegtsMuxer{
		fileName: "mystream.ts",
		vps:      forma.VPS,
		sps:      forma.SPS,
		pps:      forma.PPS,
	}
	err = muxer.initialize()
	if err != nil {
		panic(err)
	}
	defer muxer.close()
	// setup a single media
	_, err = c.Setup(desc.BaseURL, medi, 0, 0)
	if err != nil {
		panic(err)
	}
	// called when a RTP packet arrives
	c.OnPacketRTP(medi, forma, func(pkt *rtp.Packet) {
		// decode timestamp
		pts, ok := c.PacketPTS2(medi, pkt)
		if !ok {
			log.Printf("waiting for timestamp")
			return
		}
		// extract access unit from RTP packets
		au, err := rtpDec.Decode(pkt)
		if err != nil {
			if err != rtph265.ErrNonStartingPacketAndNoPrevious && err != rtph265.ErrMorePacketsNeeded {
				log.Printf("ERR: %v", err)
			}
			return
		}
		// encode the access unit into MPEG-TS
		err = muxer.writeH265(au, pts)
		if err != nil {
			log.Printf("ERR: %v", err)
			return
		}
		//log.Printf("saved TS packet")
	})
	// start playing
	_, err = c.Play(nil)
	if err != nil {
		panic(err)
	}
	// wait until a fatal error
	panic(c.Wait())
}

// MpegtsMuxer allows to save a H265 stream into a MPEG-TS file.
type MpegtsMuxer struct {
	fileName     string
	vps          []byte
	sps          []byte
	pps          []byte
	f            *os.File
	b            *bufio.Writer
	w            *mpegts.Writer
	track        *mpegts.Track
	dtsExtractor *h265.DTSExtractor
}

// initialize initializes a MpegtsMuxer.
func (e *MpegtsMuxer) initialize() error {
	var err error
	e.f, err = os.Create(e.fileName)
	if err != nil {
		return err
	}
	e.b = bufio.NewWriter(e.f)
	e.track = &mpegts.Track{
		Codec: &mpegts.CodecH265{},
	}
	e.w = mpegts.NewWriter(e.b, []*mpegts.Track{e.track})
	return nil
}

// close closes all the MpegtsMuxer resources.
func (e *MpegtsMuxer) close() {
	e.b.Flush()
	e.f.Close()
}

// writeH265 writes a H265 access unit into MPEG-TS.
func (e *MpegtsMuxer) writeH265(au [][]byte, pts int64) error {
	var filteredAU [][]byte
	isRandomAccess := false
	for _, nalu := range au {
		typ := h265.NALUType((nalu[0] >> 1) & 0b111111)
		switch typ {
		case h265.NALUType_VPS_NUT:
			e.vps = nalu
			continue
		case h265.NALUType_SPS_NUT:
			e.sps = nalu
			continue
		case h265.NALUType_PPS_NUT:
			e.pps = nalu
			continue
		case h265.NALUType_AUD_NUT:
			continue
		case h265.NALUType_IDR_W_RADL, h265.NALUType_IDR_N_LP, h265.NALUType_CRA_NUT:
			isRandomAccess = true
		}
		filteredAU = append(filteredAU, nalu)
	}
	au = filteredAU
	if au == nil {
		return nil
	}
	// add VPS, SPS and PPS before random access access unit
	if isRandomAccess {
		au = append([][]byte{e.vps, e.sps, e.pps}, au...)
	}
	if e.dtsExtractor == nil {
		// skip samples silently until we find one with a IDR
		if !isRandomAccess {
			return nil
		}
		e.dtsExtractor = h265.NewDTSExtractor()
	}
	dts, err := e.dtsExtractor.Extract(au, pts)
	if err != nil {
		return err
	}
	// encode into MPEG-TS
	return e.w.WriteH265(e.track, pts, dts, au)
}
