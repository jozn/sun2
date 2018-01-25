package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"log"
	"ms/sun2/shared/helper/uploader"
	"ms/sun2/shared/x"
	"net/http"
	"time"
)
var cnt = 0
func main() {
	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 10000; j++ {
				sendFast()
				cnt++
				if cnt% 1000 ==0 {
                    fmt.Println(cnt)
                }
			}
		}()
	}
	time.Sleep(time.Second * 100)
}

func send() {
	p := &x.PB_OtherParam_Echo{Text: "samlam"}
	b, err := proto.Marshal(p)
	m := &x.PB_CommandToServer{
		Command:      "RPC_Other.Echo",
		Data:         b,
		ClientCallId: 16565,
	}
	pb, err := proto.Marshal(m)
	request, err := uploader.NewBytesUploadRequest("http://localhost:9999/rpc", nil, "file25", pb)
	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{}
	resp, err := client.Do(request)
	fmt.Println("=================================")
	if err != nil {
		log.Fatal(err)
	} else {
		var bodyContent []byte
		fmt.Println(resp.StatusCode)
		fmt.Println(resp.Header)
		//resp.Body.Read(bodyContent)
		bts, err := ioutil.ReadAll(resp.Body)
		cmd := &x.PB_ResponseToClient{}
		err = proto.Unmarshal(bts, cmd)
		//resp.Body.Close()
		fmt.Println(bodyContent)
		fmt.Println("size: ", len(bts))
		fmt.Println("err: ", err)
		fmt.Println("cmd: ", cmd)
	}
}

var client = &http.Client{}
func sendFast() {
	p := &x.PB_OtherParam_Echo{Text: "samlam"}
	b, err := proto.Marshal(p)
	m := &x.PB_CommandToServer{
		Command:      "RPC_Other.Echo",
		Data:         b,
		ClientCallId: 16565,
	}
	pb, err := proto.Marshal(m)
	request, err := uploader.NewBytesUploadRequest("http://localhost:9999/rpc", nil, "file25", pb)
	if err != nil {
		log.Fatal(err)
	}

    resp, err := client.Do(request)
    if cnt% 1000 ==0 {
        if err != nil {
            log.Fatal(err)
        } else {
            var bodyContent []byte
            fmt.Println(resp.StatusCode)
            fmt.Println(resp.Header)
            //resp.Body.Read(bodyContent)
            bts, err := ioutil.ReadAll(resp.Body)
            cmd := &x.PB_ResponseToClient{}
            err = proto.Unmarshal(bts, cmd)
            //resp.Body.Close()
            fmt.Println(bodyContent)
            fmt.Println("size: ", len(bts))
            fmt.Println("err: ", err)
            fmt.Println("cmd: ", cmd)
        }
    }

}
