
/*Table: action  */
ALTER TABLE sun.action CHANGE COLUMN ActionId action_id bigint(20);
ALTER TABLE sun.action CHANGE COLUMN ActorUserId actor_user_id int(11);
ALTER TABLE sun.action CHANGE COLUMN ActionType action_type int(11);
ALTER TABLE sun.action CHANGE COLUMN PeerUserId peer_user_id int(20);
ALTER TABLE sun.action CHANGE COLUMN PostId post_id bigint(20);
ALTER TABLE sun.action CHANGE COLUMN CommentId comment_id bigint(20);
ALTER TABLE sun.action CHANGE COLUMN Murmur64Hash murmur64_hash bigint(20);
ALTER TABLE sun.action CHANGE COLUMN CreatedTime created_time int(11);

/*Table: blocked  */
ALTER TABLE sun.blocked CHANGE COLUMN Id id bigint(20);
ALTER TABLE sun.blocked CHANGE COLUMN UserId user_id int(11);
ALTER TABLE sun.blocked CHANGE COLUMN BlockedUserId blocked_user_id int(11);
ALTER TABLE sun.blocked CHANGE COLUMN CreatedTime created_time int(11);

/*Table: comment  */
ALTER TABLE sun.comment CHANGE COLUMN CommentId comment_id bigint(20);
ALTER TABLE sun.comment CHANGE COLUMN UserId user_id int(11);
ALTER TABLE sun.comment CHANGE COLUMN PostId post_id bigint(20);
ALTER TABLE sun.comment CHANGE COLUMN Text text text;
ALTER TABLE sun.comment CHANGE COLUMN LikesCount likes_count int(11);
ALTER TABLE sun.comment CHANGE COLUMN IsEdited is_edited tinyint(4);
ALTER TABLE sun.comment CHANGE COLUMN CreatedTime created_time int(11);

/*Table: comment_deleted  */
ALTER TABLE sun.comment_deleted CHANGE COLUMN CommentId comment_id bigint(20);
ALTER TABLE sun.comment_deleted CHANGE COLUMN UserId user_id int(11);

/*Table: event  */
ALTER TABLE sun.event CHANGE COLUMN EventId event_id bigint(20);
ALTER TABLE sun.event CHANGE COLUMN EventType event_type int(11);
ALTER TABLE sun.event CHANGE COLUMN ByUserId by_user_id bigint(20);
ALTER TABLE sun.event CHANGE COLUMN PeerUserId peer_user_id bigint(20);
ALTER TABLE sun.event CHANGE COLUMN PostId post_id bigint(20);
ALTER TABLE sun.event CHANGE COLUMN CommentId comment_id bigint(20);
ALTER TABLE sun.event CHANGE COLUMN ActionId action_id bigint(20);
ALTER TABLE sun.event CHANGE COLUMN Murmur64Hash murmur64_hash bigint(20);
ALTER TABLE sun.event CHANGE COLUMN ChatKey chat_key varchar(100);
ALTER TABLE sun.event CHANGE COLUMN MessageId message_id bigint(20);
ALTER TABLE sun.event CHANGE COLUMN ReSharedId re_shared_id bigint(20);

/*Table: followed  */
ALTER TABLE sun.followed CHANGE COLUMN Id id bigint(20);
ALTER TABLE sun.followed CHANGE COLUMN UserId user_id int(11);
ALTER TABLE sun.followed CHANGE COLUMN FollowedUserId followed_user_id int(11);
ALTER TABLE sun.followed CHANGE COLUMN CreatedTime created_time int(11);

/*Table: likes  */
ALTER TABLE sun.likes CHANGE COLUMN Id id bigint(11);
ALTER TABLE sun.likes CHANGE COLUMN PostId post_id bigint(11);
ALTER TABLE sun.likes CHANGE COLUMN PostTypeEnum post_type_enum tinyint(4);
ALTER TABLE sun.likes CHANGE COLUMN UserId user_id int(11);
ALTER TABLE sun.likes CHANGE COLUMN LikeEnum like_enum tinyint(4);
ALTER TABLE sun.likes CHANGE COLUMN CreatedTime created_time int(11);

/*Table: notify  */
ALTER TABLE sun.notify CHANGE COLUMN NotifyId notify_id bigint(10);
ALTER TABLE sun.notify CHANGE COLUMN ForUserId for_user_id int(10);
ALTER TABLE sun.notify CHANGE COLUMN ActorUserId actor_user_id int(10);
ALTER TABLE sun.notify CHANGE COLUMN NotifyType notify_type int(10);
ALTER TABLE sun.notify CHANGE COLUMN PostId post_id bigint(20);
ALTER TABLE sun.notify CHANGE COLUMN CommentId comment_id bigint(20);
ALTER TABLE sun.notify CHANGE COLUMN PeerUserId peer_user_id int(11);
ALTER TABLE sun.notify CHANGE COLUMN Murmur64Hash murmur64_hash bigint(20);
ALTER TABLE sun.notify CHANGE COLUMN SeenStatus seen_status tinyint(10);
ALTER TABLE sun.notify CHANGE COLUMN CreatedTime created_time int(10);

