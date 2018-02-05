package main

import (
	"fmt"
	"ms/sun/base"
	"ms/sun/helper"
	"ms/sun2/shared/x"
	"time"
)

var i = 0
func main() {
	base.DefultConnectToMysql()

	go f1()
	go f1()
	go f1()
	go f1()

	time.Sleep(time.Hour)
}

func f1()  {
    for {
        i++
        x, err := x.NewPost_Selector().Limit(100).GetRows(base.DB)
        helper.NoErr(err)
        fmt.Println(len(x), i)
        time.Sleep(time.Millisecond * 10)
    }
}
