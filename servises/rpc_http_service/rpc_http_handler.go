package rpc_http_service

import (
    "ms/sun/shared/config"
    "log"
    "net/http"
    "ms/sun/shared/x"
    "github.com/golang/protobuf/proto"
    "fmt"
)

type rpcHttpResponseHandler struct {
    w http.ResponseWriter
    r *http.Request
}

var cnt = 0

func (h *rpcHttpResponseHandler) HandleOfflineResult(resOut x.RpcResponseOutput) {
    //fmt.Println("rpcHttpResponseHandler")
    //helper.PertyPrint(resOut)
    cnt++
    //logHttpRpc.Println("s;" ,cnt)
    msg, ok := resOut.ResponseData.(proto.Message)
    if ok {
        bts, err := proto.Marshal(msg)
        if err != nil {
            if config.IS_DEBUG {
                log.Panic("ms- http rpc handling proto", err)
            }
            return
        }
        resToClient := &x.PB_ResponseToClient{
            ClientCallId: int64(resOut.CommandToServer.ClientCallId),
            PBClass:      resOut.PBClassName,
            RpcFullName:  resOut.RpcName,
            Data:         bts,
        }

        if true {
            fmt.Println("reslut: ", resToClient)
        }

        resToClientBts, err := proto.Marshal(resToClient)
        _ = resToClientBts
        if err != nil {
            if config.IS_DEBUG {
                log.Panic("ms- http rpc handling proto clinet", err)
            }
            return
        }
        if cnt%500 == 0 {
            if config.IS_DEBUG {
                //fmt.Println(cnt)
            }
        }
        fmt.Println(cnt)
        h.w.Write(resToClientBts)
        //h.w.Write([]byte("lgksjdglskgjslgk22222222222222222jsgl"))
    } else {
        if config.IS_DEBUG {
            log.Panic("ms- http HandleOfflineResult ResponseData is not of proto message ")
        }
    }
}

func (rpcHttpResponseHandler) IsUserOnlineResult(interface{}, error) {
    //panic("implement me")
}

//todo imple this
func (rpcHttpResponseHandler) HandelError(error) {
    //panic("implement me")
}
