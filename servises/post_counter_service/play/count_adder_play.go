package main

import (
	"math/rand"
	"ms/sun/servises/post_counter_service"
	"time"
)

func main() {
    //fmt.Println("************************88 ")
	for i := 0; i < 8; i++ {
		go func() {
            //fmt.Println("+++++++++++++++++++= ")
			for {
                //fmt.Println("go ++++++++++++ AddSeen: ")
                post_counter_service.AddSeen(rand.Intn(1000) + 1)
				time.Sleep(time.Millisecond * 1)
			}
		}()
	}
	time.Sleep(time.Hour)
}
