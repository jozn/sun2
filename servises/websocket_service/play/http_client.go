package main

import (
	"io/ioutil"
	"ms/sun/shared/helper"
	"net/http"
)

func main() {

	f1()
}

func f1() {
	res, err := http.Get("https://golang.org/pkg/log/")
	//res, err := http.Post("https://golang.org/pkg/log/")
	helper.PertyPrint(res)
	helper.PertyPrint(err)
}
