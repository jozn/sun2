package main

import (
    "fmt"
    "ms/sun/servises/rpc_http_service"
    //"ms/sun_old/models/x"
    "ms/sun/shared/x"
    "net/http"
    "ms/sun/servises/rpc_http_service/client_http"
    "time"
)

func main() {

    go func() {
        http.HandleFunc("/rpc", rpc_http_service.ServeHttpRpc)
        http.ListenAndServe("localhost:9999", nil)
    }()

    fmt.Println("1")
    for  {
        fmt.Println("--------------------------------------")
        d := &x.PB_OtherParam_Echo{
            Text: "salam",
        }

        res, err := x.RPC_AllClinetsPlay.RPC_Other.Echo(d, client_http.FN)

        fmt.Printf("result of rpc final: %#+v  %s \n",res, err)
        time.Sleep(time.Second)
    }



}
