package session_service

import (
	"ms/sun/shared/base"
	"ms/sun/shared/helper"
	"ms/sun/shared/x"
	"net/http"
	"sync"
)

type globalSessionMapType struct {
	sync.RWMutex
	mp map[string][]*x.Session
}

var lock sync.RWMutex
var mp = make(map[string][]*x.Session, 1000)

func CheckSeesionHttp(r *http.Request) int {
	sesssionUuid := r.URL.Query().Get("session")
	user_id := helper.StrToInt(r.URL.Query().Get("user_id"), 0)
	if user_id < 1 || len(sesssionUuid) == 0 {
		return 0
	}

	return CheckSeesion(sesssionUuid, user_id)
}

func CheckSeesion(uuid string, user_id int) int {
	user_id2 := checkMemForUserUuid(uuid, user_id)
	if user_id2 != 0 {
		return user_id2
	}
	loadSessionsFromTable(uuid)
	return checkMemForUserUuid(uuid, user_id)
}

func loadSessionsFromTable(uuid string) {
	rows, err := x.NewSession_Selector().SessionUuid_Eq(uuid).GetRows(base.DB)
	if err != nil {
		return
	}

	lock.Lock()
	memRows := mp[uuid]
	for _, row := range rows {
		exisit := 0
		for _, memRow := range memRows {
			if row.Id == memRow.Id {
				exisit++
			}
		}
		if exisit == 0 {
			mp[uuid] = append(mp[uuid], row)
		}
	}
	lock.Unlock()
}

func checkMemForUserUuid(uuid string, user_id int) int {
	lock.RLock()
	sesRows := mp[uuid]
	lock.RUnlock()

	for _, mRow := range sesRows {
		if mRow.UserId == user_id && mRow.SessionUuid == uuid {
			return user_id
		}
	}
	return 0
}
