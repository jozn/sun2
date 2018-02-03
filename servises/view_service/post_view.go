package view_service

import (
	"ms/sun/base"
	"ms/sun2/servises/memcache_service"
	"ms/sun2/shared/x"
)

func PostsViews(posts []*x.Post, UserId int) (viw []*x.PB_PostView) {
	// UserSlice
	var photoIds []int
	for _, p := range posts {
		if x.PostTypeEnum(p.PostTypeEnum) == x.PostTypeEnum_POST_PHOTO {
			photoIds = append(photoIds, p.MediaId)
		}
	}
	if len(photoIds) > 0 {
		x.Store.PreLoadMediaByMediaIds(photoIds)
	}

	for _, p := range posts {
		viw = append(viw, PostSingleView(p, UserId))
	}
	return viw
}

//todo improve this to just look at cashe via preloading
func PostsViewsForPostIds(postsIds []int, MeId int) (viw []*x.PB_PostView) {
	if len(postsIds) != 0 {
		posts, err := x.NewPost_Selector().PostId_In(postsIds).GetRows(base.DB)
		if err == nil {
			viw = PostsViews(posts, MeId)
		}
	}
	return
}

func PostSingleViewForPostId(postsId int, MeId int) (viw *x.PB_PostView, ok bool) {
	if postsId < 1 {
		return
	}
	p, ok := x.Store.GetPostByPostId(postsId)
	if !ok {
		return
	}

	viw = PostSingleView(p, MeId)
	if viw != nil {
		ok = true
		return
	}
	ok = false
	return
}

func PostSingleView(post *x.Post, MeId int) *x.PB_PostView {
	if post == nil {
		return nil
	}
	v := &x.PB_PostView{
		PostId:         int64(post.PostId),
		UserId:         int32(post.UserId),
		Text:           post.Text,
		RichText:       post.RichText,
		MediaCount:     int32(post.MediaCount),
		SharedTo:       int32(post.SharedTo),
		DisableComment: int32(post.DisableComment),
		HasTag:         int32(post.HasTag),
		CommentsCount:  int32(post.CommentsCount),
		LikesCount:     int32(post.LikesCount),
		ViewsCount:     int32(post.ViewsCount),
		EditedTime:     int32(post.EditedTime),
		CreatedTime:    int32(post.CreatedTime),
		ReSharedPostId: int64(post.ReSharedPostId),
		DidILiked:      false,
		DidIReShared:   false,
	}

	if MeId > 0 {
		v.DidILiked = memcache_service.DidUserLikedPost(MeId, post.PostId)
	}

	if x.MediaTypeEnum(post.PostTypeEnum) == x.MediaTypeEnum_MEDIA_IMAGE {
		m, ok := x.Store.GetMediaByMediaId(post.PostId)
		if ok {
			v.MediaView = &x.PB_MediaView{
				MediaId:     int64(m.MediaId),
				UserId:      int32(m.UserId),
				PostId:      int32(m.PostId),
				AlbumId:     int32(m.AlbumId),
				Width:       int32(m.Width),
				Height:      int32(m.Height),
				Size:        int32(m.Size),
				Duration:    int32(m.Duration),
				Color:       m.Color,
				CreatedTime: int32(m.CreatedTime),
			}
		}
	}
	v.SenderUserView = UserViewAndMe(post.UserId, MeId)
	return v
}
