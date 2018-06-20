package model_service

import (
	"ms/sun/shared/helper"
	"ms/sun/shared/x"
	"ms/sun_old/base"
)

func postSaveTags(post x.Post, parser helper.TextParser) {
	x.Store.PreLoadTag_ByNames(parser.Tags)

	tags := []*x.Tag{}
	tagPosts := []x.TagPost{}
	tagsIds := []int{}
	for _, tagName := range parser.Tags {
		tg, ok := x.Store.Tag_ByName(tagName)
		if !ok {
			tg = &x.Tag{
				TagId:       helper.NextRowsSeqId(),
				Name:        tagName,
				CreatedTime: helper.TimeNow(),
			}
			tg.Save(base.DB)
		}
		//for general all queries
		tgp := x.TagPost{
			TagId:       tg.TagId,
			PostId:      post.PostId,
			PostType:    0,
			CreatedTime: helper.TimeNow(),
		}
		tagPosts = append(tagPosts, tgp)

		//for real types queries
		tgp2 := x.TagPost{
			TagId:       tg.TagId,
			PostId:      post.PostId,
			PostType:    post.PostType,
			CreatedTime: helper.TimeNow(),
		}
		tagPosts = append(tagPosts, tgp2)

		//if id medias
		switch x.PostTypeEnum(post.PostType) {
		case x.PostTypeEnum_POST_PHOTO, x.PostTypeEnum_POST_VIDEO, x.PostTypeEnum_POST_GIF:
			tgp3 := x.TagPost{
				TagId:       tg.TagId,
				PostId:      post.PostId,
				PostType:    int(x.PostTypeEnum_POST_MEDIA),
				CreatedTime: helper.TimeNow(),
			}
			tagPosts = append(tagPosts, tgp3)
		}

		tags = append(tags, tg)
		tagPosts = append(tagPosts, tgp)
		tagsIds = append(tagsIds, tg.TagId)
	}

	if len(tagPosts) > 0 {
		x.MassReplace_TagPost(tagPosts, base.DB)
	}
	if len(tagsIds) > 0 {
		x.NewTag_Updater().Count_Increment(1).TagId_In(tagsIds).Update(base.DB)
	}
}

func postSaveMentioned(post x.Post, parser helper.TextParser) {
	x.Store.PreLoadUser_ByUserNames(parser.UserNames)

	var arr []x.ProfileMentioned
	for _, username := range parser.UserNames {
		otherUser, ok := x.Store.User_ByUserName_JustCache(username)
		if !ok {
			continue
		}
		pm := x.ProfileMentioned{
			Id: helper.NextRowsSeqId(),
			ForUserId:   otherUser.UserId,
			PostId:      post.PostId,
			PostUserId:  post.UserId,
			PostType:    post.PostType,
			CreatedTime: helper.TimeNow(),
		}
		arr = append(arr, pm)
	}
	if len(arr) > 0 {
		x.MassReplace_ProfileMentioned(arr, base.DB)
	}

	//todo add mentioned notifications
}

//tags and menthed
func postTextsRelationshipsHandle(post x.Post) {
	if len(post.Text) == 0 {
		return
	}

	parser := helper.ParseText(post.Text)

	if len(parser.Tags) > 0 {
		postSaveTags(post, parser)
	}

	if len(parser.UserNames) > 0 {
		postSaveMentioned(post, parser)
	}

}
