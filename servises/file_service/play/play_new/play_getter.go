package main

import (
	"fmt"
	"ms/sun/servises/file_service"
	"ms/sun/shared/dbs"
	"ms/sun/shared/helper"
	"ms/sun/shared/xc"
	"net/http"
	"time"
)

func main() {
	i := 0
	//os.Stdout = nil
	name := func(p *xc.FileRef) string {
		if len(p.Extension) > 0 {
			return fmt.Sprintf("%d%s", p.FileRefId, p.Extension)
		}
		return fmt.Sprintf("%d", p.FirstFileRefId)
	}

	fn := func() {
		filePosts, err := xc.NewFileRef_Selector().GetRows(dbs.CassndraSession)
		fmt.Println(err,filePosts)
		for _, filePost := range filePosts {
			u := fmt.Sprintf("http://localhost:1100/post_file/%s", name(filePost))
			fmt.Println(i, u)
			http.Get(u)
			i++
            fmt.Println(i)
		}
	}
	fmt.Println(1)
	http.HandleFunc("/hi", func(writer http.ResponseWriter, r *http.Request) {
		writer.Write([]byte("hi ========="))
	})
	go func() {
		time.Sleep(time.Second)
		http.Get("http://localhost:1100/post_file/1518506476136010007_180.jpg")
	}()

	file_service.Run()

	go func() {
		err := http.ListenAndServe(":1100", nil)
		helper.DebugErr(err)
	}()

	fn()
    time.Sleep(time.Hour)
	//go fn()
	/* for _, filePost := range filePosts {
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
	   }*/

}
