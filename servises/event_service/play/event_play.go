package main

import (
	"fmt"
	"math/rand"
	"ms/sun/helper"
	"ms/sun2/servises/event_service"
	"ms/sun2/shared/x"
	"time"
)

func main() {
	go func() {
		param := event_service.SubParam{
			LIKED_POST_EVENT: true,
		}
		sub := event_service.NewSub(param)
		for evn := range sub.LIKED_POST_EVENT {
			fmt.Println("event,", evn)
		}
	}()
	go func() {
        for {
            time.Sleep(time.Millisecond * 500)

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