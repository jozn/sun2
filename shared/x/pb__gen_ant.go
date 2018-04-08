package x

import (
	"errors"
	"log"
	"ms/sun_old/config"
	"strings"

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
	RPC_Auth   RPC_Auth
	RPC_Chat   RPC_Chat
	RPC_Other  RPC_Other
	RPC_Page   RPC_Page
	RPC_Search RPC_Search
	RPC_Social RPC_Social
	RPC_User   RPC_User
}

/////////////// Interfaces ////////////////

type RPC_Auth interface {
	CheckPhone(param *PB_UserParam_CheckUserName2, userParam RPC_UserParam) (res PB_UserResponse_CheckUserName2, err error)
	SendCode(param *PB_UserParam_CheckUserName2, userParam RPC_UserParam) (res PB_UserResponse_CheckUserName2, err error)
	SendCodeToSms(param *PB_UserParam_CheckUserName2, userParam RPC_UserParam) (res PB_UserResponse_CheckUserName2, err error)
	SendCodeToTelgram(param *PB_UserParam_CheckUserName2, userParam RPC_UserParam) (res PB_UserResponse_CheckUserName2, err error)
	SingUp(param *PB_UserParam_CheckUserName2, userParam RPC_UserParam) (res PB_UserResponse_CheckUserName2, err error)
	SingIn(param *PB_UserParam_CheckUserName2, userParam RPC_UserParam) (res PB_UserResponse_CheckUserName2, err error)
	LogOut(param *PB_UserParam_CheckUserName2, userParam RPC_UserParam) (res PB_UserResponse_CheckUserName2, err error)
}

type RPC_Chat interface {
	AddNewMessage(param *PB_ChatParam_AddNewMessage, userParam RPC_UserParam) (res PB_ChatResponse_AddNewMessage, err error)
	SetRoomActionDoing(param *PB_ChatParam_SetRoomActionDoing, userParam RPC_UserParam) (res PB_ChatResponse_SetRoomActionDoing, err error)
	SetMessagesAsReceived(param *PB_ChatParam_SetMessagesAsReceived, userParam RPC_UserParam) (res PB_ChatResponse_SetMessagesAsReceived, err error)
	SetMessagesRangeAsSeen(param *PB_ChatParam_SetChatMessagesRangeAsSeen, userParam RPC_UserParam) (res PB_ChatResponse_SetChatMessagesRangeAsSeen, err error)
	DeleteChatHistory(param *PB_ChatParam_DeleteChatHistory, userParam RPC_UserParam) (res PB_ChatResponse_DeleteChatHistory, err error)
	DeleteMessagesByIds(param *PB_ChatParam_DeleteMessagesByIds, userParam RPC_UserParam) (res PB_ChatResponse_DeleteMessagesByIds, err error)
	EditMessage(param *PB_ChatParam_EditMessage, userParam RPC_UserParam) (res PB_ChatResponse_EditMessage, err error)
	GetChatList(param *PB_ChatParam_GetChatList, userParam RPC_UserParam) (res PB_ChatResponse_GetChatList, err error)
	GetChatHistoryToOlder(param *PB_ChatParam_GetChatHistoryToOlder, userParam RPC_UserParam) (res PB_ChatResponse_GetChatHistoryToOlder, err error)
}

type RPC_Other interface {
	Echo(param *PB_OtherParam_Echo, userParam RPC_UserParam) (res PB_OtherResponse_Echo, err error)
}

type RPC_Page interface {
	GetCommentsPage(param *PB_PageParam_GetCommentsPage, userParam RPC_UserParam) (res PB_PageResponse_GetCommentsPage, err error)
	GetHomePage(param *PB_PageParam_GetHomePage, userParam RPC_UserParam) (res PB_PageResponse_GetHomePage, err error)
	GetProfilePage(param *PB_PageParam_GetProfilePage, userParam RPC_UserParam) (res PB_PageResponse_GetProfilePage, err error)
	GetLikesPage(param *PB_PageParam_GetLikesPage, userParam RPC_UserParam) (res PB_PageResponse_GetLikesPage, err error)
	GetFollowersPage(param *PB_PageParam_GetFollowersPage, userParam RPC_UserParam) (res PB_PageResponse_GetFollowersPage, err error)
	GetFollowingsPage(param *PB_PageParam_GetFollowingsPage, userParam RPC_UserParam) (res PB_PageResponse_GetFollowingsPage, err error)
	GetNotifiesPage(param *PB_PageParam_GetNotifiesPage, userParam RPC_UserParam) (res PB_PageResponse_GetNotifiesPage, err error)
	GetUserActionsPage(param *PB_PageParam_GetUserActionsPage, userParam RPC_UserParam) (res PB_PageResponse_GetUserActionsPage, err error)
	GetSuggestedPostsPage(param *PB_PageParam_GetSuggestedPostsPage, userParam RPC_UserParam) (res PB_PageResponse_GetSuggestedPostsPage, err error)
	GetSuggestedUsersPage(param *PB_PageParam_GetSuggestedUsersPage, userParam RPC_UserParam) (res PB_PageResponse_GetSuggestedUsersPage, err error)
	GetSuggestedTagsPage(param *PB_PageParam_GetSuggestedTagsPage, userParam RPC_UserParam) (res PB_PageResponse_GetSuggestedTagsPage, err error)
	GetLastPostsPage(param *PB_PageParam_GetLastPostsPage, userParam RPC_UserParam) (res PB_PageResponse_GetLastPostsPage, err error)
	GetTagPage(param *PB_PageParam_GetTagPage, userParam RPC_UserParam) (res PB_PageResponse_GetTagPage, err error)
	SearchTagsPage(param *PB_PageParam_SearchTagsPage, userParam RPC_UserParam) (res PB_PageResponse_SearchTagsPage, err error)
	SearchUsersPage(param *PB_PageParam_SearchUsersPage, userParam RPC_UserParam) (res PB_PageResponse_SearchUsersPage, err error)
}

type RPC_Search interface {
	SearchTags(param *PB_SearchResponse_AddNewC, userParam RPC_UserParam) (res PB_SearchResponse_AddNewC, err error)
}

type RPC_Social interface {
	AddComment(param *PB_SocialParam_AddComment, userParam RPC_UserParam) (res PB_SocialResponse_AddComment, err error)
	DeleteComment(param *PB_SocialParam_DeleteComment, userParam RPC_UserParam) (res PB_SocialResponse_DeleteComment, err error)
	AddPost(param *PB_SocialParam_AddPost, userParam RPC_UserParam) (res PB_SocialResponse_AddPost, err error)
	EditPost(param *PB_SocialParam_EditPost, userParam RPC_UserParam) (res PB_SocialResponse_EditPost, err error)
	DeletePost(param *PB_SocialParam_DeletePost, userParam RPC_UserParam) (res PB_SocialResponse_DeletePost, err error)
	ArchivePost(param *PB_SocialParam_ArchivePost, userParam RPC_UserParam) (res PB_SocialResponse_ArchivePost, err error)
	LikePost(param *PB_SocialParam_LikePost, userParam RPC_UserParam) (res PB_SocialResponse_LikePost, err error)
	UnLikePost(param *PB_SocialParam_UnLikePost, userParam RPC_UserParam) (res PB_SocialResponse_UnLikePost, err error)
	FollowUser(param *PB_SocialParam_FollowUser, userParam RPC_UserParam) (res PB_SocialResponse_FollowUser, err error)
	UnFollowUser(param *PB_SocialParam_UnFollowUser, userParam RPC_UserParam) (res PB_SocialResponse_UnFollowUser, err error)
}

