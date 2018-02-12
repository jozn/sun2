package rpc_service

import (
	"ms/sun2/servises/chat_service"
	"ms/sun2/shared/x"
)

type rpc_chat int

func (rpc_chat) AddNewMessage(param *x.PB_ChatParam_AddNewMessage, userParam x.RPC_UserParam) (res x.PB_ChatResponse_AddNewMessage, err error) {
    chat_service.ChatAddNewMsg(userParam.GetUserId(),param.MessageView,param.FileBlob)
    return
}

func (rpc_chat) SetRoomActionDoing(param *x.PB_ChatParam_SetRoomActionDoing, userParam x.RPC_UserParam) (res x.PB_ChatResponse_SetRoomActionDoing, err error) {
	panic("implement me")
}

func (rpc_chat) SetMessagesAsReceived(param *x.PB_ChatParam_SetMessagesAsReceived, userParam x.RPC_UserParam) (res x.PB_ChatResponse_SetMessagesAsReceived, err error) {

	return
}

func (rpc_chat) SetMessagesRangeAsSeen(param *x.PB_ChatParam_SetChatMessagesRangeAsSeen, userParam x.RPC_UserParam) (res x.PB_ChatResponse_SetChatMessagesRangeAsSeen, err error) {
	panic("implement me")
}

func (rpc_chat) DeleteChatHistory(param *x.PB_ChatParam_DeleteChatHistory, userParam x.RPC_UserParam) (res x.PB_ChatResponse_DeleteChatHistory, err error) {
	panic("implement me")
}

func (rpc_chat) DeleteMessagesByIds(param *x.PB_ChatParam_DeleteMessagesByIds, userParam x.RPC_UserParam) (res x.PB_ChatResponse_DeleteMessagesByIds, err error) {
	panic("implement me")
}

func (rpc_chat) EditMessage(param *x.PB_ChatParam_EditMessage, userParam x.RPC_UserParam) (res x.PB_ChatResponse_EditMessage, err error) {
	panic("implement me")
}

func (rpc_chat) GetChatList(param *x.PB_ChatParam_GetChatList, userParam x.RPC_UserParam) (res x.PB_ChatResponse_GetChatList, err error) {
	res.Chats = chat_service.GetChatList(userParam.GetUserId())
	return
}

func (rpc_chat) GetChatHistoryToOlder(param *x.PB_ChatParam_GetChatHistoryToOlder, userParam x.RPC_UserParam) (res x.PB_ChatResponse_GetChatHistoryToOlder, err error) {
	res.MessagesViews, res.HasMore = chat_service.GetMessageList(userParam.GetUserId(), param.ChatKey, int(param.FromTopMessageId), int(param.Limit))
	return
}
