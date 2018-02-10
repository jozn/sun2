package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	const NUM = 1000
	const BUFF = 20

	var r404 = 0
	var r200 = 0

	ch := make(chan bool, 2)
	buff := make(chan bool, BUFF)
	for i := 0; i < BUFF; i++ {
		buff <- true
	}

	for i := 0; i < NUM; i++ {
		cnt := 15000000
		go func(i int) {
			defer func() {
				if r := recover(); r != nil {
					ch <- true
					buff <- true
					fmt.Println("ERR")
				}
			}()

			<-buff

			url := fmt.Sprintf("http://wisgoon.com/pin/%d/", i+cnt)
			res, err := http.Get(url)
			if err != nil {
				fmt.Println(err)
			}
			switch res.StatusCode {
			case 200:
				r200 += 1
			case 404:
				r404 += 1
			}
			fmt.Println(res.StatusCode, url)
			ch <- true
			buff <- true
		}(i)
	}
	for i := 0; i < NUM; i++ {
		<-ch
	}
	go func() {
		time.Sleep(time.Second * 4)
		fmt.Println("==============================")
		fmt.Println(404, r404)
		fmt.Println(200, r200)
	}()

	fmt.Println("==============================")
	fmt.Println(404, r404)
	fmt.Println(200, r200)

	time.Sleep(time.Hour)

}
