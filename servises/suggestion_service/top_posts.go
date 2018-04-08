package suggestion_service

import (
	"ms/sun_old/base"
	"ms/sun_old/helper"
	"ms/sun/shared/golib/sorter"
	"ms/sun/shared/x"
	"sort"
	"time"
)

var TopPosts []int
var TopPhotoPosts []int

//////////////// Jobs /////////////////////

func job_TopTrendingPosts_Infinite() {
	defer helper.JustRecover()
	helper.SleepForDebugDelay(1)

	for {
		genTopTrendingPosts()
		time.Sleep(time.Minute * 30)
	}
}

////////////////////////////////////////////

func genTopTrendingPosts() {
	const SIZE = 100000
	postIds, err := x.NewLike_Selector().
		Select_PostId().Limit(SIZE).
		//PostTypeEnum_Eq(int(x.PostTypeEnum_POST_PHOTO)).
		OrderBy_Id_Desc().
		GetIntSlice(base.DB)
	if err != nil || len(postIds) == 0 {
		return
	}

	//number of likes per posts
	mp := make(map[int]int, SIZE)
	for _, r := range postIds {
		mp[r] += 1
	}

	arr := make([]sorter.IntWeight, 0, len(mp))
	for pId, cnt := range mp {
		arr = append(arr, sorter.IntWeight{Id: pId, Weight: cnt})
	}
	arr2 := sorter.IntWeightSlice(arr)
	sort.Sort(arr2)

	//min := math.Min(float64(len(arr2)), float64(5000))
	var topPostsTemp []int
	for i, pw := range arr2 {
		topPostsTemp = append(topPostsTemp, pw.Id)
		if i >= 5000 {
			break
		}
	}

	TopPosts = topPostsTemp

	//Top photo posts
	//for top photos check if
	postPhotoIds, err := x.NewPost_Selector().
		Select_PostId().
		PostId_In(topPostsTemp).
		PostTypeEnum_Eq(int(x.PostTypeEnum_POST_PHOTO)).
		Limit(1000).
		GetIntSlice(base.DB)
	if err != nil || len(postIds) == 0 {
		return
	}

	TopPhotoPosts = postPhotoIds

}
