package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"ms/sun/servises/file_service_old"
	"ms/sun/shared/helper"
	"ms/sun/shared/x"
	"net/http"
	"strings"
	"time"
)

var cnt int = 1
var size int = 0

func main() {
	x.LogTableSqlReq.FileMsg = false
	x.LogTableSqlReq.FilePost = false

	Insert_many(10)
	file_service_old.Run()
	http.HandleFunc("/hi", func(writer http.ResponseWriter, r *http.Request) {
		writer.Write([]byte("hi"))
	})
	go func() {
		time.Sleep(time.Second)
		http.Get("http://localhost:1100/post_file/1518506476136010007_180.jpg")
	}()
	http.ListenAndServe(":1100", nil)
}
func Insert_many(num int) {
	for i := 0; i < num; i++ {
		insertMsg()
		insertPost()
		fmt.Println("cnt: ", i, " size:(mb)", size/1000000)
	}
}

func insertPost() {
	ext, bs, err := RandFile()
	if err != nil {
		return
	}
	cas := file_service_old.Row{
		Id:        helper.NextRowsSeqId(),
		Data:      bs,
		FileType:  1,
		Extension: ext,
		DataThumb: []byte{},
	}
	cnt++
	file_service_old.SavePostFile(cas)
}

func insertMsg() {
	ext, bs, err := RandFile()
	if err != nil {
		return
	}
	cas := file_service_old.Row{
		Id:        helper.NextRowsSeqId(),
		Data:      bs,
		FileType:  1,
		Extension: ext,
		DataThumb: []byte{},
	}
	cnt++
	file_service_old.SaveMsgFile(cas)
}

func RandFile() (ext string, bs []byte, err error) {
	const dir = `C:\Go\_gopath\src\ms\sun_old\upload\all_file_samples2`
	imageFiles, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	fn := dir + "/" + imageFiles[rand.Intn(len(imageFiles))].Name()

	bs, err = ioutil.ReadFile(fn)
	if err != nil {
		log.Fatal(err)
	}
	size += len(bs)

	ind := strings.LastIndex(fn, ".")
	if ind > 0 {
		ext = fn[ind:]
	}
	return ext, bs, nil

}
