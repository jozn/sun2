package main

import (
	"log"
	"math/rand"
	"ms/sun/helper"
	"ms/sun2/servises/model_service"
    "ms/sun2/scripts/facts/fact_utils"
)

func main() {
	for i := 0; i < 1000; i++ {
		if rand.Intn(10) <= 6 {
			factPostText()
		} else {
			factPostImage()
		}
	}
}

func factPostText() {
	p := model_service.PostAddParam{
		Text:   helper.FactRandStrEmoji(100, true),
		UserId: rand.Intn(80) + 1,
	}
	model_service.AddPost(p)
}

func factPostImage() {
	_, bts, _ := fact_utils.RandImage()
	p := model_service.PostAddParam{
		Text:       helper.FactRandStrEmoji(100, true),
		UserId:     rand.Intn(80) + 1,
		MediaBytes: bts,
	}
	if rand.Intn(3) == 2 {
		p.Text = ""
	}
	_, err := model_service.AddPost(p)
	if err != nil {
		log.Fatal(err)
	}
}


