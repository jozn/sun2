package main_test

import (
	"github.com/golang/protobuf/proto"
	"ms/sun2/shared/x"
	"testing"
)

func BenchmarkPlay12(b *testing.B) {
	p := &x.PB_Post{
		//Id:             546454,
		UserId:         10,
		//TypeId:         54654,
		Text:           "ljgsdg slgjsdgl;g sd;g sdg;sdklgs ;gsd;glksd gsd;gk  sdglsdjg sgjsg sgsdjg ",
		//FormatedText:   "",
		MediaCount:     10,
		SharedTo:       20,
		DisableComment: 455654654,
		HasTag:         10,
		LikesCount:     05554565,
		CommentsCount:  0545465,
		EditedTime:     165410,
		CreatedTime:    450,
	}
	p2 := &x.PB_MessageView{}

	_ = p
	for i := 0; i < b.N; i++ {
		proto.Marshal(p2)
	}
}

func BenchmarkMap1(b *testing.B) {
    mp := make(map[int]bool,10000000)

    for i := 0; i < b.N; i++ {
        mp[i] = true
    }
}
