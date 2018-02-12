package message_adder

import (
	"ms/sun2/shared/x"
	"sync"
)

//this one must load from mysql, kafka or for now just simple loading on single node
var MessageAdderChanel = make(chan *x.PB_MessageView) //messges must have a right messageFileId
var msgRouter = msgRouterImple{
    mp: make( map[string]*chatDirect,10000),
}

func init() {
    msgRouter.handle()
}

type msgRouterImple struct {
	mp map[string]*chatDirect
	sync.RWMutex
}

//there is even no need for locking
func (r *msgRouterImple) handle() {
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
