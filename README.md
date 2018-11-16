# Backend 

This is backend for [android client](https://github.com/jozn/ms_native)
+80% is completed.
# Technologies
Websocket, Http, REST (not used anymore), Cassandra, MySQL, Custom RPC, protocol buffers , a lot of code generators. Experimental CockroachDB.

# Features
+ Like what needed for Instagram backend. post, like, comments, notifications, hot posts, personal suggestions, logging.
+ Cassandra based file storage, with on fly converting and resizing of images then caches them, converting .gifs to .mp4 on fly and reducing a lot of its size. Highly scalable. Cached result in several discs.
+ Full Chat system with ability to push.

# Guides:

+ [cassandra_walker] output: https://github.com/jozn/sun2/tree/master/shared/xc
+ [db_walker](https://github.com/jozn/db-walker) and [pb_walker](https://github.com/jozn/pb_walker) output: https://github.com/jozn/sun2/tree/master/shared/x
+ Protocol buffer files: https://github.com/jozn/sun2/tree/master/shared/proto
