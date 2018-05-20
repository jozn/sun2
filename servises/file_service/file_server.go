package file_service

import (
	"fmt"
	"ms/sun/servises/file_service/file_common"
	"ms/sun/servises/file_service/file_disk_cache"
	"ms/sun/servises/file_service/file_resizer"
	"ms/sun/shared/dbs"
	"ms/sun/shared/xc"
	"net/http"
	"os"
    "io/ioutil"
    "strings"
    "errors"
    c "github.com/patrickmn/go-cache"
    "time"
)

var RowCache = c.New(time.Second*1*3600, time.Second*60*5)

var config = &file_common.FileServingConfig{
	DiskDirs:     []string{"D:/sun/a/", "D:/sun/b/"},
	FileServerId: 1,
}

func Run() {
	fmt.Println("file_service.Run() called")

	openLoaclDb()

	for _, cat := range file_common.HttpCategories {
		//http.Handle("/"+cat.UrlPath+"/", cat)
		//try := &retryHttp{
		//	cat: cat,
		//}
		http.HandleFunc("/"+cat.UrlPath+"/", func(writer http.ResponseWriter, request *http.Request){
            try := &retryHttp{
                cat: cat,
            }
            try.ServeHTTP(writer,request)
        })
		//http.Handle("/"+cat.UrlPath+"/", try)
	}
	//http.Handle("/getFile/", &fileCategory{})
	//http.Handle("/", &fileCategory{})
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("hi uu "))
	})
}

func DeferCleanUp() {
    localdb.Close()
}
func serveOriginalHTTP(t *retryHttp) {
	fmt.Println("serveOriginalHTTP")
	//file_disk_cache.SelectDisk(row, config)
	//path, ok := file_disk_cache.IsLocalCacheAvailable(row, config)
	if t.row.IsOrginalLocalDiskCacheAvailable {
		t.httpServeFile(t.row.RowCacheOutDiskPath)
		return
	}
	err := loadOriginalFromStore(t)
	if err == nil {
		t.httpServeFile(t.row.RowCacheOutDiskPath)
	} else {
		//http.NotFound(w, r)
	}
}

func serveResizeHTTP(t *retryHttp) {
	fmt.Println("serveResizeHTTP")

	if !allowedSize[t.row.RequestedImageSize] {
		t.httpError(file_common.ErrResizeNotAllowed.Error(), http.StatusNotFound)
		return
	}

	if !allowedResizingExceptionsMap[t.row.FileExtensionWithDot] {
		t.httpError("can not resize this file.", http.StatusBadRequest)
		return
	}

	cachePathSized := t.row.RowCacheOutDiskPathSized //row.SelectedCacheDiskPath + row.RowCacheOutRelativePathSized
	if file_disk_cache.IsFileExists(cachePathSized) {
		t.httpServeFile(cachePathSized)
		return
	}
	var err error
	if file_disk_cache.IsFileExists(t.row.RowCacheOutDiskPath) {

	} else {
		err = loadOriginalFromStore(t)
	}

	if err == nil {
		res := file_service.Resizer{
			InputFullPath:  t.row.RowCacheOutDiskPath,
			OutputFullPath: cachePathSized,
			Width:          t.row.RequestedImageSize,
			Quality:        90,
		}
		res.ResizeFFMPEG()
		t.httpServeFile(cachePathSized)
	} else {
		//http.NotFound(w, t)
	}
}

func serveThumbHTTP(t *retryHttp) {
	fmt.Println("serveThumbHTTP")

	if !allowedThumbExceptionsMap[t.row.FileExtensionWithDot] {
		t.httpError("can not have thumbs of this file.", http.StatusBadRequest)
		return
	}

	cachePathThumb := t.row.RowCacheOutDiskPathThumb
	if file_disk_cache.IsFileExists(cachePathThumb) {
		t.httpServeFile(cachePathThumb)
		return
	}
	resizing := func() {
		res := file_service.Resizer{
			InputFullPath:  t.row.RowCacheOutDiskPath,
			OutputFullPath: cachePathThumb,
			Width:          150,
			Quality:        90,
		}
		res.ResizeFFMPEG()
	}
	var err error
	if t.row.IsOrginalLocalDiskCacheAvailable {
		//this means we have not thumb in localdb - so just resize the original to 100px
		resizing()
	} else {
		err = loadOriginalFromStore(t)
	}

	if err == nil {
		if file_disk_cache.IsFileExists(cachePathThumb) {
			t.httpServeFile(cachePathThumb)
		} else {
			resizing()
			t.httpServeFile(cachePathThumb)
		}
	} else {
		//http.NotFound(w, r)
	}
}

