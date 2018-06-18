package x

import (
	"errors"
	"fmt"
	"ms/sun/shared/base"
)

//TODO: WE MUST separate int from string to not let empty string "" from preloading or loading and inserting into caches

// Action - PRIMARY

// Action - ActorUserId

// Blocked - PRIMARY

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

// CommentDeleted - PRIMARY

// Event - PRIMARY

// Followed - PRIMARY

// Likes - PRIMARY

// Likes - PostId

//field//field//field

///// Generated from index 'Id'.
func (c _StoreImpl) Likes_ById(Id int) (*Likes, bool) {
	o, ok := RowCacheIndex.Get("Likes_Id:" + fmt.Sprintf("%v", Id))
	if ok {
		if obj, ok := o.(*Likes); ok {
			return obj, true
		}
	}

	row, err := NewLikes_Selector().Id_Eq(Id).GetRow(base.DB)
	if err == nil {
		RowCacheIndex.Set("Likes_Id:"+fmt.Sprintf("%v", row.Id), row, 0)
		return row, true
	}

	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) Likes_ById_JustCache(Id int) (*Likes, bool) {
	o, ok := RowCacheIndex.Get("Likes_Id:" + fmt.Sprintf("%v", Id))
	if ok {
		if obj, ok := o.(*Likes); ok {
			return obj, true
		}
	}

	XOLogErr(errors.New("_JustCache is empty for secondry index " + "Likes_Id:" + fmt.Sprintf("%v", Id)))
	return nil, false
}

