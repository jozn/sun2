package x

import "github.com/jmoiron/sqlx"

//todo this code can be used for multi secondery coulmn -- but this one interfier with secondery_template - in futer merege this two temple to one unifed that uses both primiry keys and othe seconder options and with reloadings

// ActionByActionId Generated from index 'PRIMARY' -- retrieves a row from 'sun.action' as a Action.
func ActionByActionId(db *sqlx.DB, actionId int) (*Action, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM sun.action ` +
		`WHERE ActionId = ?`

	XOLog(sqlstr, actionId)
	a := Action{
		_exists: true,
	}

	err = db.Get(&a, sqlstr, actionId)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnAction_LoadOne(&a)

	return &a, nil
}

// BlockedById Generated from index 'PRIMARY' -- retrieves a row from 'sun.blocked' as a Blocked.
func BlockedById(db *sqlx.DB, id int) (*Blocked, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM sun.blocked ` +
		`WHERE Id = ?`

	XOLog(sqlstr, id)
	b := Blocked{
		_exists: true,
	}

	err = db.Get(&b, sqlstr, id)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnBlocked_LoadOne(&b)

	return &b, nil
}

// CommentByCommentId Generated from index 'PRIMARY' -- retrieves a row from 'sun.comment' as a Comment.
func CommentByCommentId(db *sqlx.DB, commentId int) (*Comment, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM sun.comment ` +
		`WHERE CommentId = ?`

	XOLog(sqlstr, commentId)
	c := Comment{
		_exists: true,
	}

	err = db.Get(&c, sqlstr, commentId)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnComment_LoadOne(&c)

	return &c, nil
}

// CommentDeletedByCommentId Generated from index 'PRIMARY' -- retrieves a row from 'sun.comment_deleted' as a CommentDeleted.
func CommentDeletedByCommentId(db *sqlx.DB, commentId int) (*CommentDeleted, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM sun.comment_deleted ` +
		`WHERE CommentId = ?`

	XOLog(sqlstr, commentId)
	cd := CommentDeleted{
		_exists: true,
	}

	err = db.Get(&cd, sqlstr, commentId)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnCommentDeleted_LoadOne(&cd)

	return &cd, nil
}

// EventByEventId Generated from index 'PRIMARY' -- retrieves a row from 'sun.event' as a Event.
func EventByEventId(db *sqlx.DB, eventId int) (*Event, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM sun.event ` +
		`WHERE EventId = ?`

	XOLog(sqlstr, eventId)
	e := Event{
		_exists: true,
	}

	err = db.Get(&e, sqlstr, eventId)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnEvent_LoadOne(&e)

	return &e, nil
}

// FollowedById Generated from index 'PRIMARY' -- retrieves a row from 'sun.followed' as a Followed.
func FollowedById(db *sqlx.DB, id int) (*Followed, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM sun.followed ` +
		`WHERE Id = ?`

	XOLog(sqlstr, id)
	f := Followed{
		_exists: true,
	}

	err = db.Get(&f, sqlstr, id)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnFollowed_LoadOne(&f)

	return &f, nil
}

// LikeById Generated from index 'PRIMARY' -- retrieves a row from 'sun.likes' as a Like.
func LikeById(db *sqlx.DB, id int) (*Like, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM sun.likes ` +
		`WHERE Id = ?`

	XOLog(sqlstr, id)
	l := Like{
		_exists: true,
	}

	err = db.Get(&l, sqlstr, id)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnLike_LoadOne(&l)

	return &l, nil
}

// NotifyByNotifyId Generated from index 'PRIMARY' -- retrieves a row from 'sun.notify' as a Notify.
func NotifyByNotifyId(db *sqlx.DB, notifyId int) (*Notify, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM sun.notify ` +
		`WHERE NotifyId = ?`

	XOLog(sqlstr, notifyId)
	n := Notify{
		_exists: true,
	}

	err = db.Get(&n, sqlstr, notifyId)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnNotify_LoadOne(&n)

	return &n, nil
}

