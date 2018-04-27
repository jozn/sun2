package main

import (
	"ms/sun/shared/base"
	"ms/sun/shared/x"
	"time"
)

var i5 = 0

func main() {
	base.DefultConnectToMysql()

	go f5()


	time.Sleep(time.Hour)
}

func f5() {
	for {
		//x,_:=x.NewAction_Selector().GetRows(base.DB)
		x,_:=x.NewActionFanout_Selector().GetRows(base.DB)
		i5++

		if i5%1 == 0 {
			println(i5, len(x))
		}
	}
}