func (c _StoreImpl) PreLoadLikes_ByIds(Ids []int) {
	not_cached := make([]int, 0, len(Ids))

	for _, id := range Ids {
		_, ok := RowCacheIndex.Get("Likes_Id:" + fmt.Sprintf("%v", id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		rows, err := NewLikes_Selector().Id_In(not_cached).GetRows(base.DB)
		if err == nil {
			for _, row := range rows {
				RowCacheIndex.Set("Likes_Id:"+fmt.Sprintf("%v", row.Id), row, 0)
			}
		}
	}
}

//field//field//field

///// Generated from index 'PostId_2'.
func (c _StoreImpl) Likes_ByPostId(PostId int) (*Likes, bool) {
	o, ok := RowCacheIndex.Get("Likes_PostId_2:" + fmt.Sprintf("%v", PostId))
	if ok {
		if obj, ok := o.(*Likes); ok {
			return obj, true
		}
	}

	row, err := NewLikes_Selector().PostId_Eq(PostId).GetRow(base.DB)
	if err == nil {
		RowCacheIndex.Set("Likes_PostId_2:"+fmt.Sprintf("%v", row.PostId), row, 0)
		return row, true
	}

	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) Likes_ByPostId_JustCache(PostId int) (*Likes, bool) {
	o, ok := RowCacheIndex.Get("Likes_PostId_2:" + fmt.Sprintf("%v", PostId))
	if ok {
		if obj, ok := o.(*Likes); ok {
			return obj, true
		}
	}

	XOLogErr(errors.New("_JustCache is empty for secondry index " + "Likes_PostId_2:" + fmt.Sprintf("%v", PostId)))
	return nil, false
}

func (c _StoreImpl) PreLoadLikes_ByPostIds(PostIds []int) {
	not_cached := make([]int, 0, len(PostIds))

	for _, id := range PostIds {
		_, ok := RowCacheIndex.Get("Likes_PostId_2:" + fmt.Sprintf("%v", id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		rows, err := NewLikes_Selector().PostId_In(not_cached).GetRows(base.DB)
		if err == nil {
			for _, row := range rows {
				RowCacheIndex.Set("Likes_PostId_2:"+fmt.Sprintf("%v", row.PostId), row, 0)
			}
		}
	}
}

// Notify - PRIMARY

// Notify - ForUserId

// NotifyRemoved - PRIMARY

// PhoneContacts - PRIMARY

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

// PostCount - PRIMARY

// PostDeleted - PRIMARY

// PostKeys - PRIMARY

//field//field//field

///// Generated from index 'PostKey'.
func (c _StoreImpl) PostKeys_ByPostKeyStr(PostKeyStr string) (*PostKeys, bool) {
	o, ok := RowCacheIndex.Get("PostKeys_PostKey:" + fmt.Sprintf("%v", PostKeyStr))
	if ok {
		if obj, ok := o.(*PostKeys); ok {
			return obj, true
		}
	}

	row, err := NewPostKeys_Selector().PostKeyStr_Eq(PostKeyStr).GetRow(base.DB)
	if err == nil {
		RowCacheIndex.Set("PostKeys_PostKey:"+fmt.Sprintf("%v", row.PostKeyStr), row, 0)
		return row, true
	}

	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PostKeys_ByPostKeyStr_JustCache(PostKeyStr string) (*PostKeys, bool) {
	o, ok := RowCacheIndex.Get("PostKeys_PostKey:" + fmt.Sprintf("%v", PostKeyStr))
	if ok {
		if obj, ok := o.(*PostKeys); ok {
			return obj, true
		}
	}

	XOLogErr(errors.New("_JustCache is empty for secondry index " + "PostKeys_PostKey:" + fmt.Sprintf("%v", PostKeyStr)))
	return nil, false
}

func (c _StoreImpl) PreLoadPostKeys_ByPostKeyStrs(PostKeyStrs []string) {
	not_cached := make([]string, 0, len(PostKeyStrs))

	for _, id := range PostKeyStrs {
		_, ok := RowCacheIndex.Get("PostKeys_PostKey:" + fmt.Sprintf("%v", id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		rows, err := NewPostKeys_Selector().PostKeyStr_In(not_cached).GetRows(base.DB)
		if err == nil {
			for _, row := range rows {
				RowCacheIndex.Set("PostKeys_PostKey:"+fmt.Sprintf("%v", row.PostKeyStr), row, 0)
			}
		}
	}
}

//field//field//field

///// Generated from index 'Used'.
func (c _StoreImpl) PostKeys_ByUsed(Used int) (*PostKeys, bool) {
	o, ok := RowCacheIndex.Get("PostKeys_Used:" + fmt.Sprintf("%v", Used))
	if ok {
		if obj, ok := o.(*PostKeys); ok {
			return obj, true
		}
	}

	row, err := NewPostKeys_Selector().Used_Eq(Used).GetRow(base.DB)
	if err == nil {
		RowCacheIndex.Set("PostKeys_Used:"+fmt.Sprintf("%v", row.Used), row, 0)
		return row, true
	}

	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PostKeys_ByUsed_JustCache(Used int) (*PostKeys, bool) {
	o, ok := RowCacheIndex.Get("PostKeys_Used:" + fmt.Sprintf("%v", Used))
	if ok {
		if obj, ok := o.(*PostKeys); ok {
			return obj, true
		}
	}

	XOLogErr(errors.New("_JustCache is empty for secondry index " + "PostKeys_Used:" + fmt.Sprintf("%v", Used)))
	return nil, false
}

func (c _StoreImpl) PreLoadPostKeys_ByUseds(Useds []int) {
	not_cached := make([]int, 0, len(Useds))

	for _, id := range Useds {
		_, ok := RowCacheIndex.Get("PostKeys_Used:" + fmt.Sprintf("%v", id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		rows, err := NewPostKeys_Selector().Used_In(not_cached).GetRows(base.DB)
		if err == nil {
			for _, row := range rows {
				RowCacheIndex.Set("PostKeys_Used:"+fmt.Sprintf("%v", row.Used), row, 0)
			}
		}
	}
}

// PostLink - PRIMARY

// PostMedia - PRIMARY

//field//field//field

///// Generated from index 'HashMd5'.
func (c _StoreImpl) PostMedia_ByMd5Hash(Md5Hash string) (*PostMedia, bool) {
	o, ok := RowCacheIndex.Get("PostMedia_HashMd5:" + fmt.Sprintf("%v", Md5Hash))
	if ok {
		if obj, ok := o.(*PostMedia); ok {
			return obj, true
		}
	}

	row, err := NewPostMedia_Selector().Md5Hash_Eq(Md5Hash).GetRow(base.DB)
	if err == nil {
		RowCacheIndex.Set("PostMedia_HashMd5:"+fmt.Sprintf("%v", row.Md5Hash), row, 0)
		return row, true
	}

	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PostMedia_ByMd5Hash_JustCache(Md5Hash string) (*PostMedia, bool) {
	o, ok := RowCacheIndex.Get("PostMedia_HashMd5:" + fmt.Sprintf("%v", Md5Hash))
	if ok {
		if obj, ok := o.(*PostMedia); ok {
			return obj, true
		}
	}

	XOLogErr(errors.New("_JustCache is empty for secondry index " + "PostMedia_HashMd5:" + fmt.Sprintf("%v", Md5Hash)))
	return nil, false
}

func (c _StoreImpl) PreLoadPostMedia_ByMd5Hashs(Md5Hashs []string) {
	not_cached := make([]string, 0, len(Md5Hashs))

	for _, id := range Md5Hashs {
		_, ok := RowCacheIndex.Get("PostMedia_HashMd5:" + fmt.Sprintf("%v", id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		rows, err := NewPostMedia_Selector().Md5Hash_In(not_cached).GetRows(base.DB)
		if err == nil {
			for _, row := range rows {
				RowCacheIndex.Set("PostMedia_HashMd5:"+fmt.Sprintf("%v", row.Md5Hash), row, 0)
			}
		}
	}
}

//field//field//field

///// Generated from index 'CreatedTime'.
func (c _StoreImpl) PostMedia_ByCreatedTime(CreatedTime int) (*PostMedia, bool) {
	o, ok := RowCacheIndex.Get("PostMedia_CreatedTime:" + fmt.Sprintf("%v", CreatedTime))
	if ok {
		if obj, ok := o.(*PostMedia); ok {
			return obj, true
		}
	}

	row, err := NewPostMedia_Selector().CreatedTime_Eq(CreatedTime).GetRow(base.DB)
	if err == nil {
		RowCacheIndex.Set("PostMedia_CreatedTime:"+fmt.Sprintf("%v", row.CreatedTime), row, 0)
		return row, true
	}

	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PostMedia_ByCreatedTime_JustCache(CreatedTime int) (*PostMedia, bool) {
	o, ok := RowCacheIndex.Get("PostMedia_CreatedTime:" + fmt.Sprintf("%v", CreatedTime))
	if ok {
		if obj, ok := o.(*PostMedia); ok {
			return obj, true
		}
	}

	XOLogErr(errors.New("_JustCache is empty for secondry index " + "PostMedia_CreatedTime:" + fmt.Sprintf("%v", CreatedTime)))
	return nil, false
}

func (c _StoreImpl) PreLoadPostMedia_ByCreatedTimes(CreatedTimes []int) {
	not_cached := make([]int, 0, len(CreatedTimes))

	for _, id := range CreatedTimes {
		_, ok := RowCacheIndex.Get("PostMedia_CreatedTime:" + fmt.Sprintf("%v", id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		rows, err := NewPostMedia_Selector().CreatedTime_In(not_cached).GetRows(base.DB)
		if err == nil {
			for _, row := range rows {
				RowCacheIndex.Set("PostMedia_CreatedTime:"+fmt.Sprintf("%v", row.CreatedTime), row, 0)
			}
		}
	}
}

//field//field//field

///// Generated from index 'AlbumId'.
func (c _StoreImpl) PostMedia_ByAlbumId(AlbumId int) (*PostMedia, bool) {
	o, ok := RowCacheIndex.Get("PostMedia_AlbumId:" + fmt.Sprintf("%v", AlbumId))
	if ok {
		if obj, ok := o.(*PostMedia); ok {
			return obj, true
		}
	}

	row, err := NewPostMedia_Selector().AlbumId_Eq(AlbumId).GetRow(base.DB)
	if err == nil {
		RowCacheIndex.Set("PostMedia_AlbumId:"+fmt.Sprintf("%v", row.AlbumId), row, 0)
		return row, true
	}

	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PostMedia_ByAlbumId_JustCache(AlbumId int) (*PostMedia, bool) {
	o, ok := RowCacheIndex.Get("PostMedia_AlbumId:" + fmt.Sprintf("%v", AlbumId))
	if ok {
		if obj, ok := o.(*PostMedia); ok {
			return obj, true
		}
	}

	XOLogErr(errors.New("_JustCache is empty for secondry index " + "PostMedia_AlbumId:" + fmt.Sprintf("%v", AlbumId)))
	return nil, false
}

func (c _StoreImpl) PreLoadPostMedia_ByAlbumIds(AlbumIds []int) {
	not_cached := make([]int, 0, len(AlbumIds))

	for _, id := range AlbumIds {
		_, ok := RowCacheIndex.Get("PostMedia_AlbumId:" + fmt.Sprintf("%v", id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		rows, err := NewPostMedia_Selector().AlbumId_In(not_cached).GetRows(base.DB)
		if err == nil {
			for _, row := range rows {
				RowCacheIndex.Set("PostMedia_AlbumId:"+fmt.Sprintf("%v", row.AlbumId), row, 0)
			}
		}
	}
}

//field//field//field

///// Generated from index 'PostId2'.
func (c _StoreImpl) PostMedia_ByPostId(PostId int) (*PostMedia, bool) {
	o, ok := RowCacheIndex.Get("PostMedia_PostId2:" + fmt.Sprintf("%v", PostId))
	if ok {
		if obj, ok := o.(*PostMedia); ok {
			return obj, true
		}
	}

	row, err := NewPostMedia_Selector().PostId_Eq(PostId).GetRow(base.DB)
	if err == nil {
		RowCacheIndex.Set("PostMedia_PostId2:"+fmt.Sprintf("%v", row.PostId), row, 0)
		return row, true
	}

	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PostMedia_ByPostId_JustCache(PostId int) (*PostMedia, bool) {
	o, ok := RowCacheIndex.Get("PostMedia_PostId2:" + fmt.Sprintf("%v", PostId))
	if ok {
		if obj, ok := o.(*PostMedia); ok {
			return obj, true
		}
	}

	XOLogErr(errors.New("_JustCache is empty for secondry index " + "PostMedia_PostId2:" + fmt.Sprintf("%v", PostId)))
	return nil, false
}

func (c _StoreImpl) PreLoadPostMedia_ByPostIds(PostIds []int) {
	not_cached := make([]int, 0, len(PostIds))

	for _, id := range PostIds {
		_, ok := RowCacheIndex.Get("PostMedia_PostId2:" + fmt.Sprintf("%v", id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		rows, err := NewPostMedia_Selector().PostId_In(not_cached).GetRows(base.DB)
		if err == nil {
			for _, row := range rows {
				RowCacheIndex.Set("PostMedia_PostId2:"+fmt.Sprintf("%v", row.PostId), row, 0)
			}
		}
	}
}

// PostMentioned - PRIMARY

// PostReshared - PRIMARY

// Session - PRIMARY

// Session - SessionUuid_2

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

//field//field//field

///// Generated from index 'SessionUuid'.
func (c _StoreImpl) Session_BySessionUuid(SessionUuid string) (*Session, bool) {
	o, ok := RowCacheIndex.Get("Session_SessionUuid:" + fmt.Sprintf("%v", SessionUuid))
	if ok {
		if obj, ok := o.(*Session); ok {
			return obj, true
		}
	}

	row, err := NewSession_Selector().SessionUuid_Eq(SessionUuid).GetRow(base.DB)
	if err == nil {
		RowCacheIndex.Set("Session_SessionUuid:"+fmt.Sprintf("%v", row.SessionUuid), row, 0)
		return row, true
	}

	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) Session_BySessionUuid_JustCache(SessionUuid string) (*Session, bool) {
	o, ok := RowCacheIndex.Get("Session_SessionUuid:" + fmt.Sprintf("%v", SessionUuid))
	if ok {
		if obj, ok := o.(*Session); ok {
			return obj, true
		}
	}

	XOLogErr(errors.New("_JustCache is empty for secondry index " + "Session_SessionUuid:" + fmt.Sprintf("%v", SessionUuid)))
	return nil, false
}

func (c _StoreImpl) PreLoadSession_BySessionUuids(SessionUuids []string) {
	not_cached := make([]string, 0, len(SessionUuids))

	for _, id := range SessionUuids {
		_, ok := RowCacheIndex.Get("Session_SessionUuid:" + fmt.Sprintf("%v", id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		rows, err := NewSession_Selector().SessionUuid_In(not_cached).GetRows(base.DB)
		if err == nil {
			for _, row := range rows {
				RowCacheIndex.Set("Session_SessionUuid:"+fmt.Sprintf("%v", row.SessionUuid), row, 0)
			}
		}
	}
}

// SettingNotifications - PRIMARY

// Sms - PRIMARY

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

// TagPost - PRIMARY

// TagPost - TagId

// TriggerLog - PRIMARY

//field//field//field

///// Generated from index 'CreatedSe'.
func (c _StoreImpl) TriggerLog_ByCreatedSe(CreatedSe int) (*TriggerLog, bool) {
	o, ok := RowCacheIndex.Get("TriggerLog_CreatedSe:" + fmt.Sprintf("%v", CreatedSe))
	if ok {
		if obj, ok := o.(*TriggerLog); ok {
			return obj, true
		}
	}

	row, err := NewTriggerLog_Selector().CreatedSe_Eq(CreatedSe).GetRow(base.DB)
	if err == nil {
		RowCacheIndex.Set("TriggerLog_CreatedSe:"+fmt.Sprintf("%v", row.CreatedSe), row, 0)
		return row, true
	}

	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) TriggerLog_ByCreatedSe_JustCache(CreatedSe int) (*TriggerLog, bool) {
	o, ok := RowCacheIndex.Get("TriggerLog_CreatedSe:" + fmt.Sprintf("%v", CreatedSe))
	if ok {
		if obj, ok := o.(*TriggerLog); ok {
			return obj, true
		}
	}

	XOLogErr(errors.New("_JustCache is empty for secondry index " + "TriggerLog_CreatedSe:" + fmt.Sprintf("%v", CreatedSe)))
	return nil, false
}

func (c _StoreImpl) PreLoadTriggerLog_ByCreatedSes(CreatedSes []int) {
	not_cached := make([]int, 0, len(CreatedSes))

	for _, id := range CreatedSes {
		_, ok := RowCacheIndex.Get("TriggerLog_CreatedSe:" + fmt.Sprintf("%v", id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		rows, err := NewTriggerLog_Selector().CreatedSe_In(not_cached).GetRows(base.DB)
		if err == nil {
			for _, row := range rows {
				RowCacheIndex.Set("TriggerLog_CreatedSe:"+fmt.Sprintf("%v", row.CreatedSe), row, 0)
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

///// Generated from index 'Phone'.
func (c _StoreImpl) User_ByPhone(Phone string) (*User, bool) {
	o, ok := RowCacheIndex.Get("User_Phone:" + fmt.Sprintf("%v", Phone))
	if ok {
		if obj, ok := o.(*User); ok {
			return obj, true
		}
	}

	row, err := NewUser_Selector().Phone_Eq(Phone).GetRow(base.DB)
	if err == nil {
		RowCacheIndex.Set("User_Phone:"+fmt.Sprintf("%v", row.Phone), row, 0)
		return row, true
	}

	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) User_ByPhone_JustCache(Phone string) (*User, bool) {
	o, ok := RowCacheIndex.Get("User_Phone:" + fmt.Sprintf("%v", Phone))
	if ok {
		if obj, ok := o.(*User); ok {
			return obj, true
		}
	}

	XOLogErr(errors.New("_JustCache is empty for secondry index " + "User_Phone:" + fmt.Sprintf("%v", Phone)))
	return nil, false
}

func (c _StoreImpl) PreLoadUser_ByPhones(Phones []string) {
	not_cached := make([]string, 0, len(Phones))

	for _, id := range Phones {
		_, ok := RowCacheIndex.Get("User_Phone:" + fmt.Sprintf("%v", id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		rows, err := NewUser_Selector().Phone_In(not_cached).GetRows(base.DB)
		if err == nil {
			for _, row := range rows {
				RowCacheIndex.Set("User_Phone:"+fmt.Sprintf("%v", row.Phone), row, 0)
			}
		}
	}
}

//field//field//field

///// Generated from index 'UserNameLower'.
func (c _StoreImpl) User_ByUserNameLower(UserNameLower string) (*User, bool) {
	o, ok := RowCacheIndex.Get("User_UserNameLower:" + fmt.Sprintf("%v", UserNameLower))
	if ok {
		if obj, ok := o.(*User); ok {
			return obj, true
		}
	}

	row, err := NewUser_Selector().UserNameLower_Eq(UserNameLower).GetRow(base.DB)
	if err == nil {
		RowCacheIndex.Set("User_UserNameLower:"+fmt.Sprintf("%v", row.UserNameLower), row, 0)
		return row, true
	}

	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) User_ByUserNameLower_JustCache(UserNameLower string) (*User, bool) {
	o, ok := RowCacheIndex.Get("User_UserNameLower:" + fmt.Sprintf("%v", UserNameLower))
	if ok {
		if obj, ok := o.(*User); ok {
			return obj, true
		}
	}

	XOLogErr(errors.New("_JustCache is empty for secondry index " + "User_UserNameLower:" + fmt.Sprintf("%v", UserNameLower)))
	return nil, false
}

func (c _StoreImpl) PreLoadUser_ByUserNameLowers(UserNameLowers []string) {
	not_cached := make([]string, 0, len(UserNameLowers))

	for _, id := range UserNameLowers {
		_, ok := RowCacheIndex.Get("User_UserNameLower:" + fmt.Sprintf("%v", id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		rows, err := NewUser_Selector().UserNameLower_In(not_cached).GetRows(base.DB)
		if err == nil {
			for _, row := range rows {
				RowCacheIndex.Set("User_UserNameLower:"+fmt.Sprintf("%v", row.UserNameLower), row, 0)
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

// UserRelation - PRIMARY

// Chat - PRIMARY

//field//field//field

///// Generated from index 'RoomKey'.
func (c _StoreImpl) Chat_ByRoomKey(RoomKey string) (*Chat, bool) {
	o, ok := RowCacheIndex.Get("Chat_RoomKey:" + fmt.Sprintf("%v", RoomKey))
	if ok {
		if obj, ok := o.(*Chat); ok {
			return obj, true
		}
	}

	row, err := NewChat_Selector().RoomKey_Eq(RoomKey).GetRow(base.DB)
	if err == nil {
		RowCacheIndex.Set("Chat_RoomKey:"+fmt.Sprintf("%v", row.RoomKey), row, 0)
		return row, true
	}

	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) Chat_ByRoomKey_JustCache(RoomKey string) (*Chat, bool) {
	o, ok := RowCacheIndex.Get("Chat_RoomKey:" + fmt.Sprintf("%v", RoomKey))
	if ok {
		if obj, ok := o.(*Chat); ok {
			return obj, true
		}
	}

	XOLogErr(errors.New("_JustCache is empty for secondry index " + "Chat_RoomKey:" + fmt.Sprintf("%v", RoomKey)))
	return nil, false
}

func (c _StoreImpl) PreLoadChat_ByRoomKeys(RoomKeys []string) {
	not_cached := make([]string, 0, len(RoomKeys))

	for _, id := range RoomKeys {
		_, ok := RowCacheIndex.Get("Chat_RoomKey:" + fmt.Sprintf("%v", id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		rows, err := NewChat_Selector().RoomKey_In(not_cached).GetRows(base.DB)
		if err == nil {
			for _, row := range rows {
				RowCacheIndex.Set("Chat_RoomKey:"+fmt.Sprintf("%v", row.RoomKey), row, 0)
			}
		}
	}
}

// ChatDeleted - PRIMARY

// ChatDeleted - ChatId

// ChatLastMessage - PRIMARY

// ChatUserVersion - PRIMARY

// Group - PRIMARY

// GroupMember - PRIMARY

//field//field//field

///// Generated from index 'Id'.
func (c _StoreImpl) GroupMember_ByOrderId(OrderId int) (*GroupMember, bool) {
	o, ok := RowCacheIndex.Get("GroupMember_Id:" + fmt.Sprintf("%v", OrderId))
	if ok {
		if obj, ok := o.(*GroupMember); ok {
			return obj, true
		}
	}

	row, err := NewGroupMember_Selector().OrderId_Eq(OrderId).GetRow(base.DB)
	if err == nil {
		RowCacheIndex.Set("GroupMember_Id:"+fmt.Sprintf("%v", row.OrderId), row, 0)
		return row, true
	}

	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) GroupMember_ByOrderId_JustCache(OrderId int) (*GroupMember, bool) {
	o, ok := RowCacheIndex.Get("GroupMember_Id:" + fmt.Sprintf("%v", OrderId))
	if ok {
		if obj, ok := o.(*GroupMember); ok {
			return obj, true
		}
	}

	XOLogErr(errors.New("_JustCache is empty for secondry index " + "GroupMember_Id:" + fmt.Sprintf("%v", OrderId)))
	return nil, false
}

func (c _StoreImpl) PreLoadGroupMember_ByOrderIds(OrderIds []int) {
	not_cached := make([]int, 0, len(OrderIds))

	for _, id := range OrderIds {
		_, ok := RowCacheIndex.Get("GroupMember_Id:" + fmt.Sprintf("%v", id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		rows, err := NewGroupMember_Selector().OrderId_In(not_cached).GetRows(base.DB)
		if err == nil {
			for _, row := range rows {
				RowCacheIndex.Set("GroupMember_Id:"+fmt.Sprintf("%v", row.OrderId), row, 0)
			}
		}
	}
}

// GroupOrderdUser - PRIMARY

//field//field//field

///// Generated from index 'GroupId'.
func (c _StoreImpl) GroupOrderdUser_ByGroupId(GroupId int) (*GroupOrderdUser, bool) {
	o, ok := RowCacheIndex.Get("GroupOrderdUser_GroupId:" + fmt.Sprintf("%v", GroupId))
	if ok {
		if obj, ok := o.(*GroupOrderdUser); ok {
			return obj, true
		}
	}

	row, err := NewGroupOrderdUser_Selector().GroupId_Eq(GroupId).GetRow(base.DB)
	if err == nil {
		RowCacheIndex.Set("GroupOrderdUser_GroupId:"+fmt.Sprintf("%v", row.GroupId), row, 0)
		return row, true
	}

	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) GroupOrderdUser_ByGroupId_JustCache(GroupId int) (*GroupOrderdUser, bool) {
	o, ok := RowCacheIndex.Get("GroupOrderdUser_GroupId:" + fmt.Sprintf("%v", GroupId))
	if ok {
		if obj, ok := o.(*GroupOrderdUser); ok {
			return obj, true
		}
	}

	XOLogErr(errors.New("_JustCache is empty for secondry index " + "GroupOrderdUser_GroupId:" + fmt.Sprintf("%v", GroupId)))
	return nil, false
}

func (c _StoreImpl) PreLoadGroupOrderdUser_ByGroupIds(GroupIds []int) {
	not_cached := make([]int, 0, len(GroupIds))

	for _, id := range GroupIds {
		_, ok := RowCacheIndex.Get("GroupOrderdUser_GroupId:" + fmt.Sprintf("%v", id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		rows, err := NewGroupOrderdUser_Selector().GroupId_In(not_cached).GetRows(base.DB)
		if err == nil {
			for _, row := range rows {
				RowCacheIndex.Set("GroupOrderdUser_GroupId:"+fmt.Sprintf("%v", row.GroupId), row, 0)
			}
		}
	}
}

// GroupPinedMsg - PRIMARY

// FileMsg - PRIMARY

// FilePost - PRIMARY

// ActionFanout - PRIMARY

// ActionFanout - ForUserId

//field//field//field

///// Generated from index 'ActionId'.
func (c _StoreImpl) ActionFanout_ByActionId(ActionId int) (*ActionFanout, bool) {
	o, ok := RowCacheIndex.Get("ActionFanout_ActionId:" + fmt.Sprintf("%v", ActionId))
	if ok {
		if obj, ok := o.(*ActionFanout); ok {
			return obj, true
		}
	}

	row, err := NewActionFanout_Selector().ActionId_Eq(ActionId).GetRow(base.DB)
	if err == nil {
		RowCacheIndex.Set("ActionFanout_ActionId:"+fmt.Sprintf("%v", row.ActionId), row, 0)
		return row, true
	}

	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) ActionFanout_ByActionId_JustCache(ActionId int) (*ActionFanout, bool) {
	o, ok := RowCacheIndex.Get("ActionFanout_ActionId:" + fmt.Sprintf("%v", ActionId))
	if ok {
		if obj, ok := o.(*ActionFanout); ok {
			return obj, true
		}
	}

	XOLogErr(errors.New("_JustCache is empty for secondry index " + "ActionFanout_ActionId:" + fmt.Sprintf("%v", ActionId)))
	return nil, false
}

func (c _StoreImpl) PreLoadActionFanout_ByActionIds(ActionIds []int) {
	not_cached := make([]int, 0, len(ActionIds))

	for _, id := range ActionIds {
		_, ok := RowCacheIndex.Get("ActionFanout_ActionId:" + fmt.Sprintf("%v", id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		rows, err := NewActionFanout_Selector().ActionId_In(not_cached).GetRows(base.DB)
		if err == nil {
			for _, row := range rows {
				RowCacheIndex.Set("ActionFanout_ActionId:"+fmt.Sprintf("%v", row.ActionId), row, 0)
			}
		}
	}
}

// ActionFanout - ForUserId_2

// HomeFanout - PRIMARY

// HomeFanout - ForUserId_2

//field//field//field

///// Generated from index 'PostId'.
func (c _StoreImpl) HomeFanout_ByPostId(PostId int) (*HomeFanout, bool) {
	o, ok := RowCacheIndex.Get("HomeFanout_PostId:" + fmt.Sprintf("%v", PostId))
	if ok {
		if obj, ok := o.(*HomeFanout); ok {
			return obj, true
		}
	}

	row, err := NewHomeFanout_Selector().PostId_Eq(PostId).GetRow(base.DB)
	if err == nil {
		RowCacheIndex.Set("HomeFanout_PostId:"+fmt.Sprintf("%v", row.PostId), row, 0)
		return row, true
	}

	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) HomeFanout_ByPostId_JustCache(PostId int) (*HomeFanout, bool) {
	o, ok := RowCacheIndex.Get("HomeFanout_PostId:" + fmt.Sprintf("%v", PostId))
	if ok {
		if obj, ok := o.(*HomeFanout); ok {
			return obj, true
		}
	}

	XOLogErr(errors.New("_JustCache is empty for secondry index " + "HomeFanout_PostId:" + fmt.Sprintf("%v", PostId)))
	return nil, false
}

func (c _StoreImpl) PreLoadHomeFanout_ByPostIds(PostIds []int) {
	not_cached := make([]int, 0, len(PostIds))

	for _, id := range PostIds {
		_, ok := RowCacheIndex.Get("HomeFanout_PostId:" + fmt.Sprintf("%v", id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		rows, err := NewHomeFanout_Selector().PostId_In(not_cached).GetRows(base.DB)
		if err == nil {
			for _, row := range rows {
				RowCacheIndex.Set("HomeFanout_PostId:"+fmt.Sprintf("%v", row.PostId), row, 0)
			}
		}
	}
}

//field//field//field

///// Generated from index 'ForUserId'.
func (c _StoreImpl) HomeFanout_ByForUserId(ForUserId int) (*HomeFanout, bool) {
	o, ok := RowCacheIndex.Get("HomeFanout_ForUserId:" + fmt.Sprintf("%v", ForUserId))
	if ok {
		if obj, ok := o.(*HomeFanout); ok {
			return obj, true
		}
	}

	row, err := NewHomeFanout_Selector().ForUserId_Eq(ForUserId).GetRow(base.DB)
	if err == nil {
		RowCacheIndex.Set("HomeFanout_ForUserId:"+fmt.Sprintf("%v", row.ForUserId), row, 0)
		return row, true
	}

	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) HomeFanout_ByForUserId_JustCache(ForUserId int) (*HomeFanout, bool) {
	o, ok := RowCacheIndex.Get("HomeFanout_ForUserId:" + fmt.Sprintf("%v", ForUserId))
	if ok {
		if obj, ok := o.(*HomeFanout); ok {
			return obj, true
		}
	}

	XOLogErr(errors.New("_JustCache is empty for secondry index " + "HomeFanout_ForUserId:" + fmt.Sprintf("%v", ForUserId)))
	return nil, false
}

func (c _StoreImpl) PreLoadHomeFanout_ByForUserIds(ForUserIds []int) {
	not_cached := make([]int, 0, len(ForUserIds))

	for _, id := range ForUserIds {
		_, ok := RowCacheIndex.Get("HomeFanout_ForUserId:" + fmt.Sprintf("%v", id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		rows, err := NewHomeFanout_Selector().ForUserId_In(not_cached).GetRows(base.DB)
		if err == nil {
			for _, row := range rows {
				RowCacheIndex.Set("HomeFanout_ForUserId:"+fmt.Sprintf("%v", row.ForUserId), row, 0)
			}
		}
	}
}

// SuggestedTopPosts - PRIMARY

// SuggestedUser - PRIMARY

// PushChat - PRIMARY

//field//field//field

///// Generated from index 'Seq'.
func (c _StoreImpl) PushChat_BySeq(Seq int) (*PushChat, bool) {
	o, ok := RowCacheIndex.Get("PushChat_Seq:" + fmt.Sprintf("%v", Seq))
	if ok {
		if obj, ok := o.(*PushChat); ok {
			return obj, true
		}
	}

	row, err := NewPushChat_Selector().Seq_Eq(Seq).GetRow(base.DB)
	if err == nil {
		RowCacheIndex.Set("PushChat_Seq:"+fmt.Sprintf("%v", row.Seq), row, 0)
		return row, true
	}

	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PushChat_BySeq_JustCache(Seq int) (*PushChat, bool) {
	o, ok := RowCacheIndex.Get("PushChat_Seq:" + fmt.Sprintf("%v", Seq))
	if ok {
		if obj, ok := o.(*PushChat); ok {
			return obj, true
		}
	}

	XOLogErr(errors.New("_JustCache is empty for secondry index " + "PushChat_Seq:" + fmt.Sprintf("%v", Seq)))
	return nil, false
}

func (c _StoreImpl) PreLoadPushChat_BySeqs(Seqs []int) {
	not_cached := make([]int, 0, len(Seqs))

	for _, id := range Seqs {
		_, ok := RowCacheIndex.Get("PushChat_Seq:" + fmt.Sprintf("%v", id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		rows, err := NewPushChat_Selector().Seq_In(not_cached).GetRows(base.DB)
		if err == nil {
			for _, row := range rows {
				RowCacheIndex.Set("PushChat_Seq:"+fmt.Sprintf("%v", row.Seq), row, 0)
			}
		}
	}
}

// HTTPRPCLog - PRIMARY

// MetricLog - PRIMARY

// XfileServiceInfoLog - PRIMARY

// XfileServiceMetricLog - PRIMARY

// XfileServiceRequestLog - PRIMARY

// InvalidateCache - PRIMARY
