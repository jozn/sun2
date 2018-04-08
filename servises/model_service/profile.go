package model_service

import (
	"ms/sun_old/base"
	"ms/sun/shared/x"
)

type profilePager struct {
	Offset        int
	Limit         int
	ProfileUserId int
	UserName      string
}

func ProfileGetHome(pager profilePager) ([]int, error) {
	selector := x.NewPost_Selector().
		Select_PostId().
		UserId_Eq(pager.ProfileUserId).
		OrderBy_PostId_Desc()

	if pager.Limit > 0 {
		selector.PostId_LT(pager.Offset)
	}

	return selector.GetIntSlice(base.DB)
}

func ProfileGetMentioned(pager profilePager) ([]int, error) {
	selector := x.NewPostMentioned_Selector().
		Select_PostId().
		ForUserId_Eq(pager.ProfileUserId).
		OrderBy_MentionedId_Desc()

	if pager.Limit > 0 {
		selector.MentionedId_LT(pager.Offset)
	}

	return selector.GetIntSlice(base.DB)
}

func ProfileGetMedia(pager profilePager) ([]int, error) {
	return profileGetbyCat(pager, x.PostCategoryEnum_PostCat_Media)
}

func ProfileGetFile(pager profilePager) ([]int, error) {
	return profileGetbyCat(pager, x.PostCategoryEnum_PostCat_File)
}

func profileGetbyCat(pager profilePager, cat x.PostCategoryEnum) ([]int, error) {
	selector := x.NewPost_Selector().
		Select_PostId().
		PostCategoryEnum_Eq(int(cat)).
		UserId_Eq(pager.ProfileUserId).
		OrderBy_PostId_Desc()

	if pager.Limit > 0 {
		selector.PostId_LT(pager.Offset)
	}

	return selector.GetIntSlice(base.DB)
}
