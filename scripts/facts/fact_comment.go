package main

import (
	"math/rand"
	"ms/sun/helper"
	"ms/sun2/servises/model_service"
    "ms/sun2/scripts/facts/fact_utils"
)

func main() {
	for i := 0; i < 10000; i++ {
		factComment()
	}
}
func factComment() {
	pid := fact_utils.GetRandPostId()
	model_service.Comment_Add(rand.Intn(80)+1, pid, helper.FactRandStrEmoji(25, true))
}
