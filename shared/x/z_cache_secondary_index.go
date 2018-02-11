package x

import (
	"fmt"
	"ms/sun/base"
)

//TODO: WE MUST separate int from string to not let empty string "" from preloading or loading and inserting into caches

// Action - PRIMARY

// Action - ActorUserId

//field//field//field

///// Generated from index 'Seq'.
func (c _StoreImpl) Action_BySeq(Seq int) (*Action, bool) {
	o, ok := RowCacheIndex.Get("Action_Seq:" + fmt.Sprintf("%v", Seq))
	if ok {
		if obj, ok := o.(*Action); ok {
			return obj, true
		}
	}

	row, err := NewAction_Selector().Seq_Eq(Seq).GetRow(base.DB)
	if err == nil {
		RowCacheIndex.Set("Action_Seq:"+fmt.Sprintf("%v", row.Seq), row, 0)
		return row, true
	}

	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadAction_BySeqs(Seqs []int) {
	not_cached := make([]int, 0, len(Seqs))

	for _, id := range Seqs {
		_, ok := RowCacheIndex.Get("Action_Seq:" + fmt.Sprintf("%v", id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		rows, err := NewAction_Selector().Seq_In(not_cached).GetRows(base.DB)
		if err == nil {
			for _, row := range rows {
				RowCacheIndex.Set("Action_Seq:"+fmt.Sprintf("%v", row.Seq), row, 0)
			}
		}
	}
}

// Chat - PRIMARY

// ChatSync - PRIMARY

// Comment - PRIMARY

//field//field//field

///// Generated from index 'PostId'.
func (c _StoreImpl) Comment_ByPostId(PostId int) (*Comment, bool) {
	o, ok := RowCacheIndex.Get("Comment_PostId:" + fmt.Sprintf("%v", PostId))
	if ok {
		if obj, ok := o.(*Comment); ok {
			return obj, true
		}
	}

	row, err := NewComment_Selector().PostId_Eq(PostId).GetRow(base.DB)
	if err == nil {
		RowCacheIndex.Set("Comment_PostId:"+fmt.Sprintf("%v", row.PostId), row, 0)
		return row, true
	}

	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadComment_ByPostIds(PostIds []int) {
	not_cached := make([]int, 0, len(PostIds))

	for _, id := range PostIds {
		_, ok := RowCacheIndex.Get("Comment_PostId:" + fmt.Sprintf("%v", id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		rows, err := NewComment_Selector().PostId_In(not_cached).GetRows(base.DB)
		if err == nil {
			for _, row := range rows {
				RowCacheIndex.Set("Comment_PostId:"+fmt.Sprintf("%v", row.PostId), row, 0)
			}
		}
	}
}

//field//field//field

///// Generated from index 'Asa'.
func (c _StoreImpl) Comment_BySeq(Seq int) (*Comment, bool) {
	o, ok := RowCacheIndex.Get("Comment_Asa:" + fmt.Sprintf("%v", Seq))
	if ok {
		if obj, ok := o.(*Comment); ok {
			return obj, true
		}
	}

	row, err := NewComment_Selector().Seq_Eq(Seq).GetRow(base.DB)
	if err == nil {
		RowCacheIndex.Set("Comment_Asa:"+fmt.Sprintf("%v", row.Seq), row, 0)
		return row, true
	}

	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadComment_BySeqs(Seqs []int) {
	not_cached := make([]int, 0, len(Seqs))

	for _, id := range Seqs {
		_, ok := RowCacheIndex.Get("Comment_Asa:" + fmt.Sprintf("%v", id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		rows, err := NewComment_Selector().Seq_In(not_cached).GetRows(base.DB)
		if err == nil {
			for _, row := range rows {
				RowCacheIndex.Set("Comment_Asa:"+fmt.Sprintf("%v", row.Seq), row, 0)
			}
		}
	}
}

// DirectMessage - PRIMARY

// DirectMessageCopy - PRIMARY

// DirectOffline - PRIMARY

//field//field//field

///// Generated from index 'ToUserId'.
func (c _StoreImpl) DirectOffline_ByToUserId(ToUserId int) (*DirectOffline, bool) {
	o, ok := RowCacheIndex.Get("DirectOffline_ToUserId:" + fmt.Sprintf("%v", ToUserId))
	if ok {
		if obj, ok := o.(*DirectOffline); ok {
			return obj, true
		}
	}

	row, err := NewDirectOffline_Selector().ToUserId_Eq(ToUserId).GetRow(base.DB)
	if err == nil {
		RowCacheIndex.Set("DirectOffline_ToUserId:"+fmt.Sprintf("%v", row.ToUserId), row, 0)
		return row, true
	}

	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadDirectOffline_ByToUserIds(ToUserIds []int) {
	not_cached := make([]int, 0, len(ToUserIds))

	for _, id := range ToUserIds {
		_, ok := RowCacheIndex.Get("DirectOffline_ToUserId:" + fmt.Sprintf("%v", id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		rows, err := NewDirectOffline_Selector().ToUserId_In(not_cached).GetRows(base.DB)
		if err == nil {
			for _, row := range rows {
				RowCacheIndex.Set("DirectOffline_ToUserId:"+fmt.Sprintf("%v", row.ToUserId), row, 0)
			}
		}
	}
}

// DirectToMessage - PRIMARY

// DirectToMessage - ChatKey

// FollowingList - PRIMARY

// FollowingListMember - PRIMARY

// FollowingListMember - UserId

// FollowingListMember - FollowedUserId

// FollowingListMember - UserId_2

// FollowingListMemberRemoved - PRIMARY

// FollowingListMemberRemoved - UserId

// GeneralLog - PRIMARY

// Group - PRIMARY

// GroupMember - PRIMARY

//field//field//field

///// Generated from index 'Id'.
func (c _StoreImpl) GroupMember_ById(Id int) (*GroupMember, bool) {
	o, ok := RowCacheIndex.Get("GroupMember_Id:" + fmt.Sprintf("%v", Id))
	if ok {
		if obj, ok := o.(*GroupMember); ok {
			return obj, true
		}
	}

	row, err := NewGroupMember_Selector().Id_Eq(Id).GetRow(base.DB)
	if err == nil {
		RowCacheIndex.Set("GroupMember_Id:"+fmt.Sprintf("%v", row.Id), row, 0)
		return row, true
	}

	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadGroupMember_ByIds(Ids []int) {
	not_cached := make([]int, 0, len(Ids))

	for _, id := range Ids {
		_, ok := RowCacheIndex.Get("GroupMember_Id:" + fmt.Sprintf("%v", id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		rows, err := NewGroupMember_Selector().Id_In(not_cached).GetRows(base.DB)
		if err == nil {
			for _, row := range rows {
				RowCacheIndex.Set("GroupMember_Id:"+fmt.Sprintf("%v", row.Id), row, 0)
			}
		}
	}
}

// GroupMessage - PRIMARY

// Key - PRIMARY

// Like - PRIMARY

// Like - PostId

//field//field//field

///// Generated from index 'Id'.
func (c _StoreImpl) Like_ById(Id int) (*Like, bool) {
	o, ok := RowCacheIndex.Get("Like_Id:" + fmt.Sprintf("%v", Id))
	if ok {
		if obj, ok := o.(*Like); ok {
			return obj, true
		}
	}

	row, err := NewLike_Selector().Id_Eq(Id).GetRow(base.DB)
	if err == nil {
		RowCacheIndex.Set("Like_Id:"+fmt.Sprintf("%v", row.Id), row, 0)
		return row, true
	}

	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadLike_ByIds(Ids []int) {
	not_cached := make([]int, 0, len(Ids))

	for _, id := range Ids {
		_, ok := RowCacheIndex.Get("Like_Id:" + fmt.Sprintf("%v", id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		rows, err := NewLike_Selector().Id_In(not_cached).GetRows(base.DB)
		if err == nil {
			for _, row := range rows {
				RowCacheIndex.Set("Like_Id:"+fmt.Sprintf("%v", row.Id), row, 0)
			}
		}
	}
}

//field//field//field

///// Generated from index 'PostId_2'.
func (c _StoreImpl) Like_ByPostId(PostId int) (*Like, bool) {
	o, ok := RowCacheIndex.Get("Like_PostId_2:" + fmt.Sprintf("%v", PostId))
	if ok {
		if obj, ok := o.(*Like); ok {
			return obj, true
		}
	}

	row, err := NewLike_Selector().PostId_Eq(PostId).GetRow(base.DB)
	if err == nil {
		RowCacheIndex.Set("Like_PostId_2:"+fmt.Sprintf("%v", row.PostId), row, 0)
		return row, true
	}

	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadLike_ByPostIds(PostIds []int) {
	not_cached := make([]int, 0, len(PostIds))

	for _, id := range PostIds {
		_, ok := RowCacheIndex.Get("Like_PostId_2:" + fmt.Sprintf("%v", id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		rows, err := NewLike_Selector().PostId_In(not_cached).GetRows(base.DB)
		if err == nil {
			for _, row := range rows {
				RowCacheIndex.Set("Like_PostId_2:"+fmt.Sprintf("%v", row.PostId), row, 0)
			}
		}
	}
}

// Media - PRIMARY

//field//field//field

///// Generated from index 'HashMd5'.
func (c _StoreImpl) Media_ByMd5Hash(Md5Hash string) (*Media, bool) {
	o, ok := RowCacheIndex.Get("Media_HashMd5:" + fmt.Sprintf("%v", Md5Hash))
	if ok {
		if obj, ok := o.(*Media); ok {
			return obj, true
		}
	}

	row, err := NewMedia_Selector().Md5Hash_Eq(Md5Hash).GetRow(base.DB)
	if err == nil {
		RowCacheIndex.Set("Media_HashMd5:"+fmt.Sprintf("%v", row.Md5Hash), row, 0)
		return row, true
	}

	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadMedia_ByMd5Hashs(Md5Hashs []string) {
	not_cached := make([]string, 0, len(Md5Hashs))

	for _, id := range Md5Hashs {
		_, ok := RowCacheIndex.Get("Media_HashMd5:" + fmt.Sprintf("%v", id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		rows, err := NewMedia_Selector().Md5Hash_In(not_cached).GetRows(base.DB)
		if err == nil {
			for _, row := range rows {
				RowCacheIndex.Set("Media_HashMd5:"+fmt.Sprintf("%v", row.Md5Hash), row, 0)
			}
		}
	}
}

//field//field//field

///// Generated from index 'CreatedTime'.
func (c _StoreImpl) Media_ByCreatedTime(CreatedTime int) (*Media, bool) {
	o, ok := RowCacheIndex.Get("Media_CreatedTime:" + fmt.Sprintf("%v", CreatedTime))
	if ok {
		if obj, ok := o.(*Media); ok {
			return obj, true
		}
	}

	row, err := NewMedia_Selector().CreatedTime_Eq(CreatedTime).GetRow(base.DB)
	if err == nil {
		RowCacheIndex.Set("Media_CreatedTime:"+fmt.Sprintf("%v", row.CreatedTime), row, 0)
		return row, true
	}

	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadMedia_ByCreatedTimes(CreatedTimes []int) {
	not_cached := make([]int, 0, len(CreatedTimes))

	for _, id := range CreatedTimes {
		_, ok := RowCacheIndex.Get("Media_CreatedTime:" + fmt.Sprintf("%v", id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		rows, err := NewMedia_Selector().CreatedTime_In(not_cached).GetRows(base.DB)
		if err == nil {
			for _, row := range rows {
				RowCacheIndex.Set("Media_CreatedTime:"+fmt.Sprintf("%v", row.CreatedTime), row, 0)
			}
		}
	}
}

//field//field//field

///// Generated from index 'AlbumId'.
func (c _StoreImpl) Media_ByAlbumId(AlbumId int) (*Media, bool) {
	o, ok := RowCacheIndex.Get("Media_AlbumId:" + fmt.Sprintf("%v", AlbumId))
	if ok {
		if obj, ok := o.(*Media); ok {
			return obj, true
		}
	}

	row, err := NewMedia_Selector().AlbumId_Eq(AlbumId).GetRow(base.DB)
	if err == nil {
		RowCacheIndex.Set("Media_AlbumId:"+fmt.Sprintf("%v", row.AlbumId), row, 0)
		return row, true
	}

	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadMedia_ByAlbumIds(AlbumIds []int) {
	not_cached := make([]int, 0, len(AlbumIds))

	for _, id := range AlbumIds {
		_, ok := RowCacheIndex.Get("Media_AlbumId:" + fmt.Sprintf("%v", id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		rows, err := NewMedia_Selector().AlbumId_In(not_cached).GetRows(base.DB)
		if err == nil {
			for _, row := range rows {
				RowCacheIndex.Set("Media_AlbumId:"+fmt.Sprintf("%v", row.AlbumId), row, 0)
			}
		}
	}
}

//field//field//field

///// Generated from index 'PostId2'.
func (c _StoreImpl) Media_ByPostId(PostId int) (*Media, bool) {
	o, ok := RowCacheIndex.Get("Media_PostId2:" + fmt.Sprintf("%v", PostId))
	if ok {
		if obj, ok := o.(*Media); ok {
			return obj, true
		}
	}

	row, err := NewMedia_Selector().PostId_Eq(PostId).GetRow(base.DB)
	if err == nil {
		RowCacheIndex.Set("Media_PostId2:"+fmt.Sprintf("%v", row.PostId), row, 0)
		return row, true
	}

	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadMedia_ByPostIds(PostIds []int) {
	not_cached := make([]int, 0, len(PostIds))

	for _, id := range PostIds {
		_, ok := RowCacheIndex.Get("Media_PostId2:" + fmt.Sprintf("%v", id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		rows, err := NewMedia_Selector().PostId_In(not_cached).GetRows(base.DB)
		if err == nil {
			for _, row := range rows {
				RowCacheIndex.Set("Media_PostId2:"+fmt.Sprintf("%v", row.PostId), row, 0)
			}
		}
	}
}

// MessageFile - PRIMARY

// Notify - PRIMARY

// Notify - ForUserId

// NotifyRemoved - PRIMARY

// PhoneContact - PRIMARY

// Post - PRIMARY

//field//field//field

///// Generated from index 'UserId'.
func (c _StoreImpl) Post_ByUserId(UserId int) (*Post, bool) {
	o, ok := RowCacheIndex.Get("Post_UserId:" + fmt.Sprintf("%v", UserId))
	if ok {
		if obj, ok := o.(*Post); ok {
			return obj, true
		}
	}

	row, err := NewPost_Selector().UserId_Eq(UserId).GetRow(base.DB)
	if err == nil {
		RowCacheIndex.Set("Post_UserId:"+fmt.Sprintf("%v", row.UserId), row, 0)
		return row, true
	}

	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadPost_ByUserIds(UserIds []int) {
	not_cached := make([]int, 0, len(UserIds))

	for _, id := range UserIds {
		_, ok := RowCacheIndex.Get("Post_UserId:" + fmt.Sprintf("%v", id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		rows, err := NewPost_Selector().UserId_In(not_cached).GetRows(base.DB)
		if err == nil {
			for _, row := range rows {
				RowCacheIndex.Set("Post_UserId:"+fmt.Sprintf("%v", row.UserId), row, 0)
			}
		}
	}
}

// SearchClicked - PRIMARY

// Session - PRIMARY

//field//field//field

///// Generated from index 'SessionUuid2'.
func (c _StoreImpl) Session_BySessionUuid(SessionUuid string) (*Session, bool) {
	o, ok := RowCacheIndex.Get("Session_SessionUuid2:" + fmt.Sprintf("%v", SessionUuid))
	if ok {
		if obj, ok := o.(*Session); ok {
			return obj, true
		}
	}

	row, err := NewSession_Selector().SessionUuid_Eq(SessionUuid).GetRow(base.DB)
	if err == nil {
		RowCacheIndex.Set("Session_SessionUuid2:"+fmt.Sprintf("%v", row.SessionUuid), row, 0)
		return row, true
	}

	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadSession_BySessionUuids(SessionUuids []string) {
	not_cached := make([]string, 0, len(SessionUuids))

	for _, id := range SessionUuids {
		_, ok := RowCacheIndex.Get("Session_SessionUuid2:" + fmt.Sprintf("%v", id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		rows, err := NewSession_Selector().SessionUuid_In(not_cached).GetRows(base.DB)
		if err == nil {
			for _, row := range rows {
				RowCacheIndex.Set("Session_SessionUuid2:"+fmt.Sprintf("%v", row.SessionUuid), row, 0)
			}
		}
	}
}

//field//field//field

///// Generated from index 'Id'.
func (c _StoreImpl) Session_ById(Id int) (*Session, bool) {
	o, ok := RowCacheIndex.Get("Session_Id:" + fmt.Sprintf("%v", Id))
	if ok {
		if obj, ok := o.(*Session); ok {
			return obj, true
		}
	}

	row, err := NewSession_Selector().Id_Eq(Id).GetRow(base.DB)
	if err == nil {
		RowCacheIndex.Set("Session_Id:"+fmt.Sprintf("%v", row.Id), row, 0)
		return row, true
	}

	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadSession_ByIds(Ids []int) {
	not_cached := make([]int, 0, len(Ids))

	for _, id := range Ids {
		_, ok := RowCacheIndex.Get("Session_Id:" + fmt.Sprintf("%v", id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		rows, err := NewSession_Selector().Id_In(not_cached).GetRows(base.DB)
		if err == nil {
			for _, row := range rows {
				RowCacheIndex.Set("Session_Id:"+fmt.Sprintf("%v", row.Id), row, 0)
			}
		}
	}
}

//field//field//field

///// Generated from index 'UserId'.
func (c _StoreImpl) Session_ByUserId(UserId int) (*Session, bool) {
	o, ok := RowCacheIndex.Get("Session_UserId:" + fmt.Sprintf("%v", UserId))
	if ok {
		if obj, ok := o.(*Session); ok {
			return obj, true
		}
	}

	row, err := NewSession_Selector().UserId_Eq(UserId).GetRow(base.DB)
	if err == nil {
		RowCacheIndex.Set("Session_UserId:"+fmt.Sprintf("%v", row.UserId), row, 0)
		return row, true
	}

	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadSession_ByUserIds(UserIds []int) {
	not_cached := make([]int, 0, len(UserIds))

	for _, id := range UserIds {
		_, ok := RowCacheIndex.Get("Session_UserId:" + fmt.Sprintf("%v", id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		rows, err := NewSession_Selector().UserId_In(not_cached).GetRows(base.DB)
		if err == nil {
			for _, row := range rows {
				RowCacheIndex.Set("Session_UserId:"+fmt.Sprintf("%v", row.UserId), row, 0)
			}
		}
	}
}

// SettingClient - PRIMARY

// SettingNotification - PRIMARY

// SuggestedTopPost - PRIMARY

// SuggestedUser - PRIMARY

// Tag - PRIMARY

//field//field//field

///// Generated from index 'Name'.
func (c _StoreImpl) Tag_ByName(Name string) (*Tag, bool) {
	o, ok := RowCacheIndex.Get("Tag_Name:" + fmt.Sprintf("%v", Name))
	if ok {
		if obj, ok := o.(*Tag); ok {
			return obj, true
		}
	}

	row, err := NewTag_Selector().Name_Eq(Name).GetRow(base.DB)
	if err == nil {
		RowCacheIndex.Set("Tag_Name:"+fmt.Sprintf("%v", row.Name), row, 0)
		return row, true
	}

	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadTag_ByNames(Names []string) {
	not_cached := make([]string, 0, len(Names))

	for _, id := range Names {
		_, ok := RowCacheIndex.Get("Tag_Name:" + fmt.Sprintf("%v", id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		rows, err := NewTag_Selector().Name_In(not_cached).GetRows(base.DB)
		if err == nil {
			for _, row := range rows {
				RowCacheIndex.Set("Tag_Name:"+fmt.Sprintf("%v", row.Name), row, 0)
			}
		}
	}
}

// TagsPost - PRIMARY

// TagsPost - TagId

// TriggerLog - PRIMARY

//field//field//field

///// Generated from index 'Id'.
func (c _StoreImpl) TriggerLog_ById(Id int) (*TriggerLog, bool) {
	o, ok := RowCacheIndex.Get("TriggerLog_Id:" + fmt.Sprintf("%v", Id))
	if ok {
		if obj, ok := o.(*TriggerLog); ok {
			return obj, true
		}
	}

	row, err := NewTriggerLog_Selector().Id_Eq(Id).GetRow(base.DB)
	if err == nil {
		RowCacheIndex.Set("TriggerLog_Id:"+fmt.Sprintf("%v", row.Id), row, 0)
		return row, true
	}

	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadTriggerLog_ByIds(Ids []int) {
	not_cached := make([]int, 0, len(Ids))

	for _, id := range Ids {
		_, ok := RowCacheIndex.Get("TriggerLog_Id:" + fmt.Sprintf("%v", id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		rows, err := NewTriggerLog_Selector().Id_In(not_cached).GetRows(base.DB)
		if err == nil {
			for _, row := range rows {
				RowCacheIndex.Set("TriggerLog_Id:"+fmt.Sprintf("%v", row.Id), row, 0)
			}
		}
	}
}

// User - PRIMARY

//field//field//field

///// Generated from index 'UserName'.
func (c _StoreImpl) User_ByUserName(UserName string) (*User, bool) {
	o, ok := RowCacheIndex.Get("User_UserName:" + fmt.Sprintf("%v", UserName))
	if ok {
		if obj, ok := o.(*User); ok {
			return obj, true
		}
	}

	row, err := NewUser_Selector().UserName_Eq(UserName).GetRow(base.DB)
	if err == nil {
		RowCacheIndex.Set("User_UserName:"+fmt.Sprintf("%v", row.UserName), row, 0)
		return row, true
	}

	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadUser_ByUserNames(UserNames []string) {
	not_cached := make([]string, 0, len(UserNames))

	for _, id := range UserNames {
		_, ok := RowCacheIndex.Get("User_UserName:" + fmt.Sprintf("%v", id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		rows, err := NewUser_Selector().UserName_In(not_cached).GetRows(base.DB)
		if err == nil {
			for _, row := range rows {
				RowCacheIndex.Set("User_UserName:"+fmt.Sprintf("%v", row.UserName), row, 0)
			}
		}
	}
}

//field//field//field

///// Generated from index 'Email'.
func (c _StoreImpl) User_ByEmail(Email string) (*User, bool) {
	o, ok := RowCacheIndex.Get("User_Email:" + fmt.Sprintf("%v", Email))
	if ok {
		if obj, ok := o.(*User); ok {
			return obj, true
		}
	}

	row, err := NewUser_Selector().Email_Eq(Email).GetRow(base.DB)
	if err == nil {
		RowCacheIndex.Set("User_Email:"+fmt.Sprintf("%v", row.Email), row, 0)
		return row, true
	}

	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadUser_ByEmails(Emails []string) {
	not_cached := make([]string, 0, len(Emails))

	for _, id := range Emails {
		_, ok := RowCacheIndex.Get("User_Email:" + fmt.Sprintf("%v", id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		rows, err := NewUser_Selector().Email_In(not_cached).GetRows(base.DB)
		if err == nil {
			for _, row := range rows {
				RowCacheIndex.Set("User_Email:"+fmt.Sprintf("%v", row.Email), row, 0)
			}
		}
	}
}

//field//field//field

///// Generated from index 'Phone'.
func (c _StoreImpl) User_ByPhone2(Phone2 string) (*User, bool) {
	o, ok := RowCacheIndex.Get("User_Phone:" + fmt.Sprintf("%v", Phone2))
	if ok {
		if obj, ok := o.(*User); ok {
			return obj, true
		}
	}

	row, err := NewUser_Selector().Phone2_Eq(Phone2).GetRow(base.DB)
	if err == nil {
		RowCacheIndex.Set("User_Phone:"+fmt.Sprintf("%v", row.Phone2), row, 0)
		return row, true
	}

	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadUser_ByPhone2s(Phone2s []string) {
	not_cached := make([]string, 0, len(Phone2s))

	for _, id := range Phone2s {
		_, ok := RowCacheIndex.Get("User_Phone:" + fmt.Sprintf("%v", id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		rows, err := NewUser_Selector().Phone2_In(not_cached).GetRows(base.DB)
		if err == nil {
			for _, row := range rows {
				RowCacheIndex.Set("User_Phone:"+fmt.Sprintf("%v", row.Phone2), row, 0)
			}
		}
	}
}

// UserMetaInfo - PRIMARY

//field//field//field

///// Generated from index 'UserId2'.
func (c _StoreImpl) UserMetaInfo_ByUserId(UserId int) (*UserMetaInfo, bool) {
	o, ok := RowCacheIndex.Get("UserMetaInfo_UserId2:" + fmt.Sprintf("%v", UserId))
	if ok {
		if obj, ok := o.(*UserMetaInfo); ok {
			return obj, true
		}
	}

	row, err := NewUserMetaInfo_Selector().UserId_Eq(UserId).GetRow(base.DB)
	if err == nil {
		RowCacheIndex.Set("UserMetaInfo_UserId2:"+fmt.Sprintf("%v", row.UserId), row, 0)
		return row, true
	}

	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadUserMetaInfo_ByUserIds(UserIds []int) {
	not_cached := make([]int, 0, len(UserIds))

	for _, id := range UserIds {
		_, ok := RowCacheIndex.Get("UserMetaInfo_UserId2:" + fmt.Sprintf("%v", id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		rows, err := NewUserMetaInfo_Selector().UserId_In(not_cached).GetRows(base.DB)
		if err == nil {
			for _, row := range rows {
				RowCacheIndex.Set("UserMetaInfo_UserId2:"+fmt.Sprintf("%v", row.UserId), row, 0)
			}
		}
	}
}

// UserPassword - PRIMARY
