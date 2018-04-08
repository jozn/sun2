package message_adder

import (
	"ms/sun_old/helper"
	"ms/sun/shared/x"
	"sync"
	"time"
)

//this one must load from mysql, kafka or for now just simple loading on single node
var MessageAdderChanel = make(chan *x.PB_MessageView) //messges must have a right messageFileId
var msgRouter = msgRouterImple{
	mp: make(map[string]*chatDirect, 10000),
}

func init() {
	go msgRouter.handle()
	go msgRouter.cleaner()
}

type msgRouterImple struct {
	mp map[string]*chatDirect
	sync.RWMutex
}

//there is even no need for locking
func (r *msgRouterImple) handle() {
	defer helper.JustRecover()
	for msg := range MessageAdderChanel {
		r.RLock()
		cDirect := r.mp[msg.RoomKey]
		r.RUnlock()

		if cDirect == nil {
			cDirect = newChaatDirectByRoomKey(msg.RoomKey)

			r.Lock()
			r.mp[msg.RoomKey] = cDirect
			r.Unlock()
		}

		cDirect.AddChan <- msg
	}
}

func (r *msgRouterImple) cleaner() {
	defer helper.JustRecover()
	for {
		time.Sleep(time.Minute * 1)
		r.Lock()
		for RoomKey, adder := range r.mp {
			if adder.lastActiveTime < (helper.TimeNow() - 3600) {
				delete(r.mp, RoomKey)
			}
		}
		r.Unlock()
	}
}
