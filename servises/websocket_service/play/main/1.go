package main

import (
    "ms/sun2/servises/websocket_service"
    "net/http"
)

func main()  {
    http.HandleFunc("/ws", websocket_service.ServeHttpWsPB)
    http.HandleFunc("/rpc", websocket_service.ServeHttpRpc)
    http.ListenAndServe("0.0.0.0:9999", nil)
}
