package file_service

import (
	"net/http"
)

type fileCategory struct {
	urlPath       string
	cachePath     string
	setterToStore func(row Row)
	getterOfStore func(id int) (*Row, error)
}

func (c fileCategory) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//os.Stdout = nil
	//fmt.Println("sasssss")
	row := NewRowReq(c, r.URL)
	//fmt.Println(row)
	//fmt.Println(err)
	row.ServeHTTP(w, r)
}

////////////////////////////////

var httpCategories = []fileCategory{
	{
		urlPath:       "msg_file",
		cachePath:     "msg_file",
		setterToStore: mysqlSaveMsg,
		getterOfStore: mysqlGetmsgFile,
	},
	{
		urlPath:       "post_file",
		cachePath:     "post_file",
		setterToStore: mysqlSavePost,
		getterOfStore: mysqlGetPostFile,
	},
	{
		urlPath:       "avatar",
		cachePath:     "avatar",
		setterToStore: mysqlSavePost,
		getterOfStore: mysqlGetPostFile,
	},
}
