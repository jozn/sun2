package main

import (
	"github.com/peterbourgon/diskv"
	"ms/sun/shared/helper"
    "fmt"
    "os"
)

func main() {
	o := diskv.Options{
		BasePath: `C:\Go\_gopath\src\ms\sun\file_refs.db`,
	}
	db := diskv.New(o)
	for i := 0; i < 1000000; i++ {
		err := db.Write("key"+helper.IntToStr(i), []byte{byte(i)})
		helper.NoErr(err)
	}
	fmt.Println(os.Getwd())

}
