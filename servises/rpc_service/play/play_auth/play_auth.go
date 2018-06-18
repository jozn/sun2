package main

import (
	"fmt"
	"ms/sun/servises/rpc_http_service"
	"ms/sun/servises/rpc_http_service/client_http"
	"ms/sun/shared/helper"
	"ms/sun/shared/x"
	"net/http"
)

func main() {
	go func() {
		http.HandleFunc("/rpc", rpc_http_service.ServeHttpRpc)
		http.ListenAndServe("localhost:9999", nil)
	}()
	i := 0
	fmt.Println("1")
	i++
	fmt.Println("------------- send sms -------------------")

	//sendSms()
	//confrinmSms()
	register()
}

func sendSms() {
	fmt.Println("============ sendSms ===============")
	hash := helper.RandAlphaNumbericString(32)
	d := &x.RPC_Auth_Types_SendConfirmCode_Param{
		Phone:       "989164150318",
		Hash:        hash,
		CountryCode: "98",
	}

	res, err := x.RPC_AllClinetsPlay.RPC_Auth.SendConfirmCode(d, client_http.FN)

	fmt.Printf("result of rpc final: %#+v  %s \n", res, err)
}

func confrinmSms() {
	fmt.Println("============ confrinmSms ===============")

	hash := "1nvZhD6QvOCR62NnQ53L8vtlG8DmIfOi"
	d := &x.RPC_Auth_Types_ConfirmCode_Param{
		Phone: "989164150318",
		Hash:  hash,
		Code:  1887,
	}
	res, err := x.RPC_AllClinetsPlay.RPC_Auth.ConfirmCode(d, client_http.FN)

	fmt.Printf("result of rpc final: %#+v  %s \n", res, err)
}

func register() {
    fmt.Println("============ confrinmSms ===============")

    hash := "1nvZhD6QvOCR62NnQ53L8vtlG8DmIfOi"
    d := &x.RPC_Auth_Types_SingUp_Param{
        Phone: "989164150318",
        Hash:  hash,
        FirstName:  "حمید",
    }
    res, err := x.RPC_AllClinetsPlay.RPC_Auth.SingUp(d, client_http.FN)

    fmt.Printf("result of rpc final: %#+v  %s \n", res, err)
}
