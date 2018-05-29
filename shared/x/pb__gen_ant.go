package x

import (
	"errors"
	"log"
	"ms/sun/shared/config"

	"github.com/golang/protobuf/proto"
)

type RPC_UserParam interface {
	GetUserId() int
	IsUser() bool
}

type RPC_ResponseHandlerInterface interface {
	//HandleOfflineResult(dataPB interface{},PBClass string,RpcName string,cmd PB_CommandToServer,p RPC_UserParam, paramParsed interface{})
	HandleOfflineResult(resOut RpcResponseOutput)
	IsUserOnlineResult(interface{}, error)
	HandelError(error)
}

type RpcResponseOutput struct {
	UserParam       RPC_UserParam
	ResponseData    interface{}
	PBClassName     string
	RpcName         string
	CommandToServer PB_CommandToServer
	RpcParamPassed  interface{}
}

var RPC_ResponseHandler RPC_ResponseHandlerInterface

//note: rpc methods cant have equal name they must be different even in different rpc services
type RPC_AllHandlersInteract struct {
	RPC_Chat RPC_Chat
}

/////////////// Interfaces ////////////////

type RPC_Chat interface {
	AddNewMessage(param *PB_RPC_Chat_Types_AddNewMessage_Param, userParam RPC_UserParam) (res PB_RPC_Chat_Types_AddNewMessage_Param, err error)
	SetRoomActionDoing(param *PB_RPC_Chat_Types_SetRoomActionDoing_Param, userParam RPC_UserParam) (res PB_RPC_Chat_Types_SetRoomActionDoing_Param, err error)
	GetChatList(param *PB_RPC_Chat_Types_GetChatList_Param, userParam RPC_UserParam) (res PB_RPC_Chat_Types_GetChatList_Param, err error)
	GetChatHistory(param *PB_RPC_Chat_Types_GetChatHistory_Param, userParam RPC_UserParam) (res PB_RPC_Chat_Types_GetChatHistory_Param, err error)
	PushRoomsChange(param *PB_RPC_Chat_Types_PushRoomsChange_Param, userParam RPC_UserParam) (res PB_RPC_Chat_Types_PushRoomsChange_Param, err error)
	GetRoomsChange(param *PB_RPC_Chat_Types_GetRoomsChange_Param, userParam RPC_UserParam) (res PB_RPC_Chat_Types_GetRoomsChange_Param, err error)
}

func noDevErr(err error) {
	if config.IS_DEBUG && err != nil {
		log.Panic(err)
	}
}

type rpcParamHandler struct {
	cmd             PB_CommandToServer
	params          RPC_UserParam
	rpcHandler      RPC_AllHandlersInteract
	responseHandler RPC_ResponseHandlerInterface
}

//todo: this method can be outputed from x package
////////////// map of rpc methods to all
func HandleRpcs(cmd PB_CommandToServer, params RPC_UserParam, rpcHandler RPC_AllHandlersInteract, responseHandler RPC_ResponseHandlerInterface) {

	fn, ok := mpRpcMethods[cmd.Command]
	if !ok {
		if config.IS_DEBUG {
			log.Panic("HandleRpcs:  command not registerd for ", cmd.Command)
		}
	}

	p := rpcParamHandler{
		cmd:             cmd,
		params:          params,
		rpcHandler:      rpcHandler,
		responseHandler: responseHandler,
	}
	fn(p)
}

