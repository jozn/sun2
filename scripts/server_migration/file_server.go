package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	fmt.Println(os.Getwd())
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.ListenAndServe("0.0.0.0:8000", nil)
}
