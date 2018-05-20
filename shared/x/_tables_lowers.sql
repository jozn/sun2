
/*Table: action  */
ALTER TABLE sun.action CHANGE COLUMN ActionId action_id bigint(20);
ALTER TABLE sun.action CHANGE COLUMN ActorUserId actor_user_id int(11);
ALTER TABLE sun.action CHANGE COLUMN ActionTypeEnum action_type_enum int(11);
ALTER TABLE sun.action CHANGE COLUMN PeerUserId peer_user_id int(20);
ALTER TABLE sun.action CHANGE COLUMN PostId post_id bigint(20);
ALTER TABLE sun.action CHANGE COLUMN CommentId comment_id bigint(20);
ALTER TABLE sun.action CHANGE COLUMN Murmur64Hash murmur64_hash bigint(20);
ALTER TABLE sun.action CHANGE COLUMN CreatedTime created_time int(11);

/*Table: comment  */
ALTER TABLE sun.comment CHANGE COLUMN CommentId comment_id bigint(20);
ALTER TABLE sun.comment CHANGE COLUMN UserId user_id int(11);
ALTER TABLE sun.comment CHANGE COLUMN PostId post_id bigint(20);
ALTER TABLE sun.comment CHANGE COLUMN Text text text;
ALTER TABLE sun.comment CHANGE COLUMN LikesCount likes_count int(11);
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

/*Table: following_list  */
ALTER TABLE sun.following_list CHANGE COLUMN Id id int(11);
ALTER TABLE sun.following_list CHANGE COLUMN UserId user_id int(11);
ALTER TABLE sun.following_list CHANGE COLUMN ListType list_type int(11);
ALTER TABLE sun.following_list CHANGE COLUMN Name name varchar(75);
ALTER TABLE sun.following_list CHANGE COLUMN Count count int(11);
ALTER TABLE sun.following_list CHANGE COLUMN IsAuto is_auto tinyint(1);
ALTER TABLE sun.following_list CHANGE COLUMN IsPimiry is_pimiry tinyint(1);
ALTER TABLE sun.following_list CHANGE COLUMN CreatedTime created_time int(11);

/*Table: following_list_member  */
ALTER TABLE sun.following_list_member CHANGE COLUMN Id id bigint(20);
ALTER TABLE sun.following_list_member CHANGE COLUMN ListId list_id int(11);
ALTER TABLE sun.following_list_member CHANGE COLUMN UserId user_id int(11);
ALTER TABLE sun.following_list_member CHANGE COLUMN FollowedUserId followed_user_id int(11);
ALTER TABLE sun.following_list_member CHANGE COLUMN CreatedTime created_time int(11);

/*Table: following_list_member_removed  */
ALTER TABLE sun.following_list_member_removed CHANGE COLUMN Id id bigint(20);
ALTER TABLE sun.following_list_member_removed CHANGE COLUMN ListId list_id int(11);
ALTER TABLE sun.following_list_member_removed CHANGE COLUMN UserId user_id int(11);
ALTER TABLE sun.following_list_member_removed CHANGE COLUMN UnFollowedUserId un_followed_user_id int(11);
ALTER TABLE sun.following_list_member_removed CHANGE COLUMN UpdatedTime updated_time int(11);

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
ALTER TABLE sun.notify CHANGE COLUMN NotifyTypeEnum notify_type_enum int(10);
ALTER TABLE sun.notify CHANGE COLUMN PostId post_id bigint(20);
ALTER TABLE sun.notify CHANGE COLUMN CommentId comment_id bigint(20);
ALTER TABLE sun.notify CHANGE COLUMN PeerUserId peer_user_id int(11);
ALTER TABLE sun.notify CHANGE COLUMN Murmur64Hash murmur64_hash bigint(20);
ALTER TABLE sun.notify CHANGE COLUMN SeenStatus seen_status int(10);
ALTER TABLE sun.notify CHANGE COLUMN CreatedTime created_time int(10);

