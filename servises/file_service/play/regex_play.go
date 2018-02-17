package main

import (
	"fmt"
	"regexp"
)

func main() {
	url := "http://localhost:5151/msg_file/1518506476136010007.jpg"
	//url := "http://localhost:5151/msg_file/1518506476136010007_250.jpg"
	//url := "/msg_file/1518506476136010007_450.jpg"
	size := regexp.MustCompile(`/(([0-9]+)(_[[:alnum:]]+)?\.(\w+))`)
	//size := regexp.MustCompile(`/([0-9]+(_[[:alnum:]]+)?\.\w+)`)

	res := size.FindStringSubmatch(url)
	res2 := size.FindAllStringSubmatch(url,500)

	fmt.Println(res)
	fmt.Println(res2)
}

func f25() {
	url := "http://localhost:5151/msg_file/1518506476136010007_thumb.jpg"
	size := regexp.MustCompile(`.+_([[:alpha:]]+)\..+`)

	size.FindStringSubmatch(url)
}
