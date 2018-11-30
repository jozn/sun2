package main

import (
	"fmt"
	"math/rand"
	"ms/sun_old/base"
	"ms/sun/shared/helper"
	"ms/sun/shared/x"
	"time"
)

var i4 = 0

func main() {
	base.DefultConnectToMysql()

	for i := 0; i < 6; i++ {
		go f4()
	}

	time.Sleep(time.Hour)
}

func f4() {
	for {
		arr := make([]x.HomeFanout, 0, 50)
		for i := 0; i < 100; i++ {
			p := x.HomeFanout{
				OrderId:        helper.NextRowsSeqId(),
				ForUserId: rand.Intn(10000)+1,
				PostId:    helper.NextRowsSeqId(),
			}
            i4++
            if i4%1000 == 0{

                fmt.Println(i4)
            }
			arr = append(arr, p)
		}


		x.MassReplace_HomeFanout(arr, base.DB)
	}

}
