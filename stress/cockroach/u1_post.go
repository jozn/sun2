package main

import (
	"fmt"
	"ms/sun/shared/dbs"
	"ms/sun/shared/x"
	"sync/atomic"
	"time"
)

func main() {
	x.LogTableSqlReq.PostCdb = false
	i := int64(0)
	fn := func() {
        //time.Sleep(time.Millisecond  * int64(rand.Intn(10000)))
		for {
			atomic.AddInt64(&i, 1)
			//rows, err := x.PostCdbByPostId(conns.DB_PG,1526682757615011071)
			rows, err := x.NewPostCdb_Updater().UserId(int(i)).PostId_Eq(1526682757615011071).Update(dbs.DB_PG)

			if i% 10 == 0{
                fmt.Println((rows), err,"*")
            }
            time.Sleep(time.Millisecond)
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
