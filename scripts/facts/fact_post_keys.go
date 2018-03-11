package main

import (
	"ms/sun/base"
	"ms/sun/helper"
	"ms/sun2/shared/x"
)

func main() {
	for i := 0; i < 500; i++ {
		var postKeys []x.PostKey
		for j := 0; j < 200; j++ {
			k := x.PostKey{
				PostKeyStr:  helper.RandAlphaNumbericString(6),
				Used: 0,
			}
			postKeys = append(postKeys, k)

		}
		x.MassReplace_PostKey(postKeys, base.DB)
	}
}
