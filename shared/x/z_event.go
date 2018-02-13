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

//Media Events

func OnMedia_AfterInsert(row *Media) {
	RowCache.Set("Media:"+strconv.Itoa(row.MediaId), row, time.Hour*0)
}

func OnMedia_AfterUpdate(row *Media) {
	RowCache.Set("Media:"+strconv.Itoa(row.MediaId), row, time.Hour*0)
}

func OnMedia_AfterDelete(row *Media) {
	RowCache.Delete("Media:" + strconv.Itoa(row.MediaId))
}

func OnMedia_LoadOne(row *Media) {
	RowCache.Set("Media:"+strconv.Itoa(row.MediaId), row, time.Hour*0)
}

func OnMedia_LoadMany(rows []*Media) {
	for _, row := range rows {
		RowCache.Set("Media:"+strconv.Itoa(row.MediaId), row, time.Hour*0)
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

//TagsPost Events

func OnTagsPost_AfterInsert(row *TagsPost) {
	RowCache.Set("TagsPost:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnTagsPost_AfterUpdate(row *TagsPost) {
	RowCache.Set("TagsPost:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnTagsPost_AfterDelete(row *TagsPost) {
	RowCache.Delete("TagsPost:" + strconv.Itoa(row.Id))
}

func OnTagsPost_LoadOne(row *TagsPost) {
	RowCache.Set("TagsPost:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnTagsPost_LoadMany(rows []*TagsPost) {
	for _, row := range rows {
		RowCache.Set("TagsPost:"+strconv.Itoa(row.Id), row, time.Hour*0)
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

//ChatSync Events

func OnChatSync_AfterInsert(row *ChatSync) {
	RowCache.Set("ChatSync:"+strconv.Itoa(row.SyncId), row, time.Hour*0)
}

func OnChatSync_AfterUpdate(row *ChatSync) {
	RowCache.Set("ChatSync:"+strconv.Itoa(row.SyncId), row, time.Hour*0)
}

func OnChatSync_AfterDelete(row *ChatSync) {
	RowCache.Delete("ChatSync:" + strconv.Itoa(row.SyncId))
}

func OnChatSync_LoadOne(row *ChatSync) {
	RowCache.Set("ChatSync:"+strconv.Itoa(row.SyncId), row, time.Hour*0)
}

func OnChatSync_LoadMany(rows []*ChatSync) {
	for _, row := range rows {
		RowCache.Set("ChatSync:"+strconv.Itoa(row.SyncId), row, time.Hour*0)
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

//Home Events

func OnHome_AfterInsert(row *Home) {
	RowCache.Set("Home:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnHome_AfterUpdate(row *Home) {
	RowCache.Set("Home:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnHome_AfterDelete(row *Home) {
	RowCache.Delete("Home:" + strconv.Itoa(row.Id))
}

func OnHome_LoadOne(row *Home) {
	RowCache.Set("Home:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnHome_LoadMany(rows []*Home) {
	for _, row := range rows {
		RowCache.Set("Home:"+strconv.Itoa(row.Id), row, time.Hour*0)
	}
}

//MessageFile Events

func OnMessageFile_AfterInsert(row *MessageFile) {
	RowCache.Set("MessageFile:"+strconv.Itoa(row.MessageFileId), row, time.Hour*0)
}

func OnMessageFile_AfterUpdate(row *MessageFile) {
	RowCache.Set("MessageFile:"+strconv.Itoa(row.MessageFileId), row, time.Hour*0)
}

func OnMessageFile_AfterDelete(row *MessageFile) {
	RowCache.Delete("MessageFile:" + strconv.Itoa(row.MessageFileId))
}

func OnMessageFile_LoadOne(row *MessageFile) {
	RowCache.Set("MessageFile:"+strconv.Itoa(row.MessageFileId), row, time.Hour*0)
}

func OnMessageFile_LoadMany(rows []*MessageFile) {
	for _, row := range rows {
		RowCache.Set("MessageFile:"+strconv.Itoa(row.MessageFileId), row, time.Hour*0)
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
