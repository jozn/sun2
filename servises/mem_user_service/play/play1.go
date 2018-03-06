package main

import (
    "ms/sun2/servises/mem_user_service"
    "ms/sun/helper"
    "time"
)

func main() {
    helper.PertyPrint(helper.NanoTimeBeforeNowSeconds(10))
    helper.PertyPrint(helper.TimeNowNano())
	mu := mem_user_service.AllMemUserMap.GetForUser(6)
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
