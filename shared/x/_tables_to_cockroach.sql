
/*Table: action  */
CREATE TABLE IF NOT EXISTS sun.action (
    action_id int PRIMARY KEY NOT NULL ,
    actor_user_id int NOT NULL ,
    action_type_enum int NOT NULL ,
    peer_user_id int NOT NULL ,
    post_id int NOT NULL ,
    comment_id int NOT NULL ,
    murmur64_hash int NOT NULL ,
    created_time int NOT NULL ,
);

/*Table: comment  */
CREATE TABLE IF NOT EXISTS sun.comment (
    comment_id int PRIMARY KEY NOT NULL ,
    user_id int NOT NULL ,
    post_id int NOT NULL ,
    text string NOT NULL ,
    likes_count int NOT NULL ,
    created_time int NOT NULL ,
);

/*Table: comment_deleted  */
CREATE TABLE IF NOT EXISTS sun.comment_deleted (
    comment_id int PRIMARY KEY NOT NULL ,
    user_id int NOT NULL ,
);

/*Table: event  */
CREATE TABLE IF NOT EXISTS sun.event (
    event_id int PRIMARY KEY NOT NULL ,
    event_type int NOT NULL ,
    by_user_id int NOT NULL ,
    peer_user_id int NOT NULL ,
    post_id int NOT NULL ,
    comment_id int NOT NULL ,
    action_id int NOT NULL ,
    murmur64_hash int NOT NULL ,
    chat_key string NOT NULL ,
    message_id int NOT NULL ,
    re_shared_id int NOT NULL ,
);

/*Table: following_list  */
CREATE TABLE IF NOT EXISTS sun.following_list (
    id int PRIMARY KEY NOT NULL ,
    user_id int NOT NULL ,
    list_type int NOT NULL ,
    name string NOT NULL ,
    count int NOT NULL ,
    is_auto int NOT NULL ,
    is_pimiry int NOT NULL ,
    created_time int NOT NULL ,
);

/*Table: following_list_member  */
CREATE TABLE IF NOT EXISTS sun.following_list_member (
    id int PRIMARY KEY NOT NULL ,
    list_id int NOT NULL ,
    user_id int NOT NULL ,
    followed_user_id int NOT NULL ,
    created_time int NOT NULL ,
);

/*Table: following_list_member_removed  */
CREATE TABLE IF NOT EXISTS sun.following_list_member_removed (
    id int PRIMARY KEY NOT NULL ,
    list_id int NOT NULL ,
    user_id int NOT NULL ,
    un_followed_user_id int NOT NULL ,
    updated_time int NOT NULL ,
);

/*Table: likes  */
CREATE TABLE IF NOT EXISTS sun.likes (
    id int PRIMARY KEY NOT NULL ,
    post_id int NOT NULL ,
    post_type_enum int NOT NULL ,
    user_id int NOT NULL ,
    like_enum int NOT NULL ,
    created_time int NOT NULL ,
);

/*Table: notify  */
CREATE TABLE IF NOT EXISTS sun.notify (
    notify_id int PRIMARY KEY NOT NULL ,
    for_user_id int NOT NULL ,
    actor_user_id int NOT NULL ,
    notify_type_enum int NOT NULL ,
    post_id int NOT NULL ,
    comment_id int NOT NULL ,
    peer_user_id int NOT NULL ,
    murmur64_hash int NOT NULL ,
    seen_status int NOT NULL ,
    created_time int NOT NULL ,
);

/*Table: notify_removed  */
CREATE TABLE IF NOT EXISTS sun.notify_removed (
    murmur64_hash int NOT NULL ,
    for_user_id int NOT NULL ,
    id int PRIMARY KEY NOT NULL ,
);

