package x

import (
	"errors"
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

func (c _StoreImpl) GetActionByActionId_JustCache(ActionId int) (*Action, bool) {
	o, ok := RowCache.Get("Action:" + strconv.Itoa(ActionId))
	if ok {
		if obj, ok := o.(*Action); ok {
			return obj, true
		}
	}

	if LogTableSqlReq.Action {
		XOLogErr(errors.New("_JustCache is empty for Action: " + strconv.Itoa(ActionId)))
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

func (c _StoreImpl) GetCommentByCommentId_JustCache(CommentId int) (*Comment, bool) {
	o, ok := RowCache.Get("Comment:" + strconv.Itoa(CommentId))
	if ok {
		if obj, ok := o.(*Comment); ok {
			return obj, true
		}
	}

	if LogTableSqlReq.Comment {
		XOLogErr(errors.New("_JustCache is empty for Comment: " + strconv.Itoa(CommentId)))
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

func (c _StoreImpl) GetCommentDeletedByUserId(UserId int) (*CommentDeleted, bool) {
	o, ok := RowCache.Get("CommentDeleted:" + strconv.Itoa(UserId))
	if ok {
		if obj, ok := o.(*CommentDeleted); ok {
			return obj, true
		}
	}
	obj2, err := CommentDeletedByUserId(base.DB, UserId)
	if err == nil {
		return obj2, true
	}
	if LogTableSqlReq.CommentDeleted {
		XOLogErr(err)
	}
	return nil, false
}

func (c _StoreImpl) GetCommentDeletedByUserId_JustCache(UserId int) (*CommentDeleted, bool) {
	o, ok := RowCache.Get("CommentDeleted:" + strconv.Itoa(UserId))
	if ok {
		if obj, ok := o.(*CommentDeleted); ok {
			return obj, true
		}
	}

	if LogTableSqlReq.CommentDeleted {
		XOLogErr(errors.New("_JustCache is empty for CommentDeleted: " + strconv.Itoa(UserId)))
	}
	return nil, false
}

func (c _StoreImpl) PreLoadCommentDeletedByUserIds(ids []int) {
	not_cached := make([]int, 0, len(ids))

	for _, id := range ids {
		_, ok := RowCache.Get("CommentDeleted:" + strconv.Itoa(id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		NewCommentDeleted_Selector().UserId_In(not_cached).GetRows(base.DB)
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

func (c _StoreImpl) GetEventByEventId_JustCache(EventId int) (*Event, bool) {
	o, ok := RowCache.Get("Event:" + strconv.Itoa(EventId))
	if ok {
		if obj, ok := o.(*Event); ok {
			return obj, true
		}
	}

	if LogTableSqlReq.Event {
		XOLogErr(errors.New("_JustCache is empty for Event: " + strconv.Itoa(EventId)))
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

func (c _StoreImpl) GetFollowingListById_JustCache(Id int) (*FollowingList, bool) {
	o, ok := RowCache.Get("FollowingList:" + strconv.Itoa(Id))
	if ok {
		if obj, ok := o.(*FollowingList); ok {
			return obj, true
		}
	}

	if LogTableSqlReq.FollowingList {
		XOLogErr(errors.New("_JustCache is empty for FollowingList: " + strconv.Itoa(Id)))
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

func (c _StoreImpl) GetFollowingListMemberById_JustCache(Id int) (*FollowingListMember, bool) {
	o, ok := RowCache.Get("FollowingListMember:" + strconv.Itoa(Id))
	if ok {
		if obj, ok := o.(*FollowingListMember); ok {
			return obj, true
		}
	}

	if LogTableSqlReq.FollowingListMember {
		XOLogErr(errors.New("_JustCache is empty for FollowingListMember: " + strconv.Itoa(Id)))
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

func (c _StoreImpl) GetFollowingListMemberRemovedById_JustCache(Id int) (*FollowingListMemberRemoved, bool) {
	o, ok := RowCache.Get("FollowingListMemberRemoved:" + strconv.Itoa(Id))
	if ok {
		if obj, ok := o.(*FollowingListMemberRemoved); ok {
			return obj, true
		}
	}

	if LogTableSqlReq.FollowingListMemberRemoved {
		XOLogErr(errors.New("_JustCache is empty for FollowingListMemberRemoved: " + strconv.Itoa(Id)))
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

func (c _StoreImpl) GetGroupByGroupId_JustCache(GroupId int) (*Group, bool) {
	o, ok := RowCache.Get("Group:" + strconv.Itoa(GroupId))
	if ok {
		if obj, ok := o.(*Group); ok {
			return obj, true
		}
	}

	if LogTableSqlReq.Group {
		XOLogErr(errors.New("_JustCache is empty for Group: " + strconv.Itoa(GroupId)))
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

func (c _StoreImpl) GetGroupMemberById_JustCache(Id int) (*GroupMember, bool) {
	o, ok := RowCache.Get("GroupMember:" + strconv.Itoa(Id))
	if ok {
		if obj, ok := o.(*GroupMember); ok {
			return obj, true
		}
	}

	if LogTableSqlReq.GroupMember {
		XOLogErr(errors.New("_JustCache is empty for GroupMember: " + strconv.Itoa(Id)))
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

func (c _StoreImpl) GetGroupMessageByMessageId_JustCache(MessageId int) (*GroupMessage, bool) {
	o, ok := RowCache.Get("GroupMessage:" + strconv.Itoa(MessageId))
	if ok {
		if obj, ok := o.(*GroupMessage); ok {
			return obj, true
		}
	}

	if LogTableSqlReq.GroupMessage {
		XOLogErr(errors.New("_JustCache is empty for GroupMessage: " + strconv.Itoa(MessageId)))
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

func (c _StoreImpl) GetLikeById_JustCache(Id int) (*Like, bool) {
	o, ok := RowCache.Get("Like:" + strconv.Itoa(Id))
	if ok {
		if obj, ok := o.(*Like); ok {
			return obj, true
		}
	}

	if LogTableSqlReq.Like {
		XOLogErr(errors.New("_JustCache is empty for Like: " + strconv.Itoa(Id)))
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

func (c _StoreImpl) GetNotifyByNotifyId_JustCache(NotifyId int) (*Notify, bool) {
	o, ok := RowCache.Get("Notify:" + strconv.Itoa(NotifyId))
	if ok {
		if obj, ok := o.(*Notify); ok {
			return obj, true
		}
	}

	if LogTableSqlReq.Notify {
		XOLogErr(errors.New("_JustCache is empty for Notify: " + strconv.Itoa(NotifyId)))
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

func (c _StoreImpl) GetNotifyRemovedById_JustCache(Id int) (*NotifyRemoved, bool) {
	o, ok := RowCache.Get("NotifyRemoved:" + strconv.Itoa(Id))
	if ok {
		if obj, ok := o.(*NotifyRemoved); ok {
			return obj, true
		}
	}

	if LogTableSqlReq.NotifyRemoved {
		XOLogErr(errors.New("_JustCache is empty for NotifyRemoved: " + strconv.Itoa(Id)))
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

func (c _StoreImpl) GetPhoneContactById_JustCache(Id int) (*PhoneContact, bool) {
	o, ok := RowCache.Get("PhoneContact:" + strconv.Itoa(Id))
	if ok {
		if obj, ok := o.(*PhoneContact); ok {
			return obj, true
		}
	}

	if LogTableSqlReq.PhoneContact {
		XOLogErr(errors.New("_JustCache is empty for PhoneContact: " + strconv.Itoa(Id)))
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

func (c _StoreImpl) GetPostByPostId_JustCache(PostId int) (*Post, bool) {
	o, ok := RowCache.Get("Post:" + strconv.Itoa(PostId))
	if ok {
		if obj, ok := o.(*Post); ok {
			return obj, true
		}
	}

	if LogTableSqlReq.Post {
		XOLogErr(errors.New("_JustCache is empty for Post: " + strconv.Itoa(PostId)))
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

func (c _StoreImpl) GetPostCountByPostId(PostId int) (*PostCount, bool) {
	o, ok := RowCache.Get("PostCount:" + strconv.Itoa(PostId))
	if ok {
		if obj, ok := o.(*PostCount); ok {
			return obj, true
		}
	}
	obj2, err := PostCountByPostId(base.DB, PostId)
	if err == nil {
		return obj2, true
	}
	if LogTableSqlReq.PostCount {
		XOLogErr(err)
	}
	return nil, false
}

func (c _StoreImpl) GetPostCountByPostId_JustCache(PostId int) (*PostCount, bool) {
	o, ok := RowCache.Get("PostCount:" + strconv.Itoa(PostId))
	if ok {
		if obj, ok := o.(*PostCount); ok {
			return obj, true
		}
	}

	if LogTableSqlReq.PostCount {
		XOLogErr(errors.New("_JustCache is empty for PostCount: " + strconv.Itoa(PostId)))
	}
	return nil, false
}

func (c _StoreImpl) PreLoadPostCountByPostIds(ids []int) {
	not_cached := make([]int, 0, len(ids))

	for _, id := range ids {
		_, ok := RowCache.Get("PostCount:" + strconv.Itoa(id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		NewPostCount_Selector().PostId_In(not_cached).GetRows(base.DB)
	}
}

// yes 222 int

func (c _StoreImpl) GetPostDeletedByPostId(PostId int) (*PostDeleted, bool) {
	o, ok := RowCache.Get("PostDeleted:" + strconv.Itoa(PostId))
	if ok {
		if obj, ok := o.(*PostDeleted); ok {
			return obj, true
		}
	}
	obj2, err := PostDeletedByPostId(base.DB, PostId)
	if err == nil {
		return obj2, true
	}
	if LogTableSqlReq.PostDeleted {
		XOLogErr(err)
	}
	return nil, false
}

func (c _StoreImpl) GetPostDeletedByPostId_JustCache(PostId int) (*PostDeleted, bool) {
	o, ok := RowCache.Get("PostDeleted:" + strconv.Itoa(PostId))
	if ok {
		if obj, ok := o.(*PostDeleted); ok {
			return obj, true
		}
	}

	if LogTableSqlReq.PostDeleted {
		XOLogErr(errors.New("_JustCache is empty for PostDeleted: " + strconv.Itoa(PostId)))
	}
	return nil, false
}

func (c _StoreImpl) PreLoadPostDeletedByPostIds(ids []int) {
	not_cached := make([]int, 0, len(ids))

	for _, id := range ids {
		_, ok := RowCache.Get("PostDeleted:" + strconv.Itoa(id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		NewPostDeleted_Selector().PostId_In(not_cached).GetRows(base.DB)
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

func (c _StoreImpl) GetPostKeyById_JustCache(Id int) (*PostKey, bool) {
	o, ok := RowCache.Get("PostKey:" + strconv.Itoa(Id))
	if ok {
		if obj, ok := o.(*PostKey); ok {
			return obj, true
		}
	}

	if LogTableSqlReq.PostKey {
		XOLogErr(errors.New("_JustCache is empty for PostKey: " + strconv.Itoa(Id)))
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

func (c _StoreImpl) GetPostLinkByLinkId(LinkId int) (*PostLink, bool) {
	o, ok := RowCache.Get("PostLink:" + strconv.Itoa(LinkId))
	if ok {
		if obj, ok := o.(*PostLink); ok {
			return obj, true
		}
	}
	obj2, err := PostLinkByLinkId(base.DB, LinkId)
	if err == nil {
		return obj2, true
	}
	if LogTableSqlReq.PostLink {
		XOLogErr(err)
	}
	return nil, false
}

func (c _StoreImpl) GetPostLinkByLinkId_JustCache(LinkId int) (*PostLink, bool) {
	o, ok := RowCache.Get("PostLink:" + strconv.Itoa(LinkId))
	if ok {
		if obj, ok := o.(*PostLink); ok {
			return obj, true
		}
	}

	if LogTableSqlReq.PostLink {
		XOLogErr(errors.New("_JustCache is empty for PostLink: " + strconv.Itoa(LinkId)))
	}
	return nil, false
}

func (c _StoreImpl) PreLoadPostLinkByLinkIds(ids []int) {
	not_cached := make([]int, 0, len(ids))

	for _, id := range ids {
		_, ok := RowCache.Get("PostLink:" + strconv.Itoa(id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		NewPostLink_Selector().LinkId_In(not_cached).GetRows(base.DB)
	}
}

// yes 222 int

func (c _StoreImpl) GetPostMediaByMediaId(MediaId int) (*PostMedia, bool) {
	o, ok := RowCache.Get("PostMedia:" + strconv.Itoa(MediaId))
	if ok {
		if obj, ok := o.(*PostMedia); ok {
			return obj, true
		}
	}
	obj2, err := PostMediaByMediaId(base.DB, MediaId)
	if err == nil {
		return obj2, true
	}
	if LogTableSqlReq.PostMedia {
		XOLogErr(err)
	}
	return nil, false
}

func (c _StoreImpl) GetPostMediaByMediaId_JustCache(MediaId int) (*PostMedia, bool) {
	o, ok := RowCache.Get("PostMedia:" + strconv.Itoa(MediaId))
	if ok {
		if obj, ok := o.(*PostMedia); ok {
			return obj, true
		}
	}

	if LogTableSqlReq.PostMedia {
		XOLogErr(errors.New("_JustCache is empty for PostMedia: " + strconv.Itoa(MediaId)))
	}
	return nil, false
}

func (c _StoreImpl) PreLoadPostMediaByMediaIds(ids []int) {
	not_cached := make([]int, 0, len(ids))

	for _, id := range ids {
		_, ok := RowCache.Get("PostMedia:" + strconv.Itoa(id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		NewPostMedia_Selector().MediaId_In(not_cached).GetRows(base.DB)
	}
}

// yes 222 int

func (c _StoreImpl) GetPostMentionedByPostId(PostId int) (*PostMentioned, bool) {
	o, ok := RowCache.Get("PostMentioned:" + strconv.Itoa(PostId))
	if ok {
		if obj, ok := o.(*PostMentioned); ok {
			return obj, true
		}
	}
	obj2, err := PostMentionedByPostId(base.DB, PostId)
	if err == nil {
		return obj2, true
	}
	if LogTableSqlReq.PostMentioned {
		XOLogErr(err)
	}
	return nil, false
}

func (c _StoreImpl) GetPostMentionedByPostId_JustCache(PostId int) (*PostMentioned, bool) {
	o, ok := RowCache.Get("PostMentioned:" + strconv.Itoa(PostId))
	if ok {
		if obj, ok := o.(*PostMentioned); ok {
			return obj, true
		}
	}

	if LogTableSqlReq.PostMentioned {
		XOLogErr(errors.New("_JustCache is empty for PostMentioned: " + strconv.Itoa(PostId)))
	}
	return nil, false
}

func (c _StoreImpl) PreLoadPostMentionedByPostIds(ids []int) {
	not_cached := make([]int, 0, len(ids))

	for _, id := range ids {
		_, ok := RowCache.Get("PostMentioned:" + strconv.Itoa(id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		NewPostMentioned_Selector().PostId_In(not_cached).GetRows(base.DB)
	}
}

// yes 222 int

func (c _StoreImpl) GetPostResharedByResharedId(ResharedId int) (*PostReshared, bool) {
	o, ok := RowCache.Get("PostReshared:" + strconv.Itoa(ResharedId))
	if ok {
		if obj, ok := o.(*PostReshared); ok {
			return obj, true
		}
	}
	obj2, err := PostResharedByResharedId(base.DB, ResharedId)
	if err == nil {
		return obj2, true
	}
	if LogTableSqlReq.PostReshared {
		XOLogErr(err)
	}
	return nil, false
}

func (c _StoreImpl) GetPostResharedByResharedId_JustCache(ResharedId int) (*PostReshared, bool) {
	o, ok := RowCache.Get("PostReshared:" + strconv.Itoa(ResharedId))
	if ok {
		if obj, ok := o.(*PostReshared); ok {
			return obj, true
		}
	}

	if LogTableSqlReq.PostReshared {
		XOLogErr(errors.New("_JustCache is empty for PostReshared: " + strconv.Itoa(ResharedId)))
	}
	return nil, false
}

func (c _StoreImpl) PreLoadPostResharedByResharedIds(ids []int) {
	not_cached := make([]int, 0, len(ids))

	for _, id := range ids {
		_, ok := RowCache.Get("PostReshared:" + strconv.Itoa(id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		NewPostReshared_Selector().ResharedId_In(not_cached).GetRows(base.DB)
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

func (c _StoreImpl) GetSearchClickedById_JustCache(Id int) (*SearchClicked, bool) {
	o, ok := RowCache.Get("SearchClicked:" + strconv.Itoa(Id))
	if ok {
		if obj, ok := o.(*SearchClicked); ok {
			return obj, true
		}
	}

	if LogTableSqlReq.SearchClicked {
		XOLogErr(errors.New("_JustCache is empty for SearchClicked: " + strconv.Itoa(Id)))
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

func (c _StoreImpl) GetSessionBySessionUuid_JustCache(SessionUuid string) (*Session, bool) {
	o, ok := RowCache.Get("Session:" + SessionUuid)
	if ok {
		if obj, ok := o.(*Session); ok {
			return obj, true
		}
	}

	if LogTableSqlReq.Session {
		XOLogErr(errors.New("_JustCache is empty for Session: " + SessionUuid))
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

func (c _StoreImpl) GetSettingClientByUserId_JustCache(UserId int) (*SettingClient, bool) {
	o, ok := RowCache.Get("SettingClient:" + strconv.Itoa(UserId))
	if ok {
		if obj, ok := o.(*SettingClient); ok {
			return obj, true
		}
	}

	if LogTableSqlReq.SettingClient {
		XOLogErr(errors.New("_JustCache is empty for SettingClient: " + strconv.Itoa(UserId)))
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

func (c _StoreImpl) GetSettingNotificationByUserId_JustCache(UserId int) (*SettingNotification, bool) {
	o, ok := RowCache.Get("SettingNotification:" + strconv.Itoa(UserId))
	if ok {
		if obj, ok := o.(*SettingNotification); ok {
			return obj, true
		}
	}

	if LogTableSqlReq.SettingNotification {
		XOLogErr(errors.New("_JustCache is empty for SettingNotification: " + strconv.Itoa(UserId)))
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

func (c _StoreImpl) GetTagByTagId_JustCache(TagId int) (*Tag, bool) {
	o, ok := RowCache.Get("Tag:" + strconv.Itoa(TagId))
	if ok {
		if obj, ok := o.(*Tag); ok {
			return obj, true
		}
	}

	if LogTableSqlReq.Tag {
		XOLogErr(errors.New("_JustCache is empty for Tag: " + strconv.Itoa(TagId)))
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

func (c _StoreImpl) GetTagPostById(Id int) (*TagPost, bool) {
	o, ok := RowCache.Get("TagPost:" + strconv.Itoa(Id))
	if ok {
		if obj, ok := o.(*TagPost); ok {
			return obj, true
		}
	}
	obj2, err := TagPostById(base.DB, Id)
	if err == nil {
		return obj2, true
	}
	if LogTableSqlReq.TagPost {
		XOLogErr(err)
	}
	return nil, false
}

func (c _StoreImpl) GetTagPostById_JustCache(Id int) (*TagPost, bool) {
	o, ok := RowCache.Get("TagPost:" + strconv.Itoa(Id))
	if ok {
		if obj, ok := o.(*TagPost); ok {
			return obj, true
		}
	}

	if LogTableSqlReq.TagPost {
		XOLogErr(errors.New("_JustCache is empty for TagPost: " + strconv.Itoa(Id)))
	}
	return nil, false
}

func (c _StoreImpl) PreLoadTagPostByIds(ids []int) {
	not_cached := make([]int, 0, len(ids))

	for _, id := range ids {
		_, ok := RowCache.Get("TagPost:" + strconv.Itoa(id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		NewTagPost_Selector().Id_In(not_cached).GetRows(base.DB)
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

func (c _StoreImpl) GetTriggerLogById_JustCache(Id int) (*TriggerLog, bool) {
	o, ok := RowCache.Get("TriggerLog:" + strconv.Itoa(Id))
	if ok {
		if obj, ok := o.(*TriggerLog); ok {
			return obj, true
		}
	}

	if LogTableSqlReq.TriggerLog {
		XOLogErr(errors.New("_JustCache is empty for TriggerLog: " + strconv.Itoa(Id)))
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

func (c _StoreImpl) GetUserByUserId_JustCache(UserId int) (*User, bool) {
	o, ok := RowCache.Get("User:" + strconv.Itoa(UserId))
	if ok {
		if obj, ok := o.(*User); ok {
			return obj, true
		}
	}

	if LogTableSqlReq.User {
		XOLogErr(errors.New("_JustCache is empty for User: " + strconv.Itoa(UserId)))
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

func (c _StoreImpl) GetUserMetaInfoById_JustCache(Id int) (*UserMetaInfo, bool) {
	o, ok := RowCache.Get("UserMetaInfo:" + strconv.Itoa(Id))
	if ok {
		if obj, ok := o.(*UserMetaInfo); ok {
			return obj, true
		}
	}

	if LogTableSqlReq.UserMetaInfo {
		XOLogErr(errors.New("_JustCache is empty for UserMetaInfo: " + strconv.Itoa(Id)))
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

func (c _StoreImpl) GetUserPasswordByUserId_JustCache(UserId int) (*UserPassword, bool) {
	o, ok := RowCache.Get("UserPassword:" + strconv.Itoa(UserId))
	if ok {
		if obj, ok := o.(*UserPassword); ok {
			return obj, true
		}
	}

	if LogTableSqlReq.UserPassword {
		XOLogErr(errors.New("_JustCache is empty for UserPassword: " + strconv.Itoa(UserId)))
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

func (c _StoreImpl) GetChatByChatKey_JustCache(ChatKey string) (*Chat, bool) {
	o, ok := RowCache.Get("Chat:" + ChatKey)
	if ok {
		if obj, ok := o.(*Chat); ok {
			return obj, true
		}
	}

	if LogTableSqlReq.Chat {
		XOLogErr(errors.New("_JustCache is empty for Chat: " + ChatKey))
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

func (c _StoreImpl) GetChatLastMessageByChatKey_JustCache(ChatKey string) (*ChatLastMessage, bool) {
	o, ok := RowCache.Get("ChatLastMessage:" + ChatKey)
	if ok {
		if obj, ok := o.(*ChatLastMessage); ok {
			return obj, true
		}
	}

	if LogTableSqlReq.ChatLastMessage {
		XOLogErr(errors.New("_JustCache is empty for ChatLastMessage: " + ChatKey))
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

func (c _StoreImpl) GetChatSyncBySyncId_JustCache(SyncId int) (*ChatSync, bool) {
	o, ok := RowCache.Get("ChatSync:" + strconv.Itoa(SyncId))
	if ok {
		if obj, ok := o.(*ChatSync); ok {
			return obj, true
		}
	}

	if LogTableSqlReq.ChatSync {
		XOLogErr(errors.New("_JustCache is empty for ChatSync: " + strconv.Itoa(SyncId)))
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

func (c _StoreImpl) GetDirectMessageByMessageId_JustCache(MessageId int) (*DirectMessage, bool) {
	o, ok := RowCache.Get("DirectMessage:" + strconv.Itoa(MessageId))
	if ok {
		if obj, ok := o.(*DirectMessage); ok {
			return obj, true
		}
	}

	if LogTableSqlReq.DirectMessage {
		XOLogErr(errors.New("_JustCache is empty for DirectMessage: " + strconv.Itoa(MessageId)))
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

func (c _StoreImpl) GetHomeById_JustCache(Id int) (*Home, bool) {
	o, ok := RowCache.Get("Home:" + strconv.Itoa(Id))
	if ok {
		if obj, ok := o.(*Home); ok {
			return obj, true
		}
	}

	if LogTableSqlReq.Home {
		XOLogErr(errors.New("_JustCache is empty for Home: " + strconv.Itoa(Id)))
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

func (c _StoreImpl) GetMessageFileByMessageFileId_JustCache(MessageFileId int) (*MessageFile, bool) {
	o, ok := RowCache.Get("MessageFile:" + strconv.Itoa(MessageFileId))
	if ok {
		if obj, ok := o.(*MessageFile); ok {
			return obj, true
		}
	}

	if LogTableSqlReq.MessageFile {
		XOLogErr(errors.New("_JustCache is empty for MessageFile: " + strconv.Itoa(MessageFileId)))
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

func (c _StoreImpl) GetFileMsgById_JustCache(Id int) (*FileMsg, bool) {
	o, ok := RowCache.Get("FileMsg:" + strconv.Itoa(Id))
	if ok {
		if obj, ok := o.(*FileMsg); ok {
			return obj, true
		}
	}

	if LogTableSqlReq.FileMsg {
		XOLogErr(errors.New("_JustCache is empty for FileMsg: " + strconv.Itoa(Id)))
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

func (c _StoreImpl) GetFilePostById_JustCache(Id int) (*FilePost, bool) {
	o, ok := RowCache.Get("FilePost:" + strconv.Itoa(Id))
	if ok {
		if obj, ok := o.(*FilePost); ok {
			return obj, true
		}
	}

	if LogTableSqlReq.FilePost {
		XOLogErr(errors.New("_JustCache is empty for FilePost: " + strconv.Itoa(Id)))
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

func (c _StoreImpl) GetActionFanoutByOrderId(OrderId int) (*ActionFanout, bool) {
	o, ok := RowCache.Get("ActionFanout:" + strconv.Itoa(OrderId))
	if ok {
		if obj, ok := o.(*ActionFanout); ok {
			return obj, true
		}
	}
	obj2, err := ActionFanoutByOrderId(base.DB, OrderId)
	if err == nil {
		return obj2, true
	}
	if LogTableSqlReq.ActionFanout {
		XOLogErr(err)
	}
	return nil, false
}

func (c _StoreImpl) GetActionFanoutByOrderId_JustCache(OrderId int) (*ActionFanout, bool) {
	o, ok := RowCache.Get("ActionFanout:" + strconv.Itoa(OrderId))
	if ok {
		if obj, ok := o.(*ActionFanout); ok {
			return obj, true
		}
	}

	if LogTableSqlReq.ActionFanout {
		XOLogErr(errors.New("_JustCache is empty for ActionFanout: " + strconv.Itoa(OrderId)))
	}
	return nil, false
}

func (c _StoreImpl) PreLoadActionFanoutByOrderIds(ids []int) {
	not_cached := make([]int, 0, len(ids))

	for _, id := range ids {
		_, ok := RowCache.Get("ActionFanout:" + strconv.Itoa(id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		NewActionFanout_Selector().OrderId_In(not_cached).GetRows(base.DB)
	}
}

// yes 222 int

func (c _StoreImpl) GetHomeFanoutByOrderId(OrderId int) (*HomeFanout, bool) {
	o, ok := RowCache.Get("HomeFanout:" + strconv.Itoa(OrderId))
	if ok {
		if obj, ok := o.(*HomeFanout); ok {
			return obj, true
		}
	}
	obj2, err := HomeFanoutByOrderId(base.DB, OrderId)
	if err == nil {
		return obj2, true
	}
	if LogTableSqlReq.HomeFanout {
		XOLogErr(err)
	}
	return nil, false
}

func (c _StoreImpl) GetHomeFanoutByOrderId_JustCache(OrderId int) (*HomeFanout, bool) {
	o, ok := RowCache.Get("HomeFanout:" + strconv.Itoa(OrderId))
	if ok {
		if obj, ok := o.(*HomeFanout); ok {
			return obj, true
		}
	}

	if LogTableSqlReq.HomeFanout {
		XOLogErr(errors.New("_JustCache is empty for HomeFanout: " + strconv.Itoa(OrderId)))
	}
	return nil, false
}

func (c _StoreImpl) PreLoadHomeFanoutByOrderIds(ids []int) {
	not_cached := make([]int, 0, len(ids))

	for _, id := range ids {
		_, ok := RowCache.Get("HomeFanout:" + strconv.Itoa(id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		NewHomeFanout_Selector().OrderId_In(not_cached).GetRows(base.DB)
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

func (c _StoreImpl) GetSuggestedTopPostById_JustCache(Id int) (*SuggestedTopPost, bool) {
	o, ok := RowCache.Get("SuggestedTopPost:" + strconv.Itoa(Id))
	if ok {
		if obj, ok := o.(*SuggestedTopPost); ok {
			return obj, true
		}
	}

	if LogTableSqlReq.SuggestedTopPost {
		XOLogErr(errors.New("_JustCache is empty for SuggestedTopPost: " + strconv.Itoa(Id)))
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

func (c _StoreImpl) GetSuggestedUserById_JustCache(Id int) (*SuggestedUser, bool) {
	o, ok := RowCache.Get("SuggestedUser:" + strconv.Itoa(Id))
	if ok {
		if obj, ok := o.(*SuggestedUser); ok {
			return obj, true
		}
	}

	if LogTableSqlReq.SuggestedUser {
		XOLogErr(errors.New("_JustCache is empty for SuggestedUser: " + strconv.Itoa(Id)))
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
