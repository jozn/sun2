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

//Blocked Events

func OnBlocked_AfterInsert(row *Blocked) {
	RowCache.Set("Blocked:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnBlocked_AfterUpdate(row *Blocked) {
	RowCache.Set("Blocked:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnBlocked_AfterDelete(row *Blocked) {
	RowCache.Delete("Blocked:" + strconv.Itoa(row.Id))
}

func OnBlocked_LoadOne(row *Blocked) {
	RowCache.Set("Blocked:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnBlocked_LoadMany(rows []*Blocked) {
	for _, row := range rows {
		RowCache.Set("Blocked:"+strconv.Itoa(row.Id), row, time.Hour*0)
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

//Followed Events

func OnFollowed_AfterInsert(row *Followed) {
	RowCache.Set("Followed:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnFollowed_AfterUpdate(row *Followed) {
	RowCache.Set("Followed:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnFollowed_AfterDelete(row *Followed) {
	RowCache.Delete("Followed:" + strconv.Itoa(row.Id))
}

func OnFollowed_LoadOne(row *Followed) {
	RowCache.Set("Followed:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnFollowed_LoadMany(rows []*Followed) {
	for _, row := range rows {
		RowCache.Set("Followed:"+strconv.Itoa(row.Id), row, time.Hour*0)
	}
}

//Likes Events

func OnLikes_AfterInsert(row *Likes) {
	RowCache.Set("Likes:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnLikes_AfterUpdate(row *Likes) {
	RowCache.Set("Likes:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnLikes_AfterDelete(row *Likes) {
	RowCache.Delete("Likes:" + strconv.Itoa(row.Id))
}

func OnLikes_LoadOne(row *Likes) {
	RowCache.Set("Likes:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnLikes_LoadMany(rows []*Likes) {
	for _, row := range rows {
		RowCache.Set("Likes:"+strconv.Itoa(row.Id), row, time.Hour*0)
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

//PhoneContacts Events

func OnPhoneContacts_AfterInsert(row *PhoneContacts) {
	RowCache.Set("PhoneContacts:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnPhoneContacts_AfterUpdate(row *PhoneContacts) {
	RowCache.Set("PhoneContacts:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnPhoneContacts_AfterDelete(row *PhoneContacts) {
	RowCache.Delete("PhoneContacts:" + strconv.Itoa(row.Id))
}

func OnPhoneContacts_LoadOne(row *PhoneContacts) {
	RowCache.Set("PhoneContacts:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnPhoneContacts_LoadMany(rows []*PhoneContacts) {
	for _, row := range rows {
		RowCache.Set("PhoneContacts:"+strconv.Itoa(row.Id), row, time.Hour*0)
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

//PostKeys Events

func OnPostKeys_AfterInsert(row *PostKeys) {
	RowCache.Set("PostKeys:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnPostKeys_AfterUpdate(row *PostKeys) {
	RowCache.Set("PostKeys:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnPostKeys_AfterDelete(row *PostKeys) {
	RowCache.Delete("PostKeys:" + strconv.Itoa(row.Id))
}

func OnPostKeys_LoadOne(row *PostKeys) {
	RowCache.Set("PostKeys:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnPostKeys_LoadMany(rows []*PostKeys) {
	for _, row := range rows {
		RowCache.Set("PostKeys:"+strconv.Itoa(row.Id), row, time.Hour*0)
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

//PostPromoted Events

func OnPostPromoted_AfterInsert(row *PostPromoted) {
	RowCache.Set("PostPromoted:"+strconv.Itoa(row.PromoteId), row, time.Hour*0)
}

func OnPostPromoted_AfterUpdate(row *PostPromoted) {
	RowCache.Set("PostPromoted:"+strconv.Itoa(row.PromoteId), row, time.Hour*0)
}

func OnPostPromoted_AfterDelete(row *PostPromoted) {
	RowCache.Delete("PostPromoted:" + strconv.Itoa(row.PromoteId))
}

func OnPostPromoted_LoadOne(row *PostPromoted) {
	RowCache.Set("PostPromoted:"+strconv.Itoa(row.PromoteId), row, time.Hour*0)
}

func OnPostPromoted_LoadMany(rows []*PostPromoted) {
	for _, row := range rows {
		RowCache.Set("PostPromoted:"+strconv.Itoa(row.PromoteId), row, time.Hour*0)
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

//ProfileAll Events

func OnProfileAll_AfterInsert(row *ProfileAll) {
	RowCache.Set("ProfileAll:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnProfileAll_AfterUpdate(row *ProfileAll) {
	RowCache.Set("ProfileAll:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnProfileAll_AfterDelete(row *ProfileAll) {
	RowCache.Delete("ProfileAll:" + strconv.Itoa(row.Id))
}

func OnProfileAll_LoadOne(row *ProfileAll) {
	RowCache.Set("ProfileAll:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnProfileAll_LoadMany(rows []*ProfileAll) {
	for _, row := range rows {
		RowCache.Set("ProfileAll:"+strconv.Itoa(row.Id), row, time.Hour*0)
	}
}

//ProfileMedia Events

func OnProfileMedia_AfterInsert(row *ProfileMedia) {
	RowCache.Set("ProfileMedia:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnProfileMedia_AfterUpdate(row *ProfileMedia) {
	RowCache.Set("ProfileMedia:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnProfileMedia_AfterDelete(row *ProfileMedia) {
	RowCache.Delete("ProfileMedia:" + strconv.Itoa(row.Id))
}

func OnProfileMedia_LoadOne(row *ProfileMedia) {
	RowCache.Set("ProfileMedia:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnProfileMedia_LoadMany(rows []*ProfileMedia) {
	for _, row := range rows {
		RowCache.Set("ProfileMedia:"+strconv.Itoa(row.Id), row, time.Hour*0)
	}
}

//ProfileMentioned Events

func OnProfileMentioned_AfterInsert(row *ProfileMentioned) {
	RowCache.Set("ProfileMentioned:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnProfileMentioned_AfterUpdate(row *ProfileMentioned) {
	RowCache.Set("ProfileMentioned:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnProfileMentioned_AfterDelete(row *ProfileMentioned) {
	RowCache.Delete("ProfileMentioned:" + strconv.Itoa(row.Id))
}

func OnProfileMentioned_LoadOne(row *ProfileMentioned) {
	RowCache.Set("ProfileMentioned:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnProfileMentioned_LoadMany(rows []*ProfileMentioned) {
	for _, row := range rows {
		RowCache.Set("ProfileMentioned:"+strconv.Itoa(row.Id), row, time.Hour*0)
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

//SettingNotifications Events

func OnSettingNotifications_AfterInsert(row *SettingNotifications) {
	RowCache.Set("SettingNotifications:"+strconv.Itoa(row.UserId), row, time.Hour*0)
}

func OnSettingNotifications_AfterUpdate(row *SettingNotifications) {
	RowCache.Set("SettingNotifications:"+strconv.Itoa(row.UserId), row, time.Hour*0)
}

func OnSettingNotifications_AfterDelete(row *SettingNotifications) {
	RowCache.Delete("SettingNotifications:" + strconv.Itoa(row.UserId))
}

func OnSettingNotifications_LoadOne(row *SettingNotifications) {
	RowCache.Set("SettingNotifications:"+strconv.Itoa(row.UserId), row, time.Hour*0)
}

func OnSettingNotifications_LoadMany(rows []*SettingNotifications) {
	for _, row := range rows {
		RowCache.Set("SettingNotifications:"+strconv.Itoa(row.UserId), row, time.Hour*0)
	}
}

//Sms Events

func OnSms_AfterInsert(row *Sms) {
	RowCache.Set("Sms:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnSms_AfterUpdate(row *Sms) {
	RowCache.Set("Sms:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnSms_AfterDelete(row *Sms) {
	RowCache.Delete("Sms:" + strconv.Itoa(row.Id))
}

func OnSms_LoadOne(row *Sms) {
	RowCache.Set("Sms:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnSms_LoadMany(rows []*Sms) {
	for _, row := range rows {
		RowCache.Set("Sms:"+strconv.Itoa(row.Id), row, time.Hour*0)
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

//UserRelation Events

func OnUserRelation_AfterInsert(row *UserRelation) {
	RowCache.Set("UserRelation:"+strconv.Itoa(row.RelNanoId), row, time.Hour*0)
}

func OnUserRelation_AfterUpdate(row *UserRelation) {
	RowCache.Set("UserRelation:"+strconv.Itoa(row.RelNanoId), row, time.Hour*0)
}

func OnUserRelation_AfterDelete(row *UserRelation) {
	RowCache.Delete("UserRelation:" + strconv.Itoa(row.RelNanoId))
}

func OnUserRelation_LoadOne(row *UserRelation) {
	RowCache.Set("UserRelation:"+strconv.Itoa(row.RelNanoId), row, time.Hour*0)
}

func OnUserRelation_LoadMany(rows []*UserRelation) {
	for _, row := range rows {
		RowCache.Set("UserRelation:"+strconv.Itoa(row.RelNanoId), row, time.Hour*0)
	}
}

//Chat Events

func OnChat_AfterInsert(row *Chat) {
	RowCache.Set("Chat:"+strconv.Itoa(row.ChatId), row, time.Hour*0)
}

func OnChat_AfterUpdate(row *Chat) {
	RowCache.Set("Chat:"+strconv.Itoa(row.ChatId), row, time.Hour*0)
}

func OnChat_AfterDelete(row *Chat) {
	RowCache.Delete("Chat:" + strconv.Itoa(row.ChatId))
}

func OnChat_LoadOne(row *Chat) {
	RowCache.Set("Chat:"+strconv.Itoa(row.ChatId), row, time.Hour*0)
}

func OnChat_LoadMany(rows []*Chat) {
	for _, row := range rows {
		RowCache.Set("Chat:"+strconv.Itoa(row.ChatId), row, time.Hour*0)
	}
}

//ChatDeleted Events

func OnChatDeleted_AfterInsert(row *ChatDeleted) {
	RowCache.Set("ChatDeleted:"+strconv.Itoa(row.ChatId), row, time.Hour*0)
}

func OnChatDeleted_AfterUpdate(row *ChatDeleted) {
	RowCache.Set("ChatDeleted:"+strconv.Itoa(row.ChatId), row, time.Hour*0)
}

func OnChatDeleted_AfterDelete(row *ChatDeleted) {
	RowCache.Delete("ChatDeleted:" + strconv.Itoa(row.ChatId))
}

func OnChatDeleted_LoadOne(row *ChatDeleted) {
	RowCache.Set("ChatDeleted:"+strconv.Itoa(row.ChatId), row, time.Hour*0)
}

func OnChatDeleted_LoadMany(rows []*ChatDeleted) {
	for _, row := range rows {
		RowCache.Set("ChatDeleted:"+strconv.Itoa(row.ChatId), row, time.Hour*0)
	}
}

//ChatLastMessage Events

func OnChatLastMessage_AfterInsert(row *ChatLastMessage) {
	RowCache.Set("ChatLastMessage:"+row.ChatIdGroupId, row, time.Hour*0)
}

func OnChatLastMessage_AfterUpdate(row *ChatLastMessage) {
	RowCache.Set("ChatLastMessage:"+row.ChatIdGroupId, row, time.Hour*0)
}

func OnChatLastMessage_AfterDelete(row *ChatLastMessage) {
	RowCache.Delete("ChatLastMessage:" + row.ChatIdGroupId)
}

func OnChatLastMessage_LoadOne(row *ChatLastMessage) {
	RowCache.Set("ChatLastMessage:"+row.ChatIdGroupId, row, time.Hour*0)
}

func OnChatLastMessage_LoadMany(rows []*ChatLastMessage) {
	for _, row := range rows {
		RowCache.Set("ChatLastMessage:"+row.ChatIdGroupId, row, time.Hour*0)
	}
}

//ChatUserVersion Events

func OnChatUserVersion_AfterInsert(row *ChatUserVersion) {
	RowCache.Set("ChatUserVersion:"+strconv.Itoa(row.VersionTimeNano), row, time.Hour*0)
}

func OnChatUserVersion_AfterUpdate(row *ChatUserVersion) {
	RowCache.Set("ChatUserVersion:"+strconv.Itoa(row.VersionTimeNano), row, time.Hour*0)
}

func OnChatUserVersion_AfterDelete(row *ChatUserVersion) {
	RowCache.Delete("ChatUserVersion:" + strconv.Itoa(row.VersionTimeNano))
}

func OnChatUserVersion_LoadOne(row *ChatUserVersion) {
	RowCache.Set("ChatUserVersion:"+strconv.Itoa(row.VersionTimeNano), row, time.Hour*0)
}

func OnChatUserVersion_LoadMany(rows []*ChatUserVersion) {
	for _, row := range rows {
		RowCache.Set("ChatUserVersion:"+strconv.Itoa(row.VersionTimeNano), row, time.Hour*0)
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
	RowCache.Set("GroupMember:"+strconv.Itoa(row.OrderId), row, time.Hour*0)
}

func OnGroupMember_AfterUpdate(row *GroupMember) {
	RowCache.Set("GroupMember:"+strconv.Itoa(row.OrderId), row, time.Hour*0)
}

func OnGroupMember_AfterDelete(row *GroupMember) {
	RowCache.Delete("GroupMember:" + strconv.Itoa(row.OrderId))
}

func OnGroupMember_LoadOne(row *GroupMember) {
	RowCache.Set("GroupMember:"+strconv.Itoa(row.OrderId), row, time.Hour*0)
}

func OnGroupMember_LoadMany(rows []*GroupMember) {
	for _, row := range rows {
		RowCache.Set("GroupMember:"+strconv.Itoa(row.OrderId), row, time.Hour*0)
	}
}

//GroupOrderdUser Events

func OnGroupOrderdUser_AfterInsert(row *GroupOrderdUser) {
	RowCache.Set("GroupOrderdUser:"+strconv.Itoa(row.OrderId), row, time.Hour*0)
}

func OnGroupOrderdUser_AfterUpdate(row *GroupOrderdUser) {
	RowCache.Set("GroupOrderdUser:"+strconv.Itoa(row.OrderId), row, time.Hour*0)
}

func OnGroupOrderdUser_AfterDelete(row *GroupOrderdUser) {
	RowCache.Delete("GroupOrderdUser:" + strconv.Itoa(row.OrderId))
}

func OnGroupOrderdUser_LoadOne(row *GroupOrderdUser) {
	RowCache.Set("GroupOrderdUser:"+strconv.Itoa(row.OrderId), row, time.Hour*0)
}

func OnGroupOrderdUser_LoadMany(rows []*GroupOrderdUser) {
	for _, row := range rows {
		RowCache.Set("GroupOrderdUser:"+strconv.Itoa(row.OrderId), row, time.Hour*0)
	}
}

//GroupPinedMsg Events

func OnGroupPinedMsg_AfterInsert(row *GroupPinedMsg) {
	RowCache.Set("GroupPinedMsg:"+strconv.Itoa(row.MsgId), row, time.Hour*0)
}

func OnGroupPinedMsg_AfterUpdate(row *GroupPinedMsg) {
	RowCache.Set("GroupPinedMsg:"+strconv.Itoa(row.MsgId), row, time.Hour*0)
}

func OnGroupPinedMsg_AfterDelete(row *GroupPinedMsg) {
	RowCache.Delete("GroupPinedMsg:" + strconv.Itoa(row.MsgId))
}

func OnGroupPinedMsg_LoadOne(row *GroupPinedMsg) {
	RowCache.Set("GroupPinedMsg:"+strconv.Itoa(row.MsgId), row, time.Hour*0)
}

func OnGroupPinedMsg_LoadMany(rows []*GroupPinedMsg) {
	for _, row := range rows {
		RowCache.Set("GroupPinedMsg:"+strconv.Itoa(row.MsgId), row, time.Hour*0)
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

//SuggestedTopPosts Events

func OnSuggestedTopPosts_AfterInsert(row *SuggestedTopPosts) {
	RowCache.Set("SuggestedTopPosts:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnSuggestedTopPosts_AfterUpdate(row *SuggestedTopPosts) {
	RowCache.Set("SuggestedTopPosts:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnSuggestedTopPosts_AfterDelete(row *SuggestedTopPosts) {
	RowCache.Delete("SuggestedTopPosts:" + strconv.Itoa(row.Id))
}

func OnSuggestedTopPosts_LoadOne(row *SuggestedTopPosts) {
	RowCache.Set("SuggestedTopPosts:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnSuggestedTopPosts_LoadMany(rows []*SuggestedTopPosts) {
	for _, row := range rows {
		RowCache.Set("SuggestedTopPosts:"+strconv.Itoa(row.Id), row, time.Hour*0)
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

//InvalidateCache Events

func OnInvalidateCache_AfterInsert(row *InvalidateCache) {
	RowCache.Set("InvalidateCache:"+strconv.Itoa(row.OrderId), row, time.Hour*0)
}

func OnInvalidateCache_AfterUpdate(row *InvalidateCache) {
	RowCache.Set("InvalidateCache:"+strconv.Itoa(row.OrderId), row, time.Hour*0)
}

func OnInvalidateCache_AfterDelete(row *InvalidateCache) {
	RowCache.Delete("InvalidateCache:" + strconv.Itoa(row.OrderId))
}

func OnInvalidateCache_LoadOne(row *InvalidateCache) {
	RowCache.Set("InvalidateCache:"+strconv.Itoa(row.OrderId), row, time.Hour*0)
}

func OnInvalidateCache_LoadMany(rows []*InvalidateCache) {
	for _, row := range rows {
		RowCache.Set("InvalidateCache:"+strconv.Itoa(row.OrderId), row, time.Hour*0)
	}
}