/*Table: notify_removed  */
ALTER TABLE sun.notify_removed CHANGE COLUMN Murmur64Hash murmur64_hash bigint(11);
ALTER TABLE sun.notify_removed CHANGE COLUMN ForUserId for_user_id int(11);
ALTER TABLE sun.notify_removed CHANGE COLUMN Id id bigint(20);

/*Table: phone_contacts  */
ALTER TABLE sun.phone_contacts CHANGE COLUMN Id id int(10);
ALTER TABLE sun.phone_contacts CHANGE COLUMN UserId user_id int(10);
ALTER TABLE sun.phone_contacts CHANGE COLUMN Phone phone bigint(20);
ALTER TABLE sun.phone_contacts CHANGE COLUMN PhoneDisplayName phone_display_name varchar(250);
ALTER TABLE sun.phone_contacts CHANGE COLUMN PhoneFamilyName phone_family_name varchar(250);
ALTER TABLE sun.phone_contacts CHANGE COLUMN PhoneNumber phone_number varchar(250);
ALTER TABLE sun.phone_contacts CHANGE COLUMN PhoneNormalizedNumber phone_normalized_number varchar(15);
ALTER TABLE sun.phone_contacts CHANGE COLUMN PhoneContactRowId phone_contact_row_id int(10);
ALTER TABLE sun.phone_contacts CHANGE COLUMN DeviceUuidId device_uuid_id int(10);
ALTER TABLE sun.phone_contacts CHANGE COLUMN CreatedTime created_time int(10);

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
ALTER TABLE sun.post_count CHANGE COLUMN ViewsCount views_count bigint(20);

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
ALTER TABLE sun.post_mentioned CHANGE COLUMN PostTypeEnum post_type_enum tinyint(11);
ALTER TABLE sun.post_mentioned CHANGE COLUMN PostCategoryEnum post_category_enum tinyint(4);
ALTER TABLE sun.post_mentioned CHANGE COLUMN CreatedTime created_time int(11);

/*Table: post_reshared  */
ALTER TABLE sun.post_reshared CHANGE COLUMN ResharedId reshared_id bigint(20);
ALTER TABLE sun.post_reshared CHANGE COLUMN ByUserId by_user_id int(11);
ALTER TABLE sun.post_reshared CHANGE COLUMN PostId post_id bigint(20);
ALTER TABLE sun.post_reshared CHANGE COLUMN PostUserId post_user_id int(11);
ALTER TABLE sun.post_reshared CHANGE COLUMN PostTypeEnum post_type_enum tinyint(4);
ALTER TABLE sun.post_reshared CHANGE COLUMN PostCategoryEnum post_category_enum tinyint(4);
ALTER TABLE sun.post_reshared CHANGE COLUMN CreatedTime created_time int(11);

/*Table: search_clicked  */
ALTER TABLE sun.search_clicked CHANGE COLUMN Id id bigint(20);
ALTER TABLE sun.search_clicked CHANGE COLUMN Query query varchar(100);
ALTER TABLE sun.search_clicked CHANGE COLUMN ClickType click_type int(11);
ALTER TABLE sun.search_clicked CHANGE COLUMN TargetId target_id int(11);
ALTER TABLE sun.search_clicked CHANGE COLUMN UserId user_id int(11);
ALTER TABLE sun.search_clicked CHANGE COLUMN CreatedTime created_time int(11);

/*Table: session  */
ALTER TABLE sun.session CHANGE COLUMN Id id bigint(20);
ALTER TABLE sun.session CHANGE COLUMN SessionUuid session_uuid varchar(75);
ALTER TABLE sun.session CHANGE COLUMN UserId user_id int(11);
ALTER TABLE sun.session CHANGE COLUMN LastIpAddress last_ip_address varchar(75);
ALTER TABLE sun.session CHANGE COLUMN AppVersion app_version smallint(6);
ALTER TABLE sun.session CHANGE COLUMN ActiveTime active_time int(11);
ALTER TABLE sun.session CHANGE COLUMN CreatedTime created_time int(11);

