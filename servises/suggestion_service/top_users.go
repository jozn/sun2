package suggestion_service

import (
	"math"
	"ms/sun/base"
	"ms/sun/ds"
	"ms/sun/helper"
	"ms/sun2/servises/memcache_service"
	"ms/sun2/shared/x"
	"time"
)

const TOP_USERS_LIMIT = 100

var TopUsers []int

func job_TopTrendingUsersCalculate_Infinixte() {
	defer helper.JustRecover()
	helper.SleepForDebugDelay(5)

	for {
		TopUsers = genTopTrendingUsers()
		time.Sleep(time.Minute * 30)
	}
}

func Scheduler_GenUserSuggestions(ForUserId int) {
	m, ok := x.Store.UserMetaInfo_ByUserId(ForUserId)
	if ok {
		if m.LastUserRecGen+4*3600 < helper.TimeNow() {
			genUserSuggestionForUserAndSaveMeta_BG(ForUserId)
		}
	} else {
		m = &x.UserMetaInfo{
			UserId: ForUserId,
		}
		m.LastUserRecGen = 0
		m.Save(base.DB)
		genUserSuggestionForUserAndSaveMeta_BG(ForUserId)
	}
}
func genUserSuggestionForUserAndSaveMeta_BG(ForUserId int) {
	go func() {
		defer helper.JustRecover()
		genSuggestionForUser(ForUserId)
		m, ok := x.Store.UserMetaInfo_ByUserId(ForUserId)
		if ok {
			m.LastUserRecGen = helper.TimeNow()
			m.Save(base.DB)
		}
	}()

}

func genSuggestionForUser(ForUserId int) {
	//for skipping
	iFollowing, err1 := x.NewFollowingListMember_Selector().
		Select_FollowedUserId().
		UserId_Eq(ForUserId).
		GetIntSlice(base.DB)

	followedMe, err2 := x.NewFollowingListMember_Selector().
		Select_UserId().
		FollowedUserId_Eq(ForUserId).
		Limit(200).
		OrderBy_Id_Desc().
		GetIntSlice(base.DB)

	myContacts := memcache_service.GetUserContactsSlice(ForUserId)

	if err1 != nil || err2 != nil {
		return
	}
	mpFollowing := make(map[int]bool, len(iFollowing))

	for _, uid := range iFollowing {
		mpFollowing[uid] = true
	}

	suggestions := make([]x.RecommendUser, 0, 200)

	cnt1 := 0
	for _, uid := range myContacts {
		if !mpFollowing[uid] {
			row := x.RecommendUser{
				UserId:      ForUserId,
				TargetId:    uid,
				Weight:      0.5,
				CreatedTime: helper.TimeNow(),
			}
			suggestions = append(suggestions, row)
			cnt1++
		}
		if cnt1 >= 100 {
			break
		}
	}

	cnt2 := 0
	for _, uid := range followedMe {
		if !mpFollowing[uid] {
			row := x.RecommendUser{
				UserId:      ForUserId,
				TargetId:    uid,
				Weight:      0.2,
				CreatedTime: helper.TimeNow(),
			}
			suggestions = append(suggestions, row)
			cnt2++
		}
		if cnt2 >= 100 {
			break
		}
	}

	cnt3 := 0
	for _, uid := range TopUsers {
		if !mpFollowing[uid] {
			row := x.RecommendUser{
				UserId:      ForUserId,
				TargetId:    uid,
				Weight:      0.3,
				CreatedTime: helper.TimeNow(),
			}
			suggestions = append(suggestions, row)
			cnt3++
		}
		if cnt3 >= 20 {
			break
		}
	}

	x.MassInsert_RecommendUser(suggestions, base.DB)
}

func genTopTrendingUsers() []int {
	rows, err := x.NewFollowingListMember_Selector().Limit(20000).OrderBy_Id_Desc().GetRows(base.DB)
	if err != nil {
		return []int{}
	}

	mp := make(map[int]int, 20000) // map[FollowedUserId] = count
	for _, r := range rows {
		mp[r.FollowedUserId] += 1
	}

	coll := ds.New()
	for user, _ := range mp {
		coll.Add(user)
	}
	coll.SortDesc()

	min := math.Min(float64(coll.Size()), TOP_USERS_LIMIT)

	return coll.Values()[:int(min)]
}

/*
func Recommend_ReGenUsersForUser(ForUserId int) {
	contacts_coll := memcache_service.Con(ForUserId, 0) //contacts UserIds list
	//following_coll := MemoryStore.UserFollowingList_Get(ForUserId)
	following_coll := memcache_service.GetUserFollowsSlice(ForUserId)

	notUserIds := memcache_service.GetUserFollowsSlice(ForUserId) //make([]int,0,len(following_coll) + len(contacts_coll))

	//select those that they follows me but i don'f follow them
	toFollowRows, err := x.NewFollowingListMember_Selector().
		Select_UserId().
		FollowedUserId_Eq(ForUserId).
		UserId_NotIn(notUserIds).
		Limit(200).
		GetIntSlice(base.DB)

	if err != nil {
		return
	}

	toFollow := ds.New()
	toFollow.AddAndSort(toFollowRows...)

	for _, id := range contacts_coll.Values() {
		if !following_coll.Contains(id) { //if we don't follow this contacts already
			toFollow.Add(id)
		}
	}
	toFollow.SortDesc()

	var rows []x.RecommendUser
	for _, id := range toFollow.Values() {
		rows = append(rows, x.RecommendUser{
			UserId:      ForUserId,
			TargetId:    id,
			Weight:      0.5,
			CreatedTime: helper.TimeNow(),
		})
	}

	for _, id := range TopUsers {
		rows = append(rows, x.RecommendUser{
			UserId:      ForUserId,
			TargetId:    id,
			Weight:      0.2,
			CreatedTime: helper.TimeNow(),
		})
	}

	x.MassInsert_RecommendUser(rows, base.DB)

}

*/
