package main

import (
    "net/http"

    "ms/sun/servises/rpc_http_service"
)

func main()  {

    http.HandleFunc("/rpc", rpc_http_service.ServeHttpRpc)
    http.ListenAndServe("localhost:9999",nil)
}

func client()  {

}
