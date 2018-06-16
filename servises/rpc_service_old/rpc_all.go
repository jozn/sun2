package rpc_service

import "ms/sun/shared/x"

//var RpcAll = x.RPC_AllHandlersInteract{}
var RpcAll = x.RPC_AllHandlersInteract{
	RPC_Chat:  rpc_chat(0),
	RPC_Other: rpc_other(0),
}

func init() {
	//RpcAll.RPC_Msg = rpcMsg(0)
	/*RpcAll.RPC_Sync = rpcSync(0)
	  RpcAll.RPC_Chat = rpcChat(0)
	  RpcAll.RPC_Other = rpcOther(74)*/

	//RpcAll.RPC_Other = rpc_other(0)
	//RpcAll.RPC_Chat = rpc_chat(0)
}