/*Table: notify_removed  */
ALTER TABLE sun.notify_removed CHANGE COLUMN Murmur64Hash murmur64_hash bigint(11);
ALTER TABLE sun.notify_removed CHANGE COLUMN ForUserId for_user_id int(11);
ALTER TABLE sun.notify_removed CHANGE COLUMN Id id bigint(20);

/*Table: phone_contacts  */
ALTER TABLE sun.phone_contacts CHANGE COLUMN Id id bigint(10);
ALTER TABLE sun.phone_contacts CHANGE COLUMN UserId user_id int(10);
ALTER TABLE sun.phone_contacts CHANGE COLUMN ClientId client_id bigint(20);
ALTER TABLE sun.phone_contacts CHANGE COLUMN Phone phone varchar(20);
ALTER TABLE sun.phone_contacts CHANGE COLUMN FirstName first_name varchar(50);
ALTER TABLE sun.phone_contacts CHANGE COLUMN LastName last_name varchar(50);

/*Table: post  */
ALTER TABLE sun.post CHANGE COLUMN PostId post_id bigint(20);
ALTER TABLE sun.post CHANGE COLUMN UserId user_id int(11);
ALTER TABLE sun.post CHANGE COLUMN PostTypeEnum post_type_enum int(4);
ALTER TABLE sun.post CHANGE COLUMN PostCategoryEnum post_category_enum int(11);
ALTER TABLE sun.post CHANGE COLUMN MediaId media_id bigint(11);
ALTER TABLE sun.post CHANGE COLUMN PostKey post_key varchar(50);
ALTER TABLE sun.post CHANGE COLUMN Text text text;
ALTER TABLE sun.post CHANGE COLUMN RichText rich_text text;
ALTER TABLE sun.post CHANGE COLUMN MediaCount media_count tinyint(4);
ALTER TABLE sun.post CHANGE COLUMN SharedTo shared_to int(11);
ALTER TABLE sun.post CHANGE COLUMN DisableComment disable_comment tinyint(1);
ALTER TABLE sun.post CHANGE COLUMN Source source int(11);
ALTER TABLE sun.post CHANGE COLUMN HasTag has_tag int(11);
ALTER TABLE sun.post CHANGE COLUMN Seq seq int(11);
ALTER TABLE sun.post CHANGE COLUMN CommentsCount comments_count int(11);
ALTER TABLE sun.post CHANGE COLUMN LikesCount likes_count int(11);
ALTER TABLE sun.post CHANGE COLUMN ViewsCount views_count int(11);
ALTER TABLE sun.post CHANGE COLUMN EditedTime edited_time int(11);
ALTER TABLE sun.post CHANGE COLUMN CreatedTime created_time int(11);
ALTER TABLE sun.post CHANGE COLUMN ReSharedPostId re_shared_post_id bigint(20);

/*Table: post_count  */
ALTER TABLE sun.post_count CHANGE COLUMN PostId post_id bigint(20);
ALTER TABLE sun.post_count CHANGE COLUMN CommentsCount comments_count int(11);
ALTER TABLE sun.post_count CHANGE COLUMN LikesCount likes_count int(11);
ALTER TABLE sun.post_count CHANGE COLUMN ViewsCount views_count bigint(20);
ALTER TABLE sun.post_count CHANGE COLUMN ReSharedCount re_shared_count int(11);
ALTER TABLE sun.post_count CHANGE COLUMN ChatSharedCount chat_shared_count int(11);

/*Table: post_deleted  */
ALTER TABLE sun.post_deleted CHANGE COLUMN PostId post_id bigint(20);
ALTER TABLE sun.post_deleted CHANGE COLUMN UserId user_id int(11);

/*Table: post_keys  */
ALTER TABLE sun.post_keys CHANGE COLUMN Id id int(11);
ALTER TABLE sun.post_keys CHANGE COLUMN PostKeyStr post_key_str varchar(50);
ALTER TABLE sun.post_keys CHANGE COLUMN Used used tinyint(4);

/*Table: post_link  */
ALTER TABLE sun.post_link CHANGE COLUMN LinkId link_id bigint(20);
ALTER TABLE sun.post_link CHANGE COLUMN LinkUrl link_url varchar(500);

/*Table: post_media  */
ALTER TABLE sun.post_media CHANGE COLUMN MediaId media_id bigint(20);
ALTER TABLE sun.post_media CHANGE COLUMN UserId user_id int(11);
ALTER TABLE sun.post_media CHANGE COLUMN PostId post_id bigint(11);
ALTER TABLE sun.post_media CHANGE COLUMN AlbumId album_id int(11);
ALTER TABLE sun.post_media CHANGE COLUMN MediaTypeEnum media_type_enum smallint(4);
ALTER TABLE sun.post_media CHANGE COLUMN Width width smallint(6);
ALTER TABLE sun.post_media CHANGE COLUMN Height height smallint(6);
ALTER TABLE sun.post_media CHANGE COLUMN Size size int(11);
ALTER TABLE sun.post_media CHANGE COLUMN Duration duration int(11);
ALTER TABLE sun.post_media CHANGE COLUMN Extension extension varchar(50);
ALTER TABLE sun.post_media CHANGE COLUMN Md5Hash md5_hash varchar(72);
ALTER TABLE sun.post_media CHANGE COLUMN Color color varchar(10);
ALTER TABLE sun.post_media CHANGE COLUMN CreatedTime created_time int(11);
ALTER TABLE sun.post_media CHANGE COLUMN ViewCount view_count int(11);
ALTER TABLE sun.post_media CHANGE COLUMN Extra extra varchar(2000);