/*Table: phone_contacts  */
CREATE TABLE IF NOT EXISTS sun.phone_contacts (
    id int PRIMARY KEY NOT NULL ,
    user_id int NOT NULL ,
    phone int NOT NULL ,
    phone_display_name string NOT NULL ,
    phone_family_name string NOT NULL ,
    phone_number string NOT NULL ,
    phone_normalized_number string NOT NULL ,
    phone_contact_row_id int NOT NULL ,
    device_uuid_id int NOT NULL ,
    created_time int NOT NULL ,
);

/*Table: post  */
CREATE TABLE IF NOT EXISTS sun.post (
    post_id int PRIMARY KEY NOT NULL ,
    user_id int NOT NULL ,
    post_type_enum int NOT NULL ,
    post_category_enum int NOT NULL ,
    media_id int NOT NULL ,
    post_key string NOT NULL ,
    text string NOT NULL ,
    rich_text string NOT NULL ,
    media_count int NOT NULL ,
    shared_to int NOT NULL ,
    disable_comment int NOT NULL ,
    source int NOT NULL ,
    has_tag int NOT NULL ,
    seq int NOT NULL ,
    comments_count int NOT NULL ,
    likes_count int NOT NULL ,
    views_count int NOT NULL ,
    edited_time int NOT NULL ,
    created_time int NOT NULL ,
    re_shared_post_id int NOT NULL ,
);

/*Table: post_copy  */
CREATE TABLE IF NOT EXISTS sun.post_copy (
    post_id int PRIMARY KEY NOT NULL ,
    user_id int ,
    post_type_enum int ,
    post_category_enum int ,
    media_id int ,
    post_key string ,
    text string ,
    rich_text string ,
    media_count int ,
    shared_to int ,
    disable_comment int ,
    source int ,
    has_tag int ,
    seq int ,
    comments_count int ,
    likes_count int ,
    views_count int ,
    edited_time int ,
    created_time int ,
    re_shared_post_id int ,
);

/*Table: post_count  */
CREATE TABLE IF NOT EXISTS sun.post_count (
    post_id int PRIMARY KEY NOT NULL ,
    views_count int ,
);

/*Table: post_deleted  */
CREATE TABLE IF NOT EXISTS sun.post_deleted (
    post_id int PRIMARY KEY NOT NULL ,
    user_id int NOT NULL ,
);

/*Table: post_keys  */
CREATE TABLE IF NOT EXISTS sun.post_keys (
    id int PRIMARY KEY NOT NULL ,
    post_key_str string UNIQUE NOT NULL ,
    used int NOT NULL ,
);

/*Table: post_link  */
CREATE TABLE IF NOT EXISTS sun.post_link (
    link_id int PRIMARY KEY NOT NULL ,
    link_url string NOT NULL ,
);

/*Table: post_media  */
CREATE TABLE IF NOT EXISTS sun.post_media (
    media_id int PRIMARY KEY NOT NULL ,
    user_id int NOT NULL ,
    post_id int NOT NULL ,
    album_id int NOT NULL ,
    media_type_enum int NOT NULL ,
    width int NOT NULL ,
    height int NOT NULL ,
    size int NOT NULL ,
    duration int NOT NULL ,
    extension string NOT NULL ,
    md5_hash string NOT NULL ,
    color string NOT NULL ,
    created_time int NOT NULL ,
    view_count int NOT NULL ,
    extra string NOT NULL ,
);

/*Table: post_mentioned  */
CREATE TABLE IF NOT EXISTS sun.post_mentioned (
    mentioned_id int PRIMARY KEY NOT NULL ,
    for_user_id int NOT NULL ,
    post_id int NOT NULL ,
    post_user_id int NOT NULL ,
    post_type_enum int NOT NULL ,
    post_category_enum int NOT NULL ,
    created_time int NOT NULL ,
);

/*Table: post_reshared  */
CREATE TABLE IF NOT EXISTS sun.post_reshared (
    reshared_id int PRIMARY KEY NOT NULL ,
    by_user_id int NOT NULL ,
    post_id int NOT NULL ,
    post_user_id int NOT NULL ,
    post_type_enum int NOT NULL ,
    post_category_enum int NOT NULL ,
    created_time int NOT NULL ,
);

