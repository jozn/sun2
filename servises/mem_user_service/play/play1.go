package main

import (
    "ms/sun2/servises/mem_user_service"
    "ms/sun/helper"
)

func main() {
	mu := mem_user_service.AllMemUserMap.GetForUser(6)
    helper.PertyPrint(mu.GetFollowed())
    helper.PertyPrint(mu.GetFollowers())
    helper.PertyPrint(mu.GetFollowers())
    helper.PertyPrint(mu.GetFollowers())
    helper.PertyPrint(mu.GetFollowers())
    helper.PertyPrint(mu.GetFollowers())

    helper.PertyPrint(mu.GetLastPost())
    helper.PertyPrint(mu.GetLastActions())

}
