package main

import (
	"fmt"
	"math/rand"
	"ms/sun_old/base"
	"ms/sun_old/helper"
	"ms/sun/shared/x"
	"time"
)

var i4 = 0

func main() {
	base.DefultConnectToMysql()

	go f4()
	go f4()
	go f4()
	go f4()
	go f4()
	go f4()

	time.Sleep(time.Hour)
}

func f4() {
	for {
		arr := make([]x.Home, 0, 50)
		for i := 0; i < 1000; i++ {
			p := x.Home{
				Id:        helper.NextRowsSeqId(),
				ForUserId: rand.Intn(10000)+1,
				PostId:    helper.NextRowsSeqId(),
			}

			arr = append(arr, p)
		}
        i4++
		if i4%1 == 0{

            fmt.Println(i4)
        }
		x.MassReplace_Home(arr, base.DB)
	}

}
