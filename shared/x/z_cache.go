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
	if LogTableSqlReq.Action {
		XOLogErr(err)
	}
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
	if LogTableSqlReq.Comment {
		XOLogErr(err)
	}
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

func (c _StoreImpl) GetEventByEventId(EventId int) (*Event, bool) {
	o, ok := RowCache.Get("Event:" + strconv.Itoa(EventId))
	if ok {
		if obj, ok := o.(*Event); ok {
			return obj, true
		}
	}
	obj2, err := EventByEventId(base.DB, EventId)
	if err == nil {
		return obj2, true
	}
	if LogTableSqlReq.Event {
		XOLogErr(err)
	}
	return nil, false
}

func (c _StoreImpl) PreLoadEventByEventIds(ids []int) {
	not_cached := make([]int, 0, len(ids))

	for _, id := range ids {
		_, ok := RowCache.Get("Event:" + strconv.Itoa(id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		NewEvent_Selector().EventId_In(not_cached).GetRows(base.DB)
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
	if LogTableSqlReq.FollowingList {
		XOLogErr(err)
	}
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
	if LogTableSqlReq.FollowingListMember {
		XOLogErr(err)
	}
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

func (c _StoreImpl) GetFollowingListMemberRemovedById(Id int) (*FollowingListMemberRemoved, bool) {
	o, ok := RowCache.Get("FollowingListMemberRemoved:" + strconv.Itoa(Id))
	if ok {
		if obj, ok := o.(*FollowingListMemberRemoved); ok {
			return obj, true
		}
	}
	obj2, err := FollowingListMemberRemovedById(base.DB, Id)
	if err == nil {
		return obj2, true
	}
	if LogTableSqlReq.FollowingListMemberRemoved {
		XOLogErr(err)
	}
	return nil, false
}

func (c _StoreImpl) PreLoadFollowingListMemberRemovedByIds(ids []int) {
	not_cached := make([]int, 0, len(ids))

	for _, id := range ids {
		_, ok := RowCache.Get("FollowingListMemberRemoved:" + strconv.Itoa(id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		NewFollowingListMemberRemoved_Selector().Id_In(not_cached).GetRows(base.DB)
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
	if LogTableSqlReq.Group {
		XOLogErr(err)
	}
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
	if LogTableSqlReq.GroupMember {
		XOLogErr(err)
	}
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
	if LogTableSqlReq.GroupMessage {
		XOLogErr(err)
	}
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
	if LogTableSqlReq.Like {
		XOLogErr(err)
	}
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
	if LogTableSqlReq.Media {
		XOLogErr(err)
	}
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
	if LogTableSqlReq.Notify {
		XOLogErr(err)
	}
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
	if LogTableSqlReq.NotifyRemoved {
		XOLogErr(err)
	}
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
	if LogTableSqlReq.PhoneContact {
		XOLogErr(err)
	}
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
	if LogTableSqlReq.Post {
		XOLogErr(err)
	}
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

func (c _StoreImpl) GetPostKeyById(Id int) (*PostKey, bool) {
	o, ok := RowCache.Get("PostKey:" + strconv.Itoa(Id))
	if ok {
		if obj, ok := o.(*PostKey); ok {
			return obj, true
		}
	}
	obj2, err := PostKeyById(base.DB, Id)
	if err == nil {
		return obj2, true
	}
	if LogTableSqlReq.PostKey {
		XOLogErr(err)
	}
	return nil, false
}

func (c _StoreImpl) PreLoadPostKeyByIds(ids []int) {
	not_cached := make([]int, 0, len(ids))

	for _, id := range ids {
		_, ok := RowCache.Get("PostKey:" + strconv.Itoa(id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		NewPostKey_Selector().Id_In(not_cached).GetRows(base.DB)
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
	if LogTableSqlReq.SearchClicked {
		XOLogErr(err)
	}
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

func (c _StoreImpl) GetSessionBySessionUuid(SessionUuid string) (*Session, bool) {
	o, ok := RowCache.Get("Session:" + SessionUuid)
	if ok {
		if obj, ok := o.(*Session); ok {
			return obj, true
		}
	}
	obj2, err := SessionBySessionUuid(base.DB, SessionUuid)
	if err == nil {
		return obj2, true
	}
	if LogTableSqlReq.Session {
		XOLogErr(err)
	}
	return nil, false
}

func (c _StoreImpl) PreLoadSessionBySessionUuids(ids []string) {
	not_cached := make([]string, 0, len(ids))

	for _, id := range ids {
		_, ok := RowCache.Get("Session:" + id)
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		NewSession_Selector().SessionUuid_In(not_cached).GetRows(base.DB)
	}
}

// yes 222 string

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
	if LogTableSqlReq.SettingClient {
		XOLogErr(err)
	}
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
	if LogTableSqlReq.SettingNotification {
		XOLogErr(err)
	}
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

func (c _StoreImpl) GetSuggestedTopPostById(Id int) (*SuggestedTopPost, bool) {
	o, ok := RowCache.Get("SuggestedTopPost:" + strconv.Itoa(Id))
	if ok {
		if obj, ok := o.(*SuggestedTopPost); ok {
			return obj, true
		}
	}
	obj2, err := SuggestedTopPostById(base.DB, Id)
	if err == nil {
		return obj2, true
	}
	if LogTableSqlReq.SuggestedTopPost {
		XOLogErr(err)
	}
	return nil, false
}

func (c _StoreImpl) PreLoadSuggestedTopPostByIds(ids []int) {
	not_cached := make([]int, 0, len(ids))

	for _, id := range ids {
		_, ok := RowCache.Get("SuggestedTopPost:" + strconv.Itoa(id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		NewSuggestedTopPost_Selector().Id_In(not_cached).GetRows(base.DB)
	}
}

// yes 222 int

func (c _StoreImpl) GetSuggestedUserById(Id int) (*SuggestedUser, bool) {
	o, ok := RowCache.Get("SuggestedUser:" + strconv.Itoa(Id))
	if ok {
		if obj, ok := o.(*SuggestedUser); ok {
			return obj, true
		}
	}
	obj2, err := SuggestedUserById(base.DB, Id)
	if err == nil {
		return obj2, true
	}
	if LogTableSqlReq.SuggestedUser {
		XOLogErr(err)
	}
	return nil, false
}

func (c _StoreImpl) PreLoadSuggestedUserByIds(ids []int) {
	not_cached := make([]int, 0, len(ids))

	for _, id := range ids {
		_, ok := RowCache.Get("SuggestedUser:" + strconv.Itoa(id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		NewSuggestedUser_Selector().Id_In(not_cached).GetRows(base.DB)
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
	if LogTableSqlReq.Tag {
		XOLogErr(err)
	}
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
	if LogTableSqlReq.TagsPost {
		XOLogErr(err)
	}
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
	if LogTableSqlReq.TriggerLog {
		XOLogErr(err)
	}
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
	if LogTableSqlReq.User {
		XOLogErr(err)
	}
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
	if LogTableSqlReq.UserMetaInfo {
		XOLogErr(err)
	}
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
	if LogTableSqlReq.UserPassword {
		XOLogErr(err)
	}
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
	if LogTableSqlReq.Chat {
		XOLogErr(err)
	}
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

func (c _StoreImpl) GetChatLastMessageByChatKey(ChatKey string) (*ChatLastMessage, bool) {
	o, ok := RowCache.Get("ChatLastMessage:" + ChatKey)
	if ok {
		if obj, ok := o.(*ChatLastMessage); ok {
			return obj, true
		}
	}
	obj2, err := ChatLastMessageByChatKey(base.DB, ChatKey)
	if err == nil {
		return obj2, true
	}
	if LogTableSqlReq.ChatLastMessage {
		XOLogErr(err)
	}
	return nil, false
}

func (c _StoreImpl) PreLoadChatLastMessageByChatKeys(ids []string) {
	not_cached := make([]string, 0, len(ids))

	for _, id := range ids {
		_, ok := RowCache.Get("ChatLastMessage:" + id)
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		NewChatLastMessage_Selector().ChatKey_In(not_cached).GetRows(base.DB)
	}
}

// yes 222 string

func (c _StoreImpl) GetChatSyncBySyncId(SyncId int) (*ChatSync, bool) {
	o, ok := RowCache.Get("ChatSync:" + strconv.Itoa(SyncId))
	if ok {
		if obj, ok := o.(*ChatSync); ok {
			return obj, true
		}
	}
	obj2, err := ChatSyncBySyncId(base.DB, SyncId)
	if err == nil {
		return obj2, true
	}
	if LogTableSqlReq.ChatSync {
		XOLogErr(err)
	}
	return nil, false
}

func (c _StoreImpl) PreLoadChatSyncBySyncIds(ids []int) {
	not_cached := make([]int, 0, len(ids))

	for _, id := range ids {
		_, ok := RowCache.Get("ChatSync:" + strconv.Itoa(id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		NewChatSync_Selector().SyncId_In(not_cached).GetRows(base.DB)
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
	if LogTableSqlReq.DirectMessage {
		XOLogErr(err)
	}
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

func (c _StoreImpl) GetHomeById(Id int) (*Home, bool) {
	o, ok := RowCache.Get("Home:" + strconv.Itoa(Id))
	if ok {
		if obj, ok := o.(*Home); ok {
			return obj, true
		}
	}
	obj2, err := HomeById(base.DB, Id)
	if err == nil {
		return obj2, true
	}
	if LogTableSqlReq.Home {
		XOLogErr(err)
	}
	return nil, false
}

func (c _StoreImpl) PreLoadHomeByIds(ids []int) {
	not_cached := make([]int, 0, len(ids))

	for _, id := range ids {
		_, ok := RowCache.Get("Home:" + strconv.Itoa(id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		NewHome_Selector().Id_In(not_cached).GetRows(base.DB)
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
	if LogTableSqlReq.MessageFile {
		XOLogErr(err)
	}
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

func (c _StoreImpl) GetFileMsgById(Id int) (*FileMsg, bool) {
	o, ok := RowCache.Get("FileMsg:" + strconv.Itoa(Id))
	if ok {
		if obj, ok := o.(*FileMsg); ok {
			return obj, true
		}
	}
	obj2, err := FileMsgById(base.DB, Id)
	if err == nil {
		return obj2, true
	}
	if LogTableSqlReq.FileMsg {
		XOLogErr(err)
	}
	return nil, false
}

func (c _StoreImpl) PreLoadFileMsgByIds(ids []int) {
	not_cached := make([]int, 0, len(ids))

	for _, id := range ids {
		_, ok := RowCache.Get("FileMsg:" + strconv.Itoa(id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		NewFileMsg_Selector().Id_In(not_cached).GetRows(base.DB)
	}
}

// yes 222 int

func (c _StoreImpl) GetFilePostById(Id int) (*FilePost, bool) {
	o, ok := RowCache.Get("FilePost:" + strconv.Itoa(Id))
	if ok {
		if obj, ok := o.(*FilePost); ok {
			return obj, true
		}
	}
	obj2, err := FilePostById(base.DB, Id)
	if err == nil {
		return obj2, true
	}
	if LogTableSqlReq.FilePost {
		XOLogErr(err)
	}
	return nil, false
}

func (c _StoreImpl) PreLoadFilePostByIds(ids []int) {
	not_cached := make([]int, 0, len(ids))

	for _, id := range ids {
		_, ok := RowCache.Get("FilePost:" + strconv.Itoa(id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		NewFilePost_Selector().Id_In(not_cached).GetRows(base.DB)
	}
}

// yes 222 int
