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

//DirectOffline Events

func OnDirectOffline_AfterInsert(row *DirectOffline) {
	RowCache.Set("DirectOffline:"+strconv.Itoa(row.DirectOfflineId), row, time.Hour*0)
}

func OnDirectOffline_AfterUpdate(row *DirectOffline) {
	RowCache.Set("DirectOffline:"+strconv.Itoa(row.DirectOfflineId), row, time.Hour*0)
}

func OnDirectOffline_AfterDelete(row *DirectOffline) {
	RowCache.Delete("DirectOffline:" + strconv.Itoa(row.DirectOfflineId))
}

func OnDirectOffline_LoadOne(row *DirectOffline) {
	RowCache.Set("DirectOffline:"+strconv.Itoa(row.DirectOfflineId), row, time.Hour*0)
}

func OnDirectOffline_LoadMany(rows []*DirectOffline) {
	for _, row := range rows {
		RowCache.Set("DirectOffline:"+strconv.Itoa(row.DirectOfflineId), row, time.Hour*0)
	}
}

//DirectToMessage Events

func OnDirectToMessage_AfterInsert(row *DirectToMessage) {
	RowCache.Set("DirectToMessage:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnDirectToMessage_AfterUpdate(row *DirectToMessage) {
	RowCache.Set("DirectToMessage:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnDirectToMessage_AfterDelete(row *DirectToMessage) {
	RowCache.Delete("DirectToMessage:" + strconv.Itoa(row.Id))
}

func OnDirectToMessage_LoadOne(row *DirectToMessage) {
	RowCache.Set("DirectToMessage:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnDirectToMessage_LoadMany(rows []*DirectToMessage) {
	for _, row := range rows {
		RowCache.Set("DirectToMessage:"+strconv.Itoa(row.Id), row, time.Hour*0)
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

//FollowingListMemberHistory Events

func OnFollowingListMemberHistory_AfterInsert(row *FollowingListMemberHistory) {
	RowCache.Set("FollowingListMemberHistory:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnFollowingListMemberHistory_AfterUpdate(row *FollowingListMemberHistory) {
	RowCache.Set("FollowingListMemberHistory:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnFollowingListMemberHistory_AfterDelete(row *FollowingListMemberHistory) {
	RowCache.Delete("FollowingListMemberHistory:" + strconv.Itoa(row.Id))
}

func OnFollowingListMemberHistory_LoadOne(row *FollowingListMemberHistory) {
	RowCache.Set("FollowingListMemberHistory:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnFollowingListMemberHistory_LoadMany(rows []*FollowingListMemberHistory) {
	for _, row := range rows {
		RowCache.Set("FollowingListMemberHistory:"+strconv.Itoa(row.Id), row, time.Hour*0)
	}
}

//GeneralLog Events

func OnGeneralLog_AfterInsert(row *GeneralLog) {
	RowCache.Set("GeneralLog:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnGeneralLog_AfterUpdate(row *GeneralLog) {
	RowCache.Set("GeneralLog:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnGeneralLog_AfterDelete(row *GeneralLog) {
	RowCache.Delete("GeneralLog:" + strconv.Itoa(row.Id))
}

func OnGeneralLog_LoadOne(row *GeneralLog) {
	RowCache.Set("GeneralLog:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnGeneralLog_LoadMany(rows []*GeneralLog) {
	for _, row := range rows {
		RowCache.Set("GeneralLog:"+strconv.Itoa(row.Id), row, time.Hour*0)
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

//PhoneContactsCopy Events

func OnPhoneContactsCopy_AfterInsert(row *PhoneContactsCopy) {
	RowCache.Set("PhoneContactsCopy:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnPhoneContactsCopy_AfterUpdate(row *PhoneContactsCopy) {
	RowCache.Set("PhoneContactsCopy:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnPhoneContactsCopy_AfterDelete(row *PhoneContactsCopy) {
	RowCache.Delete("PhoneContactsCopy:" + strconv.Itoa(row.Id))
}

func OnPhoneContactsCopy_LoadOne(row *PhoneContactsCopy) {
	RowCache.Set("PhoneContactsCopy:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnPhoneContactsCopy_LoadMany(rows []*PhoneContactsCopy) {
	for _, row := range rows {
		RowCache.Set("PhoneContactsCopy:"+strconv.Itoa(row.Id), row, time.Hour*0)
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

//RecommendUser Events

func OnRecommendUser_AfterInsert(row *RecommendUser) {
	RowCache.Set("RecommendUser:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnRecommendUser_AfterUpdate(row *RecommendUser) {
	RowCache.Set("RecommendUser:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnRecommendUser_AfterDelete(row *RecommendUser) {
	RowCache.Delete("RecommendUser:" + strconv.Itoa(row.Id))
}

func OnRecommendUser_LoadOne(row *RecommendUser) {
	RowCache.Set("RecommendUser:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnRecommendUser_LoadMany(rows []*RecommendUser) {
	for _, row := range rows {
		RowCache.Set("RecommendUser:"+strconv.Itoa(row.Id), row, time.Hour*0)
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
