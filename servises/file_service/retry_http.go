package file_service

import (
	"fmt"
	"github.com/akrylysov/pogreb"
	"log"
	"ms/sun/servises/file_service/file_common"
	"ms/sun/servises/file_service/file_disk_cache"
	"ms/sun/shared/helper"
	"net/http"
)

type retryHttp struct {
	cat                file_common.FileCategory
	w                  http.ResponseWriter
	r                  *http.Request
	row                *file_common.RowReq
	tryed              int
	fileStoredIdLoaded bool
	success            bool
	handelded          bool
	markForRetry       bool
	wrap               *httpWirterWrapper
}

func (t *retryHttp) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.r = r
	t.w = w

	t.wrap = &httpWirterWrapper{
		w: t.w,
		r: t.r,
	}

	if _row, ok := RowCache.Get(r.URL.Path); ok {
		if row2, ok := _row.(*file_common.RowReq); ok {
			t.wrap.req = *row2

			defer func() {
				httpMonitorChan <- t.wrap
			}()
			//fmt.Println("memory", row2.MemoryCacheFinalFullPathFile, r.URL.Path)
			http.ServeFile(t.wrap, r, row2.MemoryCacheFinalFullPathFile)
			return
		}
	}

	//load real loacal db
	t.row = file_common.NewRowReq(t.cat, r.URL)
	t.wrap.req = *t.row

	defer func() {
		httpMonitorChan <- t.wrap
	}()

	if t.row.Err != nil {
		t.httpError(t.row.Err.Error(), http.StatusBadRequest)
		return
	}

	fileSoteIdStr, err := localdb.Get([]byte(fmt.Sprintf("%d", t.row.FileRefId)))
	fileSoteId := helper.StrToInt(string(fileSoteIdStr), 0)
	if err == nil && fileSoteId > 0 {
		t.fileStoredIdLoaded = true
		t.row.SetNewFileDataStoreId(fileSoteId)
		t.tryServerMore()
	} else { //load from cassandra //todo can improve this if file existes on disk
		loadOriginalFromStore(t)
		t.tryServerMore()
	}

	if t.markForRetry && !t.handelded {
		//t.tryServerMore()
	}

	if !t.handelded {
		http.NotFound(t.wrap, t.r)
	}

}

func (t *retryHttp) httpError(err string, code int) {
	t.success = false
	t.handelded = true
	http.Error(t.wrap, err, code)
}

func (t *retryHttp) httpServeFile(filePath string) {
	t.success = true
	t.handelded = true
	//cache it
	RowCache.Set(t.row.FullUrlPath, t.row, 0)
	t.row.MemoryCacheFinalFullPathFile = filePath
	http.ServeFile(t.wrap, t.r, filePath)
}

func (t *retryHttp) tryServerMore() {
	t.wrap = &httpWirterWrapper{
		w: t.w,
		r: t.r,
	}

	//row := file_common.NewRowReq(category, r.URL)
	row := t.row
	file_disk_cache.SelectDisk(row, config)
	//t.wrap.req = row

	//fmt.Println("111111111111***************************")
	//defer func() {
	//	//fmt.Println("***************************")
	//	httpMonitorChan <- t.wrap
	//}()

	//helper.PertyPrint(t.row)
	if row == nil {
		t.handelded = true
		http.NotFound(t.wrap, t.r)
		return
	}
	if row.Err != nil {
		t.httpError(row.Err.Error(), http.StatusBadRequest)
		return
	}

	switch {
	case row.RequestedImageSize > 0:
		serveResizeHTTP(t)
	case row.IsThumb:
		serveThumbHTTP(t)
	default:
		serveOriginalHTTP(t)
	}
}

var localdb *pogreb.DB

func openLoaclDb() {
	//ooption :=  pogreb.Options{
	//    BackgroundSyncInterval:
	//}
	connected := false
	var err error
	for i := 1; i < 10; i++ {
		if connected {
			break
		}
		name := fmt.Sprintf("local_file_ref%d.localdb", i)
		localdb, err = pogreb.Open(name, nil)
		if err == nil {
			//log.Fatal(err)
			connected = true
		}
	}
	if !connected {
		log.Fatal(err)
	}
}