func loadOriginalFromStore(ret *retryHttp) error {
	row := ret.row
	fmt.Println("loadOriginalFromStore", row)
	cassRowFileRef, err := xc.NewFileRef_Selector().FileRefId_Eq(row.FileDataStoreId).GetRow(dbs.CassndraSession)
	if err != nil {
		return err
	}

	localdb.Put(fileToLocalDbKey(cassRowFileRef.FileRefId), fileToLocalDbKey(cassRowFileRef.FirstFileRefId))

	ret.row.SetNewFileDataStoreId(cassRowFileRef.FirstFileRefId)
	file_disk_cache.SelectDisk(row, config)
	if row.IsOrginalLocalDiskCacheAvailable {
		ret.markForRetry = true
		return nil
	}

	fileStorage, err := xc.NewFileStorage_Selector().Md5Key_Eq(cassRowFileRef.Md5Key).GetRow(dbs.CassndraSession)
	if err != nil {
		return err
	}

    //ret.row.SetNewFileDataStoreIdAndExtntion(fileStorage.FirstFileRefId,fileStorage.Extension)
	createRowOutCacheDir(row)
	if fileStorage.Extension == row.FileExtensionWithDot {
        ioutil.WriteFile(row.RowCacheOutDiskPath, fileStorage.Zdata, os.ModePerm)
        if len(fileStorage.ZdataThumb) > 0 {
            ioutil.WriteFile(row.RowCacheOutDiskPathThumb, fileStorage.ZdataThumb, os.ModePerm)
        }
    }else if (row.FileExtensionWithDot != "") && (fileStorage.Extension != "") && (strings.Index(row.FileExtensionWithDot, ".") != -1) && (strings.Index(fileStorage.Extension, ".") != -1){
        //to do add convert if supported
        if !allowedConvertingExceptionsMap[fileStorage.Extension] || !allowedConvertingExceptionsMap[row.FileExtensionWithDot] {
            return errors.New("can't convert this files files.")
        }

        orginalExtenPath := strings.Replace(row.RowCacheOutDiskPath, row.FileExtensionWithDot, fileStorage.Extension, -1)
        ioutil.WriteFile(orginalExtenPath, fileStorage.Zdata, os.ModePerm)
        if len(fileStorage.ZdataThumb) > 0 {
            //ioutil.WriteFile(row.RowCacheOutDiskPathThumb, cassRow.DataThumb, os.ModePerm)
        }

        width := 250
        if fileStorage.Width > 0 {
            width = fileStorage.Width
        }
        res := file_service.Resizer{
            InputFullPath:  orginalExtenPath,
            OutputFullPath: row.RowCacheOutDiskPath,
            Width:          width,
            Quality:        90,
        }
        res.ResizeFFMPEG()
    } else {
        return errors.New("can't convert files, with mismatch stored and requested")
    }

	ret.markForRetry = true

	return nil
}

func fileToLocalDbKey(fileId int) []byte {
	return []byte(fmt.Sprintf("%d", fileId))
}

func createRowOutCacheDir(r *file_common.RowReq) {
	os.MkdirAll(r.RowCacheOutDiskPathDir, os.ModeDir)
}

var allowedSize = map[int]bool{
	100:  true,
	250:  true,
	500:  true,
	1080: true,
}

/*
func handler(category file_common.FileCategory) (string, func(http.ResponseWriter, *http.Request)) {
	path := "/" + category.UrlPath + "/"
	helper.PertyPrint(path)
	fn := func(w http.ResponseWriter, r *http.Request) {

		wrap := &httpWirterWrapper{
			w: w,
			r: r,
		}

		row := file_common.NewRowReq(category, r.URL)
		file_disk_cache.SelectDisk(row, config)
		wrap.req = row

		//fmt.Println("111111111111***************************")
		defer func() {
			//fmt.Println("***************************")
			httpMonitorChan <- wrap
		}()

		helper.PertyPrint(row)
		if row == nil {
			http.NotFound(wrap, r)
			return
		}
		if row.Err != nil {
			http.Error(wrap, row.Err.Error(), http.StatusBadRequest)
			return
		}

		switch {
		case row.RequestedImageSize > 0:
			serveResizeHTTP(row, wrap, r)
		case row.IsThumb:
			serveThumbHTTP(row, wrap, r)
		default:
			serveOriginalHTTP(row, wrap, r)
		}
	}
	return path, fn
}
*/
