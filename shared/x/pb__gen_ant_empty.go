package x

/////////////// Empty Sample RPC - mainly for mocking ////////////////

/////////////////// RPC_Auth  -  EmptyRPC_RPC_Auth /////////////////////
type EmptyRPC_RPC_Auth int

var Empty_RPC_RPC_Auth_Sample = EmptyRPC_RPC_Auth(0)

func (EmptyRPC_RPC_Auth) CheckPhone(i *PB_UserParam_CheckUserName2, p RPC_UserParam) (*PB_UserResponse_CheckUserName2, error) {
	return nil, nil
}
func (EmptyRPC_RPC_Auth) SendCode(i *PB_UserParam_CheckUserName2, p RPC_UserParam) (*PB_UserResponse_CheckUserName2, error) {
	return nil, nil
}
func (EmptyRPC_RPC_Auth) SendCodeToSms(i *PB_UserParam_CheckUserName2, p RPC_UserParam) (*PB_UserResponse_CheckUserName2, error) {
	return nil, nil
}
func (EmptyRPC_RPC_Auth) SendCodeToTelgram(i *PB_UserParam_CheckUserName2, p RPC_UserParam) (*PB_UserResponse_CheckUserName2, error) {
	return nil, nil
}
func (EmptyRPC_RPC_Auth) SingUp(i *PB_UserParam_CheckUserName2, p RPC_UserParam) (*PB_UserResponse_CheckUserName2, error) {
	return nil, nil
}
func (EmptyRPC_RPC_Auth) SingIn(i *PB_UserParam_CheckUserName2, p RPC_UserParam) (*PB_UserResponse_CheckUserName2, error) {
	return nil, nil
}
func (EmptyRPC_RPC_Auth) LogOut(i *PB_UserParam_CheckUserName2, p RPC_UserParam) (*PB_UserResponse_CheckUserName2, error) {
	return nil, nil
}

/////////////////// RPC_Chat  -  EmptyRPC_RPC_Chat /////////////////////
type EmptyRPC_RPC_Chat int

var Empty_RPC_RPC_Chat_Sample = EmptyRPC_RPC_Chat(0)

func (EmptyRPC_RPC_Chat) AddNewMessage(i *PB_ChatParam_AddNewMessage, p RPC_UserParam) (*PB_ChatResponse_AddNewMessage, error) {
	return nil, nil
}
func (EmptyRPC_RPC_Chat) SetRoomActionDoing(i *PB_ChatParam_SetRoomActionDoing, p RPC_UserParam) (*PB_ChatResponse_SetRoomActionDoing, error) {
	return nil, nil
}
func (EmptyRPC_RPC_Chat) SetMessagesRangeAsSeen(i *PB_ChatParam_SetChatMessagesRangeAsSeen, p RPC_UserParam) (*PB_ChatResponse_SetChatMessagesRangeAsSeen, error) {
	return nil, nil
}
func (EmptyRPC_RPC_Chat) DeleteChatHistory(i *PB_ChatParam_DeleteChatHistory, p RPC_UserParam) (*PB_ChatResponse_DeleteChatHistory, error) {
	return nil, nil
}
func (EmptyRPC_RPC_Chat) DeleteMessagesByIds(i *PB_ChatParam_DeleteMessagesByIds, p RPC_UserParam) (*PB_ChatResponse_DeleteMessagesByIds, error) {
	return nil, nil
}
func (EmptyRPC_RPC_Chat) SetMessagesAsReceived(i *PB_ChatParam_SetMessagesAsReceived, p RPC_UserParam) (*PB_ChatResponse_SetMessagesAsReceived, error) {
	return nil, nil
}
func (EmptyRPC_RPC_Chat) EditMessage(i *PB_ChatParam_EditMessage, p RPC_UserParam) (*PB_ChatResponse_EditMessage, error) {
	return nil, nil
}
func (EmptyRPC_RPC_Chat) GetChatList(i *PB_ChatParam_GetChatList, p RPC_UserParam) (*PB_ChatResponse_GetChatList, error) {
	return nil, nil
}
func (EmptyRPC_RPC_Chat) GetChatHistoryToOlder(i *PB_ChatParam_GetChatHistoryToOlder, p RPC_UserParam) (*PB_ChatResponse_GetChatHistoryToOlder, error) {
	return nil, nil
}
func (EmptyRPC_RPC_Chat) GetFreshAllDirectMessagesList(i *PB_ChatParam_GetFreshAllDirectMessagesList, p RPC_UserParam) (*PB_ChatResponse_GetFreshAllDirectMessagesList, error) {
	return nil, nil
}

