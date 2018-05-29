package x

/////////////// Empty Sample RPC - mainly for mocking ////////////////

/////////////////// RPC_Chat  -  EmptyRPC_RPC_Chat /////////////////////
type EmptyRPC_RPC_Chat int

var Empty_RPC_RPC_Chat_Sample = EmptyRPC_RPC_Chat(0)

func (EmptyRPC_RPC_Chat) AddNewMessage(i *PB_RPC_Chat_Types_AddNewMessage_Param, p RPC_UserParam) (*PB_RPC_Chat_Types_AddNewMessage_Param, error) {
	return nil, nil
}
func (EmptyRPC_RPC_Chat) SetRoomActionDoing(i *PB_RPC_Chat_Types_SetRoomActionDoing_Param, p RPC_UserParam) (*PB_RPC_Chat_Types_SetRoomActionDoing_Param, error) {
	return nil, nil
}
func (EmptyRPC_RPC_Chat) GetChatList(i *PB_RPC_Chat_Types_GetChatList_Param, p RPC_UserParam) (*PB_RPC_Chat_Types_GetChatList_Param, error) {
	return nil, nil
}
func (EmptyRPC_RPC_Chat) GetChatHistory(i *PB_RPC_Chat_Types_GetChatHistory_Param, p RPC_UserParam) (*PB_RPC_Chat_Types_GetChatHistory_Param, error) {
	return nil, nil
}
func (EmptyRPC_RPC_Chat) PushRoomsChange(i *PB_RPC_Chat_Types_PushRoomsChange_Param, p RPC_UserParam) (*PB_RPC_Chat_Types_PushRoomsChange_Param, error) {
	return nil, nil
}
func (EmptyRPC_RPC_Chat) GetRoomsChange(i *PB_RPC_Chat_Types_GetRoomsChange_Param, p RPC_UserParam) (*PB_RPC_Chat_Types_GetRoomsChange_Param, error) {
	return nil, nil
}
