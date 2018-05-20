package main

import (
	"fmt"
	"ms/sun/shared/base"
	"ms/sun/shared/x"
	"net/http"
)

func main() {
	i := 0
	name := func(p *x.FilePost) string {
		if len(p.Extension) > 0 {
			return fmt.Sprintf("%d%s", p.Id, p.Extension)
		}
		return fmt.Sprintf("%d", p.Id)
	}

	nameMsg := func(p *x.FileMsg) string {
		if len(p.Extension) > 0 {
			return fmt.Sprintf("%d%s", p.Id, p.Extension)
		}
		return fmt.Sprintf("%d", p.Id)
	}

	name500 := func(p *x.FilePost) string {
		if len(p.Extension) > 0 {
			return fmt.Sprintf("%d_500%s", p.Id, p.Extension)
		}
		return fmt.Sprintf("%d_500", p.Id)
	}

	fn := func() {
		filePosts, _ := x.NewFilePost_Selector().GetRows(base.DB)
		for _, filePost := range filePosts {
			u := fmt.Sprintf("http://localhost:1100/post_file/%s", name(filePost))
			fmt.Println(i, u)
			http.Get(u)
			i++
		}

		for _, filePost := range filePosts {
			u := fmt.Sprintf("http://localhost:1100/post_file/%s", name500(filePost))
			fmt.Println(i, u)
			//http.Get(u)
			i++
		}

		fileMsgs, _ := x.NewFileMsg_Selector().GetRows(base.DB)
		for _, id := range fileMsgs {
			u := fmt.Sprintf("http://localhost:1100/msg_file/%s", nameMsg(id))
			fmt.Println(i, u)
			http.Get(u)
			i++
		}
	}

	//go fn()
	fn()

	fmt.Println(i)

}