/*Table: post_mentioned  */
ALTER TABLE sun.post_mentioned CHANGE COLUMN MentionedId mentioned_id bigint(20);
ALTER TABLE sun.post_mentioned CHANGE COLUMN ForUserId for_user_id int(11);
ALTER TABLE sun.post_mentioned CHANGE COLUMN PostId post_id bigint(20);
ALTER TABLE sun.post_mentioned CHANGE COLUMN PostUserId post_user_id int(11);
ALTER TABLE sun.post_mentioned CHANGE COLUMN PostType post_type tinyint(11);
ALTER TABLE sun.post_mentioned CHANGE COLUMN CreatedTime created_time int(11);

/*Table: post_reshared  */
ALTER TABLE sun.post_reshared CHANGE COLUMN ResharedId reshared_id bigint(20);
ALTER TABLE sun.post_reshared CHANGE COLUMN PostId post_id bigint(20);
ALTER TABLE sun.post_reshared CHANGE COLUMN ByUserId by_user_id int(11);
ALTER TABLE sun.post_reshared CHANGE COLUMN PostUserId post_user_id int(11);
ALTER TABLE sun.post_reshared CHANGE COLUMN CreatedTime created_time int(11);

/*Table: session  */
ALTER TABLE sun.session CHANGE COLUMN Id id bigint(20);
ALTER TABLE sun.session CHANGE COLUMN SessionUuid session_uuid varchar(75);
ALTER TABLE sun.session CHANGE COLUMN UserId user_id int(11);
ALTER TABLE sun.session CHANGE COLUMN LastIpAddress last_ip_address varchar(75);
ALTER TABLE sun.session CHANGE COLUMN AppVersion app_version smallint(6);
ALTER TABLE sun.session CHANGE COLUMN ActiveTime active_time int(11);
ALTER TABLE sun.session CHANGE COLUMN CreatedTime created_time int(11);

/*Table: setting_notifications  */
ALTER TABLE sun.setting_notifications CHANGE COLUMN UserId user_id int(11);
ALTER TABLE sun.setting_notifications CHANGE COLUMN SocialLedOn social_led_on int(11);
ALTER TABLE sun.setting_notifications CHANGE COLUMN SocialLedColor social_led_color varchar(30);
ALTER TABLE sun.setting_notifications CHANGE COLUMN ReqestToFollowYou reqest_to_follow_you int(11);
ALTER TABLE sun.setting_notifications CHANGE COLUMN FollowedYou followed_you int(11);
ALTER TABLE sun.setting_notifications CHANGE COLUMN AccptedYourFollowRequest accpted_your_follow_request int(11);
ALTER TABLE sun.setting_notifications CHANGE COLUMN YourPostLiked your_post_liked int(11);
ALTER TABLE sun.setting_notifications CHANGE COLUMN YourPostCommented your_post_commented int(11);
ALTER TABLE sun.setting_notifications CHANGE COLUMN MenthenedYouInPost menthened_you_in_post int(11);
ALTER TABLE sun.setting_notifications CHANGE COLUMN MenthenedYouInComment menthened_you_in_comment int(11);
ALTER TABLE sun.setting_notifications CHANGE COLUMN YourContactsJoined your_contacts_joined int(11);
ALTER TABLE sun.setting_notifications CHANGE COLUMN DirectMessage direct_message int(11);
ALTER TABLE sun.setting_notifications CHANGE COLUMN DirectAlert direct_alert int(11);
ALTER TABLE sun.setting_notifications CHANGE COLUMN DirectPerview direct_perview int(11);
ALTER TABLE sun.setting_notifications CHANGE COLUMN DirectLedOn direct_led_on int(11);
ALTER TABLE sun.setting_notifications CHANGE COLUMN DirectLedColor direct_led_color int(11);
ALTER TABLE sun.setting_notifications CHANGE COLUMN DirectVibrate direct_vibrate int(11);
ALTER TABLE sun.setting_notifications CHANGE COLUMN DirectPopup direct_popup int(11);
ALTER TABLE sun.setting_notifications CHANGE COLUMN DirectSound direct_sound int(11);
ALTER TABLE sun.setting_notifications CHANGE COLUMN DirectPriority direct_priority int(11);

