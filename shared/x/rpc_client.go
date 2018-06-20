package x

import "github.com/golang/protobuf/proto"

//type RPCClientHandler func(cmd string, pb interface{}) interface{}
type RPCClientHandler func(cmdSre string, pbIn, pbOut proto.Message) error

// all clients struc
var RPC_AllClinetsPlay = struct {
	RPC_Auth    RPC_Auth_Client
	RPC_Chat    RPC_Chat_Client
	RPC_General RPC_General_Client
	RPC_Page    RPC_Page_Client
	RPC_Social  RPC_Social_Client
	RPC_User    RPC_User_Client
}{
	RPC_Auth:    RPC_Auth_Client(0),
	RPC_Chat:    RPC_Chat_Client(0),
	RPC_General: RPC_General_Client(0),
	RPC_Page:    RPC_Page_Client(0),
	RPC_Social:  RPC_Social_Client(0),
	RPC_User:    RPC_User_Client(0),
}

// client types defs
type RPC_Auth_Client int
type RPC_Chat_Client int
type RPC_General_Client int
type RPC_Page_Client int
type RPC_Social_Client int
type RPC_User_Client int

/////////////// methods ////////////////

// service: RPC_Auth

func (RPC_Auth_Client) SendConfirmCode(param *RPC_Auth_Types_SendConfirmCode_Param, fn RPCClientHandler) (*RPC_Auth_Types_SendConfirmCode_Response, error) {
	out := &RPC_Auth_Types_SendConfirmCode_Response{}
	err := fn("RPC_Auth.SendConfirmCode", param, out)
	if err == nil {
		return out, nil
	}
	return nil, err
}

func (RPC_Auth_Client) ConfirmCode(param *RPC_Auth_Types_ConfirmCode_Param, fn RPCClientHandler) (*RPC_Auth_Types_ConfirmCode_Response, error) {
	out := &RPC_Auth_Types_ConfirmCode_Response{}
	err := fn("RPC_Auth.ConfirmCode", param, out)
	if err == nil {
		return out, nil
	}
	return nil, err
}

func (RPC_Auth_Client) SingUp(param *RPC_Auth_Types_SingUp_Param, fn RPCClientHandler) (*RPC_Auth_Types_SingUp_Response, error) {
	out := &RPC_Auth_Types_SingUp_Response{}
	err := fn("RPC_Auth.SingUp", param, out)
	if err == nil {
		return out, nil
	}
	return nil, err
}

func (RPC_Auth_Client) SingIn(param *RPC_Auth_Types_SingIn_Param, fn RPCClientHandler) (*RPC_Auth_Types_SingIn_Response, error) {
	out := &RPC_Auth_Types_SingIn_Response{}
	err := fn("RPC_Auth.SingIn", param, out)
	if err == nil {
		return out, nil
	}
	return nil, err
}

func (RPC_Auth_Client) LogOut(param *RPC_Auth_Types_LogOut_Param, fn RPCClientHandler) (*RPC_Auth_Types_LogOut_Response, error) {
	out := &RPC_Auth_Types_LogOut_Response{}
	err := fn("RPC_Auth.LogOut", param, out)
	if err == nil {
		return out, nil
	}
	return nil, err
}

// service: RPC_Chat

func (RPC_Chat_Client) AddNewMessage(param *PB_RPC_Chat_Types_AddNewMessage_Param, fn RPCClientHandler) (*PB_RPC_Chat_Types_AddNewMessage_Response, error) {
	out := &PB_RPC_Chat_Types_AddNewMessage_Response{}
	err := fn("RPC_Chat.AddNewMessage", param, out)
	if err == nil {
		return out, nil
	}
	return nil, err
}

func (RPC_Chat_Client) SetRoomActionDoing(param *PB_RPC_Chat_Types_SetRoomActionDoing_Param, fn RPCClientHandler) (*PB_RPC_Chat_Types_SetRoomActionDoing_Response, error) {
	out := &PB_RPC_Chat_Types_SetRoomActionDoing_Response{}
	err := fn("RPC_Chat.SetRoomActionDoing", param, out)
	if err == nil {
		return out, nil
	}
	return nil, err
}

func (RPC_Chat_Client) GetChatList(param *PB_RPC_Chat_Types_GetChatList_Param, fn RPCClientHandler) (*PB_RPC_Chat_Types_GetChatList_Response, error) {
	out := &PB_RPC_Chat_Types_GetChatList_Response{}
	err := fn("RPC_Chat.GetChatList", param, out)
	if err == nil {
		return out, nil
	}
	return nil, err
}

