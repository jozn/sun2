package x

/////////////// Empty Sample RPC - mainly for mocking ////////////////

/////////////////// RPC_Auth  -  EmptyRPC_RPC_Auth /////////////////////
type EmptyRPC_RPC_Auth int

var Empty_RPC_RPC_Auth_Sample = EmptyRPC_RPC_Auth(0)

func (EmptyRPC_RPC_Auth) SendConfirmCode(i *RPC_Auth_Types_SendConfirmCode_Param, p RPC_UserParam) (*RPC_Auth_Types_SendConfirmCode_Response, error) {
	return nil, nil
}
func (EmptyRPC_RPC_Auth) ConfirmCode(i *RPC_Auth_Types_ConfirmCode_Param, p RPC_UserParam) (*RPC_Auth_Types_ConfirmCode_Response, error) {
	return nil, nil
}
func (EmptyRPC_RPC_Auth) SingUp(i *RPC_Auth_Types_SingUp_Param, p RPC_UserParam) (*RPC_Auth_Types_SingUp_Response, error) {
	return nil, nil
}
func (EmptyRPC_RPC_Auth) SingIn(i *RPC_Auth_Types_SingIn_Param, p RPC_UserParam) (*RPC_Auth_Types_SingIn_Response, error) {
	return nil, nil
}
func (EmptyRPC_RPC_Auth) LogOut(i *RPC_Auth_Types_LogOut_Param, p RPC_UserParam) (*RPC_Auth_Types_LogOut_Response, error) {
	return nil, nil
}

/////////////////// RPC_Chat  -  EmptyRPC_RPC_Chat /////////////////////
type EmptyRPC_RPC_Chat int

var Empty_RPC_RPC_Chat_Sample = EmptyRPC_RPC_Chat(0)

func (EmptyRPC_RPC_Chat) AddNewMessage(i *PB_RPC_Chat_Types_AddNewMessage_Param, p RPC_UserParam) (*PB_RPC_Chat_Types_AddNewMessage_Response, error) {
	return nil, nil
}
func (EmptyRPC_RPC_Chat) SetRoomActionDoing(i *PB_RPC_Chat_Types_SetRoomActionDoing_Param, p RPC_UserParam) (*PB_RPC_Chat_Types_SetRoomActionDoing_Response, error) {
	return nil, nil
}
func (EmptyRPC_RPC_Chat) GetChatList(i *PB_RPC_Chat_Types_GetChatList_Param, p RPC_UserParam) (*PB_RPC_Chat_Types_GetChatList_Response, error) {
	return nil, nil
}
func (EmptyRPC_RPC_Chat) GetChatHistory(i *PB_RPC_Chat_Types_GetChatHistory_Param, p RPC_UserParam) (*PB_RPC_Chat_Types_GetChatHistory_Response, error) {
	return nil, nil
}
func (EmptyRPC_RPC_Chat) PushRoomsChange(i *PB_RPC_Chat_Types_PushRoomsChange_Param, p RPC_UserParam) (*PB_RPC_Chat_Types_PushRoomsChange_Response, error) {
	return nil, nil
}
func (EmptyRPC_RPC_Chat) GetRoomsChange(i *PB_RPC_Chat_Types_GetRoomsChange_Param, p RPC_UserParam) (*PB_RPC_Chat_Types_GetRoomsChange_Response, error) {
	return nil, nil
}

/////////////////// RPC_General  -  EmptyRPC_RPC_General /////////////////////
type EmptyRPC_RPC_General int

var Empty_RPC_RPC_General_Sample = EmptyRPC_RPC_General(0)

func (EmptyRPC_RPC_General) Echo(i *RPC_General_Types_Echo_Param, p RPC_UserParam) (*RPC_General_Types_Echo_Response, error) {
	return nil, nil
}
func (EmptyRPC_RPC_General) CheckUserName(i *RPC_General_Types_CheckUserName_Param, p RPC_UserParam) (*RPC_General_Types_CheckUserName_Response, error) {
	return nil, nil
}

/////////////////// RPC_Page  -  EmptyRPC_RPC_Page /////////////////////
type EmptyRPC_RPC_Page int

var Empty_RPC_RPC_Page_Sample = EmptyRPC_RPC_Page(0)

