package pusher_service

import (
    "ms/sun/shared/x"
    "ms/sun/shared/helper"
    "fmt"
)

var pushChatRouteChan = make(chan *x.PushChat,10000)

//var pushChatGroup = make(chan *x.PushChat, 10000)
//var pushChatPosts = make(chan *x.PushChat, 10000)
var m = 1
func routePushChatLoop() {
    defer helper.JustRecover()
	for pc := range pushChatRouteChan {
	    //fmt.Println("ooooooooooo: ",len(allPipesMap.mp))
	    //helper.PertyPrint2(allPipesMap)
		devs, ok := allPipesMap.getForUser(pc.ToUserId)
		if ok {
            //fmt.Println("mmm: ",pc)

            pb := x.PB_Push{
			    LastPushId:int64(m),
			    LastChatPushId:int64(pc.ToUserId),
            }
            m++
			devs.data <- pb
			_ = pb
			_ = devs

			if m%1000 == 0 {
			    fmt.Println(m)
            }
		}
	}
}
