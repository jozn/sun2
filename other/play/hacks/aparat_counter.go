package main

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"
	"time"
    "ms/sun/shared/helper"
)

func main() {
	const NUM = 1000
	const BUFF = 30

	var r404 = 0//most of them are private, few are deleted - nazikia showes private as deleted
	var r200 = 0
	var r500 = 0

	ch := make(chan bool, 2)
	buff := make(chan bool, BUFF)
	for i := 0; i < BUFF; i++ {
		buff <- true
	}

	for i := 0; i < NUM; i++ {
		go func(i int) {
			defer func() {
				if r := recover(); r != nil {
					ch <- true
					buff <- true
					fmt.Println("ERR")
				}
			}()

			<-buff

			ex := helper.RandAlphaNumbericString(5)
			url := fmt.Sprintf("https://www.aparat.com/v/%s", ex)
			res, err := http.Get(url)
			if err != nil {
				fmt.Println(err)
			}
			buf := new(bytes.Buffer)
			buf.ReadFrom(res.Body)
			s := buf.String()
			if strings.Contains(s, "Not Found") {
				r404 += 1
			} else {
				switch res.StatusCode {
				case 200:
					r200 += 1
				case 404:
					r404 += 1
				default:
					r500 += 1
				}
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
		fmt.Println(500, r500)
	}()

	fmt.Println("==============================")
	fmt.Println(404, r404)
	fmt.Println(200, r200)
	fmt.Println(500, r500)

	time.Sleep(time.Hour)

}
