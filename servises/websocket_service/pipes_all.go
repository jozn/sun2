package websocket_service

import (
	"ms/sun_old/helper"
	"sync"

	"ms/sun/shared/x"

	"github.com/gorilla/websocket"
)

//todo change UserDevicePipe => *UserDevicePipe
type pipesMap struct {
	mp map[int]*UserDevicePipe
	m  sync.RWMutex
}

func (m *pipesMap) IsPipeOpen(UserId int) bool {
	pipe, ok := m.GetUserPipe(UserId)
	if ok && pipe.IsOpen {
		return true
	}
	return false
}

func (m *pipesMap) SendToUser(UserId int, cmd x.PB_CommandToClient) {
	pipe, ok := m.GetUserPipe(UserId)

	logPipes.DevPrintf("pipesMap.SendToUser() no callback to user:%d %v %v ", UserId, ok, cmd.Command)

	if ok && pipe.IsOpen {
		defer func() {
			if r := recover(); r != nil {
				pipe.ShutDownCompletely()
                logPipes.DevPrintf("Recovered in SendToUser: ", r)
			}
		}()
		pipe.SendToUser(cmd)
	}
}

func (m *pipesMap) SendToUserWithCallBack(UserId int, call x.PB_CommandToClient, callback func()) {
	m.SendToUserWithCallBacks(UserId, call, callback, nil)
}

func (m *pipesMap) SendToUserWithCallBacks(UserId int, call x.PB_CommandToClient, callback func(), errback func()) {
	pipe, ok := m.GetUserPipe(UserId)
	logPipes.DevPrintf("pipesMap sending to user withCallbacks:%d %v %v ", UserId, ok, call.Command)

	if ok && pipe.IsOpen {
		defer func() {
			if r := recover(); r != nil {
				pipe.ShutDownCompletely()
				logPipes.DevPrintf("Recovered in SendToUser: ", r)
			}
		}()

		resCallback := commandToClientCallback{
			serverCallId: call.ServerCallId,
			success:      callback,
			error:        errback,
			timeoutAtMs:  helper.TimeNowMs() + 5000,
		}

		commandsToClientMapRegistery.register(resCallback)
		pipe.SendToUser(call)
	} else {
		logPipes.DevPrintf(" pipesMap SendToUserWithCallBacks: user is not connected (just run errback if provided) ")
		if errback != nil {
			errback()
		}
	}
}

func (m *pipesMap) GetUserPipe(UserId int) (*UserDevicePipe, bool) {
	m.m.RLock()
	pipe, ok := m.mp[UserId]
	m.m.RUnlock()
	return pipe, ok
}

func (m *pipesMap) ShutDownUser(UserId int) {
	pipe, ok := m.GetUserPipe(UserId)
	if ok {
		pipe.ShutDown()
		m.DeleteUserPipe(UserId)
	}
}

//adds a new pip
func (m *pipesMap) ServeNewHttpWsForUser(UserId int, ws *websocket.Conn) {
	pipe := &UserDevicePipe{
		UserId:       UserId,
		ToDeviceChan: make(chan x.PB_CommandToClient, 10),
		IsOpen:       true,
		Ws:           ws,
	}

	helper.DebugPrintln("serving user ws for user: ", UserId)
	logPipes.Println("serving user ws for user: ", UserId)

	pipe.ServeIncomingCalls()
	pipe.ServeSendToUserDevice()

	m.AddUserPipe(UserId, pipe)
}

func (m *pipesMap) AddUserPipe(UserId int, pipe *UserDevicePipe) {
	logPipes.Println("pipesMap AddUserPipe UserId ", UserId)
	m.m.Lock()
	m.mp[UserId] = pipe
	m.m.Unlock()
}

func (m *pipesMap) DeleteUserPipe(UserId int) {
	logPipes.Println("pipesMap DeleteUserPipe UserId ", UserId)
	m.m.Lock()
	delete(m.mp, UserId)
	m.m.Unlock()
}
