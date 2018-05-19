package main

import (
	"fmt"
	"ms/sun/shared/conns"
	"ms/sun/shared/helper"
	"ms/sun/shared/x"
	"sync/atomic"
	"time"
    "math/rand"
)

func main() {
	x.LogTableSqlReq.PostCdb = false
	i := int64(0)
	fn := func() {
		for {
			atomic.AddInt64(&i, 1)
			rows, err := x.NewPostCdb_Selector().PostId_LT(helper.NanoRowIdAtSec(int(time.Now().Unix()) - (rand.Intn(1000)+100))).Limit(1).GetRows(conns.DB_PG)

			if i% 1000 == 0{
                fmt.Println(len(rows), err)
            }
		}
	}

	go func() {
		for {
			time.Sleep(time.Second)
			fmt.Println(i)
		}
	}()

	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	fn()

}
