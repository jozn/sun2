package x

import "github.com/golang/protobuf/proto"

//type RPCClientHandler func(cmd string, pb interface{}) interface{}
type RPCClientHandler func(cmdSre string, pbIn, pbOut proto.Message) error

// all clients struc
var RPC_AllClinetsPlay = struct {
	RPC_Auth   RPC_Auth_Client
	RPC_Chat   RPC_Chat_Client
	RPC_Other  RPC_Other_Client
	RPC_Page   RPC_Page_Client
	RPC_Search RPC_Search_Client
	RPC_Social RPC_Social_Client
	RPC_User   RPC_User_Client
}{
	RPC_Auth:   RPC_Auth_Client(0),
	RPC_Chat:   RPC_Chat_Client(0),
	RPC_Other:  RPC_Other_Client(0),
	RPC_Page:   RPC_Page_Client(0),
	RPC_Search: RPC_Search_Client(0),
	RPC_Social: RPC_Social_Client(0),
	RPC_User:   RPC_User_Client(0),
}

// client types defs
type RPC_Auth_Client int
type RPC_Chat_Client int
type RPC_Other_Client int
type RPC_Page_Client int
type RPC_Search_Client int
type RPC_Social_Client int
type RPC_User_Client int

/////////////// methods ////////////////

// service: RPC_Auth

func (RPC_Auth_Client) CheckPhone(param *PB_UserParam_CheckUserName2, fn RPCClientHandler) (*PB_UserResponse_CheckUserName2, error) {
	out := &PB_UserResponse_CheckUserName2{}
	err := fn("RPC_Auth.CheckPhone", param, out)
	if err == nil {
		return out, nil
	}
	return nil, err
}

func (RPC_Auth_Client) SendCode(param *PB_UserParam_CheckUserName2, fn RPCClientHandler) (*PB_UserResponse_CheckUserName2, error) {
	out := &PB_UserResponse_CheckUserName2{}
	err := fn("RPC_Auth.SendCode", param, out)
	if err == nil {
		return out, nil
	}
	return nil, err
}

func (RPC_Auth_Client) SendCodeToSms(param *PB_UserParam_CheckUserName2, fn RPCClientHandler) (*PB_UserResponse_CheckUserName2, error) {
	out := &PB_UserResponse_CheckUserName2{}
	err := fn("RPC_Auth.SendCodeToSms", param, out)
	if err == nil {
		return out, nil
	}
	return nil, err
}

func (RPC_Auth_Client) SendCodeToTelgram(param *PB_UserParam_CheckUserName2, fn RPCClientHandler) (*PB_UserResponse_CheckUserName2, error) {
	out := &PB_UserResponse_CheckUserName2{}
	err := fn("RPC_Auth.SendCodeToTelgram", param, out)
	if err == nil {
		return out, nil
	}
	return nil, err
}

func (RPC_Auth_Client) SingUp(param *PB_UserParam_CheckUserName2, fn RPCClientHandler) (*PB_UserResponse_CheckUserName2, error) {
	out := &PB_UserResponse_CheckUserName2{}
	err := fn("RPC_Auth.SingUp", param, out)
	if err == nil {
		return out, nil
	}
	return nil, err
}

func (RPC_Auth_Client) SingIn(param *PB_UserParam_CheckUserName2, fn RPCClientHandler) (*PB_UserResponse_CheckUserName2, error) {
	out := &PB_UserResponse_CheckUserName2{}
	err := fn("RPC_Auth.SingIn", param, out)
	if err == nil {
		return out, nil
	}
	return nil, err
}

func (RPC_Auth_Client) LogOut(param *PB_UserParam_CheckUserName2, fn RPCClientHandler) (*PB_UserResponse_CheckUserName2, error) {
	out := &PB_UserResponse_CheckUserName2{}
	err := fn("RPC_Auth.LogOut", param, out)
	if err == nil {
		return out, nil
	}
	return nil, err
}

// service: RPC_Chat

func (RPC_Chat_Client) AddNewMessage(param *PB_ChatParam_AddNewMessage, fn RPCClientHandler) (*PB_ChatResponse_AddNewMessage, error) {
	out := &PB_ChatResponse_AddNewMessage{}
	err := fn("RPC_Chat.AddNewMessage", param, out)
	if err == nil {
		return out, nil
	}
	return nil, err
}

