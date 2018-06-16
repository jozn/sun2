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
	FailRequest(error PB_Error)
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
	RPC_Auth    RPC_Auth
	RPC_Chat    RPC_Chat
	RPC_General RPC_General
	RPC_Page    RPC_Page
	RPC_Social  RPC_Social
	RPC_User    RPC_User
}

/////////////// Interfaces ////////////////

type RPC_Auth interface {
	SendConfirmCode(param *RPC_Auth_Types_SendConfirmCode_Param, userParam RPC_UserParam) (res RPC_Auth_Types_SendConfirmCode_Response, err error)
	ConfirmCode(param *RPC_Auth_Types_ConfirmCode_Param, userParam RPC_UserParam) (res RPC_Auth_Types_ConfirmCode_Response, err error)
	SingUp(param *RPC_Auth_Types_SingUp_Param, userParam RPC_UserParam) (res RPC_Auth_Types_SingUp_Response, err error)
	SingIn(param *RPC_Auth_Types_SingIn_Param, userParam RPC_UserParam) (res RPC_Auth_Types_SingIn_Response, err error)
	LogOut(param *RPC_Auth_Types_LogOut_Param, userParam RPC_UserParam) (res RPC_Auth_Types_LogOut_Response, err error)
}

type RPC_Chat interface {
	AddNewMessage(param *PB_RPC_Chat_Types_AddNewMessage_Param, userParam RPC_UserParam) (res PB_RPC_Chat_Types_AddNewMessage_Response, err error)
	SetRoomActionDoing(param *PB_RPC_Chat_Types_SetRoomActionDoing_Param, userParam RPC_UserParam) (res PB_RPC_Chat_Types_SetRoomActionDoing_Response, err error)
	GetChatList(param *PB_RPC_Chat_Types_GetChatList_Param, userParam RPC_UserParam) (res PB_RPC_Chat_Types_GetChatList_Response, err error)
	GetChatHistory(param *PB_RPC_Chat_Types_GetChatHistory_Param, userParam RPC_UserParam) (res PB_RPC_Chat_Types_GetChatHistory_Response, err error)
	PushRoomsChange(param *PB_RPC_Chat_Types_PushRoomsChange_Param, userParam RPC_UserParam) (res PB_RPC_Chat_Types_PushRoomsChange_Response, err error)
	GetRoomsChange(param *PB_RPC_Chat_Types_GetRoomsChange_Param, userParam RPC_UserParam) (res PB_RPC_Chat_Types_GetRoomsChange_Response, err error)
}

type RPC_General interface {
	Echo(param *RPC_General_Types_Echo_Param, userParam RPC_UserParam) (res RPC_General_Types_Echo_Response, err error)
	CheckUserName(param *RPC_General_Types_CheckUserName_Param, userParam RPC_UserParam) (res RPC_General_Types_CheckUserName_Response, err error)
}

type RPC_Page interface {
	GetCommentsPage(param *RPC_Page_Types_GetCommentsPage_Param, userParam RPC_UserParam) (res RPC_Page_Types_GetCommentsPage_Response, err error)
	GetHomePage(param *RPC_Page_Types_GetHomePage_Param, userParam RPC_UserParam) (res RPC_Page_Types_GetHomePage_Response, err error)
	GetProfileAbout(param *RPC_Page_Types_GetProfileAbout_Param, userParam RPC_UserParam) (res RPC_Page_Types_GetProfileAbout_Response, err error)
	GetProfileAllShared(param *RPC_Page_Types_GetProfileAllShared_Param, userParam RPC_UserParam) (res RPC_Page_Types_GetProfileAllShared_Response, err error)
	GetProfileByCategoryPage(param *RPC_Page_Types_GetProfileByCategoryPage_Param, userParam RPC_UserParam) (res RPC_Page_Types_GetProfileByCategoryPage_Response, err error)
	GetLikesPage(param *RPC_Page_Types_GetLikesPage_Param, userParam RPC_UserParam) (res RPC_Page_Types_GetLikesPage_Response, err error)
	GetFollowersPage(param *RPC_Page_Types_GetFollowersPage_Param, userParam RPC_UserParam) (res RPC_Page_Types_GetFollowersPage_Response, err error)
	GetFollowingsPage(param *RPC_Page_Types_GetFollowingsPage_Param, userParam RPC_UserParam) (res RPC_Page_Types_GetFollowingsPage_Response, err error)
	GetNotifiesPage(param *RPC_Page_Types_GetNotifiesPage_Param, userParam RPC_UserParam) (res RPC_Page_Types_GetNotifiesPage_Response, err error)
	GetUserActionsPage(param *RPC_Page_Types_GetUserActionsPage_Param, userParam RPC_UserParam) (res RPC_Page_Types_GetUserActionsPage_Response, err error)
	GetPromotedPostsPage(param *RPC_Page_Types_GetPromotedPostsPage_Param, userParam RPC_UserParam) (res RPC_Page_Types_GetPromotedPostsPage_Response, err error)
	GetSuggestedUsersPage(param *RPC_Page_Types_GetSuggestedUsersPage_Param, userParam RPC_UserParam) (res RPC_Page_Types_GetSuggestedUsersPage_Response, err error)
	GetSuggestedTagsPage(param *RPC_Page_Types_GetSuggestedTagsPage_Param, userParam RPC_UserParam) (res RPC_Page_Types_GetSuggestedTagsPage_Response, err error)
	GetLastPostsPage(param *RPC_Page_Types_GetLastPostsPage_Param, userParam RPC_UserParam) (res RPC_Page_Types_GetLastPostsPage_Response, err error)
	GetLastTagPage(param *RPC_Page_Types_GetLastTagPage_Param, userParam RPC_UserParam) (res RPC_Page_Types_GetLastTagPage_Response, err error)
	SearchTagsPage(param *RPC_Page_Types_SearchTagsPage_Param, userParam RPC_UserParam) (res RPC_Page_Types_SearchTagsPage_Response, err error)
	SearchUsersPage(param *RPC_Page_Types_SearchUsersPage_Param, userParam RPC_UserParam) (res RPC_Page_Types_SearchUsersPage_Response, err error)
}

