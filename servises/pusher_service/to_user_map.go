package pusher_service

import (
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
	"io"
	"log"
	"ms/sun/shared/config"
	"ms/sun/shared/helper"
	"ms/sun/shared/x"
	"sync"
	"time"
    "fmt"
)

type userMultiDevicePusher struct {
	userId    int
	data      chan x.PB_Push
	devicesWS []*wsSingleDevicePipe
}

func (ud *userMultiDevicePusher) startLoop_go() {
	for pb_push := range ud.data {
		for _, dev := range ud.devicesWS {
			if dev != nil && dev.isOpen {
				dev.data <- pb_push
			}
		}
	}
}

func (ud *userMultiDevicePusher) shouldClean() bool {
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

////////////////////
var pipId = 0
var cnt  = 0

type wsSingleDevicePipe struct {
	pipId    int
	userId   int
	isOpen   bool
	data     chan x.PB_Push
	ws       *websocket.Conn
	ioStream io.Writer //if debug ws is irrelevant
}

func (pipe *wsSingleDevicePipe) startLoop_go() {
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
		logPush.Printf("wsSingleDevicePipe data: %v \n", r)
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

///////////////////
var allPipesMap *pipesMap = newPipesMap()

type pipesMap struct {
	mp map[int]*userMultiDevicePusher
	m  sync.RWMutex
}

func newPipesMap() *pipesMap {
	return &pipesMap{
		mp: make(map[int]*userMultiDevicePusher, 1000),
	}
}

func (m *pipesMap) getForUser(userId int) (ud *userMultiDevicePusher, ok bool) {
	m.m.RLock()
	defer m.m.RUnlock()
	ud, ok = m.mp[userId]
	return ud, ok
}

func (m *pipesMap) addUserDevice(userId int, ws *websocket.Conn, ioStream io.Writer) {
	pipId++
	pipe := &wsSingleDevicePipe{
		pipId:  pipId,
		userId: userId,
		isOpen: true,
		data:   make(chan x.PB_Push, 1),
		ws:     ws,
	}
	go pipe.startLoop_go()
	if ioStream != nil {
		pipe.ioStream = ioStream
	}
	m.m.Lock()
	defer m.m.Unlock()
	ud, ok := m.mp[userId]
	if !ok {
		ud = &userMultiDevicePusher{
			userId: userId,
			data:   make(chan x.PB_Push, 100),
		}
		go ud.startLoop_go()
		m.mp[userId] = ud
		logPush.Printf("add userMultiDevicePusher to AllMapPipes: user: %d \n", userId)
	}
	ud.devicesWS = append(ud.devicesWS, pipe)

	logPush.Printf("addUserDevice() : user: %d , userMultiDevicePusher: %v \n", userId, ud)

}

func (m *pipesMap) cleanUnConnectedUsersLoop() {
	for {
		time.Sleep(time.Second * 1)
		m.m.Lock()
		for userId, ud := range m.mp {
			if ud.shouldClean() {
				delete(m.mp, userId)
				logPushPipes.Println("cleaned userid:", userId)
			}
		}
		m.m.Unlock()
	}
}
