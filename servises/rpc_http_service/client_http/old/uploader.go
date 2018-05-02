package main

import (
    "net/http"
    "os"
    "io/ioutil"
    "bytes"
    "mime/multipart"
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

// Creates a new file upload http request with optional extra params
func newBytesUploadRequest2(uri string, params map[string]string, btsArr []byte) (*http.Request, error) {

    body := new(bytes.Buffer)
    writer := multipart.NewWriter(body)
    part, err := writer.CreateFormField("file")//, fi.Name())
    if err != nil {
        return nil, err
    }
    part.Write(btsArr)

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


