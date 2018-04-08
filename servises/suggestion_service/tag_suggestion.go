package suggestion_service

import (
	"ms/sun_old/base"
	"ms/sun_old/helper"
	"ms/sun/servises/view_service"
	"ms/sun/shared/config"
	"ms/sun/shared/x"
	"time"
)

const TOP_TAGS_LIMIT = 30

func Tags_RepeatedlyJobs() {
	//top tags
	go func() {
		defer helper.JustRecover()
		for {
			if config.DEBUG_DELAY_RUN_STARTUPS { //just don't make the log files messy for this at each startups
				time.Sleep(time.Minute * 5)
			}
			reloadTopTags()
			reloadTopPostsForTopTags()
			time.Sleep(time.Minute * config.TAGS_RELOAD_TOP_INTERVAL_MINS)
		}
	}()
}

var TopTagsViews []*x.PB_TagView
var TopTagsWithPostsResult = make([]*x.PB_TopTagWithSamplePosts, 0, 50)

func reloadTopTags() {
	tags, err := x.NewTag_Selector().OrderBy_Count_Desc().Limit(TOP_TAGS_LIMIT).GetRows(base.DB)
	if err != nil {
		return
	}
	var topTagsViews = make([]*x.PB_TagView, 0, 50)
	for _, m := range tags {
		pb := &x.PB_TagView{
			TagId:         int64(m.TagId),
			Name:          m.Name,
			Count:         int32(m.Count),
			TagStatusEnum: int32(m.TagStatusEnum),
			CreatedTime:   int32(m.CreatedTime),
		}
		topTagsViews = append(topTagsViews, pb)
	}
	TopTagsViews = topTagsViews
}

func reloadTopPostsForTopTags() {
	tags := TopTagsViews
	var newTopTagsWithPosts = make([]*x.PB_TopTagWithSamplePosts, 0, 50)

	for _, t := range tags {
		postsIds, err := x.NewTagsPost_Selector().Select_PostId().TagId_Eq(int(t.TagId)).
			PostTypeEnum_Eq(int(x.PostTypeEnum_POST_PHOTO)).
			Limit(4).
			OrderBy_Id_Desc().
			GetIntSlice(base.DB)
		if err != nil {
			continue
		}

		v := &x.PB_TopTagWithSamplePosts{
			TagView:      t,
			PostViewList: view_service.PostsViewsForPostIds(postsIds, 0),
		}

		newTopTagsWithPosts = append(newTopTagsWithPosts, v)
	}

	TopTagsWithPostsResult = newTopTagsWithPosts
}
