package main

import (
    "net/http"
    "time"
    "ms/sun/servises/file_service"
)

func main()  {
    file_service.Run()
    defer file_service.DeferCleanUp()
    http.HandleFunc("/hi", func(writer http.ResponseWriter, r *http.Request) {
        writer.Write([]byte("hi ========="))
    })
    go func() {
        time.Sleep(time.Second)
        http.Get("http://localhost:1100/post_file/1518506476136010007_180.jpg")
    }()
    http.ListenAndServe(":1100", nil)

    time.Sleep(time.Hour)
}