/////////////////// RPC_Other  -  EmptyRPC_RPC_Other /////////////////////
type EmptyRPC_RPC_Other int

var Empty_RPC_RPC_Other_Sample = EmptyRPC_RPC_Other(0)

func (EmptyRPC_RPC_Other) Echo(i *PB_OtherParam_Echo, p RPC_UserParam) (*PB_OtherResponse_Echo, error) {
	return nil, nil
}

/////////////////// RPC_Page  -  EmptyRPC_RPC_Page /////////////////////
type EmptyRPC_RPC_Page int

var Empty_RPC_RPC_Page_Sample = EmptyRPC_RPC_Page(0)

func (EmptyRPC_RPC_Page) GetCommentsPage(i *PB_PageParam_GetCommentsPage, p RPC_UserParam) (*PB_PageResponse_GetCommentsPage, error) {
	return nil, nil
}
func (EmptyRPC_RPC_Page) GetHomePage(i *PB_PageParam_GetHomePage, p RPC_UserParam) (*PB_PageResponse_GetHomePage, error) {
	return nil, nil
}
func (EmptyRPC_RPC_Page) GetProfilePage(i *PB_PageParam_GetProfilePage, p RPC_UserParam) (*PB_PageResponse_GetProfilePage, error) {
	return nil, nil
}
func (EmptyRPC_RPC_Page) GetLikesPage(i *PB_PageParam_GetLikesPage, p RPC_UserParam) (*PB_PageResponse_GetLikesPage, error) {
	return nil, nil
}
func (EmptyRPC_RPC_Page) GetFollowersPage(i *PB_PageParam_GetFollowersPage, p RPC_UserParam) (*PB_PageResponse_GetFollowersPage, error) {
	return nil, nil
}
func (EmptyRPC_RPC_Page) GetFollowingsPage(i *PB_PageParam_GetFollowingsPage, p RPC_UserParam) (*PB_PageResponse_GetFollowingsPage, error) {
	return nil, nil
}
func (EmptyRPC_RPC_Page) GetNotifiesPage(i *PB_PageParam_GetNotifiesPage, p RPC_UserParam) (*PB_PageResponse_GetNotifiesPage, error) {
	return nil, nil
}
func (EmptyRPC_RPC_Page) GetUserActionsPage(i *PB_PageParam_GetUserActionsPage, p RPC_UserParam) (*PB_PageResponse_GetUserActionsPage, error) {
	return nil, nil
}
func (EmptyRPC_RPC_Page) GetSuggestedPostsPage(i *PB_PageParam_GetSuggestedPostsPage, p RPC_UserParam) (*PB_PageResponse_GetSuggestedPostsPage, error) {
	return nil, nil
}
func (EmptyRPC_RPC_Page) GetSuggestedUsersPage(i *PB_PageParam_GetSuggestedUsersPage, p RPC_UserParam) (*PB_PageResponse_GetSuggestedUsersPage, error) {
	return nil, nil
}
func (EmptyRPC_RPC_Page) GetSuggestedTagsPage(i *PB_PageParam_GetSuggestedTagsPage, p RPC_UserParam) (*PB_PageResponse_GetSuggestedTagsPage, error) {
	return nil, nil
}
func (EmptyRPC_RPC_Page) GetLastPostsPage(i *PB_PageParam_GetLastPostsPage, p RPC_UserParam) (*PB_PageResponse_GetLastPostsPage, error) {
	return nil, nil
}
func (EmptyRPC_RPC_Page) GetTagPage(i *PB_PageParam_GetTagPage, p RPC_UserParam) (*PB_PageResponse_GetTagPage, error) {
	return nil, nil
}
func (EmptyRPC_RPC_Page) SearchTagsPage(i *PB_PageParam_SearchTagsPage, p RPC_UserParam) (*PB_PageResponse_SearchTagsPage, error) {
	return nil, nil
}
func (EmptyRPC_RPC_Page) SearchUsersPage(i *PB_PageParam_SearchUsersPage, p RPC_UserParam) (*PB_PageResponse_SearchUsersPage, error) {
	return nil, nil
}

/////////////////// RPC_Search  -  EmptyRPC_RPC_Search /////////////////////
type EmptyRPC_RPC_Search int

var Empty_RPC_RPC_Search_Sample = EmptyRPC_RPC_Search(0)