/*Table: sms  */
ALTER TABLE sun.sms CHANGE COLUMN Id id int(10) unsigned;
ALTER TABLE sun.sms CHANGE COLUMN Hash hash varchar(50);
ALTER TABLE sun.sms CHANGE COLUMN ClientPhone client_phone varchar(50);
ALTER TABLE sun.sms CHANGE COLUMN GenratedCode genrated_code int(11);
ALTER TABLE sun.sms CHANGE COLUMN SmsSenderNumber sms_sender_number varchar(50);
ALTER TABLE sun.sms CHANGE COLUMN SmsSendStatues sms_send_statues varchar(50);
ALTER TABLE sun.sms CHANGE COLUMN Carrier carrier varchar(50);
ALTER TABLE sun.sms CHANGE COLUMN Country country varbinary(50);
ALTER TABLE sun.sms CHANGE COLUMN IsValidPhone is_valid_phone tinyint(4);
ALTER TABLE sun.sms CHANGE COLUMN IsConfirmed is_confirmed tinyint(4);
ALTER TABLE sun.sms CHANGE COLUMN IsLogin is_login tinyint(4);
ALTER TABLE sun.sms CHANGE COLUMN IsRegister is_register tinyint(4);
ALTER TABLE sun.sms CHANGE COLUMN RetriedCount retried_count tinyint(4);

/*Table: tag  */
ALTER TABLE sun.tag CHANGE COLUMN TagId tag_id bigint(11);
ALTER TABLE sun.tag CHANGE COLUMN Name name varchar(100);
ALTER TABLE sun.tag CHANGE COLUMN Count count int(11);
ALTER TABLE sun.tag CHANGE COLUMN TagStatusEnum tag_status_enum int(1);
ALTER TABLE sun.tag CHANGE COLUMN CreatedTime created_time int(11);

/*Table: tag_post  */
ALTER TABLE sun.tag_post CHANGE COLUMN Id id bigint(11);
ALTER TABLE sun.tag_post CHANGE COLUMN TagId tag_id int(11);
ALTER TABLE sun.tag_post CHANGE COLUMN PostId post_id int(11);
ALTER TABLE sun.tag_post CHANGE COLUMN PostTypeEnum post_type_enum int(11);
ALTER TABLE sun.tag_post CHANGE COLUMN PostCategoryEnum post_category_enum int(11);
ALTER TABLE sun.tag_post CHANGE COLUMN CreatedTime created_time int(11);

/*Table: trigger_log  */
ALTER TABLE sun.trigger_log CHANGE COLUMN Id id bigint(20);
ALTER TABLE sun.trigger_log CHANGE COLUMN ModelName model_name varchar(50);
ALTER TABLE sun.trigger_log CHANGE COLUMN ChangeType change_type varchar(50);
ALTER TABLE sun.trigger_log CHANGE COLUMN TargetId target_id bigint(20);
ALTER TABLE sun.trigger_log CHANGE COLUMN TargetStr target_str varchar(100);
ALTER TABLE sun.trigger_log CHANGE COLUMN CreatedSe created_se int(11);

/*Table: user  */
ALTER TABLE sun.user CHANGE COLUMN UserId user_id int(10);
ALTER TABLE sun.user CHANGE COLUMN UserName user_name varchar(250);
ALTER TABLE sun.user CHANGE COLUMN UserNameLower user_name_lower varchar(250);
ALTER TABLE sun.user CHANGE COLUMN FirstName first_name varchar(250);
ALTER TABLE sun.user CHANGE COLUMN LastName last_name varchar(250);
ALTER TABLE sun.user CHANGE COLUMN IsVerified is_verified tinyint(4);
ALTER TABLE sun.user CHANGE COLUMN AvatarId avatar_id bigint(20);
ALTER TABLE sun.user CHANGE COLUMN ProfilePrivacy profile_privacy int(10);
ALTER TABLE sun.user CHANGE COLUMN OnlinePrivacy online_privacy int(11);
ALTER TABLE sun.user CHANGE COLUMN CallPrivacy call_privacy int(4);
ALTER TABLE sun.user CHANGE COLUMN AddToGroupPrivacy add_to_group_privacy int(11);
ALTER TABLE sun.user CHANGE COLUMN SeenMessagePrivacy seen_message_privacy int(11);
ALTER TABLE sun.user CHANGE COLUMN Phone phone bigint(20);
ALTER TABLE sun.user CHANGE COLUMN Email email varchar(250);
ALTER TABLE sun.user CHANGE COLUMN About about varchar(500);
ALTER TABLE sun.user CHANGE COLUMN PasswordHash password_hash varchar(250);
ALTER TABLE sun.user CHANGE COLUMN PasswordSalt password_salt varchar(250);
ALTER TABLE sun.user CHANGE COLUMN PostSeq post_seq int(11);
ALTER TABLE sun.user CHANGE COLUMN FollowersCount followers_count int(10);
ALTER TABLE sun.user CHANGE COLUMN FollowingCount following_count int(10);
ALTER TABLE sun.user CHANGE COLUMN PostsCount posts_count int(10);
ALTER TABLE sun.user CHANGE COLUMN MediaCount media_count int(10);
ALTER TABLE sun.user CHANGE COLUMN PhotoCount photo_count int(11);
ALTER TABLE sun.user CHANGE COLUMN VideoCount video_count int(11);
ALTER TABLE sun.user CHANGE COLUMN GifCount gif_count int(11);
ALTER TABLE sun.user CHANGE COLUMN AudioCount audio_count int(11);
ALTER TABLE sun.user CHANGE COLUMN VoiceCount voice_count int(11);
ALTER TABLE sun.user CHANGE COLUMN FileCount file_count int(11);
ALTER TABLE sun.user CHANGE COLUMN LinkCount link_count int(11);
ALTER TABLE sun.user CHANGE COLUMN BoardCount board_count int(11);
ALTER TABLE sun.user CHANGE COLUMN PinedCount pined_count int(11);
ALTER TABLE sun.user CHANGE COLUMN LikesCount likes_count int(10);
ALTER TABLE sun.user CHANGE COLUMN ResharedCount reshared_count int(10);
ALTER TABLE sun.user CHANGE COLUMN LastPostTime last_post_time int(11);
ALTER TABLE sun.user CHANGE COLUMN CreatedTime created_time int(10);
ALTER TABLE sun.user CHANGE COLUMN VersionTime version_time int(10);
ALTER TABLE sun.user CHANGE COLUMN IsDeleted is_deleted tinyint(4);
ALTER TABLE sun.user CHANGE COLUMN IsBanned is_banned tinyint(4);

