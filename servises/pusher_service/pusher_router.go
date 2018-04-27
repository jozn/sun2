package pusher_service

import "ms/sun/shared/x"

var pushChatRouteChan = make(chan *x.PushChat, 10000)

//var pushChatGroup = make(chan *x.PushChat, 10000)
//var pushChatPosts = make(chan *x.PushChat, 10000)

func routePushChatLoop() {
	for pc := range pushChatRouteChan {
		devs, ok := allPipesMap.getForUser(pc.ToUserId)
		if ok {
			pb := x.PB_Push{}
			devs.data <- pb
		}
	}
}
