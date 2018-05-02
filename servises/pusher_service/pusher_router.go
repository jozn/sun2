package pusher_service

import (
	"ms/sun/shared/helper"
	"ms/sun/shared/x"
    "ms/sun/servises/pusher_service/push_translator"
)

var pushChatRouteChan = make(chan *x.PushChat, 10000)

//var pushChatGroup = make(chan *x.PushChat, 10000)
//var pushChatPosts = make(chan *x.PushChat, 10000)


func routePushChatLoop() {
	defer helper.JustRecover()
	for pc := range pushChatRouteChan {
		//fmt.Println("ooooooooooo: ",len(allPipesMap.mp))
		//helper.PertyPrint2(allPipesMap)
		devs, ok := allPipesMap.getForUser(pc.ToUserId)
		if ok {
			//fmt.Println("mmm: ",pc)
			Monitor.TotalRowsProceedWiltActiveUser += 1
			Monitor.TotalChatRowsProceedWiltActiveUser += 1
			/*pb := x.PB_Push{
				LastPushId:     int64(m),
				LastChatPushId: int64(pc.ToUserId),
			}*/
			pb := push_translator.ChatPushToPbChat(pc)

			devs.data <- pb
			_ = pb
			_ = devs

		}
	}
}
