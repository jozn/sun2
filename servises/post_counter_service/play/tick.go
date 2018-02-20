package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(time.Millisecond * 500)
	for {
		select {
		case t:= <-tick:
			fmt.Println("tick",t)
		}
	}
}