/*Table: user_relation  */
ALTER TABLE sun.user_relation CHANGE COLUMN RelNanoId rel_nano_id bigint(20);
ALTER TABLE sun.user_relation CHANGE COLUMN UserId user_id int(11);
ALTER TABLE sun.user_relation CHANGE COLUMN PeerUserId peer_user_id int(11);
ALTER TABLE sun.user_relation CHANGE COLUMN Follwing follwing tinyint(4);
ALTER TABLE sun.user_relation CHANGE COLUMN Followed followed tinyint(4);
ALTER TABLE sun.user_relation CHANGE COLUMN InContacts in_contacts tinyint(4);
ALTER TABLE sun.user_relation CHANGE COLUMN MutualContact mutual_contact tinyint(4);
ALTER TABLE sun.user_relation CHANGE COLUMN IsFavorite is_favorite tinyint(4);
ALTER TABLE sun.user_relation CHANGE COLUMN Notify notify tinyint(4);

/*Table: chat  */
ALTER TABLE sun_chat.chat CHANGE COLUMN ChatId chat_id bigint(20);
ALTER TABLE sun_chat.chat CHANGE COLUMN ChatKey chat_key varchar(50);
ALTER TABLE sun_chat.chat CHANGE COLUMN RoomKey room_key varchar(50);
ALTER TABLE sun_chat.chat CHANGE COLUMN RoomType room_type smallint(6);
ALTER TABLE sun_chat.chat CHANGE COLUMN UserId user_id int(11);
ALTER TABLE sun_chat.chat CHANGE COLUMN PeerUserId peer_user_id int(11);
ALTER TABLE sun_chat.chat CHANGE COLUMN GroupId group_id bigint(20);
ALTER TABLE sun_chat.chat CHANGE COLUMN HashTagId hash_tag_id bigint(20);
ALTER TABLE sun_chat.chat CHANGE COLUMN Title title varchar(100);
ALTER TABLE sun_chat.chat CHANGE COLUMN PinTimeMs pin_time_ms bigint(20);
ALTER TABLE sun_chat.chat CHANGE COLUMN FromMsgId from_msg_id bigint(20);
ALTER TABLE sun_chat.chat CHANGE COLUMN UnseenCount unseen_count int(11);
ALTER TABLE sun_chat.chat CHANGE COLUMN Seq seq int(11);
ALTER TABLE sun_chat.chat CHANGE COLUMN LastMsgId last_msg_id bigint(20);
ALTER TABLE sun_chat.chat CHANGE COLUMN LastMyMsgStatus last_my_msg_status tinyint(4);
ALTER TABLE sun_chat.chat CHANGE COLUMN MyLastSeenSeq my_last_seen_seq int(11);
ALTER TABLE sun_chat.chat CHANGE COLUMN MyLastSeenMsgId my_last_seen_msg_id bigint(20);
ALTER TABLE sun_chat.chat CHANGE COLUMN PeerLastSeenMsgId peer_last_seen_msg_id bigint(20);
ALTER TABLE sun_chat.chat CHANGE COLUMN MyLastDeliveredSeq my_last_delivered_seq int(11);
ALTER TABLE sun_chat.chat CHANGE COLUMN MyLastDeliveredMsgId my_last_delivered_msg_id bigint(20);
ALTER TABLE sun_chat.chat CHANGE COLUMN PeerLastDeliveredMsgId peer_last_delivered_msg_id bigint(20);
ALTER TABLE sun_chat.chat CHANGE COLUMN IsActive is_active tinyint(4);
ALTER TABLE sun_chat.chat CHANGE COLUMN IsLeft is_left tinyint(4);
ALTER TABLE sun_chat.chat CHANGE COLUMN IsCreator is_creator tinyint(4);
ALTER TABLE sun_chat.chat CHANGE COLUMN IsKicked is_kicked tinyint(4);
ALTER TABLE sun_chat.chat CHANGE COLUMN IsAdmin is_admin tinyint(4);
ALTER TABLE sun_chat.chat CHANGE COLUMN IsDeactivated is_deactivated tinyint(4);
ALTER TABLE sun_chat.chat CHANGE COLUMN IsMuted is_muted tinyint(4);
ALTER TABLE sun_chat.chat CHANGE COLUMN MuteUntil mute_until int(11);
ALTER TABLE sun_chat.chat CHANGE COLUMN VersionTimeMs version_time_ms bigint(20);
ALTER TABLE sun_chat.chat CHANGE COLUMN VersionSeq version_seq int(11);
ALTER TABLE sun_chat.chat CHANGE COLUMN OrderTime order_time int(11);
ALTER TABLE sun_chat.chat CHANGE COLUMN CreatedTime created_time int(11);
ALTER TABLE sun_chat.chat CHANGE COLUMN DraftText draft_text varchar(5000);
ALTER TABLE sun_chat.chat CHANGE COLUMN DratReplyToMsgId drat_reply_to_msg_id bigint(20);

