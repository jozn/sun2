package main

import (
	"math/rand"
	"ms/sun_old/base"
	"ms/sun/shared/x"
	"time"
)

var i5 = 0

func main() {
	base.DefultConnectToMysql()

	go f5()
	go f5()
	go f5()
	go f5()
	go f5()
	go f5()

	time.Sleep(time.Hour)
}

func f5() {
	for {
		x,_:=x.NewHome_Selector().ForUserId_Eq(rand.Intn(10000)).GetRows(base.DB)
		i5++

		if i5%100 == 0 {
			println(i5, len(x))
		}
	}
}
