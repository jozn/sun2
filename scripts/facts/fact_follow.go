package main

import (
	"math/rand"
	"ms/sun2/servises/model_service"
    "ms/sun2/scripts/facts/fact_utils"
)

func main() {
	for i := 0; i < 10000; i++ {
        if rand.Intn(10) < 2 {//20%
            factUnFollow()
        }else {
            factFollow()
        }
	}
}
func factFollow() {
	model_service.Follow(fact_utils.GetRandUserId(),fact_utils.GetRandUserId())
}

func factUnFollow() {
    model_service.UnFollow(fact_utils.GetRandUserId(),fact_utils.GetRandUserId())
}