/*Table: chat_deleted  */
ALTER TABLE sun_chat.chat_deleted CHANGE COLUMN ChatId chat_id bigint(20);
ALTER TABLE sun_chat.chat_deleted CHANGE COLUMN RoomKey room_key varchar(50);

/*Table: chat_last_message  */
ALTER TABLE sun_chat.chat_last_message CHANGE COLUMN ChatIdGroupId chat_id_group_id varchar(75);
ALTER TABLE sun_chat.chat_last_message CHANGE COLUMN LastMsgPb last_msg_pb blob;

/*Table: chat_user_version  */
ALTER TABLE sun_chat.chat_user_version CHANGE COLUMN UserId user_id int(11);
ALTER TABLE sun_chat.chat_user_version CHANGE COLUMN ChatId chat_id bigint(20);
ALTER TABLE sun_chat.chat_user_version CHANGE COLUMN VersionTimeNano version_time_nano int(11);

/*Table: group  */
ALTER TABLE sun_chat.group CHANGE COLUMN GroupId group_id bigint(20);
ALTER TABLE sun_chat.group CHANGE COLUMN GroupKey group_key varchar(50);
ALTER TABLE sun_chat.group CHANGE COLUMN GroupName group_name varchar(200);
ALTER TABLE sun_chat.group CHANGE COLUMN UserName user_name varchar(30);
ALTER TABLE sun_chat.group CHANGE COLUMN IsSuperGroup is_super_group tinyint(4);
ALTER TABLE sun_chat.group CHANGE COLUMN HashTagId hash_tag_id bigint(20);
ALTER TABLE sun_chat.group CHANGE COLUMN CreatorUserId creator_user_id int(20);
ALTER TABLE sun_chat.group CHANGE COLUMN GroupPrivacy group_privacy tinyint(4);
ALTER TABLE sun_chat.group CHANGE COLUMN HistoryViewAble history_view_able tinyint(4);
ALTER TABLE sun_chat.group CHANGE COLUMN Seq seq bigint(20);
ALTER TABLE sun_chat.group CHANGE COLUMN LastMsgId last_msg_id bigint(20);
ALTER TABLE sun_chat.group CHANGE COLUMN PinedMsgId pined_msg_id bigint(20);
ALTER TABLE sun_chat.group CHANGE COLUMN AvatarRefId avatar_ref_id bigint(20);
ALTER TABLE sun_chat.group CHANGE COLUMN AvatarCount avatar_count int(11);
ALTER TABLE sun_chat.group CHANGE COLUMN About about varchar(1000);
ALTER TABLE sun_chat.group CHANGE COLUMN InviteLinkHash invite_link_hash varchar(150);
ALTER TABLE sun_chat.group CHANGE COLUMN MembersCount members_count int(11);
ALTER TABLE sun_chat.group CHANGE COLUMN AdminsCount admins_count int(11);
ALTER TABLE sun_chat.group CHANGE COLUMN ModeratorCounts moderator_counts int(11);
ALTER TABLE sun_chat.group CHANGE COLUMN SortTime sort_time int(20);
ALTER TABLE sun_chat.group CHANGE COLUMN CreatedTime created_time int(20);
ALTER TABLE sun_chat.group CHANGE COLUMN IsMute is_mute tinyint(4);
ALTER TABLE sun_chat.group CHANGE COLUMN IsDeleted is_deleted tinyint(4);
ALTER TABLE sun_chat.group CHANGE COLUMN IsBanned is_banned tinyint(4);

/*Table: group_member  */
ALTER TABLE sun_chat.group_member CHANGE COLUMN OrderId order_id bigint(20);
ALTER TABLE sun_chat.group_member CHANGE COLUMN GroupId group_id bigint(20);
ALTER TABLE sun_chat.group_member CHANGE COLUMN UserId user_id int(11);
ALTER TABLE sun_chat.group_member CHANGE COLUMN ByUserId by_user_id int(11);
ALTER TABLE sun_chat.group_member CHANGE COLUMN GroupRole group_role tinyint(4);
ALTER TABLE sun_chat.group_member CHANGE COLUMN CreatedTime created_time int(11);