type RPC_User interface {
	BlockUser(param *PB_UserParam_BlockUser, userParam RPC_UserParam) (res PB_UserResponse_BlockUser, err error)
	UnBlockUser(param *PB_UserParam_UnBlockUser, userParam RPC_UserParam) (res PB_UserResponse_UnBlockUser, err error)
	GetBlockedList(param *PB_UserParam_BlockedList, userParam RPC_UserParam) (res PB_UserResponse_BlockedList, err error)
	UpdateAbout(param *PB_UserParam_UpdateAbout, userParam RPC_UserParam) (res PB_UserResponse_UpdateAbout, err error)
	UpdateUserName(param *PB_UserParam_UpdateUserName, userParam RPC_UserParam) (res PB_UserResponse_UpdateUserName, err error)
	ChangePrivacy(param *PB_UserParam_ChangePrivacy, userParam RPC_UserParam) (res PB_UserResponseOffline_ChangePrivacy, err error)
	ChangeAvatar(param *PB_UserParam_ChangeAvatar, userParam RPC_UserParam) (res PB_UserResponse_ChangeAvatar, err error)
	CheckUserName(param *PB_UserParam_CheckUserName, userParam RPC_UserParam) (res PB_UserResponse_CheckUserName, err error)
}

func noDevErr(err error) {
	if config.IS_DEBUG && err != nil {
		log.Panic(err)
	}
}

