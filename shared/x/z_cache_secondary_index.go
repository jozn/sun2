package x

import (
	"errors"
	"fmt"
	"ms/sun/base"
)

//TODO: WE MUST separate int from string to not let empty string "" from preloading or loading and inserting into caches

// Action - PRIMARY

// Action - ActorUserId

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

func (c _StoreImpl) Comment_ByPostId_JustCache(PostId int) (*Comment, bool) {
	o, ok := RowCacheIndex.Get("Comment_PostId:" + fmt.Sprintf("%v", PostId))
	if ok {
		if obj, ok := o.(*Comment); ok {
			return obj, true
		}
	}

	XOLogErr(errors.New("_JustCache is empty for secondry index " + "Comment_PostId:" + fmt.Sprintf("%v", PostId)))
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

// Event - PRIMARY

// FollowingList - PRIMARY

// FollowingListMember - PRIMARY

// FollowingListMember - UserId

// FollowingListMember - FollowedUserId

// FollowingListMember - UserId_2

// FollowingListMemberRemoved - PRIMARY

// FollowingListMemberRemoved - UserId

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

func (c _StoreImpl) GroupMember_ById_JustCache(Id int) (*GroupMember, bool) {
	o, ok := RowCacheIndex.Get("GroupMember_Id:" + fmt.Sprintf("%v", Id))
	if ok {
		if obj, ok := o.(*GroupMember); ok {
			return obj, true
		}
	}

	XOLogErr(errors.New("_JustCache is empty for secondry index " + "GroupMember_Id:" + fmt.Sprintf("%v", Id)))
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

func (c _StoreImpl) Like_ById_JustCache(Id int) (*Like, bool) {
	o, ok := RowCacheIndex.Get("Like_Id:" + fmt.Sprintf("%v", Id))
	if ok {
		if obj, ok := o.(*Like); ok {
			return obj, true
		}
	}

	XOLogErr(errors.New("_JustCache is empty for secondry index " + "Like_Id:" + fmt.Sprintf("%v", Id)))
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

func (c _StoreImpl) Like_ByPostId_JustCache(PostId int) (*Like, bool) {
	o, ok := RowCacheIndex.Get("Like_PostId_2:" + fmt.Sprintf("%v", PostId))
	if ok {
		if obj, ok := o.(*Like); ok {
			return obj, true
		}
	}

	XOLogErr(errors.New("_JustCache is empty for secondry index " + "Like_PostId_2:" + fmt.Sprintf("%v", PostId)))
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

func (c _StoreImpl) Media_ByMd5Hash_JustCache(Md5Hash string) (*Media, bool) {
	o, ok := RowCacheIndex.Get("Media_HashMd5:" + fmt.Sprintf("%v", Md5Hash))
	if ok {
		if obj, ok := o.(*Media); ok {
			return obj, true
		}
	}

	XOLogErr(errors.New("_JustCache is empty for secondry index " + "Media_HashMd5:" + fmt.Sprintf("%v", Md5Hash)))
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

func (c _StoreImpl) Media_ByCreatedTime_JustCache(CreatedTime int) (*Media, bool) {
	o, ok := RowCacheIndex.Get("Media_CreatedTime:" + fmt.Sprintf("%v", CreatedTime))
	if ok {
		if obj, ok := o.(*Media); ok {
			return obj, true
		}
	}

	XOLogErr(errors.New("_JustCache is empty for secondry index " + "Media_CreatedTime:" + fmt.Sprintf("%v", CreatedTime)))
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

func (c _StoreImpl) Media_ByAlbumId_JustCache(AlbumId int) (*Media, bool) {
	o, ok := RowCacheIndex.Get("Media_AlbumId:" + fmt.Sprintf("%v", AlbumId))
	if ok {
		if obj, ok := o.(*Media); ok {
			return obj, true
		}
	}

	XOLogErr(errors.New("_JustCache is empty for secondry index " + "Media_AlbumId:" + fmt.Sprintf("%v", AlbumId)))
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

func (c _StoreImpl) Media_ByPostId_JustCache(PostId int) (*Media, bool) {
	o, ok := RowCacheIndex.Get("Media_PostId2:" + fmt.Sprintf("%v", PostId))
	if ok {
		if obj, ok := o.(*Media); ok {
			return obj, true
		}
	}

	XOLogErr(errors.New("_JustCache is empty for secondry index " + "Media_PostId2:" + fmt.Sprintf("%v", PostId)))
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

func (c _StoreImpl) Post_ByUserId_JustCache(UserId int) (*Post, bool) {
	o, ok := RowCacheIndex.Get("Post_UserId:" + fmt.Sprintf("%v", UserId))
	if ok {
		if obj, ok := o.(*Post); ok {
			return obj, true
		}
	}

	XOLogErr(errors.New("_JustCache is empty for secondry index " + "Post_UserId:" + fmt.Sprintf("%v", UserId)))
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

// PostKey - PRIMARY

//field//field//field

///// Generated from index 'PostKey'.
func (c _StoreImpl) PostKey_ByPostKeyStr(PostKeyStr string) (*PostKey, bool) {
	o, ok := RowCacheIndex.Get("PostKey_PostKey:" + fmt.Sprintf("%v", PostKeyStr))
	if ok {
		if obj, ok := o.(*PostKey); ok {
			return obj, true
		}
	}

	row, err := NewPostKey_Selector().PostKeyStr_Eq(PostKeyStr).GetRow(base.DB)
	if err == nil {
		RowCacheIndex.Set("PostKey_PostKey:"+fmt.Sprintf("%v", row.PostKeyStr), row, 0)
		return row, true
	}

	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PostKey_ByPostKeyStr_JustCache(PostKeyStr string) (*PostKey, bool) {
	o, ok := RowCacheIndex.Get("PostKey_PostKey:" + fmt.Sprintf("%v", PostKeyStr))
	if ok {
		if obj, ok := o.(*PostKey); ok {
			return obj, true
		}
	}

	XOLogErr(errors.New("_JustCache is empty for secondry index " + "PostKey_PostKey:" + fmt.Sprintf("%v", PostKeyStr)))
	return nil, false
}

func (c _StoreImpl) PreLoadPostKey_ByPostKeyStrs(PostKeyStrs []string) {
	not_cached := make([]string, 0, len(PostKeyStrs))

	for _, id := range PostKeyStrs {
		_, ok := RowCacheIndex.Get("PostKey_PostKey:" + fmt.Sprintf("%v", id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		rows, err := NewPostKey_Selector().PostKeyStr_In(not_cached).GetRows(base.DB)
		if err == nil {
			for _, row := range rows {
				RowCacheIndex.Set("PostKey_PostKey:"+fmt.Sprintf("%v", row.PostKeyStr), row, 0)
			}
		}
	}
}

//field//field//field

///// Generated from index 'Used'.
func (c _StoreImpl) PostKey_ByUsed(Used int) (*PostKey, bool) {
	o, ok := RowCacheIndex.Get("PostKey_Used:" + fmt.Sprintf("%v", Used))
	if ok {
		if obj, ok := o.(*PostKey); ok {
			return obj, true
		}
	}

	row, err := NewPostKey_Selector().Used_Eq(Used).GetRow(base.DB)
	if err == nil {
		RowCacheIndex.Set("PostKey_Used:"+fmt.Sprintf("%v", row.Used), row, 0)
		return row, true
	}

	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PostKey_ByUsed_JustCache(Used int) (*PostKey, bool) {
	o, ok := RowCacheIndex.Get("PostKey_Used:" + fmt.Sprintf("%v", Used))
	if ok {
		if obj, ok := o.(*PostKey); ok {
			return obj, true
		}
	}

	XOLogErr(errors.New("_JustCache is empty for secondry index " + "PostKey_Used:" + fmt.Sprintf("%v", Used)))
	return nil, false
}

func (c _StoreImpl) PreLoadPostKey_ByUseds(Useds []int) {
	not_cached := make([]int, 0, len(Useds))

	for _, id := range Useds {
		_, ok := RowCacheIndex.Get("PostKey_Used:" + fmt.Sprintf("%v", id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		rows, err := NewPostKey_Selector().Used_In(not_cached).GetRows(base.DB)
		if err == nil {
			for _, row := range rows {
				RowCacheIndex.Set("PostKey_Used:"+fmt.Sprintf("%v", row.Used), row, 0)
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

func (c _StoreImpl) Session_BySessionUuid_JustCache(SessionUuid string) (*Session, bool) {
	o, ok := RowCacheIndex.Get("Session_SessionUuid2:" + fmt.Sprintf("%v", SessionUuid))
	if ok {
		if obj, ok := o.(*Session); ok {
			return obj, true
		}
	}

	XOLogErr(errors.New("_JustCache is empty for secondry index " + "Session_SessionUuid2:" + fmt.Sprintf("%v", SessionUuid)))
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

func (c _StoreImpl) Session_ByUserId_JustCache(UserId int) (*Session, bool) {
	o, ok := RowCacheIndex.Get("Session_UserId:" + fmt.Sprintf("%v", UserId))
	if ok {
		if obj, ok := o.(*Session); ok {
			return obj, true
		}
	}

	XOLogErr(errors.New("_JustCache is empty for secondry index " + "Session_UserId:" + fmt.Sprintf("%v", UserId)))
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

func (c _StoreImpl) Tag_ByName_JustCache(Name string) (*Tag, bool) {
	o, ok := RowCacheIndex.Get("Tag_Name:" + fmt.Sprintf("%v", Name))
	if ok {
		if obj, ok := o.(*Tag); ok {
			return obj, true
		}
	}

	XOLogErr(errors.New("_JustCache is empty for secondry index " + "Tag_Name:" + fmt.Sprintf("%v", Name)))
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

func (c _StoreImpl) TriggerLog_ById_JustCache(Id int) (*TriggerLog, bool) {
	o, ok := RowCacheIndex.Get("TriggerLog_Id:" + fmt.Sprintf("%v", Id))
	if ok {
		if obj, ok := o.(*TriggerLog); ok {
			return obj, true
		}
	}

	XOLogErr(errors.New("_JustCache is empty for secondry index " + "TriggerLog_Id:" + fmt.Sprintf("%v", Id)))
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

func (c _StoreImpl) User_ByUserName_JustCache(UserName string) (*User, bool) {
	o, ok := RowCacheIndex.Get("User_UserName:" + fmt.Sprintf("%v", UserName))
	if ok {
		if obj, ok := o.(*User); ok {
			return obj, true
		}
	}

	XOLogErr(errors.New("_JustCache is empty for secondry index " + "User_UserName:" + fmt.Sprintf("%v", UserName)))
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

func (c _StoreImpl) User_ByEmail_JustCache(Email string) (*User, bool) {
	o, ok := RowCacheIndex.Get("User_Email:" + fmt.Sprintf("%v", Email))
	if ok {
		if obj, ok := o.(*User); ok {
			return obj, true
		}
	}

	XOLogErr(errors.New("_JustCache is empty for secondry index " + "User_Email:" + fmt.Sprintf("%v", Email)))
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

func (c _StoreImpl) User_ByPhone2_JustCache(Phone2 string) (*User, bool) {
	o, ok := RowCacheIndex.Get("User_Phone:" + fmt.Sprintf("%v", Phone2))
	if ok {
		if obj, ok := o.(*User); ok {
			return obj, true
		}
	}

	XOLogErr(errors.New("_JustCache is empty for secondry index " + "User_Phone:" + fmt.Sprintf("%v", Phone2)))
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

func (c _StoreImpl) UserMetaInfo_ByUserId_JustCache(UserId int) (*UserMetaInfo, bool) {
	o, ok := RowCacheIndex.Get("UserMetaInfo_UserId2:" + fmt.Sprintf("%v", UserId))
	if ok {
		if obj, ok := o.(*UserMetaInfo); ok {
			return obj, true
		}
	}

	XOLogErr(errors.New("_JustCache is empty for secondry index " + "UserMetaInfo_UserId2:" + fmt.Sprintf("%v", UserId)))
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

// Chat - PRIMARY

// ChatLastMessage - PRIMARY

// ChatSync - PRIMARY

// DirectMessage - PRIMARY

// Home - PRIMARY

//field//field//field

///// Generated from index 'ForUserId'.
func (c _StoreImpl) Home_ByForUserId(ForUserId int) (*Home, bool) {
	o, ok := RowCacheIndex.Get("Home_ForUserId:" + fmt.Sprintf("%v", ForUserId))
	if ok {
		if obj, ok := o.(*Home); ok {
			return obj, true
		}
	}

	row, err := NewHome_Selector().ForUserId_Eq(ForUserId).GetRow(base.DB)
	if err == nil {
		RowCacheIndex.Set("Home_ForUserId:"+fmt.Sprintf("%v", row.ForUserId), row, 0)
		return row, true
	}

	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) Home_ByForUserId_JustCache(ForUserId int) (*Home, bool) {
	o, ok := RowCacheIndex.Get("Home_ForUserId:" + fmt.Sprintf("%v", ForUserId))
	if ok {
		if obj, ok := o.(*Home); ok {
			return obj, true
		}
	}

	XOLogErr(errors.New("_JustCache is empty for secondry index " + "Home_ForUserId:" + fmt.Sprintf("%v", ForUserId)))
	return nil, false
}

func (c _StoreImpl) PreLoadHome_ByForUserIds(ForUserIds []int) {
	not_cached := make([]int, 0, len(ForUserIds))

	for _, id := range ForUserIds {
		_, ok := RowCacheIndex.Get("Home_ForUserId:" + fmt.Sprintf("%v", id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		rows, err := NewHome_Selector().ForUserId_In(not_cached).GetRows(base.DB)
		if err == nil {
			for _, row := range rows {
				RowCacheIndex.Set("Home_ForUserId:"+fmt.Sprintf("%v", row.ForUserId), row, 0)
			}
		}
	}
}

//field//field//field

///// Generated from index 'PostId'.
func (c _StoreImpl) Home_ByPostId(PostId int) (*Home, bool) {
	o, ok := RowCacheIndex.Get("Home_PostId:" + fmt.Sprintf("%v", PostId))
	if ok {
		if obj, ok := o.(*Home); ok {
			return obj, true
		}
	}

	row, err := NewHome_Selector().PostId_Eq(PostId).GetRow(base.DB)
	if err == nil {
		RowCacheIndex.Set("Home_PostId:"+fmt.Sprintf("%v", row.PostId), row, 0)
		return row, true
	}

	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) Home_ByPostId_JustCache(PostId int) (*Home, bool) {
	o, ok := RowCacheIndex.Get("Home_PostId:" + fmt.Sprintf("%v", PostId))
	if ok {
		if obj, ok := o.(*Home); ok {
			return obj, true
		}
	}

	XOLogErr(errors.New("_JustCache is empty for secondry index " + "Home_PostId:" + fmt.Sprintf("%v", PostId)))
	return nil, false
}

func (c _StoreImpl) PreLoadHome_ByPostIds(PostIds []int) {
	not_cached := make([]int, 0, len(PostIds))

	for _, id := range PostIds {
		_, ok := RowCacheIndex.Get("Home_PostId:" + fmt.Sprintf("%v", id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		rows, err := NewHome_Selector().PostId_In(not_cached).GetRows(base.DB)
		if err == nil {
			for _, row := range rows {
				RowCacheIndex.Set("Home_PostId:"+fmt.Sprintf("%v", row.PostId), row, 0)
			}
		}
	}
}

// MessageFile - PRIMARY

// FileMsg - PRIMARY

// FilePost - PRIMARY
