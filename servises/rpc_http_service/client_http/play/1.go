package main

import (
	"fmt"
	"ms/sun/servises/rpc_http_service"
	//"ms/sun_old/models/x"
	"ms/sun/servises/rpc_http_service/client_http"
	"ms/sun/shared/x"
	"net/http"
	"time"
)

func main() {

	go func() {
		http.HandleFunc("/rpc", rpc_http_service.ServeHttpRpc)
		http.ListenAndServe("localhost:9999", nil)
	}()
	i := 0
	fmt.Println("1")
	for {
		i++
		fmt.Println("--------------------------------------")
		d := &x.PB_OtherParam_Echo{
			Text: fmt.Sprintf("Hi From: %d ", i),
		}

		res, err := x.RPC_AllClinetsPlay.RPC_Other.Echo(d, client_http.FN)

		fmt.Printf("result of rpc final: %#+v  %s \n", res, err)
		time.Sleep(time.Millisecond * 300)

	}

}