/*Table: search_clicked  */
CREATE TABLE IF NOT EXISTS sun.search_clicked (
    id int PRIMARY KEY NOT NULL ,
    query string NOT NULL ,
    click_type int NOT NULL ,
    target_id int NOT NULL ,
    user_id int NOT NULL ,
    created_time int NOT NULL ,
);

/*Table: session  */
CREATE TABLE IF NOT EXISTS sun.session (
    id int PRIMARY KEY NOT NULL ,
    session_uuid string NOT NULL ,
    user_id int NOT NULL ,
    last_ip_address string NOT NULL ,
    app_version int NOT NULL ,
    active_time int NOT NULL ,
    created_time int NOT NULL ,
);

/*Table: setting_client  */
CREATE TABLE IF NOT EXISTS sun.setting_client (
    user_id int PRIMARY KEY NOT NULL ,
    auto_download_wifi_voice int NOT NULL ,
    auto_download_wifi_image int NOT NULL ,
    auto_download_wifi_video int NOT NULL ,
    auto_download_wifi_file int NOT NULL ,
    auto_download_wifi_music int NOT NULL ,
    auto_download_wifi_gif int NOT NULL ,
    auto_download_cellular_voice int NOT NULL ,
    auto_download_cellular_image int NOT NULL ,
    auto_download_cellular_video int NOT NULL ,
    auto_download_cellular_file int NOT NULL ,
    auto_download_cellular_music int NOT NULL ,
    auto_download_cellular_gif int NOT NULL ,
    auto_download_roaming_voice int NOT NULL ,
    auto_download_roaming_image int NOT NULL ,
    auto_download_roaming_video int NOT NULL ,
    auto_download_roaming_file int NOT NULL ,
    auto_download_roaming_music int NOT NULL ,
    auto_download_roaming_gif int NOT NULL ,
    save_to_gallery int NOT NULL ,
);

/*Table: setting_notifications  */
CREATE TABLE IF NOT EXISTS sun.setting_notifications (
    user_id int PRIMARY KEY NOT NULL ,
    social_led_on int NOT NULL ,
    social_led_color string NOT NULL ,
    reqest_to_follow_you int NOT NULL ,
    followed_you int NOT NULL ,
    accpted_your_follow_request int NOT NULL ,
    your_post_liked int NOT NULL ,
    your_post_commented int NOT NULL ,
    menthened_you_in_post int NOT NULL ,
    menthened_you_in_comment int NOT NULL ,
    your_contacts_joined int NOT NULL ,
    direct_message int NOT NULL ,
    direct_alert int NOT NULL ,
    direct_perview int NOT NULL ,
    direct_led_on int NOT NULL ,
    direct_led_color int NOT NULL ,
    direct_vibrate int NOT NULL ,
    direct_popup int NOT NULL ,
    direct_sound int NOT NULL ,
    direct_priority int NOT NULL ,
);

/*Table: tag  */
CREATE TABLE IF NOT EXISTS sun.tag (
    tag_id int PRIMARY KEY NOT NULL ,
    name string UNIQUE NOT NULL ,
    count int NOT NULL ,
    tag_status_enum int NOT NULL ,
    created_time int NOT NULL ,
);

/*Table: tag_post  */
CREATE TABLE IF NOT EXISTS sun.tag_post (
    id int PRIMARY KEY NOT NULL ,
    tag_id int NOT NULL ,
    post_id int NOT NULL ,
    post_type_enum int NOT NULL ,
    post_category_enum int NOT NULL ,
    created_time int NOT NULL ,
);

/*Table: trigger_log  */
CREATE TABLE IF NOT EXISTS sun.trigger_log (
    id int PRIMARY KEY NOT NULL ,
    model_name string NOT NULL ,
    change_type string NOT NULL ,
    target_id int NOT NULL ,
    target_str string NOT NULL ,
    created_se int NOT NULL ,
);