// NotifyRemovedById Generated from index 'PRIMARY' -- retrieves a row from 'sun.notify_removed' as a NotifyRemoved.
func NotifyRemovedById(db *sqlx.DB, id int) (*NotifyRemoved, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM sun.notify_removed ` +
		`WHERE Id = ?`

	XOLog(sqlstr, id)
	nr := NotifyRemoved{
		_exists: true,
	}

	err = db.Get(&nr, sqlstr, id)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnNotifyRemoved_LoadOne(&nr)

	return &nr, nil
}

// PhoneContactById Generated from index 'PRIMARY' -- retrieves a row from 'sun.phone_contacts' as a PhoneContact.
func PhoneContactById(db *sqlx.DB, id int) (*PhoneContact, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM sun.phone_contacts ` +
		`WHERE Id = ?`

	XOLog(sqlstr, id)
	pc := PhoneContact{
		_exists: true,
	}

	err = db.Get(&pc, sqlstr, id)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnPhoneContact_LoadOne(&pc)

	return &pc, nil
}

// PostByPostId Generated from index 'PRIMARY' -- retrieves a row from 'sun.post' as a Post.
func PostByPostId(db *sqlx.DB, postId int) (*Post, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM sun.post ` +
		`WHERE PostId = ?`

	XOLog(sqlstr, postId)
	p := Post{
		_exists: true,
	}

	err = db.Get(&p, sqlstr, postId)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnPost_LoadOne(&p)

	return &p, nil
}

// PostCountByPostId Generated from index 'PRIMARY' -- retrieves a row from 'sun.post_count' as a PostCount.
func PostCountByPostId(db *sqlx.DB, postId int) (*PostCount, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM sun.post_count ` +
		`WHERE PostId = ?`

	XOLog(sqlstr, postId)
	pc := PostCount{
		_exists: true,
	}

	err = db.Get(&pc, sqlstr, postId)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnPostCount_LoadOne(&pc)

	return &pc, nil
}

// PostDeletedByPostId Generated from index 'PRIMARY' -- retrieves a row from 'sun.post_deleted' as a PostDeleted.
func PostDeletedByPostId(db *sqlx.DB, postId int) (*PostDeleted, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM sun.post_deleted ` +
		`WHERE PostId = ?`

	XOLog(sqlstr, postId)
	pd := PostDeleted{
		_exists: true,
	}

	err = db.Get(&pd, sqlstr, postId)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnPostDeleted_LoadOne(&pd)

	return &pd, nil
}

// PostKeyById Generated from index 'PRIMARY' -- retrieves a row from 'sun.post_keys' as a PostKey.
func PostKeyById(db *sqlx.DB, id int) (*PostKey, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM sun.post_keys ` +
		`WHERE Id = ?`

	XOLog(sqlstr, id)
	pk := PostKey{
		_exists: true,
	}

	err = db.Get(&pk, sqlstr, id)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnPostKey_LoadOne(&pk)

	return &pk, nil
}

// PostLinkByLinkId Generated from index 'PRIMARY' -- retrieves a row from 'sun.post_link' as a PostLink.
func PostLinkByLinkId(db *sqlx.DB, linkId int) (*PostLink, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM sun.post_link ` +
		`WHERE LinkId = ?`

	XOLog(sqlstr, linkId)
	pl := PostLink{
		_exists: true,
	}

	err = db.Get(&pl, sqlstr, linkId)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnPostLink_LoadOne(&pl)

	return &pl, nil
}

// PostMediaByMediaId Generated from index 'PRIMARY' -- retrieves a row from 'sun.post_media' as a PostMedia.
func PostMediaByMediaId(db *sqlx.DB, mediaId int) (*PostMedia, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM sun.post_media ` +
		`WHERE MediaId = ?`

	XOLog(sqlstr, mediaId)
	pm := PostMedia{
		_exists: true,
	}

	err = db.Get(&pm, sqlstr, mediaId)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnPostMedia_LoadOne(&pm)

	return &pm, nil
}