/*Table: setting_client  */
ALTER TABLE sun.setting_client CHANGE COLUMN UserId user_id int(11);
ALTER TABLE sun.setting_client CHANGE COLUMN AutoDownloadWifiVoice auto_download_wifi_voice int(11);
ALTER TABLE sun.setting_client CHANGE COLUMN AutoDownloadWifiImage auto_download_wifi_image int(11);
ALTER TABLE sun.setting_client CHANGE COLUMN AutoDownloadWifiVideo auto_download_wifi_video int(11);
ALTER TABLE sun.setting_client CHANGE COLUMN AutoDownloadWifiFile auto_download_wifi_file int(11);
ALTER TABLE sun.setting_client CHANGE COLUMN AutoDownloadWifiMusic auto_download_wifi_music int(11);
ALTER TABLE sun.setting_client CHANGE COLUMN AutoDownloadWifiGif auto_download_wifi_gif int(11);
ALTER TABLE sun.setting_client CHANGE COLUMN AutoDownloadCellularVoice auto_download_cellular_voice int(11);
ALTER TABLE sun.setting_client CHANGE COLUMN AutoDownloadCellularImage auto_download_cellular_image int(11);
ALTER TABLE sun.setting_client CHANGE COLUMN AutoDownloadCellularVideo auto_download_cellular_video int(11);
ALTER TABLE sun.setting_client CHANGE COLUMN AutoDownloadCellularFile auto_download_cellular_file int(11);
ALTER TABLE sun.setting_client CHANGE COLUMN AutoDownloadCellularMusic auto_download_cellular_music int(11);
ALTER TABLE sun.setting_client CHANGE COLUMN AutoDownloadCellularGif auto_download_cellular_gif int(11);
ALTER TABLE sun.setting_client CHANGE COLUMN AutoDownloadRoamingVoice auto_download_roaming_voice int(11);
ALTER TABLE sun.setting_client CHANGE COLUMN AutoDownloadRoamingImage auto_download_roaming_image int(11);
ALTER TABLE sun.setting_client CHANGE COLUMN AutoDownloadRoamingVideo auto_download_roaming_video int(11);
ALTER TABLE sun.setting_client CHANGE COLUMN AutoDownloadRoamingFile auto_download_roaming_file int(11);
ALTER TABLE sun.setting_client CHANGE COLUMN AutoDownloadRoamingMusic auto_download_roaming_music int(11);
ALTER TABLE sun.setting_client CHANGE COLUMN AutoDownloadRoamingGif auto_download_roaming_gif int(11);
ALTER TABLE sun.setting_client CHANGE COLUMN SaveToGallery save_to_gallery int(11);

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
ALTER TABLE sun.user CHANGE COLUMN UserTypeEnum user_type_enum int(11);
ALTER TABLE sun.user CHANGE COLUMN UserLevelEnum user_level_enum int(11);
ALTER TABLE sun.user CHANGE COLUMN AvatarId avatar_id bigint(20);
ALTER TABLE sun.user CHANGE COLUMN ProfilePrivacyEnum profile_privacy_enum tinyint(10);
ALTER TABLE sun.user CHANGE COLUMN Phone phone bigint(20);
ALTER TABLE sun.user CHANGE COLUMN About about varchar(500);
ALTER TABLE sun.user CHANGE COLUMN Email email varchar(250);
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
ALTER TABLE sun.user CHANGE COLUMN LastActionTime last_action_time int(11);
ALTER TABLE sun.user CHANGE COLUMN LastPostTime last_post_time int(11);
ALTER TABLE sun.user CHANGE COLUMN PrimaryFollowingList primary_following_list int(10);
ALTER TABLE sun.user CHANGE COLUMN CreatedSe created_se int(10);
ALTER TABLE sun.user CHANGE COLUMN UpdatedMs updated_ms bigint(10);
ALTER TABLE sun.user CHANGE COLUMN OnlinePrivacyEnum online_privacy_enum int(11);
ALTER TABLE sun.user CHANGE COLUMN LastActivityTime last_activity_time int(11);
ALTER TABLE sun.user CHANGE COLUMN Phone2 phone2 varchar(250);