/*Table: user  */
CREATE TABLE IF NOT EXISTS sun.user (
    user_id int PRIMARY KEY NOT NULL ,
    user_name string UNIQUE NOT NULL ,
    user_name_lower string NOT NULL ,
    first_name string NOT NULL ,
    last_name string NOT NULL ,
    user_type_enum int NOT NULL ,
    user_level_enum int NOT NULL ,
    avatar_id int NOT NULL ,
    profile_privacy_enum int NOT NULL ,
    phone int NOT NULL ,
    about string NOT NULL ,
    email string UNIQUE NOT NULL ,
    password_hash string NOT NULL ,
    password_salt string NOT NULL ,
    post_seq int NOT NULL ,
    followers_count int NOT NULL ,
    following_count int NOT NULL ,
    posts_count int NOT NULL ,
    media_count int NOT NULL ,
    photo_count int NOT NULL ,
    video_count int NOT NULL ,
    gif_count int NOT NULL ,
    audio_count int NOT NULL ,
    voice_count int NOT NULL ,
    file_count int NOT NULL ,
    link_count int NOT NULL ,
    board_count int NOT NULL ,
    pined_count int NOT NULL ,
    likes_count int NOT NULL ,
    reshared_count int NOT NULL ,
    last_action_time int NOT NULL ,
    last_post_time int NOT NULL ,
    primary_following_list int NOT NULL ,
    created_se int NOT NULL ,
    updated_ms int NOT NULL ,
    online_privacy_enum int NOT NULL ,
    last_activity_time int NOT NULL ,
    phone2 string NOT NULL ,
);

/*Table: user_meta_info  */
CREATE TABLE IF NOT EXISTS sun.user_meta_info (
    id int PRIMARY KEY NOT NULL ,
    user_id int UNIQUE NOT NULL ,
    is_notification_dirty int NOT NULL ,
    last_user_rec_gen int NOT NULL ,
);

/*Table: user_password  */
CREATE TABLE IF NOT EXISTS sun.user_password (
    user_id int PRIMARY KEY NOT NULL ,
    password string NOT NULL ,
    created_time int NOT NULL ,
);

/*Table: chat  */
CREATE TABLE IF NOT EXISTS sun_chat.chat (
    chat_key string PRIMARY KEY NOT NULL ,
    room_key string NOT NULL ,
    room_type_enum int NOT NULL ,
    user_id int NOT NULL ,
    peer_user_id int NOT NULL ,
    group_id int NOT NULL ,
    created_time int NOT NULL ,
    seq int NOT NULL ,
    seen_seq int NOT NULL ,
    updated_ms int NOT NULL ,
);

/*Table: chat_last_message  */
CREATE TABLE IF NOT EXISTS sun_chat.chat_last_message (
    chat_key string PRIMARY KEY NOT NULL ,
    for_user_id int NOT NULL ,
    last_msg_pb bytes NOT NULL ,
    last_msg_json string NOT NULL ,
);

/*Table: direct_message  */
CREATE TABLE IF NOT EXISTS sun_chat.direct_message (
    chat_key string NOT NULL ,
    message_id int PRIMARY KEY NOT NULL ,
    room_key string NOT NULL ,
    user_id int NOT NULL ,
    message_file_id int NOT NULL ,
    message_type_enum int NOT NULL ,
    text string NOT NULL ,
    created_time int NOT NULL ,
    seq int NOT NULL ,
    deliviry_status_enum int NOT NULL ,
    extra_pb bytes NOT NULL ,
);

/*Table: group  */
CREATE TABLE IF NOT EXISTS sun_chat.group (
    group_id int PRIMARY KEY NOT NULL ,
    group_name string NOT NULL ,
    members_count int NOT NULL ,
    group_privacy_enum int NOT NULL ,
    creator_user_id int NOT NULL ,
    created_time int NOT NULL ,
    updated_ms int NOT NULL ,
    current_seq int NOT NULL ,
);

