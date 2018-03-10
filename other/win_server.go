package main

import (
    "fmt"
    "net/http"
    "os"
)

func main() {
    fmt.Println(os.Getwd())

    http.Handle("/e/", http.FileServer(http.Dir(`E:\`)))
    //http.Handle("/e/", http.FileServer(http.Dir(`.`)))
    http.Handle("/c/", http.FileServer(http.Dir("C:/")))
    http.Handle("/d", http.FileServer(http.Dir("D:/")))
    http.Handle("/", http.FileServer(http.Dir(`E:\`)))
    http.Handle("/h", http.FileServer(http.Dir("/")))
    err := http.ListenAndServe("0.0.0.0:2000", nil)
    if err != nil {
        print(err)
    }
}

