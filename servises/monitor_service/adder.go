package monitor_service

var rpcChan = make(chan string, 1)

func AddRpc(rpc string) {
	rpcChan <- rpc
}

func addRpChanHanle_go() {
    //fmt.Println("*********************",HttpRpcMetrics)
	for method := range rpcChan {
		HttpRpcMetrics.TotalCall += 1
		HttpRpcMetrics.MethodsCalls[method] += 1
		//fmt.Println("*********************",HttpRpcMetrics)
	}
}