/*Table: group_member  */
CREATE TABLE IF NOT EXISTS sun_chat.group_member (
    id int PRIMARY KEY NOT NULL ,
    group_id int NOT NULL ,
    group_key string NOT NULL ,
    user_id int NOT NULL ,
    by_user_id int NOT NULL ,
    group_role_enum_id int NOT NULL ,
    created_time int NOT NULL ,
);

/*Table: group_message  */
CREATE TABLE IF NOT EXISTS sun_chat.group_message (
    message_id int PRIMARY KEY NOT NULL ,
    room_key string NOT NULL ,
    user_id int NOT NULL ,
    message_file_id int NOT NULL ,
    message_type_enum int NOT NULL ,
    text string NOT NULL ,
    created_ms int NOT NULL ,
    delivery_status_enum int NOT NULL ,
);

/*Table: file_msg  */
CREATE TABLE IF NOT EXISTS sun_file.file_msg (
    id int PRIMARY KEY NOT NULL ,
    access_hash int NOT NULL ,
    file_type int NOT NULL ,
    width int NOT NULL ,
    height int NOT NULL ,
    extension string NOT NULL ,
    user_id int NOT NULL ,
    data_thumb bytes NOT NULL ,
    data bytes NOT NULL ,
);

/*Table: file_post  */
CREATE TABLE IF NOT EXISTS sun_file.file_post (
    id int PRIMARY KEY NOT NULL ,
    access_hash int NOT NULL ,
    file_type int NOT NULL ,
    width int NOT NULL ,
    height int NOT NULL ,
    extension string NOT NULL ,
    user_id int NOT NULL ,
    data_thumb bytes NOT NULL ,
    data bytes NOT NULL ,
);

/*Table: action_fanout  */
CREATE TABLE IF NOT EXISTS sun_meta.action_fanout (
    order_id int PRIMARY KEY NOT NULL ,
    for_user_id int NOT NULL ,
    action_id int NOT NULL ,
    actor_user_id int NOT NULL ,
);

/*Table: home_fanout  */
CREATE TABLE IF NOT EXISTS sun_meta.home_fanout (
    order_id int PRIMARY KEY NOT NULL ,
    for_user_id int NOT NULL ,
    post_id int NOT NULL ,
    post_user_id int NOT NULL ,
    reshared_id int NOT NULL ,
);

/*Table: suggested_top_posts  */
CREATE TABLE IF NOT EXISTS sun_meta.suggested_top_posts (
    id int PRIMARY KEY NOT NULL ,
    post_id int NOT NULL ,
);

/*Table: suggested_user  */
CREATE TABLE IF NOT EXISTS sun_meta.suggested_user (
    id int PRIMARY KEY NOT NULL ,
    user_id int NOT NULL ,
    target_id int NOT NULL ,
    weight float64 NOT NULL ,
    created_time int NOT NULL ,
);

/*Table: chat_sync2  */
CREATE TABLE IF NOT EXISTS sun_push.chat_sync2 (
    sync_id int PRIMARY KEY NOT NULL ,
    to_user_id int ,
    chat_sync_type_id int ,
    room_key string ,
    chat_key string ,
    from_high_message_id int ,
    to_low_message_id int ,
    message_id int ,
    message_pb bytes ,
    message_json string ,
    created_time int ,
);

/*Table: lower_table  */
CREATE TABLE IF NOT EXISTS sun_push.lower_table (
    id int PRIMARY KEY NOT NULL ,
    text string NOT NULL ,
    time_stamp int NOT NULL ,
    any_thing_more_ int NOT NULL ,
);

