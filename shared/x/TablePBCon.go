package x

/*
func PBConvPB__Action_To_Action( o *PB_Action) *Action {
     n := &Action{
      ActionId: int ( o.ActionId ),
      ActorUserId: int ( o.ActorUserId ),
      ActionTypeEnum: int ( o.ActionTypeEnum ),
      PeerUserId: int ( o.PeerUserId ),
      PostId: int ( o.PostId ),
      CommentId: int ( o.CommentId ),
      Murmur64Hash: int ( o.Murmur64Hash ),
      CreatedTime: int ( o.CreatedTime ),
      Seq: int ( o.Seq ),
    }
    return n
}

func PBConvPB_Action_To_Action ( o *Action) *PB_Action {
     n := &PB_Action{
      ActionId: int64 ( o.ActionId ),
      ActorUserId: int32 ( o.ActorUserId ),
      ActionTypeEnum: int32 ( o.ActionTypeEnum ),
      PeerUserId: int32 ( o.PeerUserId ),
      PostId: int64 ( o.PostId ),
      CommentId: int64 ( o.CommentId ),
      Murmur64Hash: int64 ( o.Murmur64Hash ),
      CreatedTime: int32 ( o.CreatedTime ),
      Seq: int32 ( o.Seq ),
    }
    return n
}
*/
/*
func PBConvPB__Chat_To_Chat( o *PB_Chat) *Chat {
     n := &Chat{
      ChatKey: string ( o.ChatKey ),
      RoomKey: string ( o.RoomKey ),
      RoomTypeEnum: int ( o.RoomTypeEnum ),
      UserId: int ( o.UserId ),
      PeerUserId: int ( o.PeerUserId ),
      GroupId: int ( o.GroupId ),
      CreatedTime: int ( o.CreatedTime ),
      StartMessageIdFrom: int ( o.StartMessageIdFrom ),
      LastMessageId: int ( o.LastMessageId ),
      LastSeenMessageId: int ( o.LastSeenMessageId ),
      UpdatedMs: int ( o.UpdatedMs ),
    }
    return n
}

func PBConvPB_Chat_To_Chat ( o *Chat) *PB_Chat {
     n := &PB_Chat{
      ChatKey: string ( o.ChatKey ),
      RoomKey: string ( o.RoomKey ),
      RoomTypeEnum: int32 ( o.RoomTypeEnum ),
      UserId: int32 ( o.UserId ),
      PeerUserId: int32 ( o.PeerUserId ),
      GroupId: int64 ( o.GroupId ),
      CreatedTime: int32 ( o.CreatedTime ),
      StartMessageIdFrom: int64 ( o.StartMessageIdFrom ),
      LastMessageId: int64 ( o.LastMessageId ),
      LastSeenMessageId: int64 ( o.LastSeenMessageId ),
      UpdatedMs: int64 ( o.UpdatedMs ),
    }
    return n
}
*/
/*
func PBConvPB__Comment_To_Comment( o *PB_Comment) *Comment {
     n := &Comment{
      CommentId: int ( o.CommentId ),
      UserId: int ( o.UserId ),
      PostId: int ( o.PostId ),
      Text: string ( o.Text ),
      LikesCount: int ( o.LikesCount ),
      CreatedTime: int ( o.CreatedTime ),
      Seq: int ( o.Seq ),
    }
    return n
}

func PBConvPB_Comment_To_Comment ( o *Comment) *PB_Comment {
     n := &PB_Comment{
      CommentId: int64 ( o.CommentId ),
      UserId: int32 ( o.UserId ),
      PostId: int32 ( o.PostId ),
      Text: string ( o.Text ),
      LikesCount: int32 ( o.LikesCount ),
      CreatedTime: int32 ( o.CreatedTime ),
      Seq: int32 ( o.Seq ),
    }
    return n
}
*/
/*
func PBConvPB__DirectMessage_To_DirectMessage( o *PB_DirectMessage) *DirectMessage {
     n := &DirectMessage{
      MessageId: int ( o.MessageId ),
      MessageKey: string ( o.MessageKey ),
      RoomKey: string ( o.RoomKey ),
      UserId: int ( o.UserId ),
      MessageFileId: int ( o.MessageFileId ),
      MessageTypeEnumId: int ( o.MessageTypeEnumId ),
      Text: string ( o.Text ),
      CreatedSe: int ( o.CreatedSe ),
      PeerReceivedTime: int ( o.PeerReceivedTime ),
      PeerSeenTime: int ( o.PeerSeenTime ),
      DeliviryStatusEnumId: int ( o.DeliviryStatusEnumId ),
    }
    return n
}

func PBConvPB_DirectMessage_To_DirectMessage ( o *DirectMessage) *PB_DirectMessage {
     n := &PB_DirectMessage{
      MessageId: int64 ( o.MessageId ),
      MessageKey: string ( o.MessageKey ),
      RoomKey: string ( o.RoomKey ),
      UserId: int32 ( o.UserId ),
      MessageFileId: int64 ( o.MessageFileId ),
      MessageTypeEnumId: int32 ( o.MessageTypeEnumId ),
      Text: string ( o.Text ),
      CreatedSe: int32 ( o.CreatedSe ),
      PeerReceivedTime: int32 ( o.PeerReceivedTime ),
      PeerSeenTime: int32 ( o.PeerSeenTime ),
      DeliviryStatusEnumId: int32 ( o.DeliviryStatusEnumId ),
    }
    return n
}
*/
/*
func PBConvPB__DirectOffline_To_DirectOffline( o *PB_DirectOffline) *DirectOffline {
     n := &DirectOffline{
      DirectOfflineId: int ( o.DirectOfflineId ),
      ToUserId: int ( o.ToUserId ),
      ChatKey: string ( o.ChatKey ),
      MessageId: int ( o.MessageId ),
      MessageFileId: int ( o.MessageFileId ),
      PBClass: string ( o.PBClass ),
      DataPB: []byte ( o.DataPB ),
      DataJson: string ( o.DataJson ),
      DataTemp: string ( o.DataTemp ),
      AtTimeMs: int ( o.AtTimeMs ),
    }
    return n
}

func PBConvPB_DirectOffline_To_DirectOffline ( o *DirectOffline) *PB_DirectOffline {
     n := &PB_DirectOffline{
      DirectOfflineId: int64 ( o.DirectOfflineId ),
      ToUserId: int32 ( o.ToUserId ),
      ChatKey: string ( o.ChatKey ),
      MessageId: int64 ( o.MessageId ),
      MessageFileId: int64 ( o.MessageFileId ),
      PBClass: string ( o.PBClass ),
      DataPB: []byte ( o.DataPB ),
      DataJson: string ( o.DataJson ),
      DataTemp: string ( o.DataTemp ),
      AtTimeMs: int64 ( o.AtTimeMs ),
    }
    return n
}
*/
/*
func PBConvPB__DirectToMessage_To_DirectToMessage( o *PB_DirectToMessage) *DirectToMessage {
     n := &DirectToMessage{
      Id: int ( o.Id ),
      ChatKey: string ( o.ChatKey ),
      MessageId: int ( o.MessageId ),
      SourceEnumId: int ( o.SourceEnumId ),
    }
    return n
}

func PBConvPB_DirectToMessage_To_DirectToMessage ( o *DirectToMessage) *PB_DirectToMessage {
     n := &PB_DirectToMessage{
      Id: int64 ( o.Id ),
      ChatKey: string ( o.ChatKey ),
      MessageId: int64 ( o.MessageId ),
      SourceEnumId: int32 ( o.SourceEnumId ),
    }
    return n
}
*/
/*
func PBConvPB__Feed_To_Feed( o *PB_Feed) *Feed {
     n := &Feed{
      UserId: int ( o.UserId ),
      PostId: int ( o.PostId ),
      FeedId: int ( o.FeedId ),
    }
    return n
}

func PBConvPB_Feed_To_Feed ( o *Feed) *PB_Feed {
     n := &PB_Feed{
      UserId: int64 ( o.UserId ),
      PostId: int64 ( o.PostId ),
      FeedId: int32 ( o.FeedId ),
    }
    return n
}
*/
/*
func PBConvPB__FollowingList_To_FollowingList( o *PB_FollowingList) *FollowingList {
     n := &FollowingList{
      Id: int ( o.Id ),
      UserId: int ( o.UserId ),
      ListType: int ( o.ListType ),
      Name: string ( o.Name ),
      Count: int ( o.Count ),
      IsAuto: int ( o.IsAuto ),
      IsPimiry: int ( o.IsPimiry ),
      CreatedTime: int ( o.CreatedTime ),
    }
    return n
}

func PBConvPB_FollowingList_To_FollowingList ( o *FollowingList) *PB_FollowingList {
     n := &PB_FollowingList{
      Id: int32 ( o.Id ),
      UserId: int32 ( o.UserId ),
      ListType: int32 ( o.ListType ),
      Name: string ( o.Name ),
      Count: int32 ( o.Count ),
      IsAuto: int32 ( o.IsAuto ),
      IsPimiry: int32 ( o.IsPimiry ),
      CreatedTime: int32 ( o.CreatedTime ),
    }
    return n
}
*/
/*
func PBConvPB__FollowingListMember_To_FollowingListMember( o *PB_FollowingListMember) *FollowingListMember {
     n := &FollowingListMember{
      Id: int ( o.Id ),
      ListId: int ( o.ListId ),
      UserId: int ( o.UserId ),
      FollowedUserId: int ( o.FollowedUserId ),
      FollowType: int ( o.FollowType ),
      UpdatedTimeMs: int ( o.UpdatedTimeMs ),
    }
    return n
}

func PBConvPB_FollowingListMember_To_FollowingListMember ( o *FollowingListMember) *PB_FollowingListMember {
     n := &PB_FollowingListMember{
      Id: int64 ( o.Id ),
      ListId: int32 ( o.ListId ),
      UserId: int32 ( o.UserId ),
      FollowedUserId: int32 ( o.FollowedUserId ),
      FollowType: int32 ( o.FollowType ),
      UpdatedTimeMs: int64 ( o.UpdatedTimeMs ),
    }
    return n
}
*/
/*
func PBConvPB__FollowingListMemberHistory_To_FollowingListMemberHistory( o *PB_FollowingListMemberHistory) *FollowingListMemberHistory {
     n := &FollowingListMemberHistory{
      Id: int ( o.Id ),
      ListId: int ( o.ListId ),
      UserId: int ( o.UserId ),
      FollowedUserId: int ( o.FollowedUserId ),
      FollowType: int ( o.FollowType ),
      UpdatedTimeMs: int ( o.UpdatedTimeMs ),
      FollowId: int ( o.FollowId ),
    }
    return n
}

func PBConvPB_FollowingListMemberHistory_To_FollowingListMemberHistory ( o *FollowingListMemberHistory) *PB_FollowingListMemberHistory {
     n := &PB_FollowingListMemberHistory{
      Id: int64 ( o.Id ),
      ListId: int32 ( o.ListId ),
      UserId: int32 ( o.UserId ),
      FollowedUserId: int32 ( o.FollowedUserId ),
      FollowType: int32 ( o.FollowType ),
      UpdatedTimeMs: int64 ( o.UpdatedTimeMs ),
      FollowId: int32 ( o.FollowId ),
    }
    return n
}
*/
/*
func PBConvPB__GeneralLog_To_GeneralLog( o *PB_GeneralLog) *GeneralLog {
     n := &GeneralLog{
      Id: int ( o.Id ),
      ToUserId: int ( o.ToUserId ),
      TargetId: int ( o.TargetId ),
      LogTypeId: int ( o.LogTypeId ),
      ExtraPB: []byte ( o.ExtraPB ),
      ExtraJson: string ( o.ExtraJson ),
      CreatedMs: int ( o.CreatedMs ),
    }
    return n
}

func PBConvPB_GeneralLog_To_GeneralLog ( o *GeneralLog) *PB_GeneralLog {
     n := &PB_GeneralLog{
      Id: int64 ( o.Id ),
      ToUserId: int32 ( o.ToUserId ),
      TargetId: int32 ( o.TargetId ),
      LogTypeId: int32 ( o.LogTypeId ),
      ExtraPB: []byte ( o.ExtraPB ),
      ExtraJson: string ( o.ExtraJson ),
      CreatedMs: int64 ( o.CreatedMs ),
    }
    return n
}
*/
/*
func PBConvPB__Group_To_Group( o *PB_Group) *Group {
     n := &Group{
      GroupId: int ( o.GroupId ),
      GroupName: string ( o.GroupName ),
      MembersCount: int ( o.MembersCount ),
      GroupPrivacyEnum: int ( o.GroupPrivacyEnum ),
      CreatorUserId: int ( o.CreatorUserId ),
      CreatedTime: int ( o.CreatedTime ),
      UpdatedMs: int ( o.UpdatedMs ),
      CurrentSeq: int ( o.CurrentSeq ),
    }
    return n
}

func PBConvPB_Group_To_Group ( o *Group) *PB_Group {
     n := &PB_Group{
      GroupId: int64 ( o.GroupId ),
      GroupName: string ( o.GroupName ),
      MembersCount: int32 ( o.MembersCount ),
      GroupPrivacyEnum: int32 ( o.GroupPrivacyEnum ),
      CreatorUserId: int32 ( o.CreatorUserId ),
      CreatedTime: int32 ( o.CreatedTime ),
      UpdatedMs: int64 ( o.UpdatedMs ),
      CurrentSeq: int32 ( o.CurrentSeq ),
    }
    return n
}
*/
/*
func PBConvPB__GroupMember_To_GroupMember( o *PB_GroupMember) *GroupMember {
     n := &GroupMember{
      Id: int ( o.Id ),
      GroupId: int ( o.GroupId ),
      GroupKey: string ( o.GroupKey ),
      UserId: int ( o.UserId ),
      ByUserId: int ( o.ByUserId ),
      GroupRoleEnumId: int ( o.GroupRoleEnumId ),
      CreatedTime: int ( o.CreatedTime ),
    }
    return n
}

func PBConvPB_GroupMember_To_GroupMember ( o *GroupMember) *PB_GroupMember {
     n := &PB_GroupMember{
      Id: int64 ( o.Id ),
      GroupId: int64 ( o.GroupId ),
      GroupKey: string ( o.GroupKey ),
      UserId: int32 ( o.UserId ),
      ByUserId: int32 ( o.ByUserId ),
      GroupRoleEnumId: int32 ( o.GroupRoleEnumId ),
      CreatedTime: int32 ( o.CreatedTime ),
    }
    return n
}
*/
/*
func PBConvPB__GroupMessage_To_GroupMessage( o *PB_GroupMessage) *GroupMessage {
     n := &GroupMessage{
      MessageId: int ( o.MessageId ),
      RoomKey: string ( o.RoomKey ),
      UserId: int ( o.UserId ),
      MessageFileId: int ( o.MessageFileId ),
      MessageTypeEnum: int ( o.MessageTypeEnum ),
      Text: string ( o.Text ),
      CreatedMs: int ( o.CreatedMs ),
      DeliveryStatusEnum: int ( o.DeliveryStatusEnum ),
    }
    return n
}

func PBConvPB_GroupMessage_To_GroupMessage ( o *GroupMessage) *PB_GroupMessage {
     n := &PB_GroupMessage{
      MessageId: int64 ( o.MessageId ),
      RoomKey: string ( o.RoomKey ),
      UserId: int32 ( o.UserId ),
      MessageFileId: int64 ( o.MessageFileId ),
      MessageTypeEnum: int32 ( o.MessageTypeEnum ),
      Text: string ( o.Text ),
      CreatedMs: int64 ( o.CreatedMs ),
      DeliveryStatusEnum: int32 ( o.DeliveryStatusEnum ),
    }
    return n
}
*/
/*
func PBConvPB__Like_To_Like( o *PB_Like) *Like {
     n := &Like{
      Id: int ( o.Id ),
      PostId: int ( o.PostId ),
      PostTypeEnum: int ( o.PostTypeEnum ),
      UserId: int ( o.UserId ),
      LikeEnum: int ( o.LikeEnum ),
      CreatedTime: int ( o.CreatedTime ),
    }
    return n
}

func PBConvPB_Like_To_Like ( o *Like) *PB_Like {
     n := &PB_Like{
      Id: int64 ( o.Id ),
      PostId: int64 ( o.PostId ),
      PostTypeEnum: int32 ( o.PostTypeEnum ),
      UserId: int32 ( o.UserId ),
      LikeEnum: int32 ( o.LikeEnum ),
      CreatedTime: int32 ( o.CreatedTime ),
    }
    return n
}
*/
/*
func PBConvPB__Media_To_Media( o *PB_Media) *Media {
     n := &Media{
      MediaId: int ( o.MediaId ),
      UserId: int ( o.UserId ),
      PostId: int ( o.PostId ),
      AlbumId: int ( o.AlbumId ),
      MediaTypeEnum: int ( o.MediaTypeEnum ),
      Width: int ( o.Width ),
      Height: int ( o.Height ),
      Size: int ( o.Size ),
      Duration: int ( o.Duration ),
      Md5Hash: string ( o.Md5Hash ),
      Color: string ( o.Color ),
      CreatedTime: int ( o.CreatedTime ),
    }
    return n
}

func PBConvPB_Media_To_Media ( o *Media) *PB_Media {
     n := &PB_Media{
      MediaId: int64 ( o.MediaId ),
      UserId: int32 ( o.UserId ),
      PostId: int32 ( o.PostId ),
      AlbumId: int32 ( o.AlbumId ),
      MediaTypeEnum: int32 ( o.MediaTypeEnum ),
      Width: int32 ( o.Width ),
      Height: int32 ( o.Height ),
      Size: int32 ( o.Size ),
      Duration: int32 ( o.Duration ),
      Md5Hash: string ( o.Md5Hash ),
      Color: string ( o.Color ),
      CreatedTime: int32 ( o.CreatedTime ),
    }
    return n
}
*/
/*
func PBConvPB__MessageFile_To_MessageFile( o *PB_MessageFile) *MessageFile {
     n := &MessageFile{
      MessageFileId: int ( o.MessageFileId ),
      MessageFileKey: string ( o.MessageFileKey ),
      UserId: int ( o.UserId ),
      Title: string ( o.Title ),
      Size: int ( o.Size ),
      FileTypeEnum: int ( o.FileTypeEnum ),
      Width: int ( o.Width ),
      Height: int ( o.Height ),
      Duration: int ( o.Duration ),
      Extension: string ( o.Extension ),
      Md5Hash: string ( o.Md5Hash ),
      CreatedTime: int ( o.CreatedTime ),
    }
    return n
}

func PBConvPB_MessageFile_To_MessageFile ( o *MessageFile) *PB_MessageFile {
     n := &PB_MessageFile{
      MessageFileId: int64 ( o.MessageFileId ),
      MessageFileKey: string ( o.MessageFileKey ),
      UserId: int32 ( o.UserId ),
      Title: string ( o.Title ),
      Size: int32 ( o.Size ),
      FileTypeEnum: int32 ( o.FileTypeEnum ),
      Width: int32 ( o.Width ),
      Height: int32 ( o.Height ),
      Duration: int32 ( o.Duration ),
      Extension: string ( o.Extension ),
      Md5Hash: string ( o.Md5Hash ),
      CreatedTime: int32 ( o.CreatedTime ),
    }
    return n
}
*/
/*
func PBConvPB__Notify_To_Notify( o *PB_Notify) *Notify {
     n := &Notify{
      NotifyId: int ( o.NotifyId ),
      ForUserId: int ( o.ForUserId ),
      ActorUserId: int ( o.ActorUserId ),
      NotiyTypeEnum: int ( o.NotiyTypeEnum ),
      PostId: int ( o.PostId ),
      CommentId: int ( o.CommentId ),
      PeerUserId: int ( o.PeerUserId ),
      Murmur64Hash: int ( o.Murmur64Hash ),
      SeenStatus: int ( o.SeenStatus ),
      CreatedTime: int ( o.CreatedTime ),
      Seq: int ( o.Seq ),
    }
    return n
}

func PBConvPB_Notify_To_Notify ( o *Notify) *PB_Notify {
     n := &PB_Notify{
      NotifyId: int64 ( o.NotifyId ),
      ForUserId: int32 ( o.ForUserId ),
      ActorUserId: int32 ( o.ActorUserId ),
      NotiyTypeEnum: int32 ( o.NotiyTypeEnum ),
      PostId: int64 ( o.PostId ),
      CommentId: int64 ( o.CommentId ),
      PeerUserId: int32 ( o.PeerUserId ),
      Murmur64Hash: int64 ( o.Murmur64Hash ),
      SeenStatus: int32 ( o.SeenStatus ),
      CreatedTime: int32 ( o.CreatedTime ),
      Seq: int32 ( o.Seq ),
    }
    return n
}
*/
/*
func PBConvPB__NotifyRemoved_To_NotifyRemoved( o *PB_NotifyRemoved) *NotifyRemoved {
     n := &NotifyRemoved{
      Murmur64Hash: int ( o.Murmur64Hash ),
      ForUserId: int ( o.ForUserId ),
      Id: int ( o.Id ),
    }
    return n
}

func PBConvPB_NotifyRemoved_To_NotifyRemoved ( o *NotifyRemoved) *PB_NotifyRemoved {
     n := &PB_NotifyRemoved{
      Murmur64Hash: int64 ( o.Murmur64Hash ),
      ForUserId: int32 ( o.ForUserId ),
      Id: int64 ( o.Id ),
    }
    return n
}
*/
/*
func PBConvPB__PhoneContact_To_PhoneContact( o *PB_PhoneContact) *PhoneContact {
     n := &PhoneContact{
      Id: int ( o.Id ),
      PhoneDisplayName: string ( o.PhoneDisplayName ),
      PhoneFamilyName: string ( o.PhoneFamilyName ),
      PhoneNumber: string ( o.PhoneNumber ),
      PhoneNormalizedNumber: string ( o.PhoneNormalizedNumber ),
      PhoneContactRowId: int ( o.PhoneContactRowId ),
      UserId: int ( o.UserId ),
      DeviceUuidId: int ( o.DeviceUuidId ),
      CreatedTime: int ( o.CreatedTime ),
      UpdatedTime: int ( o.UpdatedTime ),
    }
    return n
}

func PBConvPB_PhoneContact_To_PhoneContact ( o *PhoneContact) *PB_PhoneContact {
     n := &PB_PhoneContact{
      Id: int32 ( o.Id ),
      PhoneDisplayName: string ( o.PhoneDisplayName ),
      PhoneFamilyName: string ( o.PhoneFamilyName ),
      PhoneNumber: string ( o.PhoneNumber ),
      PhoneNormalizedNumber: string ( o.PhoneNormalizedNumber ),
      PhoneContactRowId: int32 ( o.PhoneContactRowId ),
      UserId: int32 ( o.UserId ),
      DeviceUuidId: int32 ( o.DeviceUuidId ),
      CreatedTime: int32 ( o.CreatedTime ),
      UpdatedTime: int32 ( o.UpdatedTime ),
    }
    return n
}
*/
/*
func PBConvPB__Post_To_Post( o *PB_Post) *Post {
     n := &Post{
      PostId: int ( o.PostId ),
      UserId: int ( o.UserId ),
      PostTypeEnum: int ( o.PostTypeEnum ),
      MediaId: int ( o.MediaId ),
      Text: string ( o.Text ),
      RichText: string ( o.RichText ),
      MediaCount: int ( o.MediaCount ),
      SharedTo: int ( o.SharedTo ),
      DisableComment: int ( o.DisableComment ),
      HasTag: int ( o.HasTag ),
      CommentsCount: int ( o.CommentsCount ),
      LikesCount: int ( o.LikesCount ),
      ViewsCount: int ( o.ViewsCount ),
      EditedTime: int ( o.EditedTime ),
      CreatedTime: int ( o.CreatedTime ),
      ReSharedPostId: int ( o.ReSharedPostId ),
    }
    return n
}

func PBConvPB_Post_To_Post ( o *Post) *PB_Post {
     n := &PB_Post{
      PostId: int64 ( o.PostId ),
      UserId: int32 ( o.UserId ),
      PostTypeEnum: int32 ( o.PostTypeEnum ),
      MediaId: int32 ( o.MediaId ),
      Text: string ( o.Text ),
      RichText: string ( o.RichText ),
      MediaCount: int32 ( o.MediaCount ),
      SharedTo: int32 ( o.SharedTo ),
      DisableComment: int32 ( o.DisableComment ),
      HasTag: int32 ( o.HasTag ),
      CommentsCount: int32 ( o.CommentsCount ),
      LikesCount: int32 ( o.LikesCount ),
      ViewsCount: int32 ( o.ViewsCount ),
      EditedTime: int32 ( o.EditedTime ),
      CreatedTime: int32 ( o.CreatedTime ),
      ReSharedPostId: int64 ( o.ReSharedPostId ),
    }
    return n
}
*/
/*
func PBConvPB__RecommendUser_To_RecommendUser( o *PB_RecommendUser) *RecommendUser {
     n := &RecommendUser{
      Id: int ( o.Id ),
      UserId: int ( o.UserId ),
      TargetId: int ( o.TargetId ),
      Weight: float32 ( o.Weight ),
      CreatedTime: int ( o.CreatedTime ),
    }
    return n
}

func PBConvPB_RecommendUser_To_RecommendUser ( o *RecommendUser) *PB_RecommendUser {
     n := &PB_RecommendUser{
      Id: int32 ( o.Id ),
      UserId: int32 ( o.UserId ),
      TargetId: int32 ( o.TargetId ),
      Weight: float32 ( o.Weight ),
      CreatedTime: int32 ( o.CreatedTime ),
    }
    return n
}
*/
/*
func PBConvPB__SearchClicked_To_SearchClicked( o *PB_SearchClicked) *SearchClicked {
     n := &SearchClicked{
      Id: int ( o.Id ),
      Query: string ( o.Query ),
      ClickType: int ( o.ClickType ),
      TargetId: int ( o.TargetId ),
      UserId: int ( o.UserId ),
      CreatedTime: int ( o.CreatedTime ),
    }
    return n
}

func PBConvPB_SearchClicked_To_SearchClicked ( o *SearchClicked) *PB_SearchClicked {
     n := &PB_SearchClicked{
      Id: int64 ( o.Id ),
      Query: string ( o.Query ),
      ClickType: int32 ( o.ClickType ),
      TargetId: int32 ( o.TargetId ),
      UserId: int32 ( o.UserId ),
      CreatedTime: int32 ( o.CreatedTime ),
    }
    return n
}
*/
/*
func PBConvPB__Session_To_Session( o *PB_Session) *Session {
     n := &Session{
      Id: int ( o.Id ),
      UserId: int ( o.UserId ),
      SessionUuid: string ( o.SessionUuid ),
      ClientUuid: string ( o.ClientUuid ),
      DeviceUuid: string ( o.DeviceUuid ),
      LastActivityTime: int ( o.LastActivityTime ),
      LastIpAddress: string ( o.LastIpAddress ),
      LastWifiMacAddress: string ( o.LastWifiMacAddress ),
      LastNetworkType: string ( o.LastNetworkType ),
      LastNetworkTypeEnumId: int ( o.LastNetworkTypeEnumId ),
      AppVersion: int ( o.AppVersion ),
      UpdatedTime: int ( o.UpdatedTime ),
      CreatedTime: int ( o.CreatedTime ),
    }
    return n
}

func PBConvPB_Session_To_Session ( o *Session) *PB_Session {
     n := &PB_Session{
      Id: int64 ( o.Id ),
      UserId: int32 ( o.UserId ),
      SessionUuid: string ( o.SessionUuid ),
      ClientUuid: string ( o.ClientUuid ),
      DeviceUuid: string ( o.DeviceUuid ),
      LastActivityTime: int32 ( o.LastActivityTime ),
      LastIpAddress: string ( o.LastIpAddress ),
      LastWifiMacAddress: string ( o.LastWifiMacAddress ),
      LastNetworkType: string ( o.LastNetworkType ),
      LastNetworkTypeEnumId: int32 ( o.LastNetworkTypeEnumId ),
      AppVersion: int32 ( o.AppVersion ),
      UpdatedTime: int32 ( o.UpdatedTime ),
      CreatedTime: int32 ( o.CreatedTime ),
    }
    return n
}
*/
/*
func PBConvPB__SettingClient_To_SettingClient( o *PB_SettingClient) *SettingClient {
     n := &SettingClient{
      UserId: int ( o.UserId ),
      AutoDownloadWifiVoice: int ( o.AutoDownloadWifiVoice ),
      AutoDownloadWifiImage: int ( o.AutoDownloadWifiImage ),
      AutoDownloadWifiVideo: int ( o.AutoDownloadWifiVideo ),
      AutoDownloadWifiFile: int ( o.AutoDownloadWifiFile ),
      AutoDownloadWifiMusic: int ( o.AutoDownloadWifiMusic ),
      AutoDownloadWifiGif: int ( o.AutoDownloadWifiGif ),
      AutoDownloadCellularVoice: int ( o.AutoDownloadCellularVoice ),
      AutoDownloadCellularImage: int ( o.AutoDownloadCellularImage ),
      AutoDownloadCellularVideo: int ( o.AutoDownloadCellularVideo ),
      AutoDownloadCellularFile: int ( o.AutoDownloadCellularFile ),
      AutoDownloadCellularMusic: int ( o.AutoDownloadCellularMusic ),
      AutoDownloadCellularGif: int ( o.AutoDownloadCellularGif ),
      AutoDownloadRoamingVoice: int ( o.AutoDownloadRoamingVoice ),
      AutoDownloadRoamingImage: int ( o.AutoDownloadRoamingImage ),
      AutoDownloadRoamingVideo: int ( o.AutoDownloadRoamingVideo ),
      AutoDownloadRoamingFile: int ( o.AutoDownloadRoamingFile ),
      AutoDownloadRoamingMusic: int ( o.AutoDownloadRoamingMusic ),
      AutoDownloadRoamingGif: int ( o.AutoDownloadRoamingGif ),
      SaveToGallery: int ( o.SaveToGallery ),
    }
    return n
}

func PBConvPB_SettingClient_To_SettingClient ( o *SettingClient) *PB_SettingClient {
     n := &PB_SettingClient{
      UserId: int32 ( o.UserId ),
      AutoDownloadWifiVoice: int32 ( o.AutoDownloadWifiVoice ),
      AutoDownloadWifiImage: int32 ( o.AutoDownloadWifiImage ),
      AutoDownloadWifiVideo: int32 ( o.AutoDownloadWifiVideo ),
      AutoDownloadWifiFile: int32 ( o.AutoDownloadWifiFile ),
      AutoDownloadWifiMusic: int32 ( o.AutoDownloadWifiMusic ),
      AutoDownloadWifiGif: int32 ( o.AutoDownloadWifiGif ),
      AutoDownloadCellularVoice: int32 ( o.AutoDownloadCellularVoice ),
      AutoDownloadCellularImage: int32 ( o.AutoDownloadCellularImage ),
      AutoDownloadCellularVideo: int32 ( o.AutoDownloadCellularVideo ),
      AutoDownloadCellularFile: int32 ( o.AutoDownloadCellularFile ),
      AutoDownloadCellularMusic: int32 ( o.AutoDownloadCellularMusic ),
      AutoDownloadCellularGif: int32 ( o.AutoDownloadCellularGif ),
      AutoDownloadRoamingVoice: int32 ( o.AutoDownloadRoamingVoice ),
      AutoDownloadRoamingImage: int32 ( o.AutoDownloadRoamingImage ),
      AutoDownloadRoamingVideo: int32 ( o.AutoDownloadRoamingVideo ),
      AutoDownloadRoamingFile: int32 ( o.AutoDownloadRoamingFile ),
      AutoDownloadRoamingMusic: int32 ( o.AutoDownloadRoamingMusic ),
      AutoDownloadRoamingGif: int32 ( o.AutoDownloadRoamingGif ),
      SaveToGallery: int32 ( o.SaveToGallery ),
    }
    return n
}
*/
/*
func PBConvPB__SettingNotification_To_SettingNotification( o *PB_SettingNotification) *SettingNotification {
     n := &SettingNotification{
      UserId: int ( o.UserId ),
      SocialLedOn: int ( o.SocialLedOn ),
      SocialLedColor: string ( o.SocialLedColor ),
      ReqestToFollowYou: int ( o.ReqestToFollowYou ),
      FollowedYou: int ( o.FollowedYou ),
      AccptedYourFollowRequest: int ( o.AccptedYourFollowRequest ),
      YourPostLiked: int ( o.YourPostLiked ),
      YourPostCommented: int ( o.YourPostCommented ),
      MenthenedYouInPost: int ( o.MenthenedYouInPost ),
      MenthenedYouInComment: int ( o.MenthenedYouInComment ),
      YourContactsJoined: int ( o.YourContactsJoined ),
      DirectMessage: int ( o.DirectMessage ),
      DirectAlert: int ( o.DirectAlert ),
      DirectPerview: int ( o.DirectPerview ),
      DirectLedOn: int ( o.DirectLedOn ),
      DirectLedColor: int ( o.DirectLedColor ),
      DirectVibrate: int ( o.DirectVibrate ),
      DirectPopup: int ( o.DirectPopup ),
      DirectSound: int ( o.DirectSound ),
      DirectPriority: int ( o.DirectPriority ),
    }
    return n
}

func PBConvPB_SettingNotification_To_SettingNotification ( o *SettingNotification) *PB_SettingNotification {
     n := &PB_SettingNotification{
      UserId: int32 ( o.UserId ),
      SocialLedOn: int32 ( o.SocialLedOn ),
      SocialLedColor: string ( o.SocialLedColor ),
      ReqestToFollowYou: int32 ( o.ReqestToFollowYou ),
      FollowedYou: int32 ( o.FollowedYou ),
      AccptedYourFollowRequest: int32 ( o.AccptedYourFollowRequest ),
      YourPostLiked: int32 ( o.YourPostLiked ),
      YourPostCommented: int32 ( o.YourPostCommented ),
      MenthenedYouInPost: int32 ( o.MenthenedYouInPost ),
      MenthenedYouInComment: int32 ( o.MenthenedYouInComment ),
      YourContactsJoined: int32 ( o.YourContactsJoined ),
      DirectMessage: int32 ( o.DirectMessage ),
      DirectAlert: int32 ( o.DirectAlert ),
      DirectPerview: int32 ( o.DirectPerview ),
      DirectLedOn: int32 ( o.DirectLedOn ),
      DirectLedColor: int32 ( o.DirectLedColor ),
      DirectVibrate: int32 ( o.DirectVibrate ),
      DirectPopup: int32 ( o.DirectPopup ),
      DirectSound: int32 ( o.DirectSound ),
      DirectPriority: int32 ( o.DirectPriority ),
    }
    return n
}
*/
/*
func PBConvPB__Tag_To_Tag( o *PB_Tag) *Tag {
     n := &Tag{
      TagId: int ( o.TagId ),
      Name: string ( o.Name ),
      Count: int ( o.Count ),
      TagStatusEnum: int ( o.TagStatusEnum ),
      CreatedTime: int ( o.CreatedTime ),
    }
    return n
}

func PBConvPB_Tag_To_Tag ( o *Tag) *PB_Tag {
     n := &PB_Tag{
      TagId: int64 ( o.TagId ),
      Name: string ( o.Name ),
      Count: int32 ( o.Count ),
      TagStatusEnum: int32 ( o.TagStatusEnum ),
      CreatedTime: int32 ( o.CreatedTime ),
    }
    return n
}
*/
/*
func PBConvPB__TagsPost_To_TagsPost( o *PB_TagsPost) *TagsPost {
     n := &TagsPost{
      Id: int ( o.Id ),
      TagId: int ( o.TagId ),
      PostId: int ( o.PostId ),
      PostTypeEnum: int ( o.PostTypeEnum ),
      CreatedTime: int ( o.CreatedTime ),
    }
    return n
}

func PBConvPB_TagsPost_To_TagsPost ( o *TagsPost) *PB_TagsPost {
     n := &PB_TagsPost{
      Id: int64 ( o.Id ),
      TagId: int32 ( o.TagId ),
      PostId: int32 ( o.PostId ),
      PostTypeEnum: int32 ( o.PostTypeEnum ),
      CreatedTime: int32 ( o.CreatedTime ),
    }
    return n
}
*/
/*
func PBConvPB__TriggerLog_To_TriggerLog( o *PB_TriggerLog) *TriggerLog {
     n := &TriggerLog{
      Id: int ( o.Id ),
      ModelName: string ( o.ModelName ),
      ChangeType: string ( o.ChangeType ),
      TargetId: int ( o.TargetId ),
      TargetStr: string ( o.TargetStr ),
      CreatedSe: int ( o.CreatedSe ),
    }
    return n
}

func PBConvPB_TriggerLog_To_TriggerLog ( o *TriggerLog) *PB_TriggerLog {
     n := &PB_TriggerLog{
      Id: int64 ( o.Id ),
      ModelName: string ( o.ModelName ),
      ChangeType: string ( o.ChangeType ),
      TargetId: int64 ( o.TargetId ),
      TargetStr: string ( o.TargetStr ),
      CreatedSe: int32 ( o.CreatedSe ),
    }
    return n
}
*/
/*
func PBConvPB__User_To_User( o *PB_User) *User {
     n := &User{
      UserId: int ( o.UserId ),
      UserName: string ( o.UserName ),
      UserNameLower: string ( o.UserNameLower ),
      FirstName: string ( o.FirstName ),
      LastName: string ( o.LastName ),
      UserTypeEnum: int ( o.UserTypeEnum ),
      UserLevelEnum: int ( o.UserLevelEnum ),
      AvatarId: int ( o.AvatarId ),
      ProfilePrivacyEnum: int ( o.ProfilePrivacyEnum ),
      Phone: int ( o.Phone ),
      About: string ( o.About ),
      Email: string ( o.Email ),
      PasswordHash: string ( o.PasswordHash ),
      PasswordSalt: string ( o.PasswordSalt ),
      FollowersCount: int ( o.FollowersCount ),
      FollowingCount: int ( o.FollowingCount ),
      PostsCount: int ( o.PostsCount ),
      MediaCount: int ( o.MediaCount ),
      LikesCount: int ( o.LikesCount ),
      ResharedCount: int ( o.ResharedCount ),
      LastActionTime: int ( o.LastActionTime ),
      LastPostTime: int ( o.LastPostTime ),
      PrimaryFollowingList: int ( o.PrimaryFollowingList ),
      CreatedSe: int ( o.CreatedSe ),
      UpdatedMs: int ( o.UpdatedMs ),
      OnlinePrivacyEnum: int ( o.OnlinePrivacyEnum ),
      LastActivityTime: int ( o.LastActivityTime ),
      Phone2: string ( o.Phone2 ),
    }
    return n
}

func PBConvPB_User_To_User ( o *User) *PB_User {
     n := &PB_User{
      UserId: int32 ( o.UserId ),
      UserName: string ( o.UserName ),
      UserNameLower: string ( o.UserNameLower ),
      FirstName: string ( o.FirstName ),
      LastName: string ( o.LastName ),
      UserTypeEnum: int32 ( o.UserTypeEnum ),
      UserLevelEnum: int32 ( o.UserLevelEnum ),
      AvatarId: int64 ( o.AvatarId ),
      ProfilePrivacyEnum: int32 ( o.ProfilePrivacyEnum ),
      Phone: int64 ( o.Phone ),
      About: string ( o.About ),
      Email: string ( o.Email ),
      PasswordHash: string ( o.PasswordHash ),
      PasswordSalt: string ( o.PasswordSalt ),
      FollowersCount: int32 ( o.FollowersCount ),
      FollowingCount: int32 ( o.FollowingCount ),
      PostsCount: int32 ( o.PostsCount ),
      MediaCount: int32 ( o.MediaCount ),
      LikesCount: int32 ( o.LikesCount ),
      ResharedCount: int32 ( o.ResharedCount ),
      LastActionTime: int32 ( o.LastActionTime ),
      LastPostTime: int32 ( o.LastPostTime ),
      PrimaryFollowingList: int32 ( o.PrimaryFollowingList ),
      CreatedSe: int32 ( o.CreatedSe ),
      UpdatedMs: int64 ( o.UpdatedMs ),
      OnlinePrivacyEnum: int32 ( o.OnlinePrivacyEnum ),
      LastActivityTime: int32 ( o.LastActivityTime ),
      Phone2: string ( o.Phone2 ),
    }
    return n
}
*/
/*
func PBConvPB__UserMetaInfo_To_UserMetaInfo( o *PB_UserMetaInfo) *UserMetaInfo {
     n := &UserMetaInfo{
      Id: int ( o.Id ),
      UserId: int ( o.UserId ),
      IsNotificationDirty: int ( o.IsNotificationDirty ),
      LastUserRecGen: int ( o.LastUserRecGen ),
    }
    return n
}

func PBConvPB_UserMetaInfo_To_UserMetaInfo ( o *UserMetaInfo) *PB_UserMetaInfo {
     n := &PB_UserMetaInfo{
      Id: int32 ( o.Id ),
      UserId: int32 ( o.UserId ),
      IsNotificationDirty: int32 ( o.IsNotificationDirty ),
      LastUserRecGen: int32 ( o.LastUserRecGen ),
    }
    return n
}
*/
/*
func PBConvPB__UserPassword_To_UserPassword( o *PB_UserPassword) *UserPassword {
     n := &UserPassword{
      UserId: int ( o.UserId ),
      Password: string ( o.Password ),
      CreatedTime: int ( o.CreatedTime ),
    }
    return n
}

func PBConvPB_UserPassword_To_UserPassword ( o *UserPassword) *PB_UserPassword {
     n := &PB_UserPassword{
      UserId: int32 ( o.UserId ),
      Password: string ( o.Password ),
      CreatedTime: int32 ( o.CreatedTime ),
    }
    return n
}
*/
