package main

import (
	"math/rand"
	"ms/sun2/servises/model_service"
    "ms/sun2/scripts/facts/fact_utils"
)

func main() {
	for i := 0; i < 10000; i++ {
        if rand.Intn(10) < 2 {
            factUnLike()
        }else {
            factLike()
        }
	}
}
func factLike() {
	model_service.Like_LikePost(fact_utils.GetRandUserId(),fact_utils.GetRandPostId())
}

func factUnLike() {
    model_service.Like_UnlikePost(fact_utils.GetRandUserId(),fact_utils.GetRandPostId())
}