type RPC_Social interface {
	AddComment(param *RPC_Social_Types_AddComment_Param, userParam RPC_UserParam) (res RPC_Social_Types_AddComment_Response, err error)
	DeleteComment(param *RPC_Social_Types_DeleteComment_Param, userParam RPC_UserParam) (res RPC_Social_Types_DeleteComment_Response, err error)
	EditComment(param *RPC_Social_Types_EditComment_Param, userParam RPC_UserParam) (res RPC_Social_Types_EditComment_Response, err error)
	AddPost(param *RPC_Social_Types_AddPost_Param, userParam RPC_UserParam) (res RPC_Social_Types_AddPost_Response, err error)
	EditPost(param *RPC_Social_Types_EditPost_Param, userParam RPC_UserParam) (res RPC_Social_Types_EditPost_Response, err error)
	DeletePost(param *RPC_Social_Types_DeletePost_Param, userParam RPC_UserParam) (res RPC_Social_Types_DeletePost_Response, err error)
	ArchivePost(param *RPC_Social_Types_ArchivePost_Param, userParam RPC_UserParam) (res RPC_Social_Types_ArchivePost_Response, err error)
	LikePost(param *RPC_Social_Types_LikePost_Param, userParam RPC_UserParam) (res RPC_Social_Types_LikePost_Response, err error)
	UnLikePost(param *RPC_Social_Types_UnLikePost_Param, userParam RPC_UserParam) (res RPC_Social_Types_UnLikePost_Response, err error)
	FollowUser(param *RPC_Social_Types_FollowUser_Param, userParam RPC_UserParam) (res RPC_Social_Types_FollowUser_Response, err error)
	UnFollowUser(param *RPC_Social_Types_UnFollowUser_Param, userParam RPC_UserParam) (res RPC_Social_Types_UnFollowUser_Response, err error)
	BlockUser(param *RPC_Social_Types_BlockUser_Param, userParam RPC_UserParam) (res RPC_Social_Types_BlockUser_Response, err error)
	UnBlockUser(param *RPC_Social_Types_UnBlockUser_Param, userParam RPC_UserParam) (res RPC_Social_Types_UnBlockUser_Response, err error)
	AddSeenPosts(param *RPC_Social_Types_AddSeenPosts_Param, userParam RPC_UserParam) (res RPC_Social_Types_AddSeenPosts_Response, err error)
}

