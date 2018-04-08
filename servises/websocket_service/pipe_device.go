package websocket_service

import (
	"sync"

	"github.com/gorilla/websocket"

	"github.com/golang/protobuf/proto"
	"log"
	"ms/sun_old/config"
	"ms/sun_old/helper"
	"ms/sun/shared/x"
)

const PB_CommandReceivedToServer = "PB_CommandReceivedToServer"
const PB_CommandReceivedToClient = "PB_CommandReceivedToClient"

type UserDevicePipe struct {
	UserId         int
	IsOpen         bool
	Ws             *websocket.Conn
	ToDeviceChan   chan x.PB_CommandToClient
	FromDeviceChan chan x.PB_CommandToServer // NOT NEEDED?
	m              sync.RWMutex
	hasShutDown    bool
}

func (pipe *UserDevicePipe) ServeIncomingCalls() {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				if config.IS_DEBUG {
					log.Panic("recverd in ServeSendToUserDevice err: ", r)
				}
				logPipes.DevPrintln("Recovered in ws messaging clinet request", r)
				pipe.ShutDownCompletely()
			}
		}()

		for {
			messageType, bytes, err := pipe.Ws.ReadMessage() //blocking

			logPipes.DevPrintln("messageType: ", " ::", messageType, string(bytes))
			if messageType == websocket.CloseMessage || err != nil {
				pipe.ShutDownCompletely()
				logPipes.DevPrintf("closeing pip for userId: %v , messageType:%v , err: %v", pipe.UserId, messageType, err)
				return
			}

			if messageType == websocket.TextMessage {

			}

			if messageType == websocket.BinaryMessage {
				pb := &x.PB_CommandToServer{}
				err := proto.Unmarshal(bytes, pb)
				if err == nil {
					if config.IS_DEBUG {
						//wsDebugLog.DevPrintln("-> from device:", pb.Command, helper.ToJson(pb))
						logPipes.DevPrintln("-> from device:", pb.Command, helper.ToJson(pb))
					}
					serverWSReqCalls(pb, pipe)
				}
			}
		}
	}()
}

func (pipe *UserDevicePipe) ServeSendToUserDevice() {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				if config.IS_DEBUG {
					log.Panic("recverd in ServeSendToUserDevice err: ", r)
				} else {
					pipe.ShutDownCompletely()
				}
			}
		}()
		for r := range pipe.ToDeviceChan {
			if config.IS_DEBUG {
				//wsDebugLog.DevPrintln("<- to device:", pipe.UserId, helper.ToJson(r))
				logPipes.DevPrintln("<- to device:", pipe.UserId, helper.ToJson(r))
			}
			bts, err := proto.Marshal(&r)
			if err == nil {
				pipe.Ws.WriteMessage(websocket.BinaryMessage, bts)
			}
		}
		pipe.ShutDownCompletely()
	}()
}

func (pipe *UserDevicePipe) SendToUser(resCall x.PB_CommandToClient) {
	if pipe.IsOpen {
		pipe.ToDeviceChan <- resCall
	}
}

func (pipe *UserDevicePipe) ShutDown() {
	if pipe.hasShutDown == true {
		return
	}
	pipe.m.Lock()
	pipe.Ws.Close()
	close(pipe.ToDeviceChan)
	pipe.hasShutDown = true
	pipe.m.Unlock()
}

func (pipe *UserDevicePipe) ShutDownCompletely() {
	defer helper.JustRecover() //maybe double close chanel
	pipe.ShutDown()
	AllPipesMap.ShutDownUser(pipe.UserId)
}
