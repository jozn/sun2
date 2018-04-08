package main

import (
	"fmt"
	"ms/sun_old/base"
	"ms/sun/shared/x"
)

func main() {
	n := 0
	next := func() int {
		n++
		return n
	}
	work := func() {
        for  {
            //_, err := x.HomeFanoutByOrderId(base.DB, rand.Intn(40000))
            _, err := x.HomeFanoutByOrderId(base.DB, 1000)
            if n := next(); n%10000 == 0 {
                fmt.Println(n , err)
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