////////////// map of rpc methods to all
func HandleRpcs(cmd PB_CommandToServer, params RPC_UserParam, rpcHandler RPC_AllHandlersInteract, responseHandler RPC_ResponseHandlerInterface) {

	splits := strings.Split(cmd.Command, ".")

	if len(splits) != 2 {
		noDevErr(errors.New("HandleRpcs: splic is not 2 parts"))
		return
	}

	switch splits[0] {

	case "RPC_Auth":

		//rpc,ok := rpcHandler.RPC_Auth
		rpc := rpcHandler.RPC_Auth
		/*if !ok {
		    e:=errors.New("rpcHandler could not be cast to : RPC_Auth")
		    noDevErr(e)
		    RPC_ResponseHandler.HandelError(e)
		    return
		}*/

		switch splits[1] {
		case "CheckPhone": //each pb_service_method
			load := &PB_UserParam_CheckUserName2{}
			err := proto.Unmarshal(cmd.Data, load)
			if err == nil {
				res, err := rpc.CheckPhone(load, params)
				if err == nil {
					out := RpcResponseOutput{
						RpcName:         "RPC_Auth.CheckPhone",
						UserParam:       params,
						CommandToServer: cmd,
						PBClassName:     "PB_UserResponse_CheckUserName2",
						ResponseData:    &res,
						RpcParamPassed:  load,
					}
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_UserResponse_CheckUserName2",cmd, params)
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_UserResponse_CheckUserName2","RPC_Auth.CheckPhone",cmd, params , load)
					responseHandler.HandleOfflineResult(out)
				} else {
					responseHandler.HandelError(err)
				}
			} else {
				responseHandler.HandelError(err)
			}
		case "SendCode": //each pb_service_method
			load := &PB_UserParam_CheckUserName2{}
			err := proto.Unmarshal(cmd.Data, load)
			if err == nil {
				res, err := rpc.SendCode(load, params)
				if err == nil {
					out := RpcResponseOutput{
						RpcName:         "RPC_Auth.SendCode",
						UserParam:       params,
						CommandToServer: cmd,
						PBClassName:     "PB_UserResponse_CheckUserName2",
						ResponseData:    &res,
						RpcParamPassed:  load,
					}
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_UserResponse_CheckUserName2",cmd, params)
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_UserResponse_CheckUserName2","RPC_Auth.SendCode",cmd, params , load)
					responseHandler.HandleOfflineResult(out)
				} else {
					responseHandler.HandelError(err)
				}
			} else {
				responseHandler.HandelError(err)
			}
		case "SendCodeToSms": //each pb_service_method
			load := &PB_UserParam_CheckUserName2{}
			err := proto.Unmarshal(cmd.Data, load)
			if err == nil {
				res, err := rpc.SendCodeToSms(load, params)
				if err == nil {
					out := RpcResponseOutput{
						RpcName:         "RPC_Auth.SendCodeToSms",
						UserParam:       params,
						CommandToServer: cmd,
						PBClassName:     "PB_UserResponse_CheckUserName2",
						ResponseData:    &res,
						RpcParamPassed:  load,
					}
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_UserResponse_CheckUserName2",cmd, params)
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_UserResponse_CheckUserName2","RPC_Auth.SendCodeToSms",cmd, params , load)
					responseHandler.HandleOfflineResult(out)
				} else {
					responseHandler.HandelError(err)
				}
			} else {
				responseHandler.HandelError(err)
			}
		case "SendCodeToTelgram": //each pb_service_method
			load := &PB_UserParam_CheckUserName2{}
			err := proto.Unmarshal(cmd.Data, load)
			if err == nil {
				res, err := rpc.SendCodeToTelgram(load, params)
				if err == nil {
					out := RpcResponseOutput{
						RpcName:         "RPC_Auth.SendCodeToTelgram",
						UserParam:       params,
						CommandToServer: cmd,
						PBClassName:     "PB_UserResponse_CheckUserName2",
						ResponseData:    &res,
						RpcParamPassed:  load,
					}
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_UserResponse_CheckUserName2",cmd, params)
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_UserResponse_CheckUserName2","RPC_Auth.SendCodeToTelgram",cmd, params , load)
					responseHandler.HandleOfflineResult(out)
				} else {
					responseHandler.HandelError(err)
				}
			} else {
				responseHandler.HandelError(err)
			}
		case "SingUp": //each pb_service_method
			load := &PB_UserParam_CheckUserName2{}
			err := proto.Unmarshal(cmd.Data, load)
			if err == nil {
				res, err := rpc.SingUp(load, params)
				if err == nil {
					out := RpcResponseOutput{
						RpcName:         "RPC_Auth.SingUp",
						UserParam:       params,
						CommandToServer: cmd,
						PBClassName:     "PB_UserResponse_CheckUserName2",
						ResponseData:    &res,
						RpcParamPassed:  load,
					}
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_UserResponse_CheckUserName2",cmd, params)
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_UserResponse_CheckUserName2","RPC_Auth.SingUp",cmd, params , load)
					responseHandler.HandleOfflineResult(out)
				} else {
					responseHandler.HandelError(err)
				}
			} else {
				responseHandler.HandelError(err)
			}
		case "SingIn": //each pb_service_method
			load := &PB_UserParam_CheckUserName2{}
			err := proto.Unmarshal(cmd.Data, load)
			if err == nil {
				res, err := rpc.SingIn(load, params)
				if err == nil {
					out := RpcResponseOutput{
						RpcName:         "RPC_Auth.SingIn",
						UserParam:       params,
						CommandToServer: cmd,
						PBClassName:     "PB_UserResponse_CheckUserName2",
						ResponseData:    &res,
						RpcParamPassed:  load,
					}
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_UserResponse_CheckUserName2",cmd, params)
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_UserResponse_CheckUserName2","RPC_Auth.SingIn",cmd, params , load)
					responseHandler.HandleOfflineResult(out)
				} else {
					responseHandler.HandelError(err)
				}
			} else {
				responseHandler.HandelError(err)
			}
		case "LogOut": //each pb_service_method
			load := &PB_UserParam_CheckUserName2{}
			err := proto.Unmarshal(cmd.Data, load)
			if err == nil {
				res, err := rpc.LogOut(load, params)
				if err == nil {
					out := RpcResponseOutput{
						RpcName:         "RPC_Auth.LogOut",
						UserParam:       params,
						CommandToServer: cmd,
						PBClassName:     "PB_UserResponse_CheckUserName2",
						ResponseData:    &res,
						RpcParamPassed:  load,
					}
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_UserResponse_CheckUserName2",cmd, params)
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_UserResponse_CheckUserName2","RPC_Auth.LogOut",cmd, params , load)
					responseHandler.HandleOfflineResult(out)
				} else {
					responseHandler.HandelError(err)
				}
			} else {
				responseHandler.HandelError(err)
			}
		default:
			noDevErr(errors.New("rpc method is does not exist: " + cmd.Command))
		}
	case "RPC_Chat":

		//rpc,ok := rpcHandler.RPC_Chat
		rpc := rpcHandler.RPC_Chat
		/*if !ok {
		    e:=errors.New("rpcHandler could not be cast to : RPC_Chat")
		    noDevErr(e)
		    RPC_ResponseHandler.HandelError(e)
		    return
		}*/

		switch splits[1] {
		case "AddNewMessage": //each pb_service_method
			load := &PB_ChatParam_AddNewMessage{}
			err := proto.Unmarshal(cmd.Data, load)
			if err == nil {
				res, err := rpc.AddNewMessage(load, params)
				if err == nil {
					out := RpcResponseOutput{
						RpcName:         "RPC_Chat.AddNewMessage",
						UserParam:       params,
						CommandToServer: cmd,
						PBClassName:     "PB_ChatResponse_AddNewMessage",
						ResponseData:    &res,
						RpcParamPassed:  load,
					}
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_ChatResponse_AddNewMessage",cmd, params)
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_ChatResponse_AddNewMessage","RPC_Chat.AddNewMessage",cmd, params , load)
					responseHandler.HandleOfflineResult(out)
				} else {
					responseHandler.HandelError(err)
				}
			} else {
				responseHandler.HandelError(err)
			}
		case "SetRoomActionDoing": //each pb_service_method
			load := &PB_ChatParam_SetRoomActionDoing{}
			err := proto.Unmarshal(cmd.Data, load)
			if err == nil {
				res, err := rpc.SetRoomActionDoing(load, params)
				if err == nil {
					out := RpcResponseOutput{
						RpcName:         "RPC_Chat.SetRoomActionDoing",
						UserParam:       params,
						CommandToServer: cmd,
						PBClassName:     "PB_ChatResponse_SetRoomActionDoing",
						ResponseData:    &res,
						RpcParamPassed:  load,
					}
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_ChatResponse_SetRoomActionDoing",cmd, params)
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_ChatResponse_SetRoomActionDoing","RPC_Chat.SetRoomActionDoing",cmd, params , load)
					responseHandler.HandleOfflineResult(out)
				} else {
					responseHandler.HandelError(err)
				}
			} else {
				responseHandler.HandelError(err)
			}
		case "SetMessagesAsReceived": //each pb_service_method
			load := &PB_ChatParam_SetMessagesAsReceived{}
			err := proto.Unmarshal(cmd.Data, load)
			if err == nil {
				res, err := rpc.SetMessagesAsReceived(load, params)
				if err == nil {
					out := RpcResponseOutput{
						RpcName:         "RPC_Chat.SetMessagesAsReceived",
						UserParam:       params,
						CommandToServer: cmd,
						PBClassName:     "PB_ChatResponse_SetMessagesAsReceived",
						ResponseData:    &res,
						RpcParamPassed:  load,
					}
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_ChatResponse_SetMessagesAsReceived",cmd, params)
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_ChatResponse_SetMessagesAsReceived","RPC_Chat.SetMessagesAsReceived",cmd, params , load)
					responseHandler.HandleOfflineResult(out)
				} else {
					responseHandler.HandelError(err)
				}
			} else {
				responseHandler.HandelError(err)
			}
		case "SetMessagesRangeAsSeen": //each pb_service_method
			load := &PB_ChatParam_SetChatMessagesRangeAsSeen{}
			err := proto.Unmarshal(cmd.Data, load)
			if err == nil {
				res, err := rpc.SetMessagesRangeAsSeen(load, params)
				if err == nil {
					out := RpcResponseOutput{
						RpcName:         "RPC_Chat.SetMessagesRangeAsSeen",
						UserParam:       params,
						CommandToServer: cmd,
						PBClassName:     "PB_ChatResponse_SetChatMessagesRangeAsSeen",
						ResponseData:    &res,
						RpcParamPassed:  load,
					}
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_ChatResponse_SetChatMessagesRangeAsSeen",cmd, params)
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_ChatResponse_SetChatMessagesRangeAsSeen","RPC_Chat.SetMessagesRangeAsSeen",cmd, params , load)
					responseHandler.HandleOfflineResult(out)
				} else {
					responseHandler.HandelError(err)
				}
			} else {
				responseHandler.HandelError(err)
			}
		case "DeleteChatHistory": //each pb_service_method
			load := &PB_ChatParam_DeleteChatHistory{}
			err := proto.Unmarshal(cmd.Data, load)
			if err == nil {
				res, err := rpc.DeleteChatHistory(load, params)
				if err == nil {
					out := RpcResponseOutput{
						RpcName:         "RPC_Chat.DeleteChatHistory",
						UserParam:       params,
						CommandToServer: cmd,
						PBClassName:     "PB_ChatResponse_DeleteChatHistory",
						ResponseData:    &res,
						RpcParamPassed:  load,
					}
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_ChatResponse_DeleteChatHistory",cmd, params)
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_ChatResponse_DeleteChatHistory","RPC_Chat.DeleteChatHistory",cmd, params , load)
					responseHandler.HandleOfflineResult(out)
				} else {
					responseHandler.HandelError(err)
				}
			} else {
				responseHandler.HandelError(err)
			}
		case "DeleteMessagesByIds": //each pb_service_method
			load := &PB_ChatParam_DeleteMessagesByIds{}
			err := proto.Unmarshal(cmd.Data, load)
			if err == nil {
				res, err := rpc.DeleteMessagesByIds(load, params)
				if err == nil {
					out := RpcResponseOutput{
						RpcName:         "RPC_Chat.DeleteMessagesByIds",
						UserParam:       params,
						CommandToServer: cmd,
						PBClassName:     "PB_ChatResponse_DeleteMessagesByIds",
						ResponseData:    &res,
						RpcParamPassed:  load,
					}
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_ChatResponse_DeleteMessagesByIds",cmd, params)
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_ChatResponse_DeleteMessagesByIds","RPC_Chat.DeleteMessagesByIds",cmd, params , load)
					responseHandler.HandleOfflineResult(out)
				} else {
					responseHandler.HandelError(err)
				}
			} else {
				responseHandler.HandelError(err)
			}
		case "EditMessage": //each pb_service_method
			load := &PB_ChatParam_EditMessage{}
			err := proto.Unmarshal(cmd.Data, load)
			if err == nil {
				res, err := rpc.EditMessage(load, params)
				if err == nil {
					out := RpcResponseOutput{
						RpcName:         "RPC_Chat.EditMessage",
						UserParam:       params,
						CommandToServer: cmd,
						PBClassName:     "PB_ChatResponse_EditMessage",
						ResponseData:    &res,
						RpcParamPassed:  load,
					}
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_ChatResponse_EditMessage",cmd, params)
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_ChatResponse_EditMessage","RPC_Chat.EditMessage",cmd, params , load)
					responseHandler.HandleOfflineResult(out)
				} else {
					responseHandler.HandelError(err)
				}
			} else {
				responseHandler.HandelError(err)
			}
		case "GetChatList": //each pb_service_method
			load := &PB_ChatParam_GetChatList{}
			err := proto.Unmarshal(cmd.Data, load)
			if err == nil {
				res, err := rpc.GetChatList(load, params)
				if err == nil {
					out := RpcResponseOutput{
						RpcName:         "RPC_Chat.GetChatList",
						UserParam:       params,
						CommandToServer: cmd,
						PBClassName:     "PB_ChatResponse_GetChatList",
						ResponseData:    &res,
						RpcParamPassed:  load,
					}
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_ChatResponse_GetChatList",cmd, params)
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_ChatResponse_GetChatList","RPC_Chat.GetChatList",cmd, params , load)
					responseHandler.HandleOfflineResult(out)
				} else {
					responseHandler.HandelError(err)
				}
			} else {
				responseHandler.HandelError(err)
			}
		case "GetChatHistoryToOlder": //each pb_service_method
			load := &PB_ChatParam_GetChatHistoryToOlder{}
			err := proto.Unmarshal(cmd.Data, load)
			if err == nil {
				res, err := rpc.GetChatHistoryToOlder(load, params)
				if err == nil {
					out := RpcResponseOutput{
						RpcName:         "RPC_Chat.GetChatHistoryToOlder",
						UserParam:       params,
						CommandToServer: cmd,
						PBClassName:     "PB_ChatResponse_GetChatHistoryToOlder",
						ResponseData:    &res,
						RpcParamPassed:  load,
					}
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_ChatResponse_GetChatHistoryToOlder",cmd, params)
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_ChatResponse_GetChatHistoryToOlder","RPC_Chat.GetChatHistoryToOlder",cmd, params , load)
					responseHandler.HandleOfflineResult(out)
				} else {
					responseHandler.HandelError(err)
				}
			} else {
				responseHandler.HandelError(err)
			}
		default:
			noDevErr(errors.New("rpc method is does not exist: " + cmd.Command))
		}
	case "RPC_Other":

		//rpc,ok := rpcHandler.RPC_Other
		rpc := rpcHandler.RPC_Other
		/*if !ok {
		    e:=errors.New("rpcHandler could not be cast to : RPC_Other")
		    noDevErr(e)
		    RPC_ResponseHandler.HandelError(e)
		    return
		}*/

		switch splits[1] {
		case "Echo": //each pb_service_method
			load := &PB_OtherParam_Echo{}
			err := proto.Unmarshal(cmd.Data, load)
			if err == nil {
				res, err := rpc.Echo(load, params)
				if err == nil {
					out := RpcResponseOutput{
						RpcName:         "RPC_Other.Echo",
						UserParam:       params,
						CommandToServer: cmd,
						PBClassName:     "PB_OtherResponse_Echo",
						ResponseData:    &res,
						RpcParamPassed:  load,
					}
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_OtherResponse_Echo",cmd, params)
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_OtherResponse_Echo","RPC_Other.Echo",cmd, params , load)
					responseHandler.HandleOfflineResult(out)
				} else {
					responseHandler.HandelError(err)
				}
			} else {
				responseHandler.HandelError(err)
			}
		default:
			noDevErr(errors.New("rpc method is does not exist: " + cmd.Command))
		}
	case "RPC_Page":

		//rpc,ok := rpcHandler.RPC_Page
		rpc := rpcHandler.RPC_Page
		/*if !ok {
		    e:=errors.New("rpcHandler could not be cast to : RPC_Page")
		    noDevErr(e)
		    RPC_ResponseHandler.HandelError(e)
		    return
		}*/

		switch splits[1] {
		case "GetCommentsPage": //each pb_service_method
			load := &PB_PageParam_GetCommentsPage{}
			err := proto.Unmarshal(cmd.Data, load)
			if err == nil {
				res, err := rpc.GetCommentsPage(load, params)
				if err == nil {
					out := RpcResponseOutput{
						RpcName:         "RPC_Page.GetCommentsPage",
						UserParam:       params,
						CommandToServer: cmd,
						PBClassName:     "PB_PageResponse_GetCommentsPage",
						ResponseData:    &res,
						RpcParamPassed:  load,
					}
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_PageResponse_GetCommentsPage",cmd, params)
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_PageResponse_GetCommentsPage","RPC_Page.GetCommentsPage",cmd, params , load)
					responseHandler.HandleOfflineResult(out)
				} else {
					responseHandler.HandelError(err)
				}
			} else {
				responseHandler.HandelError(err)
			}
		case "GetHomePage": //each pb_service_method
			load := &PB_PageParam_GetHomePage{}
			err := proto.Unmarshal(cmd.Data, load)
			if err == nil {
				res, err := rpc.GetHomePage(load, params)
				if err == nil {
					out := RpcResponseOutput{
						RpcName:         "RPC_Page.GetHomePage",
						UserParam:       params,
						CommandToServer: cmd,
						PBClassName:     "PB_PageResponse_GetHomePage",
						ResponseData:    &res,
						RpcParamPassed:  load,
					}
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_PageResponse_GetHomePage",cmd, params)
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_PageResponse_GetHomePage","RPC_Page.GetHomePage",cmd, params , load)
					responseHandler.HandleOfflineResult(out)
				} else {
					responseHandler.HandelError(err)
				}
			} else {
				responseHandler.HandelError(err)
			}
		case "GetProfilePage": //each pb_service_method
			load := &PB_PageParam_GetProfilePage{}
			err := proto.Unmarshal(cmd.Data, load)
			if err == nil {
				res, err := rpc.GetProfilePage(load, params)
				if err == nil {
					out := RpcResponseOutput{
						RpcName:         "RPC_Page.GetProfilePage",
						UserParam:       params,
						CommandToServer: cmd,
						PBClassName:     "PB_PageResponse_GetProfilePage",
						ResponseData:    &res,
						RpcParamPassed:  load,
					}
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_PageResponse_GetProfilePage",cmd, params)
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_PageResponse_GetProfilePage","RPC_Page.GetProfilePage",cmd, params , load)
					responseHandler.HandleOfflineResult(out)
				} else {
					responseHandler.HandelError(err)
				}
			} else {
				responseHandler.HandelError(err)
			}
		case "GetLikesPage": //each pb_service_method
			load := &PB_PageParam_GetLikesPage{}
			err := proto.Unmarshal(cmd.Data, load)
			if err == nil {
				res, err := rpc.GetLikesPage(load, params)
				if err == nil {
					out := RpcResponseOutput{
						RpcName:         "RPC_Page.GetLikesPage",
						UserParam:       params,
						CommandToServer: cmd,
						PBClassName:     "PB_PageResponse_GetLikesPage",
						ResponseData:    &res,
						RpcParamPassed:  load,
					}
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_PageResponse_GetLikesPage",cmd, params)
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_PageResponse_GetLikesPage","RPC_Page.GetLikesPage",cmd, params , load)
					responseHandler.HandleOfflineResult(out)
				} else {
					responseHandler.HandelError(err)
				}
			} else {
				responseHandler.HandelError(err)
			}
		case "GetFollowersPage": //each pb_service_method
			load := &PB_PageParam_GetFollowersPage{}
			err := proto.Unmarshal(cmd.Data, load)
			if err == nil {
				res, err := rpc.GetFollowersPage(load, params)
				if err == nil {
					out := RpcResponseOutput{
						RpcName:         "RPC_Page.GetFollowersPage",
						UserParam:       params,
						CommandToServer: cmd,
						PBClassName:     "PB_PageResponse_GetFollowersPage",
						ResponseData:    &res,
						RpcParamPassed:  load,
					}
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_PageResponse_GetFollowersPage",cmd, params)
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_PageResponse_GetFollowersPage","RPC_Page.GetFollowersPage",cmd, params , load)
					responseHandler.HandleOfflineResult(out)
				} else {
					responseHandler.HandelError(err)
				}
			} else {
				responseHandler.HandelError(err)
			}
		case "GetFollowingsPage": //each pb_service_method
			load := &PB_PageParam_GetFollowingsPage{}
			err := proto.Unmarshal(cmd.Data, load)
			if err == nil {
				res, err := rpc.GetFollowingsPage(load, params)
				if err == nil {
					out := RpcResponseOutput{
						RpcName:         "RPC_Page.GetFollowingsPage",
						UserParam:       params,
						CommandToServer: cmd,
						PBClassName:     "PB_PageResponse_GetFollowingsPage",
						ResponseData:    &res,
						RpcParamPassed:  load,
					}
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_PageResponse_GetFollowingsPage",cmd, params)
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_PageResponse_GetFollowingsPage","RPC_Page.GetFollowingsPage",cmd, params , load)
					responseHandler.HandleOfflineResult(out)
				} else {
					responseHandler.HandelError(err)
				}
			} else {
				responseHandler.HandelError(err)
			}
		case "GetNotifiesPage": //each pb_service_method
			load := &PB_PageParam_GetNotifiesPage{}
			err := proto.Unmarshal(cmd.Data, load)
			if err == nil {
				res, err := rpc.GetNotifiesPage(load, params)
				if err == nil {
					out := RpcResponseOutput{
						RpcName:         "RPC_Page.GetNotifiesPage",
						UserParam:       params,
						CommandToServer: cmd,
						PBClassName:     "PB_PageResponse_GetNotifiesPage",
						ResponseData:    &res,
						RpcParamPassed:  load,
					}
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_PageResponse_GetNotifiesPage",cmd, params)
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_PageResponse_GetNotifiesPage","RPC_Page.GetNotifiesPage",cmd, params , load)
					responseHandler.HandleOfflineResult(out)
				} else {
					responseHandler.HandelError(err)
				}
			} else {
				responseHandler.HandelError(err)
			}
		case "GetUserActionsPage": //each pb_service_method
			load := &PB_PageParam_GetUserActionsPage{}
			err := proto.Unmarshal(cmd.Data, load)
			if err == nil {
				res, err := rpc.GetUserActionsPage(load, params)
				if err == nil {
					out := RpcResponseOutput{
						RpcName:         "RPC_Page.GetUserActionsPage",
						UserParam:       params,
						CommandToServer: cmd,
						PBClassName:     "PB_PageResponse_GetUserActionsPage",
						ResponseData:    &res,
						RpcParamPassed:  load,
					}
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_PageResponse_GetUserActionsPage",cmd, params)
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_PageResponse_GetUserActionsPage","RPC_Page.GetUserActionsPage",cmd, params , load)
					responseHandler.HandleOfflineResult(out)
				} else {
					responseHandler.HandelError(err)
				}
			} else {
				responseHandler.HandelError(err)
			}
		case "GetSuggestedPostsPage": //each pb_service_method
			load := &PB_PageParam_GetSuggestedPostsPage{}
			err := proto.Unmarshal(cmd.Data, load)
			if err == nil {
				res, err := rpc.GetSuggestedPostsPage(load, params)
				if err == nil {
					out := RpcResponseOutput{
						RpcName:         "RPC_Page.GetSuggestedPostsPage",
						UserParam:       params,
						CommandToServer: cmd,
						PBClassName:     "PB_PageResponse_GetSuggestedPostsPage",
						ResponseData:    &res,
						RpcParamPassed:  load,
					}
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_PageResponse_GetSuggestedPostsPage",cmd, params)
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_PageResponse_GetSuggestedPostsPage","RPC_Page.GetSuggestedPostsPage",cmd, params , load)
					responseHandler.HandleOfflineResult(out)
				} else {
					responseHandler.HandelError(err)
				}
			} else {
				responseHandler.HandelError(err)
			}
		case "GetSuggestedUsersPage": //each pb_service_method
			load := &PB_PageParam_GetSuggestedUsersPage{}
			err := proto.Unmarshal(cmd.Data, load)
			if err == nil {
				res, err := rpc.GetSuggestedUsersPage(load, params)
				if err == nil {
					out := RpcResponseOutput{
						RpcName:         "RPC_Page.GetSuggestedUsersPage",
						UserParam:       params,
						CommandToServer: cmd,
						PBClassName:     "PB_PageResponse_GetSuggestedUsersPage",
						ResponseData:    &res,
						RpcParamPassed:  load,
					}
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_PageResponse_GetSuggestedUsersPage",cmd, params)
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_PageResponse_GetSuggestedUsersPage","RPC_Page.GetSuggestedUsersPage",cmd, params , load)
					responseHandler.HandleOfflineResult(out)
				} else {
					responseHandler.HandelError(err)
				}
			} else {
				responseHandler.HandelError(err)
			}
		case "GetSuggestedTagsPage": //each pb_service_method
			load := &PB_PageParam_GetSuggestedTagsPage{}
			err := proto.Unmarshal(cmd.Data, load)
			if err == nil {
				res, err := rpc.GetSuggestedTagsPage(load, params)
				if err == nil {
					out := RpcResponseOutput{
						RpcName:         "RPC_Page.GetSuggestedTagsPage",
						UserParam:       params,
						CommandToServer: cmd,
						PBClassName:     "PB_PageResponse_GetSuggestedTagsPage",
						ResponseData:    &res,
						RpcParamPassed:  load,
					}
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_PageResponse_GetSuggestedTagsPage",cmd, params)
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_PageResponse_GetSuggestedTagsPage","RPC_Page.GetSuggestedTagsPage",cmd, params , load)
					responseHandler.HandleOfflineResult(out)
				} else {
					responseHandler.HandelError(err)
				}
			} else {
				responseHandler.HandelError(err)
			}
		case "GetLastPostsPage": //each pb_service_method
			load := &PB_PageParam_GetLastPostsPage{}
			err := proto.Unmarshal(cmd.Data, load)
			if err == nil {
				res, err := rpc.GetLastPostsPage(load, params)
				if err == nil {
					out := RpcResponseOutput{
						RpcName:         "RPC_Page.GetLastPostsPage",
						UserParam:       params,
						CommandToServer: cmd,
						PBClassName:     "PB_PageResponse_GetLastPostsPage",
						ResponseData:    &res,
						RpcParamPassed:  load,
					}
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_PageResponse_GetLastPostsPage",cmd, params)
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_PageResponse_GetLastPostsPage","RPC_Page.GetLastPostsPage",cmd, params , load)
					responseHandler.HandleOfflineResult(out)
				} else {
					responseHandler.HandelError(err)
				}
			} else {
				responseHandler.HandelError(err)
			}
		case "GetTagPage": //each pb_service_method
			load := &PB_PageParam_GetTagPage{}
			err := proto.Unmarshal(cmd.Data, load)
			if err == nil {
				res, err := rpc.GetTagPage(load, params)
				if err == nil {
					out := RpcResponseOutput{
						RpcName:         "RPC_Page.GetTagPage",
						UserParam:       params,
						CommandToServer: cmd,
						PBClassName:     "PB_PageResponse_GetTagPage",
						ResponseData:    &res,
						RpcParamPassed:  load,
					}
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_PageResponse_GetTagPage",cmd, params)
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_PageResponse_GetTagPage","RPC_Page.GetTagPage",cmd, params , load)
					responseHandler.HandleOfflineResult(out)
				} else {
					responseHandler.HandelError(err)
				}
			} else {
				responseHandler.HandelError(err)
			}
		case "SearchTagsPage": //each pb_service_method
			load := &PB_PageParam_SearchTagsPage{}
			err := proto.Unmarshal(cmd.Data, load)
			if err == nil {
				res, err := rpc.SearchTagsPage(load, params)
				if err == nil {
					out := RpcResponseOutput{
						RpcName:         "RPC_Page.SearchTagsPage",
						UserParam:       params,
						CommandToServer: cmd,
						PBClassName:     "PB_PageResponse_SearchTagsPage",
						ResponseData:    &res,
						RpcParamPassed:  load,
					}
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_PageResponse_SearchTagsPage",cmd, params)
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_PageResponse_SearchTagsPage","RPC_Page.SearchTagsPage",cmd, params , load)
					responseHandler.HandleOfflineResult(out)
				} else {
					responseHandler.HandelError(err)
				}
			} else {
				responseHandler.HandelError(err)
			}
		case "SearchUsersPage": //each pb_service_method
			load := &PB_PageParam_SearchUsersPage{}
			err := proto.Unmarshal(cmd.Data, load)
			if err == nil {
				res, err := rpc.SearchUsersPage(load, params)
				if err == nil {
					out := RpcResponseOutput{
						RpcName:         "RPC_Page.SearchUsersPage",
						UserParam:       params,
						CommandToServer: cmd,
						PBClassName:     "PB_PageResponse_SearchUsersPage",
						ResponseData:    &res,
						RpcParamPassed:  load,
					}
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_PageResponse_SearchUsersPage",cmd, params)
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_PageResponse_SearchUsersPage","RPC_Page.SearchUsersPage",cmd, params , load)
					responseHandler.HandleOfflineResult(out)
				} else {
					responseHandler.HandelError(err)
				}
			} else {
				responseHandler.HandelError(err)
			}
		default:
			noDevErr(errors.New("rpc method is does not exist: " + cmd.Command))
		}
	case "RPC_Search":

		//rpc,ok := rpcHandler.RPC_Search
		rpc := rpcHandler.RPC_Search
		/*if !ok {
		    e:=errors.New("rpcHandler could not be cast to : RPC_Search")
		    noDevErr(e)
		    RPC_ResponseHandler.HandelError(e)
		    return
		}*/

		switch splits[1] {
		case "SearchTags": //each pb_service_method
			load := &PB_SearchResponse_AddNewC{}
			err := proto.Unmarshal(cmd.Data, load)
			if err == nil {
				res, err := rpc.SearchTags(load, params)
				if err == nil {
					out := RpcResponseOutput{
						RpcName:         "RPC_Search.SearchTags",
						UserParam:       params,
						CommandToServer: cmd,
						PBClassName:     "PB_SearchResponse_AddNewC",
						ResponseData:    &res,
						RpcParamPassed:  load,
					}
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_SearchResponse_AddNewC",cmd, params)
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_SearchResponse_AddNewC","RPC_Search.SearchTags",cmd, params , load)
					responseHandler.HandleOfflineResult(out)
				} else {
					responseHandler.HandelError(err)
				}
			} else {
				responseHandler.HandelError(err)
			}
		default:
			noDevErr(errors.New("rpc method is does not exist: " + cmd.Command))
		}
	case "RPC_Social":

		//rpc,ok := rpcHandler.RPC_Social
		rpc := rpcHandler.RPC_Social
		/*if !ok {
		    e:=errors.New("rpcHandler could not be cast to : RPC_Social")
		    noDevErr(e)
		    RPC_ResponseHandler.HandelError(e)
		    return
		}*/

		switch splits[1] {
		case "AddComment": //each pb_service_method
			load := &PB_SocialParam_AddComment{}
			err := proto.Unmarshal(cmd.Data, load)
			if err == nil {
				res, err := rpc.AddComment(load, params)
				if err == nil {
					out := RpcResponseOutput{
						RpcName:         "RPC_Social.AddComment",
						UserParam:       params,
						CommandToServer: cmd,
						PBClassName:     "PB_SocialResponse_AddComment",
						ResponseData:    &res,
						RpcParamPassed:  load,
					}
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_SocialResponse_AddComment",cmd, params)
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_SocialResponse_AddComment","RPC_Social.AddComment",cmd, params , load)
					responseHandler.HandleOfflineResult(out)
				} else {
					responseHandler.HandelError(err)
				}
			} else {
				responseHandler.HandelError(err)
			}
		case "DeleteComment": //each pb_service_method
			load := &PB_SocialParam_DeleteComment{}
			err := proto.Unmarshal(cmd.Data, load)
			if err == nil {
				res, err := rpc.DeleteComment(load, params)
				if err == nil {
					out := RpcResponseOutput{
						RpcName:         "RPC_Social.DeleteComment",
						UserParam:       params,
						CommandToServer: cmd,
						PBClassName:     "PB_SocialResponse_DeleteComment",
						ResponseData:    &res,
						RpcParamPassed:  load,
					}
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_SocialResponse_DeleteComment",cmd, params)
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_SocialResponse_DeleteComment","RPC_Social.DeleteComment",cmd, params , load)
					responseHandler.HandleOfflineResult(out)
				} else {
					responseHandler.HandelError(err)
				}
			} else {
				responseHandler.HandelError(err)
			}
		case "AddPost": //each pb_service_method
			load := &PB_SocialParam_AddPost{}
			err := proto.Unmarshal(cmd.Data, load)
			if err == nil {
				res, err := rpc.AddPost(load, params)
				if err == nil {
					out := RpcResponseOutput{
						RpcName:         "RPC_Social.AddPost",
						UserParam:       params,
						CommandToServer: cmd,
						PBClassName:     "PB_SocialResponse_AddPost",
						ResponseData:    &res,
						RpcParamPassed:  load,
					}
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_SocialResponse_AddPost",cmd, params)
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_SocialResponse_AddPost","RPC_Social.AddPost",cmd, params , load)
					responseHandler.HandleOfflineResult(out)
				} else {
					responseHandler.HandelError(err)
				}
			} else {
				responseHandler.HandelError(err)
			}
		case "EditPost": //each pb_service_method
			load := &PB_SocialParam_EditPost{}
			err := proto.Unmarshal(cmd.Data, load)
			if err == nil {
				res, err := rpc.EditPost(load, params)
				if err == nil {
					out := RpcResponseOutput{
						RpcName:         "RPC_Social.EditPost",
						UserParam:       params,
						CommandToServer: cmd,
						PBClassName:     "PB_SocialResponse_EditPost",
						ResponseData:    &res,
						RpcParamPassed:  load,
					}
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_SocialResponse_EditPost",cmd, params)
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_SocialResponse_EditPost","RPC_Social.EditPost",cmd, params , load)
					responseHandler.HandleOfflineResult(out)
				} else {
					responseHandler.HandelError(err)
				}
			} else {
				responseHandler.HandelError(err)
			}
		case "DeletePost": //each pb_service_method
			load := &PB_SocialParam_DeletePost{}
			err := proto.Unmarshal(cmd.Data, load)
			if err == nil {
				res, err := rpc.DeletePost(load, params)
				if err == nil {
					out := RpcResponseOutput{
						RpcName:         "RPC_Social.DeletePost",
						UserParam:       params,
						CommandToServer: cmd,
						PBClassName:     "PB_SocialResponse_DeletePost",
						ResponseData:    &res,
						RpcParamPassed:  load,
					}
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_SocialResponse_DeletePost",cmd, params)
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_SocialResponse_DeletePost","RPC_Social.DeletePost",cmd, params , load)
					responseHandler.HandleOfflineResult(out)
				} else {
					responseHandler.HandelError(err)
				}
			} else {
				responseHandler.HandelError(err)
			}
		case "ArchivePost": //each pb_service_method
			load := &PB_SocialParam_ArchivePost{}
			err := proto.Unmarshal(cmd.Data, load)
			if err == nil {
				res, err := rpc.ArchivePost(load, params)
				if err == nil {
					out := RpcResponseOutput{
						RpcName:         "RPC_Social.ArchivePost",
						UserParam:       params,
						CommandToServer: cmd,
						PBClassName:     "PB_SocialResponse_ArchivePost",
						ResponseData:    &res,
						RpcParamPassed:  load,
					}
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_SocialResponse_ArchivePost",cmd, params)
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_SocialResponse_ArchivePost","RPC_Social.ArchivePost",cmd, params , load)
					responseHandler.HandleOfflineResult(out)
				} else {
					responseHandler.HandelError(err)
				}
			} else {
				responseHandler.HandelError(err)
			}
		case "LikePost": //each pb_service_method
			load := &PB_SocialParam_LikePost{}
			err := proto.Unmarshal(cmd.Data, load)
			if err == nil {
				res, err := rpc.LikePost(load, params)
				if err == nil {
					out := RpcResponseOutput{
						RpcName:         "RPC_Social.LikePost",
						UserParam:       params,
						CommandToServer: cmd,
						PBClassName:     "PB_SocialResponse_LikePost",
						ResponseData:    &res,
						RpcParamPassed:  load,
					}
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_SocialResponse_LikePost",cmd, params)
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_SocialResponse_LikePost","RPC_Social.LikePost",cmd, params , load)
					responseHandler.HandleOfflineResult(out)
				} else {
					responseHandler.HandelError(err)
				}
			} else {
				responseHandler.HandelError(err)
			}
		case "UnLikePost": //each pb_service_method
			load := &PB_SocialParam_UnLikePost{}
			err := proto.Unmarshal(cmd.Data, load)
			if err == nil {
				res, err := rpc.UnLikePost(load, params)
				if err == nil {
					out := RpcResponseOutput{
						RpcName:         "RPC_Social.UnLikePost",
						UserParam:       params,
						CommandToServer: cmd,
						PBClassName:     "PB_SocialResponse_UnLikePost",
						ResponseData:    &res,
						RpcParamPassed:  load,
					}
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_SocialResponse_UnLikePost",cmd, params)
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_SocialResponse_UnLikePost","RPC_Social.UnLikePost",cmd, params , load)
					responseHandler.HandleOfflineResult(out)
				} else {
					responseHandler.HandelError(err)
				}
			} else {
				responseHandler.HandelError(err)
			}
		case "FollowUser": //each pb_service_method
			load := &PB_SocialParam_FollowUser{}
			err := proto.Unmarshal(cmd.Data, load)
			if err == nil {
				res, err := rpc.FollowUser(load, params)
				if err == nil {
					out := RpcResponseOutput{
						RpcName:         "RPC_Social.FollowUser",
						UserParam:       params,
						CommandToServer: cmd,
						PBClassName:     "PB_SocialResponse_FollowUser",
						ResponseData:    &res,
						RpcParamPassed:  load,
					}
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_SocialResponse_FollowUser",cmd, params)
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_SocialResponse_FollowUser","RPC_Social.FollowUser",cmd, params , load)
					responseHandler.HandleOfflineResult(out)
				} else {
					responseHandler.HandelError(err)
				}
			} else {
				responseHandler.HandelError(err)
			}
		case "UnFollowUser": //each pb_service_method
			load := &PB_SocialParam_UnFollowUser{}
			err := proto.Unmarshal(cmd.Data, load)
			if err == nil {
				res, err := rpc.UnFollowUser(load, params)
				if err == nil {
					out := RpcResponseOutput{
						RpcName:         "RPC_Social.UnFollowUser",
						UserParam:       params,
						CommandToServer: cmd,
						PBClassName:     "PB_SocialResponse_UnFollowUser",
						ResponseData:    &res,
						RpcParamPassed:  load,
					}
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_SocialResponse_UnFollowUser",cmd, params)
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_SocialResponse_UnFollowUser","RPC_Social.UnFollowUser",cmd, params , load)
					responseHandler.HandleOfflineResult(out)
				} else {
					responseHandler.HandelError(err)
				}
			} else {
				responseHandler.HandelError(err)
			}
		default:
			noDevErr(errors.New("rpc method is does not exist: " + cmd.Command))
		}
	case "RPC_User":

		//rpc,ok := rpcHandler.RPC_User
		rpc := rpcHandler.RPC_User
		/*if !ok {
		    e:=errors.New("rpcHandler could not be cast to : RPC_User")
		    noDevErr(e)
		    RPC_ResponseHandler.HandelError(e)
		    return
		}*/

		switch splits[1] {
		case "BlockUser": //each pb_service_method
			load := &PB_UserParam_BlockUser{}
			err := proto.Unmarshal(cmd.Data, load)
			if err == nil {
				res, err := rpc.BlockUser(load, params)
				if err == nil {
					out := RpcResponseOutput{
						RpcName:         "RPC_User.BlockUser",
						UserParam:       params,
						CommandToServer: cmd,
						PBClassName:     "PB_UserResponse_BlockUser",
						ResponseData:    &res,
						RpcParamPassed:  load,
					}
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_UserResponse_BlockUser",cmd, params)
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_UserResponse_BlockUser","RPC_User.BlockUser",cmd, params , load)
					responseHandler.HandleOfflineResult(out)
				} else {
					responseHandler.HandelError(err)
				}
			} else {
				responseHandler.HandelError(err)
			}
		case "UnBlockUser": //each pb_service_method
			load := &PB_UserParam_UnBlockUser{}
			err := proto.Unmarshal(cmd.Data, load)
			if err == nil {
				res, err := rpc.UnBlockUser(load, params)
				if err == nil {
					out := RpcResponseOutput{
						RpcName:         "RPC_User.UnBlockUser",
						UserParam:       params,
						CommandToServer: cmd,
						PBClassName:     "PB_UserResponse_UnBlockUser",
						ResponseData:    &res,
						RpcParamPassed:  load,
					}
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_UserResponse_UnBlockUser",cmd, params)
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_UserResponse_UnBlockUser","RPC_User.UnBlockUser",cmd, params , load)
					responseHandler.HandleOfflineResult(out)
				} else {
					responseHandler.HandelError(err)
				}
			} else {
				responseHandler.HandelError(err)
			}
		case "GetBlockedList": //each pb_service_method
			load := &PB_UserParam_BlockedList{}
			err := proto.Unmarshal(cmd.Data, load)
			if err == nil {
				res, err := rpc.GetBlockedList(load, params)
				if err == nil {
					out := RpcResponseOutput{
						RpcName:         "RPC_User.GetBlockedList",
						UserParam:       params,
						CommandToServer: cmd,
						PBClassName:     "PB_UserResponse_BlockedList",
						ResponseData:    &res,
						RpcParamPassed:  load,
					}
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_UserResponse_BlockedList",cmd, params)
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_UserResponse_BlockedList","RPC_User.GetBlockedList",cmd, params , load)
					responseHandler.HandleOfflineResult(out)
				} else {
					responseHandler.HandelError(err)
				}
			} else {
				responseHandler.HandelError(err)
			}
		case "UpdateAbout": //each pb_service_method
			load := &PB_UserParam_UpdateAbout{}
			err := proto.Unmarshal(cmd.Data, load)
			if err == nil {
				res, err := rpc.UpdateAbout(load, params)
				if err == nil {
					out := RpcResponseOutput{
						RpcName:         "RPC_User.UpdateAbout",
						UserParam:       params,
						CommandToServer: cmd,
						PBClassName:     "PB_UserResponse_UpdateAbout",
						ResponseData:    &res,
						RpcParamPassed:  load,
					}
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_UserResponse_UpdateAbout",cmd, params)
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_UserResponse_UpdateAbout","RPC_User.UpdateAbout",cmd, params , load)
					responseHandler.HandleOfflineResult(out)
				} else {
					responseHandler.HandelError(err)
				}
			} else {
				responseHandler.HandelError(err)
			}
		case "UpdateUserName": //each pb_service_method
			load := &PB_UserParam_UpdateUserName{}
			err := proto.Unmarshal(cmd.Data, load)
			if err == nil {
				res, err := rpc.UpdateUserName(load, params)
				if err == nil {
					out := RpcResponseOutput{
						RpcName:         "RPC_User.UpdateUserName",
						UserParam:       params,
						CommandToServer: cmd,
						PBClassName:     "PB_UserResponse_UpdateUserName",
						ResponseData:    &res,
						RpcParamPassed:  load,
					}
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_UserResponse_UpdateUserName",cmd, params)
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_UserResponse_UpdateUserName","RPC_User.UpdateUserName",cmd, params , load)
					responseHandler.HandleOfflineResult(out)
				} else {
					responseHandler.HandelError(err)
				}
			} else {
				responseHandler.HandelError(err)
			}
		case "ChangePrivacy": //each pb_service_method
			load := &PB_UserParam_ChangePrivacy{}
			err := proto.Unmarshal(cmd.Data, load)
			if err == nil {
				res, err := rpc.ChangePrivacy(load, params)
				if err == nil {
					out := RpcResponseOutput{
						RpcName:         "RPC_User.ChangePrivacy",
						UserParam:       params,
						CommandToServer: cmd,
						PBClassName:     "PB_UserResponseOffline_ChangePrivacy",
						ResponseData:    &res,
						RpcParamPassed:  load,
					}
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_UserResponseOffline_ChangePrivacy",cmd, params)
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_UserResponseOffline_ChangePrivacy","RPC_User.ChangePrivacy",cmd, params , load)
					responseHandler.HandleOfflineResult(out)
				} else {
					responseHandler.HandelError(err)
				}
			} else {
				responseHandler.HandelError(err)
			}
		case "ChangeAvatar": //each pb_service_method
			load := &PB_UserParam_ChangeAvatar{}
			err := proto.Unmarshal(cmd.Data, load)
			if err == nil {
				res, err := rpc.ChangeAvatar(load, params)
				if err == nil {
					out := RpcResponseOutput{
						RpcName:         "RPC_User.ChangeAvatar",
						UserParam:       params,
						CommandToServer: cmd,
						PBClassName:     "PB_UserResponse_ChangeAvatar",
						ResponseData:    &res,
						RpcParamPassed:  load,
					}
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_UserResponse_ChangeAvatar",cmd, params)
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_UserResponse_ChangeAvatar","RPC_User.ChangeAvatar",cmd, params , load)
					responseHandler.HandleOfflineResult(out)
				} else {
					responseHandler.HandelError(err)
				}
			} else {
				responseHandler.HandelError(err)
			}
		case "CheckUserName": //each pb_service_method
			load := &PB_UserParam_CheckUserName{}
			err := proto.Unmarshal(cmd.Data, load)
			if err == nil {
				res, err := rpc.CheckUserName(load, params)
				if err == nil {
					out := RpcResponseOutput{
						RpcName:         "RPC_User.CheckUserName",
						UserParam:       params,
						CommandToServer: cmd,
						PBClassName:     "PB_UserResponse_CheckUserName",
						ResponseData:    &res,
						RpcParamPassed:  load,
					}
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_UserResponse_CheckUserName",cmd, params)
					//RPC_ResponseHandler.HandleOfflineResult(res,"PB_UserResponse_CheckUserName","RPC_User.CheckUserName",cmd, params , load)
					responseHandler.HandleOfflineResult(out)
				} else {
					responseHandler.HandelError(err)
				}
			} else {
				responseHandler.HandelError(err)
			}
		default:
			noDevErr(errors.New("rpc method is does not exist: " + cmd.Command))
		}
	default:
		noDevErr(errors.New("rpc dosent exisit for: " + cmd.Command))
	}
}

