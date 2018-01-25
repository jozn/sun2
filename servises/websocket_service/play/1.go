package main

import (
    "net/http"
    "ms/sun2/servises/websocket_service"
)

func main()  {

    http.HandleFunc("/ws", websocket_service.ServeHttpWsPB)
    http.HandleFunc("/", websocket_service.ServeHttpRpc)
    http.ListenAndServe("localhost:9999",nil)
}

func client()  {

}