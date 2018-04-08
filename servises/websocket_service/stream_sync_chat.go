package websocket_service

import (
	"ms/sun_old/base"
	"ms/sun_old/helper"
	"ms/sun/servises/chat_service"
	"ms/sun/shared/x"
	"time"
)

var syncChatPipePusher = make(chan *x.ChatSync, 100000)
var lastSyncLoaded = helper.NextRowsSeqId()

func syncChaFetcher() {
    defer helper.JustRecover()
	for {
		time.Sleep(time.Millisecond * 50)
		syncs, err := x.NewChatSync_Selector().SyncId_GE(lastSyncLoaded).OrderBy_SyncId_Asc().GetRows(base.DB)
		if err == nil || len(syncs) == 0 {
			continue
		}
		lastSyncLoaded = syncs[0].SyncId

		for _, sync := range syncs {
			syncChatPipePusher <- sync
		}
	}
}

//can causes panic
func syncChatPusherToDevicesFramer() {
    defer helper.JustRecover()
	const siz = 100000
	arr := make([]*x.ChatSync, 0, siz)
	cnt := 0

	ticker := time.NewTicker(10 * time.Millisecond)
	for {
		select {
		case m := <-syncChatPipePusher:
			arr = append(arr, m)

		case <-ticker.C:
			if len(arr) > 0 {
				cnt++
				logChatSync.Printf("batch of syncChatPusherToDevicesFramer - cnt:%d - len:%d \n", cnt, len(arr))
				pre := arr
				arr = make([]*x.ChatSync, 0, siz)
				go func(syncs []*x.ChatSync) {
					defer helper.JustRecover()
					mp := make(map[int][]*x.ChatSync)
					for _, sync := range syncs {
						mp[sync.ToUserId] = append(mp[sync.ToUserId], sync)
					}

					for uid, syncs2 := range mp {
						if AllPipesMap.IsPipeOpen(uid) {
							up := chat_service.ToUpdates(uid, syncs2)
							cmd := newPB_CommandToClient_WithData("Updates", up)
							AllPipesMap.SendToUser(uid, cmd)
						}
					}
				}(pre)
			}
		}
	}
}