/////////////// Direct in PB_CommandToClient /////////////
/*


 RPC_Auth.CheckPhone
 RPC_Auth.SendCode
 RPC_Auth.SendCodeToSms
 RPC_Auth.SendCodeToTelgram
 RPC_Auth.SingUp
 RPC_Auth.SingIn
 RPC_Auth.LogOut



 RPC_Chat.AddNewMessage
 RPC_Chat.SetRoomActionDoing
 RPC_Chat.SetMessagesAsReceived
 RPC_Chat.SetMessagesRangeAsSeen
 RPC_Chat.DeleteChatHistory
 RPC_Chat.DeleteMessagesByIds
 RPC_Chat.EditMessage
 RPC_Chat.GetChatList
 RPC_Chat.GetChatHistoryToOlder



 RPC_Other.Echo



 RPC_Page.GetCommentsPage
 RPC_Page.GetHomePage
 RPC_Page.GetProfilePage
 RPC_Page.GetLikesPage
 RPC_Page.GetFollowersPage
 RPC_Page.GetFollowingsPage
 RPC_Page.GetNotifiesPage
 RPC_Page.GetUserActionsPage
 RPC_Page.GetSuggestedPostsPage
 RPC_Page.GetSuggestedUsersPage
 RPC_Page.GetSuggestedTagsPage
 RPC_Page.GetLastPostsPage
 RPC_Page.GetTagPage
 RPC_Page.SearchTagsPage
 RPC_Page.SearchUsersPage



 RPC_Search.SearchTags



 RPC_Social.AddComment
 RPC_Social.DeleteComment
 RPC_Social.AddPost
 RPC_Social.EditPost
 RPC_Social.DeletePost
 RPC_Social.ArchivePost
 RPC_Social.LikePost
 RPC_Social.UnLikePost
 RPC_Social.FollowUser
 RPC_Social.UnFollowUser



 RPC_User.BlockUser
 RPC_User.UnBlockUser
 RPC_User.GetBlockedList
 RPC_User.UpdateAbout
 RPC_User.UpdateUserName
 RPC_User.ChangePrivacy
 RPC_User.ChangeAvatar
 RPC_User.CheckUserName


*/