func (EmptyRPC_RPC_Page) GetCommentsPage(i *RPC_Page_Types_GetCommentsPage_Param, p RPC_UserParam) (*RPC_Page_Types_GetCommentsPage_Response, error) {
	return nil, nil
}
func (EmptyRPC_RPC_Page) GetHomePage(i *RPC_Page_Types_GetHomePage_Param, p RPC_UserParam) (*RPC_Page_Types_GetHomePage_Response, error) {
	return nil, nil
}
func (EmptyRPC_RPC_Page) GetProfileAbout(i *RPC_Page_Types_GetProfileAbout_Param, p RPC_UserParam) (*RPC_Page_Types_GetProfileAbout_Response, error) {
	return nil, nil
}
func (EmptyRPC_RPC_Page) GetProfileAllShared(i *RPC_Page_Types_GetProfileAllShared_Param, p RPC_UserParam) (*RPC_Page_Types_GetProfileAllShared_Response, error) {
	return nil, nil
}
func (EmptyRPC_RPC_Page) GetProfileByCategoryPage(i *RPC_Page_Types_GetProfileByCategoryPage_Param, p RPC_UserParam) (*RPC_Page_Types_GetProfileByCategoryPage_Response, error) {
	return nil, nil
}
func (EmptyRPC_RPC_Page) GetLikesPage(i *RPC_Page_Types_GetLikesPage_Param, p RPC_UserParam) (*RPC_Page_Types_GetLikesPage_Response, error) {
	return nil, nil
}
func (EmptyRPC_RPC_Page) GetFollowersPage(i *RPC_Page_Types_GetFollowersPage_Param, p RPC_UserParam) (*RPC_Page_Types_GetFollowersPage_Response, error) {
	return nil, nil
}
func (EmptyRPC_RPC_Page) GetFollowingsPage(i *RPC_Page_Types_GetFollowingsPage_Param, p RPC_UserParam) (*RPC_Page_Types_GetFollowingsPage_Response, error) {
	return nil, nil
}
func (EmptyRPC_RPC_Page) GetNotifiesPage(i *RPC_Page_Types_GetNotifiesPage_Param, p RPC_UserParam) (*RPC_Page_Types_GetNotifiesPage_Response, error) {
	return nil, nil
}
func (EmptyRPC_RPC_Page) GetUserActionsPage(i *RPC_Page_Types_GetUserActionsPage_Param, p RPC_UserParam) (*RPC_Page_Types_GetUserActionsPage_Response, error) {
	return nil, nil
}
func (EmptyRPC_RPC_Page) GetPromotedPostsPage(i *RPC_Page_Types_GetPromotedPostsPage_Param, p RPC_UserParam) (*RPC_Page_Types_GetPromotedPostsPage_Response, error) {
	return nil, nil
}
func (EmptyRPC_RPC_Page) GetSuggestedUsersPage(i *RPC_Page_Types_GetSuggestedUsersPage_Param, p RPC_UserParam) (*RPC_Page_Types_GetSuggestedUsersPage_Response, error) {
	return nil, nil
}
func (EmptyRPC_RPC_Page) GetSuggestedTagsPage(i *RPC_Page_Types_GetSuggestedTagsPage_Param, p RPC_UserParam) (*RPC_Page_Types_GetSuggestedTagsPage_Response, error) {
	return nil, nil
}
func (EmptyRPC_RPC_Page) GetLastPostsPage(i *RPC_Page_Types_GetLastPostsPage_Param, p RPC_UserParam) (*RPC_Page_Types_GetLastPostsPage_Response, error) {
	return nil, nil
}
func (EmptyRPC_RPC_Page) GetLastTagPage(i *RPC_Page_Types_GetLastTagPage_Param, p RPC_UserParam) (*RPC_Page_Types_GetLastTagPage_Response, error) {
	return nil, nil
}
func (EmptyRPC_RPC_Page) SearchTagsPage(i *RPC_Page_Types_SearchTagsPage_Param, p RPC_UserParam) (*RPC_Page_Types_SearchTagsPage_Response, error) {
	return nil, nil
}
func (EmptyRPC_RPC_Page) SearchUsersPage(i *RPC_Page_Types_SearchUsersPage_Param, p RPC_UserParam) (*RPC_Page_Types_SearchUsersPage_Response, error) {
	return nil, nil
}

/////////////////// RPC_Social  -  EmptyRPC_RPC_Social /////////////////////
type EmptyRPC_RPC_Social int

var Empty_RPC_RPC_Social_Sample = EmptyRPC_RPC_Social(0)

