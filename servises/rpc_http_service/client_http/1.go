package client_http

import (
	"fmt"
	"ms/sun/servises/rpc_http_service"
	//"ms/sun_old/models/x"
	"ms/sun/shared/x"
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

	res, err := x.RPC_AllClinetsPlay.RPC_Other.Echo(d, FN)

	fmt.Println(res, err)

	time.Sleep(time.Minute)
}
