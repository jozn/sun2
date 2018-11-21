# Sun2 - Backend  (A clean restart of old sun codebase with focus for modular and service ointed )

This is backend for [android client](https://github.com/jozn/ms_native)

+80% is completed.

A note for small numbers of commits: I started fresh from old "sun" project, because my technology stack has evolved a lot and i started fresh and common copied the codes into this project and changed them as necessary.  

# Technologies
Service oriented. Small services with ability to bundle multi service at once into one executable file.
Websocket, Http, REST (not used anymore), Cassandra, MySQL, Custom RPC, protocol buffers , a lot of code generators. Experimental CockroachDB.

# Features
+ Like what needed for Instagram backend. post, like, comments, notifications, hot posts, personal suggestions, logging.
+ Cassandra based file storage and file serving , with on fly converting and resizing of images then caching them on local disks, converting .gifs to .mp4 on fly and reducing a lot of its size. Highly scalable. Cached result in several discs.
+ Full Chat system with ability to push.

# Guides:

+ [cassandra_walker] output: https://github.com/jozn/sun2/tree/master/shared/xc
+ [db_walker](https://github.com/jozn/db-walker) and [pb_walker](https://github.com/jozn/pb_walker) output: https://github.com/jozn/sun2/tree/master/shared/x
+ Protocol buffer files: https://github.com/jozn/sun2/tree/master/shared/proto

# Object Storage
+ source: https://github.com/jozn/sun2/tree/master/servises/file_service
+ Based on Cassandra. High scalability. Similar to what Walmart has done. This tool can stores any files, every files splits into smaller chunks before inserting to Cassandra. A group of custom http servers, sits between users and Cassandra and cache the output on their local disks. This http servers resize and convert images based on their http query parameters and extensions. This servers also convert .gif and .webp to .mp4 when requested, this significantly reduce animated .gif sizes. This server is planned to stream .mp4 with HTTP Live Streaming (HLS).
