package uploader

import (
    "os"
    "io/ioutil"
    "mime/multipart"
    "net/http"
    "bytes"
)

// Creates a new file upload http request with optional extra params
func NewFileUploadRequest(uri string, params map[string]string, paramName, path string) (*http.Request, error) {
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
func NewBytesUploadRequest(uri string, params map[string]string, paramName string, bytes1 []byte) (*http.Request, error) {
    body := &bytes.Buffer{}
    writer := multipart.NewWriter(body)
    part, err := writer.CreateFormFile("file","filename_sample.3gp")
    if err != nil {
        return nil, err
    }
    _,err = part.Write(bytes1)
    //fmt.Println("size", i,err)

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
