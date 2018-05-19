package x

import (
	"strconv"
	"time"
)

//Action Events

func OnAction_AfterInsert(row *Action) {
	RowCache.Set("Action:"+strconv.Itoa(row.ActionId), row, time.Hour*0)
}

func OnAction_AfterUpdate(row *Action) {
	RowCache.Set("Action:"+strconv.Itoa(row.ActionId), row, time.Hour*0)
}

func OnAction_AfterDelete(row *Action) {
	RowCache.Delete("Action:" + strconv.Itoa(row.ActionId))
}

func OnAction_LoadOne(row *Action) {
	RowCache.Set("Action:"+strconv.Itoa(row.ActionId), row, time.Hour*0)
}

func OnAction_LoadMany(rows []*Action) {
	for _, row := range rows {
		RowCache.Set("Action:"+strconv.Itoa(row.ActionId), row, time.Hour*0)
	}
}

//Comment Events

func OnComment_AfterInsert(row *Comment) {
	RowCache.Set("Comment:"+strconv.Itoa(row.CommentId), row, time.Hour*0)
}

func OnComment_AfterUpdate(row *Comment) {
	RowCache.Set("Comment:"+strconv.Itoa(row.CommentId), row, time.Hour*0)
}

func OnComment_AfterDelete(row *Comment) {
	RowCache.Delete("Comment:" + strconv.Itoa(row.CommentId))
}

func OnComment_LoadOne(row *Comment) {
	RowCache.Set("Comment:"+strconv.Itoa(row.CommentId), row, time.Hour*0)
}

func OnComment_LoadMany(rows []*Comment) {
	for _, row := range rows {
		RowCache.Set("Comment:"+strconv.Itoa(row.CommentId), row, time.Hour*0)
	}
}

//CommentDeleted Events

func OnCommentDeleted_AfterInsert(row *CommentDeleted) {
	RowCache.Set("CommentDeleted:"+strconv.Itoa(row.CommentId), row, time.Hour*0)
}

func OnCommentDeleted_AfterUpdate(row *CommentDeleted) {
	RowCache.Set("CommentDeleted:"+strconv.Itoa(row.CommentId), row, time.Hour*0)
}

func OnCommentDeleted_AfterDelete(row *CommentDeleted) {
	RowCache.Delete("CommentDeleted:" + strconv.Itoa(row.CommentId))
}

func OnCommentDeleted_LoadOne(row *CommentDeleted) {
	RowCache.Set("CommentDeleted:"+strconv.Itoa(row.CommentId), row, time.Hour*0)
}

func OnCommentDeleted_LoadMany(rows []*CommentDeleted) {
	for _, row := range rows {
		RowCache.Set("CommentDeleted:"+strconv.Itoa(row.CommentId), row, time.Hour*0)
	}
}

//Event Events

func OnEvent_AfterInsert(row *Event) {
	RowCache.Set("Event:"+strconv.Itoa(row.EventId), row, time.Hour*0)
}

func OnEvent_AfterUpdate(row *Event) {
	RowCache.Set("Event:"+strconv.Itoa(row.EventId), row, time.Hour*0)
}

func OnEvent_AfterDelete(row *Event) {
	RowCache.Delete("Event:" + strconv.Itoa(row.EventId))
}

func OnEvent_LoadOne(row *Event) {
	RowCache.Set("Event:"+strconv.Itoa(row.EventId), row, time.Hour*0)
}

func OnEvent_LoadMany(rows []*Event) {
	for _, row := range rows {
		RowCache.Set("Event:"+strconv.Itoa(row.EventId), row, time.Hour*0)
	}
}

//FollowingList Events