// PostMentionedByMentionedId Generated from index 'PRIMARY' -- retrieves a row from 'sun.post_mentioned' as a PostMentioned.
func PostMentionedByMentionedId(db *sqlx.DB, mentionedId int) (*PostMentioned, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM sun.post_mentioned ` +
		`WHERE MentionedId = ?`

	XOLog(sqlstr, mentionedId)
	pm := PostMentioned{
		_exists: true,
	}

	err = db.Get(&pm, sqlstr, mentionedId)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnPostMentioned_LoadOne(&pm)

	return &pm, nil
}

// PostResharedByResharedId Generated from index 'PRIMARY' -- retrieves a row from 'sun.post_reshared' as a PostReshared.
func PostResharedByResharedId(db *sqlx.DB, resharedId int) (*PostReshared, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM sun.post_reshared ` +
		`WHERE ResharedId = ?`

	XOLog(sqlstr, resharedId)
	pr := PostReshared{
		_exists: true,
	}

	err = db.Get(&pr, sqlstr, resharedId)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnPostReshared_LoadOne(&pr)

	return &pr, nil
}

// SessionById Generated from index 'PRIMARY' -- retrieves a row from 'sun.session' as a Session.
func SessionById(db *sqlx.DB, id int) (*Session, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM sun.session ` +
		`WHERE Id = ?`

	XOLog(sqlstr, id)
	s := Session{
		_exists: true,
	}

	err = db.Get(&s, sqlstr, id)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnSession_LoadOne(&s)

	return &s, nil
}

// SettingNotificationByUserId Generated from index 'PRIMARY' -- retrieves a row from 'sun.setting_notifications' as a SettingNotification.
func SettingNotificationByUserId(db *sqlx.DB, userId int) (*SettingNotification, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM sun.setting_notifications ` +
		`WHERE UserId = ?`

	XOLog(sqlstr, userId)
	sn := SettingNotification{
		_exists: true,
	}

	err = db.Get(&sn, sqlstr, userId)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnSettingNotification_LoadOne(&sn)

	return &sn, nil
}

// SmById Generated from index 'PRIMARY' -- retrieves a row from 'sun.sms' as a Sm.
func SmById(db *sqlx.DB, id uint) (*Sm, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM sun.sms ` +
		`WHERE Id = ?`

	XOLog(sqlstr, id)
	s := Sm{
		_exists: true,
	}

	err = db.Get(&s, sqlstr, id)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnSm_LoadOne(&s)

	return &s, nil
}

// TagByTagId Generated from index 'PRIMARY' -- retrieves a row from 'sun.tag' as a Tag.
func TagByTagId(db *sqlx.DB, tagId int) (*Tag, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM sun.tag ` +
		`WHERE TagId = ?`

	XOLog(sqlstr, tagId)
	t := Tag{
		_exists: true,
	}

	err = db.Get(&t, sqlstr, tagId)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnTag_LoadOne(&t)

	return &t, nil
}

// TagPostById Generated from index 'PRIMARY' -- retrieves a row from 'sun.tag_post' as a TagPost.
func TagPostById(db *sqlx.DB, id int) (*TagPost, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM sun.tag_post ` +
		`WHERE Id = ?`

	XOLog(sqlstr, id)
	tp := TagPost{
		_exists: true,
	}

	err = db.Get(&tp, sqlstr, id)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnTagPost_LoadOne(&tp)

	return &tp, nil
}

// TriggerLogById Generated from index 'PRIMARY' -- retrieves a row from 'sun.trigger_log' as a TriggerLog.
func TriggerLogById(db *sqlx.DB, id int) (*TriggerLog, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM sun.trigger_log ` +
		`WHERE Id = ?`

	XOLog(sqlstr, id)
	tl := TriggerLog{
		_exists: true,
	}

	err = db.Get(&tl, sqlstr, id)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnTriggerLog_LoadOne(&tl)

	return &tl, nil
}

// UserByUserId Generated from index 'PRIMARY' -- retrieves a row from 'sun.user' as a User.
func UserByUserId(db *sqlx.DB, userId int) (*User, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM sun.user ` +
		`WHERE UserId = ?`

	XOLog(sqlstr, userId)
	u := User{
		_exists: true,
	}

	err = db.Get(&u, sqlstr, userId)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnUser_LoadOne(&u)

	return &u, nil
}

// UserRelationByRelNanoId Generated from index 'PRIMARY' -- retrieves a row from 'sun.user_relation' as a UserRelation.
func UserRelationByRelNanoId(db *sqlx.DB, relNanoId int) (*UserRelation, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM sun.user_relation ` +
		`WHERE RelNanoId = ?`

	XOLog(sqlstr, relNanoId)
	ur := UserRelation{
		_exists: true,
	}

	err = db.Get(&ur, sqlstr, relNanoId)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnUserRelation_LoadOne(&ur)

	return &ur, nil
}

