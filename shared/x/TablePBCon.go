package x

/*
func PBConvPB__Action_To_Action( o *PB_Action) *Action {
     n := &Action{
      ActionId: int ( o.ActionId ),
      ActorUserId: int ( o.ActorUserId ),
      ActionType: int ( o.ActionType ),
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
      ActionType: int32 ( o.ActionType ),
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
func PBConvPB__Blocked_To_Blocked( o *PB_Blocked) *Blocked {
     n := &Blocked{
      Id: int ( o.Id ),
      UserId: int ( o.UserId ),
      BlockedUserId: int ( o.BlockedUserId ),
      CreatedTime: int ( o.CreatedTime ),
    }
    return n
}

func PBConvPB_Blocked_To_Blocked ( o *Blocked) *PB_Blocked {
     n := &PB_Blocked{
      Id: int64 ( o.Id ),
      UserId: int32 ( o.UserId ),
      BlockedUserId: int32 ( o.BlockedUserId ),
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
      IsEdited: int ( o.IsEdited ),
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
      IsEdited: int32 ( o.IsEdited ),
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
func PBConvPB__Followed_To_Followed( o *PB_Followed) *Followed {
     n := &Followed{
      Id: int ( o.Id ),
      UserId: int ( o.UserId ),
      FollowedUserId: int ( o.FollowedUserId ),
      CreatedTime: int ( o.CreatedTime ),
    }
    return n
}

func PBConvPB_Followed_To_Followed ( o *Followed) *PB_Followed {
     n := &PB_Followed{
      Id: int64 ( o.Id ),
      UserId: int32 ( o.UserId ),
      FollowedUserId: int32 ( o.FollowedUserId ),
      CreatedTime: int32 ( o.CreatedTime ),
    }
    return n
}
*/
/*
func PBConvPB__Likes_To_Likes( o *PB_Likes) *Likes {
     n := &Likes{
      Id: int ( o.Id ),
      PostId: int ( o.PostId ),
      PostTypeEnum: int ( o.PostTypeEnum ),
      UserId: int ( o.UserId ),
      LikeEnum: int ( o.LikeEnum ),
      CreatedTime: int ( o.CreatedTime ),
    }
    return n
}

func PBConvPB_Likes_To_Likes ( o *Likes) *PB_Likes {
     n := &PB_Likes{
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
      NotifyType: int ( o.NotifyType ),
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
      NotifyType: int32 ( o.NotifyType ),
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
func PBConvPB__PhoneContacts_To_PhoneContacts( o *PB_PhoneContacts) *PhoneContacts {
     n := &PhoneContacts{
      Id: int ( o.Id ),
      UserId: int ( o.UserId ),
      ClientId: int ( o.ClientId ),
      Phone: string ( o.Phone ),
      FirstName: string ( o.FirstName ),
      LastName: string ( o.LastName ),
    }
    return n
}

func PBConvPB_PhoneContacts_To_PhoneContacts ( o *PhoneContacts) *PB_PhoneContacts {
     n := &PB_PhoneContacts{
      Id: int64 ( o.Id ),
      UserId: int32 ( o.UserId ),
      ClientId: int64 ( o.ClientId ),
      Phone: string ( o.Phone ),
      FirstName: string ( o.FirstName ),
      LastName: string ( o.LastName ),
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
      CommentsCount: int ( o.CommentsCount ),
      LikesCount: int ( o.LikesCount ),
      ViewsCount: int ( o.ViewsCount ),
      ReSharedCount: int ( o.ReSharedCount ),
      ChatSharedCount: int ( o.ChatSharedCount ),
    }
    return n
}

func PBConvPB_PostCount_To_PostCount ( o *PostCount) *PB_PostCount {
     n := &PB_PostCount{
      PostId: int64 ( o.PostId ),
      CommentsCount: int32 ( o.CommentsCount ),
      LikesCount: int32 ( o.LikesCount ),
      ViewsCount: int64 ( o.ViewsCount ),
      ReSharedCount: int32 ( o.ReSharedCount ),
      ChatSharedCount: int32 ( o.ChatSharedCount ),
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
func PBConvPB__PostKeys_To_PostKeys( o *PB_PostKeys) *PostKeys {
     n := &PostKeys{
      Id: int ( o.Id ),
      PostKeyStr: string ( o.PostKeyStr ),
      Used: int ( o.Used ),
    }
    return n
}

func PBConvPB_PostKeys_To_PostKeys ( o *PostKeys) *PB_PostKeys {
     n := &PB_PostKeys{
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
      PostType: int ( o.PostType ),
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
      PostType: int32 ( o.PostType ),
      CreatedTime: int32 ( o.CreatedTime ),
    }
    return n
}
*/
/*
func PBConvPB__PostReshared_To_PostReshared( o *PB_PostReshared) *PostReshared {
     n := &PostReshared{
      ResharedId: int ( o.ResharedId ),
      PostId: int ( o.PostId ),
      ByUserId: int ( o.ByUserId ),
      PostUserId: int ( o.PostUserId ),
      CreatedTime: int ( o.CreatedTime ),
    }
    return n
}

func PBConvPB_PostReshared_To_PostReshared ( o *PostReshared) *PB_PostReshared {
     n := &PB_PostReshared{
      ResharedId: int64 ( o.ResharedId ),
      PostId: int64 ( o.PostId ),
      ByUserId: int32 ( o.ByUserId ),
      PostUserId: int32 ( o.PostUserId ),
      CreatedTime: int32 ( o.CreatedTime ),
    }
    return n
}
*/
/*
func PBConvPB__Session_To_Session( o *PB_Session) *Session {
     n := &Session{
      Id: int ( o.Id ),
      SessionUuid: string ( o.SessionUuid ),
      UserId: int ( o.UserId ),
      LastIpAddress: string ( o.LastIpAddress ),
      AppVersion: int ( o.AppVersion ),
      ActiveTime: int ( o.ActiveTime ),
      CreatedTime: int ( o.CreatedTime ),
    }
    return n
}

func PBConvPB_Session_To_Session ( o *Session) *PB_Session {
     n := &PB_Session{
      Id: int64 ( o.Id ),
      SessionUuid: string ( o.SessionUuid ),
      UserId: int32 ( o.UserId ),
      LastIpAddress: string ( o.LastIpAddress ),
      AppVersion: int32 ( o.AppVersion ),
      ActiveTime: int32 ( o.ActiveTime ),
      CreatedTime: int32 ( o.CreatedTime ),
    }
    return n
}
*/
/*
func PBConvPB__SettingNotifications_To_SettingNotifications( o *PB_SettingNotifications) *SettingNotifications {
     n := &SettingNotifications{
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

func PBConvPB_SettingNotifications_To_SettingNotifications ( o *SettingNotifications) *PB_SettingNotifications {
     n := &PB_SettingNotifications{
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
func PBConvPB__Sms_To_Sms( o *PB_Sms) *Sms {
     n := &Sms{
      Id: int ( o.Id ),
      Hash: string ( o.Hash ),
      AppUuid: string ( o.AppUuid ),
      ClientPhone: string ( o.ClientPhone ),
      GenratedCode: int ( o.GenratedCode ),
      SmsSenderNumber: string ( o.SmsSenderNumber ),
      SmsSendStatues: string ( o.SmsSendStatues ),
      SmsHttpBody: string ( o.SmsHttpBody ),
      Err: string ( o.Err ),
      Carrier: string ( o.Carrier ),
      Country: []byte ( o.Country ),
      IsValidPhone: int ( o.IsValidPhone ),
      IsConfirmed: int ( o.IsConfirmed ),
      IsLogin: int ( o.IsLogin ),
      IsRegister: int ( o.IsRegister ),
      RetriedCount: int ( o.RetriedCount ),
      TTL: int ( o.TTL ),
    }
    return n
}

func PBConvPB_Sms_To_Sms ( o *Sms) *PB_Sms {
     n := &PB_Sms{
      Id: int32 ( o.Id ),
      Hash: string ( o.Hash ),
      AppUuid: string ( o.AppUuid ),
      ClientPhone: string ( o.ClientPhone ),
      GenratedCode: int32 ( o.GenratedCode ),
      SmsSenderNumber: string ( o.SmsSenderNumber ),
      SmsSendStatues: string ( o.SmsSendStatues ),
      SmsHttpBody: string ( o.SmsHttpBody ),
      Err: string ( o.Err ),
      Carrier: string ( o.Carrier ),
      Country: []byte ( o.Country ),
      IsValidPhone: int32 ( o.IsValidPhone ),
      IsConfirmed: int32 ( o.IsConfirmed ),
      IsLogin: int32 ( o.IsLogin ),
      IsRegister: int32 ( o.IsRegister ),
      RetriedCount: int32 ( o.RetriedCount ),
      TTL: int32 ( o.TTL ),
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
      IsVerified: int ( o.IsVerified ),
      AvatarId: int ( o.AvatarId ),
      AccessHash: int ( o.AccessHash ),
      ProfilePrivacy: int ( o.ProfilePrivacy ),
      OnlinePrivacy: int ( o.OnlinePrivacy ),
      CallPrivacy: int ( o.CallPrivacy ),
      AddToGroupPrivacy: int ( o.AddToGroupPrivacy ),
      SeenMessagePrivacy: int ( o.SeenMessagePrivacy ),
      Phone: string ( o.Phone ),
      Email: string ( o.Email ),
      About: string ( o.About ),
      DefaultUserName: string ( o.DefaultUserName ),
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
      LastPostTime: int ( o.LastPostTime ),
      CreatedTime: int ( o.CreatedTime ),
      VersionTime: int ( o.VersionTime ),
      IsDeleted: int ( o.IsDeleted ),
      IsBanned: int ( o.IsBanned ),
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
      IsVerified: int32 ( o.IsVerified ),
      AvatarId: int64 ( o.AvatarId ),
      AccessHash: int32 ( o.AccessHash ),
      ProfilePrivacy: int32 ( o.ProfilePrivacy ),
      OnlinePrivacy: int32 ( o.OnlinePrivacy ),
      CallPrivacy: int32 ( o.CallPrivacy ),
      AddToGroupPrivacy: int32 ( o.AddToGroupPrivacy ),
      SeenMessagePrivacy: int32 ( o.SeenMessagePrivacy ),
      Phone: string ( o.Phone ),
      Email: string ( o.Email ),
      About: string ( o.About ),
      DefaultUserName: string ( o.DefaultUserName ),
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
      LastPostTime: int32 ( o.LastPostTime ),
      CreatedTime: int32 ( o.CreatedTime ),
      VersionTime: int32 ( o.VersionTime ),
      IsDeleted: int32 ( o.IsDeleted ),
      IsBanned: int32 ( o.IsBanned ),
    }
    return n
}
*/
/*
func PBConvPB__UserRelation_To_UserRelation( o *PB_UserRelation) *UserRelation {
     n := &UserRelation{
      RelNanoId: int ( o.RelNanoId ),
      UserId: int ( o.UserId ),
      PeerUserId: int ( o.PeerUserId ),
      Follwing: int ( o.Follwing ),
      Followed: int ( o.Followed ),
      InContacts: int ( o.InContacts ),
      MutualContact: int ( o.MutualContact ),
      IsFavorite: int ( o.IsFavorite ),
      Notify: int ( o.Notify ),
    }
    return n
}

func PBConvPB_UserRelation_To_UserRelation ( o *UserRelation) *PB_UserRelation {
     n := &PB_UserRelation{
      RelNanoId: int64 ( o.RelNanoId ),
      UserId: int32 ( o.UserId ),
      PeerUserId: int32 ( o.PeerUserId ),
      Follwing: int32 ( o.Follwing ),
      Followed: int32 ( o.Followed ),
      InContacts: int32 ( o.InContacts ),
      MutualContact: int32 ( o.MutualContact ),
      IsFavorite: int32 ( o.IsFavorite ),
      Notify: int32 ( o.Notify ),
    }
    return n
}
*/
/*
func PBConvPB__Chat_To_Chat( o *PB_Chat) *Chat {
     n := &Chat{
      ChatId: int ( o.ChatId ),
      ChatKey: string ( o.ChatKey ),
      RoomKey: string ( o.RoomKey ),
      RoomType: int ( o.RoomType ),
      UserId: int ( o.UserId ),
      PeerUserId: int ( o.PeerUserId ),
      GroupId: int ( o.GroupId ),
      HashTagId: int ( o.HashTagId ),
      Title: string ( o.Title ),
      PinTimeMs: int ( o.PinTimeMs ),
      FromMsgId: int ( o.FromMsgId ),
      UnseenCount: int ( o.UnseenCount ),
      Seq: int ( o.Seq ),
      LastMsgId: int ( o.LastMsgId ),
      LastMyMsgStatus: int ( o.LastMyMsgStatus ),
      MyLastSeenSeq: int ( o.MyLastSeenSeq ),
      MyLastSeenMsgId: int ( o.MyLastSeenMsgId ),
      PeerLastSeenMsgId: int ( o.PeerLastSeenMsgId ),
      MyLastDeliveredSeq: int ( o.MyLastDeliveredSeq ),
      MyLastDeliveredMsgId: int ( o.MyLastDeliveredMsgId ),
      PeerLastDeliveredMsgId: int ( o.PeerLastDeliveredMsgId ),
      IsActive: int ( o.IsActive ),
      IsLeft: int ( o.IsLeft ),
      IsCreator: int ( o.IsCreator ),
      IsKicked: int ( o.IsKicked ),
      IsAdmin: int ( o.IsAdmin ),
      IsDeactivated: int ( o.IsDeactivated ),
      IsMuted: int ( o.IsMuted ),
      MuteUntil: int ( o.MuteUntil ),
      VersionTimeMs: int ( o.VersionTimeMs ),
      VersionSeq: int ( o.VersionSeq ),
      OrderTime: int ( o.OrderTime ),
      CreatedTime: int ( o.CreatedTime ),
      DraftText: string ( o.DraftText ),
      DratReplyToMsgId: int ( o.DratReplyToMsgId ),
    }
    return n
}

func PBConvPB_Chat_To_Chat ( o *Chat) *PB_Chat {
     n := &PB_Chat{
      ChatId: int64 ( o.ChatId ),
      ChatKey: string ( o.ChatKey ),
      RoomKey: string ( o.RoomKey ),
      RoomType: int32 ( o.RoomType ),
      UserId: int32 ( o.UserId ),
      PeerUserId: int32 ( o.PeerUserId ),
      GroupId: int64 ( o.GroupId ),
      HashTagId: int64 ( o.HashTagId ),
      Title: string ( o.Title ),
      PinTimeMs: int64 ( o.PinTimeMs ),
      FromMsgId: int64 ( o.FromMsgId ),
      UnseenCount: int32 ( o.UnseenCount ),
      Seq: int32 ( o.Seq ),
      LastMsgId: int64 ( o.LastMsgId ),
      LastMyMsgStatus: int32 ( o.LastMyMsgStatus ),
      MyLastSeenSeq: int32 ( o.MyLastSeenSeq ),
      MyLastSeenMsgId: int64 ( o.MyLastSeenMsgId ),
      PeerLastSeenMsgId: int64 ( o.PeerLastSeenMsgId ),
      MyLastDeliveredSeq: int32 ( o.MyLastDeliveredSeq ),
      MyLastDeliveredMsgId: int64 ( o.MyLastDeliveredMsgId ),
      PeerLastDeliveredMsgId: int64 ( o.PeerLastDeliveredMsgId ),
      IsActive: int32 ( o.IsActive ),
      IsLeft: int32 ( o.IsLeft ),
      IsCreator: int32 ( o.IsCreator ),
      IsKicked: int32 ( o.IsKicked ),
      IsAdmin: int32 ( o.IsAdmin ),
      IsDeactivated: int32 ( o.IsDeactivated ),
      IsMuted: int32 ( o.IsMuted ),
      MuteUntil: int32 ( o.MuteUntil ),
      VersionTimeMs: int64 ( o.VersionTimeMs ),
      VersionSeq: int32 ( o.VersionSeq ),
      OrderTime: int32 ( o.OrderTime ),
      CreatedTime: int32 ( o.CreatedTime ),
      DraftText: string ( o.DraftText ),
      DratReplyToMsgId: int64 ( o.DratReplyToMsgId ),
    }
    return n
}
*/
/*
func PBConvPB__ChatDeleted_To_ChatDeleted( o *PB_ChatDeleted) *ChatDeleted {
     n := &ChatDeleted{
      ChatId: int ( o.ChatId ),
      RoomKey: string ( o.RoomKey ),
    }
    return n
}

func PBConvPB_ChatDeleted_To_ChatDeleted ( o *ChatDeleted) *PB_ChatDeleted {
     n := &PB_ChatDeleted{
      ChatId: int64 ( o.ChatId ),
      RoomKey: string ( o.RoomKey ),
    }
    return n
}
*/
/*
func PBConvPB__ChatLastMessage_To_ChatLastMessage( o *PB_ChatLastMessage) *ChatLastMessage {
     n := &ChatLastMessage{
      ChatIdGroupId: string ( o.ChatIdGroupId ),
      LastMsgPb: []byte ( o.LastMsgPb ),
    }
    return n
}

func PBConvPB_ChatLastMessage_To_ChatLastMessage ( o *ChatLastMessage) *PB_ChatLastMessage {
     n := &PB_ChatLastMessage{
      ChatIdGroupId: string ( o.ChatIdGroupId ),
      LastMsgPb: []byte ( o.LastMsgPb ),
    }
    return n
}
*/
/*
func PBConvPB__ChatUserVersion_To_ChatUserVersion( o *PB_ChatUserVersion) *ChatUserVersion {
     n := &ChatUserVersion{
      UserId: int ( o.UserId ),
      ChatId: int ( o.ChatId ),
      VersionTimeNano: int ( o.VersionTimeNano ),
    }
    return n
}

func PBConvPB_ChatUserVersion_To_ChatUserVersion ( o *ChatUserVersion) *PB_ChatUserVersion {
     n := &PB_ChatUserVersion{
      UserId: int32 ( o.UserId ),
      ChatId: int64 ( o.ChatId ),
      VersionTimeNano: int32 ( o.VersionTimeNano ),
    }
    return n
}
*/
/*
func PBConvPB__Group_To_Group( o *PB_Group) *Group {
     n := &Group{
      GroupId: int ( o.GroupId ),
      GroupKey: string ( o.GroupKey ),
      GroupName: string ( o.GroupName ),
      UserName: string ( o.UserName ),
      IsSuperGroup: int ( o.IsSuperGroup ),
      HashTagId: int ( o.HashTagId ),
      CreatorUserId: int ( o.CreatorUserId ),
      GroupPrivacy: int ( o.GroupPrivacy ),
      HistoryViewAble: int ( o.HistoryViewAble ),
      Seq: int ( o.Seq ),
      LastMsgId: int ( o.LastMsgId ),
      PinedMsgId: int ( o.PinedMsgId ),
      AvatarRefId: int ( o.AvatarRefId ),
      AvatarCount: int ( o.AvatarCount ),
      About: string ( o.About ),
      InviteLinkHash: string ( o.InviteLinkHash ),
      MembersCount: int ( o.MembersCount ),
      AdminsCount: int ( o.AdminsCount ),
      ModeratorCounts: int ( o.ModeratorCounts ),
      SortTime: int ( o.SortTime ),
      CreatedTime: int ( o.CreatedTime ),
      IsMute: int ( o.IsMute ),
      IsDeleted: int ( o.IsDeleted ),
      IsBanned: int ( o.IsBanned ),
    }
    return n
}

func PBConvPB_Group_To_Group ( o *Group) *PB_Group {
     n := &PB_Group{
      GroupId: int64 ( o.GroupId ),
      GroupKey: string ( o.GroupKey ),
      GroupName: string ( o.GroupName ),
      UserName: string ( o.UserName ),
      IsSuperGroup: int32 ( o.IsSuperGroup ),
      HashTagId: int64 ( o.HashTagId ),
      CreatorUserId: int32 ( o.CreatorUserId ),
      GroupPrivacy: int32 ( o.GroupPrivacy ),
      HistoryViewAble: int32 ( o.HistoryViewAble ),
      Seq: int64 ( o.Seq ),
      LastMsgId: int64 ( o.LastMsgId ),
      PinedMsgId: int64 ( o.PinedMsgId ),
      AvatarRefId: int64 ( o.AvatarRefId ),
      AvatarCount: int32 ( o.AvatarCount ),
      About: string ( o.About ),
      InviteLinkHash: string ( o.InviteLinkHash ),
      MembersCount: int32 ( o.MembersCount ),
      AdminsCount: int32 ( o.AdminsCount ),
      ModeratorCounts: int32 ( o.ModeratorCounts ),
      SortTime: int32 ( o.SortTime ),
      CreatedTime: int32 ( o.CreatedTime ),
      IsMute: int32 ( o.IsMute ),
      IsDeleted: int32 ( o.IsDeleted ),
      IsBanned: int32 ( o.IsBanned ),
    }
    return n
}
*/
/*
func PBConvPB__GroupMember_To_GroupMember( o *PB_GroupMember) *GroupMember {
     n := &GroupMember{
      OrderId: int ( o.OrderId ),
      GroupId: int ( o.GroupId ),
      UserId: int ( o.UserId ),
      ByUserId: int ( o.ByUserId ),
      GroupRole: int ( o.GroupRole ),
      CreatedTime: int ( o.CreatedTime ),
    }
    return n
}

func PBConvPB_GroupMember_To_GroupMember ( o *GroupMember) *PB_GroupMember {
     n := &PB_GroupMember{
      OrderId: int64 ( o.OrderId ),
      GroupId: int64 ( o.GroupId ),
      UserId: int32 ( o.UserId ),
      ByUserId: int32 ( o.ByUserId ),
      GroupRole: int32 ( o.GroupRole ),
      CreatedTime: int32 ( o.CreatedTime ),
    }
    return n
}
*/
/*
func PBConvPB__GroupOrderdUser_To_GroupOrderdUser( o *PB_GroupOrderdUser) *GroupOrderdUser {
     n := &GroupOrderdUser{
      OrderId: int ( o.OrderId ),
      GroupId: int ( o.GroupId ),
      UserId: int ( o.UserId ),
    }
    return n
}

func PBConvPB_GroupOrderdUser_To_GroupOrderdUser ( o *GroupOrderdUser) *PB_GroupOrderdUser {
     n := &PB_GroupOrderdUser{
      OrderId: int64 ( o.OrderId ),
      GroupId: int64 ( o.GroupId ),
      UserId: int32 ( o.UserId ),
    }
    return n
}
*/
/*
func PBConvPB__GroupPinedMsg_To_GroupPinedMsg( o *PB_GroupPinedMsg) *GroupPinedMsg {
     n := &GroupPinedMsg{
      MsgId: int ( o.MsgId ),
      MsgPb: []byte ( o.MsgPb ),
    }
    return n
}

func PBConvPB_GroupPinedMsg_To_GroupPinedMsg ( o *GroupPinedMsg) *PB_GroupPinedMsg {
     n := &PB_GroupPinedMsg{
      MsgId: int64 ( o.MsgId ),
      MsgPb: []byte ( o.MsgPb ),
    }
    return n
}
*/
/*
func PBConvPB__FileMsg_To_FileMsg( o *PB_FileMsg) *FileMsg {
     n := &FileMsg{
      Id: int ( o.Id ),
      AccessHash: int ( o.AccessHash ),
      FileType: int ( o.FileType ),
      Width: int ( o.Width ),
      Height: int ( o.Height ),
      Extension: string ( o.Extension ),
      UserId: int ( o.UserId ),
      DataThumb: []byte ( o.DataThumb ),
      Data: []byte ( o.Data ),
    }
    return n
}

func PBConvPB_FileMsg_To_FileMsg ( o *FileMsg) *PB_FileMsg {
     n := &PB_FileMsg{
      Id: int64 ( o.Id ),
      AccessHash: int32 ( o.AccessHash ),
      FileType: int32 ( o.FileType ),
      Width: int32 ( o.Width ),
      Height: int32 ( o.Height ),
      Extension: string ( o.Extension ),
      UserId: int32 ( o.UserId ),
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
      AccessHash: int ( o.AccessHash ),
      FileType: int ( o.FileType ),
      Width: int ( o.Width ),
      Height: int ( o.Height ),
      Extension: string ( o.Extension ),
      UserId: int ( o.UserId ),
      DataThumb: []byte ( o.DataThumb ),
      Data: []byte ( o.Data ),
    }
    return n
}

func PBConvPB_FilePost_To_FilePost ( o *FilePost) *PB_FilePost {
     n := &PB_FilePost{
      Id: int64 ( o.Id ),
      AccessHash: int32 ( o.AccessHash ),
      FileType: int32 ( o.FileType ),
      Width: int32 ( o.Width ),
      Height: int32 ( o.Height ),
      Extension: string ( o.Extension ),
      UserId: int32 ( o.UserId ),
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
      ResharedId: int ( o.ResharedId ),
    }
    return n
}

func PBConvPB_HomeFanout_To_HomeFanout ( o *HomeFanout) *PB_HomeFanout {
     n := &PB_HomeFanout{
      OrderId: int64 ( o.OrderId ),
      ForUserId: int64 ( o.ForUserId ),
      PostId: int64 ( o.PostId ),
      PostUserId: int64 ( o.PostUserId ),
      ResharedId: int64 ( o.ResharedId ),
    }
    return n
}
*/
/*
func PBConvPB__SuggestedTopPosts_To_SuggestedTopPosts( o *PB_SuggestedTopPosts) *SuggestedTopPosts {
     n := &SuggestedTopPosts{
      Id: int ( o.Id ),
      PostId: int ( o.PostId ),
    }
    return n
}

func PBConvPB_SuggestedTopPosts_To_SuggestedTopPosts ( o *SuggestedTopPosts) *PB_SuggestedTopPosts {
     n := &PB_SuggestedTopPosts{
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
/*
func PBConvPB__PushChat_To_PushChat( o *PB_PushChat) *PushChat {
     n := &PushChat{
      PushId: int ( o.PushId ),
      ToUserId: int ( o.ToUserId ),
      PushTypeId: int ( o.PushTypeId ),
      RoomKey: string ( o.RoomKey ),
      ChatKey: string ( o.ChatKey ),
      Seq: int ( o.Seq ),
      UnseenCount: int ( o.UnseenCount ),
      FromHighMessageId: int ( o.FromHighMessageId ),
      ToLowMessageId: int ( o.ToLowMessageId ),
      MessageId: int ( o.MessageId ),
      MessageFileId: int ( o.MessageFileId ),
      MessagePb: []byte ( o.MessagePb ),
      MessageJson: string ( o.MessageJson ),
      CreatedTime: int ( o.CreatedTime ),
    }
    return n
}

func PBConvPB_PushChat_To_PushChat ( o *PushChat) *PB_PushChat {
     n := &PB_PushChat{
      PushId: int64 ( o.PushId ),
      ToUserId: int32 ( o.ToUserId ),
      PushTypeId: int32 ( o.PushTypeId ),
      RoomKey: string ( o.RoomKey ),
      ChatKey: string ( o.ChatKey ),
      Seq: int32 ( o.Seq ),
      UnseenCount: int32 ( o.UnseenCount ),
      FromHighMessageId: int64 ( o.FromHighMessageId ),
      ToLowMessageId: int64 ( o.ToLowMessageId ),
      MessageId: int64 ( o.MessageId ),
      MessageFileId: int64 ( o.MessageFileId ),
      MessagePb: []byte ( o.MessagePb ),
      MessageJson: string ( o.MessageJson ),
      CreatedTime: int32 ( o.CreatedTime ),
    }
    return n
}
*/
/*
func PBConvPB__HTTPRPCLog_To_HTTPRPCLog( o *PB_HTTPRPCLog) *HTTPRPCLog {
     n := &HTTPRPCLog{
      Id: int ( o.Id ),
      Time: string ( o.Time ),
      MethodFull: string ( o.MethodFull ),
      MethodParent: string ( o.MethodParent ),
      UserId: int ( o.UserId ),
      SessionId: string ( o.SessionId ),
      StatusCode: int ( o.StatusCode ),
      InputSize: int ( o.InputSize ),
      OutputSize: int ( o.OutputSize ),
      ReqestJson: string ( o.ReqestJson ),
      ResponseJson: string ( o.ResponseJson ),
      ReqestParamJson: string ( o.ReqestParamJson ),
      ResponseMsgJson: string ( o.ResponseMsgJson ),
    }
    return n
}

func PBConvPB_HTTPRPCLog_To_HTTPRPCLog ( o *HTTPRPCLog) *PB_HTTPRPCLog {
     n := &PB_HTTPRPCLog{
      Id: int32 ( o.Id ),
      Time: string ( o.Time ),
      MethodFull: string ( o.MethodFull ),
      MethodParent: string ( o.MethodParent ),
      UserId: int32 ( o.UserId ),
      SessionId: string ( o.SessionId ),
      StatusCode: int32 ( o.StatusCode ),
      InputSize: int32 ( o.InputSize ),
      OutputSize: int32 ( o.OutputSize ),
      ReqestJson: string ( o.ReqestJson ),
      ResponseJson: string ( o.ResponseJson ),
      ReqestParamJson: string ( o.ReqestParamJson ),
      ResponseMsgJson: string ( o.ResponseMsgJson ),
    }
    return n
}
*/
/*
func PBConvPB__MetricLog_To_MetricLog( o *PB_MetricLog) *MetricLog {
     n := &MetricLog{
      Id: int ( o.Id ),
      InstanceId: int ( o.InstanceId ),
      StartFrom: string ( o.StartFrom ),
      EndTo: string ( o.EndTo ),
      StartTime: int ( o.StartTime ),
      Duration: string ( o.Duration ),
      MetericsJson: string ( o.MetericsJson ),
    }
    return n
}

func PBConvPB_MetricLog_To_MetricLog ( o *MetricLog) *PB_MetricLog {
     n := &PB_MetricLog{
      Id: int32 ( o.Id ),
      InstanceId: int32 ( o.InstanceId ),
      StartFrom: string ( o.StartFrom ),
      EndTo: string ( o.EndTo ),
      StartTime: int32 ( o.StartTime ),
      Duration: string ( o.Duration ),
      MetericsJson: string ( o.MetericsJson ),
    }
    return n
}
*/
/*
func PBConvPB__XfileServiceInfoLog_To_XfileServiceInfoLog( o *PB_XfileServiceInfoLog) *XfileServiceInfoLog {
     n := &XfileServiceInfoLog{
      Id: int ( o.Id ),
      InstanceId: int ( o.InstanceId ),
      Url: string ( o.Url ),
      CreatedTime: string ( o.CreatedTime ),
    }
    return n
}

func PBConvPB_XfileServiceInfoLog_To_XfileServiceInfoLog ( o *XfileServiceInfoLog) *PB_XfileServiceInfoLog {
     n := &PB_XfileServiceInfoLog{
      Id: int64 ( o.Id ),
      InstanceId: int32 ( o.InstanceId ),
      Url: string ( o.Url ),
      CreatedTime: string ( o.CreatedTime ),
    }
    return n
}
*/
/*
func PBConvPB__XfileServiceMetricLog_To_XfileServiceMetricLog( o *PB_XfileServiceMetricLog) *XfileServiceMetricLog {
     n := &XfileServiceMetricLog{
      Id: int ( o.Id ),
      InstanceId: int ( o.InstanceId ),
      MetricJson: string ( o.MetricJson ),
    }
    return n
}

func PBConvPB_XfileServiceMetricLog_To_XfileServiceMetricLog ( o *XfileServiceMetricLog) *PB_XfileServiceMetricLog {
     n := &PB_XfileServiceMetricLog{
      Id: int64 ( o.Id ),
      InstanceId: int32 ( o.InstanceId ),
      MetricJson: string ( o.MetricJson ),
    }
    return n
}
*/
/*
func PBConvPB__XfileServiceRequestLog_To_XfileServiceRequestLog( o *PB_XfileServiceRequestLog) *XfileServiceRequestLog {
     n := &XfileServiceRequestLog{
      Id: int ( o.Id ),
      LocalSeq: int ( o.LocalSeq ),
      InstanceId: int ( o.InstanceId ),
      Url: string ( o.Url ),
      HttpCode: int ( o.HttpCode ),
      CreatedTime: string ( o.CreatedTime ),
    }
    return n
}

func PBConvPB_XfileServiceRequestLog_To_XfileServiceRequestLog ( o *XfileServiceRequestLog) *PB_XfileServiceRequestLog {
     n := &PB_XfileServiceRequestLog{
      Id: int64 ( o.Id ),
      LocalSeq: int32 ( o.LocalSeq ),
      InstanceId: int32 ( o.InstanceId ),
      Url: string ( o.Url ),
      HttpCode: int32 ( o.HttpCode ),
      CreatedTime: string ( o.CreatedTime ),
    }
    return n
}
*/
/*
func PBConvPB__InvalidateCache_To_InvalidateCache( o *PB_InvalidateCache) *InvalidateCache {
     n := &InvalidateCache{
      OrderId: int ( o.OrderId ),
      CacheKey: string ( o.CacheKey ),
    }
    return n
}

func PBConvPB_InvalidateCache_To_InvalidateCache ( o *InvalidateCache) *PB_InvalidateCache {
     n := &PB_InvalidateCache{
      OrderId: int64 ( o.OrderId ),
      CacheKey: string ( o.CacheKey ),
    }
    return n
}
*/