func (EmptyRPC_RPC_Social) AddComment(i *RPC_Social_Types_AddComment_Param, p RPC_UserParam) (*RPC_Social_Types_AddComment_Response, error) {
	return nil, nil
}
func (EmptyRPC_RPC_Social) DeleteComment(i *RPC_Social_Types_DeleteComment_Param, p RPC_UserParam) (*RPC_Social_Types_DeleteComment_Response, error) {
	return nil, nil
}
func (EmptyRPC_RPC_Social) EditComment(i *RPC_Social_Types_EditComment_Param, p RPC_UserParam) (*RPC_Social_Types_EditComment_Response, error) {
	return nil, nil
}
func (EmptyRPC_RPC_Social) AddPost(i *RPC_Social_Types_AddPost_Param, p RPC_UserParam) (*RPC_Social_Types_AddPost_Response, error) {
	return nil, nil
}
func (EmptyRPC_RPC_Social) EditPost(i *RPC_Social_Types_EditPost_Param, p RPC_UserParam) (*RPC_Social_Types_EditPost_Response, error) {
	return nil, nil
}
func (EmptyRPC_RPC_Social) DeletePost(i *RPC_Social_Types_DeletePost_Param, p RPC_UserParam) (*RPC_Social_Types_DeletePost_Response, error) {
	return nil, nil
}
func (EmptyRPC_RPC_Social) ArchivePost(i *RPC_Social_Types_ArchivePost_Param, p RPC_UserParam) (*RPC_Social_Types_ArchivePost_Response, error) {
	return nil, nil
}
func (EmptyRPC_RPC_Social) LikePost(i *RPC_Social_Types_LikePost_Param, p RPC_UserParam) (*RPC_Social_Types_LikePost_Response, error) {
	return nil, nil
}
func (EmptyRPC_RPC_Social) UnLikePost(i *RPC_Social_Types_UnLikePost_Param, p RPC_UserParam) (*RPC_Social_Types_UnLikePost_Response, error) {
	return nil, nil
}
func (EmptyRPC_RPC_Social) FollowUser(i *RPC_Social_Types_FollowUser_Param, p RPC_UserParam) (*RPC_Social_Types_FollowUser_Response, error) {
	return nil, nil
}
func (EmptyRPC_RPC_Social) UnFollowUser(i *RPC_Social_Types_UnFollowUser_Param, p RPC_UserParam) (*RPC_Social_Types_UnFollowUser_Response, error) {
	return nil, nil
}
func (EmptyRPC_RPC_Social) BlockUser(i *RPC_Social_Types_BlockUser_Param, p RPC_UserParam) (*RPC_Social_Types_BlockUser_Response, error) {
	return nil, nil
}
func (EmptyRPC_RPC_Social) UnBlockUser(i *RPC_Social_Types_UnBlockUser_Param, p RPC_UserParam) (*RPC_Social_Types_UnBlockUser_Response, error) {
	return nil, nil
}
func (EmptyRPC_RPC_Social) AddSeenPosts(i *RPC_Social_Types_AddSeenPosts_Param, p RPC_UserParam) (*RPC_Social_Types_AddSeenPosts_Response, error) {
	return nil, nil
}

/////////////////// RPC_User  -  EmptyRPC_RPC_User /////////////////////
type EmptyRPC_RPC_User int

var Empty_RPC_RPC_User_Sample = EmptyRPC_RPC_User(0)

func (EmptyRPC_RPC_User) UpdateAbout(i *RPC_User_Types_UpdateAbout_Param, p RPC_UserParam) (*RPC_User_Types_UpdateAbout_Response, error) {
	return nil, nil
}
func (EmptyRPC_RPC_User) UpdateUserName(i *RPC_User_Types_UpdateUserName_Param, p RPC_UserParam) (*RPC_User_Types_UpdateUserName_Response, error) {
	return nil, nil
}
func (EmptyRPC_RPC_User) ChangePrivacy(i *RPC_User_Types_ChangePrivacy_Param, p RPC_UserParam) (*RPC_User_Types_ChangePrivacy_Response, error) {
	return nil, nil
}
func (EmptyRPC_RPC_User) ChangeAvatar(i *RPC_User_Types_ChangeAvatar_Param, p RPC_UserParam) (*RPC_User_Types_ChangeAvatar_Response, error) {
	return nil, nil
}
