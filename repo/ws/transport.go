package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"strconv"
	"time"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"go.uber.org/zap"
)

func (c *Client) online() error {
	c.logger.Debug("websocket online")
	statusData := c.GenerateWebsocketWriteParam(DeviceOnlineStatus, "", EndSignEnd, nil)
	b, _ := json.Marshal(statusData)
	return wsutil.WriteClientMessage(c.conn, ws.OpText, b)
}

const (
	DeviceOnlineStatus    = "device_online_status"
	DeviceOnlineReady     = "device_online_ready"
	DeviceWorkflowExecute = "device_workflow_execute"
	DeviceCommandCallback = "command_callback"
	DeviceAsrAudioText    = "device_asr_audio_text"
)

const (
	EndSignStart   = "start"
	EndSignEnd     = "end"
	EndSignPause   = "pause"
	EndSignRecover = "recover"
	EndSignProcess = "process"
)

const (
	Source  = "DEVICE"
	BizCode = "toy"
)

func (c *Client) GenerateWebsocketWriteParam(action, data, endSign string, extra *WriteBodyExtra) WriteBody {
	digestSecret, err := AesDecrypt(c.cipherSecret, c.cipherPublicKey)
	if err != nil {
		c.logger.Error("AutoCheckVersion:AesDecrypt", zap.Error(err))
		return WriteBody{}
	}
	if digestSecret == "" {
		c.logger.Error("digest secret is empty")
		return WriteBody{}
	}
	var stamp = strconv.FormatInt(time.Now().Unix(), 10)
	if extra == nil {
		extra = &WriteBodyExtra{}
	}
	return WriteBody{
		Action:    action,
		BizCode:   BizCode,
		DeviceNo:  c.deviceNo,
		Source:    Source,
		Timestamp: stamp,
		EndSign:   endSign,
		ExtraMap:  extra,
		Data:      data,
		Digest:    DigestSecret(data, digestSecret, stamp),
	}
}

func DigestSecret(data, secretNo, timestamp string) string {
	return GetMD5([]byte(data + secretNo + timestamp))
}

func GetMD5(src []byte) string {
	h := md5.New()
	_, _ = h.Write(src)
	return hex.EncodeToString(h.Sum(nil))
}

// Aes 解密
func AesDecrypt(cipherStr, key string) (string, error) {
	ciphertext, err := base64.StdEncoding.DecodeString(cipherStr)
	if err != nil {
		return "", err
	}
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}
	if len(ciphertext) < aes.BlockSize {
		return "", err
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(ciphertext, ciphertext)
	// 删除PKCS#7填充
	plaintext := unpad(ciphertext, aes.BlockSize)
	return string(plaintext), nil
}

// unpad 删除PKCS#7填充的数据
func unpad(buf []byte, blockSize int) []byte {
	if len(buf)%blockSize != 0 {
		return nil
	}
	padding := int(buf[len(buf)-1])
	return buf[:len(buf)-padding]
}

// Upload
// @description 通用接口，上传数据到远端
func (c *Client) Upload(action, body, signStatus string, extra *WriteBodyExtra) {
	c.logger.Debug("Websocket.Upload", zap.Any("action", action), zap.Any("signStatus", signStatus), zap.Any("extra", extra))
	writeBody := c.GenerateWebsocketWriteParam(action, body, signStatus, extra)
	writeBys, err := json.Marshal(writeBody)
	if err != nil {
		c.logger.Error("Websocket WriteBody Marshal", zap.Error(err))
		return
	}
	c.WriteMessage(writeBys)
}
