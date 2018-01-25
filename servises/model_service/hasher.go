package model_service

import (
	"fmt"
	"github.com/spaolacci/murmur3"
)

func hashPostCommented(PostId, CommentId int) int {
	key := fmt.Sprintf("post_commented_%d_%d", PostId, CommentId)
	return int(murmur3.Sum64([]byte(key)))
}

func hashFollowed(UserId, FollowedPeerUserId int) int {
	key := fmt.Sprintf("user_follwed_%d_%d", UserId, FollowedPeerUserId)
	return int(murmur3.Sum64([]byte(key)))
}

func hashPostLiked(UserId, PostId int) int {
	key := fmt.Sprintf("post_liked_%d_%d", UserId, PostId)
	return int(murmur3.Sum64([]byte(key)))
}