func (RPC_Chat_Client) SetRoomActionDoing(param *PB_ChatParam_SetRoomActionDoing, fn RPCClientHandler) (*PB_ChatResponse_SetRoomActionDoing, error) {
	out := &PB_ChatResponse_SetRoomActionDoing{}
	err := fn("RPC_Chat.SetRoomActionDoing", param, out)
	if err == nil {
		return out, nil
	}
	return nil, err
}

func (RPC_Chat_Client) SetMessagesAsReceived(param *PB_ChatParam_SetMessagesAsReceived, fn RPCClientHandler) (*PB_ChatResponse_SetMessagesAsReceived, error) {
	out := &PB_ChatResponse_SetMessagesAsReceived{}
	err := fn("RPC_Chat.SetMessagesAsReceived", param, out)
	if err == nil {
		return out, nil
	}
	return nil, err
}

func (RPC_Chat_Client) SetMessagesRangeAsSeen(param *PB_ChatParam_SetChatMessagesRangeAsSeen, fn RPCClientHandler) (*PB_ChatResponse_SetChatMessagesRangeAsSeen, error) {
	out := &PB_ChatResponse_SetChatMessagesRangeAsSeen{}
	err := fn("RPC_Chat.SetMessagesRangeAsSeen", param, out)
	if err == nil {
		return out, nil
	}
	return nil, err
}

func (RPC_Chat_Client) DeleteChatHistory(param *PB_ChatParam_DeleteChatHistory, fn RPCClientHandler) (*PB_ChatResponse_DeleteChatHistory, error) {
	out := &PB_ChatResponse_DeleteChatHistory{}
	err := fn("RPC_Chat.DeleteChatHistory", param, out)
	if err == nil {
		return out, nil
	}
	return nil, err
}

func (RPC_Chat_Client) DeleteMessagesByIds(param *PB_ChatParam_DeleteMessagesByIds, fn RPCClientHandler) (*PB_ChatResponse_DeleteMessagesByIds, error) {
	out := &PB_ChatResponse_DeleteMessagesByIds{}
	err := fn("RPC_Chat.DeleteMessagesByIds", param, out)
	if err == nil {
		return out, nil
	}
	return nil, err
}

func (RPC_Chat_Client) EditMessage(param *PB_ChatParam_EditMessage, fn RPCClientHandler) (*PB_ChatResponse_EditMessage, error) {
	out := &PB_ChatResponse_EditMessage{}
	err := fn("RPC_Chat.EditMessage", param, out)
	if err == nil {
		return out, nil
	}
	return nil, err
}

func (RPC_Chat_Client) GetChatList(param *PB_ChatParam_GetChatList, fn RPCClientHandler) (*PB_ChatResponse_GetChatList, error) {
	out := &PB_ChatResponse_GetChatList{}
	err := fn("RPC_Chat.GetChatList", param, out)
	if err == nil {
		return out, nil
	}
	return nil, err
}

func (RPC_Chat_Client) GetChatHistoryToOlder(param *PB_ChatParam_GetChatHistoryToOlder, fn RPCClientHandler) (*PB_ChatResponse_GetChatHistoryToOlder, error) {
	out := &PB_ChatResponse_GetChatHistoryToOlder{}
	err := fn("RPC_Chat.GetChatHistoryToOlder", param, out)
	if err == nil {
		return out, nil
	}
	return nil, err
}

// service: RPC_Other

func (RPC_Other_Client) Echo(param *PB_OtherParam_Echo, fn RPCClientHandler) (*PB_OtherResponse_Echo, error) {
	out := &PB_OtherResponse_Echo{}
	err := fn("RPC_Other.Echo", param, out)
	if err == nil {
		return out, nil
	}
	return nil, err
}

// service: RPC_Page

func (RPC_Page_Client) GetCommentsPage(param *PB_PageParam_GetCommentsPage, fn RPCClientHandler) (*PB_PageResponse_GetCommentsPage, error) {
	out := &PB_PageResponse_GetCommentsPage{}
	err := fn("RPC_Page.GetCommentsPage", param, out)
	if err == nil {
		return out, nil
	}
	return nil, err
}

