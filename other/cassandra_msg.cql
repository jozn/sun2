drop keyspace sunc_msg;
create keyspace sunc_msg WITH REPLICATION = { 'class' : 'SimpleStrategy', 'replication_factor' : 3 };

/* this to mysql */
create table sunc_msg.chat (
	chat_key varchar primary key,
	user_id int,
	room_type int,
	room_key varchar,
	peer_user_id int,
	group_id bigint,
	hashtag_id bigint,
	seq int,
	seen_seq int,
	is_mute int,
	created_time int,
	updated_ms bigint,
	last_msg blob,
	last_msg_status int,
); 

create table sunc_msg.message (
	chat_key varchar,
	message_id bigint,
	room_key varchar,
	user_id int,
	file_ref_id bigint,
	message_type int,
	text text,
	hiden int,
	seq int,
	fwd_msg_id bigint,
	post_id bigint,
	sticker_id bigint,
	created_time int,
	delivered_time int,
	seen_time int,
	deliviry_status int,
	reply_to_message_id bigint,
	views_count int,
	edit_time int,
	ttl int,
	extra_pb blob,
	primary key (chat_key, message_id)
); 

create table sunc_msg.message_media (
	chat_key varchar,
	message_cat varchar,
	message_id bigint,
	file_ref_id bigint,
	post_id bigint,
	primary key ((chat_key, message_cat), message_id)
); 
