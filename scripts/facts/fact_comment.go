package main

import (
	"math/rand"
	"ms/sun/shared/helper"
	"ms/sun/servises/model_service"
    "ms/sun/scripts/facts/fact_utils"
)

func main() {
	for i := 0; i < 100000; i++ {
		factComment()
	}
}
func factComment() {
	pid := fact_utils.GetRandPostId()
	model_service.Comment_Add(rand.Intn(80)+1, pid, helper.FactRandStrEmoji(25, true))
}
