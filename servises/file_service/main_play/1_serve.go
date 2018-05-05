package main

import (
    "ms/sun/servises/file_service"
    "net/http"
    "time"
)

func main()  {
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