func (RPC_Page_Client) GetHomePage(param *PB_PageParam_GetHomePage, fn RPCClientHandler) (*PB_PageResponse_GetHomePage, error) {
	out := &PB_PageResponse_GetHomePage{}
	err := fn("RPC_Page.GetHomePage", param, out)
	if err == nil {
		return out, nil
	}
	return nil, err
}

func (RPC_Page_Client) GetProfilePage(param *PB_PageParam_GetProfilePage, fn RPCClientHandler) (*PB_PageResponse_GetProfilePage, error) {
	out := &PB_PageResponse_GetProfilePage{}
	err := fn("RPC_Page.GetProfilePage", param, out)
	if err == nil {
		return out, nil
	}
	return nil, err
}

func (RPC_Page_Client) GetLikesPage(param *PB_PageParam_GetLikesPage, fn RPCClientHandler) (*PB_PageResponse_GetLikesPage, error) {
	out := &PB_PageResponse_GetLikesPage{}
	err := fn("RPC_Page.GetLikesPage", param, out)
	if err == nil {
		return out, nil
	}
	return nil, err
}

func (RPC_Page_Client) GetFollowersPage(param *PB_PageParam_GetFollowersPage, fn RPCClientHandler) (*PB_PageResponse_GetFollowersPage, error) {
	out := &PB_PageResponse_GetFollowersPage{}
	err := fn("RPC_Page.GetFollowersPage", param, out)
	if err == nil {
		return out, nil
	}
	return nil, err
}

func (RPC_Page_Client) GetFollowingsPage(param *PB_PageParam_GetFollowingsPage, fn RPCClientHandler) (*PB_PageResponse_GetFollowingsPage, error) {
	out := &PB_PageResponse_GetFollowingsPage{}
	err := fn("RPC_Page.GetFollowingsPage", param, out)
	if err == nil {
		return out, nil
	}
	return nil, err
}

func (RPC_Page_Client) GetNotifiesPage(param *PB_PageParam_GetNotifiesPage, fn RPCClientHandler) (*PB_PageResponse_GetNotifiesPage, error) {
	out := &PB_PageResponse_GetNotifiesPage{}
	err := fn("RPC_Page.GetNotifiesPage", param, out)
	if err == nil {
		return out, nil
	}
	return nil, err
}

func (RPC_Page_Client) GetUserActionsPage(param *PB_PageParam_GetUserActionsPage, fn RPCClientHandler) (*PB_PageResponse_GetUserActionsPage, error) {
	out := &PB_PageResponse_GetUserActionsPage{}
	err := fn("RPC_Page.GetUserActionsPage", param, out)
	if err == nil {
		return out, nil
	}
	return nil, err
}

func (RPC_Page_Client) GetSuggestedPostsPage(param *PB_PageParam_GetSuggestedPostsPage, fn RPCClientHandler) (*PB_PageResponse_GetSuggestedPostsPage, error) {
	out := &PB_PageResponse_GetSuggestedPostsPage{}
	err := fn("RPC_Page.GetSuggestedPostsPage", param, out)
	if err == nil {
		return out, nil
	}
	return nil, err
}

func (RPC_Page_Client) GetSuggestedUsersPage(param *PB_PageParam_GetSuggestedUsersPage, fn RPCClientHandler) (*PB_PageResponse_GetSuggestedUsersPage, error) {
	out := &PB_PageResponse_GetSuggestedUsersPage{}
	err := fn("RPC_Page.GetSuggestedUsersPage", param, out)
	if err == nil {
		return out, nil
	}
	return nil, err
}

func (RPC_Page_Client) GetSuggestedTagsPage(param *PB_PageParam_GetSuggestedTagsPage, fn RPCClientHandler) (*PB_PageResponse_GetSuggestedTagsPage, error) {
	out := &PB_PageResponse_GetSuggestedTagsPage{}
	err := fn("RPC_Page.GetSuggestedTagsPage", param, out)
	if err == nil {
		return out, nil
	}
	return nil, err
}

