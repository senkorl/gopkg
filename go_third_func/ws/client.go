package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"runtime"
	"sync"
	"time"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"go.uber.org/zap"
)

var (
	client   *Client
	WriteErr = errors.New("websocket write error")
	ReadErr  = errors.New("websocket read error")
)

const (
	pingWait     = time.Second * 5
	readTimeout  = time.Second * 15
	writeTimeout = time.Second * 5
)

type Client struct {
	context context.Context
	cancel  context.CancelFunc
	//done    chan struct{}
	logger *zap.Logger
	mu     sync.RWMutex

	url             string
	deviceNo        string
	cipherSecret    string
	cipherPublicKey string

	conn               net.Conn
	writeChan          chan []byte
	readChan           chan []byte
	errChan            chan error
	pingWait           time.Duration
	isConnected        bool
	writeSleepDuration int
	pingHandler        func()
}

func NewClient(ctx context.Context, logger *zap.Logger, url, deviceNo, cipherSecret, cipherPublicKey string) *Client {
	client = &Client{
		context:         ctx,
		url:             url,
		deviceNo:        deviceNo,
		cipherSecret:    cipherSecret,
		cipherPublicKey: cipherPublicKey,
		writeChan:       make(chan []byte, 1024),
		readChan:        make(chan []byte, 1024),
		errChan:         make(chan error),
		//done:               make(chan struct{}),
		pingWait:           pingWait,
		logger:             logger,
		writeSleepDuration: 1,
		pingHandler:        func() {},
	}
	return client
}

func (c *Client) Run() {
	ctx, cancel := context.WithCancel(c.context)
	c.cancel = cancel
	go c.readLoop(ctx)
	go c.writeLoop(ctx)
}

func (c *Client) Connect() error {
	c.mu.Lock()
	defer c.mu.Unlock()
	//if c.done == nil {
	//	return errors.New("websocket client is done")
	//}
	if c.conn != nil {
		_ = c.conn.Close()
	}
	var err error
	serverUrl := fmt.Sprintf("%s?deviceNo=%s", c.url, c.deviceNo)
	conn, _, _, err := ws.Dial(c.context, serverUrl)
	if err != nil {
		c.logger.Error("Connect", zap.Error(err))
		c.isConnected = false
		return err
	}
	runtime.SetFinalizer(conn, func(conn net.Conn) {
		_ = conn.Close()
	})
	c.logger.Info("connected:", zap.String("url", serverUrl), zap.String("local_addr", conn.LocalAddr().String()))
	c.conn = conn
	if err = c.online(); err != nil {
		return err
	}
	c.isConnected = true
	return nil
}

// IsConnected returns the WebSocket connection state
func (c *Client) IsConnected() bool {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return c.isConnected
}

func (c *Client) setIsConnected(isConnected bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.isConnected = isConnected
}

func (c *Client) ReceiveMessage() <-chan []byte {
	return c.readChan
}

func (c *Client) WriteMessage(writeArgs []byte) {
	c.writeChan <- writeArgs
}

func (c *Client) ReceiveErrorResult() <-chan error {
	return c.errChan
}

func (c *Client) readLoop(ctx context.Context) {
	defer Recover(c.logger)
	defer func() {
		c.logger.Debug("websocket.read.loop.done")
	}()
	for {
		select {
		case <-ctx.Done():
			return
		default:
			if !c.IsConnected() {
				time.Sleep(time.Second)
				continue
			}
			var msg []wsutil.Message
			_ = c.conn.SetReadDeadline(time.Now().Add(readTimeout))
			msg, err := wsutil.ReadServerMessage(c.conn, msg)
			if err != nil {
				c.setIsConnected(false)
				//if c.isDone() {
				//	return
				//}
				c.logger.Error("readLoop:websocket IsUnexpectedCloseError", zap.Error(err))
				//if err = c.Connect(); err == nil {
				//	continue
				//}
				//if c.errChan != nil {
				//	c.errChan <- ReadErr
				//}
				continue
			}
			if len(msg) == 0 {
				continue
			}
			// 将接收消息推送出去
			for _, m := range msg {
				if m.OpCode.IsControl() {
					_ = wsutil.HandleServerControlMessage(c.conn, m)
					continue
				}
				c.readChan <- m.Payload
				if chanLen := len(c.readChan); chanLen > 1 {
					c.logger.Debug("ws.read.block", zap.Any("len", chanLen))
				}
			}
		}
	}
}

type PingExt struct {
	DeviceNo string `json:"deviceNo"`
}

func (c *Client) writeLoop(cxt context.Context) {
	defer Recover(c.logger)
	ticker := time.NewTicker(c.pingWait)
	defer func() {
		ticker.Stop()
		c.logger.Debug("websocket.write.loop.done")
	}()
	// 写消息到 WebSocket 服务器
	for {
		select {
		case <-cxt.Done():
			return
		case <-ticker.C:
			if !c.IsConnected() {
				_ = c.Connect()
				continue
			}
			_ = c.conn.SetWriteDeadline(time.Now().Add(writeTimeout))
			var pingData = PingExt{DeviceNo: c.deviceNo}
			pingBys, err := json.Marshal(pingData)
			if err != nil {
				continue
			}
			c.logger.Debug("websocket write loop ticker")
			if err = wsutil.WriteClientMessage(c.conn, ws.OpPing, pingBys); err != nil {
				c.setIsConnected(false)
				c.logger.Error("writeLoop:ticker failed", zap.Error(err))
			}
			c.pingHandler()
		case msg := <-c.writeChan:
			if !c.IsConnected() {
				c.logger.Error("conn is nil or is not connected")
				continue
			}
			_ = c.conn.SetWriteDeadline(time.Now().Add(writeTimeout))
		Resend:
			err := wsutil.WriteClientMessage(c.conn, ws.OpText, msg)
			if err != nil {
				c.setIsConnected(false)
				//if c.isDone() {
				//	c.logger.Debug("websocket.write.loop.done")
				//	return
				//}
				c.logger.Error("websocket write message failed", zap.Error(err))
				err = c.Connect()
				if err == nil {
					goto Resend
				}
				if err != nil && c.errChan != nil {
					c.errChan <- WriteErr
					continue
				}
			}
		}
	}
}

//helper_func (c *Client) Done() chan struct{} {
//	return c.done
//}

//helper_func (c *Client) isDone() bool {
//	c.mu.Lock()
//	defer c.mu.Unlock()
//	return c.done == nil
//}

func (c *Client) Close() {
	c.mu.Lock()
	defer c.mu.Unlock()
	//if c.done == nil {
	//	return
	//}
	//close(c.done)
	//c.done = nil
	c.cancel()
	c.isConnected = false
	frame := ws.NewCloseFrame(ws.NewCloseFrameBody(
		ws.StatusNormalClosure, "",
	))
	frame = ws.MaskFrameInPlace(frame)
	if err := ws.WriteFrame(c.conn, frame); err != nil {
		c.logger.Error("Websocket.Close.WriteFrame", zap.Error(err))
	}
	if err := c.conn.Close(); err != nil {
		c.logger.Error("Websocket.Close.Conn", zap.Error(err))
	}
}

func (c *Client) SetPingHandler(f func()) {
	c.pingHandler = f
}
