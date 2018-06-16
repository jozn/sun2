
/*Table: action  */
CREATE TABLE IF NOT EXISTS sun.action (
    action_id int PRIMARY KEY NOT NULL ,
    actor_user_id int NOT NULL ,
    action_type int NOT NULL ,
    peer_user_id int NOT NULL ,
    post_id int NOT NULL ,
    comment_id int NOT NULL ,
    murmur64_hash int NOT NULL ,
    created_time int NOT NULL ,
);

/*Table: blocked  */
CREATE TABLE IF NOT EXISTS sun.blocked (
    id int PRIMARY KEY NOT NULL ,
    user_id int NOT NULL ,
    blocked_user_id int NOT NULL ,
    created_time int NOT NULL ,
);

/*Table: comment  */
CREATE TABLE IF NOT EXISTS sun.comment (
    comment_id int PRIMARY KEY NOT NULL ,
    user_id int NOT NULL ,
    post_id int NOT NULL ,
    text string NOT NULL ,
    likes_count int NOT NULL ,
    is_edited int NOT NULL ,
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

/*Table: followed  */
CREATE TABLE IF NOT EXISTS sun.followed (
    id int PRIMARY KEY NOT NULL ,
    user_id int NOT NULL ,
    followed_user_id int NOT NULL ,
    created_time int NOT NULL ,
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
    notify_type int NOT NULL ,
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
    client_id int NOT NULL ,
    phone string NOT NULL ,
    first_name string NOT NULL ,
    last_name string NOT NULL ,
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

/*Table: post_count  */
CREATE TABLE IF NOT EXISTS sun.post_count (
    post_id int PRIMARY KEY NOT NULL ,
    comments_count int NOT NULL ,
    likes_count int NOT NULL ,
    views_count int NOT NULL ,
    re_shared_count int NOT NULL ,
    chat_shared_count int NOT NULL ,
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
    post_type int NOT NULL ,
    created_time int NOT NULL ,
);

/*Table: post_reshared  */
CREATE TABLE IF NOT EXISTS sun.post_reshared (
    reshared_id int PRIMARY KEY NOT NULL ,
    post_id int NOT NULL ,
    by_user_id int NOT NULL ,
    post_user_id int NOT NULL ,
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

/*Table: sms  */
CREATE TABLE IF NOT EXISTS sun.sms (
    id UNKNOWN_sqlToGo__ PRIMARY KEY NOT NULL ,
    hash string NOT NULL ,
    client_phone string NOT NULL ,
    genrated_code int NOT NULL ,
    sms_sender_number string NOT NULL ,
    sms_send_statues string NOT NULL ,
    carrier string NOT NULL ,
    country bytes NOT NULL ,
    is_valid_phone int NOT NULL ,
    is_confirmed int NOT NULL ,
    is_login int NOT NULL ,
    is_register int NOT NULL ,
    retried_count int NOT NULL ,
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
    is_verified int NOT NULL ,
    avatar_id int NOT NULL ,
    profile_privacy int NOT NULL ,
    online_privacy int NOT NULL ,
    call_privacy int NOT NULL ,
    add_to_group_privacy int NOT NULL ,
    seen_message_privacy int NOT NULL ,
    phone int NOT NULL ,
    email string UNIQUE NOT NULL ,
    about string NOT NULL ,
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
    last_post_time int NOT NULL ,
    created_time int NOT NULL ,
    version_time int NOT NULL ,
    is_deleted int NOT NULL ,
    is_banned int NOT NULL ,
);

/*Table: user_relation  */
CREATE TABLE IF NOT EXISTS sun.user_relation (
    rel_nano_id int PRIMARY KEY NOT NULL ,
    user_id int NOT NULL ,
    peer_user_id int NOT NULL ,
    follwing int NOT NULL ,
    followed int NOT NULL ,
    in_contacts int NOT NULL ,
    mutual_contact int NOT NULL ,
    is_favorite int NOT NULL ,
    notify int NOT NULL ,
);

/*Table: chat  */
CREATE TABLE IF NOT EXISTS sun_chat.chat (
    chat_id int PRIMARY KEY NOT NULL ,
    chat_key string NOT NULL ,
    room_key string NOT NULL ,
    room_type int NOT NULL ,
    user_id int NOT NULL ,
    peer_user_id int NOT NULL ,
    group_id int NOT NULL ,
    hash_tag_id int NOT NULL ,
    title string NOT NULL ,
    pin_time_ms int NOT NULL ,
    from_msg_id int NOT NULL ,
    unseen_count int NOT NULL ,
    seq int NOT NULL ,
    last_msg_id int NOT NULL ,
    last_my_msg_status int NOT NULL ,
    my_last_seen_seq int NOT NULL ,
    my_last_seen_msg_id int NOT NULL ,
    peer_last_seen_msg_id int NOT NULL ,
    my_last_delivered_seq int NOT NULL ,
    my_last_delivered_msg_id int NOT NULL ,
    peer_last_delivered_msg_id int NOT NULL ,
    is_active int NOT NULL ,
    is_left int NOT NULL ,
    is_creator int NOT NULL ,
    is_kicked int NOT NULL ,
    is_admin int NOT NULL ,
    is_deactivated int NOT NULL ,
    is_muted int NOT NULL ,
    mute_until int NOT NULL ,
    version_time_ms int NOT NULL ,
    version_seq int NOT NULL ,
    order_time int NOT NULL ,
    created_time int NOT NULL ,
    draft_text string NOT NULL ,
    drat_reply_to_msg_id int NOT NULL ,
);

/*Table: chat_deleted  */
CREATE TABLE IF NOT EXISTS sun_chat.chat_deleted (
    chat_id int PRIMARY KEY NOT NULL ,
    room_key string PRIMARY KEY NOT NULL ,
);

/*Table: chat_last_message  */
CREATE TABLE IF NOT EXISTS sun_chat.chat_last_message (
    chat_id_group_id string PRIMARY KEY NOT NULL ,
    last_msg_pb bytes NOT NULL ,
);

/*Table: chat_user_version  */
CREATE TABLE IF NOT EXISTS sun_chat.chat_user_version (
    user_id int NOT NULL ,
    chat_id int NOT NULL ,
    version_time_nano int PRIMARY KEY NOT NULL ,
);

/*Table: group  */
CREATE TABLE IF NOT EXISTS sun_chat.group (
    group_id int PRIMARY KEY NOT NULL ,
    group_key string NOT NULL ,
    group_name string NOT NULL ,
    user_name string NOT NULL ,
    is_super_group int NOT NULL ,
    hash_tag_id int NOT NULL ,
    creator_user_id int NOT NULL ,
    group_privacy int NOT NULL ,
    history_view_able int NOT NULL ,
    seq int NOT NULL ,
    last_msg_id int NOT NULL ,
    pined_msg_id int NOT NULL ,
    avatar_ref_id int NOT NULL ,
    avatar_count int NOT NULL ,
    about string NOT NULL ,
    invite_link_hash string NOT NULL ,
    members_count int NOT NULL ,
    admins_count int NOT NULL ,
    moderator_counts int NOT NULL ,
    sort_time int NOT NULL ,
    created_time int NOT NULL ,
    is_mute int NOT NULL ,
    is_deleted int NOT NULL ,
    is_banned int NOT NULL ,
);

/*Table: group_member  */
CREATE TABLE IF NOT EXISTS sun_chat.group_member (
    order_id int PRIMARY KEY NOT NULL ,
    group_id int NOT NULL ,
    user_id int NOT NULL ,
    by_user_id int NOT NULL ,
    group_role int NOT NULL ,
    created_time int NOT NULL ,
);

/*Table: group_orderd_user  */
CREATE TABLE IF NOT EXISTS sun_chat.group_orderd_user (
    order_id int PRIMARY KEY NOT NULL ,
    group_id int NOT NULL ,
    user_id int NOT NULL ,
);

/*Table: group_pined_msg  */
CREATE TABLE IF NOT EXISTS sun_chat.group_pined_msg (
    msg_id int PRIMARY KEY NOT NULL ,
    msg_pb bytes NOT NULL ,
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

/*Table: invalidate_cache  */
CREATE TABLE IF NOT EXISTS sun_internal.invalidate_cache (
    order_id int PRIMARY KEY NOT NULL ,
    cache_key string NOT NULL ,
);

}