/*Table: group_orderd_user  */
ALTER TABLE sun_chat.group_orderd_user CHANGE COLUMN OrderId order_id bigint(20);
ALTER TABLE sun_chat.group_orderd_user CHANGE COLUMN GroupId group_id bigint(20);
ALTER TABLE sun_chat.group_orderd_user CHANGE COLUMN UserId user_id int(11);

/*Table: group_pined_msg  */
ALTER TABLE sun_chat.group_pined_msg CHANGE COLUMN MsgId msg_id bigint(20);
ALTER TABLE sun_chat.group_pined_msg CHANGE COLUMN MsgPb msg_pb blob;

/*Table: file_msg  */
ALTER TABLE sun_file.file_msg CHANGE COLUMN Id id bigint(20);
ALTER TABLE sun_file.file_msg CHANGE COLUMN AccessHash access_hash int(11);
ALTER TABLE sun_file.file_msg CHANGE COLUMN FileType file_type int(11);
ALTER TABLE sun_file.file_msg CHANGE COLUMN Width width int(11);
ALTER TABLE sun_file.file_msg CHANGE COLUMN Height height int(11);
ALTER TABLE sun_file.file_msg CHANGE COLUMN Extension extension varchar(50);
ALTER TABLE sun_file.file_msg CHANGE COLUMN UserId user_id int(11);
ALTER TABLE sun_file.file_msg CHANGE COLUMN DataThumb data_thumb longblob;
ALTER TABLE sun_file.file_msg CHANGE COLUMN Data data longblob;

/*Table: file_post  */
ALTER TABLE sun_file.file_post CHANGE COLUMN Id id bigint(20);
ALTER TABLE sun_file.file_post CHANGE COLUMN AccessHash access_hash int(11);
ALTER TABLE sun_file.file_post CHANGE COLUMN FileType file_type int(11);
ALTER TABLE sun_file.file_post CHANGE COLUMN Width width int(11);
ALTER TABLE sun_file.file_post CHANGE COLUMN Height height int(11);
ALTER TABLE sun_file.file_post CHANGE COLUMN Extension extension varchar(50);
ALTER TABLE sun_file.file_post CHANGE COLUMN UserId user_id int(11);
ALTER TABLE sun_file.file_post CHANGE COLUMN DataThumb data_thumb longblob;
ALTER TABLE sun_file.file_post CHANGE COLUMN Data data longblob;

/*Table: action_fanout  */
ALTER TABLE sun_meta.action_fanout CHANGE COLUMN OrderId order_id bigint(20);
ALTER TABLE sun_meta.action_fanout CHANGE COLUMN ForUserId for_user_id int(20);
ALTER TABLE sun_meta.action_fanout CHANGE COLUMN ActionId action_id bigint(20);
ALTER TABLE sun_meta.action_fanout CHANGE COLUMN ActorUserId actor_user_id int(11);

/*Table: home_fanout  */
ALTER TABLE sun_meta.home_fanout CHANGE COLUMN OrderId order_id bigint(20);
ALTER TABLE sun_meta.home_fanout CHANGE COLUMN ForUserId for_user_id bigint(20);
ALTER TABLE sun_meta.home_fanout CHANGE COLUMN PostId post_id bigint(20);
ALTER TABLE sun_meta.home_fanout CHANGE COLUMN PostUserId post_user_id bigint(20);
ALTER TABLE sun_meta.home_fanout CHANGE COLUMN ResharedId reshared_id bigint(20);

/*Table: suggested_top_posts  */
ALTER TABLE sun_meta.suggested_top_posts CHANGE COLUMN Id id bigint(20);
ALTER TABLE sun_meta.suggested_top_posts CHANGE COLUMN PostId post_id bigint(20);

/*Table: suggested_user  */
ALTER TABLE sun_meta.suggested_user CHANGE COLUMN Id id int(11);
ALTER TABLE sun_meta.suggested_user CHANGE COLUMN UserId user_id int(11);
ALTER TABLE sun_meta.suggested_user CHANGE COLUMN TargetId target_id int(11);
ALTER TABLE sun_meta.suggested_user CHANGE COLUMN Weight weight float;
ALTER TABLE sun_meta.suggested_user CHANGE COLUMN CreatedTime created_time int(11);

