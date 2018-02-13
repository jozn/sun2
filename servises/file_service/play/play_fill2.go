package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"ms/sun/helper"
	"ms/sun2/servises/file_service"
	"net/http"
)

var cnt int = 1
var size int = 0

func main() {
	Insert_many(10)
	file_service.Run()
	http.HandleFunc("/hi", func(writer http.ResponseWriter, r *http.Request) {
		writer.Write([]byte("hi"))
	})
	http.ListenAndServe(":5151", nil)
}
func Insert_many(num int) {
	for i := 0; i < num; i++ {
		insert()
		fmt.Println("cnt: ", i, " size:(mb)", size/1000000)
	}
}

func insert() {
	_, bs, err := RandImage()
	if err != nil {
		return
	}
	cas := file_service.Row{
		Id:        helper.NextRowsSeqId(),
		Data:      bs,
		FileType:  1,
		Extension: ".jpg",
		DataThumb: []byte{},
	}
	cnt++
	file_service.SavePostFile(cas)
}

func RandImage() (fn string, bs []byte, err error) {
	const dir = `C:\Go\_gopath\src\ms\sun\upload\samples`
	imageFiles, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	fn = dir + "/" + imageFiles[rand.Intn(len(imageFiles))].Name()

	bs, err = ioutil.ReadFile(fn)
	if err != nil {
		log.Fatal(err)
	}
	size += len(bs)
	return fn[2:], bs, nil

}
