package pusher_service

import (
	"ms/sun/shared/base"
	"ms/sun/shared/x"
	"time"
    "ms/sun/shared/config"
    "ms/sun/shared/helper"
)

var chatLastSelected = 0

func loadChatPushersLoop() {
    defer helper.JustRecover()
	for {
		time.Sleep(time.Millisecond * 500)
		selector := x.NewPushChat_Selector()
		if chatLastSelected > 0 {
			selector.PushId_GT(chatLastSelected)
		}

		rows, err := selector.OrderBy_PushId_Asc().GetRows(base.DB)
        if config.IS_DEBUG {
            logPushLoader.Printf("loaded from push_chat len: %d \n", len(rows))
        }

		if err != nil || len(rows) == 0 {
			continue
		}

		chatLastSelected = rows[len(rows)-1].PushId
		for _, r := range rows {
		    //fmt.Println("ppp: ", r)
			pushChatRouteChan <- r
		}
	}
}
