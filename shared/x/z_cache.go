package x

import (
	"ms/sun/base"
	"strconv"
)

func (c _StoreImpl) GetActionByActionId(ActionId int) (*Action, bool) {
	o, ok := RowCache.Get("Action:" + strconv.Itoa(ActionId))
	if ok {
		if obj, ok := o.(*Action); ok {
			return obj, true
		}
	}
	obj2, err := ActionByActionId(base.DB, ActionId)
	if err == nil {
		return obj2, true
	}
	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadActionByActionIds(ids []int) {
	not_cached := make([]int, 0, len(ids))

	for _, id := range ids {
		_, ok := RowCache.Get("Action:" + strconv.Itoa(id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		NewAction_Selector().ActionId_In(not_cached).GetRows(base.DB)
	}
}

// yes 222 int

func (c _StoreImpl) GetChatByChatKey(ChatKey string) (*Chat, bool) {
	o, ok := RowCache.Get("Chat:" + ChatKey)
	if ok {
		if obj, ok := o.(*Chat); ok {
			return obj, true
		}
	}
	obj2, err := ChatByChatKey(base.DB, ChatKey)
	if err == nil {
		return obj2, true
	}
	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadChatByChatKeys(ids []string) {
	not_cached := make([]string, 0, len(ids))

	for _, id := range ids {
		_, ok := RowCache.Get("Chat:" + id)
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		NewChat_Selector().ChatKey_In(not_cached).GetRows(base.DB)
	}
}

// yes 222 string

func (c _StoreImpl) GetCommentByCommentId(CommentId int) (*Comment, bool) {
	o, ok := RowCache.Get("Comment:" + strconv.Itoa(CommentId))
	if ok {
		if obj, ok := o.(*Comment); ok {
			return obj, true
		}
	}
	obj2, err := CommentByCommentId(base.DB, CommentId)
	if err == nil {
		return obj2, true
	}
	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadCommentByCommentIds(ids []int) {
	not_cached := make([]int, 0, len(ids))

	for _, id := range ids {
		_, ok := RowCache.Get("Comment:" + strconv.Itoa(id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		NewComment_Selector().CommentId_In(not_cached).GetRows(base.DB)
	}
}

// yes 222 int

func (c _StoreImpl) GetDirectMessageByMessageId(MessageId int) (*DirectMessage, bool) {
	o, ok := RowCache.Get("DirectMessage:" + strconv.Itoa(MessageId))
	if ok {
		if obj, ok := o.(*DirectMessage); ok {
			return obj, true
		}
	}
	obj2, err := DirectMessageByMessageId(base.DB, MessageId)
	if err == nil {
		return obj2, true
	}
	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadDirectMessageByMessageIds(ids []int) {
	not_cached := make([]int, 0, len(ids))

	for _, id := range ids {
		_, ok := RowCache.Get("DirectMessage:" + strconv.Itoa(id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		NewDirectMessage_Selector().MessageId_In(not_cached).GetRows(base.DB)
	}
}

// yes 222 int

func (c _StoreImpl) GetDirectOfflineByDirectOfflineId(DirectOfflineId int) (*DirectOffline, bool) {
	o, ok := RowCache.Get("DirectOffline:" + strconv.Itoa(DirectOfflineId))
	if ok {
		if obj, ok := o.(*DirectOffline); ok {
			return obj, true
		}
	}
	obj2, err := DirectOfflineByDirectOfflineId(base.DB, DirectOfflineId)
	if err == nil {
		return obj2, true
	}
	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadDirectOfflineByDirectOfflineIds(ids []int) {
	not_cached := make([]int, 0, len(ids))

	for _, id := range ids {
		_, ok := RowCache.Get("DirectOffline:" + strconv.Itoa(id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		NewDirectOffline_Selector().DirectOfflineId_In(not_cached).GetRows(base.DB)
	}
}

// yes 222 int

func (c _StoreImpl) GetDirectToMessageById(Id int) (*DirectToMessage, bool) {
	o, ok := RowCache.Get("DirectToMessage:" + strconv.Itoa(Id))
	if ok {
		if obj, ok := o.(*DirectToMessage); ok {
			return obj, true
		}
	}
	obj2, err := DirectToMessageById(base.DB, Id)
	if err == nil {
		return obj2, true
	}
	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadDirectToMessageByIds(ids []int) {
	not_cached := make([]int, 0, len(ids))

	for _, id := range ids {
		_, ok := RowCache.Get("DirectToMessage:" + strconv.Itoa(id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		NewDirectToMessage_Selector().Id_In(not_cached).GetRows(base.DB)
	}
}

// yes 222 int

func (c _StoreImpl) GetFollowingListById(Id int) (*FollowingList, bool) {
	o, ok := RowCache.Get("FollowingList:" + strconv.Itoa(Id))
	if ok {
		if obj, ok := o.(*FollowingList); ok {
			return obj, true
		}
	}
	obj2, err := FollowingListById(base.DB, Id)
	if err == nil {
		return obj2, true
	}
	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadFollowingListByIds(ids []int) {
	not_cached := make([]int, 0, len(ids))

	for _, id := range ids {
		_, ok := RowCache.Get("FollowingList:" + strconv.Itoa(id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		NewFollowingList_Selector().Id_In(not_cached).GetRows(base.DB)
	}
}

// yes 222 int

func (c _StoreImpl) GetFollowingListMemberById(Id int) (*FollowingListMember, bool) {
	o, ok := RowCache.Get("FollowingListMember:" + strconv.Itoa(Id))
	if ok {
		if obj, ok := o.(*FollowingListMember); ok {
			return obj, true
		}
	}
	obj2, err := FollowingListMemberById(base.DB, Id)
	if err == nil {
		return obj2, true
	}
	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadFollowingListMemberByIds(ids []int) {
	not_cached := make([]int, 0, len(ids))

	for _, id := range ids {
		_, ok := RowCache.Get("FollowingListMember:" + strconv.Itoa(id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		NewFollowingListMember_Selector().Id_In(not_cached).GetRows(base.DB)
	}
}

// yes 222 int

func (c _StoreImpl) GetFollowingListMemberHistoryById(Id int) (*FollowingListMemberHistory, bool) {
	o, ok := RowCache.Get("FollowingListMemberHistory:" + strconv.Itoa(Id))
	if ok {
		if obj, ok := o.(*FollowingListMemberHistory); ok {
			return obj, true
		}
	}
	obj2, err := FollowingListMemberHistoryById(base.DB, Id)
	if err == nil {
		return obj2, true
	}
	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadFollowingListMemberHistoryByIds(ids []int) {
	not_cached := make([]int, 0, len(ids))

	for _, id := range ids {
		_, ok := RowCache.Get("FollowingListMemberHistory:" + strconv.Itoa(id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		NewFollowingListMemberHistory_Selector().Id_In(not_cached).GetRows(base.DB)
	}
}

// yes 222 int

func (c _StoreImpl) GetGeneralLogById(Id int) (*GeneralLog, bool) {
	o, ok := RowCache.Get("GeneralLog:" + strconv.Itoa(Id))
	if ok {
		if obj, ok := o.(*GeneralLog); ok {
			return obj, true
		}
	}
	obj2, err := GeneralLogById(base.DB, Id)
	if err == nil {
		return obj2, true
	}
	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadGeneralLogByIds(ids []int) {
	not_cached := make([]int, 0, len(ids))

	for _, id := range ids {
		_, ok := RowCache.Get("GeneralLog:" + strconv.Itoa(id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		NewGeneralLog_Selector().Id_In(not_cached).GetRows(base.DB)
	}
}

// yes 222 int

func (c _StoreImpl) GetGroupByGroupId(GroupId int) (*Group, bool) {
	o, ok := RowCache.Get("Group:" + strconv.Itoa(GroupId))
	if ok {
		if obj, ok := o.(*Group); ok {
			return obj, true
		}
	}
	obj2, err := GroupByGroupId(base.DB, GroupId)
	if err == nil {
		return obj2, true
	}
	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadGroupByGroupIds(ids []int) {
	not_cached := make([]int, 0, len(ids))

	for _, id := range ids {
		_, ok := RowCache.Get("Group:" + strconv.Itoa(id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		NewGroup_Selector().GroupId_In(not_cached).GetRows(base.DB)
	}
}

// yes 222 int

func (c _StoreImpl) GetGroupMemberById(Id int) (*GroupMember, bool) {
	o, ok := RowCache.Get("GroupMember:" + strconv.Itoa(Id))
	if ok {
		if obj, ok := o.(*GroupMember); ok {
			return obj, true
		}
	}
	obj2, err := GroupMemberById(base.DB, Id)
	if err == nil {
		return obj2, true
	}
	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadGroupMemberByIds(ids []int) {
	not_cached := make([]int, 0, len(ids))

	for _, id := range ids {
		_, ok := RowCache.Get("GroupMember:" + strconv.Itoa(id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		NewGroupMember_Selector().Id_In(not_cached).GetRows(base.DB)
	}
}

// yes 222 int

func (c _StoreImpl) GetGroupMessageByMessageId(MessageId int) (*GroupMessage, bool) {
	o, ok := RowCache.Get("GroupMessage:" + strconv.Itoa(MessageId))
	if ok {
		if obj, ok := o.(*GroupMessage); ok {
			return obj, true
		}
	}
	obj2, err := GroupMessageByMessageId(base.DB, MessageId)
	if err == nil {
		return obj2, true
	}
	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadGroupMessageByMessageIds(ids []int) {
	not_cached := make([]int, 0, len(ids))

	for _, id := range ids {
		_, ok := RowCache.Get("GroupMessage:" + strconv.Itoa(id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		NewGroupMessage_Selector().MessageId_In(not_cached).GetRows(base.DB)
	}
}

// yes 222 int

func (c _StoreImpl) GetLikeById(Id int) (*Like, bool) {
	o, ok := RowCache.Get("Like:" + strconv.Itoa(Id))
	if ok {
		if obj, ok := o.(*Like); ok {
			return obj, true
		}
	}
	obj2, err := LikeById(base.DB, Id)
	if err == nil {
		return obj2, true
	}
	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadLikeByIds(ids []int) {
	not_cached := make([]int, 0, len(ids))

	for _, id := range ids {
		_, ok := RowCache.Get("Like:" + strconv.Itoa(id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		NewLike_Selector().Id_In(not_cached).GetRows(base.DB)
	}
}

// yes 222 int

func (c _StoreImpl) GetMediaByMediaId(MediaId int) (*Media, bool) {
	o, ok := RowCache.Get("Media:" + strconv.Itoa(MediaId))
	if ok {
		if obj, ok := o.(*Media); ok {
			return obj, true
		}
	}
	obj2, err := MediaByMediaId(base.DB, MediaId)
	if err == nil {
		return obj2, true
	}
	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadMediaByMediaIds(ids []int) {
	not_cached := make([]int, 0, len(ids))

	for _, id := range ids {
		_, ok := RowCache.Get("Media:" + strconv.Itoa(id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		NewMedia_Selector().MediaId_In(not_cached).GetRows(base.DB)
	}
}

// yes 222 int

func (c _StoreImpl) GetMessageFileByMessageFileId(MessageFileId int) (*MessageFile, bool) {
	o, ok := RowCache.Get("MessageFile:" + strconv.Itoa(MessageFileId))
	if ok {
		if obj, ok := o.(*MessageFile); ok {
			return obj, true
		}
	}
	obj2, err := MessageFileByMessageFileId(base.DB, MessageFileId)
	if err == nil {
		return obj2, true
	}
	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadMessageFileByMessageFileIds(ids []int) {
	not_cached := make([]int, 0, len(ids))

	for _, id := range ids {
		_, ok := RowCache.Get("MessageFile:" + strconv.Itoa(id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		NewMessageFile_Selector().MessageFileId_In(not_cached).GetRows(base.DB)
	}
}

// yes 222 int

func (c _StoreImpl) GetNotifyByNotifyId(NotifyId int) (*Notify, bool) {
	o, ok := RowCache.Get("Notify:" + strconv.Itoa(NotifyId))
	if ok {
		if obj, ok := o.(*Notify); ok {
			return obj, true
		}
	}
	obj2, err := NotifyByNotifyId(base.DB, NotifyId)
	if err == nil {
		return obj2, true
	}
	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadNotifyByNotifyIds(ids []int) {
	not_cached := make([]int, 0, len(ids))

	for _, id := range ids {
		_, ok := RowCache.Get("Notify:" + strconv.Itoa(id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		NewNotify_Selector().NotifyId_In(not_cached).GetRows(base.DB)
	}
}

// yes 222 int

func (c _StoreImpl) GetNotifyRemovedById(Id int) (*NotifyRemoved, bool) {
	o, ok := RowCache.Get("NotifyRemoved:" + strconv.Itoa(Id))
	if ok {
		if obj, ok := o.(*NotifyRemoved); ok {
			return obj, true
		}
	}
	obj2, err := NotifyRemovedById(base.DB, Id)
	if err == nil {
		return obj2, true
	}
	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadNotifyRemovedByIds(ids []int) {
	not_cached := make([]int, 0, len(ids))

	for _, id := range ids {
		_, ok := RowCache.Get("NotifyRemoved:" + strconv.Itoa(id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		NewNotifyRemoved_Selector().Id_In(not_cached).GetRows(base.DB)
	}
}

// yes 222 int

func (c _StoreImpl) GetPhoneContactById(Id int) (*PhoneContact, bool) {
	o, ok := RowCache.Get("PhoneContact:" + strconv.Itoa(Id))
	if ok {
		if obj, ok := o.(*PhoneContact); ok {
			return obj, true
		}
	}
	obj2, err := PhoneContactById(base.DB, Id)
	if err == nil {
		return obj2, true
	}
	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadPhoneContactByIds(ids []int) {
	not_cached := make([]int, 0, len(ids))

	for _, id := range ids {
		_, ok := RowCache.Get("PhoneContact:" + strconv.Itoa(id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		NewPhoneContact_Selector().Id_In(not_cached).GetRows(base.DB)
	}
}

// yes 222 int

func (c _StoreImpl) GetPhoneContactsCopyById(Id int) (*PhoneContactsCopy, bool) {
	o, ok := RowCache.Get("PhoneContactsCopy:" + strconv.Itoa(Id))
	if ok {
		if obj, ok := o.(*PhoneContactsCopy); ok {
			return obj, true
		}
	}
	obj2, err := PhoneContactsCopyById(base.DB, Id)
	if err == nil {
		return obj2, true
	}
	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadPhoneContactsCopyByIds(ids []int) {
	not_cached := make([]int, 0, len(ids))

	for _, id := range ids {
		_, ok := RowCache.Get("PhoneContactsCopy:" + strconv.Itoa(id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		NewPhoneContactsCopy_Selector().Id_In(not_cached).GetRows(base.DB)
	}
}

// yes 222 int

func (c _StoreImpl) GetPostByPostId(PostId int) (*Post, bool) {
	o, ok := RowCache.Get("Post:" + strconv.Itoa(PostId))
	if ok {
		if obj, ok := o.(*Post); ok {
			return obj, true
		}
	}
	obj2, err := PostByPostId(base.DB, PostId)
	if err == nil {
		return obj2, true
	}
	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadPostByPostIds(ids []int) {
	not_cached := make([]int, 0, len(ids))

	for _, id := range ids {
		_, ok := RowCache.Get("Post:" + strconv.Itoa(id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		NewPost_Selector().PostId_In(not_cached).GetRows(base.DB)
	}
}

// yes 222 int

func (c _StoreImpl) GetRecommendUserById(Id int) (*RecommendUser, bool) {
	o, ok := RowCache.Get("RecommendUser:" + strconv.Itoa(Id))
	if ok {
		if obj, ok := o.(*RecommendUser); ok {
			return obj, true
		}
	}
	obj2, err := RecommendUserById(base.DB, Id)
	if err == nil {
		return obj2, true
	}
	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadRecommendUserByIds(ids []int) {
	not_cached := make([]int, 0, len(ids))

	for _, id := range ids {
		_, ok := RowCache.Get("RecommendUser:" + strconv.Itoa(id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		NewRecommendUser_Selector().Id_In(not_cached).GetRows(base.DB)
	}
}

// yes 222 int

func (c _StoreImpl) GetSearchClickedById(Id int) (*SearchClicked, bool) {
	o, ok := RowCache.Get("SearchClicked:" + strconv.Itoa(Id))
	if ok {
		if obj, ok := o.(*SearchClicked); ok {
			return obj, true
		}
	}
	obj2, err := SearchClickedById(base.DB, Id)
	if err == nil {
		return obj2, true
	}
	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadSearchClickedByIds(ids []int) {
	not_cached := make([]int, 0, len(ids))

	for _, id := range ids {
		_, ok := RowCache.Get("SearchClicked:" + strconv.Itoa(id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		NewSearchClicked_Selector().Id_In(not_cached).GetRows(base.DB)
	}
}

// yes 222 int

func (c _StoreImpl) GetSessionById(Id int) (*Session, bool) {
	o, ok := RowCache.Get("Session:" + strconv.Itoa(Id))
	if ok {
		if obj, ok := o.(*Session); ok {
			return obj, true
		}
	}
	obj2, err := SessionById(base.DB, Id)
	if err == nil {
		return obj2, true
	}
	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadSessionByIds(ids []int) {
	not_cached := make([]int, 0, len(ids))

	for _, id := range ids {
		_, ok := RowCache.Get("Session:" + strconv.Itoa(id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		NewSession_Selector().Id_In(not_cached).GetRows(base.DB)
	}
}

// yes 222 int

func (c _StoreImpl) GetSettingClientByUserId(UserId int) (*SettingClient, bool) {
	o, ok := RowCache.Get("SettingClient:" + strconv.Itoa(UserId))
	if ok {
		if obj, ok := o.(*SettingClient); ok {
			return obj, true
		}
	}
	obj2, err := SettingClientByUserId(base.DB, UserId)
	if err == nil {
		return obj2, true
	}
	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadSettingClientByUserIds(ids []int) {
	not_cached := make([]int, 0, len(ids))

	for _, id := range ids {
		_, ok := RowCache.Get("SettingClient:" + strconv.Itoa(id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		NewSettingClient_Selector().UserId_In(not_cached).GetRows(base.DB)
	}
}

// yes 222 int

func (c _StoreImpl) GetSettingNotificationByUserId(UserId int) (*SettingNotification, bool) {
	o, ok := RowCache.Get("SettingNotification:" + strconv.Itoa(UserId))
	if ok {
		if obj, ok := o.(*SettingNotification); ok {
			return obj, true
		}
	}
	obj2, err := SettingNotificationByUserId(base.DB, UserId)
	if err == nil {
		return obj2, true
	}
	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadSettingNotificationByUserIds(ids []int) {
	not_cached := make([]int, 0, len(ids))

	for _, id := range ids {
		_, ok := RowCache.Get("SettingNotification:" + strconv.Itoa(id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		NewSettingNotification_Selector().UserId_In(not_cached).GetRows(base.DB)
	}
}

// yes 222 int

func (c _StoreImpl) GetTagByTagId(TagId int) (*Tag, bool) {
	o, ok := RowCache.Get("Tag:" + strconv.Itoa(TagId))
	if ok {
		if obj, ok := o.(*Tag); ok {
			return obj, true
		}
	}
	obj2, err := TagByTagId(base.DB, TagId)
	if err == nil {
		return obj2, true
	}
	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadTagByTagIds(ids []int) {
	not_cached := make([]int, 0, len(ids))

	for _, id := range ids {
		_, ok := RowCache.Get("Tag:" + strconv.Itoa(id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		NewTag_Selector().TagId_In(not_cached).GetRows(base.DB)
	}
}

// yes 222 int

func (c _StoreImpl) GetTagsPostById(Id int) (*TagsPost, bool) {
	o, ok := RowCache.Get("TagsPost:" + strconv.Itoa(Id))
	if ok {
		if obj, ok := o.(*TagsPost); ok {
			return obj, true
		}
	}
	obj2, err := TagsPostById(base.DB, Id)
	if err == nil {
		return obj2, true
	}
	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadTagsPostByIds(ids []int) {
	not_cached := make([]int, 0, len(ids))

	for _, id := range ids {
		_, ok := RowCache.Get("TagsPost:" + strconv.Itoa(id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		NewTagsPost_Selector().Id_In(not_cached).GetRows(base.DB)
	}
}

// yes 222 int

func (c _StoreImpl) GetTriggerLogById(Id int) (*TriggerLog, bool) {
	o, ok := RowCache.Get("TriggerLog:" + strconv.Itoa(Id))
	if ok {
		if obj, ok := o.(*TriggerLog); ok {
			return obj, true
		}
	}
	obj2, err := TriggerLogById(base.DB, Id)
	if err == nil {
		return obj2, true
	}
	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadTriggerLogByIds(ids []int) {
	not_cached := make([]int, 0, len(ids))

	for _, id := range ids {
		_, ok := RowCache.Get("TriggerLog:" + strconv.Itoa(id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		NewTriggerLog_Selector().Id_In(not_cached).GetRows(base.DB)
	}
}

// yes 222 int

func (c _StoreImpl) GetUserByUserId(UserId int) (*User, bool) {
	o, ok := RowCache.Get("User:" + strconv.Itoa(UserId))
	if ok {
		if obj, ok := o.(*User); ok {
			return obj, true
		}
	}
	obj2, err := UserByUserId(base.DB, UserId)
	if err == nil {
		return obj2, true
	}
	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadUserByUserIds(ids []int) {
	not_cached := make([]int, 0, len(ids))

	for _, id := range ids {
		_, ok := RowCache.Get("User:" + strconv.Itoa(id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		NewUser_Selector().UserId_In(not_cached).GetRows(base.DB)
	}
}

// yes 222 int

func (c _StoreImpl) GetUserMetaInfoById(Id int) (*UserMetaInfo, bool) {
	o, ok := RowCache.Get("UserMetaInfo:" + strconv.Itoa(Id))
	if ok {
		if obj, ok := o.(*UserMetaInfo); ok {
			return obj, true
		}
	}
	obj2, err := UserMetaInfoById(base.DB, Id)
	if err == nil {
		return obj2, true
	}
	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadUserMetaInfoByIds(ids []int) {
	not_cached := make([]int, 0, len(ids))

	for _, id := range ids {
		_, ok := RowCache.Get("UserMetaInfo:" + strconv.Itoa(id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		NewUserMetaInfo_Selector().Id_In(not_cached).GetRows(base.DB)
	}
}

// yes 222 int

func (c _StoreImpl) GetUserPasswordByUserId(UserId int) (*UserPassword, bool) {
	o, ok := RowCache.Get("UserPassword:" + strconv.Itoa(UserId))
	if ok {
		if obj, ok := o.(*UserPassword); ok {
			return obj, true
		}
	}
	obj2, err := UserPasswordByUserId(base.DB, UserId)
	if err == nil {
		return obj2, true
	}
	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadUserPasswordByUserIds(ids []int) {
	not_cached := make([]int, 0, len(ids))

	for _, id := range ids {
		_, ok := RowCache.Get("UserPassword:" + strconv.Itoa(id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		NewUserPassword_Selector().UserId_In(not_cached).GetRows(base.DB)
	}
}

// yes 222 int
