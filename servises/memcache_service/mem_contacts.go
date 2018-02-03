package memcache_service

import (
	"fmt"
	c "github.com/patrickmn/go-cache"
	"ms/sun/base"
	"ms/sun2/shared/helper/go_map"
	"ms/sun2/shared/x"
	"time"
)

var contactsCache = c.New(time.Hour*24, time.Minute*1)

func DoesUserInPhoneContacts(MeId, PeerId int) bool {
	//fmt.Println("mmm", MeId,PeerId)
	if MeId < 1 || PeerId < 1 {
		return false
	}
	mp, ok := contactsCache.Get(getContactKey(MeId))
	var m *go_map.ConcurrentIntMap
	if ok {
		m, ok = mp.(*go_map.ConcurrentIntMap)
		if !ok { //should never happens
			return false
		}
	} else {
		m = fillContactsCacheForUser(MeId)
	}

	val := m.GetVal(PeerId)
	return val == 1
}

func GetUserContactsSlice(UserId int) []int {
	if UserId < 1 {
		return []int{}
	}
	mp, ok := contactsCache.Get(getContactKey(UserId))
	var m *go_map.ConcurrentIntMap
	if ok {
		m, ok = mp.(*go_map.ConcurrentIntMap)
		if ok { //should never happens
			return []int{}
		}
	} else {
		m = fillContactsCacheForUser(UserId)
	}

	return m.GetAllKeys()
}

/*
func AddToUserContacts(UserId, PeerId int) {
	if UserId < 1 || PeerId < 1 {
		return
	}
	mp, ok := contactsCache.Get(getContactKey(UserId))
	if !ok {
		return
	}
	m, ok := mp.(*go_map.ConcurrentIntMap)
	if !ok {
		return
	}
	m.SetVal(PeerId, int(x.FollowingEnum_FOLLOWING)) //todo fix me add enums
}

func DeleteFromUserFollwings(UserId, PeerId int) {
	if UserId < 1 || PeerId < 1 {
		return
	}
	mp, ok := contactsCache.Get(getContactKey(UserId))
	if !ok {
		return
	}
	m, ok := mp.(*go_map.ConcurrentIntMap)
	if !ok {
		return
	}
	m.Delete(PeerId)
}
*/

//////// Internals ////////
func getContactKey(UserId int) string {
	return fmt.Sprintf("_Contacts_%d", UserId)
}

func fillContactsCacheForUser(userId int) *go_map.ConcurrentIntMap {
	phones, err := x.NewPhoneContact_Selector().Select_Phone().UserId_Eq(userId).GetIntSlice(base.DB)
	if err != nil && len(phones) == 0 {
		//just return something not to painc
		return go_map.NewConcurrentIntMap(0)
	}
	userIds, err := x.NewUser_Selector().Select_UserId().Phone_In(phones).GetIntSlice(base.DB)
	m := go_map.NewConcurrentIntMap(len(userIds))
	m.SetKeys(phones, 1)
	contactsCache.Set(getContactKey(userId), m, 0)
	return m
}

/*
func fillContactsCacheForUser1(userId int) *go_map.ConcurrentIntMap {
	postIds, err := x.NewFollowingListMember_Selector().Select_FollowedUserId().UserId_Eq(userId).GetIntSlice(base.DB)
	if err != nil {
		//just return something not to painc
		return go_map.NewConcurrentIntMap(0)
	}
	m := go_map.NewConcurrentIntMap(len(postIds))
	m.SetKeys(postIds, 1)
	contactsCache.Set(getContactKey(userId), m, 0)
	return m
}

func Contacts_GetChachesContactsUserIdsForUserId(UserId, LastTime int) *ds.IntList {
	key := fmt.Sprintf("UserContacts:%d", UserId) //in Views.GetPhoneForUserIfIsContact()

	val, ok := Cacher.Get(key)
	if ok {
		collection, ok := val.(*ds.IntList)
		if ok {
			return collection
		}
	}

	var contactsUsers []int
	collection := ds.New()

	sel := x.NewPhoneContact_Selector().Select_PhoneNormalizedNumber().
		UserId_Eq(UserId).
		PhoneNormalizedNumber_NotEq("")

	if LastTime > 0 {
		sel.CreatedTime_GE(LastTime)
	}

	phones, err := sel.GetStringSlice(base.DB)
	if err != nil {
		return collection
	}

	if len(phones) > 0 {
		contactsUsers, err = x.NewUser_Selector().Select_Id().Phone_In(phones).OrderBy_Id_Desc().GetIntSlice(base.DB)
		if err != nil {
			return collection
		}

		collection.AddAndSort(contactsUsers...)
		//this is used in MemoryStore.GetPhoneForUserIfIsContact()
		Cacher.Set(key, collection, time.Hour*4)
	}
	return collection
}
*/