func (RPC_Chat_Client) GetChatHistory(param *PB_RPC_Chat_Types_GetChatHistory_Param, fn RPCClientHandler) (*PB_RPC_Chat_Types_GetChatHistory_Response, error) {
	out := &PB_RPC_Chat_Types_GetChatHistory_Response{}
	err := fn("RPC_Chat.GetChatHistory", param, out)
	if err == nil {
		return out, nil
	}
	return nil, err
}

func (RPC_Chat_Client) AddRoomsChange(param *PB_RPC_Chat_Types_AddRoomsChange_Param, fn RPCClientHandler) (*PB_RPC_Chat_Types_AddRoomsChange_Response, error) {
	out := &PB_RPC_Chat_Types_AddRoomsChange_Response{}
	err := fn("RPC_Chat.AddRoomsChange", param, out)
	if err == nil {
		return out, nil
	}
	return nil, err
}

func (RPC_Chat_Client) GetRoomsChange(param *PB_RPC_Chat_Types_GetRoomsChange_Param, fn RPCClientHandler) (*PB_RPC_Chat_Types_GetRoomsChange_Response, error) {
	out := &PB_RPC_Chat_Types_GetRoomsChange_Response{}
	err := fn("RPC_Chat.GetRoomsChange", param, out)
	if err == nil {
		return out, nil
	}
	return nil, err
}

// service: RPC_General

func (RPC_General_Client) Echo(param *RPC_General_Types_Echo_Param, fn RPCClientHandler) (*RPC_General_Types_Echo_Response, error) {
	out := &RPC_General_Types_Echo_Response{}
	err := fn("RPC_General.Echo", param, out)
	if err == nil {
		return out, nil
	}
	return nil, err
}

func (RPC_General_Client) CheckUserName(param *RPC_General_Types_CheckUserName_Param, fn RPCClientHandler) (*RPC_General_Types_CheckUserName_Response, error) {
	out := &RPC_General_Types_CheckUserName_Response{}
	err := fn("RPC_General.CheckUserName", param, out)
	if err == nil {
		return out, nil
	}
	return nil, err
}

// service: RPC_Page

func (RPC_Page_Client) GetCommentsPage(param *RPC_Page_Types_GetCommentsPage_Param, fn RPCClientHandler) (*RPC_Page_Types_GetCommentsPage_Response, error) {
	out := &RPC_Page_Types_GetCommentsPage_Response{}
	err := fn("RPC_Page.GetCommentsPage", param, out)
	if err == nil {
		return out, nil
	}
	return nil, err
}

func (RPC_Page_Client) GetHomePage(param *RPC_Page_Types_GetHomePage_Param, fn RPCClientHandler) (*RPC_Page_Types_GetHomePage_Response, error) {
	out := &RPC_Page_Types_GetHomePage_Response{}
	err := fn("RPC_Page.GetHomePage", param, out)
	if err == nil {
		return out, nil
	}
	return nil, err
}

func (RPC_Page_Client) GetProfileAbout(param *RPC_Page_Types_GetProfileAbout_Param, fn RPCClientHandler) (*RPC_Page_Types_GetProfileAbout_Response, error) {
	out := &RPC_Page_Types_GetProfileAbout_Response{}
	err := fn("RPC_Page.GetProfileAbout", param, out)
	if err == nil {
		return out, nil
	}
	return nil, err
}

func (RPC_Page_Client) GetProfileAllShared(param *RPC_Page_Types_GetProfileAllShared_Param, fn RPCClientHandler) (*RPC_Page_Types_GetProfileAllShared_Response, error) {
	out := &RPC_Page_Types_GetProfileAllShared_Response{}
	err := fn("RPC_Page.GetProfileAllShared", param, out)
	if err == nil {
		return out, nil
	}
	return nil, err
}

func (RPC_Page_Client) GetProfileByCategoryPage(param *RPC_Page_Types_GetProfileByCategoryPage_Param, fn RPCClientHandler) (*RPC_Page_Types_GetProfileByCategoryPage_Response, error) {
	out := &RPC_Page_Types_GetProfileByCategoryPage_Response{}
	err := fn("RPC_Page.GetProfileByCategoryPage", param, out)
	if err == nil {
		return out, nil
	}
	return nil, err
}