func (RPC_Page_Client) GetLastPostsPage(param *PB_PageParam_GetLastPostsPage, fn RPCClientHandler) (*PB_PageResponse_GetLastPostsPage, error) {
	out := &PB_PageResponse_GetLastPostsPage{}
	err := fn("RPC_Page.GetLastPostsPage", param, out)
	if err == nil {
		return out, nil
	}
	return nil, err
}

func (RPC_Page_Client) GetTagPage(param *PB_PageParam_GetTagPage, fn RPCClientHandler) (*PB_PageResponse_GetTagPage, error) {
	out := &PB_PageResponse_GetTagPage{}
	err := fn("RPC_Page.GetTagPage", param, out)
	if err == nil {
		return out, nil
	}
	return nil, err
}

func (RPC_Page_Client) SearchTagsPage(param *PB_PageParam_SearchTagsPage, fn RPCClientHandler) (*PB_PageResponse_SearchTagsPage, error) {
	out := &PB_PageResponse_SearchTagsPage{}
	err := fn("RPC_Page.SearchTagsPage", param, out)
	if err == nil {
		return out, nil
	}
	return nil, err
}

func (RPC_Page_Client) SearchUsersPage(param *PB_PageParam_SearchUsersPage, fn RPCClientHandler) (*PB_PageResponse_SearchUsersPage, error) {
	out := &PB_PageResponse_SearchUsersPage{}
	err := fn("RPC_Page.SearchUsersPage", param, out)
	if err == nil {
		return out, nil
	}
	return nil, err
}

// service: RPC_Search

func (RPC_Search_Client) SearchTags(param *PB_SearchResponse_AddNewC, fn RPCClientHandler) (*PB_SearchResponse_AddNewC, error) {
	out := &PB_SearchResponse_AddNewC{}
	err := fn("RPC_Search.SearchTags", param, out)
	if err == nil {
		return out, nil
	}
	return nil, err
}

// service: RPC_Social

func (RPC_Social_Client) AddComment(param *PB_SocialParam_AddComment, fn RPCClientHandler) (*PB_SocialResponse_AddComment, error) {
	out := &PB_SocialResponse_AddComment{}
	err := fn("RPC_Social.AddComment", param, out)
	if err == nil {
		return out, nil
	}
	return nil, err
}

func (RPC_Social_Client) DeleteComment(param *PB_SocialParam_DeleteComment, fn RPCClientHandler) (*PB_SocialResponse_DeleteComment, error) {
	out := &PB_SocialResponse_DeleteComment{}
	err := fn("RPC_Social.DeleteComment", param, out)
	if err == nil {
		return out, nil
	}
	return nil, err
}

func (RPC_Social_Client) AddPost(param *PB_SocialParam_AddPost, fn RPCClientHandler) (*PB_SocialResponse_AddPost, error) {
	out := &PB_SocialResponse_AddPost{}
	err := fn("RPC_Social.AddPost", param, out)
	if err == nil {
		return out, nil
	}
	return nil, err
}

func (RPC_Social_Client) EditPost(param *PB_SocialParam_EditPost, fn RPCClientHandler) (*PB_SocialResponse_EditPost, error) {
	out := &PB_SocialResponse_EditPost{}
	err := fn("RPC_Social.EditPost", param, out)
	if err == nil {
		return out, nil
	}
	return nil, err
}

func (RPC_Social_Client) DeletePost(param *PB_SocialParam_DeletePost, fn RPCClientHandler) (*PB_SocialResponse_DeletePost, error) {
	out := &PB_SocialResponse_DeletePost{}
	err := fn("RPC_Social.DeletePost", param, out)
	if err == nil {
		return out, nil
	}
	return nil, err
}

func (RPC_Social_Client) ArchivePost(param *PB_SocialParam_ArchivePost, fn RPCClientHandler) (*PB_SocialResponse_ArchivePost, error) {
	out := &PB_SocialResponse_ArchivePost{}
	err := fn("RPC_Social.ArchivePost", param, out)
	if err == nil {
		return out, nil
	}
	return nil, err
}

func (RPC_Social_Client) LikePost(param *PB_SocialParam_LikePost, fn RPCClientHandler) (*PB_SocialResponse_LikePost, error) {
	out := &PB_SocialResponse_LikePost{}
	err := fn("RPC_Social.LikePost", param, out)
	if err == nil {
		return out, nil
	}
	return nil, err
}