// ChatByChatId Generated from index 'PRIMARY' -- retrieves a row from 'sun_chat.chat' as a Chat.
func ChatByChatId(db *sqlx.DB, chatId int) (*Chat, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM sun_chat.chat ` +
		`WHERE ChatId = ?`

	XOLog(sqlstr, chatId)
	c := Chat{
		_exists: true,
	}

	err = db.Get(&c, sqlstr, chatId)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnChat_LoadOne(&c)

	return &c, nil
}

// ChatDeletedByChatIdAndRoomKey Generated from index 'PRIMARY' -- retrieves a row from 'sun_chat.chat_deleted' as a ChatDeleted.
func ChatDeletedByChatIdAndRoomKey(db *sqlx.DB, chatId int, roomKey string) (*ChatDeleted, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM sun_chat.chat_deleted ` +
		`WHERE ChatId = ? AND RoomKey = ?`

	XOLog(sqlstr, chatId, roomKey)
	cd := ChatDeleted{
		_exists: true,
	}

	err = db.Get(&cd, sqlstr, chatId, roomKey)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnChatDeleted_LoadOne(&cd)

	return &cd, nil
}

// ChatLastMessageByChatIdGroupId Generated from index 'PRIMARY' -- retrieves a row from 'sun_chat.chat_last_message' as a ChatLastMessage.
func ChatLastMessageByChatIdGroupId(db *sqlx.DB, chatIdGroupId string) (*ChatLastMessage, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM sun_chat.chat_last_message ` +
		`WHERE ChatIdGroupId = ?`

	XOLog(sqlstr, chatIdGroupId)
	clm := ChatLastMessage{
		_exists: true,
	}

	err = db.Get(&clm, sqlstr, chatIdGroupId)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnChatLastMessage_LoadOne(&clm)

	return &clm, nil
}

// ChatUserVersionByVersionTimeNano Generated from index 'PRIMARY' -- retrieves a row from 'sun_chat.chat_user_version' as a ChatUserVersion.
func ChatUserVersionByVersionTimeNano(db *sqlx.DB, versionTimeNano int) (*ChatUserVersion, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM sun_chat.chat_user_version ` +
		`WHERE VersionTimeNano = ?`

	XOLog(sqlstr, versionTimeNano)
	cuv := ChatUserVersion{
		_exists: true,
	}

	err = db.Get(&cuv, sqlstr, versionTimeNano)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnChatUserVersion_LoadOne(&cuv)

	return &cuv, nil
}

// GroupByGroupId Generated from index 'PRIMARY' -- retrieves a row from 'sun_chat.group' as a Group.
func GroupByGroupId(db *sqlx.DB, groupId int) (*Group, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM sun_chat.group ` +
		`WHERE GroupId = ?`

	XOLog(sqlstr, groupId)
	g := Group{
		_exists: true,
	}

	err = db.Get(&g, sqlstr, groupId)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnGroup_LoadOne(&g)

	return &g, nil
}

// GroupMemberByOrderId Generated from index 'PRIMARY' -- retrieves a row from 'sun_chat.group_member' as a GroupMember.
func GroupMemberByOrderId(db *sqlx.DB, orderId int) (*GroupMember, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM sun_chat.group_member ` +
		`WHERE OrderId = ?`

	XOLog(sqlstr, orderId)
	gm := GroupMember{
		_exists: true,
	}

	err = db.Get(&gm, sqlstr, orderId)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnGroupMember_LoadOne(&gm)

	return &gm, nil
}