/*Table: push_chat  */
CREATE TABLE IF NOT EXISTS sun_push.push_chat (
    push_id int PRIMARY KEY NOT NULL ,
    to_user_id int NOT NULL ,
    push_type_id int NOT NULL ,
    room_key string NOT NULL ,
    chat_key string NOT NULL ,
    seq int NOT NULL ,
    unseen_count int NOT NULL ,
    from_high_message_id int NOT NULL ,
    to_low_message_id int NOT NULL ,
    message_id int NOT NULL ,
    message_file_id int NOT NULL ,
    message_pb bytes NOT NULL ,
    message_json string NOT NULL ,
    created_time int NOT NULL ,
);

/*Table: push_chat2  */
CREATE TABLE IF NOT EXISTS sun_push.push_chat2 (
    sync_id int PRIMARY KEY NOT NULL ,
    to_user_id int ,
    chat_sync_type_id int ,
    room_key string ,
    chat_key string ,
    from_high_message_id int ,
    to_low_message_id int ,
    message_id int ,
    message_pb bytes ,
    message_json string ,
    created_time int ,
);

/*Table: http_rpc_log  */
CREATE TABLE IF NOT EXISTS sun_log.http_rpc_log (
    id int PRIMARY KEY NOT NULL ,
    time string NOT NULL ,
    method_full string NOT NULL ,
    method_parent string NOT NULL ,
    user_id int NOT NULL ,
    session_id string NOT NULL ,
    status_code int NOT NULL ,
    input_size int NOT NULL ,
    output_size int NOT NULL ,
    reqest_json string NOT NULL ,
    response_json string NOT NULL ,
    reqest_param_json string NOT NULL ,
    response_msg_json string NOT NULL ,
);

/*Table: metric_log  */
CREATE TABLE IF NOT EXISTS sun_log.metric_log (
    id int PRIMARY KEY NOT NULL ,
    instance_id int NOT NULL ,
    start_from string NOT NULL ,
    end_to string NOT NULL ,
    start_time int NOT NULL ,
    duration string NOT NULL ,
    meterics_json string NOT NULL ,
);

/*Table: xfile_service_info_log  */
CREATE TABLE IF NOT EXISTS sun_log.xfile_service_info_log (
    id int PRIMARY KEY NOT NULL ,
    instance_id int NOT NULL ,
    url string NOT NULL ,
    created_time string NOT NULL ,
);

/*Table: xfile_service_metric_log  */
CREATE TABLE IF NOT EXISTS sun_log.xfile_service_metric_log (
    id int PRIMARY KEY NOT NULL ,
    instance_id int NOT NULL ,
    metric_json string NOT NULL ,
);

/*Table: xfile_service_request_log  */
CREATE TABLE IF NOT EXISTS sun_log.xfile_service_request_log (
    id int PRIMARY KEY NOT NULL ,
    local_seq int NOT NULL ,
    instance_id int NOT NULL ,
    url string NOT NULL ,
    http_code int NOT NULL ,
    created_time string NOT NULL ,
);

/*Table: accounts  */
CREATE TABLE IF NOT EXISTS suncdb.accounts (
    id INT NOT NULL ,
    balance DECIMAL ,
);

/*Table: post_cdb  */
CREATE TABLE IF NOT EXISTS suncdb.post_cdb (
    post_id INT NOT NULL ,
    user_id INT NOT NULL ,
    post_type_enum INT NOT NULL ,
    post_category_enum INT NOT NULL ,
    media_id INT NOT NULL ,
    post_key STRING NOT NULL ,
    text STRING NOT NULL ,
    rich_text STRING NOT NULL ,
    media_count INT NOT NULL ,
    shared_to INT NOT NULL ,
    disable_comment INT NOT NULL ,
    source INT NOT NULL ,
    has_tag INT NOT NULL ,
    seq INT NOT NULL ,
    comments_count INT NOT NULL ,
    likes_count INT NOT NULL ,
    views_count INT NOT NULL ,
    edited_time INT NOT NULL ,
    created_time INT NOT NULL ,
    re_shared_post_id INT NOT NULL ,
);

}