package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	//"ms/sun/servises/file_service_old"
	"ms/sun/shared/helper"
	"net/http"
	"strings"
	"time"
    "ms/sun/shared/xc"
    "ms/sun/servises/file_service/file_store"
)

var cnt int = 1
var size int = 0

func main() {

	Insert_many(1000)
	//file_service_old.Run()
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
		insertPost()
		fmt.Println("cnt: ", i, " size:(mb)", size/1000000)
	}
}

func insertPost() {
	ext, bs, err := RandFile()
	if err != nil {
		return
	}
	ref := &xc.FileRef{
        Extension: ext,
        FileRefId: helper.NanoRowIdSeq(),
        Height: 0,
        Length: 0,
        Md5Key: "",
        MimeType: "",
        Name: "file_sampl."+ext,
        StorageUserId: 0,
        StoreType: file_store.SOTRE_TYPE_DIRECT_FILE,
        UserId: 5,
        Width: 0,
    }
	file_store.SavaFile(ref,bs)
	cnt++
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