// GroupOrderdUserByOrderId Generated from index 'PRIMARY' -- retrieves a row from 'sun_chat.group_orderd_user' as a GroupOrderdUser.
func GroupOrderdUserByOrderId(db *sqlx.DB, orderId int) (*GroupOrderdUser, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM sun_chat.group_orderd_user ` +
		`WHERE OrderId = ?`

	XOLog(sqlstr, orderId)
	gou := GroupOrderdUser{
		_exists: true,
	}

	err = db.Get(&gou, sqlstr, orderId)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnGroupOrderdUser_LoadOne(&gou)

	return &gou, nil
}

// GroupPinedMsgByMsgId Generated from index 'PRIMARY' -- retrieves a row from 'sun_chat.group_pined_msg' as a GroupPinedMsg.
func GroupPinedMsgByMsgId(db *sqlx.DB, msgId int) (*GroupPinedMsg, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM sun_chat.group_pined_msg ` +
		`WHERE MsgId = ?`

	XOLog(sqlstr, msgId)
	gpm := GroupPinedMsg{
		_exists: true,
	}

	err = db.Get(&gpm, sqlstr, msgId)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnGroupPinedMsg_LoadOne(&gpm)

	return &gpm, nil
}

// FileMsgById Generated from index 'PRIMARY' -- retrieves a row from 'sun_file.file_msg' as a FileMsg.
func FileMsgById(db *sqlx.DB, id int) (*FileMsg, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM sun_file.file_msg ` +
		`WHERE Id = ?`

	XOLog(sqlstr, id)
	fm := FileMsg{
		_exists: true,
	}

	err = db.Get(&fm, sqlstr, id)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnFileMsg_LoadOne(&fm)

	return &fm, nil
}

// FilePostById Generated from index 'PRIMARY' -- retrieves a row from 'sun_file.file_post' as a FilePost.
func FilePostById(db *sqlx.DB, id int) (*FilePost, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM sun_file.file_post ` +
		`WHERE Id = ?`

	XOLog(sqlstr, id)
	fp := FilePost{
		_exists: true,
	}

	err = db.Get(&fp, sqlstr, id)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnFilePost_LoadOne(&fp)

	return &fp, nil
}

// ActionFanoutByOrderId Generated from index 'PRIMARY' -- retrieves a row from 'sun_meta.action_fanout' as a ActionFanout.
func ActionFanoutByOrderId(db *sqlx.DB, orderId int) (*ActionFanout, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM sun_meta.action_fanout ` +
		`WHERE OrderId = ?`

	XOLog(sqlstr, orderId)
	af := ActionFanout{
		_exists: true,
	}

	err = db.Get(&af, sqlstr, orderId)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnActionFanout_LoadOne(&af)

	return &af, nil
}

// HomeFanoutByOrderId Generated from index 'PRIMARY' -- retrieves a row from 'sun_meta.home_fanout' as a HomeFanout.
func HomeFanoutByOrderId(db *sqlx.DB, orderId int) (*HomeFanout, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM sun_meta.home_fanout ` +
		`WHERE OrderId = ?`

	XOLog(sqlstr, orderId)
	hf := HomeFanout{
		_exists: true,
	}

	err = db.Get(&hf, sqlstr, orderId)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnHomeFanout_LoadOne(&hf)

	return &hf, nil
}

// SuggestedTopPostById Generated from index 'PRIMARY' -- retrieves a row from 'sun_meta.suggested_top_posts' as a SuggestedTopPost.
func SuggestedTopPostById(db *sqlx.DB, id int) (*SuggestedTopPost, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM sun_meta.suggested_top_posts ` +
		`WHERE Id = ?`

	XOLog(sqlstr, id)
	stp := SuggestedTopPost{
		_exists: true,
	}

	err = db.Get(&stp, sqlstr, id)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnSuggestedTopPost_LoadOne(&stp)

	return &stp, nil
}

// SuggestedUserById Generated from index 'PRIMARY' -- retrieves a row from 'sun_meta.suggested_user' as a SuggestedUser.
func SuggestedUserById(db *sqlx.DB, id int) (*SuggestedUser, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM sun_meta.suggested_user ` +
		`WHERE Id = ?`

	XOLog(sqlstr, id)
	su := SuggestedUser{
		_exists: true,
	}

	err = db.Get(&su, sqlstr, id)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnSuggestedUser_LoadOne(&su)

	return &su, nil
}

