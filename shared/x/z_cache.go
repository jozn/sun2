package x

import (
	"errors"
	"ms/sun/shared/base"
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

func (c _StoreImpl) GetBlockedById(Id int) (*Blocked, bool) {
	o, ok := RowCache.Get("Blocked:" + strconv.Itoa(Id))
	if ok {
		if obj, ok := o.(*Blocked); ok {
			return obj, true
		}
	}
	obj2, err := BlockedById(base.DB, Id)
	if err == nil {
		return obj2, true
	}
	if LogTableSqlReq.Blocked {
		XOLogErr(err)
	}
	return nil, false
}

func (c _StoreImpl) GetBlockedById_JustCache(Id int) (*Blocked, bool) {
	o, ok := RowCache.Get("Blocked:" + strconv.Itoa(Id))
	if ok {
		if obj, ok := o.(*Blocked); ok {
			return obj, true
		}
	}

	if LogTableSqlReq.Blocked {
		XOLogErr(errors.New("_JustCache is empty for Blocked: " + strconv.Itoa(Id)))
	}
	return nil, false
}

func (c _StoreImpl) PreLoadBlockedByIds(ids []int) {
	not_cached := make([]int, 0, len(ids))

	for _, id := range ids {
		_, ok := RowCache.Get("Blocked:" + strconv.Itoa(id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		NewBlocked_Selector().Id_In(not_cached).GetRows(base.DB)
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

func (c _StoreImpl) GetCommentDeletedByCommentId(CommentId int) (*CommentDeleted, bool) {
	o, ok := RowCache.Get("CommentDeleted:" + strconv.Itoa(CommentId))
	if ok {
		if obj, ok := o.(*CommentDeleted); ok {
			return obj, true
		}
	}
	obj2, err := CommentDeletedByCommentId(base.DB, CommentId)
	if err == nil {
		return obj2, true
	}
	if LogTableSqlReq.CommentDeleted {
		XOLogErr(err)
	}
	return nil, false
}

func (c _StoreImpl) GetCommentDeletedByCommentId_JustCache(CommentId int) (*CommentDeleted, bool) {
	o, ok := RowCache.Get("CommentDeleted:" + strconv.Itoa(CommentId))
	if ok {
		if obj, ok := o.(*CommentDeleted); ok {
			return obj, true
		}
	}

	if LogTableSqlReq.CommentDeleted {
		XOLogErr(errors.New("_JustCache is empty for CommentDeleted: " + strconv.Itoa(CommentId)))
	}
	return nil, false
}

func (c _StoreImpl) PreLoadCommentDeletedByCommentIds(ids []int) {
	not_cached := make([]int, 0, len(ids))

	for _, id := range ids {
		_, ok := RowCache.Get("CommentDeleted:" + strconv.Itoa(id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		NewCommentDeleted_Selector().CommentId_In(not_cached).GetRows(base.DB)
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

func (c _StoreImpl) GetFollowedById(Id int) (*Followed, bool) {
	o, ok := RowCache.Get("Followed:" + strconv.Itoa(Id))
	if ok {
		if obj, ok := o.(*Followed); ok {
			return obj, true
		}
	}
	obj2, err := FollowedById(base.DB, Id)
	if err == nil {
		return obj2, true
	}
	if LogTableSqlReq.Followed {
		XOLogErr(err)
	}
	return nil, false
}

func (c _StoreImpl) GetFollowedById_JustCache(Id int) (*Followed, bool) {
	o, ok := RowCache.Get("Followed:" + strconv.Itoa(Id))
	if ok {
		if obj, ok := o.(*Followed); ok {
			return obj, true
		}
	}

	if LogTableSqlReq.Followed {
		XOLogErr(errors.New("_JustCache is empty for Followed: " + strconv.Itoa(Id)))
	}
	return nil, false
}

func (c _StoreImpl) PreLoadFollowedByIds(ids []int) {
	not_cached := make([]int, 0, len(ids))

	for _, id := range ids {
		_, ok := RowCache.Get("Followed:" + strconv.Itoa(id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		NewFollowed_Selector().Id_In(not_cached).GetRows(base.DB)
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

func (c _StoreImpl) GetPostMentionedByMentionedId(MentionedId int) (*PostMentioned, bool) {
	o, ok := RowCache.Get("PostMentioned:" + strconv.Itoa(MentionedId))
	if ok {
		if obj, ok := o.(*PostMentioned); ok {
			return obj, true
		}
	}
	obj2, err := PostMentionedByMentionedId(base.DB, MentionedId)
	if err == nil {
		return obj2, true
	}
	if LogTableSqlReq.PostMentioned {
		XOLogErr(err)
	}
	return nil, false
}

func (c _StoreImpl) GetPostMentionedByMentionedId_JustCache(MentionedId int) (*PostMentioned, bool) {
	o, ok := RowCache.Get("PostMentioned:" + strconv.Itoa(MentionedId))
	if ok {
		if obj, ok := o.(*PostMentioned); ok {
			return obj, true
		}
	}

	if LogTableSqlReq.PostMentioned {
		XOLogErr(errors.New("_JustCache is empty for PostMentioned: " + strconv.Itoa(MentionedId)))
	}
	return nil, false
}

func (c _StoreImpl) PreLoadPostMentionedByMentionedIds(ids []int) {
	not_cached := make([]int, 0, len(ids))

	for _, id := range ids {
		_, ok := RowCache.Get("PostMentioned:" + strconv.Itoa(id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		NewPostMentioned_Selector().MentionedId_In(not_cached).GetRows(base.DB)
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
	if LogTableSqlReq.Session {
		XOLogErr(err)
	}
	return nil, false
}

func (c _StoreImpl) GetSessionById_JustCache(Id int) (*Session, bool) {
	o, ok := RowCache.Get("Session:" + strconv.Itoa(Id))
	if ok {
		if obj, ok := o.(*Session); ok {
			return obj, true
		}
	}

	if LogTableSqlReq.Session {
		XOLogErr(errors.New("_JustCache is empty for Session: " + strconv.Itoa(Id)))
	}
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

// yes 222 uint

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

func (c _StoreImpl) GetUserRelationByRelNanoId(RelNanoId int) (*UserRelation, bool) {
	o, ok := RowCache.Get("UserRelation:" + strconv.Itoa(RelNanoId))
	if ok {
		if obj, ok := o.(*UserRelation); ok {
			return obj, true
		}
	}
	obj2, err := UserRelationByRelNanoId(base.DB, RelNanoId)
	if err == nil {
		return obj2, true
	}
	if LogTableSqlReq.UserRelation {
		XOLogErr(err)
	}
	return nil, false
}

func (c _StoreImpl) GetUserRelationByRelNanoId_JustCache(RelNanoId int) (*UserRelation, bool) {
	o, ok := RowCache.Get("UserRelation:" + strconv.Itoa(RelNanoId))
	if ok {
		if obj, ok := o.(*UserRelation); ok {
			return obj, true
		}
	}

	if LogTableSqlReq.UserRelation {
		XOLogErr(errors.New("_JustCache is empty for UserRelation: " + strconv.Itoa(RelNanoId)))
	}
	return nil, false
}

func (c _StoreImpl) PreLoadUserRelationByRelNanoIds(ids []int) {
	not_cached := make([]int, 0, len(ids))

	for _, id := range ids {
		_, ok := RowCache.Get("UserRelation:" + strconv.Itoa(id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		NewUserRelation_Selector().RelNanoId_In(not_cached).GetRows(base.DB)
	}
}

// yes 222 int

func (c _StoreImpl) GetChatByChatId(ChatId int) (*Chat, bool) {
	o, ok := RowCache.Get("Chat:" + strconv.Itoa(ChatId))
	if ok {
		if obj, ok := o.(*Chat); ok {
			return obj, true
		}
	}
	obj2, err := ChatByChatId(base.DB, ChatId)
	if err == nil {
		return obj2, true
	}
	if LogTableSqlReq.Chat {
		XOLogErr(err)
	}
	return nil, false
}

func (c _StoreImpl) GetChatByChatId_JustCache(ChatId int) (*Chat, bool) {
	o, ok := RowCache.Get("Chat:" + strconv.Itoa(ChatId))
	if ok {
		if obj, ok := o.(*Chat); ok {
			return obj, true
		}
	}

	if LogTableSqlReq.Chat {
		XOLogErr(errors.New("_JustCache is empty for Chat: " + strconv.Itoa(ChatId)))
	}
	return nil, false
}

func (c _StoreImpl) PreLoadChatByChatIds(ids []int) {
	not_cached := make([]int, 0, len(ids))

	for _, id := range ids {
		_, ok := RowCache.Get("Chat:" + strconv.Itoa(id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		NewChat_Selector().ChatId_In(not_cached).GetRows(base.DB)
	}
}

// yes 222 int

func (c _StoreImpl) GetChatDeletedByRoomKey(RoomKey string) (*ChatDeleted, bool) {
	o, ok := RowCache.Get("ChatDeleted:" + RoomKey)
	if ok {
		if obj, ok := o.(*ChatDeleted); ok {
			return obj, true
		}
	}
	obj2, err := ChatDeletedByRoomKey(base.DB, RoomKey)
	if err == nil {
		return obj2, true
	}
	if LogTableSqlReq.ChatDeleted {
		XOLogErr(err)
	}
	return nil, false
}

func (c _StoreImpl) GetChatDeletedByRoomKey_JustCache(RoomKey string) (*ChatDeleted, bool) {
	o, ok := RowCache.Get("ChatDeleted:" + RoomKey)
	if ok {
		if obj, ok := o.(*ChatDeleted); ok {
			return obj, true
		}
	}

	if LogTableSqlReq.ChatDeleted {
		XOLogErr(errors.New("_JustCache is empty for ChatDeleted: " + RoomKey))
	}
	return nil, false
}

func (c _StoreImpl) PreLoadChatDeletedByRoomKeys(ids []string) {
	not_cached := make([]string, 0, len(ids))

	for _, id := range ids {
		_, ok := RowCache.Get("ChatDeleted:" + id)
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		NewChatDeleted_Selector().RoomKey_In(not_cached).GetRows(base.DB)
	}
}

// yes 222 string

func (c _StoreImpl) GetChatLastMessageByChatIdGroupId(ChatIdGroupId string) (*ChatLastMessage, bool) {
	o, ok := RowCache.Get("ChatLastMessage:" + ChatIdGroupId)
	if ok {
		if obj, ok := o.(*ChatLastMessage); ok {
			return obj, true
		}
	}
	obj2, err := ChatLastMessageByChatIdGroupId(base.DB, ChatIdGroupId)
	if err == nil {
		return obj2, true
	}
	if LogTableSqlReq.ChatLastMessage {
		XOLogErr(err)
	}
	return nil, false
}

func (c _StoreImpl) GetChatLastMessageByChatIdGroupId_JustCache(ChatIdGroupId string) (*ChatLastMessage, bool) {
	o, ok := RowCache.Get("ChatLastMessage:" + ChatIdGroupId)
	if ok {
		if obj, ok := o.(*ChatLastMessage); ok {
			return obj, true
		}
	}

	if LogTableSqlReq.ChatLastMessage {
		XOLogErr(errors.New("_JustCache is empty for ChatLastMessage: " + ChatIdGroupId))
	}
	return nil, false
}

func (c _StoreImpl) PreLoadChatLastMessageByChatIdGroupIds(ids []string) {
	not_cached := make([]string, 0, len(ids))

	for _, id := range ids {
		_, ok := RowCache.Get("ChatLastMessage:" + id)
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		NewChatLastMessage_Selector().ChatIdGroupId_In(not_cached).GetRows(base.DB)
	}
}

// yes 222 string

func (c _StoreImpl) GetChatUserVersionByVersionTimeNano(VersionTimeNano int) (*ChatUserVersion, bool) {
	o, ok := RowCache.Get("ChatUserVersion:" + strconv.Itoa(VersionTimeNano))
	if ok {
		if obj, ok := o.(*ChatUserVersion); ok {
			return obj, true
		}
	}
	obj2, err := ChatUserVersionByVersionTimeNano(base.DB, VersionTimeNano)
	if err == nil {
		return obj2, true
	}
	if LogTableSqlReq.ChatUserVersion {
		XOLogErr(err)
	}
	return nil, false
}

func (c _StoreImpl) GetChatUserVersionByVersionTimeNano_JustCache(VersionTimeNano int) (*ChatUserVersion, bool) {
	o, ok := RowCache.Get("ChatUserVersion:" + strconv.Itoa(VersionTimeNano))
	if ok {
		if obj, ok := o.(*ChatUserVersion); ok {
			return obj, true
		}
	}

	if LogTableSqlReq.ChatUserVersion {
		XOLogErr(errors.New("_JustCache is empty for ChatUserVersion: " + strconv.Itoa(VersionTimeNano)))
	}
	return nil, false
}

func (c _StoreImpl) PreLoadChatUserVersionByVersionTimeNanos(ids []int) {
	not_cached := make([]int, 0, len(ids))

	for _, id := range ids {
		_, ok := RowCache.Get("ChatUserVersion:" + strconv.Itoa(id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		NewChatUserVersion_Selector().VersionTimeNano_In(not_cached).GetRows(base.DB)
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

func (c _StoreImpl) GetGroupMemberByOrderId(OrderId int) (*GroupMember, bool) {
	o, ok := RowCache.Get("GroupMember:" + strconv.Itoa(OrderId))
	if ok {
		if obj, ok := o.(*GroupMember); ok {
			return obj, true
		}
	}
	obj2, err := GroupMemberByOrderId(base.DB, OrderId)
	if err == nil {
		return obj2, true
	}
	if LogTableSqlReq.GroupMember {
		XOLogErr(err)
	}
	return nil, false
}

func (c _StoreImpl) GetGroupMemberByOrderId_JustCache(OrderId int) (*GroupMember, bool) {
	o, ok := RowCache.Get("GroupMember:" + strconv.Itoa(OrderId))
	if ok {
		if obj, ok := o.(*GroupMember); ok {
			return obj, true
		}
	}

	if LogTableSqlReq.GroupMember {
		XOLogErr(errors.New("_JustCache is empty for GroupMember: " + strconv.Itoa(OrderId)))
	}
	return nil, false
}

func (c _StoreImpl) PreLoadGroupMemberByOrderIds(ids []int) {
	not_cached := make([]int, 0, len(ids))

	for _, id := range ids {
		_, ok := RowCache.Get("GroupMember:" + strconv.Itoa(id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		NewGroupMember_Selector().OrderId_In(not_cached).GetRows(base.DB)
	}
}

// yes 222 int

func (c _StoreImpl) GetGroupOrderdUserByOrderId(OrderId int) (*GroupOrderdUser, bool) {
	o, ok := RowCache.Get("GroupOrderdUser:" + strconv.Itoa(OrderId))
	if ok {
		if obj, ok := o.(*GroupOrderdUser); ok {
			return obj, true
		}
	}
	obj2, err := GroupOrderdUserByOrderId(base.DB, OrderId)
	if err == nil {
		return obj2, true
	}
	if LogTableSqlReq.GroupOrderdUser {
		XOLogErr(err)
	}
	return nil, false
}

func (c _StoreImpl) GetGroupOrderdUserByOrderId_JustCache(OrderId int) (*GroupOrderdUser, bool) {
	o, ok := RowCache.Get("GroupOrderdUser:" + strconv.Itoa(OrderId))
	if ok {
		if obj, ok := o.(*GroupOrderdUser); ok {
			return obj, true
		}
	}

	if LogTableSqlReq.GroupOrderdUser {
		XOLogErr(errors.New("_JustCache is empty for GroupOrderdUser: " + strconv.Itoa(OrderId)))
	}
	return nil, false
}

func (c _StoreImpl) PreLoadGroupOrderdUserByOrderIds(ids []int) {
	not_cached := make([]int, 0, len(ids))

	for _, id := range ids {
		_, ok := RowCache.Get("GroupOrderdUser:" + strconv.Itoa(id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		NewGroupOrderdUser_Selector().OrderId_In(not_cached).GetRows(base.DB)
	}
}

// yes 222 int

func (c _StoreImpl) GetGroupPinedMsgByMsgId(MsgId int) (*GroupPinedMsg, bool) {
	o, ok := RowCache.Get("GroupPinedMsg:" + strconv.Itoa(MsgId))
	if ok {
		if obj, ok := o.(*GroupPinedMsg); ok {
			return obj, true
		}
	}
	obj2, err := GroupPinedMsgByMsgId(base.DB, MsgId)
	if err == nil {
		return obj2, true
	}
	if LogTableSqlReq.GroupPinedMsg {
		XOLogErr(err)
	}
	return nil, false
}

func (c _StoreImpl) GetGroupPinedMsgByMsgId_JustCache(MsgId int) (*GroupPinedMsg, bool) {
	o, ok := RowCache.Get("GroupPinedMsg:" + strconv.Itoa(MsgId))
	if ok {
		if obj, ok := o.(*GroupPinedMsg); ok {
			return obj, true
		}
	}

	if LogTableSqlReq.GroupPinedMsg {
		XOLogErr(errors.New("_JustCache is empty for GroupPinedMsg: " + strconv.Itoa(MsgId)))
	}
	return nil, false
}

func (c _StoreImpl) PreLoadGroupPinedMsgByMsgIds(ids []int) {
	not_cached := make([]int, 0, len(ids))

	for _, id := range ids {
		_, ok := RowCache.Get("GroupPinedMsg:" + strconv.Itoa(id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		NewGroupPinedMsg_Selector().MsgId_In(not_cached).GetRows(base.DB)
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

func (c _StoreImpl) GetPushChatByPushId(PushId int) (*PushChat, bool) {
	o, ok := RowCache.Get("PushChat:" + strconv.Itoa(PushId))
	if ok {
		if obj, ok := o.(*PushChat); ok {
			return obj, true
		}
	}
	obj2, err := PushChatByPushId(base.DB, PushId)
	if err == nil {
		return obj2, true
	}
	if LogTableSqlReq.PushChat {
		XOLogErr(err)
	}
	return nil, false
}

func (c _StoreImpl) GetPushChatByPushId_JustCache(PushId int) (*PushChat, bool) {
	o, ok := RowCache.Get("PushChat:" + strconv.Itoa(PushId))
	if ok {
		if obj, ok := o.(*PushChat); ok {
			return obj, true
		}
	}

	if LogTableSqlReq.PushChat {
		XOLogErr(errors.New("_JustCache is empty for PushChat: " + strconv.Itoa(PushId)))
	}
	return nil, false
}

func (c _StoreImpl) PreLoadPushChatByPushIds(ids []int) {
	not_cached := make([]int, 0, len(ids))

	for _, id := range ids {
		_, ok := RowCache.Get("PushChat:" + strconv.Itoa(id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		NewPushChat_Selector().PushId_In(not_cached).GetRows(base.DB)
	}
}

// yes 222 int

func (c _StoreImpl) GetHTTPRPCLogById(Id int) (*HTTPRPCLog, bool) {
	o, ok := RowCache.Get("HTTPRPCLog:" + strconv.Itoa(Id))
	if ok {
		if obj, ok := o.(*HTTPRPCLog); ok {
			return obj, true
		}
	}
	obj2, err := HTTPRPCLogById(base.DB, Id)
	if err == nil {
		return obj2, true
	}
	if LogTableSqlReq.HTTPRPCLog {
		XOLogErr(err)
	}
	return nil, false
}

func (c _StoreImpl) GetHTTPRPCLogById_JustCache(Id int) (*HTTPRPCLog, bool) {
	o, ok := RowCache.Get("HTTPRPCLog:" + strconv.Itoa(Id))
	if ok {
		if obj, ok := o.(*HTTPRPCLog); ok {
			return obj, true
		}
	}

	if LogTableSqlReq.HTTPRPCLog {
		XOLogErr(errors.New("_JustCache is empty for HTTPRPCLog: " + strconv.Itoa(Id)))
	}
	return nil, false
}

func (c _StoreImpl) PreLoadHTTPRPCLogByIds(ids []int) {
	not_cached := make([]int, 0, len(ids))

	for _, id := range ids {
		_, ok := RowCache.Get("HTTPRPCLog:" + strconv.Itoa(id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		NewHTTPRPCLog_Selector().Id_In(not_cached).GetRows(base.DB)
	}
}

// yes 222 int

func (c _StoreImpl) GetMetricLogById(Id int) (*MetricLog, bool) {
	o, ok := RowCache.Get("MetricLog:" + strconv.Itoa(Id))
	if ok {
		if obj, ok := o.(*MetricLog); ok {
			return obj, true
		}
	}
	obj2, err := MetricLogById(base.DB, Id)
	if err == nil {
		return obj2, true
	}
	if LogTableSqlReq.MetricLog {
		XOLogErr(err)
	}
	return nil, false
}

func (c _StoreImpl) GetMetricLogById_JustCache(Id int) (*MetricLog, bool) {
	o, ok := RowCache.Get("MetricLog:" + strconv.Itoa(Id))
	if ok {
		if obj, ok := o.(*MetricLog); ok {
			return obj, true
		}
	}

	if LogTableSqlReq.MetricLog {
		XOLogErr(errors.New("_JustCache is empty for MetricLog: " + strconv.Itoa(Id)))
	}
	return nil, false
}

func (c _StoreImpl) PreLoadMetricLogByIds(ids []int) {
	not_cached := make([]int, 0, len(ids))

	for _, id := range ids {
		_, ok := RowCache.Get("MetricLog:" + strconv.Itoa(id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		NewMetricLog_Selector().Id_In(not_cached).GetRows(base.DB)
	}
}

// yes 222 int

func (c _StoreImpl) GetXfileServiceInfoLogById(Id int) (*XfileServiceInfoLog, bool) {
	o, ok := RowCache.Get("XfileServiceInfoLog:" + strconv.Itoa(Id))
	if ok {
		if obj, ok := o.(*XfileServiceInfoLog); ok {
			return obj, true
		}
	}
	obj2, err := XfileServiceInfoLogById(base.DB, Id)
	if err == nil {
		return obj2, true
	}
	if LogTableSqlReq.XfileServiceInfoLog {
		XOLogErr(err)
	}
	return nil, false
}

func (c _StoreImpl) GetXfileServiceInfoLogById_JustCache(Id int) (*XfileServiceInfoLog, bool) {
	o, ok := RowCache.Get("XfileServiceInfoLog:" + strconv.Itoa(Id))
	if ok {
		if obj, ok := o.(*XfileServiceInfoLog); ok {
			return obj, true
		}
	}

	if LogTableSqlReq.XfileServiceInfoLog {
		XOLogErr(errors.New("_JustCache is empty for XfileServiceInfoLog: " + strconv.Itoa(Id)))
	}
	return nil, false
}

func (c _StoreImpl) PreLoadXfileServiceInfoLogByIds(ids []int) {
	not_cached := make([]int, 0, len(ids))

	for _, id := range ids {
		_, ok := RowCache.Get("XfileServiceInfoLog:" + strconv.Itoa(id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		NewXfileServiceInfoLog_Selector().Id_In(not_cached).GetRows(base.DB)
	}
}

// yes 222 int

func (c _StoreImpl) GetXfileServiceMetricLogById(Id int) (*XfileServiceMetricLog, bool) {
	o, ok := RowCache.Get("XfileServiceMetricLog:" + strconv.Itoa(Id))
	if ok {
		if obj, ok := o.(*XfileServiceMetricLog); ok {
			return obj, true
		}
	}
	obj2, err := XfileServiceMetricLogById(base.DB, Id)
	if err == nil {
		return obj2, true
	}
	if LogTableSqlReq.XfileServiceMetricLog {
		XOLogErr(err)
	}
	return nil, false
}

func (c _StoreImpl) GetXfileServiceMetricLogById_JustCache(Id int) (*XfileServiceMetricLog, bool) {
	o, ok := RowCache.Get("XfileServiceMetricLog:" + strconv.Itoa(Id))
	if ok {
		if obj, ok := o.(*XfileServiceMetricLog); ok {
			return obj, true
		}
	}

	if LogTableSqlReq.XfileServiceMetricLog {
		XOLogErr(errors.New("_JustCache is empty for XfileServiceMetricLog: " + strconv.Itoa(Id)))
	}
	return nil, false
}

func (c _StoreImpl) PreLoadXfileServiceMetricLogByIds(ids []int) {
	not_cached := make([]int, 0, len(ids))

	for _, id := range ids {
		_, ok := RowCache.Get("XfileServiceMetricLog:" + strconv.Itoa(id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		NewXfileServiceMetricLog_Selector().Id_In(not_cached).GetRows(base.DB)
	}
}

// yes 222 int

func (c _StoreImpl) GetXfileServiceRequestLogById(Id int) (*XfileServiceRequestLog, bool) {
	o, ok := RowCache.Get("XfileServiceRequestLog:" + strconv.Itoa(Id))
	if ok {
		if obj, ok := o.(*XfileServiceRequestLog); ok {
			return obj, true
		}
	}
	obj2, err := XfileServiceRequestLogById(base.DB, Id)
	if err == nil {
		return obj2, true
	}
	if LogTableSqlReq.XfileServiceRequestLog {
		XOLogErr(err)
	}
	return nil, false
}

func (c _StoreImpl) GetXfileServiceRequestLogById_JustCache(Id int) (*XfileServiceRequestLog, bool) {
	o, ok := RowCache.Get("XfileServiceRequestLog:" + strconv.Itoa(Id))
	if ok {
		if obj, ok := o.(*XfileServiceRequestLog); ok {
			return obj, true
		}
	}

	if LogTableSqlReq.XfileServiceRequestLog {
		XOLogErr(errors.New("_JustCache is empty for XfileServiceRequestLog: " + strconv.Itoa(Id)))
	}
	return nil, false
}

func (c _StoreImpl) PreLoadXfileServiceRequestLogByIds(ids []int) {
	not_cached := make([]int, 0, len(ids))

	for _, id := range ids {
		_, ok := RowCache.Get("XfileServiceRequestLog:" + strconv.Itoa(id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		NewXfileServiceRequestLog_Selector().Id_In(not_cached).GetRows(base.DB)
	}
}

// yes 222 int

func (c _StoreImpl) GetInvalidateCacheByOrderId(OrderId int) (*InvalidateCache, bool) {
	o, ok := RowCache.Get("InvalidateCache:" + strconv.Itoa(OrderId))
	if ok {
		if obj, ok := o.(*InvalidateCache); ok {
			return obj, true
		}
	}
	obj2, err := InvalidateCacheByOrderId(base.DB, OrderId)
	if err == nil {
		return obj2, true
	}
	if LogTableSqlReq.InvalidateCache {
		XOLogErr(err)
	}
	return nil, false
}

func (c _StoreImpl) GetInvalidateCacheByOrderId_JustCache(OrderId int) (*InvalidateCache, bool) {
	o, ok := RowCache.Get("InvalidateCache:" + strconv.Itoa(OrderId))
	if ok {
		if obj, ok := o.(*InvalidateCache); ok {
			return obj, true
		}
	}

	if LogTableSqlReq.InvalidateCache {
		XOLogErr(errors.New("_JustCache is empty for InvalidateCache: " + strconv.Itoa(OrderId)))
	}
	return nil, false
}

func (c _StoreImpl) PreLoadInvalidateCacheByOrderIds(ids []int) {
	not_cached := make([]int, 0, len(ids))

	for _, id := range ids {
		_, ok := RowCache.Get("InvalidateCache:" + strconv.Itoa(id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		NewInvalidateCache_Selector().OrderId_In(not_cached).GetRows(base.DB)
	}
}

// yes 222 int