/*Table: user_meta_info  */
ALTER TABLE sun.user_meta_info CHANGE COLUMN Id id int(11);
ALTER TABLE sun.user_meta_info CHANGE COLUMN UserId user_id int(11);
ALTER TABLE sun.user_meta_info CHANGE COLUMN IsNotificationDirty is_notification_dirty tinyint(4);
ALTER TABLE sun.user_meta_info CHANGE COLUMN LastUserRecGen last_user_rec_gen int(11);

/*Table: user_password  */
ALTER TABLE sun.user_password CHANGE COLUMN UserId user_id int(11);
ALTER TABLE sun.user_password CHANGE COLUMN Password password varchar(100);
ALTER TABLE sun.user_password CHANGE COLUMN CreatedTime created_time int(11);

/*Table: chat  */
ALTER TABLE sun_chat.chat CHANGE COLUMN ChatKey chat_key varchar(50);
ALTER TABLE sun_chat.chat CHANGE COLUMN RoomKey room_key varchar(50);
ALTER TABLE sun_chat.chat CHANGE COLUMN RoomTypeEnum room_type_enum smallint(6);
ALTER TABLE sun_chat.chat CHANGE COLUMN UserId user_id int(11);
ALTER TABLE sun_chat.chat CHANGE COLUMN PeerUserId peer_user_id int(11);
ALTER TABLE sun_chat.chat CHANGE COLUMN GroupId group_id bigint(20);
ALTER TABLE sun_chat.chat CHANGE COLUMN CreatedTime created_time int(11);
ALTER TABLE sun_chat.chat CHANGE COLUMN Seq seq int(11);
ALTER TABLE sun_chat.chat CHANGE COLUMN SeenSeq seen_seq int(11);
ALTER TABLE sun_chat.chat CHANGE COLUMN UpdatedMs updated_ms bigint(20);

/*Table: chat_last_message  */
ALTER TABLE sun_chat.chat_last_message CHANGE COLUMN ChatKey chat_key varchar(75);
ALTER TABLE sun_chat.chat_last_message CHANGE COLUMN ForUserId for_user_id int(11);
ALTER TABLE sun_chat.chat_last_message CHANGE COLUMN LastMsgPb last_msg_pb blob;
ALTER TABLE sun_chat.chat_last_message CHANGE COLUMN LastMsgJson last_msg_json text;

/*Table: direct_message  */
ALTER TABLE sun_chat.direct_message CHANGE COLUMN ChatKey chat_key varchar(75);
ALTER TABLE sun_chat.direct_message CHANGE COLUMN MessageId message_id bigint(20);
ALTER TABLE sun_chat.direct_message CHANGE COLUMN RoomKey room_key varchar(70);
ALTER TABLE sun_chat.direct_message CHANGE COLUMN UserId user_id int(20);
ALTER TABLE sun_chat.direct_message CHANGE COLUMN MessageFileId message_file_id bigint(20);
ALTER TABLE sun_chat.direct_message CHANGE COLUMN MessageTypeEnum message_type_enum tinyint(4);
ALTER TABLE sun_chat.direct_message CHANGE COLUMN Text text varchar(16000);
ALTER TABLE sun_chat.direct_message CHANGE COLUMN CreatedTime created_time int(11);
ALTER TABLE sun_chat.direct_message CHANGE COLUMN Seq seq int(11);
ALTER TABLE sun_chat.direct_message CHANGE COLUMN DeliviryStatusEnum deliviry_status_enum tinyint(4);
ALTER TABLE sun_chat.direct_message CHANGE COLUMN ExtraPB extra_pb blob;