// PushChatByPushId Generated from index 'PRIMARY' -- retrieves a row from 'sun_push.push_chat' as a PushChat.
func PushChatByPushId(db *sqlx.DB, pushId int) (*PushChat, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM sun_push.push_chat ` +
		`WHERE PushId = ?`

	XOLog(sqlstr, pushId)
	pc := PushChat{
		_exists: true,
	}

	err = db.Get(&pc, sqlstr, pushId)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnPushChat_LoadOne(&pc)

	return &pc, nil
}

// HTTPRPCLogById Generated from index 'PRIMARY' -- retrieves a row from 'sun_log.http_rpc_log' as a HTTPRPCLog.
func HTTPRPCLogById(db *sqlx.DB, id int) (*HTTPRPCLog, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM sun_log.http_rpc_log ` +
		`WHERE Id = ?`

	XOLog(sqlstr, id)
	hrl := HTTPRPCLog{
		_exists: true,
	}

	err = db.Get(&hrl, sqlstr, id)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnHTTPRPCLog_LoadOne(&hrl)

	return &hrl, nil
}

// MetricLogById Generated from index 'PRIMARY' -- retrieves a row from 'sun_log.metric_log' as a MetricLog.
func MetricLogById(db *sqlx.DB, id int) (*MetricLog, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM sun_log.metric_log ` +
		`WHERE Id = ?`

	XOLog(sqlstr, id)
	ml := MetricLog{
		_exists: true,
	}

	err = db.Get(&ml, sqlstr, id)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnMetricLog_LoadOne(&ml)

	return &ml, nil
}

// XfileServiceInfoLogById Generated from index 'PRIMARY' -- retrieves a row from 'sun_log.xfile_service_info_log' as a XfileServiceInfoLog.
func XfileServiceInfoLogById(db *sqlx.DB, id int) (*XfileServiceInfoLog, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM sun_log.xfile_service_info_log ` +
		`WHERE Id = ?`

	XOLog(sqlstr, id)
	xsil := XfileServiceInfoLog{
		_exists: true,
	}

	err = db.Get(&xsil, sqlstr, id)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnXfileServiceInfoLog_LoadOne(&xsil)

	return &xsil, nil
}

// XfileServiceMetricLogById Generated from index 'PRIMARY' -- retrieves a row from 'sun_log.xfile_service_metric_log' as a XfileServiceMetricLog.
func XfileServiceMetricLogById(db *sqlx.DB, id int) (*XfileServiceMetricLog, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM sun_log.xfile_service_metric_log ` +
		`WHERE Id = ?`

	XOLog(sqlstr, id)
	xsml := XfileServiceMetricLog{
		_exists: true,
	}

	err = db.Get(&xsml, sqlstr, id)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnXfileServiceMetricLog_LoadOne(&xsml)

	return &xsml, nil
}

// XfileServiceRequestLogById Generated from index 'PRIMARY' -- retrieves a row from 'sun_log.xfile_service_request_log' as a XfileServiceRequestLog.
func XfileServiceRequestLogById(db *sqlx.DB, id int) (*XfileServiceRequestLog, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM sun_log.xfile_service_request_log ` +
		`WHERE Id = ?`

	XOLog(sqlstr, id)
	xsrl := XfileServiceRequestLog{
		_exists: true,
	}

	err = db.Get(&xsrl, sqlstr, id)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnXfileServiceRequestLog_LoadOne(&xsrl)

	return &xsrl, nil
}

// InvalidateCacheByOrderId Generated from index 'PRIMARY' -- retrieves a row from 'sun_internal.invalidate_cache' as a InvalidateCache.
func InvalidateCacheByOrderId(db *sqlx.DB, orderId int) (*InvalidateCache, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM sun_internal.invalidate_cache ` +
		`WHERE OrderId = ?`

	XOLog(sqlstr, orderId)
	ic := InvalidateCache{
		_exists: true,
	}

	err = db.Get(&ic, sqlstr, orderId)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnInvalidateCache_LoadOne(&ic)

	return &ic, nil
}
