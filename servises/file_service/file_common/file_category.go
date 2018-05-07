package file_common

import (
	"ms/sun/servises/file_service/file_store"
	"net/http"
)

type FileCategory struct {
	UrlPath       string
	cachePath     string
	SetterToStore func(row file_store.Row)
	GetterOfStore func(id int) (*file_store.Row, error)
}

func (c FileCategory) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	row := NewRowReq(c, r.URL)

	_ = row
	//row.ServeHTTP(w, r)
}

////////////////////////////////

var HttpCategories = []FileCategory{
	{
		UrlPath:       "msg_file",
		cachePath:     "msg_file",
		SetterToStore: file_store.MysqlSaveMsg,
		GetterOfStore: file_store.MysqlGetmsgFile,
	},
	{
		UrlPath:       "post_file",
		cachePath:     "post_file",
		SetterToStore: file_store.MysqlSavePost,
		GetterOfStore: file_store.MysqlGetPostFile,
	},
	{
		UrlPath:       "avatar",
		cachePath:     "avatar",
		SetterToStore: file_store.MysqlSavePost,
		GetterOfStore: file_store.MysqlGetPostFile,
	},
}
