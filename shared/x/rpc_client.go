package x

import "github.com/golang/protobuf/proto"

//type RPCClientHandler func(cmd string, pb interface{}) interface{}
type RPCClientHandler func(cmdSre string, pbIn, pbOut proto.Message) error

// all clients struc
var RPC_AllClinetsPlay = struct {
	RPC_Chat RPC_Chat_Client
}{
	RPC_Chat: RPC_Chat_Client(0),
}

// client types defs
type RPC_Chat_Client int

/////////////// methods ////////////////

// service: RPC_Chat

func (RPC_Chat_Client) AddNewMessage(param *PB_RPC_Chat_Types_AddNewMessage_Param, fn RPCClientHandler) (*PB_RPC_Chat_Types_AddNewMessage_Param, error) {
	out := &PB_RPC_Chat_Types_AddNewMessage_Param{}
	err := fn("RPC_Chat.AddNewMessage", param, out)
	if err == nil {
		return out, nil
	}
	return nil, err
}

func (RPC_Chat_Client) SetRoomActionDoing(param *PB_RPC_Chat_Types_SetRoomActionDoing_Param, fn RPCClientHandler) (*PB_RPC_Chat_Types_SetRoomActionDoing_Param, error) {
	out := &PB_RPC_Chat_Types_SetRoomActionDoing_Param{}
	err := fn("RPC_Chat.SetRoomActionDoing", param, out)
	if err == nil {
		return out, nil
	}
	return nil, err
}

func (RPC_Chat_Client) GetChatList(param *PB_RPC_Chat_Types_GetChatList_Param, fn RPCClientHandler) (*PB_RPC_Chat_Types_GetChatList_Param, error) {
	out := &PB_RPC_Chat_Types_GetChatList_Param{}
	err := fn("RPC_Chat.GetChatList", param, out)
	if err == nil {
		return out, nil
	}
	return nil, err
}

func (RPC_Chat_Client) GetChatHistory(param *PB_RPC_Chat_Types_GetChatHistory_Param, fn RPCClientHandler) (*PB_RPC_Chat_Types_GetChatHistory_Param, error) {
	out := &PB_RPC_Chat_Types_GetChatHistory_Param{}
	err := fn("RPC_Chat.GetChatHistory", param, out)
	if err == nil {
		return out, nil
	}
	return nil, err
}

func (RPC_Chat_Client) PushRoomsChange(param *PB_RPC_Chat_Types_PushRoomsChange_Param, fn RPCClientHandler) (*PB_RPC_Chat_Types_PushRoomsChange_Param, error) {
	out := &PB_RPC_Chat_Types_PushRoomsChange_Param{}
	err := fn("RPC_Chat.PushRoomsChange", param, out)
	if err == nil {
		return out, nil
	}
	return nil, err
}

func (RPC_Chat_Client) GetRoomsChange(param *PB_RPC_Chat_Types_GetRoomsChange_Param, fn RPCClientHandler) (*PB_RPC_Chat_Types_GetRoomsChange_Param, error) {
	out := &PB_RPC_Chat_Types_GetRoomsChange_Param{}
	err := fn("RPC_Chat.GetRoomsChange", param, out)
	if err == nil {
		return out, nil
	}
	return nil, err
}
