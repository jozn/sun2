package main

import (
	"fmt"
	"ms/sun/servises/session_service"
	"ms/sun/shared/base"
	"ms/sun/shared/x"
)

func main() {
	for i := 1; i < 10; i++ {
		uuid := fmt.Sprintf("z%d", i)
		s := x.Session{
			SessionUuid:   uuid,
			UserId:        6,
			LastIpAddress: "",
			AppVersion:    1,
			ActiveTime:    0,
			CreatedTime:   0,
		}
		s.Save(base.DB)
	}

	for i := -4; i < 10; i++ {
		uuid := fmt.Sprintf("z%d", i)
		uid := session_service.CheckSeesion(uuid, 6)
		fmt.Println("result: ", uid, i)
	}
}