func (RPC_Page_Client) GetLikesPage(param *RPC_Page_Types_GetLikesPage_Param, fn RPCClientHandler) (*RPC_Page_Types_GetLikesPage_Response, error) {
	out := &RPC_Page_Types_GetLikesPage_Response{}
	err := fn("RPC_Page.GetLikesPage", param, out)
	if err == nil {
		return out, nil
	}
	return nil, err
}

func (RPC_Page_Client) GetFollowersPage(param *RPC_Page_Types_GetFollowersPage_Param, fn RPCClientHandler) (*RPC_Page_Types_GetFollowersPage_Response, error) {
	out := &RPC_Page_Types_GetFollowersPage_Response{}
	err := fn("RPC_Page.GetFollowersPage", param, out)
	if err == nil {
		return out, nil
	}
	return nil, err
}

func (RPC_Page_Client) GetFollowingsPage(param *RPC_Page_Types_GetFollowingsPage_Param, fn RPCClientHandler) (*RPC_Page_Types_GetFollowingsPage_Response, error) {
	out := &RPC_Page_Types_GetFollowingsPage_Response{}
	err := fn("RPC_Page.GetFollowingsPage", param, out)
	if err == nil {
		return out, nil
	}
	return nil, err
}

func (RPC_Page_Client) GetNotifiesPage(param *RPC_Page_Types_GetNotifiesPage_Param, fn RPCClientHandler) (*RPC_Page_Types_GetNotifiesPage_Response, error) {
	out := &RPC_Page_Types_GetNotifiesPage_Response{}
	err := fn("RPC_Page.GetNotifiesPage", param, out)
	if err == nil {
		return out, nil
	}
	return nil, err
}

func (RPC_Page_Client) GetUserActionsPage(param *RPC_Page_Types_GetUserActionsPage_Param, fn RPCClientHandler) (*RPC_Page_Types_GetUserActionsPage_Response, error) {
	out := &RPC_Page_Types_GetUserActionsPage_Response{}
	err := fn("RPC_Page.GetUserActionsPage", param, out)
	if err == nil {
		return out, nil
	}
	return nil, err
}

func (RPC_Page_Client) GetPromotedPostsPage(param *RPC_Page_Types_GetPromotedPostsPage_Param, fn RPCClientHandler) (*RPC_Page_Types_GetPromotedPostsPage_Response, error) {
	out := &RPC_Page_Types_GetPromotedPostsPage_Response{}
	err := fn("RPC_Page.GetPromotedPostsPage", param, out)
	if err == nil {
		return out, nil
	}
	return nil, err
}

func (RPC_Page_Client) GetSuggestedUsersPage(param *RPC_Page_Types_GetSuggestedUsersPage_Param, fn RPCClientHandler) (*RPC_Page_Types_GetSuggestedUsersPage_Response, error) {
	out := &RPC_Page_Types_GetSuggestedUsersPage_Response{}
	err := fn("RPC_Page.GetSuggestedUsersPage", param, out)
	if err == nil {
		return out, nil
	}
	return nil, err
}

func (RPC_Page_Client) GetSuggestedTagsPage(param *RPC_Page_Types_GetSuggestedTagsPage_Param, fn RPCClientHandler) (*RPC_Page_Types_GetSuggestedTagsPage_Response, error) {
	out := &RPC_Page_Types_GetSuggestedTagsPage_Response{}
	err := fn("RPC_Page.GetSuggestedTagsPage", param, out)
	if err == nil {
		return out, nil
	}
	return nil, err
}

func (RPC_Page_Client) GetLastPostsPage(param *RPC_Page_Types_GetLastPostsPage_Param, fn RPCClientHandler) (*RPC_Page_Types_GetLastPostsPage_Response, error) {
	out := &RPC_Page_Types_GetLastPostsPage_Response{}
	err := fn("RPC_Page.GetLastPostsPage", param, out)
	if err == nil {
		return out, nil
	}
	return nil, err
}

func (RPC_Page_Client) GetLastTagPage(param *RPC_Page_Types_GetLastTagPage_Param, fn RPCClientHandler) (*RPC_Page_Types_GetLastTagPage_Response, error) {
	out := &RPC_Page_Types_GetLastTagPage_Response{}
	err := fn("RPC_Page.GetLastTagPage", param, out)
	if err == nil {
		return out, nil
	}
	return nil, err
}

