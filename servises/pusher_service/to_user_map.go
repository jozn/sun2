package pusher_service

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
	"io"
	"log"
	"ms/sun/shared/config"
	"ms/sun/shared/helper"
	"ms/sun/shared/x"
	"sync"
	"time"
)

type userMultiPipes struct {
	userId    int
	data      chan x.PB_Push
	devicesWS []*wsSinglePipe
}

func (ud *userMultiPipes) start() {
	go ud.startLoop_go()
}

func (ud *userMultiPipes) startLoop_go() {
	for pb_push := range ud.data {
		for _, dev := range ud.devicesWS {
			if dev != nil && dev.isOpen {
				dev.data <- pb_push
			}
		}
	}
}

func (ud *userMultiPipes) shouldClean() bool {
	i := 0
	for _, dev := range ud.devicesWS {
		if dev != nil && dev.isOpen {
			i++
		}
	}
	if i == 0 {
		return true
	}
	return false
}

func (ud *userMultiPipes) finish() {

}

////////////////////
var pipId = 0
var cnt = 0

type wsSinglePipe struct {
	pipId    int
	userId   int
	isOpen   bool
	data     chan x.PB_Push
	ws       *websocket.Conn
	ioStream io.Writer //if debug ws is irrelevant
}

func (pipe *wsSinglePipe) start() {
	go pipe.startLoop_go()
	go pipe.startRecivingFrames_go()
}

func (pipe *wsSinglePipe) startLoop_go() {
	logPush.Printf("startLoop_go() : user: %d \n", pipe.userId)
	defer func() {
		if r := recover(); r != nil {
			if config.IS_DEBUG {
				log.Panic("recverd in serveSendToUserDevice err: ", r)
			}
			pipe.isOpen = false
			if pipe.ws != nil {
				pipe.ws.Close()
			}
		}
	}()
	for r := range pipe.data {
		logPush.Printf("wsSinglePipe data: %v \n", r)
		if !pipe.isOpen {
			continue
		}
		var strout = ""
		if config.IS_DEBUG {
			strout = fmt.Sprintf("<- to device: %d ,%s \n", pipe.userId, helper.ToJson(r))
			logPushPipes.DevPrint(strout)
		}
		bts, err := proto.Marshal(&r)
		if err == nil {
			if pipe.ws != nil {
				pipe.ws.WriteMessage(websocket.BinaryMessage, bts)
			}
			if pipe.ioStream != nil {
				pipe.ioStream.Write([]byte(strout))
				//pipe.ioStream.Write(bts)
			}
		}
	}
}

func (pipe *wsSinglePipe) startRecivingFrames_go() {
	defer func() {
		if r := recover(); r != nil {
			Monitor.TotalWebSocketsReceivedPanics += 1
			if config.IS_DEBUG {
				log.Panic("recverd in serveSendToUserDevice err: ", r)
			}
			logPushPipes.DevPrintln("Recovered in ws messaging clinet request", r)
			pipe.isOpen = false
		}
	}()

	if pipe.ws != nil {
		for {
			messageType, bytes, err := pipe.ws.ReadMessage() //blocking

			logPushPipes.DevPrintln("messageType: ", " ::", messageType, string(bytes))
			if messageType == websocket.CloseMessage || err != nil {
				Monitor.TotalWebSocketsClosed += 1
				pipe.isOpen = false
				logPushPipes.DevPrintf("closeing pip for userId: %v , messageType:%v , err: %v", pipe.userId, messageType, err)
				return
			}

			if messageType == websocket.TextMessage {
				Monitor.TotalWebSocketsReceivedText += 1
			}

			if messageType == websocket.BinaryMessage {
				Monitor.TotalWebSocketsReceivedBinary += 1
			}
		}
	}
}

///////////////////
var allPipesMap *pipesMap = newPipesMap()

type pipesMap struct {
	mp map[int]*userMultiPipes
	m  sync.RWMutex
}

func newPipesMap() *pipesMap {
	return &pipesMap{
		mp: make(map[int]*userMultiPipes, 1000),
	}
}

func (m *pipesMap) getForUser(userId int) (ud *userMultiPipes, ok bool) {
	m.m.RLock()
	defer m.m.RUnlock()
	ud, ok = m.mp[userId]
	return ud, ok
}

func (m *pipesMap) addUserDevice(userId int, ws *websocket.Conn, ioStream io.Writer) {
	Monitor.TotalUsersAdded += 1
	pipId++
	pipe := &wsSinglePipe{
		pipId:  pipId,
		userId: userId,
		isOpen: true,
		data:   make(chan x.PB_Push, 10),
		ws:     ws,
	}
	pipe.start()
	if ioStream != nil {
		pipe.ioStream = ioStream
	}
	m.m.Lock()
	defer m.m.Unlock()
	ud, ok := m.mp[userId]
	if !ok {
		ud = &userMultiPipes{
			userId: userId,
			data:   make(chan x.PB_Push, 100),
		}
		ud.start()
		m.mp[userId] = ud
		logPush.DevPrintf("add userMultiPipes to AllMapPipes: user: %d \n", userId)
	}
	ud.devicesWS = append(ud.devicesWS, pipe)

	logPush.DevPrintf("addUserDevice() : user: %d , userMultiPipes: %v \n", userId, ud)

}

func (m *pipesMap) cleanUnConnectedUsersLoop() {
	for {
		time.Sleep(time.Second * 1)
		m.m.Lock()
		for userId, ud := range m.mp {
			if ud.shouldClean() {
				delete(m.mp, userId)
				ud.finish()
				logPushPipes.DevPrintln("cleaned userid:", userId)
			}
		}
		m.m.Unlock()
	}
}
