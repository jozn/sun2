package client_http

import (
    "net/http"
    "fmt"
    "io/ioutil"
    "ms/sun/shared/x"
    "github.com/golang/protobuf/proto"
    "log"
)

var cnt = 0
var FN = func(cmdSre string, pbIn,pbOut proto.Message) ( error) {
    cnt++
    /*pbMsg, ok := pbIn.(*proto.Message)
    if !ok {
        return  errors.New("cannot convert to pbMsg")
    }*/

    bts, err := proto.Marshal(pbIn)
    if err != nil {
        return  err
    }

    cmd := &x.PB_CommandToServer{
        ClientCallId: int64(cnt),
        Command:      cmdSre,
        Data:         bts,
    }

    upBts, err := proto.Marshal(cmd)
    if err != nil {
        return  err
    }

    mp := make(map[string]string)
    req, err := newBytesUploadRequest("http://localhost:9999/rpc", mp, upBts)
    if err != nil {
        //return nil, err
        log.Fatalln("in req : ", err)
    }
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        //return nil, err
        log.Println("after do: ", err)
    }

    fmt.Println(resp.StatusCode)
    fmt.Println(resp.Header)
    bodyVts, _ := ioutil.ReadAll(resp.Body)
    resp.Body.Close()
    fmt.Printf("body: %s \n", bodyVts)
    pbRes := &x.PB_ResponseToClient{}
    err = proto.Unmarshal(bodyVts, pbRes)
    fmt.Printf("body Cmd: %#+v  %s\n", pbIn, err)

    err = proto.Unmarshal(pbRes.Data, pbOut)

    fmt.Println("end")
    return nil
}


