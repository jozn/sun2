package rpc_http_service

import (
	"ms/sun/shared/base"
	"ms/sun/shared/x"
	"time"
)

var rpcLogSaveChan = make(chan x.HTTPRPCLog, 1000)

func saveToDbLoggs_go() {
	tick := time.NewTicker(time.Second)
	var arr []x.HTTPRPCLog
	for {
		select {
		case c := <-rpcLogSaveChan:
			arr = append(arr, c)
		case <-tick.C:
			arr2 := arr
			arr = []x.HTTPRPCLog{}
			x.MassInsert_HTTPRPCLog(arr2, base.DB)
		}
	}
}
