package main

import (
	"fmt"
	"ms/sun_old/base"
	"ms/sun/shared/x"
	"time"
    "math/rand"
    "ms/sun/shared/helper"
)

var i3 = 0

func main() {
	base.DefultConnectToMysql()

	go f3()
	go f3()
	go f3()
	go f3()
	go f3()
	go f3()

	time.Sleep(time.Hour)
}

func f3() {
	for {
		arr := make([]x.ChatLastMessage, 0, 50)
		for i := 0; i < 100; i++ {
			p := x.ChatLastMessage{
                ChatKey:  helper.RandAlphaNumbericString(15),
                ForUserId: rand.Intn(10000000),
                LastMsgPb: []byte{},
                LastMsgJson: helper.RandString(24),
			}

			arr = append(arr, p)
		}
		i3++
		fmt.Println(i3)
		x.MassReplace_ChatLastMessage(arr, base.DB)
	}

}
