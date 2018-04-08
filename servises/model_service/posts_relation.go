package model_service

import (
	"ms/sun/base"
	helper2 "ms/sun/helper"
	"ms/sun2/shared/helper"
	"ms/sun2/shared/x"
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
				TagId:       helper2.NextRowsSeqId(),
				Name:        tagName,
				CreatedTime: helper2.TimeNow(),
			}
			tg.Save(base.DB)
		}
		tgp := x.TagPost{
			TagId:        tg.TagId,
			PostId:       post.PostId,
			PostTypeEnum: post.PostTypeEnum,
			CreatedTime:  helper2.TimeNow(),
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

func postSaveMenthened(post x.Post, parser helper.TextParser) {
	x.Store.PreLoadUser_ByUserNames(parser.UserNames)

	var arr []x.PostMentioned
	for _, username := range parser.UserNames {
		otherUser, ok := x.Store.User_ByUserName_JustCache(username)
		if !ok {
			continue
		}
		pm := x.PostMentioned{
			MentionedId:      helper2.NextRowsSeqId(),
			ForUserId:        otherUser.UserId,
			PostId:           post.PostId,
			PostUserId:       post.UserId,
			PostTypeEnum:     post.PostTypeEnum,
			PostCategoryEnum: post.PostCategoryEnum,
			CreatedTime:      helper2.TimeNow(),
		}
		arr = append(arr, pm)
	}
	if len(arr) > 0 {
		x.MassReplace_PostMentioned(arr, base.DB)
	}
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
		postSaveMenthened(post, parser)
	}

}