/*Table: push_chat  */
ALTER TABLE sun_push.push_chat CHANGE COLUMN PushId push_id bigint(20);
ALTER TABLE sun_push.push_chat CHANGE COLUMN ToUserId to_user_id int(11);
ALTER TABLE sun_push.push_chat CHANGE COLUMN PushTypeId push_type_id int(11);
ALTER TABLE sun_push.push_chat CHANGE COLUMN RoomKey room_key varchar(75);
ALTER TABLE sun_push.push_chat CHANGE COLUMN ChatKey chat_key varchar(75);
ALTER TABLE sun_push.push_chat CHANGE COLUMN Seq seq int(11);
ALTER TABLE sun_push.push_chat CHANGE COLUMN UnseenCount unseen_count int(11);
ALTER TABLE sun_push.push_chat CHANGE COLUMN FromHighMessageId from_high_message_id bigint(20);
ALTER TABLE sun_push.push_chat CHANGE COLUMN ToLowMessageId to_low_message_id bigint(20);
ALTER TABLE sun_push.push_chat CHANGE COLUMN MessageId message_id bigint(20);
ALTER TABLE sun_push.push_chat CHANGE COLUMN MessageFileId message_file_id bigint(20);
ALTER TABLE sun_push.push_chat CHANGE COLUMN MessagePb message_pb blob;
ALTER TABLE sun_push.push_chat CHANGE COLUMN MessageJson message_json text;
ALTER TABLE sun_push.push_chat CHANGE COLUMN CreatedTime created_time int(11);

/*Table: http_rpc_log  */
ALTER TABLE sun_log.http_rpc_log CHANGE COLUMN Id id int(11);
ALTER TABLE sun_log.http_rpc_log CHANGE COLUMN Time time varchar(100);
ALTER TABLE sun_log.http_rpc_log CHANGE COLUMN MethodFull method_full varchar(100);
ALTER TABLE sun_log.http_rpc_log CHANGE COLUMN MethodParent method_parent varchar(100);
ALTER TABLE sun_log.http_rpc_log CHANGE COLUMN UserId user_id int(11);
ALTER TABLE sun_log.http_rpc_log CHANGE COLUMN SessionId session_id varchar(100);
ALTER TABLE sun_log.http_rpc_log CHANGE COLUMN StatusCode status_code int(11);
ALTER TABLE sun_log.http_rpc_log CHANGE COLUMN InputSize input_size int(11);
ALTER TABLE sun_log.http_rpc_log CHANGE COLUMN OutputSize output_size int(11);
ALTER TABLE sun_log.http_rpc_log CHANGE COLUMN ReqestJson reqest_json text;
ALTER TABLE sun_log.http_rpc_log CHANGE COLUMN ResponseJson response_json text;
ALTER TABLE sun_log.http_rpc_log CHANGE COLUMN ReqestParamJson reqest_param_json text;
ALTER TABLE sun_log.http_rpc_log CHANGE COLUMN ResponseMsgJson response_msg_json text;

/*Table: metric_log  */
ALTER TABLE sun_log.metric_log CHANGE COLUMN Id id int(11);
ALTER TABLE sun_log.metric_log CHANGE COLUMN InstanceId instance_id int(11);
ALTER TABLE sun_log.metric_log CHANGE COLUMN StartFrom start_from varchar(100);
ALTER TABLE sun_log.metric_log CHANGE COLUMN EndTo end_to varchar(100);
ALTER TABLE sun_log.metric_log CHANGE COLUMN StartTime start_time int(11);
ALTER TABLE sun_log.metric_log CHANGE COLUMN Duration duration varchar(100);
ALTER TABLE sun_log.metric_log CHANGE COLUMN MetericsJson meterics_json text;

/*Table: xfile_service_info_log  */
ALTER TABLE sun_log.xfile_service_info_log CHANGE COLUMN Id id bigint(20);
ALTER TABLE sun_log.xfile_service_info_log CHANGE COLUMN InstanceId instance_id int(11);
ALTER TABLE sun_log.xfile_service_info_log CHANGE COLUMN Url url varchar(100);
ALTER TABLE sun_log.xfile_service_info_log CHANGE COLUMN CreatedTime created_time varchar(100);

/*Table: xfile_service_metric_log  */
ALTER TABLE sun_log.xfile_service_metric_log CHANGE COLUMN Id id bigint(20);
ALTER TABLE sun_log.xfile_service_metric_log CHANGE COLUMN InstanceId instance_id int(11);
ALTER TABLE sun_log.xfile_service_metric_log CHANGE COLUMN MetricJson metric_json text;

/*Table: xfile_service_request_log  */
ALTER TABLE sun_log.xfile_service_request_log CHANGE COLUMN Id id bigint(20);
ALTER TABLE sun_log.xfile_service_request_log CHANGE COLUMN LocalSeq local_seq int(11);
ALTER TABLE sun_log.xfile_service_request_log CHANGE COLUMN InstanceId instance_id int(11);
ALTER TABLE sun_log.xfile_service_request_log CHANGE COLUMN Url url varchar(1000);
ALTER TABLE sun_log.xfile_service_request_log CHANGE COLUMN HttpCode http_code int(11);
ALTER TABLE sun_log.xfile_service_request_log CHANGE COLUMN CreatedTime created_time varchar(100);

/*Table: invalidate_cache  */
ALTER TABLE sun_internal.invalidate_cache CHANGE COLUMN OrderId order_id bigint(20);
ALTER TABLE sun_internal.invalidate_cache CHANGE COLUMN CacheKey cache_key varchar(100);

}