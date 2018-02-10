package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	cnt := 0
	mp := make(map[int]chan int)
	add := make(chan chan int)
	for i := 0; i < 100000; i++ {
		go func(i int) {
			ch := make(chan int)
			add <- ch
			j := 0
			for {
				j++
				<-ch
				cnt++
				if cnt%100000 == 0{
                    fmt.Println("existing: ", j, cnt)
                }
			}
		}(i)
	}
	go func() {
		for i := 0; i < 100000; i++ {
			c := <-add
			mp[i] = c
		}
        for i:=0;i<100000;i++{
            go func() {
                for {
                    time.Sleep(time.Millisecond * 100)
                    mp[rand.Intn(100000)] <- 1
                }
            }()
        }
	}()
	time.Sleep(time.Hour * 4)
}
