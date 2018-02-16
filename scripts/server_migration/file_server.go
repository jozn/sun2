package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	fmt.Println(os.Getwd())
	go func() {
        http.Handle("/dl", http.FileServer(http.Dir("/dl")))
        err := http.ListenAndServe("0.0.0.0:9000", nil)
        if err != nil {
            print(err)
        }
    }()

    go func() {
        http.Handle("/d2", http.FileServer(http.Dir(".")))
        err := http.ListenAndServe("0.0.0.0:7000", nil)
        if err != nil {
            print(err)
        }
    }()
	http.Handle("/", http.FileServer(http.Dir(".")))
	err := http.ListenAndServe("0.0.0.0:8080", nil)
	if err != nil {
	    print(err)
    }
}
