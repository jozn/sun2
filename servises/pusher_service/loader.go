package pusher_service

import (
	"ms/sun/shared/base"
	"ms/sun/shared/x"
	"time"
)

var chatLastSelected = 0

func loadChatPushersLoop() {
	for {
		time.Sleep(time.Millisecond * 100)
		selector := x.NewPushChat_Selector()
		if chatLastSelected > 0 {
			selector.PushId_GT(chatLastSelected)
		}

		rows, err := selector.OrderBy_PushId_Asc().GetRows(base.DB)
		if err != nil || len(rows) == 0 {
			continue
		}
		chatLastSelected = rows[len(rows)-1].PushId
		for _, r := range rows {
			pushChatRouteChan <- r
		}
	}
}
