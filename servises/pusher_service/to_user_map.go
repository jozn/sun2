package pusher_service

import (
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
	"log"
	"ms/sun/shared/config"
	"ms/sun/shared/helper"
	"ms/sun/shared/x"
	"sync"
	"time"
)

type userDevicesPusher struct {
	userId    int
	data      chan x.PB_Push
	devicesWS []*wsDevicePipe
}

func (ud *userDevicesPusher) sendToDevicesLoop() {
	for pb_push := range ud.data {
		for _, dev := range ud.devicesWS {
			if dev != nil && dev.isOpen {
				_ = pb_push
			}
		}
	}
}

func (ud *userDevicesPusher) shouldClean() bool {
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
type wsDevicePipe struct {
	userId int
	isOpen bool
	data   chan x.PB_Push
	ws     *websocket.Conn
}

func (wd *wsDevicePipe) getForUser(userId int) (*userDevicesPusher, bool) {
	return nil, true
}
func (pipe *wsDevicePipe) serveSendToUserDevice_go() {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				if config.IS_DEBUG {
					log.Panic("recverd in serveSendToUserDevice err: ", r)
				}
				pipe.isOpen = false
				pipe.ws.Close()
			}
		}()
		for r := range pipe.data {
			if !pipe.isOpen {
				continue
			}
			if config.IS_DEBUG {
				logPipes.DevPrintln("<- to device:", pipe.userId, helper.ToJson(r))
			}
			bts, err := proto.Marshal(&r)
			if err == nil {
				pipe.ws.WriteMessage(websocket.BinaryMessage, bts)
			}
		}
	}()
}

///////////////////
var allPipesMap *pipesMap =newPipesMap()

type pipesMap struct {
	mp map[int]*userDevicesPusher
	m  sync.RWMutex
}

func newPipesMap() *pipesMap {
    return &pipesMap{
        mp:make(map[int]*userDevicesPusher,1000),
    }
}

func (m *pipesMap) getForUser(userId int) (ud *userDevicesPusher, ok bool) {
	m.m.RLock()
	ud, ok = m.mp[userId]
	m.m.RUnlock()
	return ud, ok
}

func (m *pipesMap) addUserDevice(userId int, ws *websocket.Conn) {
	pipe := &wsDevicePipe{
		userId: userId,
		isOpen: true,
		data:   make(chan x.PB_Push, 100),
		ws:     ws,
	}
	m.m.Lock()
	defer m.m.Unlock()
	ud, ok := m.mp[userId]
	if !ok {
		ud = &userDevicesPusher{
			userId: userId,
			data:   make(chan x.PB_Push, 100),
		}
		m.mp[userId] = ud
	}
	ud.devicesWS = append(ud.devicesWS, pipe)

}

func (m *pipesMap) cleanUnConnectedUsersLoop() {
	for {
		time.Sleep(time.Second * 10)
		m.m.Lock()
		for userId, ud := range m.mp {
			if ud.shouldClean() {
				delete(m.mp, userId)
				logPipes.Println("cleaned userid:", userId)
			}
		}
	}
}
