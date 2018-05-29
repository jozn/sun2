drop keyspace sunc_file;
create keyspace sunc_file WITH REPLICATION = { 'class' : 'SimpleStrategy', 'replication_factor' : 3 };
/*create keyspace sunc_file WITH REPLICATION = {'class' : 'NetworkTopologyStrategy', 'dc1' : 3, 'dc2' : 2};*/


create table sunc_file.file_storage (
	md5_key varchar primary key,
	user_id int,
	first_file_ref_id bigint,
	mime_type varchar,
	width int,
	height int,
	length int,
	extension varchar,
	zdata_thumb blob,
	zdata blob
); 

create table sunc_file.file_ref (
	file_ref_id bigint primary key,
	user_id int,
	name varchar,
	store_type varchar,
	mime_type varchar,
	width int,
	height int,
	duration int,
	extension varchar,
	md5_key varchar,
	storage_user_id int,
	first_file_ref_id bigint
); 


create table sunc_file.storage_usage (
	md5_key varchar,
	file_ref_id bigint,
	store_type varchar,
	primary key (md5_key, file_ref_id)
); 

create table sunc_file.storage_gc (
	shard varchar,
	md5_key varchar,
	file_ref_id bigint,
	store_type varchar,
	primary key (md5_key, file_ref_id)
); 

create table sunc_file.storage_shadow (
	shadow_md5_key varchar primary key,
	storage_md5_key varchar,
	mime_type varchar
);



/*
store type: "social" "msg_direct" "msg_group" "avatar"

storage_shadow: not used now
/*

/* deprecated */
/*create table sunc_file.file_msg (
	file_msg_id bigint primary key,
	user_id int,
	store_type varchar,
	mime_type varchar,
	width int,
	height int,
	length int,
	extension varchar,
	md5_key varchar,
	storage_user_id int
); 
*/