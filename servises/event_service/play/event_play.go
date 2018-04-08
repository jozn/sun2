package main

import (
	"fmt"
	"math/rand"
	"ms/sun_old/helper"
	"ms/sun/servises/event_service"
	"ms/sun/shared/x"
	"time"
)

func main() {
    //x.LogTableSqlReq.Event = false
	go func() {
		param := event_service.SubParam{
			Liked_Post_Event: true,
		}
		sub := event_service.NewSub(param)
		for evn := range sub.Liked_Post_Event {
			fmt.Println("event,", evn)
		}
	}()
	go func() {
        for {
            time.Sleep(time.Millisecond * 100)

            evnt := x.Event{
                EventId:    helper.NextRowsSeqId(),
                EventType:  0,
                ByUserId:   rand.Intn(80) + 1,
                PeerUserId: 0,
                PostId:     12,
                CommentId:  0,
                ActionId:   0,
            }
            event_service.SaveEvent(event_service.LIKED_POST_EVENT, evnt)
        }
    }()
    time.Sleep(time.Hour)
}
