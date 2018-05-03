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

// FollowingListById Generated from index 'PRIMARY' -- retrieves a row from 'sun.following_list' as a FollowingList.
func FollowingListById(db *sqlx.DB, id int) (*FollowingList, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM sun.following_list ` +
		`WHERE Id = ?`

	XOLog(sqlstr, id)
	fl := FollowingList{
		_exists: true,
	}

	err = db.Get(&fl, sqlstr, id)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnFollowingList_LoadOne(&fl)

	return &fl, nil
}

// FollowingListMemberById Generated from index 'PRIMARY' -- retrieves a row from 'sun.following_list_member' as a FollowingListMember.
func FollowingListMemberById(db *sqlx.DB, id int) (*FollowingListMember, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM sun.following_list_member ` +
		`WHERE Id = ?`

	XOLog(sqlstr, id)
	flm := FollowingListMember{
		_exists: true,
	}

	err = db.Get(&flm, sqlstr, id)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnFollowingListMember_LoadOne(&flm)

	return &flm, nil
}

// FollowingListMemberRemovedById Generated from index 'PRIMARY' -- retrieves a row from 'sun.following_list_member_removed' as a FollowingListMemberRemoved.
func FollowingListMemberRemovedById(db *sqlx.DB, id int) (*FollowingListMemberRemoved, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM sun.following_list_member_removed ` +
		`WHERE Id = ?`

	XOLog(sqlstr, id)
	flmr := FollowingListMemberRemoved{
		_exists: true,
	}

	err = db.Get(&flmr, sqlstr, id)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnFollowingListMemberRemoved_LoadOne(&flmr)

	return &flmr, nil
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

// SearchClickedById Generated from index 'PRIMARY' -- retrieves a row from 'sun.search_clicked' as a SearchClicked.
func SearchClickedById(db *sqlx.DB, id int) (*SearchClicked, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM sun.search_clicked ` +
		`WHERE Id = ?`

	XOLog(sqlstr, id)
	sc := SearchClicked{
		_exists: true,
	}

	err = db.Get(&sc, sqlstr, id)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnSearchClicked_LoadOne(&sc)

	return &sc, nil
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

// SessionCopyBySessionUuid Generated from index 'PRIMARY' -- retrieves a row from 'sun.session_copy' as a SessionCopy.
func SessionCopyBySessionUuid(db *sqlx.DB, sessionUuid string) (*SessionCopy, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM sun.session_copy ` +
		`WHERE SessionUuid = ?`

	XOLog(sqlstr, sessionUuid)
	sc := SessionCopy{
		_exists: true,
	}

	err = db.Get(&sc, sqlstr, sessionUuid)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnSessionCopy_LoadOne(&sc)

	return &sc, nil
}

// SettingClientByUserId Generated from index 'PRIMARY' -- retrieves a row from 'sun.setting_client' as a SettingClient.
func SettingClientByUserId(db *sqlx.DB, userId int) (*SettingClient, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM sun.setting_client ` +
		`WHERE UserId = ?`

	XOLog(sqlstr, userId)
	sc := SettingClient{
		_exists: true,
	}

	err = db.Get(&sc, sqlstr, userId)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnSettingClient_LoadOne(&sc)

	return &sc, nil
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

// UserMetaInfoById Generated from index 'PRIMARY' -- retrieves a row from 'sun.user_meta_info' as a UserMetaInfo.
func UserMetaInfoById(db *sqlx.DB, id int) (*UserMetaInfo, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM sun.user_meta_info ` +
		`WHERE Id = ?`

	XOLog(sqlstr, id)
	umi := UserMetaInfo{
		_exists: true,
	}

	err = db.Get(&umi, sqlstr, id)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnUserMetaInfo_LoadOne(&umi)

	return &umi, nil
}

// UserPasswordByUserId Generated from index 'PRIMARY' -- retrieves a row from 'sun.user_password' as a UserPassword.
func UserPasswordByUserId(db *sqlx.DB, userId int) (*UserPassword, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM sun.user_password ` +
		`WHERE UserId = ?`

	XOLog(sqlstr, userId)
	up := UserPassword{
		_exists: true,
	}

	err = db.Get(&up, sqlstr, userId)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnUserPassword_LoadOne(&up)

	return &up, nil
}

// ChatByChatKey Generated from index 'PRIMARY' -- retrieves a row from 'sun_chat.chat' as a Chat.
func ChatByChatKey(db *sqlx.DB, chatKey string) (*Chat, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM sun_chat.chat ` +
		`WHERE ChatKey = ?`

	XOLog(sqlstr, chatKey)
	c := Chat{
		_exists: true,
	}

	err = db.Get(&c, sqlstr, chatKey)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnChat_LoadOne(&c)

	return &c, nil
}

// ChatLastMessageByChatKey Generated from index 'PRIMARY' -- retrieves a row from 'sun_chat.chat_last_message' as a ChatLastMessage.
func ChatLastMessageByChatKey(db *sqlx.DB, chatKey string) (*ChatLastMessage, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM sun_chat.chat_last_message ` +
		`WHERE ChatKey = ?`

	XOLog(sqlstr, chatKey)
	clm := ChatLastMessage{
		_exists: true,
	}

	err = db.Get(&clm, sqlstr, chatKey)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnChatLastMessage_LoadOne(&clm)

	return &clm, nil
}

// DirectMessageByMessageId Generated from index 'PRIMARY' -- retrieves a row from 'sun_chat.direct_message' as a DirectMessage.
func DirectMessageByMessageId(db *sqlx.DB, messageId int) (*DirectMessage, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM sun_chat.direct_message ` +
		`WHERE MessageId = ?`

	XOLog(sqlstr, messageId)
	dm := DirectMessage{
		_exists: true,
	}

	err = db.Get(&dm, sqlstr, messageId)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnDirectMessage_LoadOne(&dm)

	return &dm, nil
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

// GroupMemberById Generated from index 'PRIMARY' -- retrieves a row from 'sun_chat.group_member' as a GroupMember.
func GroupMemberById(db *sqlx.DB, id int) (*GroupMember, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM sun_chat.group_member ` +
		`WHERE Id = ?`

	XOLog(sqlstr, id)
	gm := GroupMember{
		_exists: true,
	}

	err = db.Get(&gm, sqlstr, id)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnGroupMember_LoadOne(&gm)

	return &gm, nil
}

// GroupMessageByMessageId Generated from index 'PRIMARY' -- retrieves a row from 'sun_chat.group_message' as a GroupMessage.
func GroupMessageByMessageId(db *sqlx.DB, messageId int) (*GroupMessage, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM sun_chat.group_message ` +
		`WHERE MessageId = ?`

	XOLog(sqlstr, messageId)
	gm := GroupMessage{
		_exists: true,
	}

	err = db.Get(&gm, sqlstr, messageId)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnGroupMessage_LoadOne(&gm)

	return &gm, nil
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

// ChatSync2BySyncId Generated from index 'PRIMARY' -- retrieves a row from 'sun_push.chat_sync2' as a ChatSync2.
func ChatSync2BySyncId(db *sqlx.DB, syncId int) (*ChatSync2, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM sun_push.chat_sync2 ` +
		`WHERE SyncId = ?`

	XOLog(sqlstr, syncId)
	cs := ChatSync2{
		_exists: true,
	}

	err = db.Get(&cs, sqlstr, syncId)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnChatSync2_LoadOne(&cs)

	return &cs, nil
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

// PushChat2BySyncId Generated from index 'PRIMARY' -- retrieves a row from 'sun_push.push_chat2' as a PushChat2.
func PushChat2BySyncId(db *sqlx.DB, syncId int) (*PushChat2, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM sun_push.push_chat2 ` +
		`WHERE SyncId = ?`

	XOLog(sqlstr, syncId)
	pc := PushChat2{
		_exists: true,
	}

	err = db.Get(&pc, sqlstr, syncId)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnPushChat2_LoadOne(&pc)

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