func OnFollowingList_AfterInsert(row *FollowingList) {
	RowCache.Set("FollowingList:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnFollowingList_AfterUpdate(row *FollowingList) {
	RowCache.Set("FollowingList:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnFollowingList_AfterDelete(row *FollowingList) {
	RowCache.Delete("FollowingList:" + strconv.Itoa(row.Id))
}

func OnFollowingList_LoadOne(row *FollowingList) {
	RowCache.Set("FollowingList:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnFollowingList_LoadMany(rows []*FollowingList) {
	for _, row := range rows {
		RowCache.Set("FollowingList:"+strconv.Itoa(row.Id), row, time.Hour*0)
	}
}

//FollowingListMember Events

func OnFollowingListMember_AfterInsert(row *FollowingListMember) {
	RowCache.Set("FollowingListMember:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnFollowingListMember_AfterUpdate(row *FollowingListMember) {
	RowCache.Set("FollowingListMember:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnFollowingListMember_AfterDelete(row *FollowingListMember) {
	RowCache.Delete("FollowingListMember:" + strconv.Itoa(row.Id))
}

func OnFollowingListMember_LoadOne(row *FollowingListMember) {
	RowCache.Set("FollowingListMember:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnFollowingListMember_LoadMany(rows []*FollowingListMember) {
	for _, row := range rows {
		RowCache.Set("FollowingListMember:"+strconv.Itoa(row.Id), row, time.Hour*0)
	}
}

//FollowingListMemberRemoved Events

func OnFollowingListMemberRemoved_AfterInsert(row *FollowingListMemberRemoved) {
	RowCache.Set("FollowingListMemberRemoved:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnFollowingListMemberRemoved_AfterUpdate(row *FollowingListMemberRemoved) {
	RowCache.Set("FollowingListMemberRemoved:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnFollowingListMemberRemoved_AfterDelete(row *FollowingListMemberRemoved) {
	RowCache.Delete("FollowingListMemberRemoved:" + strconv.Itoa(row.Id))
}

func OnFollowingListMemberRemoved_LoadOne(row *FollowingListMemberRemoved) {
	RowCache.Set("FollowingListMemberRemoved:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnFollowingListMemberRemoved_LoadMany(rows []*FollowingListMemberRemoved) {
	for _, row := range rows {
		RowCache.Set("FollowingListMemberRemoved:"+strconv.Itoa(row.Id), row, time.Hour*0)
	}
}

//Like Events

func OnLike_AfterInsert(row *Like) {
	RowCache.Set("Like:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnLike_AfterUpdate(row *Like) {
	RowCache.Set("Like:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnLike_AfterDelete(row *Like) {
	RowCache.Delete("Like:" + strconv.Itoa(row.Id))
}

func OnLike_LoadOne(row *Like) {
	RowCache.Set("Like:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnLike_LoadMany(rows []*Like) {
	for _, row := range rows {
		RowCache.Set("Like:"+strconv.Itoa(row.Id), row, time.Hour*0)
	}
}

//Notify Events

func OnNotify_AfterInsert(row *Notify) {
	RowCache.Set("Notify:"+strconv.Itoa(row.NotifyId), row, time.Hour*0)
}

func OnNotify_AfterUpdate(row *Notify) {
	RowCache.Set("Notify:"+strconv.Itoa(row.NotifyId), row, time.Hour*0)
}

func OnNotify_AfterDelete(row *Notify) {
	RowCache.Delete("Notify:" + strconv.Itoa(row.NotifyId))
}

func OnNotify_LoadOne(row *Notify) {
	RowCache.Set("Notify:"+strconv.Itoa(row.NotifyId), row, time.Hour*0)
}

func OnNotify_LoadMany(rows []*Notify) {
	for _, row := range rows {
		RowCache.Set("Notify:"+strconv.Itoa(row.NotifyId), row, time.Hour*0)
	}
}

//NotifyRemoved Events

func OnNotifyRemoved_AfterInsert(row *NotifyRemoved) {
	RowCache.Set("NotifyRemoved:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnNotifyRemoved_AfterUpdate(row *NotifyRemoved) {
	RowCache.Set("NotifyRemoved:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnNotifyRemoved_AfterDelete(row *NotifyRemoved) {
	RowCache.Delete("NotifyRemoved:" + strconv.Itoa(row.Id))
}

func OnNotifyRemoved_LoadOne(row *NotifyRemoved) {
	RowCache.Set("NotifyRemoved:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnNotifyRemoved_LoadMany(rows []*NotifyRemoved) {
	for _, row := range rows {
		RowCache.Set("NotifyRemoved:"+strconv.Itoa(row.Id), row, time.Hour*0)
	}
}

//PhoneContact Events

func OnPhoneContact_AfterInsert(row *PhoneContact) {
	RowCache.Set("PhoneContact:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnPhoneContact_AfterUpdate(row *PhoneContact) {
	RowCache.Set("PhoneContact:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnPhoneContact_AfterDelete(row *PhoneContact) {
	RowCache.Delete("PhoneContact:" + strconv.Itoa(row.Id))
}

func OnPhoneContact_LoadOne(row *PhoneContact) {
	RowCache.Set("PhoneContact:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnPhoneContact_LoadMany(rows []*PhoneContact) {
	for _, row := range rows {
		RowCache.Set("PhoneContact:"+strconv.Itoa(row.Id), row, time.Hour*0)
	}
}

//Post Events

func OnPost_AfterInsert(row *Post) {
	RowCache.Set("Post:"+strconv.Itoa(row.PostId), row, time.Hour*0)
}

func OnPost_AfterUpdate(row *Post) {
	RowCache.Set("Post:"+strconv.Itoa(row.PostId), row, time.Hour*0)
}

func OnPost_AfterDelete(row *Post) {
	RowCache.Delete("Post:" + strconv.Itoa(row.PostId))
}

func OnPost_LoadOne(row *Post) {
	RowCache.Set("Post:"+strconv.Itoa(row.PostId), row, time.Hour*0)
}

func OnPost_LoadMany(rows []*Post) {
	for _, row := range rows {
		RowCache.Set("Post:"+strconv.Itoa(row.PostId), row, time.Hour*0)
	}
}

//PostCopy Events

func OnPostCopy_AfterInsert(row *PostCopy) {
	RowCache.Set("PostCopy:"+strconv.Itoa(row.PostId), row, time.Hour*0)
}

func OnPostCopy_AfterUpdate(row *PostCopy) {
	RowCache.Set("PostCopy:"+strconv.Itoa(row.PostId), row, time.Hour*0)
}

func OnPostCopy_AfterDelete(row *PostCopy) {
	RowCache.Delete("PostCopy:" + strconv.Itoa(row.PostId))
}

func OnPostCopy_LoadOne(row *PostCopy) {
	RowCache.Set("PostCopy:"+strconv.Itoa(row.PostId), row, time.Hour*0)
}

func OnPostCopy_LoadMany(rows []*PostCopy) {
	for _, row := range rows {
		RowCache.Set("PostCopy:"+strconv.Itoa(row.PostId), row, time.Hour*0)
	}
}

//PostCount Events

func OnPostCount_AfterInsert(row *PostCount) {
	RowCache.Set("PostCount:"+strconv.Itoa(row.PostId), row, time.Hour*0)
}

func OnPostCount_AfterUpdate(row *PostCount) {
	RowCache.Set("PostCount:"+strconv.Itoa(row.PostId), row, time.Hour*0)
}

func OnPostCount_AfterDelete(row *PostCount) {
	RowCache.Delete("PostCount:" + strconv.Itoa(row.PostId))
}

func OnPostCount_LoadOne(row *PostCount) {
	RowCache.Set("PostCount:"+strconv.Itoa(row.PostId), row, time.Hour*0)
}

func OnPostCount_LoadMany(rows []*PostCount) {
	for _, row := range rows {
		RowCache.Set("PostCount:"+strconv.Itoa(row.PostId), row, time.Hour*0)
	}
}

//PostDeleted Events

func OnPostDeleted_AfterInsert(row *PostDeleted) {
	RowCache.Set("PostDeleted:"+strconv.Itoa(row.PostId), row, time.Hour*0)
}

func OnPostDeleted_AfterUpdate(row *PostDeleted) {
	RowCache.Set("PostDeleted:"+strconv.Itoa(row.PostId), row, time.Hour*0)
}

func OnPostDeleted_AfterDelete(row *PostDeleted) {
	RowCache.Delete("PostDeleted:" + strconv.Itoa(row.PostId))
}

func OnPostDeleted_LoadOne(row *PostDeleted) {
	RowCache.Set("PostDeleted:"+strconv.Itoa(row.PostId), row, time.Hour*0)
}

func OnPostDeleted_LoadMany(rows []*PostDeleted) {
	for _, row := range rows {
		RowCache.Set("PostDeleted:"+strconv.Itoa(row.PostId), row, time.Hour*0)
	}
}

//PostKey Events

func OnPostKey_AfterInsert(row *PostKey) {
	RowCache.Set("PostKey:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnPostKey_AfterUpdate(row *PostKey) {
	RowCache.Set("PostKey:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnPostKey_AfterDelete(row *PostKey) {
	RowCache.Delete("PostKey:" + strconv.Itoa(row.Id))
}

func OnPostKey_LoadOne(row *PostKey) {
	RowCache.Set("PostKey:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnPostKey_LoadMany(rows []*PostKey) {
	for _, row := range rows {
		RowCache.Set("PostKey:"+strconv.Itoa(row.Id), row, time.Hour*0)
	}
}

//PostLink Events

func OnPostLink_AfterInsert(row *PostLink) {
	RowCache.Set("PostLink:"+strconv.Itoa(row.LinkId), row, time.Hour*0)
}

func OnPostLink_AfterUpdate(row *PostLink) {
	RowCache.Set("PostLink:"+strconv.Itoa(row.LinkId), row, time.Hour*0)
}

func OnPostLink_AfterDelete(row *PostLink) {
	RowCache.Delete("PostLink:" + strconv.Itoa(row.LinkId))
}

func OnPostLink_LoadOne(row *PostLink) {
	RowCache.Set("PostLink:"+strconv.Itoa(row.LinkId), row, time.Hour*0)
}

func OnPostLink_LoadMany(rows []*PostLink) {
	for _, row := range rows {
		RowCache.Set("PostLink:"+strconv.Itoa(row.LinkId), row, time.Hour*0)
	}
}

//PostMedia Events

func OnPostMedia_AfterInsert(row *PostMedia) {
	RowCache.Set("PostMedia:"+strconv.Itoa(row.MediaId), row, time.Hour*0)
}

func OnPostMedia_AfterUpdate(row *PostMedia) {
	RowCache.Set("PostMedia:"+strconv.Itoa(row.MediaId), row, time.Hour*0)
}

func OnPostMedia_AfterDelete(row *PostMedia) {
	RowCache.Delete("PostMedia:" + strconv.Itoa(row.MediaId))
}

func OnPostMedia_LoadOne(row *PostMedia) {
	RowCache.Set("PostMedia:"+strconv.Itoa(row.MediaId), row, time.Hour*0)
}

func OnPostMedia_LoadMany(rows []*PostMedia) {
	for _, row := range rows {
		RowCache.Set("PostMedia:"+strconv.Itoa(row.MediaId), row, time.Hour*0)
	}
}

//PostMentioned Events

func OnPostMentioned_AfterInsert(row *PostMentioned) {
	RowCache.Set("PostMentioned:"+strconv.Itoa(row.MentionedId), row, time.Hour*0)
}

func OnPostMentioned_AfterUpdate(row *PostMentioned) {
	RowCache.Set("PostMentioned:"+strconv.Itoa(row.MentionedId), row, time.Hour*0)
}

func OnPostMentioned_AfterDelete(row *PostMentioned) {
	RowCache.Delete("PostMentioned:" + strconv.Itoa(row.MentionedId))
}

func OnPostMentioned_LoadOne(row *PostMentioned) {
	RowCache.Set("PostMentioned:"+strconv.Itoa(row.MentionedId), row, time.Hour*0)
}

func OnPostMentioned_LoadMany(rows []*PostMentioned) {
	for _, row := range rows {
		RowCache.Set("PostMentioned:"+strconv.Itoa(row.MentionedId), row, time.Hour*0)
	}
}

//PostReshared Events

func OnPostReshared_AfterInsert(row *PostReshared) {
	RowCache.Set("PostReshared:"+strconv.Itoa(row.ResharedId), row, time.Hour*0)
}

func OnPostReshared_AfterUpdate(row *PostReshared) {
	RowCache.Set("PostReshared:"+strconv.Itoa(row.ResharedId), row, time.Hour*0)
}

func OnPostReshared_AfterDelete(row *PostReshared) {
	RowCache.Delete("PostReshared:" + strconv.Itoa(row.ResharedId))
}

func OnPostReshared_LoadOne(row *PostReshared) {
	RowCache.Set("PostReshared:"+strconv.Itoa(row.ResharedId), row, time.Hour*0)
}

func OnPostReshared_LoadMany(rows []*PostReshared) {
	for _, row := range rows {
		RowCache.Set("PostReshared:"+strconv.Itoa(row.ResharedId), row, time.Hour*0)
	}
}

//SearchClicked Events

func OnSearchClicked_AfterInsert(row *SearchClicked) {
	RowCache.Set("SearchClicked:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnSearchClicked_AfterUpdate(row *SearchClicked) {
	RowCache.Set("SearchClicked:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnSearchClicked_AfterDelete(row *SearchClicked) {
	RowCache.Delete("SearchClicked:" + strconv.Itoa(row.Id))
}

func OnSearchClicked_LoadOne(row *SearchClicked) {
	RowCache.Set("SearchClicked:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnSearchClicked_LoadMany(rows []*SearchClicked) {
	for _, row := range rows {
		RowCache.Set("SearchClicked:"+strconv.Itoa(row.Id), row, time.Hour*0)
	}
}

//Session Events

func OnSession_AfterInsert(row *Session) {
	RowCache.Set("Session:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnSession_AfterUpdate(row *Session) {
	RowCache.Set("Session:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnSession_AfterDelete(row *Session) {
	RowCache.Delete("Session:" + strconv.Itoa(row.Id))
}

func OnSession_LoadOne(row *Session) {
	RowCache.Set("Session:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnSession_LoadMany(rows []*Session) {
	for _, row := range rows {
		RowCache.Set("Session:"+strconv.Itoa(row.Id), row, time.Hour*0)
	}
}

//SettingClient Events

func OnSettingClient_AfterInsert(row *SettingClient) {
	RowCache.Set("SettingClient:"+strconv.Itoa(row.UserId), row, time.Hour*0)
}

func OnSettingClient_AfterUpdate(row *SettingClient) {
	RowCache.Set("SettingClient:"+strconv.Itoa(row.UserId), row, time.Hour*0)
}

func OnSettingClient_AfterDelete(row *SettingClient) {
	RowCache.Delete("SettingClient:" + strconv.Itoa(row.UserId))
}

func OnSettingClient_LoadOne(row *SettingClient) {
	RowCache.Set("SettingClient:"+strconv.Itoa(row.UserId), row, time.Hour*0)
}

func OnSettingClient_LoadMany(rows []*SettingClient) {
	for _, row := range rows {
		RowCache.Set("SettingClient:"+strconv.Itoa(row.UserId), row, time.Hour*0)
	}
}

//SettingNotification Events

func OnSettingNotification_AfterInsert(row *SettingNotification) {
	RowCache.Set("SettingNotification:"+strconv.Itoa(row.UserId), row, time.Hour*0)
}

func OnSettingNotification_AfterUpdate(row *SettingNotification) {
	RowCache.Set("SettingNotification:"+strconv.Itoa(row.UserId), row, time.Hour*0)
}

func OnSettingNotification_AfterDelete(row *SettingNotification) {
	RowCache.Delete("SettingNotification:" + strconv.Itoa(row.UserId))
}

func OnSettingNotification_LoadOne(row *SettingNotification) {
	RowCache.Set("SettingNotification:"+strconv.Itoa(row.UserId), row, time.Hour*0)
}

func OnSettingNotification_LoadMany(rows []*SettingNotification) {
	for _, row := range rows {
		RowCache.Set("SettingNotification:"+strconv.Itoa(row.UserId), row, time.Hour*0)
	}
}

//Tag Events

func OnTag_AfterInsert(row *Tag) {
	RowCache.Set("Tag:"+strconv.Itoa(row.TagId), row, time.Hour*0)
}

func OnTag_AfterUpdate(row *Tag) {
	RowCache.Set("Tag:"+strconv.Itoa(row.TagId), row, time.Hour*0)
}

func OnTag_AfterDelete(row *Tag) {
	RowCache.Delete("Tag:" + strconv.Itoa(row.TagId))
}

func OnTag_LoadOne(row *Tag) {
	RowCache.Set("Tag:"+strconv.Itoa(row.TagId), row, time.Hour*0)
}

func OnTag_LoadMany(rows []*Tag) {
	for _, row := range rows {
		RowCache.Set("Tag:"+strconv.Itoa(row.TagId), row, time.Hour*0)
	}
}

//TagPost Events

func OnTagPost_AfterInsert(row *TagPost) {
	RowCache.Set("TagPost:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnTagPost_AfterUpdate(row *TagPost) {
	RowCache.Set("TagPost:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnTagPost_AfterDelete(row *TagPost) {
	RowCache.Delete("TagPost:" + strconv.Itoa(row.Id))
}

func OnTagPost_LoadOne(row *TagPost) {
	RowCache.Set("TagPost:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnTagPost_LoadMany(rows []*TagPost) {
	for _, row := range rows {
		RowCache.Set("TagPost:"+strconv.Itoa(row.Id), row, time.Hour*0)
	}
}

//TriggerLog Events

func OnTriggerLog_AfterInsert(row *TriggerLog) {
	RowCache.Set("TriggerLog:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnTriggerLog_AfterUpdate(row *TriggerLog) {
	RowCache.Set("TriggerLog:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnTriggerLog_AfterDelete(row *TriggerLog) {
	RowCache.Delete("TriggerLog:" + strconv.Itoa(row.Id))
}

func OnTriggerLog_LoadOne(row *TriggerLog) {
	RowCache.Set("TriggerLog:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnTriggerLog_LoadMany(rows []*TriggerLog) {
	for _, row := range rows {
		RowCache.Set("TriggerLog:"+strconv.Itoa(row.Id), row, time.Hour*0)
	}
}

//User Events

func OnUser_AfterInsert(row *User) {
	RowCache.Set("User:"+strconv.Itoa(row.UserId), row, time.Hour*0)
}

func OnUser_AfterUpdate(row *User) {
	RowCache.Set("User:"+strconv.Itoa(row.UserId), row, time.Hour*0)
}

func OnUser_AfterDelete(row *User) {
	RowCache.Delete("User:" + strconv.Itoa(row.UserId))
}

func OnUser_LoadOne(row *User) {
	RowCache.Set("User:"+strconv.Itoa(row.UserId), row, time.Hour*0)
}

func OnUser_LoadMany(rows []*User) {
	for _, row := range rows {
		RowCache.Set("User:"+strconv.Itoa(row.UserId), row, time.Hour*0)
	}
}

//UserMetaInfo Events

func OnUserMetaInfo_AfterInsert(row *UserMetaInfo) {
	RowCache.Set("UserMetaInfo:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnUserMetaInfo_AfterUpdate(row *UserMetaInfo) {
	RowCache.Set("UserMetaInfo:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnUserMetaInfo_AfterDelete(row *UserMetaInfo) {
	RowCache.Delete("UserMetaInfo:" + strconv.Itoa(row.Id))
}

func OnUserMetaInfo_LoadOne(row *UserMetaInfo) {
	RowCache.Set("UserMetaInfo:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnUserMetaInfo_LoadMany(rows []*UserMetaInfo) {
	for _, row := range rows {
		RowCache.Set("UserMetaInfo:"+strconv.Itoa(row.Id), row, time.Hour*0)
	}
}

//UserPassword Events

func OnUserPassword_AfterInsert(row *UserPassword) {
	RowCache.Set("UserPassword:"+strconv.Itoa(row.UserId), row, time.Hour*0)
}

func OnUserPassword_AfterUpdate(row *UserPassword) {
	RowCache.Set("UserPassword:"+strconv.Itoa(row.UserId), row, time.Hour*0)
}

func OnUserPassword_AfterDelete(row *UserPassword) {
	RowCache.Delete("UserPassword:" + strconv.Itoa(row.UserId))
}

func OnUserPassword_LoadOne(row *UserPassword) {
	RowCache.Set("UserPassword:"+strconv.Itoa(row.UserId), row, time.Hour*0)
}

func OnUserPassword_LoadMany(rows []*UserPassword) {
	for _, row := range rows {
		RowCache.Set("UserPassword:"+strconv.Itoa(row.UserId), row, time.Hour*0)
	}
}

//Chat Events

func OnChat_AfterInsert(row *Chat) {
	RowCache.Set("Chat:"+row.ChatKey, row, time.Hour*0)
}

func OnChat_AfterUpdate(row *Chat) {
	RowCache.Set("Chat:"+row.ChatKey, row, time.Hour*0)
}

func OnChat_AfterDelete(row *Chat) {
	RowCache.Delete("Chat:" + row.ChatKey)
}

func OnChat_LoadOne(row *Chat) {
	RowCache.Set("Chat:"+row.ChatKey, row, time.Hour*0)
}

func OnChat_LoadMany(rows []*Chat) {
	for _, row := range rows {
		RowCache.Set("Chat:"+row.ChatKey, row, time.Hour*0)
	}
}

//ChatLastMessage Events

func OnChatLastMessage_AfterInsert(row *ChatLastMessage) {
	RowCache.Set("ChatLastMessage:"+row.ChatKey, row, time.Hour*0)
}

func OnChatLastMessage_AfterUpdate(row *ChatLastMessage) {
	RowCache.Set("ChatLastMessage:"+row.ChatKey, row, time.Hour*0)
}

func OnChatLastMessage_AfterDelete(row *ChatLastMessage) {
	RowCache.Delete("ChatLastMessage:" + row.ChatKey)
}

func OnChatLastMessage_LoadOne(row *ChatLastMessage) {
	RowCache.Set("ChatLastMessage:"+row.ChatKey, row, time.Hour*0)
}

func OnChatLastMessage_LoadMany(rows []*ChatLastMessage) {
	for _, row := range rows {
		RowCache.Set("ChatLastMessage:"+row.ChatKey, row, time.Hour*0)
	}
}

//DirectMessage Events

func OnDirectMessage_AfterInsert(row *DirectMessage) {
	RowCache.Set("DirectMessage:"+strconv.Itoa(row.MessageId), row, time.Hour*0)
}

func OnDirectMessage_AfterUpdate(row *DirectMessage) {
	RowCache.Set("DirectMessage:"+strconv.Itoa(row.MessageId), row, time.Hour*0)
}

func OnDirectMessage_AfterDelete(row *DirectMessage) {
	RowCache.Delete("DirectMessage:" + strconv.Itoa(row.MessageId))
}

func OnDirectMessage_LoadOne(row *DirectMessage) {
	RowCache.Set("DirectMessage:"+strconv.Itoa(row.MessageId), row, time.Hour*0)
}

func OnDirectMessage_LoadMany(rows []*DirectMessage) {
	for _, row := range rows {
		RowCache.Set("DirectMessage:"+strconv.Itoa(row.MessageId), row, time.Hour*0)
	}
}

//Group Events

func OnGroup_AfterInsert(row *Group) {
	RowCache.Set("Group:"+strconv.Itoa(row.GroupId), row, time.Hour*0)
}

func OnGroup_AfterUpdate(row *Group) {
	RowCache.Set("Group:"+strconv.Itoa(row.GroupId), row, time.Hour*0)
}

func OnGroup_AfterDelete(row *Group) {
	RowCache.Delete("Group:" + strconv.Itoa(row.GroupId))
}

func OnGroup_LoadOne(row *Group) {
	RowCache.Set("Group:"+strconv.Itoa(row.GroupId), row, time.Hour*0)
}

func OnGroup_LoadMany(rows []*Group) {
	for _, row := range rows {
		RowCache.Set("Group:"+strconv.Itoa(row.GroupId), row, time.Hour*0)
	}
}

//GroupMember Events

func OnGroupMember_AfterInsert(row *GroupMember) {
	RowCache.Set("GroupMember:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnGroupMember_AfterUpdate(row *GroupMember) {
	RowCache.Set("GroupMember:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnGroupMember_AfterDelete(row *GroupMember) {
	RowCache.Delete("GroupMember:" + strconv.Itoa(row.Id))
}

func OnGroupMember_LoadOne(row *GroupMember) {
	RowCache.Set("GroupMember:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnGroupMember_LoadMany(rows []*GroupMember) {
	for _, row := range rows {
		RowCache.Set("GroupMember:"+strconv.Itoa(row.Id), row, time.Hour*0)
	}
}

//GroupMessage Events

func OnGroupMessage_AfterInsert(row *GroupMessage) {
	RowCache.Set("GroupMessage:"+strconv.Itoa(row.MessageId), row, time.Hour*0)
}

func OnGroupMessage_AfterUpdate(row *GroupMessage) {
	RowCache.Set("GroupMessage:"+strconv.Itoa(row.MessageId), row, time.Hour*0)
}

func OnGroupMessage_AfterDelete(row *GroupMessage) {
	RowCache.Delete("GroupMessage:" + strconv.Itoa(row.MessageId))
}

func OnGroupMessage_LoadOne(row *GroupMessage) {
	RowCache.Set("GroupMessage:"+strconv.Itoa(row.MessageId), row, time.Hour*0)
}

func OnGroupMessage_LoadMany(rows []*GroupMessage) {
	for _, row := range rows {
		RowCache.Set("GroupMessage:"+strconv.Itoa(row.MessageId), row, time.Hour*0)
	}
}

//FileMsg Events

func OnFileMsg_AfterInsert(row *FileMsg) {
	RowCache.Set("FileMsg:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnFileMsg_AfterUpdate(row *FileMsg) {
	RowCache.Set("FileMsg:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnFileMsg_AfterDelete(row *FileMsg) {
	RowCache.Delete("FileMsg:" + strconv.Itoa(row.Id))
}

func OnFileMsg_LoadOne(row *FileMsg) {
	RowCache.Set("FileMsg:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnFileMsg_LoadMany(rows []*FileMsg) {
	for _, row := range rows {
		RowCache.Set("FileMsg:"+strconv.Itoa(row.Id), row, time.Hour*0)
	}
}

//FilePost Events

func OnFilePost_AfterInsert(row *FilePost) {
	RowCache.Set("FilePost:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnFilePost_AfterUpdate(row *FilePost) {
	RowCache.Set("FilePost:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnFilePost_AfterDelete(row *FilePost) {
	RowCache.Delete("FilePost:" + strconv.Itoa(row.Id))
}

func OnFilePost_LoadOne(row *FilePost) {
	RowCache.Set("FilePost:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnFilePost_LoadMany(rows []*FilePost) {
	for _, row := range rows {
		RowCache.Set("FilePost:"+strconv.Itoa(row.Id), row, time.Hour*0)
	}
}

//ActionFanout Events

func OnActionFanout_AfterInsert(row *ActionFanout) {
	RowCache.Set("ActionFanout:"+strconv.Itoa(row.OrderId), row, time.Hour*0)
}

func OnActionFanout_AfterUpdate(row *ActionFanout) {
	RowCache.Set("ActionFanout:"+strconv.Itoa(row.OrderId), row, time.Hour*0)
}

func OnActionFanout_AfterDelete(row *ActionFanout) {
	RowCache.Delete("ActionFanout:" + strconv.Itoa(row.OrderId))
}

func OnActionFanout_LoadOne(row *ActionFanout) {
	RowCache.Set("ActionFanout:"+strconv.Itoa(row.OrderId), row, time.Hour*0)
}

func OnActionFanout_LoadMany(rows []*ActionFanout) {
	for _, row := range rows {
		RowCache.Set("ActionFanout:"+strconv.Itoa(row.OrderId), row, time.Hour*0)
	}
}

//HomeFanout Events

func OnHomeFanout_AfterInsert(row *HomeFanout) {
	RowCache.Set("HomeFanout:"+strconv.Itoa(row.OrderId), row, time.Hour*0)
}

func OnHomeFanout_AfterUpdate(row *HomeFanout) {
	RowCache.Set("HomeFanout:"+strconv.Itoa(row.OrderId), row, time.Hour*0)
}

func OnHomeFanout_AfterDelete(row *HomeFanout) {
	RowCache.Delete("HomeFanout:" + strconv.Itoa(row.OrderId))
}

func OnHomeFanout_LoadOne(row *HomeFanout) {
	RowCache.Set("HomeFanout:"+strconv.Itoa(row.OrderId), row, time.Hour*0)
}

func OnHomeFanout_LoadMany(rows []*HomeFanout) {
	for _, row := range rows {
		RowCache.Set("HomeFanout:"+strconv.Itoa(row.OrderId), row, time.Hour*0)
	}
}

//SuggestedTopPost Events

func OnSuggestedTopPost_AfterInsert(row *SuggestedTopPost) {
	RowCache.Set("SuggestedTopPost:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnSuggestedTopPost_AfterUpdate(row *SuggestedTopPost) {
	RowCache.Set("SuggestedTopPost:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnSuggestedTopPost_AfterDelete(row *SuggestedTopPost) {
	RowCache.Delete("SuggestedTopPost:" + strconv.Itoa(row.Id))
}

func OnSuggestedTopPost_LoadOne(row *SuggestedTopPost) {
	RowCache.Set("SuggestedTopPost:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnSuggestedTopPost_LoadMany(rows []*SuggestedTopPost) {
	for _, row := range rows {
		RowCache.Set("SuggestedTopPost:"+strconv.Itoa(row.Id), row, time.Hour*0)
	}
}

//SuggestedUser Events

func OnSuggestedUser_AfterInsert(row *SuggestedUser) {
	RowCache.Set("SuggestedUser:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnSuggestedUser_AfterUpdate(row *SuggestedUser) {
	RowCache.Set("SuggestedUser:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnSuggestedUser_AfterDelete(row *SuggestedUser) {
	RowCache.Delete("SuggestedUser:" + strconv.Itoa(row.Id))
}

func OnSuggestedUser_LoadOne(row *SuggestedUser) {
	RowCache.Set("SuggestedUser:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnSuggestedUser_LoadMany(rows []*SuggestedUser) {
	for _, row := range rows {
		RowCache.Set("SuggestedUser:"+strconv.Itoa(row.Id), row, time.Hour*0)
	}
}

//ChatSync2 Events

func OnChatSync2_AfterInsert(row *ChatSync2) {
	RowCache.Set("ChatSync2:"+strconv.Itoa(row.SyncId), row, time.Hour*0)
}

func OnChatSync2_AfterUpdate(row *ChatSync2) {
	RowCache.Set("ChatSync2:"+strconv.Itoa(row.SyncId), row, time.Hour*0)
}

func OnChatSync2_AfterDelete(row *ChatSync2) {
	RowCache.Delete("ChatSync2:" + strconv.Itoa(row.SyncId))
}

func OnChatSync2_LoadOne(row *ChatSync2) {
	RowCache.Set("ChatSync2:"+strconv.Itoa(row.SyncId), row, time.Hour*0)
}

func OnChatSync2_LoadMany(rows []*ChatSync2) {
	for _, row := range rows {
		RowCache.Set("ChatSync2:"+strconv.Itoa(row.SyncId), row, time.Hour*0)
	}
}

//LowerTable Events

func OnLowerTable_AfterInsert(row *LowerTable) {
	RowCache.Set("LowerTable:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnLowerTable_AfterUpdate(row *LowerTable) {
	RowCache.Set("LowerTable:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnLowerTable_AfterDelete(row *LowerTable) {
	RowCache.Delete("LowerTable:" + strconv.Itoa(row.Id))
}

func OnLowerTable_LoadOne(row *LowerTable) {
	RowCache.Set("LowerTable:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnLowerTable_LoadMany(rows []*LowerTable) {
	for _, row := range rows {
		RowCache.Set("LowerTable:"+strconv.Itoa(row.Id), row, time.Hour*0)
	}
}

//PushChat Events

func OnPushChat_AfterInsert(row *PushChat) {
	RowCache.Set("PushChat:"+strconv.Itoa(row.PushId), row, time.Hour*0)
}

func OnPushChat_AfterUpdate(row *PushChat) {
	RowCache.Set("PushChat:"+strconv.Itoa(row.PushId), row, time.Hour*0)
}

func OnPushChat_AfterDelete(row *PushChat) {
	RowCache.Delete("PushChat:" + strconv.Itoa(row.PushId))
}

func OnPushChat_LoadOne(row *PushChat) {
	RowCache.Set("PushChat:"+strconv.Itoa(row.PushId), row, time.Hour*0)
}

func OnPushChat_LoadMany(rows []*PushChat) {
	for _, row := range rows {
		RowCache.Set("PushChat:"+strconv.Itoa(row.PushId), row, time.Hour*0)
	}
}

//PushChat2 Events

func OnPushChat2_AfterInsert(row *PushChat2) {
	RowCache.Set("PushChat2:"+strconv.Itoa(row.SyncId), row, time.Hour*0)
}

func OnPushChat2_AfterUpdate(row *PushChat2) {
	RowCache.Set("PushChat2:"+strconv.Itoa(row.SyncId), row, time.Hour*0)
}

func OnPushChat2_AfterDelete(row *PushChat2) {
	RowCache.Delete("PushChat2:" + strconv.Itoa(row.SyncId))
}

func OnPushChat2_LoadOne(row *PushChat2) {
	RowCache.Set("PushChat2:"+strconv.Itoa(row.SyncId), row, time.Hour*0)
}

func OnPushChat2_LoadMany(rows []*PushChat2) {
	for _, row := range rows {
		RowCache.Set("PushChat2:"+strconv.Itoa(row.SyncId), row, time.Hour*0)
	}
}

//HTTPRPCLog Events

func OnHTTPRPCLog_AfterInsert(row *HTTPRPCLog) {
	RowCache.Set("HTTPRPCLog:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnHTTPRPCLog_AfterUpdate(row *HTTPRPCLog) {
	RowCache.Set("HTTPRPCLog:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnHTTPRPCLog_AfterDelete(row *HTTPRPCLog) {
	RowCache.Delete("HTTPRPCLog:" + strconv.Itoa(row.Id))
}

func OnHTTPRPCLog_LoadOne(row *HTTPRPCLog) {
	RowCache.Set("HTTPRPCLog:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnHTTPRPCLog_LoadMany(rows []*HTTPRPCLog) {
	for _, row := range rows {
		RowCache.Set("HTTPRPCLog:"+strconv.Itoa(row.Id), row, time.Hour*0)
	}
}

//MetricLog Events

func OnMetricLog_AfterInsert(row *MetricLog) {
	RowCache.Set("MetricLog:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnMetricLog_AfterUpdate(row *MetricLog) {
	RowCache.Set("MetricLog:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnMetricLog_AfterDelete(row *MetricLog) {
	RowCache.Delete("MetricLog:" + strconv.Itoa(row.Id))
}

func OnMetricLog_LoadOne(row *MetricLog) {
	RowCache.Set("MetricLog:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnMetricLog_LoadMany(rows []*MetricLog) {
	for _, row := range rows {
		RowCache.Set("MetricLog:"+strconv.Itoa(row.Id), row, time.Hour*0)
	}
}

//XfileServiceInfoLog Events

func OnXfileServiceInfoLog_AfterInsert(row *XfileServiceInfoLog) {
	RowCache.Set("XfileServiceInfoLog:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnXfileServiceInfoLog_AfterUpdate(row *XfileServiceInfoLog) {
	RowCache.Set("XfileServiceInfoLog:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnXfileServiceInfoLog_AfterDelete(row *XfileServiceInfoLog) {
	RowCache.Delete("XfileServiceInfoLog:" + strconv.Itoa(row.Id))
}

func OnXfileServiceInfoLog_LoadOne(row *XfileServiceInfoLog) {
	RowCache.Set("XfileServiceInfoLog:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnXfileServiceInfoLog_LoadMany(rows []*XfileServiceInfoLog) {
	for _, row := range rows {
		RowCache.Set("XfileServiceInfoLog:"+strconv.Itoa(row.Id), row, time.Hour*0)
	}
}

//XfileServiceMetricLog Events

func OnXfileServiceMetricLog_AfterInsert(row *XfileServiceMetricLog) {
	RowCache.Set("XfileServiceMetricLog:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnXfileServiceMetricLog_AfterUpdate(row *XfileServiceMetricLog) {
	RowCache.Set("XfileServiceMetricLog:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnXfileServiceMetricLog_AfterDelete(row *XfileServiceMetricLog) {
	RowCache.Delete("XfileServiceMetricLog:" + strconv.Itoa(row.Id))
}

func OnXfileServiceMetricLog_LoadOne(row *XfileServiceMetricLog) {
	RowCache.Set("XfileServiceMetricLog:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnXfileServiceMetricLog_LoadMany(rows []*XfileServiceMetricLog) {
	for _, row := range rows {
		RowCache.Set("XfileServiceMetricLog:"+strconv.Itoa(row.Id), row, time.Hour*0)
	}
}

//XfileServiceRequestLog Events

func OnXfileServiceRequestLog_AfterInsert(row *XfileServiceRequestLog) {
	RowCache.Set("XfileServiceRequestLog:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnXfileServiceRequestLog_AfterUpdate(row *XfileServiceRequestLog) {
	RowCache.Set("XfileServiceRequestLog:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnXfileServiceRequestLog_AfterDelete(row *XfileServiceRequestLog) {
	RowCache.Delete("XfileServiceRequestLog:" + strconv.Itoa(row.Id))
}

func OnXfileServiceRequestLog_LoadOne(row *XfileServiceRequestLog) {
	RowCache.Set("XfileServiceRequestLog:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnXfileServiceRequestLog_LoadMany(rows []*XfileServiceRequestLog) {
	for _, row := range rows {
		RowCache.Set("XfileServiceRequestLog:"+strconv.Itoa(row.Id), row, time.Hour*0)
	}
}

//Account Events

func OnAccount_AfterInsert(row *Account) {
	RowCache.Set("Account:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnAccount_AfterUpdate(row *Account) {
	RowCache.Set("Account:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnAccount_AfterDelete(row *Account) {
	RowCache.Delete("Account:" + strconv.Itoa(row.Id))
}

func OnAccount_LoadOne(row *Account) {
	RowCache.Set("Account:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnAccount_LoadMany(rows []*Account) {
	for _, row := range rows {
		RowCache.Set("Account:"+strconv.Itoa(row.Id), row, time.Hour*0)
	}
}

//PostCdb Events

func OnPostCdb_AfterInsert(row *PostCdb) {
	RowCache.Set("PostCdb:"+strconv.Itoa(row.PostId), row, time.Hour*0)
}

func OnPostCdb_AfterUpdate(row *PostCdb) {
	RowCache.Set("PostCdb:"+strconv.Itoa(row.PostId), row, time.Hour*0)
}

func OnPostCdb_AfterDelete(row *PostCdb) {
	RowCache.Delete("PostCdb:" + strconv.Itoa(row.PostId))
}

func OnPostCdb_LoadOne(row *PostCdb) {
	RowCache.Set("PostCdb:"+strconv.Itoa(row.PostId), row, time.Hour*0)
}

func OnPostCdb_LoadMany(rows []*PostCdb) {
	for _, row := range rows {
		RowCache.Set("PostCdb:"+strconv.Itoa(row.PostId), row, time.Hour*0)
	}
}