/*Table: group  */
ALTER TABLE sun_chat.group CHANGE COLUMN GroupId group_id bigint(20);
ALTER TABLE sun_chat.group CHANGE COLUMN GroupName group_name varchar(200);
ALTER TABLE sun_chat.group CHANGE COLUMN MembersCount members_count int(11);
ALTER TABLE sun_chat.group CHANGE COLUMN GroupPrivacyEnum group_privacy_enum tinyint(4);
ALTER TABLE sun_chat.group CHANGE COLUMN CreatorUserId creator_user_id int(20);
ALTER TABLE sun_chat.group CHANGE COLUMN CreatedTime created_time int(20);
ALTER TABLE sun_chat.group CHANGE COLUMN UpdatedMs updated_ms bigint(20);
ALTER TABLE sun_chat.group CHANGE COLUMN CurrentSeq current_seq int(11);

/*Table: group_member  */
ALTER TABLE sun_chat.group_member CHANGE COLUMN Id id bigint(20);
ALTER TABLE sun_chat.group_member CHANGE COLUMN GroupId group_id bigint(20);
ALTER TABLE sun_chat.group_member CHANGE COLUMN GroupKey group_key varchar(100);
ALTER TABLE sun_chat.group_member CHANGE COLUMN UserId user_id int(11);
ALTER TABLE sun_chat.group_member CHANGE COLUMN ByUserId by_user_id int(11);
ALTER TABLE sun_chat.group_member CHANGE COLUMN GroupRoleEnumId group_role_enum_id tinyint(4);
ALTER TABLE sun_chat.group_member CHANGE COLUMN CreatedTime created_time int(11);

/*Table: group_message  */
ALTER TABLE sun_chat.group_message CHANGE COLUMN MessageId message_id bigint(20);
ALTER TABLE sun_chat.group_message CHANGE COLUMN RoomKey room_key varchar(70);
ALTER TABLE sun_chat.group_message CHANGE COLUMN UserId user_id int(20);
ALTER TABLE sun_chat.group_message CHANGE COLUMN MessageFileId message_file_id bigint(20);
ALTER TABLE sun_chat.group_message CHANGE COLUMN MessageTypeEnum message_type_enum tinyint(4);
ALTER TABLE sun_chat.group_message CHANGE COLUMN Text text text;
ALTER TABLE sun_chat.group_message CHANGE COLUMN CreatedMs created_ms bigint(20);
ALTER TABLE sun_chat.group_message CHANGE COLUMN DeliveryStatusEnum delivery_status_enum tinyint(4);

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

/*Table: chat_sync2  */
ALTER TABLE sun_push.chat_sync2 CHANGE COLUMN sync_id sync_id bigint(20);
ALTER TABLE sun_push.chat_sync2 CHANGE COLUMN to_user_id to_user_id int(11);
ALTER TABLE sun_push.chat_sync2 CHANGE COLUMN chat_sync_type_id chat_sync_type_id int(11);
ALTER TABLE sun_push.chat_sync2 CHANGE COLUMN room_key room_key varchar(75);
ALTER TABLE sun_push.chat_sync2 CHANGE COLUMN chat_key chat_key varchar(75);
ALTER TABLE sun_push.chat_sync2 CHANGE COLUMN from_high_message_id from_high_message_id bigint(20);
ALTER TABLE sun_push.chat_sync2 CHANGE COLUMN to_low_message_id to_low_message_id bigint(20);
ALTER TABLE sun_push.chat_sync2 CHANGE COLUMN message_id message_id bigint(20);
ALTER TABLE sun_push.chat_sync2 CHANGE COLUMN message_pb message_pb blob;
ALTER TABLE sun_push.chat_sync2 CHANGE COLUMN message_json message_json text;
ALTER TABLE sun_push.chat_sync2 CHANGE COLUMN created_time created_time int(11);

