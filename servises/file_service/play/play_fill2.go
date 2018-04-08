package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"ms/sun_old/helper"
	"ms/sun/servises/file_service"
	"net/http"
    "time"
)

var cnt int = 1
var size int = 0

func main() {
	Insert_many(0)
	file_service.Run()
	http.HandleFunc("/hi", func(writer http.ResponseWriter, r *http.Request) {
		writer.Write([]byte("hi"))
	})
	go func() {
	    time.Sleep(time.Second)
	    http.Get("http://localhost:5151/post_file/1518506476136010007_180.jpg")
    }()
	http.ListenAndServe(":5151", nil)
}
func Insert_many(num int) {
	for i := 0; i < num; i++ {
		insertMsg()
		insertPost()
		fmt.Println("cnt: ", i, " size:(mb)", size/1000000)
	}
}

func insertPost() {
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

func insertMsg() {
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
    file_service.SaveMsgFile(cas)
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