type RPC_User interface {
	UpdateAbout(param *RPC_User_Types_UpdateAbout_Param, userParam RPC_UserParam) (res RPC_User_Types_UpdateAbout_Response, err error)
	UpdateUserName(param *RPC_User_Types_UpdateUserName_Param, userParam RPC_UserParam) (res RPC_User_Types_UpdateUserName_Response, err error)
	ChangePrivacy(param *RPC_User_Types_ChangePrivacy_Param, userParam RPC_UserParam) (res RPC_User_Types_ChangePrivacy_Response, err error)
	ChangeAvatar(param *RPC_User_Types_ChangeAvatar_Param, userParam RPC_UserParam) (res RPC_User_Types_ChangeAvatar_Response, err error)
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

	// rpc: RPC_Auth

	"RPC_Auth.SendConfirmCode": func(p rpcParamHandler) {
		if p.rpcHandler.RPC_Auth == nil {
			noDevErr(errors.New("rpc service is null for: p.rpcHandler.RPC_Auth.SendConfirmCode"))
			return
		}
		load := &RPC_Auth_Types_SendConfirmCode_Param{}
		err := proto.Unmarshal(p.cmd.Data, load)
		if err == nil {
			res, err := p.rpcHandler.RPC_Auth.SendConfirmCode(load, p.params)
			if err == nil {
				out := RpcResponseOutput{
					RpcName:         "RPC_Auth.SendConfirmCode",
					UserParam:       p.params,
					CommandToServer: p.cmd,
					PBClassName:     "RPC_Auth_Types.SendConfirmCode.Response",
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
	"RPC_Auth.ConfirmCode": func(p rpcParamHandler) {
		if p.rpcHandler.RPC_Auth == nil {
			noDevErr(errors.New("rpc service is null for: p.rpcHandler.RPC_Auth.ConfirmCode"))
			return
		}
		load := &RPC_Auth_Types_ConfirmCode_Param{}
		err := proto.Unmarshal(p.cmd.Data, load)
		if err == nil {
			res, err := p.rpcHandler.RPC_Auth.ConfirmCode(load, p.params)
			if err == nil {
				out := RpcResponseOutput{
					RpcName:         "RPC_Auth.ConfirmCode",
					UserParam:       p.params,
					CommandToServer: p.cmd,
					PBClassName:     "RPC_Auth_Types.ConfirmCode.Response",
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
		load := &RPC_Auth_Types_SingUp_Param{}
		err := proto.Unmarshal(p.cmd.Data, load)
		if err == nil {
			res, err := p.rpcHandler.RPC_Auth.SingUp(load, p.params)
			if err == nil {
				out := RpcResponseOutput{
					RpcName:         "RPC_Auth.SingUp",
					UserParam:       p.params,
					CommandToServer: p.cmd,
					PBClassName:     "RPC_Auth_Types.SingUp.Response",
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
		load := &RPC_Auth_Types_SingIn_Param{}
		err := proto.Unmarshal(p.cmd.Data, load)
		if err == nil {
			res, err := p.rpcHandler.RPC_Auth.SingIn(load, p.params)
			if err == nil {
				out := RpcResponseOutput{
					RpcName:         "RPC_Auth.SingIn",
					UserParam:       p.params,
					CommandToServer: p.cmd,
					PBClassName:     "RPC_Auth_Types.SingIn.Response",
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
		load := &RPC_Auth_Types_LogOut_Param{}
		err := proto.Unmarshal(p.cmd.Data, load)
		if err == nil {
			res, err := p.rpcHandler.RPC_Auth.LogOut(load, p.params)
			if err == nil {
				out := RpcResponseOutput{
					RpcName:         "RPC_Auth.LogOut",
					UserParam:       p.params,
					CommandToServer: p.cmd,
					PBClassName:     "RPC_Auth_Types.LogOut.Response",
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

	// rpc: RPC_General

	"RPC_General.Echo": func(p rpcParamHandler) {
		if p.rpcHandler.RPC_General == nil {
			noDevErr(errors.New("rpc service is null for: p.rpcHandler.RPC_General.Echo"))
			return
		}
		load := &RPC_General_Types_Echo_Param{}
		err := proto.Unmarshal(p.cmd.Data, load)
		if err == nil {
			res, err := p.rpcHandler.RPC_General.Echo(load, p.params)
			if err == nil {
				out := RpcResponseOutput{
					RpcName:         "RPC_General.Echo",
					UserParam:       p.params,
					CommandToServer: p.cmd,
					PBClassName:     "RPC_General_Types.Echo.Response",
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
	"RPC_General.CheckUserName": func(p rpcParamHandler) {
		if p.rpcHandler.RPC_General == nil {
			noDevErr(errors.New("rpc service is null for: p.rpcHandler.RPC_General.CheckUserName"))
			return
		}
		load := &RPC_General_Types_CheckUserName_Param{}
		err := proto.Unmarshal(p.cmd.Data, load)
		if err == nil {
			res, err := p.rpcHandler.RPC_General.CheckUserName(load, p.params)
			if err == nil {
				out := RpcResponseOutput{
					RpcName:         "RPC_General.CheckUserName",
					UserParam:       p.params,
					CommandToServer: p.cmd,
					PBClassName:     "RPC_General_Types.CheckUserName.Response",
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
		load := &RPC_Page_Types_GetCommentsPage_Param{}
		err := proto.Unmarshal(p.cmd.Data, load)
		if err == nil {
			res, err := p.rpcHandler.RPC_Page.GetCommentsPage(load, p.params)
			if err == nil {
				out := RpcResponseOutput{
					RpcName:         "RPC_Page.GetCommentsPage",
					UserParam:       p.params,
					CommandToServer: p.cmd,
					PBClassName:     "RPC_Page_Types.GetCommentsPage.Response",
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
		load := &RPC_Page_Types_GetHomePage_Param{}
		err := proto.Unmarshal(p.cmd.Data, load)
		if err == nil {
			res, err := p.rpcHandler.RPC_Page.GetHomePage(load, p.params)
			if err == nil {
				out := RpcResponseOutput{
					RpcName:         "RPC_Page.GetHomePage",
					UserParam:       p.params,
					CommandToServer: p.cmd,
					PBClassName:     "RPC_Page_Types.GetHomePage.Response",
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
	"RPC_Page.GetProfileAbout": func(p rpcParamHandler) {
		if p.rpcHandler.RPC_Page == nil {
			noDevErr(errors.New("rpc service is null for: p.rpcHandler.RPC_Page.GetProfileAbout"))
			return
		}
		load := &RPC_Page_Types_GetProfileAbout_Param{}
		err := proto.Unmarshal(p.cmd.Data, load)
		if err == nil {
			res, err := p.rpcHandler.RPC_Page.GetProfileAbout(load, p.params)
			if err == nil {
				out := RpcResponseOutput{
					RpcName:         "RPC_Page.GetProfileAbout",
					UserParam:       p.params,
					CommandToServer: p.cmd,
					PBClassName:     "RPC_Page_Types.GetProfileAbout.Response",
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
	"RPC_Page.GetProfileAllShared": func(p rpcParamHandler) {
		if p.rpcHandler.RPC_Page == nil {
			noDevErr(errors.New("rpc service is null for: p.rpcHandler.RPC_Page.GetProfileAllShared"))
			return
		}
		load := &RPC_Page_Types_GetProfileAllShared_Param{}
		err := proto.Unmarshal(p.cmd.Data, load)
		if err == nil {
			res, err := p.rpcHandler.RPC_Page.GetProfileAllShared(load, p.params)
			if err == nil {
				out := RpcResponseOutput{
					RpcName:         "RPC_Page.GetProfileAllShared",
					UserParam:       p.params,
					CommandToServer: p.cmd,
					PBClassName:     "RPC_Page_Types.GetProfileAllShared.Response",
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
	"RPC_Page.GetProfileByCategoryPage": func(p rpcParamHandler) {
		if p.rpcHandler.RPC_Page == nil {
			noDevErr(errors.New("rpc service is null for: p.rpcHandler.RPC_Page.GetProfileByCategoryPage"))
			return
		}
		load := &RPC_Page_Types_GetProfileByCategoryPage_Param{}
		err := proto.Unmarshal(p.cmd.Data, load)
		if err == nil {
			res, err := p.rpcHandler.RPC_Page.GetProfileByCategoryPage(load, p.params)
			if err == nil {
				out := RpcResponseOutput{
					RpcName:         "RPC_Page.GetProfileByCategoryPage",
					UserParam:       p.params,
					CommandToServer: p.cmd,
					PBClassName:     "RPC_Page_Types.GetProfileByCategoryPage.Response",
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
		load := &RPC_Page_Types_GetLikesPage_Param{}
		err := proto.Unmarshal(p.cmd.Data, load)
		if err == nil {
			res, err := p.rpcHandler.RPC_Page.GetLikesPage(load, p.params)
			if err == nil {
				out := RpcResponseOutput{
					RpcName:         "RPC_Page.GetLikesPage",
					UserParam:       p.params,
					CommandToServer: p.cmd,
					PBClassName:     "RPC_Page_Types.GetLikesPage.Response",
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
		load := &RPC_Page_Types_GetFollowersPage_Param{}
		err := proto.Unmarshal(p.cmd.Data, load)
		if err == nil {
			res, err := p.rpcHandler.RPC_Page.GetFollowersPage(load, p.params)
			if err == nil {
				out := RpcResponseOutput{
					RpcName:         "RPC_Page.GetFollowersPage",
					UserParam:       p.params,
					CommandToServer: p.cmd,
					PBClassName:     "RPC_Page_Types.GetFollowersPage.Response",
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
		load := &RPC_Page_Types_GetFollowingsPage_Param{}
		err := proto.Unmarshal(p.cmd.Data, load)
		if err == nil {
			res, err := p.rpcHandler.RPC_Page.GetFollowingsPage(load, p.params)
			if err == nil {
				out := RpcResponseOutput{
					RpcName:         "RPC_Page.GetFollowingsPage",
					UserParam:       p.params,
					CommandToServer: p.cmd,
					PBClassName:     "RPC_Page_Types.GetFollowingsPage.Response",
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
		load := &RPC_Page_Types_GetNotifiesPage_Param{}
		err := proto.Unmarshal(p.cmd.Data, load)
		if err == nil {
			res, err := p.rpcHandler.RPC_Page.GetNotifiesPage(load, p.params)
			if err == nil {
				out := RpcResponseOutput{
					RpcName:         "RPC_Page.GetNotifiesPage",
					UserParam:       p.params,
					CommandToServer: p.cmd,
					PBClassName:     "RPC_Page_Types.GetNotifiesPage.Response",
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
		load := &RPC_Page_Types_GetUserActionsPage_Param{}
		err := proto.Unmarshal(p.cmd.Data, load)
		if err == nil {
			res, err := p.rpcHandler.RPC_Page.GetUserActionsPage(load, p.params)
			if err == nil {
				out := RpcResponseOutput{
					RpcName:         "RPC_Page.GetUserActionsPage",
					UserParam:       p.params,
					CommandToServer: p.cmd,
					PBClassName:     "RPC_Page_Types.GetUserActionsPage.Response",
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
	"RPC_Page.GetPromotedPostsPage": func(p rpcParamHandler) {
		if p.rpcHandler.RPC_Page == nil {
			noDevErr(errors.New("rpc service is null for: p.rpcHandler.RPC_Page.GetPromotedPostsPage"))
			return
		}
		load := &RPC_Page_Types_GetPromotedPostsPage_Param{}
		err := proto.Unmarshal(p.cmd.Data, load)
		if err == nil {
			res, err := p.rpcHandler.RPC_Page.GetPromotedPostsPage(load, p.params)
			if err == nil {
				out := RpcResponseOutput{
					RpcName:         "RPC_Page.GetPromotedPostsPage",
					UserParam:       p.params,
					CommandToServer: p.cmd,
					PBClassName:     "RPC_Page_Types.GetPromotedPostsPage.Response",
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
		load := &RPC_Page_Types_GetSuggestedUsersPage_Param{}
		err := proto.Unmarshal(p.cmd.Data, load)
		if err == nil {
			res, err := p.rpcHandler.RPC_Page.GetSuggestedUsersPage(load, p.params)
			if err == nil {
				out := RpcResponseOutput{
					RpcName:         "RPC_Page.GetSuggestedUsersPage",
					UserParam:       p.params,
					CommandToServer: p.cmd,
					PBClassName:     "RPC_Page_Types.GetSuggestedUsersPage.Response",
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
		load := &RPC_Page_Types_GetSuggestedTagsPage_Param{}
		err := proto.Unmarshal(p.cmd.Data, load)
		if err == nil {
			res, err := p.rpcHandler.RPC_Page.GetSuggestedTagsPage(load, p.params)
			if err == nil {
				out := RpcResponseOutput{
					RpcName:         "RPC_Page.GetSuggestedTagsPage",
					UserParam:       p.params,
					CommandToServer: p.cmd,
					PBClassName:     "RPC_Page_Types.GetSuggestedTagsPage.Response",
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
		load := &RPC_Page_Types_GetLastPostsPage_Param{}
		err := proto.Unmarshal(p.cmd.Data, load)
		if err == nil {
			res, err := p.rpcHandler.RPC_Page.GetLastPostsPage(load, p.params)
			if err == nil {
				out := RpcResponseOutput{
					RpcName:         "RPC_Page.GetLastPostsPage",
					UserParam:       p.params,
					CommandToServer: p.cmd,
					PBClassName:     "RPC_Page_Types.GetLastPostsPage.Response",
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
	"RPC_Page.GetLastTagPage": func(p rpcParamHandler) {
		if p.rpcHandler.RPC_Page == nil {
			noDevErr(errors.New("rpc service is null for: p.rpcHandler.RPC_Page.GetLastTagPage"))
			return
		}
		load := &RPC_Page_Types_GetLastTagPage_Param{}
		err := proto.Unmarshal(p.cmd.Data, load)
		if err == nil {
			res, err := p.rpcHandler.RPC_Page.GetLastTagPage(load, p.params)
			if err == nil {
				out := RpcResponseOutput{
					RpcName:         "RPC_Page.GetLastTagPage",
					UserParam:       p.params,
					CommandToServer: p.cmd,
					PBClassName:     "RPC_Page_Types.GetLastTagPage.Response",
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
		load := &RPC_Page_Types_SearchTagsPage_Param{}
		err := proto.Unmarshal(p.cmd.Data, load)
		if err == nil {
			res, err := p.rpcHandler.RPC_Page.SearchTagsPage(load, p.params)
			if err == nil {
				out := RpcResponseOutput{
					RpcName:         "RPC_Page.SearchTagsPage",
					UserParam:       p.params,
					CommandToServer: p.cmd,
					PBClassName:     "RPC_Page_Types.SearchTagsPage.Response",
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
		load := &RPC_Page_Types_SearchUsersPage_Param{}
		err := proto.Unmarshal(p.cmd.Data, load)
		if err == nil {
			res, err := p.rpcHandler.RPC_Page.SearchUsersPage(load, p.params)
			if err == nil {
				out := RpcResponseOutput{
					RpcName:         "RPC_Page.SearchUsersPage",
					UserParam:       p.params,
					CommandToServer: p.cmd,
					PBClassName:     "RPC_Page_Types.SearchUsersPage.Response",
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
		load := &RPC_Social_Types_AddComment_Param{}
		err := proto.Unmarshal(p.cmd.Data, load)
		if err == nil {
			res, err := p.rpcHandler.RPC_Social.AddComment(load, p.params)
			if err == nil {
				out := RpcResponseOutput{
					RpcName:         "RPC_Social.AddComment",
					UserParam:       p.params,
					CommandToServer: p.cmd,
					PBClassName:     "RPC_Social_Types.AddComment.Response",
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
		load := &RPC_Social_Types_DeleteComment_Param{}
		err := proto.Unmarshal(p.cmd.Data, load)
		if err == nil {
			res, err := p.rpcHandler.RPC_Social.DeleteComment(load, p.params)
			if err == nil {
				out := RpcResponseOutput{
					RpcName:         "RPC_Social.DeleteComment",
					UserParam:       p.params,
					CommandToServer: p.cmd,
					PBClassName:     "RPC_Social_Types.DeleteComment.Response",
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
	"RPC_Social.EditComment": func(p rpcParamHandler) {
		if p.rpcHandler.RPC_Social == nil {
			noDevErr(errors.New("rpc service is null for: p.rpcHandler.RPC_Social.EditComment"))
			return
		}
		load := &RPC_Social_Types_EditComment_Param{}
		err := proto.Unmarshal(p.cmd.Data, load)
		if err == nil {
			res, err := p.rpcHandler.RPC_Social.EditComment(load, p.params)
			if err == nil {
				out := RpcResponseOutput{
					RpcName:         "RPC_Social.EditComment",
					UserParam:       p.params,
					CommandToServer: p.cmd,
					PBClassName:     "RPC_Social_Types.EditComment.Response",
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
		load := &RPC_Social_Types_AddPost_Param{}
		err := proto.Unmarshal(p.cmd.Data, load)
		if err == nil {
			res, err := p.rpcHandler.RPC_Social.AddPost(load, p.params)
			if err == nil {
				out := RpcResponseOutput{
					RpcName:         "RPC_Social.AddPost",
					UserParam:       p.params,
					CommandToServer: p.cmd,
					PBClassName:     "RPC_Social_Types.AddPost.Response",
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
		load := &RPC_Social_Types_EditPost_Param{}
		err := proto.Unmarshal(p.cmd.Data, load)
		if err == nil {
			res, err := p.rpcHandler.RPC_Social.EditPost(load, p.params)
			if err == nil {
				out := RpcResponseOutput{
					RpcName:         "RPC_Social.EditPost",
					UserParam:       p.params,
					CommandToServer: p.cmd,
					PBClassName:     "RPC_Social_Types.EditPost.Response",
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
		load := &RPC_Social_Types_DeletePost_Param{}
		err := proto.Unmarshal(p.cmd.Data, load)
		if err == nil {
			res, err := p.rpcHandler.RPC_Social.DeletePost(load, p.params)
			if err == nil {
				out := RpcResponseOutput{
					RpcName:         "RPC_Social.DeletePost",
					UserParam:       p.params,
					CommandToServer: p.cmd,
					PBClassName:     "RPC_Social_Types.DeletePost.Response",
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
		load := &RPC_Social_Types_ArchivePost_Param{}
		err := proto.Unmarshal(p.cmd.Data, load)
		if err == nil {
			res, err := p.rpcHandler.RPC_Social.ArchivePost(load, p.params)
			if err == nil {
				out := RpcResponseOutput{
					RpcName:         "RPC_Social.ArchivePost",
					UserParam:       p.params,
					CommandToServer: p.cmd,
					PBClassName:     "RPC_Social_Types.ArchivePost.Response",
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
		load := &RPC_Social_Types_LikePost_Param{}
		err := proto.Unmarshal(p.cmd.Data, load)
		if err == nil {
			res, err := p.rpcHandler.RPC_Social.LikePost(load, p.params)
			if err == nil {
				out := RpcResponseOutput{
					RpcName:         "RPC_Social.LikePost",
					UserParam:       p.params,
					CommandToServer: p.cmd,
					PBClassName:     "RPC_Social_Types.LikePost.Response",
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
		load := &RPC_Social_Types_UnLikePost_Param{}
		err := proto.Unmarshal(p.cmd.Data, load)
		if err == nil {
			res, err := p.rpcHandler.RPC_Social.UnLikePost(load, p.params)
			if err == nil {
				out := RpcResponseOutput{
					RpcName:         "RPC_Social.UnLikePost",
					UserParam:       p.params,
					CommandToServer: p.cmd,
					PBClassName:     "RPC_Social_Types.UnLikePost.Response",
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
		load := &RPC_Social_Types_FollowUser_Param{}
		err := proto.Unmarshal(p.cmd.Data, load)
		if err == nil {
			res, err := p.rpcHandler.RPC_Social.FollowUser(load, p.params)
			if err == nil {
				out := RpcResponseOutput{
					RpcName:         "RPC_Social.FollowUser",
					UserParam:       p.params,
					CommandToServer: p.cmd,
					PBClassName:     "RPC_Social_Types.FollowUser.Response",
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
		load := &RPC_Social_Types_UnFollowUser_Param{}
		err := proto.Unmarshal(p.cmd.Data, load)
		if err == nil {
			res, err := p.rpcHandler.RPC_Social.UnFollowUser(load, p.params)
			if err == nil {
				out := RpcResponseOutput{
					RpcName:         "RPC_Social.UnFollowUser",
					UserParam:       p.params,
					CommandToServer: p.cmd,
					PBClassName:     "RPC_Social_Types.UnFollowUser.Response",
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
	"RPC_Social.BlockUser": func(p rpcParamHandler) {
		if p.rpcHandler.RPC_Social == nil {
			noDevErr(errors.New("rpc service is null for: p.rpcHandler.RPC_Social.BlockUser"))
			return
		}
		load := &RPC_Social_Types_BlockUser_Param{}
		err := proto.Unmarshal(p.cmd.Data, load)
		if err == nil {
			res, err := p.rpcHandler.RPC_Social.BlockUser(load, p.params)
			if err == nil {
				out := RpcResponseOutput{
					RpcName:         "RPC_Social.BlockUser",
					UserParam:       p.params,
					CommandToServer: p.cmd,
					PBClassName:     "RPC_Social_Types.BlockUser.Response",
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
	"RPC_Social.UnBlockUser": func(p rpcParamHandler) {
		if p.rpcHandler.RPC_Social == nil {
			noDevErr(errors.New("rpc service is null for: p.rpcHandler.RPC_Social.UnBlockUser"))
			return
		}
		load := &RPC_Social_Types_UnBlockUser_Param{}
		err := proto.Unmarshal(p.cmd.Data, load)
		if err == nil {
			res, err := p.rpcHandler.RPC_Social.UnBlockUser(load, p.params)
			if err == nil {
				out := RpcResponseOutput{
					RpcName:         "RPC_Social.UnBlockUser",
					UserParam:       p.params,
					CommandToServer: p.cmd,
					PBClassName:     "RPC_Social_Types.UnBlockUser.Response",
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
	"RPC_Social.AddSeenPosts": func(p rpcParamHandler) {
		if p.rpcHandler.RPC_Social == nil {
			noDevErr(errors.New("rpc service is null for: p.rpcHandler.RPC_Social.AddSeenPosts"))
			return
		}
		load := &RPC_Social_Types_AddSeenPosts_Param{}
		err := proto.Unmarshal(p.cmd.Data, load)
		if err == nil {
			res, err := p.rpcHandler.RPC_Social.AddSeenPosts(load, p.params)
			if err == nil {
				out := RpcResponseOutput{
					RpcName:         "RPC_Social.AddSeenPosts",
					UserParam:       p.params,
					CommandToServer: p.cmd,
					PBClassName:     "RPC_Social_Types.AddSeenPosts.Response",
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

	"RPC_User.UpdateAbout": func(p rpcParamHandler) {
		if p.rpcHandler.RPC_User == nil {
			noDevErr(errors.New("rpc service is null for: p.rpcHandler.RPC_User.UpdateAbout"))
			return
		}
		load := &RPC_User_Types_UpdateAbout_Param{}
		err := proto.Unmarshal(p.cmd.Data, load)
		if err == nil {
			res, err := p.rpcHandler.RPC_User.UpdateAbout(load, p.params)
			if err == nil {
				out := RpcResponseOutput{
					RpcName:         "RPC_User.UpdateAbout",
					UserParam:       p.params,
					CommandToServer: p.cmd,
					PBClassName:     "RPC_User_Types.UpdateAbout.Response",
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
		load := &RPC_User_Types_UpdateUserName_Param{}
		err := proto.Unmarshal(p.cmd.Data, load)
		if err == nil {
			res, err := p.rpcHandler.RPC_User.UpdateUserName(load, p.params)
			if err == nil {
				out := RpcResponseOutput{
					RpcName:         "RPC_User.UpdateUserName",
					UserParam:       p.params,
					CommandToServer: p.cmd,
					PBClassName:     "RPC_User_Types.UpdateUserName.Response",
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
		load := &RPC_User_Types_ChangePrivacy_Param{}
		err := proto.Unmarshal(p.cmd.Data, load)
		if err == nil {
			res, err := p.rpcHandler.RPC_User.ChangePrivacy(load, p.params)
			if err == nil {
				out := RpcResponseOutput{
					RpcName:         "RPC_User.ChangePrivacy",
					UserParam:       p.params,
					CommandToServer: p.cmd,
					PBClassName:     "RPC_User_Types.ChangePrivacy.Response",
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
		load := &RPC_User_Types_ChangeAvatar_Param{}
		err := proto.Unmarshal(p.cmd.Data, load)
		if err == nil {
			res, err := p.rpcHandler.RPC_User.ChangeAvatar(load, p.params)
			if err == nil {
				out := RpcResponseOutput{
					RpcName:         "RPC_User.ChangeAvatar",
					UserParam:       p.params,
					CommandToServer: p.cmd,
					PBClassName:     "RPC_User_Types.ChangeAvatar.Response",
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


 RPC_Auth.SendConfirmCode
 RPC_Auth.ConfirmCode
 RPC_Auth.SingUp
 RPC_Auth.SingIn
 RPC_Auth.LogOut



 RPC_Chat.AddNewMessage
 RPC_Chat.SetRoomActionDoing
 RPC_Chat.GetChatList
 RPC_Chat.GetChatHistory
 RPC_Chat.PushRoomsChange
 RPC_Chat.GetRoomsChange



 RPC_General.Echo
 RPC_General.CheckUserName



 RPC_Page.GetCommentsPage
 RPC_Page.GetHomePage
 RPC_Page.GetProfileAbout
 RPC_Page.GetProfileAllShared
 RPC_Page.GetProfileByCategoryPage
 RPC_Page.GetLikesPage
 RPC_Page.GetFollowersPage
 RPC_Page.GetFollowingsPage
 RPC_Page.GetNotifiesPage
 RPC_Page.GetUserActionsPage
 RPC_Page.GetPromotedPostsPage
 RPC_Page.GetSuggestedUsersPage
 RPC_Page.GetSuggestedTagsPage
 RPC_Page.GetLastPostsPage
 RPC_Page.GetLastTagPage
 RPC_Page.SearchTagsPage
 RPC_Page.SearchUsersPage



 RPC_Social.AddComment
 RPC_Social.DeleteComment
 RPC_Social.EditComment
 RPC_Social.AddPost
 RPC_Social.EditPost
 RPC_Social.DeletePost
 RPC_Social.ArchivePost
 RPC_Social.LikePost
 RPC_Social.UnLikePost
 RPC_Social.FollowUser
 RPC_Social.UnFollowUser
 RPC_Social.BlockUser
 RPC_Social.UnBlockUser
 RPC_Social.AddSeenPosts



 RPC_User.UpdateAbout
 RPC_User.UpdateUserName
 RPC_User.ChangePrivacy
 RPC_User.ChangeAvatar


*/