/*Table: lower_table  */
ALTER TABLE sun_push.lower_table CHANGE COLUMN id id int(11);
ALTER TABLE sun_push.lower_table CHANGE COLUMN text text text;
ALTER TABLE sun_push.lower_table CHANGE COLUMN time_stamp time_stamp int(11);
ALTER TABLE sun_push.lower_table CHANGE COLUMN any_thing_more_ any_thing_more_ int(11);

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

/*Table: push_chat2  */
ALTER TABLE sun_push.push_chat2 CHANGE COLUMN sync_id sync_id bigint(20);
ALTER TABLE sun_push.push_chat2 CHANGE COLUMN to_user_id to_user_id int(11);
ALTER TABLE sun_push.push_chat2 CHANGE COLUMN chat_sync_type_id chat_sync_type_id int(11);
ALTER TABLE sun_push.push_chat2 CHANGE COLUMN room_key room_key varchar(75);
ALTER TABLE sun_push.push_chat2 CHANGE COLUMN chat_key chat_key varchar(75);
ALTER TABLE sun_push.push_chat2 CHANGE COLUMN from_high_message_id from_high_message_id bigint(20);
ALTER TABLE sun_push.push_chat2 CHANGE COLUMN to_low_message_id to_low_message_id bigint(20);
ALTER TABLE sun_push.push_chat2 CHANGE COLUMN message_id message_id bigint(20);
ALTER TABLE sun_push.push_chat2 CHANGE COLUMN message_pb message_pb blob;
ALTER TABLE sun_push.push_chat2 CHANGE COLUMN message_json message_json text;
ALTER TABLE sun_push.push_chat2 CHANGE COLUMN created_time created_time int(11);

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

/*Table: accounts  */
ALTER TABLE suncdb.accounts CHANGE COLUMN id id INT;
ALTER TABLE suncdb.accounts CHANGE COLUMN balance balance DECIMAL;

/*Table: post_cdb  */
ALTER TABLE suncdb.post_cdb CHANGE COLUMN post_id post_id INT;
ALTER TABLE suncdb.post_cdb CHANGE COLUMN user_id user_id INT;
ALTER TABLE suncdb.post_cdb CHANGE COLUMN post_type_enum post_type_enum INT;
ALTER TABLE suncdb.post_cdb CHANGE COLUMN post_category_enum post_category_enum INT;
ALTER TABLE suncdb.post_cdb CHANGE COLUMN media_id media_id INT;
ALTER TABLE suncdb.post_cdb CHANGE COLUMN post_key post_key STRING;
ALTER TABLE suncdb.post_cdb CHANGE COLUMN text text STRING;
ALTER TABLE suncdb.post_cdb CHANGE COLUMN rich_text rich_text STRING;
ALTER TABLE suncdb.post_cdb CHANGE COLUMN media_count media_count INT;
ALTER TABLE suncdb.post_cdb CHANGE COLUMN shared_to shared_to INT;
ALTER TABLE suncdb.post_cdb CHANGE COLUMN disable_comment disable_comment INT;
ALTER TABLE suncdb.post_cdb CHANGE COLUMN source source INT;
ALTER TABLE suncdb.post_cdb CHANGE COLUMN has_tag has_tag INT;
ALTER TABLE suncdb.post_cdb CHANGE COLUMN seq seq INT;
ALTER TABLE suncdb.post_cdb CHANGE COLUMN comments_count comments_count INT;
ALTER TABLE suncdb.post_cdb CHANGE COLUMN likes_count likes_count INT;
ALTER TABLE suncdb.post_cdb CHANGE COLUMN views_count views_count INT;
ALTER TABLE suncdb.post_cdb CHANGE COLUMN edited_time edited_time INT;
ALTER TABLE suncdb.post_cdb CHANGE COLUMN created_time created_time INT;
ALTER TABLE suncdb.post_cdb CHANGE COLUMN re_shared_post_id re_shared_post_id INT;

}