var mpRpcMethods = map[string]func(p rpcParamHandler){

	// rpc: RPC_Chat

	"RPC_Chat.AddNewMessage": func(p rpcParamHandler) {
		if p.rpcHandler.RPC_Chat == nil {
			noDevErr(errors.New("rpc service is null for: p.rpcHandler.RPC_Chat.AddNewMessage"))
			return
		}
		load := &PB_RPC_Chat_Types_AddNewMessage_Param{}
		err := proto.Unmarshal(p.cmd.Data, load)
		if err == nil {
			res, err := p.rpcHandler.RPC_Chat.AddNewMessage(load, p.params)
			if err == nil {
				out := RpcResponseOutput{
					RpcName:         "RPC_Chat.AddNewMessage",
					UserParam:       p.params,
					CommandToServer: p.cmd,
					PBClassName:     "PB_RPC_Chat_Types.AddNewMessage.Response",
					ResponseData:    &res,
					RpcParamPassed:  load,
				}
				p.responseHandler.HandleOfflineResult(out)
			} else {
				p.responseHandler.HandelError(err)
			}
		} else {
			p.responseHandler.HandelError(err)
		}
	},
	"RPC_Chat.SetRoomActionDoing": func(p rpcParamHandler) {
		if p.rpcHandler.RPC_Chat == nil {
			noDevErr(errors.New("rpc service is null for: p.rpcHandler.RPC_Chat.SetRoomActionDoing"))
			return
		}
		load := &PB_RPC_Chat_Types_SetRoomActionDoing_Param{}
		err := proto.Unmarshal(p.cmd.Data, load)
		if err == nil {
			res, err := p.rpcHandler.RPC_Chat.SetRoomActionDoing(load, p.params)
			if err == nil {
				out := RpcResponseOutput{
					RpcName:         "RPC_Chat.SetRoomActionDoing",
					UserParam:       p.params,
					CommandToServer: p.cmd,
					PBClassName:     "PB_RPC_Chat_Types.SetRoomActionDoing.Response",
					ResponseData:    &res,
					RpcParamPassed:  load,
				}
				p.responseHandler.HandleOfflineResult(out)
			} else {
				p.responseHandler.HandelError(err)
			}
		} else {
			p.responseHandler.HandelError(err)
		}
	},
	"RPC_Chat.GetChatList": func(p rpcParamHandler) {
		if p.rpcHandler.RPC_Chat == nil {
			noDevErr(errors.New("rpc service is null for: p.rpcHandler.RPC_Chat.GetChatList"))
			return
		}
		load := &PB_RPC_Chat_Types_GetChatList_Param{}
		err := proto.Unmarshal(p.cmd.Data, load)
		if err == nil {
			res, err := p.rpcHandler.RPC_Chat.GetChatList(load, p.params)
			if err == nil {
				out := RpcResponseOutput{
					RpcName:         "RPC_Chat.GetChatList",
					UserParam:       p.params,
					CommandToServer: p.cmd,
					PBClassName:     "PB_RPC_Chat_Types.GetChatList.Response",
					ResponseData:    &res,
					RpcParamPassed:  load,
				}
				p.responseHandler.HandleOfflineResult(out)
			} else {
				p.responseHandler.HandelError(err)
			}
		} else {
			p.responseHandler.HandelError(err)
		}
	},
	"RPC_Chat.GetChatHistory": func(p rpcParamHandler) {
		if p.rpcHandler.RPC_Chat == nil {
			noDevErr(errors.New("rpc service is null for: p.rpcHandler.RPC_Chat.GetChatHistory"))
			return
		}
		load := &PB_RPC_Chat_Types_GetChatHistory_Param{}
		err := proto.Unmarshal(p.cmd.Data, load)
		if err == nil {
			res, err := p.rpcHandler.RPC_Chat.GetChatHistory(load, p.params)
			if err == nil {
				out := RpcResponseOutput{
					RpcName:         "RPC_Chat.GetChatHistory",
					UserParam:       p.params,
					CommandToServer: p.cmd,
					PBClassName:     "PB_RPC_Chat_Types.GetChatHistory.Response",
					ResponseData:    &res,
					RpcParamPassed:  load,
				}
				p.responseHandler.HandleOfflineResult(out)
			} else {
				p.responseHandler.HandelError(err)
			}
		} else {
			p.responseHandler.HandelError(err)
		}
	},
	"RPC_Chat.PushRoomsChange": func(p rpcParamHandler) {
		if p.rpcHandler.RPC_Chat == nil {
			noDevErr(errors.New("rpc service is null for: p.rpcHandler.RPC_Chat.PushRoomsChange"))
			return
		}
		load := &PB_RPC_Chat_Types_PushRoomsChange_Param{}
		err := proto.Unmarshal(p.cmd.Data, load)
		if err == nil {
			res, err := p.rpcHandler.RPC_Chat.PushRoomsChange(load, p.params)
			if err == nil {
				out := RpcResponseOutput{
					RpcName:         "RPC_Chat.PushRoomsChange",
					UserParam:       p.params,
					CommandToServer: p.cmd,
					PBClassName:     "PB_RPC_Chat_Types.PushRoomsChange.Response",
					ResponseData:    &res,
					RpcParamPassed:  load,
				}
				p.responseHandler.HandleOfflineResult(out)
			} else {
				p.responseHandler.HandelError(err)
			}
		} else {
			p.responseHandler.HandelError(err)
		}
	},
	"RPC_Chat.GetRoomsChange": func(p rpcParamHandler) {
		if p.rpcHandler.RPC_Chat == nil {
			noDevErr(errors.New("rpc service is null for: p.rpcHandler.RPC_Chat.GetRoomsChange"))
			return
		}
		load := &PB_RPC_Chat_Types_GetRoomsChange_Param{}
		err := proto.Unmarshal(p.cmd.Data, load)
		if err == nil {
			res, err := p.rpcHandler.RPC_Chat.GetRoomsChange(load, p.params)
			if err == nil {
				out := RpcResponseOutput{
					RpcName:         "RPC_Chat.GetRoomsChange",
					UserParam:       p.params,
					CommandToServer: p.cmd,
					PBClassName:     "PB_RPC_Chat_Types.GetRoomsChange.Response",
					ResponseData:    &res,
					RpcParamPassed:  load,
				}
				p.responseHandler.HandleOfflineResult(out)
			} else {
				p.responseHandler.HandelError(err)
			}
		} else {
			p.responseHandler.HandelError(err)
		}
	},
}

/////////////// Direct in PB_CommandToClient /////////////
/*


 RPC_Chat.AddNewMessage
 RPC_Chat.SetRoomActionDoing
 RPC_Chat.GetChatList
 RPC_Chat.GetChatHistory
 RPC_Chat.PushRoomsChange
 RPC_Chat.GetRoomsChange


*/
