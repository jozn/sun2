package main

import (
	"math/rand"
	"ms/sun/shared/x"
	"ms/sun_old/base"
	"time"
)

var i56 = 0

func main() {
	base.DefultConnectToMysql()
	for i := 0; i < 6; i++ {

		go f6()
	}

	time.Sleep(time.Hour)
}

func f6() {
	for {
		arr := []int{}
		for i := 0; i < 50; i++ {
			arr = append(arr, rand.Intn(10000))
		}
		x, _ := x.NewHome_Selector().ForUserId_In(arr).Limit(150).OrderBy_Id_Desc().GetRows(base.DB)
		i56++

		if i56%100 == 0 {
			println(i56, len(x))
		}
	}
}