func (EmptyRPC_RPC_Search) SearchTags(i *PB_SearchResponse_AddNewC, p RPC_UserParam) (*PB_SearchResponse_AddNewC, error) {
	return nil, nil
}

/////////////////// RPC_Social  -  EmptyRPC_RPC_Social /////////////////////
type EmptyRPC_RPC_Social int

var Empty_RPC_RPC_Social_Sample = EmptyRPC_RPC_Social(0)

func (EmptyRPC_RPC_Social) AddComment(i *PB_SocialParam_AddComment, p RPC_UserParam) (*PB_SocialResponse_AddComment, error) {
	return nil, nil
}
func (EmptyRPC_RPC_Social) DeleteComment(i *PB_SocialParam_DeleteComment, p RPC_UserParam) (*PB_SocialResponse_DeleteComment, error) {
	return nil, nil
}
func (EmptyRPC_RPC_Social) AddPost(i *PB_SocialParam_AddPost, p RPC_UserParam) (*PB_SocialResponse_AddPost, error) {
	return nil, nil
}
func (EmptyRPC_RPC_Social) EditPost(i *PB_SocialParam_EditPost, p RPC_UserParam) (*PB_SocialResponse_EditPost, error) {
	return nil, nil
}
func (EmptyRPC_RPC_Social) DeletePost(i *PB_SocialParam_DeletePost, p RPC_UserParam) (*PB_SocialResponse_DeletePost, error) {
	return nil, nil
}
func (EmptyRPC_RPC_Social) ArchivePost(i *PB_SocialParam_ArchivePost, p RPC_UserParam) (*PB_SocialResponse_ArchivePost, error) {
	return nil, nil
}
func (EmptyRPC_RPC_Social) LikePost(i *PB_SocialParam_LikePost, p RPC_UserParam) (*PB_SocialResponse_LikePost, error) {
	return nil, nil
}
func (EmptyRPC_RPC_Social) UnLikePost(i *PB_SocialParam_UnLikePost, p RPC_UserParam) (*PB_SocialResponse_UnLikePost, error) {
	return nil, nil
}
func (EmptyRPC_RPC_Social) FollowUser(i *PB_SocialParam_FollowUser, p RPC_UserParam) (*PB_SocialResponse_FollowUser, error) {
	return nil, nil
}
func (EmptyRPC_RPC_Social) UnFollowUser(i *PB_SocialParam_UnFollowUser, p RPC_UserParam) (*PB_SocialResponse_UnFollowUser, error) {
	return nil, nil
}

/////////////////// RPC_User  -  EmptyRPC_RPC_User /////////////////////
type EmptyRPC_RPC_User int

var Empty_RPC_RPC_User_Sample = EmptyRPC_RPC_User(0)

func (EmptyRPC_RPC_User) BlockUser(i *PB_UserParam_BlockUser, p RPC_UserParam) (*PB_UserResponse_BlockUser, error) {
	return nil, nil
}
func (EmptyRPC_RPC_User) UnBlockUser(i *PB_UserParam_UnBlockUser, p RPC_UserParam) (*PB_UserResponse_UnBlockUser, error) {
	return nil, nil
}
func (EmptyRPC_RPC_User) GetBlockedList(i *PB_UserParam_BlockedList, p RPC_UserParam) (*PB_UserResponse_BlockedList, error) {
	return nil, nil
}
func (EmptyRPC_RPC_User) UpdateAbout(i *PB_UserParam_UpdateAbout, p RPC_UserParam) (*PB_UserResponse_UpdateAbout, error) {
	return nil, nil
}
func (EmptyRPC_RPC_User) UpdateUserName(i *PB_UserParam_UpdateUserName, p RPC_UserParam) (*PB_UserResponse_UpdateUserName, error) {
	return nil, nil
}
func (EmptyRPC_RPC_User) ChangePrivacy(i *PB_UserParam_ChangePrivacy, p RPC_UserParam) (*PB_UserResponseOffline_ChangePrivacy, error) {
	return nil, nil
}
func (EmptyRPC_RPC_User) ChangeAvatar(i *PB_UserParam_ChangeAvatar, p RPC_UserParam) (*PB_UserResponse_ChangeAvatar, error) {
	return nil, nil
}
func (EmptyRPC_RPC_User) CheckUserName(i *PB_UserParam_CheckUserName, p RPC_UserParam) (*PB_UserResponse_CheckUserName, error) {
	return nil, nil
}
