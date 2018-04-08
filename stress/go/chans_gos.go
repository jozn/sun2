package main

import (
	"fmt"
	"time"
)

func main() {
	const NUM = 10000
	f := func(id int, in, out chan int) {
		go func() {
			for {
				//fmt.Println("=",id)
				i := <-in
				//fmt.Println("recived: ", id, " i: ", i)
				out <- i * 2
			}
		}()
	}

	in := make(chan int,100)
	inFist := in
	out := make(chan int,100)
	for i := 0; i < NUM; i++ {
		f(i, in, out)
		in = out
		out = make(chan int,100)
	}
    go func() {
        for {
            <- in
        }
    }()
	cnt := 0
	go func() {
		for {
			cnt++
			inFist <- 1
			if cnt % 10000 == 0 {
                fmt.Println("cnt: ", cnt, " all: ", NUM*cnt)
                //time.Sleep(time.Millisecond)

            }
			//time.Sleep(time.Millisecond)
		}
	}()

	//fmt.Println("last: ", <-out)

	time.Sleep(time.Second*400)

}
