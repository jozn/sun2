package main

import (
	//"github.com/golang/protobuf/proto"
	"bytes"
	"fmt"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"log"
	"math/rand"
	"mime/multipart"
	"ms/sun/servises/rpc_http_service"
	"ms/sun_old/models/x"
	"net/http"
	"time"
)

func main() {

	go func() {
		http.HandleFunc("/rpc", rpc_http_service.ServeHttpRpc)
		http.ListenAndServe("localhost:9999", nil)
	}()

	fmt.Println("1")
	d := &x.PB_OtherParam_Echo{
		Text: "salam",
	}

	bts, _ := proto.Marshal(d)

	cmd := &x.PB_CommandToServer{
		ClientCallId: int64(rand.Intn(1000)),
		Command:      "RPC_Other.Echo",
		Data:         bts,
	}

	upBts, _ := proto.Marshal(cmd)

	mp := make(map[string]string)
	req, err := newBytesUploadRequest("http://localhost:9999/rpc", mp, upBts)
	if err != nil {
		log.Fatalln("in req : ", err)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("after do: ", err)
	} else {
		//var bodyContent []byte
		//bodyContent := make([]byte,500)
		fmt.Println(resp.StatusCode)
		fmt.Println(resp.Header)
		//resp.Body.Read(bodyContent)
		bodyVts, _ := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		fmt.Printf("body: %s \n", bodyVts)
		pb := &x.PB_ResponseToClient{}
		err = proto.Unmarshal(bodyVts, pb)
		fmt.Printf("body Cmd: %#+v  %s\n", pb, err)

	}

	fmt.Println("end")

	time.Sleep(time.Minute)
}

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

/*/////////////////////////////
type RPCClientHandler func(cmd string, pb interface{}) interface{}

var RPC_AllClinetPlay = struct {
	RPC_Auth RPC_Auth_Client
	//RPC_Chat   RPC_Chat
	//  RPC_Other  RPC_Other
	//  RPC_Page   RPC_Page
	//  RPC_Search RPC_Search
	//  RPC_Social RPC_Social
	//  RPC_User   RPC_User
}{
	RPC_Auth: RPC_Auth_Client(0),
}

type RPC_Auth_Client int

func (RPC_Auth_Client) CheckPhone(param *x.PB_UserParam_CheckUserName2, fn RPCClientHandler) (res *x.PB_UserResponse_CheckUserName2, err error) {
	r := fn("RPC_Auth.CheckPhone", param)
	res, ok := r.(*x.PB_UserResponse_CheckUserName2)
	if ok {
		return res, nil
	}
	return nil, errors.New("err in converitn in CheckPhone")
}
*/
/*
type RPC_Auth interface {
    CheckPhone(param *PB_UserParam_CheckUserName2) (res PB_UserResponse_CheckUserName2, err error)
    SendCode(param *PB_UserParam_CheckUserName2, userParam RPC_UserParam) (res PB_UserResponse_CheckUserName2, err error)
    SendCodeToSms(param *PB_UserParam_CheckUserName2, userParam RPC_UserParam) (res PB_UserResponse_CheckUserName2, err error)
    SendCodeToTelgram(param *PB_UserParam_CheckUserName2, userParam RPC_UserParam) (res PB_UserResponse_CheckUserName2, err error)
    SingUp(param *PB_UserParam_CheckUserName2, userParam RPC_UserParam) (res PB_UserResponse_CheckUserName2, err error)
    SingIn(param *PB_UserParam_CheckUserName2, userParam RPC_UserParam) (res PB_UserResponse_CheckUserName2, err error)
    LogOut(param *PB_UserParam_CheckUserName2, userParam RPC_UserParam) (res PB_UserResponse_CheckUserName2, err error)
}*/
