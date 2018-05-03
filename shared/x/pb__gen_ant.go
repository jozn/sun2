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

type rpcParamHandler struct {
	cmd             PB_CommandToServer
	params          RPC_UserParam
	rpcHandler      RPC_AllHandlersInteract
	responseHandler RPC_ResponseHandlerInterface
}

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

	// rpc: RPC_Auth

	"RPC_Auth.CheckPhone": func(p rpcParamHandler) {
		if p.rpcHandler.RPC_Auth == nil {
			noDevErr(errors.New("rpc service is null for: p.rpcHandler.RPC_Auth.CheckPhone"))
			return
		}
		load := &PB_UserParam_CheckUserName2{}
		err := proto.Unmarshal(p.cmd.Data, load)
		if err == nil {
			res, err := p.rpcHandler.RPC_Auth.CheckPhone(load, p.params)
			if err == nil {
				out := RpcResponseOutput{
					RpcName:         "RPC_Auth.CheckPhone",
					UserParam:       p.params,
					CommandToServer: p.cmd,
					PBClassName:     "PB_UserResponse_CheckUserName2",
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
	"RPC_Auth.SendCode": func(p rpcParamHandler) {
		if p.rpcHandler.RPC_Auth == nil {
			noDevErr(errors.New("rpc service is null for: p.rpcHandler.RPC_Auth.SendCode"))
			return
		}
		load := &PB_UserParam_CheckUserName2{}
		err := proto.Unmarshal(p.cmd.Data, load)
		if err == nil {
			res, err := p.rpcHandler.RPC_Auth.SendCode(load, p.params)
			if err == nil {
				out := RpcResponseOutput{
					RpcName:         "RPC_Auth.SendCode",
					UserParam:       p.params,
					CommandToServer: p.cmd,
					PBClassName:     "PB_UserResponse_CheckUserName2",
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
	"RPC_Auth.SendCodeToSms": func(p rpcParamHandler) {
		if p.rpcHandler.RPC_Auth == nil {
			noDevErr(errors.New("rpc service is null for: p.rpcHandler.RPC_Auth.SendCodeToSms"))
			return
		}
		load := &PB_UserParam_CheckUserName2{}
		err := proto.Unmarshal(p.cmd.Data, load)
		if err == nil {
			res, err := p.rpcHandler.RPC_Auth.SendCodeToSms(load, p.params)
			if err == nil {
				out := RpcResponseOutput{
					RpcName:         "RPC_Auth.SendCodeToSms",
					UserParam:       p.params,
					CommandToServer: p.cmd,
					PBClassName:     "PB_UserResponse_CheckUserName2",
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
	"RPC_Auth.SendCodeToTelgram": func(p rpcParamHandler) {
		if p.rpcHandler.RPC_Auth == nil {
			noDevErr(errors.New("rpc service is null for: p.rpcHandler.RPC_Auth.SendCodeToTelgram"))
			return
		}
		load := &PB_UserParam_CheckUserName2{}
		err := proto.Unmarshal(p.cmd.Data, load)
		if err == nil {
			res, err := p.rpcHandler.RPC_Auth.SendCodeToTelgram(load, p.params)
			if err == nil {
				out := RpcResponseOutput{
					RpcName:         "RPC_Auth.SendCodeToTelgram",
					UserParam:       p.params,
					CommandToServer: p.cmd,
					PBClassName:     "PB_UserResponse_CheckUserName2",
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
	"RPC_Auth.SingUp": func(p rpcParamHandler) {
		if p.rpcHandler.RPC_Auth == nil {
			noDevErr(errors.New("rpc service is null for: p.rpcHandler.RPC_Auth.SingUp"))
			return
		}
		load := &PB_UserParam_CheckUserName2{}
		err := proto.Unmarshal(p.cmd.Data, load)
		if err == nil {
			res, err := p.rpcHandler.RPC_Auth.SingUp(load, p.params)
			if err == nil {
				out := RpcResponseOutput{
					RpcName:         "RPC_Auth.SingUp",
					UserParam:       p.params,
					CommandToServer: p.cmd,
					PBClassName:     "PB_UserResponse_CheckUserName2",
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
	"RPC_Auth.SingIn": func(p rpcParamHandler) {
		if p.rpcHandler.RPC_Auth == nil {
			noDevErr(errors.New("rpc service is null for: p.rpcHandler.RPC_Auth.SingIn"))
			return
		}
		load := &PB_UserParam_CheckUserName2{}
		err := proto.Unmarshal(p.cmd.Data, load)
		if err == nil {
			res, err := p.rpcHandler.RPC_Auth.SingIn(load, p.params)
			if err == nil {
				out := RpcResponseOutput{
					RpcName:         "RPC_Auth.SingIn",
					UserParam:       p.params,
					CommandToServer: p.cmd,
					PBClassName:     "PB_UserResponse_CheckUserName2",
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
	"RPC_Auth.LogOut": func(p rpcParamHandler) {
		if p.rpcHandler.RPC_Auth == nil {
			noDevErr(errors.New("rpc service is null for: p.rpcHandler.RPC_Auth.LogOut"))
			return
		}
		load := &PB_UserParam_CheckUserName2{}
		err := proto.Unmarshal(p.cmd.Data, load)
		if err == nil {
			res, err := p.rpcHandler.RPC_Auth.LogOut(load, p.params)
			if err == nil {
				out := RpcResponseOutput{
					RpcName:         "RPC_Auth.LogOut",
					UserParam:       p.params,
					CommandToServer: p.cmd,
					PBClassName:     "PB_UserResponse_CheckUserName2",
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

	// rpc: RPC_Chat

	"RPC_Chat.AddNewMessage": func(p rpcParamHandler) {
		if p.rpcHandler.RPC_Chat == nil {
			noDevErr(errors.New("rpc service is null for: p.rpcHandler.RPC_Chat.AddNewMessage"))
			return
		}
		load := &PB_ChatParam_AddNewMessage{}
		err := proto.Unmarshal(p.cmd.Data, load)
		if err == nil {
			res, err := p.rpcHandler.RPC_Chat.AddNewMessage(load, p.params)
			if err == nil {
				out := RpcResponseOutput{
					RpcName:         "RPC_Chat.AddNewMessage",
					UserParam:       p.params,
					CommandToServer: p.cmd,
					PBClassName:     "PB_ChatResponse_AddNewMessage",
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
		load := &PB_ChatParam_SetRoomActionDoing{}
		err := proto.Unmarshal(p.cmd.Data, load)
		if err == nil {
			res, err := p.rpcHandler.RPC_Chat.SetRoomActionDoing(load, p.params)
			if err == nil {
				out := RpcResponseOutput{
					RpcName:         "RPC_Chat.SetRoomActionDoing",
					UserParam:       p.params,
					CommandToServer: p.cmd,
					PBClassName:     "PB_ChatResponse_SetRoomActionDoing",
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
	"RPC_Chat.SetMessagesAsReceived": func(p rpcParamHandler) {
		if p.rpcHandler.RPC_Chat == nil {
			noDevErr(errors.New("rpc service is null for: p.rpcHandler.RPC_Chat.SetMessagesAsReceived"))
			return
		}
		load := &PB_ChatParam_SetMessagesAsReceived{}
		err := proto.Unmarshal(p.cmd.Data, load)
		if err == nil {
			res, err := p.rpcHandler.RPC_Chat.SetMessagesAsReceived(load, p.params)
			if err == nil {
				out := RpcResponseOutput{
					RpcName:         "RPC_Chat.SetMessagesAsReceived",
					UserParam:       p.params,
					CommandToServer: p.cmd,
					PBClassName:     "PB_ChatResponse_SetMessagesAsReceived",
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
	"RPC_Chat.SetMessagesRangeAsSeen": func(p rpcParamHandler) {
		if p.rpcHandler.RPC_Chat == nil {
			noDevErr(errors.New("rpc service is null for: p.rpcHandler.RPC_Chat.SetMessagesRangeAsSeen"))
			return
		}
		load := &PB_ChatParam_SetChatMessagesRangeAsSeen{}
		err := proto.Unmarshal(p.cmd.Data, load)
		if err == nil {
			res, err := p.rpcHandler.RPC_Chat.SetMessagesRangeAsSeen(load, p.params)
			if err == nil {
				out := RpcResponseOutput{
					RpcName:         "RPC_Chat.SetMessagesRangeAsSeen",
					UserParam:       p.params,
					CommandToServer: p.cmd,
					PBClassName:     "PB_ChatResponse_SetChatMessagesRangeAsSeen",
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
	"RPC_Chat.DeleteChatHistory": func(p rpcParamHandler) {
		if p.rpcHandler.RPC_Chat == nil {
			noDevErr(errors.New("rpc service is null for: p.rpcHandler.RPC_Chat.DeleteChatHistory"))
			return
		}
		load := &PB_ChatParam_DeleteChatHistory{}
		err := proto.Unmarshal(p.cmd.Data, load)
		if err == nil {
			res, err := p.rpcHandler.RPC_Chat.DeleteChatHistory(load, p.params)
			if err == nil {
				out := RpcResponseOutput{
					RpcName:         "RPC_Chat.DeleteChatHistory",
					UserParam:       p.params,
					CommandToServer: p.cmd,
					PBClassName:     "PB_ChatResponse_DeleteChatHistory",
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
	"RPC_Chat.DeleteMessagesByIds": func(p rpcParamHandler) {
		if p.rpcHandler.RPC_Chat == nil {
			noDevErr(errors.New("rpc service is null for: p.rpcHandler.RPC_Chat.DeleteMessagesByIds"))
			return
		}
		load := &PB_ChatParam_DeleteMessagesByIds{}
		err := proto.Unmarshal(p.cmd.Data, load)
		if err == nil {
			res, err := p.rpcHandler.RPC_Chat.DeleteMessagesByIds(load, p.params)
			if err == nil {
				out := RpcResponseOutput{
					RpcName:         "RPC_Chat.DeleteMessagesByIds",
					UserParam:       p.params,
					CommandToServer: p.cmd,
					PBClassName:     "PB_ChatResponse_DeleteMessagesByIds",
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
	"RPC_Chat.EditMessage": func(p rpcParamHandler) {
		if p.rpcHandler.RPC_Chat == nil {
			noDevErr(errors.New("rpc service is null for: p.rpcHandler.RPC_Chat.EditMessage"))
			return
		}
		load := &PB_ChatParam_EditMessage{}
		err := proto.Unmarshal(p.cmd.Data, load)
		if err == nil {
			res, err := p.rpcHandler.RPC_Chat.EditMessage(load, p.params)
			if err == nil {
				out := RpcResponseOutput{
					RpcName:         "RPC_Chat.EditMessage",
					UserParam:       p.params,
					CommandToServer: p.cmd,
					PBClassName:     "PB_ChatResponse_EditMessage",
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
		load := &PB_ChatParam_GetChatList{}
		err := proto.Unmarshal(p.cmd.Data, load)
		if err == nil {
			res, err := p.rpcHandler.RPC_Chat.GetChatList(load, p.params)
			if err == nil {
				out := RpcResponseOutput{
					RpcName:         "RPC_Chat.GetChatList",
					UserParam:       p.params,
					CommandToServer: p.cmd,
					PBClassName:     "PB_ChatResponse_GetChatList",
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
	"RPC_Chat.GetChatHistoryToOlder": func(p rpcParamHandler) {
		if p.rpcHandler.RPC_Chat == nil {
			noDevErr(errors.New("rpc service is null for: p.rpcHandler.RPC_Chat.GetChatHistoryToOlder"))
			return
		}
		load := &PB_ChatParam_GetChatHistoryToOlder{}
		err := proto.Unmarshal(p.cmd.Data, load)
		if err == nil {
			res, err := p.rpcHandler.RPC_Chat.GetChatHistoryToOlder(load, p.params)
			if err == nil {
				out := RpcResponseOutput{
					RpcName:         "RPC_Chat.GetChatHistoryToOlder",
					UserParam:       p.params,
					CommandToServer: p.cmd,
					PBClassName:     "PB_ChatResponse_GetChatHistoryToOlder",
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

	// rpc: RPC_Other

	"RPC_Other.Echo": func(p rpcParamHandler) {
		if p.rpcHandler.RPC_Other == nil {
			noDevErr(errors.New("rpc service is null for: p.rpcHandler.RPC_Other.Echo"))
			return
		}
		load := &PB_OtherParam_Echo{}
		err := proto.Unmarshal(p.cmd.Data, load)
		if err == nil {
			res, err := p.rpcHandler.RPC_Other.Echo(load, p.params)
			if err == nil {
				out := RpcResponseOutput{
					RpcName:         "RPC_Other.Echo",
					UserParam:       p.params,
					CommandToServer: p.cmd,
					PBClassName:     "PB_OtherResponse_Echo",
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

	// rpc: RPC_Page

	"RPC_Page.GetCommentsPage": func(p rpcParamHandler) {
		if p.rpcHandler.RPC_Page == nil {
			noDevErr(errors.New("rpc service is null for: p.rpcHandler.RPC_Page.GetCommentsPage"))
			return
		}
		load := &PB_PageParam_GetCommentsPage{}
		err := proto.Unmarshal(p.cmd.Data, load)
		if err == nil {
			res, err := p.rpcHandler.RPC_Page.GetCommentsPage(load, p.params)
			if err == nil {
				out := RpcResponseOutput{
					RpcName:         "RPC_Page.GetCommentsPage",
					UserParam:       p.params,
					CommandToServer: p.cmd,
					PBClassName:     "PB_PageResponse_GetCommentsPage",
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
	"RPC_Page.GetHomePage": func(p rpcParamHandler) {
		if p.rpcHandler.RPC_Page == nil {
			noDevErr(errors.New("rpc service is null for: p.rpcHandler.RPC_Page.GetHomePage"))
			return
		}
		load := &PB_PageParam_GetHomePage{}
		err := proto.Unmarshal(p.cmd.Data, load)
		if err == nil {
			res, err := p.rpcHandler.RPC_Page.GetHomePage(load, p.params)
			if err == nil {
				out := RpcResponseOutput{
					RpcName:         "RPC_Page.GetHomePage",
					UserParam:       p.params,
					CommandToServer: p.cmd,
					PBClassName:     "PB_PageResponse_GetHomePage",
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
	"RPC_Page.GetProfilePage": func(p rpcParamHandler) {
		if p.rpcHandler.RPC_Page == nil {
			noDevErr(errors.New("rpc service is null for: p.rpcHandler.RPC_Page.GetProfilePage"))
			return
		}
		load := &PB_PageParam_GetProfilePage{}
		err := proto.Unmarshal(p.cmd.Data, load)
		if err == nil {
			res, err := p.rpcHandler.RPC_Page.GetProfilePage(load, p.params)
			if err == nil {
				out := RpcResponseOutput{
					RpcName:         "RPC_Page.GetProfilePage",
					UserParam:       p.params,
					CommandToServer: p.cmd,
					PBClassName:     "PB_PageResponse_GetProfilePage",
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
	"RPC_Page.GetLikesPage": func(p rpcParamHandler) {
		if p.rpcHandler.RPC_Page == nil {
			noDevErr(errors.New("rpc service is null for: p.rpcHandler.RPC_Page.GetLikesPage"))
			return
		}
		load := &PB_PageParam_GetLikesPage{}
		err := proto.Unmarshal(p.cmd.Data, load)
		if err == nil {
			res, err := p.rpcHandler.RPC_Page.GetLikesPage(load, p.params)
			if err == nil {
				out := RpcResponseOutput{
					RpcName:         "RPC_Page.GetLikesPage",
					UserParam:       p.params,
					CommandToServer: p.cmd,
					PBClassName:     "PB_PageResponse_GetLikesPage",
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
	"RPC_Page.GetFollowersPage": func(p rpcParamHandler) {
		if p.rpcHandler.RPC_Page == nil {
			noDevErr(errors.New("rpc service is null for: p.rpcHandler.RPC_Page.GetFollowersPage"))
			return
		}
		load := &PB_PageParam_GetFollowersPage{}
		err := proto.Unmarshal(p.cmd.Data, load)
		if err == nil {
			res, err := p.rpcHandler.RPC_Page.GetFollowersPage(load, p.params)
			if err == nil {
				out := RpcResponseOutput{
					RpcName:         "RPC_Page.GetFollowersPage",
					UserParam:       p.params,
					CommandToServer: p.cmd,
					PBClassName:     "PB_PageResponse_GetFollowersPage",
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
	"RPC_Page.GetFollowingsPage": func(p rpcParamHandler) {
		if p.rpcHandler.RPC_Page == nil {
			noDevErr(errors.New("rpc service is null for: p.rpcHandler.RPC_Page.GetFollowingsPage"))
			return
		}
		load := &PB_PageParam_GetFollowingsPage{}
		err := proto.Unmarshal(p.cmd.Data, load)
		if err == nil {
			res, err := p.rpcHandler.RPC_Page.GetFollowingsPage(load, p.params)
			if err == nil {
				out := RpcResponseOutput{
					RpcName:         "RPC_Page.GetFollowingsPage",
					UserParam:       p.params,
					CommandToServer: p.cmd,
					PBClassName:     "PB_PageResponse_GetFollowingsPage",
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
	"RPC_Page.GetNotifiesPage": func(p rpcParamHandler) {
		if p.rpcHandler.RPC_Page == nil {
			noDevErr(errors.New("rpc service is null for: p.rpcHandler.RPC_Page.GetNotifiesPage"))
			return
		}
		load := &PB_PageParam_GetNotifiesPage{}
		err := proto.Unmarshal(p.cmd.Data, load)
		if err == nil {
			res, err := p.rpcHandler.RPC_Page.GetNotifiesPage(load, p.params)
			if err == nil {
				out := RpcResponseOutput{
					RpcName:         "RPC_Page.GetNotifiesPage",
					UserParam:       p.params,
					CommandToServer: p.cmd,
					PBClassName:     "PB_PageResponse_GetNotifiesPage",
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
	"RPC_Page.GetUserActionsPage": func(p rpcParamHandler) {
		if p.rpcHandler.RPC_Page == nil {
			noDevErr(errors.New("rpc service is null for: p.rpcHandler.RPC_Page.GetUserActionsPage"))
			return
		}
		load := &PB_PageParam_GetUserActionsPage{}
		err := proto.Unmarshal(p.cmd.Data, load)
		if err == nil {
			res, err := p.rpcHandler.RPC_Page.GetUserActionsPage(load, p.params)
			if err == nil {
				out := RpcResponseOutput{
					RpcName:         "RPC_Page.GetUserActionsPage",
					UserParam:       p.params,
					CommandToServer: p.cmd,
					PBClassName:     "PB_PageResponse_GetUserActionsPage",
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
	"RPC_Page.GetSuggestedPostsPage": func(p rpcParamHandler) {
		if p.rpcHandler.RPC_Page == nil {
			noDevErr(errors.New("rpc service is null for: p.rpcHandler.RPC_Page.GetSuggestedPostsPage"))
			return
		}
		load := &PB_PageParam_GetSuggestedPostsPage{}
		err := proto.Unmarshal(p.cmd.Data, load)
		if err == nil {
			res, err := p.rpcHandler.RPC_Page.GetSuggestedPostsPage(load, p.params)
			if err == nil {
				out := RpcResponseOutput{
					RpcName:         "RPC_Page.GetSuggestedPostsPage",
					UserParam:       p.params,
					CommandToServer: p.cmd,
					PBClassName:     "PB_PageResponse_GetSuggestedPostsPage",
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
	"RPC_Page.GetSuggestedUsersPage": func(p rpcParamHandler) {
		if p.rpcHandler.RPC_Page == nil {
			noDevErr(errors.New("rpc service is null for: p.rpcHandler.RPC_Page.GetSuggestedUsersPage"))
			return
		}
		load := &PB_PageParam_GetSuggestedUsersPage{}
		err := proto.Unmarshal(p.cmd.Data, load)
		if err == nil {
			res, err := p.rpcHandler.RPC_Page.GetSuggestedUsersPage(load, p.params)
			if err == nil {
				out := RpcResponseOutput{
					RpcName:         "RPC_Page.GetSuggestedUsersPage",
					UserParam:       p.params,
					CommandToServer: p.cmd,
					PBClassName:     "PB_PageResponse_GetSuggestedUsersPage",
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
	"RPC_Page.GetSuggestedTagsPage": func(p rpcParamHandler) {
		if p.rpcHandler.RPC_Page == nil {
			noDevErr(errors.New("rpc service is null for: p.rpcHandler.RPC_Page.GetSuggestedTagsPage"))
			return
		}
		load := &PB_PageParam_GetSuggestedTagsPage{}
		err := proto.Unmarshal(p.cmd.Data, load)
		if err == nil {
			res, err := p.rpcHandler.RPC_Page.GetSuggestedTagsPage(load, p.params)
			if err == nil {
				out := RpcResponseOutput{
					RpcName:         "RPC_Page.GetSuggestedTagsPage",
					UserParam:       p.params,
					CommandToServer: p.cmd,
					PBClassName:     "PB_PageResponse_GetSuggestedTagsPage",
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
	"RPC_Page.GetLastPostsPage": func(p rpcParamHandler) {
		if p.rpcHandler.RPC_Page == nil {
			noDevErr(errors.New("rpc service is null for: p.rpcHandler.RPC_Page.GetLastPostsPage"))
			return
		}
		load := &PB_PageParam_GetLastPostsPage{}
		err := proto.Unmarshal(p.cmd.Data, load)
		if err == nil {
			res, err := p.rpcHandler.RPC_Page.GetLastPostsPage(load, p.params)
			if err == nil {
				out := RpcResponseOutput{
					RpcName:         "RPC_Page.GetLastPostsPage",
					UserParam:       p.params,
					CommandToServer: p.cmd,
					PBClassName:     "PB_PageResponse_GetLastPostsPage",
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
	"RPC_Page.GetTagPage": func(p rpcParamHandler) {
		if p.rpcHandler.RPC_Page == nil {
			noDevErr(errors.New("rpc service is null for: p.rpcHandler.RPC_Page.GetTagPage"))
			return
		}
		load := &PB_PageParam_GetTagPage{}
		err := proto.Unmarshal(p.cmd.Data, load)
		if err == nil {
			res, err := p.rpcHandler.RPC_Page.GetTagPage(load, p.params)
			if err == nil {
				out := RpcResponseOutput{
					RpcName:         "RPC_Page.GetTagPage",
					UserParam:       p.params,
					CommandToServer: p.cmd,
					PBClassName:     "PB_PageResponse_GetTagPage",
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
	"RPC_Page.SearchTagsPage": func(p rpcParamHandler) {
		if p.rpcHandler.RPC_Page == nil {
			noDevErr(errors.New("rpc service is null for: p.rpcHandler.RPC_Page.SearchTagsPage"))
			return
		}
		load := &PB_PageParam_SearchTagsPage{}
		err := proto.Unmarshal(p.cmd.Data, load)
		if err == nil {
			res, err := p.rpcHandler.RPC_Page.SearchTagsPage(load, p.params)
			if err == nil {
				out := RpcResponseOutput{
					RpcName:         "RPC_Page.SearchTagsPage",
					UserParam:       p.params,
					CommandToServer: p.cmd,
					PBClassName:     "PB_PageResponse_SearchTagsPage",
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
	"RPC_Page.SearchUsersPage": func(p rpcParamHandler) {
		if p.rpcHandler.RPC_Page == nil {
			noDevErr(errors.New("rpc service is null for: p.rpcHandler.RPC_Page.SearchUsersPage"))
			return
		}
		load := &PB_PageParam_SearchUsersPage{}
		err := proto.Unmarshal(p.cmd.Data, load)
		if err == nil {
			res, err := p.rpcHandler.RPC_Page.SearchUsersPage(load, p.params)
			if err == nil {
				out := RpcResponseOutput{
					RpcName:         "RPC_Page.SearchUsersPage",
					UserParam:       p.params,
					CommandToServer: p.cmd,
					PBClassName:     "PB_PageResponse_SearchUsersPage",
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

	// rpc: RPC_Search

	"RPC_Search.SearchTags": func(p rpcParamHandler) {
		if p.rpcHandler.RPC_Search == nil {
			noDevErr(errors.New("rpc service is null for: p.rpcHandler.RPC_Search.SearchTags"))
			return
		}
		load := &PB_SearchResponse_AddNewC{}
		err := proto.Unmarshal(p.cmd.Data, load)
		if err == nil {
			res, err := p.rpcHandler.RPC_Search.SearchTags(load, p.params)
			if err == nil {
				out := RpcResponseOutput{
					RpcName:         "RPC_Search.SearchTags",
					UserParam:       p.params,
					CommandToServer: p.cmd,
					PBClassName:     "PB_SearchResponse_AddNewC",
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

	// rpc: RPC_Social

	"RPC_Social.AddComment": func(p rpcParamHandler) {
		if p.rpcHandler.RPC_Social == nil {
			noDevErr(errors.New("rpc service is null for: p.rpcHandler.RPC_Social.AddComment"))
			return
		}
		load := &PB_SocialParam_AddComment{}
		err := proto.Unmarshal(p.cmd.Data, load)
		if err == nil {
			res, err := p.rpcHandler.RPC_Social.AddComment(load, p.params)
			if err == nil {
				out := RpcResponseOutput{
					RpcName:         "RPC_Social.AddComment",
					UserParam:       p.params,
					CommandToServer: p.cmd,
					PBClassName:     "PB_SocialResponse_AddComment",
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
	"RPC_Social.DeleteComment": func(p rpcParamHandler) {
		if p.rpcHandler.RPC_Social == nil {
			noDevErr(errors.New("rpc service is null for: p.rpcHandler.RPC_Social.DeleteComment"))
			return
		}
		load := &PB_SocialParam_DeleteComment{}
		err := proto.Unmarshal(p.cmd.Data, load)
		if err == nil {
			res, err := p.rpcHandler.RPC_Social.DeleteComment(load, p.params)
			if err == nil {
				out := RpcResponseOutput{
					RpcName:         "RPC_Social.DeleteComment",
					UserParam:       p.params,
					CommandToServer: p.cmd,
					PBClassName:     "PB_SocialResponse_DeleteComment",
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
	"RPC_Social.AddPost": func(p rpcParamHandler) {
		if p.rpcHandler.RPC_Social == nil {
			noDevErr(errors.New("rpc service is null for: p.rpcHandler.RPC_Social.AddPost"))
			return
		}
		load := &PB_SocialParam_AddPost{}
		err := proto.Unmarshal(p.cmd.Data, load)
		if err == nil {
			res, err := p.rpcHandler.RPC_Social.AddPost(load, p.params)
			if err == nil {
				out := RpcResponseOutput{
					RpcName:         "RPC_Social.AddPost",
					UserParam:       p.params,
					CommandToServer: p.cmd,
					PBClassName:     "PB_SocialResponse_AddPost",
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
	"RPC_Social.EditPost": func(p rpcParamHandler) {
		if p.rpcHandler.RPC_Social == nil {
			noDevErr(errors.New("rpc service is null for: p.rpcHandler.RPC_Social.EditPost"))
			return
		}
		load := &PB_SocialParam_EditPost{}
		err := proto.Unmarshal(p.cmd.Data, load)
		if err == nil {
			res, err := p.rpcHandler.RPC_Social.EditPost(load, p.params)
			if err == nil {
				out := RpcResponseOutput{
					RpcName:         "RPC_Social.EditPost",
					UserParam:       p.params,
					CommandToServer: p.cmd,
					PBClassName:     "PB_SocialResponse_EditPost",
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
	"RPC_Social.DeletePost": func(p rpcParamHandler) {
		if p.rpcHandler.RPC_Social == nil {
			noDevErr(errors.New("rpc service is null for: p.rpcHandler.RPC_Social.DeletePost"))
			return
		}
		load := &PB_SocialParam_DeletePost{}
		err := proto.Unmarshal(p.cmd.Data, load)
		if err == nil {
			res, err := p.rpcHandler.RPC_Social.DeletePost(load, p.params)
			if err == nil {
				out := RpcResponseOutput{
					RpcName:         "RPC_Social.DeletePost",
					UserParam:       p.params,
					CommandToServer: p.cmd,
					PBClassName:     "PB_SocialResponse_DeletePost",
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
	"RPC_Social.ArchivePost": func(p rpcParamHandler) {
		if p.rpcHandler.RPC_Social == nil {
			noDevErr(errors.New("rpc service is null for: p.rpcHandler.RPC_Social.ArchivePost"))
			return
		}
		load := &PB_SocialParam_ArchivePost{}
		err := proto.Unmarshal(p.cmd.Data, load)
		if err == nil {
			res, err := p.rpcHandler.RPC_Social.ArchivePost(load, p.params)
			if err == nil {
				out := RpcResponseOutput{
					RpcName:         "RPC_Social.ArchivePost",
					UserParam:       p.params,
					CommandToServer: p.cmd,
					PBClassName:     "PB_SocialResponse_ArchivePost",
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
	"RPC_Social.LikePost": func(p rpcParamHandler) {
		if p.rpcHandler.RPC_Social == nil {
			noDevErr(errors.New("rpc service is null for: p.rpcHandler.RPC_Social.LikePost"))
			return
		}
		load := &PB_SocialParam_LikePost{}
		err := proto.Unmarshal(p.cmd.Data, load)
		if err == nil {
			res, err := p.rpcHandler.RPC_Social.LikePost(load, p.params)
			if err == nil {
				out := RpcResponseOutput{
					RpcName:         "RPC_Social.LikePost",
					UserParam:       p.params,
					CommandToServer: p.cmd,
					PBClassName:     "PB_SocialResponse_LikePost",
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
	"RPC_Social.UnLikePost": func(p rpcParamHandler) {
		if p.rpcHandler.RPC_Social == nil {
			noDevErr(errors.New("rpc service is null for: p.rpcHandler.RPC_Social.UnLikePost"))
			return
		}
		load := &PB_SocialParam_UnLikePost{}
		err := proto.Unmarshal(p.cmd.Data, load)
		if err == nil {
			res, err := p.rpcHandler.RPC_Social.UnLikePost(load, p.params)
			if err == nil {
				out := RpcResponseOutput{
					RpcName:         "RPC_Social.UnLikePost",
					UserParam:       p.params,
					CommandToServer: p.cmd,
					PBClassName:     "PB_SocialResponse_UnLikePost",
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
	"RPC_Social.FollowUser": func(p rpcParamHandler) {
		if p.rpcHandler.RPC_Social == nil {
			noDevErr(errors.New("rpc service is null for: p.rpcHandler.RPC_Social.FollowUser"))
			return
		}
		load := &PB_SocialParam_FollowUser{}
		err := proto.Unmarshal(p.cmd.Data, load)
		if err == nil {
			res, err := p.rpcHandler.RPC_Social.FollowUser(load, p.params)
			if err == nil {
				out := RpcResponseOutput{
					RpcName:         "RPC_Social.FollowUser",
					UserParam:       p.params,
					CommandToServer: p.cmd,
					PBClassName:     "PB_SocialResponse_FollowUser",
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
	"RPC_Social.UnFollowUser": func(p rpcParamHandler) {
		if p.rpcHandler.RPC_Social == nil {
			noDevErr(errors.New("rpc service is null for: p.rpcHandler.RPC_Social.UnFollowUser"))
			return
		}
		load := &PB_SocialParam_UnFollowUser{}
		err := proto.Unmarshal(p.cmd.Data, load)
		if err == nil {
			res, err := p.rpcHandler.RPC_Social.UnFollowUser(load, p.params)
			if err == nil {
				out := RpcResponseOutput{
					RpcName:         "RPC_Social.UnFollowUser",
					UserParam:       p.params,
					CommandToServer: p.cmd,
					PBClassName:     "PB_SocialResponse_UnFollowUser",
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

	// rpc: RPC_User

	"RPC_User.BlockUser": func(p rpcParamHandler) {
		if p.rpcHandler.RPC_User == nil {
			noDevErr(errors.New("rpc service is null for: p.rpcHandler.RPC_User.BlockUser"))
			return
		}
		load := &PB_UserParam_BlockUser{}
		err := proto.Unmarshal(p.cmd.Data, load)
		if err == nil {
			res, err := p.rpcHandler.RPC_User.BlockUser(load, p.params)
			if err == nil {
				out := RpcResponseOutput{
					RpcName:         "RPC_User.BlockUser",
					UserParam:       p.params,
					CommandToServer: p.cmd,
					PBClassName:     "PB_UserResponse_BlockUser",
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
	"RPC_User.UnBlockUser": func(p rpcParamHandler) {
		if p.rpcHandler.RPC_User == nil {
			noDevErr(errors.New("rpc service is null for: p.rpcHandler.RPC_User.UnBlockUser"))
			return
		}
		load := &PB_UserParam_UnBlockUser{}
		err := proto.Unmarshal(p.cmd.Data, load)
		if err == nil {
			res, err := p.rpcHandler.RPC_User.UnBlockUser(load, p.params)
			if err == nil {
				out := RpcResponseOutput{
					RpcName:         "RPC_User.UnBlockUser",
					UserParam:       p.params,
					CommandToServer: p.cmd,
					PBClassName:     "PB_UserResponse_UnBlockUser",
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
	"RPC_User.GetBlockedList": func(p rpcParamHandler) {
		if p.rpcHandler.RPC_User == nil {
			noDevErr(errors.New("rpc service is null for: p.rpcHandler.RPC_User.GetBlockedList"))
			return
		}
		load := &PB_UserParam_BlockedList{}
		err := proto.Unmarshal(p.cmd.Data, load)
		if err == nil {
			res, err := p.rpcHandler.RPC_User.GetBlockedList(load, p.params)
			if err == nil {
				out := RpcResponseOutput{
					RpcName:         "RPC_User.GetBlockedList",
					UserParam:       p.params,
					CommandToServer: p.cmd,
					PBClassName:     "PB_UserResponse_BlockedList",
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
	"RPC_User.UpdateAbout": func(p rpcParamHandler) {
		if p.rpcHandler.RPC_User == nil {
			noDevErr(errors.New("rpc service is null for: p.rpcHandler.RPC_User.UpdateAbout"))
			return
		}
		load := &PB_UserParam_UpdateAbout{}
		err := proto.Unmarshal(p.cmd.Data, load)
		if err == nil {
			res, err := p.rpcHandler.RPC_User.UpdateAbout(load, p.params)
			if err == nil {
				out := RpcResponseOutput{
					RpcName:         "RPC_User.UpdateAbout",
					UserParam:       p.params,
					CommandToServer: p.cmd,
					PBClassName:     "PB_UserResponse_UpdateAbout",
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
	"RPC_User.UpdateUserName": func(p rpcParamHandler) {
		if p.rpcHandler.RPC_User == nil {
			noDevErr(errors.New("rpc service is null for: p.rpcHandler.RPC_User.UpdateUserName"))
			return
		}
		load := &PB_UserParam_UpdateUserName{}
		err := proto.Unmarshal(p.cmd.Data, load)
		if err == nil {
			res, err := p.rpcHandler.RPC_User.UpdateUserName(load, p.params)
			if err == nil {
				out := RpcResponseOutput{
					RpcName:         "RPC_User.UpdateUserName",
					UserParam:       p.params,
					CommandToServer: p.cmd,
					PBClassName:     "PB_UserResponse_UpdateUserName",
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
	"RPC_User.ChangePrivacy": func(p rpcParamHandler) {
		if p.rpcHandler.RPC_User == nil {
			noDevErr(errors.New("rpc service is null for: p.rpcHandler.RPC_User.ChangePrivacy"))
			return
		}
		load := &PB_UserParam_ChangePrivacy{}
		err := proto.Unmarshal(p.cmd.Data, load)
		if err == nil {
			res, err := p.rpcHandler.RPC_User.ChangePrivacy(load, p.params)
			if err == nil {
				out := RpcResponseOutput{
					RpcName:         "RPC_User.ChangePrivacy",
					UserParam:       p.params,
					CommandToServer: p.cmd,
					PBClassName:     "PB_UserResponseOffline_ChangePrivacy",
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
	"RPC_User.ChangeAvatar": func(p rpcParamHandler) {
		if p.rpcHandler.RPC_User == nil {
			noDevErr(errors.New("rpc service is null for: p.rpcHandler.RPC_User.ChangeAvatar"))
			return
		}
		load := &PB_UserParam_ChangeAvatar{}
		err := proto.Unmarshal(p.cmd.Data, load)
		if err == nil {
			res, err := p.rpcHandler.RPC_User.ChangeAvatar(load, p.params)
			if err == nil {
				out := RpcResponseOutput{
					RpcName:         "RPC_User.ChangeAvatar",
					UserParam:       p.params,
					CommandToServer: p.cmd,
					PBClassName:     "PB_UserResponse_ChangeAvatar",
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
	"RPC_User.CheckUserName": func(p rpcParamHandler) {
		if p.rpcHandler.RPC_User == nil {
			noDevErr(errors.New("rpc service is null for: p.rpcHandler.RPC_User.CheckUserName"))
			return
		}
		load := &PB_UserParam_CheckUserName{}
		err := proto.Unmarshal(p.cmd.Data, load)
		if err == nil {
			res, err := p.rpcHandler.RPC_User.CheckUserName(load, p.params)
			if err == nil {
				out := RpcResponseOutput{
					RpcName:         "RPC_User.CheckUserName",
					UserParam:       p.params,
					CommandToServer: p.cmd,
					PBClassName:     "PB_UserResponse_CheckUserName",
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
