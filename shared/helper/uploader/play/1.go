package main

import (
    "fmt"
    "os"
    "net/http"
    "ms/sun2/shared/helper/uploader"
    "log"
)

func main() {
    go func() {
        http.HandleFunc("/upload", func(writer http.ResponseWriter, r *http.Request) {
            fmt.Println("===================================")
            r.ParseMultipartForm(560000)
            fmt.Println(r.Body)
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
    request, err := uploader.NewFileUploadRequest("http://localhost:1111/upload", extraParams, "file", `C:\Go\_gopath\src\ms\sun\upload\1.mp4`)
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

