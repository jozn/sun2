package websocket_service

import (
	"errors"
	"ms/sun_old/config"
	"ms/sun/shared/helper"
	"sync"
	"time"
)

type commandToClientCallback struct {
	success      func()
	error        func()
	timeoutAtMs  int   // time second // now + 5 sec
	serverCallId int64 // time nano
}

//////////////////////////////////

type commandsToClientMap struct {
	sync.RWMutex
	mp               map[int64]commandToClientCallback
	isRunningTimeout bool
}

func (m *commandsToClientMap) register(callback commandToClientCallback) {
	if callback.serverCallId == 0 {
		logPipes.DevPrintln("ERROr: In commandToClientCallback, callback.serverCallId must not be 0")
		callback.serverCallId = int64(helper.RandomSeqUid())
	}

	logPipes.DevPrintln("commandsToClientMapRegistery register() of: ", callback.serverCallId)

	callback.timeoutAtMs = helper.TimeNowMs() + 5000
	m.Lock()
	m.mp[callback.serverCallId] = callback
	m.Unlock()
}

func (m *commandsToClientMap) get(serverCallId int64) (*commandToClientCallback, error) {
	if serverCallId == 0 {
		return nil, errors.New(" serverCallId could not be 0")
	}
	m.RLock()
	callback, ok := m.mp[serverCallId]
	m.RUnlock()
	if !ok {
		return nil, errors.New(" serverCallId not found in  map")
	}
	return &callback, nil
}

func (m *commandsToClientMap) getAndRemove(serverCallId int64) (*commandToClientCallback, error) {
	if serverCallId == 0 {
		return nil, errors.New(" serverCallId could not be 0")
	}
	m.Lock()
	callback, ok := m.mp[serverCallId]
	delete(m.mp, serverCallId)
	m.Unlock()
	if !ok {
		return nil, errors.New(" serverCallId not found in  map")
	}
	return &callback, nil
}

func (m *commandsToClientMap) remove(serverCallId int64) {
	m.Lock()
	delete(m.mp, serverCallId)
	m.Unlock()
}

func (m *commandsToClientMap) runSucceded(serverCallId int64) {
	logPipes.DevPrintln("*commandsToClientMapRegistery runSucceded() of: ", serverCallId)
	callback, err := m.getAndRemove(serverCallId)
	if err != nil {
		return
	}
	if callback.success != nil {
		callback.success()
	}
}

func (m *commandsToClientMap) runError(serverCallId int64) {
	logPipes.DevPrintln("commandsToClientMapRegistery runError() of: ", serverCallId)
	callback, err := m.getAndRemove(serverCallId)
	if err != nil {
		return
	}
	if callback.error != nil {
		callback.error()
	}
}

func (m *commandsToClientMap) setIsRunningTimeout(is bool) {
	m.Lock()
	m.isRunningTimeout = is
	m.Unlock()
}

func (m *commandsToClientMap) runErrorOfTimeouts() {
	if m.isRunningTimeout {
		return
	}
	m.setIsRunningTimeout(true)
	defer m.setIsRunningTimeout(false)

	var arr []commandToClientCallback

	m.Lock()
	for _, v := range m.mp {
		if v.timeoutAtMs < helper.TimeNowMs() {
			arr = append(arr, v)
		}
	}
	m.Unlock()

	m.Lock()
	for _, v := range arr {
		delete(m.mp, v.serverCallId)
	}
	m.Unlock()

	for _, v := range arr {
		logPipes.DevPrintln("commandsToClientMapRegistery clean errorBack of: %d - has errorBack: %v", v.serverCallId, v.error != nil)
		if v.error != nil {
			v.error()
		}
	}
}

//TODO: infuter make this run in parralll
func intervalRunCommandsCallbacksTimeOutChecker(registery *commandsToClientMap) {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				if config.IS_DEBUG {
					logPipes.DevPrintln("ERROR PANICED RECOVRED -intervalRunCommandsCallbacksTimeOutChecker - ERR:: ", r)
				}
				intervalRunCommandsCallbacksTimeOutChecker(registery)
			}
		}()

		for {
			time.Sleep(time.Second * 1)
			registery.runErrorOfTimeouts()
		}
	}()
}
