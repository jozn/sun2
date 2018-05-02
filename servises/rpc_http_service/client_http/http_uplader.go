package client_http

import (
    "net/http"
    "bytes"
    "mime/multipart"
)

func newBytesUploadRequest(uri string, params map[string]string, btsArr []byte) (*http.Request, error) {

    body := new(bytes.Buffer)
    writer := multipart.NewWriter(body)
    part, err := writer.CreateFormFile("file", "file_something") //, fi.Name())
    //part, err := writer.CreateFormField("file")//, fi.Name())
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
