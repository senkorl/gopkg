package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

var (
	url             = "wss://aitoy-dgw-server.girobot.net/ws/dgw-secure"
	dev             = "G0101JK0A00324AB"
	cipherSecret    = "ZikIP+eh8eu+zEerzH5h2sEcMV/kdMO0nUSQeAtZutdEMnKm8xZt9EtNbO55dGrFNAfm8SZqlqUVwNB4cB/QtQ=="
	cipherPublicKey = "8BoQZoCAoJ8gVdAsYChQSZBg0EPApAN8"
)

func main() {
	l := NewLogger(&Config{
		Level:      "debug",
		RootDir:    "./log",
		Filename:   "toy.log",
		MaxSize:    10,
		MaxBackups: 5,
		MaxAge:     1,
		Compress:   false,
		Env:        "prod",
		TZ:         "Asia/Shanghai",
	})
	wsCli := NewClient(context.Background(), l, url, dev, cipherSecret, cipherPublicKey)
	if err := wsCli.Connect(); err != nil {
		panic(err)
	}
	wsCli.Run()
	cc := make(chan struct{})
	go func() {
		tk := time.NewTicker(128 * time.Millisecond)
		for {
			select {
			case <-cc:
				return
			case <-tk.C:
				asrMode := "normal"
				bodyBys, err := json.Marshal(WriteBodyData{})
				if err != nil {
					return
				}
				conversationId := "MAIN_11111111"
				wsCli.Upload(DeviceAsrAudioText, string(bodyBys), EndSignEnd, &WriteBodyExtra{ModelSwitch: asrMode, ConversationId: conversationId})
			}
		}
	}()
	tm := time.NewTimer(30 * time.Second)
	for {
		select {
		case receiveArgs, ok := <-wsCli.ReceiveMessage():
			if !ok {
				return
			}
			fmt.Println(string(receiveArgs))
		case <-tm.C:
			wsCli.Close()
			cc <- struct{}{}
			time.Sleep(15 * time.Second)
			return
		}
	}
}
