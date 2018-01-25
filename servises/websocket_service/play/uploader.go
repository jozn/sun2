package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

// Creates a new file upload http request with optional extra params
func newfileUploadRequest(uri string, params map[string]string, paramName, path string) (*http.Request, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	fileContents, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	fi, err := file.Stat()
	if err != nil {
		return nil, err
	}
	file.Close()

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(paramName, fi.Name())
	if err != nil {
		return nil, err
	}
	part.Write(fileContents)

	for key, val := range params {
		_ = writer.WriteField(key, val)
	}
	err = writer.Close()
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", uri, body)
	request.Header.Add("Content-Type", writer.FormDataContentType())
	return request, err
}

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
	request, err := newfileUploadRequest("http://localhost:1111/upload", extraParams, "file", `C:\Go\_gopath\src\ms\sun\upload\1.mp4`)
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
