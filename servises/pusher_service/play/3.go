package main

import (
	"ms/sun/servises/pusher_service"
	"time"
    "runtime"
    "ms/sun/shared/x"
    "fmt"
    "ms/sun/shared/helper"
)

func main() {
	//i := 0

	go func() {
		for i := 1; i < 100; i += 5 {
			//pusher_service.ServeMockStream(i, os.Stderr)
			pusher_service.ServeMockStream(i, nil)
		}
	}()

    go func() {
        for {
            time.Sleep(time.Hour * 10)
            fmt.Println("Gc + len rows" ,len(x.RowCache.Items()))
            x.RowCache.Flush()
            runtime.GC()
            m := &runtime.MemStats{}
            runtime.ReadMemStats(m)
            helper.PertyPrintNew(m)
        }
    }()



	go func() {
		/*for {
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
		}*/
	}()
	time.Sleep(time.Hour)
}
