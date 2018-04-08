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
    }
    return n
}

func PBConvPB_Comment_To_Comment ( o *Comment) *PB_Comment {
     n := &PB_Comment{
      CommentId: int64 ( o.CommentId ),
      UserId: int32 ( o.UserId ),
      PostId: int64 ( o.PostId ),
      Text: string ( o.Text ),
      LikesCount: int32 ( o.LikesCount ),
      CreatedTime: int32 ( o.CreatedTime ),
    }
    return n
}
*/
/*
func PBConvPB__CommentDeleted_To_CommentDeleted( o *PB_CommentDeleted) *CommentDeleted {
     n := &CommentDeleted{
      CommentId: int ( o.CommentId ),
      UserId: int ( o.UserId ),
    }
    return n
}

func PBConvPB_CommentDeleted_To_CommentDeleted ( o *CommentDeleted) *PB_CommentDeleted {
     n := &PB_CommentDeleted{
      CommentId: int64 ( o.CommentId ),
      UserId: int32 ( o.UserId ),
    }
    return n
}
*/
/*
func PBConvPB__Event_To_Event( o *PB_Event) *Event {
     n := &Event{
      EventId: int ( o.EventId ),
      EventType: int ( o.EventType ),
      ByUserId: int ( o.ByUserId ),
      PeerUserId: int ( o.PeerUserId ),
      PostId: int ( o.PostId ),
      CommentId: int ( o.CommentId ),
      ActionId: int ( o.ActionId ),
      Murmur64Hash: int ( o.Murmur64Hash ),
      ChatKey: string ( o.ChatKey ),
      MessageId: int ( o.MessageId ),
      ReSharedId: int ( o.ReSharedId ),
    }
    return n
}

func PBConvPB_Event_To_Event ( o *Event) *PB_Event {
     n := &PB_Event{
      EventId: int64 ( o.EventId ),
      EventType: int32 ( o.EventType ),
      ByUserId: int64 ( o.ByUserId ),
      PeerUserId: int64 ( o.PeerUserId ),
      PostId: int64 ( o.PostId ),
      CommentId: int64 ( o.CommentId ),
      ActionId: int64 ( o.ActionId ),
      Murmur64Hash: int64 ( o.Murmur64Hash ),
      ChatKey: string ( o.ChatKey ),
      MessageId: int64 ( o.MessageId ),
      ReSharedId: int64 ( o.ReSharedId ),
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
      CreatedTime: int ( o.CreatedTime ),
    }
    return n
}

func PBConvPB_FollowingListMember_To_FollowingListMember ( o *FollowingListMember) *PB_FollowingListMember {
     n := &PB_FollowingListMember{
      Id: int64 ( o.Id ),
      ListId: int32 ( o.ListId ),
      UserId: int32 ( o.UserId ),
      FollowedUserId: int32 ( o.FollowedUserId ),
      CreatedTime: int32 ( o.CreatedTime ),
    }
    return n
}
*/
/*
func PBConvPB__FollowingListMemberRemoved_To_FollowingListMemberRemoved( o *PB_FollowingListMemberRemoved) *FollowingListMemberRemoved {
     n := &FollowingListMemberRemoved{
      Id: int ( o.Id ),
      ListId: int ( o.ListId ),
      UserId: int ( o.UserId ),
      UnFollowedUserId: int ( o.UnFollowedUserId ),
      UpdatedTime: int ( o.UpdatedTime ),
    }
    return n
}

func PBConvPB_FollowingListMemberRemoved_To_FollowingListMemberRemoved ( o *FollowingListMemberRemoved) *PB_FollowingListMemberRemoved {
     n := &PB_FollowingListMemberRemoved{
      Id: int64 ( o.Id ),
      ListId: int32 ( o.ListId ),
      UserId: int32 ( o.UserId ),
      UnFollowedUserId: int32 ( o.UnFollowedUserId ),
      UpdatedTime: int32 ( o.UpdatedTime ),
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
func PBConvPB__Notify_To_Notify( o *PB_Notify) *Notify {
     n := &Notify{
      NotifyId: int ( o.NotifyId ),
      ForUserId: int ( o.ForUserId ),
      ActorUserId: int ( o.ActorUserId ),
      NotifyTypeEnum: int ( o.NotifyTypeEnum ),
      PostId: int ( o.PostId ),
      CommentId: int ( o.CommentId ),
      PeerUserId: int ( o.PeerUserId ),
      Murmur64Hash: int ( o.Murmur64Hash ),
      SeenStatus: int ( o.SeenStatus ),
      CreatedTime: int ( o.CreatedTime ),
    }
    return n
}

func PBConvPB_Notify_To_Notify ( o *Notify) *PB_Notify {
     n := &PB_Notify{
      NotifyId: int64 ( o.NotifyId ),
      ForUserId: int32 ( o.ForUserId ),
      ActorUserId: int32 ( o.ActorUserId ),
      NotifyTypeEnum: int32 ( o.NotifyTypeEnum ),
      PostId: int64 ( o.PostId ),
      CommentId: int64 ( o.CommentId ),
      PeerUserId: int32 ( o.PeerUserId ),
      Murmur64Hash: int64 ( o.Murmur64Hash ),
      SeenStatus: int32 ( o.SeenStatus ),
      CreatedTime: int32 ( o.CreatedTime ),
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
      UserId: int ( o.UserId ),
      Phone: int ( o.Phone ),
      PhoneDisplayName: string ( o.PhoneDisplayName ),
      PhoneFamilyName: string ( o.PhoneFamilyName ),
      PhoneNumber: string ( o.PhoneNumber ),
      PhoneNormalizedNumber: string ( o.PhoneNormalizedNumber ),
      PhoneContactRowId: int ( o.PhoneContactRowId ),
      DeviceUuidId: int ( o.DeviceUuidId ),
      CreatedTime: int ( o.CreatedTime ),
    }
    return n
}

func PBConvPB_PhoneContact_To_PhoneContact ( o *PhoneContact) *PB_PhoneContact {
     n := &PB_PhoneContact{
      Id: int32 ( o.Id ),
      UserId: int32 ( o.UserId ),
      Phone: int64 ( o.Phone ),
      PhoneDisplayName: string ( o.PhoneDisplayName ),
      PhoneFamilyName: string ( o.PhoneFamilyName ),
      PhoneNumber: string ( o.PhoneNumber ),
      PhoneNormalizedNumber: string ( o.PhoneNormalizedNumber ),
      PhoneContactRowId: int32 ( o.PhoneContactRowId ),
      DeviceUuidId: int32 ( o.DeviceUuidId ),
      CreatedTime: int32 ( o.CreatedTime ),
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
      PostCategoryEnum: int ( o.PostCategoryEnum ),
      MediaId: int ( o.MediaId ),
      PostKey: string ( o.PostKey ),
      Text: string ( o.Text ),
      RichText: string ( o.RichText ),
      MediaCount: int ( o.MediaCount ),
      SharedTo: int ( o.SharedTo ),
      DisableComment: int ( o.DisableComment ),
      Source: int ( o.Source ),
      HasTag: int ( o.HasTag ),
      Seq: int ( o.Seq ),
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
      PostCategoryEnum: int32 ( o.PostCategoryEnum ),
      MediaId: int64 ( o.MediaId ),
      PostKey: string ( o.PostKey ),
      Text: string ( o.Text ),
      RichText: string ( o.RichText ),
      MediaCount: int32 ( o.MediaCount ),
      SharedTo: int32 ( o.SharedTo ),
      DisableComment: int32 ( o.DisableComment ),
      Source: int32 ( o.Source ),
      HasTag: int32 ( o.HasTag ),
      Seq: int32 ( o.Seq ),
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
func PBConvPB__PostCount_To_PostCount( o *PB_PostCount) *PostCount {
     n := &PostCount{
      PostId: int ( o.PostId ),
      ViewsCount: int ( o.ViewsCount ),
    }
    return n
}

func PBConvPB_PostCount_To_PostCount ( o *PostCount) *PB_PostCount {
     n := &PB_PostCount{
      PostId: int64 ( o.PostId ),
      ViewsCount: int64 ( o.ViewsCount ),
    }
    return n
}
*/
/*
func PBConvPB__PostDeleted_To_PostDeleted( o *PB_PostDeleted) *PostDeleted {
     n := &PostDeleted{
      PostId: int ( o.PostId ),
      UserId: int ( o.UserId ),
    }
    return n
}

func PBConvPB_PostDeleted_To_PostDeleted ( o *PostDeleted) *PB_PostDeleted {
     n := &PB_PostDeleted{
      PostId: int64 ( o.PostId ),
      UserId: int32 ( o.UserId ),
    }
    return n
}
*/
/*
func PBConvPB__PostKey_To_PostKey( o *PB_PostKey) *PostKey {
     n := &PostKey{
      Id: int ( o.Id ),
      PostKeyStr: string ( o.PostKeyStr ),
      Used: int ( o.Used ),
    }
    return n
}

func PBConvPB_PostKey_To_PostKey ( o *PostKey) *PB_PostKey {
     n := &PB_PostKey{
      Id: int32 ( o.Id ),
      PostKeyStr: string ( o.PostKeyStr ),
      Used: int32 ( o.Used ),
    }
    return n
}
*/
/*
func PBConvPB__PostLink_To_PostLink( o *PB_PostLink) *PostLink {
     n := &PostLink{
      LinkId: int ( o.LinkId ),
      LinkUrl: string ( o.LinkUrl ),
    }
    return n
}

func PBConvPB_PostLink_To_PostLink ( o *PostLink) *PB_PostLink {
     n := &PB_PostLink{
      LinkId: int64 ( o.LinkId ),
      LinkUrl: string ( o.LinkUrl ),
    }
    return n
}
*/
/*
func PBConvPB__PostMedia_To_PostMedia( o *PB_PostMedia) *PostMedia {
     n := &PostMedia{
      MediaId: int ( o.MediaId ),
      UserId: int ( o.UserId ),
      PostId: int ( o.PostId ),
      AlbumId: int ( o.AlbumId ),
      MediaTypeEnum: int ( o.MediaTypeEnum ),
      Width: int ( o.Width ),
      Height: int ( o.Height ),
      Size: int ( o.Size ),
      Duration: int ( o.Duration ),
      Extension: string ( o.Extension ),
      Md5Hash: string ( o.Md5Hash ),
      Color: string ( o.Color ),
      CreatedTime: int ( o.CreatedTime ),
      ViewCount: int ( o.ViewCount ),
      Extra: string ( o.Extra ),
    }
    return n
}

func PBConvPB_PostMedia_To_PostMedia ( o *PostMedia) *PB_PostMedia {
     n := &PB_PostMedia{
      MediaId: int64 ( o.MediaId ),
      UserId: int32 ( o.UserId ),
      PostId: int64 ( o.PostId ),
      AlbumId: int32 ( o.AlbumId ),
      MediaTypeEnum: int32 ( o.MediaTypeEnum ),
      Width: int32 ( o.Width ),
      Height: int32 ( o.Height ),
      Size: int32 ( o.Size ),
      Duration: int32 ( o.Duration ),
      Extension: string ( o.Extension ),
      Md5Hash: string ( o.Md5Hash ),
      Color: string ( o.Color ),
      CreatedTime: int32 ( o.CreatedTime ),
      ViewCount: int32 ( o.ViewCount ),
      Extra: string ( o.Extra ),
    }
    return n
}
*/
/*
func PBConvPB__PostMentioned_To_PostMentioned( o *PB_PostMentioned) *PostMentioned {
     n := &PostMentioned{
      MentionedId: int ( o.MentionedId ),
      ForUserId: int ( o.ForUserId ),
      PostId: int ( o.PostId ),
      PostUserId: int ( o.PostUserId ),
      PostTypeEnum: int ( o.PostTypeEnum ),
      PostCategoryEnum: int ( o.PostCategoryEnum ),
      CreatedTime: int ( o.CreatedTime ),
    }
    return n
}

func PBConvPB_PostMentioned_To_PostMentioned ( o *PostMentioned) *PB_PostMentioned {
     n := &PB_PostMentioned{
      MentionedId: int64 ( o.MentionedId ),
      ForUserId: int32 ( o.ForUserId ),
      PostId: int64 ( o.PostId ),
      PostUserId: int32 ( o.PostUserId ),
      PostTypeEnum: int32 ( o.PostTypeEnum ),
      PostCategoryEnum: int32 ( o.PostCategoryEnum ),
      CreatedTime: int32 ( o.CreatedTime ),
    }
    return n
}
*/
/*
func PBConvPB__PostReshared_To_PostReshared( o *PB_PostReshared) *PostReshared {
     n := &PostReshared{
      ResharedId: int ( o.ResharedId ),
      ByUserId: int ( o.ByUserId ),
      PostId: int ( o.PostId ),
      PostUserId: int ( o.PostUserId ),
      PostTypeEnum: int ( o.PostTypeEnum ),
      PostCategoryEnum: int ( o.PostCategoryEnum ),
      CreatedTime: int ( o.CreatedTime ),
    }
    return n
}

func PBConvPB_PostReshared_To_PostReshared ( o *PostReshared) *PB_PostReshared {
     n := &PB_PostReshared{
      ResharedId: int64 ( o.ResharedId ),
      ByUserId: int32 ( o.ByUserId ),
      PostId: int64 ( o.PostId ),
      PostUserId: int32 ( o.PostUserId ),
      PostTypeEnum: int32 ( o.PostTypeEnum ),
      PostCategoryEnum: int32 ( o.PostCategoryEnum ),
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
      SessionUuid: string ( o.SessionUuid ),
      UserId: int ( o.UserId ),
      ClientUuid: string ( o.ClientUuid ),
      DeviceUuid: string ( o.DeviceUuid ),
      LastActivityTime: int ( o.LastActivityTime ),
      LastIpAddress: string ( o.LastIpAddress ),
      LastWifiMacAddress: string ( o.LastWifiMacAddress ),
      LastNetworkType: string ( o.LastNetworkType ),
      LastNetworkTypeEnum: int ( o.LastNetworkTypeEnum ),
      AppVersion: int ( o.AppVersion ),
      UpdatedTime: int ( o.UpdatedTime ),
      CreatedTime: int ( o.CreatedTime ),
    }
    return n
}

func PBConvPB_Session_To_Session ( o *Session) *PB_Session {
     n := &PB_Session{
      SessionUuid: string ( o.SessionUuid ),
      UserId: int32 ( o.UserId ),
      ClientUuid: string ( o.ClientUuid ),
      DeviceUuid: string ( o.DeviceUuid ),
      LastActivityTime: int32 ( o.LastActivityTime ),
      LastIpAddress: string ( o.LastIpAddress ),
      LastWifiMacAddress: string ( o.LastWifiMacAddress ),
      LastNetworkType: string ( o.LastNetworkType ),
      LastNetworkTypeEnum: int32 ( o.LastNetworkTypeEnum ),
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
func PBConvPB__TagPost_To_TagPost( o *PB_TagPost) *TagPost {
     n := &TagPost{
      Id: int ( o.Id ),
      TagId: int ( o.TagId ),
      PostId: int ( o.PostId ),
      PostTypeEnum: int ( o.PostTypeEnum ),
      PostCategoryEnum: int ( o.PostCategoryEnum ),
      CreatedTime: int ( o.CreatedTime ),
    }
    return n
}

func PBConvPB_TagPost_To_TagPost ( o *TagPost) *PB_TagPost {
     n := &PB_TagPost{
      Id: int64 ( o.Id ),
      TagId: int32 ( o.TagId ),
      PostId: int32 ( o.PostId ),
      PostTypeEnum: int32 ( o.PostTypeEnum ),
      PostCategoryEnum: int32 ( o.PostCategoryEnum ),
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
      PostSeq: int ( o.PostSeq ),
      FollowersCount: int ( o.FollowersCount ),
      FollowingCount: int ( o.FollowingCount ),
      PostsCount: int ( o.PostsCount ),
      MediaCount: int ( o.MediaCount ),
      PhotoCount: int ( o.PhotoCount ),
      VideoCount: int ( o.VideoCount ),
      GifCount: int ( o.GifCount ),
      AudioCount: int ( o.AudioCount ),
      VoiceCount: int ( o.VoiceCount ),
      FileCount: int ( o.FileCount ),
      LinkCount: int ( o.LinkCount ),
      BoardCount: int ( o.BoardCount ),
      PinedCount: int ( o.PinedCount ),
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
      PostSeq: int32 ( o.PostSeq ),
      FollowersCount: int32 ( o.FollowersCount ),
      FollowingCount: int32 ( o.FollowingCount ),
      PostsCount: int32 ( o.PostsCount ),
      MediaCount: int32 ( o.MediaCount ),
      PhotoCount: int32 ( o.PhotoCount ),
      VideoCount: int32 ( o.VideoCount ),
      GifCount: int32 ( o.GifCount ),
      AudioCount: int32 ( o.AudioCount ),
      VoiceCount: int32 ( o.VoiceCount ),
      FileCount: int32 ( o.FileCount ),
      LinkCount: int32 ( o.LinkCount ),
      BoardCount: int32 ( o.BoardCount ),
      PinedCount: int32 ( o.PinedCount ),
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
      Seq: int ( o.Seq ),
      SeenSeq: int ( o.SeenSeq ),
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
      Seq: int32 ( o.Seq ),
      SeenSeq: int32 ( o.SeenSeq ),
      UpdatedMs: int64 ( o.UpdatedMs ),
    }
    return n
}
*/
/*
func PBConvPB__ChatLastMessage_To_ChatLastMessage( o *PB_ChatLastMessage) *ChatLastMessage {
     n := &ChatLastMessage{
      ChatKey: string ( o.ChatKey ),
      ForUserId: int ( o.ForUserId ),
      LastMsgPb: []byte ( o.LastMsgPb ),
      LastMsgJson: string ( o.LastMsgJson ),
    }
    return n
}

func PBConvPB_ChatLastMessage_To_ChatLastMessage ( o *ChatLastMessage) *PB_ChatLastMessage {
     n := &PB_ChatLastMessage{
      ChatKey: string ( o.ChatKey ),
      ForUserId: int32 ( o.ForUserId ),
      LastMsgPb: []byte ( o.LastMsgPb ),
      LastMsgJson: string ( o.LastMsgJson ),
    }
    return n
}
*/
/*
func PBConvPB__ChatSync_To_ChatSync( o *PB_ChatSync) *ChatSync {
     n := &ChatSync{
      SyncId: int ( o.SyncId ),
      ToUserId: int ( o.ToUserId ),
      ChatSyncTypeId: int ( o.ChatSyncTypeId ),
      RoomKey: string ( o.RoomKey ),
      ChatKey: string ( o.ChatKey ),
      FromHighMessageId: int ( o.FromHighMessageId ),
      ToLowMessageId: int ( o.ToLowMessageId ),
      MessageId: int ( o.MessageId ),
      MessagePb: []byte ( o.MessagePb ),
      MessageJson: string ( o.MessageJson ),
      CreatedTime: int ( o.CreatedTime ),
    }
    return n
}

func PBConvPB_ChatSync_To_ChatSync ( o *ChatSync) *PB_ChatSync {
     n := &PB_ChatSync{
      SyncId: int64 ( o.SyncId ),
      ToUserId: int32 ( o.ToUserId ),
      ChatSyncTypeId: int32 ( o.ChatSyncTypeId ),
      RoomKey: string ( o.RoomKey ),
      ChatKey: string ( o.ChatKey ),
      FromHighMessageId: int64 ( o.FromHighMessageId ),
      ToLowMessageId: int64 ( o.ToLowMessageId ),
      MessageId: int64 ( o.MessageId ),
      MessagePb: []byte ( o.MessagePb ),
      MessageJson: string ( o.MessageJson ),
      CreatedTime: int32 ( o.CreatedTime ),
    }
    return n
}
*/
/*
func PBConvPB__DirectMessage_To_DirectMessage( o *PB_DirectMessage) *DirectMessage {
     n := &DirectMessage{
      ChatKey: string ( o.ChatKey ),
      MessageId: int ( o.MessageId ),
      RoomKey: string ( o.RoomKey ),
      UserId: int ( o.UserId ),
      MessageFileId: int ( o.MessageFileId ),
      MessageTypeEnum: int ( o.MessageTypeEnum ),
      Text: string ( o.Text ),
      CreatedTime: int ( o.CreatedTime ),
      Seq: int ( o.Seq ),
      DeliviryStatusEnum: int ( o.DeliviryStatusEnum ),
      ExtraPB: []byte ( o.ExtraPB ),
    }
    return n
}

func PBConvPB_DirectMessage_To_DirectMessage ( o *DirectMessage) *PB_DirectMessage {
     n := &PB_DirectMessage{
      ChatKey: string ( o.ChatKey ),
      MessageId: int64 ( o.MessageId ),
      RoomKey: string ( o.RoomKey ),
      UserId: int32 ( o.UserId ),
      MessageFileId: int64 ( o.MessageFileId ),
      MessageTypeEnum: int32 ( o.MessageTypeEnum ),
      Text: string ( o.Text ),
      CreatedTime: int32 ( o.CreatedTime ),
      Seq: int32 ( o.Seq ),
      DeliviryStatusEnum: int32 ( o.DeliviryStatusEnum ),
      ExtraPB: []byte ( o.ExtraPB ),
    }
    return n
}
*/
/*
func PBConvPB__Home_To_Home( o *PB_Home) *Home {
     n := &Home{
      Id: int ( o.Id ),
      ForUserId: int ( o.ForUserId ),
      PostId: int ( o.PostId ),
    }
    return n
}

func PBConvPB_Home_To_Home ( o *Home) *PB_Home {
     n := &PB_Home{
      Id: int64 ( o.Id ),
      ForUserId: int32 ( o.ForUserId ),
      PostId: int64 ( o.PostId ),
    }
    return n
}
*/
/*
func PBConvPB__MessageFile_To_MessageFile( o *PB_MessageFile) *MessageFile {
     n := &MessageFile{
      MessageFileId: int ( o.MessageFileId ),
      FileTypeEnum: int ( o.FileTypeEnum ),
      UserId: int ( o.UserId ),
      Title: string ( o.Title ),
      Size: int ( o.Size ),
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
      FileTypeEnum: int32 ( o.FileTypeEnum ),
      UserId: int32 ( o.UserId ),
      Title: string ( o.Title ),
      Size: int32 ( o.Size ),
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
func PBConvPB__FileMsg_To_FileMsg( o *PB_FileMsg) *FileMsg {
     n := &FileMsg{
      Id: int ( o.Id ),
      FileType: int ( o.FileType ),
      Extension: string ( o.Extension ),
      DataThumb: []byte ( o.DataThumb ),
      Data: []byte ( o.Data ),
    }
    return n
}

func PBConvPB_FileMsg_To_FileMsg ( o *FileMsg) *PB_FileMsg {
     n := &PB_FileMsg{
      Id: int64 ( o.Id ),
      FileType: int32 ( o.FileType ),
      Extension: string ( o.Extension ),
      DataThumb: []byte ( o.DataThumb ),
      Data: []byte ( o.Data ),
    }
    return n
}
*/
/*
func PBConvPB__FilePost_To_FilePost( o *PB_FilePost) *FilePost {
     n := &FilePost{
      Id: int ( o.Id ),
      FileType: int ( o.FileType ),
      Extension: string ( o.Extension ),
      DataThumb: []byte ( o.DataThumb ),
      Data: []byte ( o.Data ),
    }
    return n
}

func PBConvPB_FilePost_To_FilePost ( o *FilePost) *PB_FilePost {
     n := &PB_FilePost{
      Id: int64 ( o.Id ),
      FileType: int32 ( o.FileType ),
      Extension: string ( o.Extension ),
      DataThumb: []byte ( o.DataThumb ),
      Data: []byte ( o.Data ),
    }
    return n
}
*/
/*
func PBConvPB__ActionFanout_To_ActionFanout( o *PB_ActionFanout) *ActionFanout {
     n := &ActionFanout{
      OrderId: int ( o.OrderId ),
      ForUserId: int ( o.ForUserId ),
      ActionId: int ( o.ActionId ),
      ActorUserId: int ( o.ActorUserId ),
    }
    return n
}

func PBConvPB_ActionFanout_To_ActionFanout ( o *ActionFanout) *PB_ActionFanout {
     n := &PB_ActionFanout{
      OrderId: int64 ( o.OrderId ),
      ForUserId: int32 ( o.ForUserId ),
      ActionId: int64 ( o.ActionId ),
      ActorUserId: int32 ( o.ActorUserId ),
    }
    return n
}
*/
/*
func PBConvPB__HomeFanout_To_HomeFanout( o *PB_HomeFanout) *HomeFanout {
     n := &HomeFanout{
      OrderId: int ( o.OrderId ),
      ForUserId: int ( o.ForUserId ),
      PostId: int ( o.PostId ),
      PostUserId: int ( o.PostUserId ),
    }
    return n
}

func PBConvPB_HomeFanout_To_HomeFanout ( o *HomeFanout) *PB_HomeFanout {
     n := &PB_HomeFanout{
      OrderId: int64 ( o.OrderId ),
      ForUserId: int64 ( o.ForUserId ),
      PostId: int64 ( o.PostId ),
      PostUserId: int64 ( o.PostUserId ),
    }
    return n
}
*/
/*
func PBConvPB__SuggestedTopPost_To_SuggestedTopPost( o *PB_SuggestedTopPost) *SuggestedTopPost {
     n := &SuggestedTopPost{
      Id: int ( o.Id ),
      PostId: int ( o.PostId ),
    }
    return n
}

func PBConvPB_SuggestedTopPost_To_SuggestedTopPost ( o *SuggestedTopPost) *PB_SuggestedTopPost {
     n := &PB_SuggestedTopPost{
      Id: int64 ( o.Id ),
      PostId: int64 ( o.PostId ),
    }
    return n
}
*/
/*
func PBConvPB__SuggestedUser_To_SuggestedUser( o *PB_SuggestedUser) *SuggestedUser {
     n := &SuggestedUser{
      Id: int ( o.Id ),
      UserId: int ( o.UserId ),
      TargetId: int ( o.TargetId ),
      Weight: float32 ( o.Weight ),
      CreatedTime: int ( o.CreatedTime ),
    }
    return n
}

func PBConvPB_SuggestedUser_To_SuggestedUser ( o *SuggestedUser) *PB_SuggestedUser {
     n := &PB_SuggestedUser{
      Id: int32 ( o.Id ),
      UserId: int32 ( o.UserId ),
      TargetId: int32 ( o.TargetId ),
      Weight: float32 ( o.Weight ),
      CreatedTime: int32 ( o.CreatedTime ),
    }
    return n
}
*/
