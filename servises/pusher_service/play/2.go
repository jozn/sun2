package main

import (
	"ms/sun/servises/pusher_service"
	"ms/sun/shared/base"
	"ms/sun/shared/helper"
	"ms/sun/shared/x"
	"os"
	"time"
    "math/rand"
)

func main() {
	i := 0

	go func() {
		for i := 0; i < 100; i += 2 {
			pusher_service.ServeMockStream(i, os.Stderr)
		}
	}()

	go func() {
		for {
			i++
			//time.Sleep(time.Millisecond * 50)
			arr := []x.PushChat{}
			for i := 0; i < 100; i++ {
				p := x.PushChat{
					PushId:            helper.NanoRowIdSeq(),
					ToUserId:          rand.Intn(120),
					PushTypeId:        10,
					RoomKey:           "sd",
					ChatKey:           "gs",
					Seq:               i,
					UnseenCount:       0,
					FromHighMessageId: 0,
					ToLowMessageId:    0,
					MessageId:         2555,
					MessageFileId:     0,
					MessagePb:         []byte{},
					MessageJson:       "",
					CreatedTime:       helper.TimeNow(),
				}
				arr = append(arr, p)
			}

			x.MassReplace_PushChat(arr, base.DB)
		}
	}()
	time.Sleep(time.Hour)
}
