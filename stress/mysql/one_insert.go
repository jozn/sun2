package main

import (
	"fmt"
	"ms/sun_old/base"
	"ms/sun/shared/x"
)

func main() {
    n:=0
    next := func() int {
        n++
        return n
    }
    work := func() {
        for i := 0; i < 100000; i++ {
            m := next()
            p := x.HomeFanout{
                OrderId:    m,
                ForUserId:  m,
                PostId:     m,
                PostUserId: m,
            }
            p.Save(base.DB)
            if m%10000 == 0 {
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