func (RPC_Page_Client) SearchTagsPage(param *RPC_Page_Types_SearchTagsPage_Param, fn RPCClientHandler) (*RPC_Page_Types_SearchTagsPage_Response, error) {
	out := &RPC_Page_Types_SearchTagsPage_Response{}
	err := fn("RPC_Page.SearchTagsPage", param, out)
	if err == nil {
		return out, nil
	}
	return nil, err
}

func (RPC_Page_Client) SearchUsersPage(param *RPC_Page_Types_SearchUsersPage_Param, fn RPCClientHandler) (*RPC_Page_Types_SearchUsersPage_Response, error) {
	out := &RPC_Page_Types_SearchUsersPage_Response{}
	err := fn("RPC_Page.SearchUsersPage", param, out)
	if err == nil {
		return out, nil
	}
	return nil, err
}

// service: RPC_Social

func (RPC_Social_Client) AddComment(param *RPC_Social_Types_AddComment_Param, fn RPCClientHandler) (*RPC_Social_Types_AddComment_Response, error) {
	out := &RPC_Social_Types_AddComment_Response{}
	err := fn("RPC_Social.AddComment", param, out)
	if err == nil {
		return out, nil
	}
	return nil, err
}

func (RPC_Social_Client) DeleteComment(param *RPC_Social_Types_DeleteComment_Param, fn RPCClientHandler) (*RPC_Social_Types_DeleteComment_Response, error) {
	out := &RPC_Social_Types_DeleteComment_Response{}
	err := fn("RPC_Social.DeleteComment", param, out)
	if err == nil {
		return out, nil
	}
	return nil, err
}

func (RPC_Social_Client) EditComment(param *RPC_Social_Types_EditComment_Param, fn RPCClientHandler) (*RPC_Social_Types_EditComment_Response, error) {
	out := &RPC_Social_Types_EditComment_Response{}
	err := fn("RPC_Social.EditComment", param, out)
	if err == nil {
		return out, nil
	}
	return nil, err
}

func (RPC_Social_Client) LikeComment(param *RPC_Social_Types_LikeComment_Param, fn RPCClientHandler) (*RPC_Social_Types_LikeComment_Response, error) {
	out := &RPC_Social_Types_LikeComment_Response{}
	err := fn("RPC_Social.LikeComment", param, out)
	if err == nil {
		return out, nil
	}
	return nil, err
}

func (RPC_Social_Client) AddPost(param *RPC_Social_Types_AddPost_Param, fn RPCClientHandler) (*RPC_Social_Types_AddPost_Response, error) {
	out := &RPC_Social_Types_AddPost_Response{}
	err := fn("RPC_Social.AddPost", param, out)
	if err == nil {
		return out, nil
	}
	return nil, err
}

func (RPC_Social_Client) EditPost(param *RPC_Social_Types_EditPost_Param, fn RPCClientHandler) (*RPC_Social_Types_EditPost_Response, error) {
	out := &RPC_Social_Types_EditPost_Response{}
	err := fn("RPC_Social.EditPost", param, out)
	if err == nil {
		return out, nil
	}
	return nil, err
}

func (RPC_Social_Client) DeletePost(param *RPC_Social_Types_DeletePost_Param, fn RPCClientHandler) (*RPC_Social_Types_DeletePost_Response, error) {
	out := &RPC_Social_Types_DeletePost_Response{}
	err := fn("RPC_Social.DeletePost", param, out)
	if err == nil {
		return out, nil
	}
	return nil, err
}

func (RPC_Social_Client) ArchivePost(param *RPC_Social_Types_ArchivePost_Param, fn RPCClientHandler) (*RPC_Social_Types_ArchivePost_Response, error) {
	out := &RPC_Social_Types_ArchivePost_Response{}
	err := fn("RPC_Social.ArchivePost", param, out)
	if err == nil {
		return out, nil
	}
	return nil, err
}

func (RPC_Social_Client) PromotePost(param *RPC_Social_Types_PromotePost_Param, fn RPCClientHandler) (*RPC_Social_Types_PromotePost_Response, error) {
	out := &RPC_Social_Types_PromotePost_Response{}
	err := fn("RPC_Social.PromotePost", param, out)
	if err == nil {
		return out, nil
	}
	return nil, err
}

