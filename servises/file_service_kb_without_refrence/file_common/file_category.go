package file_common

import (

	"net/http"
)

type FileCategory struct {
	UrlPath       string
	cachePath     string
	//SetterToStore func(row *xc.FileRef)
	//GetterOfStore func(id int) (*xc.FileRef, error)
}

func (c FileCategory) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	row := NewRowReq(c, r.URL)

	_ = row
	//row.ServeHTTP(w, r)
}

////////////////////////////////

var MsgFileCategory = FileCategory{
	UrlPath:       "msg_file",
	cachePath:     "msg_file",
	//SetterToStore: file_store.SavaFile,
	//GetterOfStore: file_store_dep.MysqlGetmsgFile,
}

var PostFileCategory = FileCategory{
	UrlPath:       "post_file",
	cachePath:     "post_file",
	//SetterToStore: file_store_dep.MysqlSavePost,
	//GetterOfStore: file_store_dep.MysqlGetPostFile,
}

var AvatarFileCategory = FileCategory{
	UrlPath:       "avatar",
	cachePath:     "avatar",
	//SetterToStore: file_store_dep.MysqlSavePost,
	//GetterOfStore: file_store_dep.MysqlGetPostFile,
}

var HttpCategories = []FileCategory{
	MsgFileCategory, PostFileCategory, AvatarFileCategory,
}

/*

var MsgFileCategory = FileCategory{
	UrlPath:       "msg_file",
	cachePath:     "msg_file",
	SetterToStore: file_store_dep.MysqlSaveMsg,
	GetterOfStore: file_store_dep.MysqlGetmsgFile,
}

var PostFileCategory = FileCategory{
	UrlPath:       "post_file",
	cachePath:     "post_file",
	SetterToStore: file_store_dep.MysqlSavePost,
	GetterOfStore: file_store_dep.MysqlGetPostFile,
}

var AvatarFileCategory = FileCategory{
	UrlPath:       "avatar",
	cachePath:     "avatar",
	SetterToStore: file_store_dep.MysqlSavePost,
	GetterOfStore: file_store_dep.MysqlGetPostFile,
}

*/