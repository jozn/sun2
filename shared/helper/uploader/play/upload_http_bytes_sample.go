package main

import (
    "fmt"
    "os"
    "net/http"
    "ms/sun/shared/helper/uploader"
    "log"
    "github.com/jozn/protobuf/proto"
    "ms/sun/shared/x"
)

func main() {
    go func() {
        http.HandleFunc("/upload", func(writer http.ResponseWriter, r *http.Request) {
            fmt.Println("===================================")
            r.ParseMultipartForm(560000)
            fmt.Println(r.Body)
            fmt.Printf("%v \n",r.Body)
            fmt.Println(r.FormFile("file"))
            fmt.Println("===================================")
        })
        http.ListenAndServe("0.0.0.0:1111", nil)
    }()
    path, _ := os.Getwd()
    path += "/test.pdf"
    extraParams := map[string]string{
        "title":       "My Document",
        "author":      "Matt Aimonetti",
        "description": "A document with all the Go programming language secrets",
    }
    m:= &x.PB_MessageView{
        Text:"sdsgs",
        MessageFileId:64656,
    }
    pb ,err:= proto.Marshal(m)
    request, err := uploader.NewBytesUploadRequest("http://localhost:1111/upload", extraParams, "file25", pb)
    if err != nil {
        log.Fatal(err)
    }

    client := &http.Client{}
    resp, err := client.Do(request)
    if err != nil {
        log.Fatal(err)
    } else {
        var bodyContent []byte
        fmt.Println(resp.StatusCode)
        fmt.Println(resp.Header)
        resp.Body.Read(bodyContent)
        resp.Body.Close()
        fmt.Println(bodyContent)
    }
}

