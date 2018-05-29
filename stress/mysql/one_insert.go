package main

import (
	"fmt"
	//"ms/sun_old/base"
	"ms/sun/shared/x"
    "ms/sun/shared/base"
)

func main() {
    //x.LogTableSqlReq
    n:=0
    next := func() int {
        n++
        return n
    }
    i := 0
    work := func() {
        for ; i < 1000000; i++ {
            m := next()
            p := x.HomeFanout{
                OrderId:    m,
                ForUserId:  m,
                PostId:     m,
                PostUserId: m,
            }
            p.Save(base.DB)
            if m%1000 == 0 {
                fmt.Println("i: ", m)
            }
        }
    }

    go work()
    go work()
    go work()
    go work()
    go work()
    go work()
    go work()
    go work()
    go work()
    go work()
    go work()
    go work()
    work()

}
