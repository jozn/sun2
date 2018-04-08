package main

import (
    "ms/sun/servises/mem_user_service"
    "ms/sun_old/helper"
    "time"
)

func main() {
    helper.PertyPrint(helper.NanoTimeBeforeNowSeconds(10))
    helper.PertyPrint(helper.TimeNowNano())
	mu := mem_user_service.allMemUserMap.GetForUser(6)
    helper.PertyPrint(mu.GetFollowed())
    helper.PertyPrint(mu.GetFollowers())
    helper.PertyPrint(mu.GetFollowers())
    helper.PertyPrint(mu.GetFollowers())
    helper.PertyPrint(mu.GetFollowers())
    helper.PertyPrint(mu.GetFollowers())

    helper.PertyPrint(mu.GetLastPost())
    helper.PertyPrint(mu.GetLastActions())

    time.Sleep(time.Minute)

}