func (RPC_Social_Client) UnLikePost(param *PB_SocialParam_UnLikePost, fn RPCClientHandler) (*PB_SocialResponse_UnLikePost, error) {
	out := &PB_SocialResponse_UnLikePost{}
	err := fn("RPC_Social.UnLikePost", param, out)
	if err == nil {
		return out, nil
	}
	return nil, err
}

func (RPC_Social_Client) FollowUser(param *PB_SocialParam_FollowUser, fn RPCClientHandler) (*PB_SocialResponse_FollowUser, error) {
	out := &PB_SocialResponse_FollowUser{}
	err := fn("RPC_Social.FollowUser", param, out)
	if err == nil {
		return out, nil
	}
	return nil, err
}

func (RPC_Social_Client) UnFollowUser(param *PB_SocialParam_UnFollowUser, fn RPCClientHandler) (*PB_SocialResponse_UnFollowUser, error) {
	out := &PB_SocialResponse_UnFollowUser{}
	err := fn("RPC_Social.UnFollowUser", param, out)
	if err == nil {
		return out, nil
	}
	return nil, err
}

// service: RPC_User

func (RPC_User_Client) BlockUser(param *PB_UserParam_BlockUser, fn RPCClientHandler) (*PB_UserResponse_BlockUser, error) {
	out := &PB_UserResponse_BlockUser{}
	err := fn("RPC_User.BlockUser", param, out)
	if err == nil {
		return out, nil
	}
	return nil, err
}

func (RPC_User_Client) UnBlockUser(param *PB_UserParam_UnBlockUser, fn RPCClientHandler) (*PB_UserResponse_UnBlockUser, error) {
	out := &PB_UserResponse_UnBlockUser{}
	err := fn("RPC_User.UnBlockUser", param, out)
	if err == nil {
		return out, nil
	}
	return nil, err
}

func (RPC_User_Client) GetBlockedList(param *PB_UserParam_BlockedList, fn RPCClientHandler) (*PB_UserResponse_BlockedList, error) {
	out := &PB_UserResponse_BlockedList{}
	err := fn("RPC_User.GetBlockedList", param, out)
	if err == nil {
		return out, nil
	}
	return nil, err
}

func (RPC_User_Client) UpdateAbout(param *PB_UserParam_UpdateAbout, fn RPCClientHandler) (*PB_UserResponse_UpdateAbout, error) {
	out := &PB_UserResponse_UpdateAbout{}
	err := fn("RPC_User.UpdateAbout", param, out)
	if err == nil {
		return out, nil
	}
	return nil, err
}

func (RPC_User_Client) UpdateUserName(param *PB_UserParam_UpdateUserName, fn RPCClientHandler) (*PB_UserResponse_UpdateUserName, error) {
	out := &PB_UserResponse_UpdateUserName{}
	err := fn("RPC_User.UpdateUserName", param, out)
	if err == nil {
		return out, nil
	}
	return nil, err
}

func (RPC_User_Client) ChangePrivacy(param *PB_UserParam_ChangePrivacy, fn RPCClientHandler) (*PB_UserResponseOffline_ChangePrivacy, error) {
	out := &PB_UserResponseOffline_ChangePrivacy{}
	err := fn("RPC_User.ChangePrivacy", param, out)
	if err == nil {
		return out, nil
	}
	return nil, err
}

func (RPC_User_Client) ChangeAvatar(param *PB_UserParam_ChangeAvatar, fn RPCClientHandler) (*PB_UserResponse_ChangeAvatar, error) {
	out := &PB_UserResponse_ChangeAvatar{}
	err := fn("RPC_User.ChangeAvatar", param, out)
	if err == nil {
		return out, nil
	}
	return nil, err
}

func (RPC_User_Client) CheckUserName(param *PB_UserParam_CheckUserName, fn RPCClientHandler) (*PB_UserResponse_CheckUserName, error) {
	out := &PB_UserResponse_CheckUserName{}
	err := fn("RPC_User.CheckUserName", param, out)
	if err == nil {
		return out, nil
	}
	return nil, err
}