func (RPC_Social_Client) LikePost(param *RPC_Social_Types_LikePost_Param, fn RPCClientHandler) (*RPC_Social_Types_LikePost_Response, error) {
	out := &RPC_Social_Types_LikePost_Response{}
	err := fn("RPC_Social.LikePost", param, out)
	if err == nil {
		return out, nil
	}
	return nil, err
}

func (RPC_Social_Client) UnLikePost(param *RPC_Social_Types_UnLikePost_Param, fn RPCClientHandler) (*RPC_Social_Types_UnLikePost_Response, error) {
	out := &RPC_Social_Types_UnLikePost_Response{}
	err := fn("RPC_Social.UnLikePost", param, out)
	if err == nil {
		return out, nil
	}
	return nil, err
}

func (RPC_Social_Client) FollowUser(param *RPC_Social_Types_FollowUser_Param, fn RPCClientHandler) (*RPC_Social_Types_FollowUser_Response, error) {
	out := &RPC_Social_Types_FollowUser_Response{}
	err := fn("RPC_Social.FollowUser", param, out)
	if err == nil {
		return out, nil
	}
	return nil, err
}

func (RPC_Social_Client) UnFollowUser(param *RPC_Social_Types_UnFollowUser_Param, fn RPCClientHandler) (*RPC_Social_Types_UnFollowUser_Response, error) {
	out := &RPC_Social_Types_UnFollowUser_Response{}
	err := fn("RPC_Social.UnFollowUser", param, out)
	if err == nil {
		return out, nil
	}
	return nil, err
}

func (RPC_Social_Client) BlockUser(param *RPC_Social_Types_BlockUser_Param, fn RPCClientHandler) (*RPC_Social_Types_BlockUser_Response, error) {
	out := &RPC_Social_Types_BlockUser_Response{}
	err := fn("RPC_Social.BlockUser", param, out)
	if err == nil {
		return out, nil
	}
	return nil, err
}

func (RPC_Social_Client) UnBlockUser(param *RPC_Social_Types_UnBlockUser_Param, fn RPCClientHandler) (*RPC_Social_Types_UnBlockUser_Response, error) {
	out := &RPC_Social_Types_UnBlockUser_Response{}
	err := fn("RPC_Social.UnBlockUser", param, out)
	if err == nil {
		return out, nil
	}
	return nil, err
}

func (RPC_Social_Client) AddSeenPosts(param *RPC_Social_Types_AddSeenPosts_Param, fn RPCClientHandler) (*RPC_Social_Types_AddSeenPosts_Response, error) {
	out := &RPC_Social_Types_AddSeenPosts_Response{}
	err := fn("RPC_Social.AddSeenPosts", param, out)
	if err == nil {
		return out, nil
	}
	return nil, err
}

// service: RPC_User

func (RPC_User_Client) UpdateAbout(param *RPC_User_Types_UpdateAbout_Param, fn RPCClientHandler) (*RPC_User_Types_UpdateAbout_Response, error) {
	out := &RPC_User_Types_UpdateAbout_Response{}
	err := fn("RPC_User.UpdateAbout", param, out)
	if err == nil {
		return out, nil
	}
	return nil, err
}

func (RPC_User_Client) UpdateUserName(param *RPC_User_Types_UpdateUserName_Param, fn RPCClientHandler) (*RPC_User_Types_UpdateUserName_Response, error) {
	out := &RPC_User_Types_UpdateUserName_Response{}
	err := fn("RPC_User.UpdateUserName", param, out)
	if err == nil {
		return out, nil
	}
	return nil, err
}

func (RPC_User_Client) ChangePrivacy(param *RPC_User_Types_ChangePrivacy_Param, fn RPCClientHandler) (*RPC_User_Types_ChangePrivacy_Response, error) {
	out := &RPC_User_Types_ChangePrivacy_Response{}
	err := fn("RPC_User.ChangePrivacy", param, out)
	if err == nil {
		return out, nil
	}
	return nil, err
}

func (RPC_User_Client) ChangeAvatar(param *RPC_User_Types_ChangeAvatar_Param, fn RPCClientHandler) (*RPC_User_Types_ChangeAvatar_Response, error) {
	out := &RPC_User_Types_ChangeAvatar_Response{}
	err := fn("RPC_User.ChangeAvatar", param, out)
	if err == nil {
		return out, nil
	}
	return nil, err
}
