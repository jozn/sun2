package file_service

import (
	"fmt"
	"io/ioutil"
	"ms/sun/servises/file_service/file_common"
	"ms/sun/servises/file_service/file_disk_cache"
	"ms/sun/servises/file_service/file_resizer"
	"ms/sun/shared/helper"
	"net/http"
	"os"
)

var config = &file_common.FileServingConfig{
	DiskDirs:     []string{"D:/sun/a/", "D:/sun/b/"},
	FileServerId: 1,
}

func Run() {
	fmt.Println("file_service.Run() called")
	for _, cat := range file_common.HttpCategories {
		//http.Handle("/"+cat.UrlPath+"/", cat)
		http.HandleFunc(handler(cat))
	}
	//http.Handle("/getFile/", &fileCategory{})
	//http.Handle("/", &fileCategory{})
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("hi uu "))
	})
}
func handler(category file_common.FileCategory) (string, func(http.ResponseWriter, *http.Request)) {
	path := "/" + category.UrlPath + "/"
	helper.PertyPrint(path)
	fn := func(w http.ResponseWriter, r *http.Request) {
		row := file_common.NewRowReq(category, r.URL)
		file_disk_cache.SelectDisk(row, config)

		helper.PertyPrint(row)
		if row == nil {
			http.NotFound(w, r)
			return
		}
		if row.Err != nil {
			http.Error(w, row.Err.Error(), http.StatusBadRequest)
			return
		}

		switch {
		case row.RequestedImageSize > 0:
			serveResizeHTTP(row, w, r)
		case row.IsThumb:
			serveThumbHTTP(row, w, r)
		default:
			serveOriginalHTTP(row, w, r)
		}
	}
	return path, fn
}

func serveOriginalHTTP(row *file_common.RowReq, w http.ResponseWriter, r *http.Request) {
	fmt.Println("serveOriginalHTTP")
	//file_disk_cache.SelectDisk(row, config)
	//path, ok := file_disk_cache.IsLocalCacheAvailable(row, config)
	if row.IsOrginalLocalDiskCacheAvailable {
		http.ServeFile(w, r, row.RowCacheOutDiskPath)
		return
	}
	err := loadOriginalFromStore(row)
	if err == nil {
		http.ServeFile(w, r, row.RowCacheOutDiskPath)
	} else {
		http.NotFound(w, r)
	}
}

func serveResizeHTTP(row *file_common.RowReq, w http.ResponseWriter, r *http.Request) {
	fmt.Println("serveResizeHTTP")

	if !allowedSize[row.RequestedImageSize] {
		http.Error(w, file_common.ErrResizeNotAllowed.Error(), http.StatusNotFound)
		return
	}
	cachePathSized := row.RowCacheOutDiskPathSized //row.SelectedCacheDiskPath + row.RowCacheOutRelativePathSized
	if file_disk_cache.IsFileExists(cachePathSized) {
		http.ServeFile(w, r, cachePathSized)
		return
	}
	var err error
	if file_disk_cache.IsFileExists(row.RowCacheOutDiskPath) {

	} else {
		err = loadOriginalFromStore(row)
	}

	if err == nil {
		res := file_service.Resizer{
			InputFullPath:  row.RowCacheOutDiskPath,
			OutputFullPath: cachePathSized,
			Width:          row.RequestedImageSize,
			Quality:        90,
		}
		res.ResizeFFMPEG()
		http.ServeFile(w, r, cachePathSized)
	} else {
		http.NotFound(w, r)
	}
}

func serveThumbHTTP(row *file_common.RowReq, w http.ResponseWriter, r *http.Request) {
	fmt.Println("serveThumbHTTP")

	cachePathThumb := row.RowCacheOutDiskPathThumb
	if file_disk_cache.IsFileExists(cachePathThumb) {
		http.ServeFile(w, r, cachePathThumb)
		return
	}
	resizing := func() {
		res := file_service.Resizer{
			InputFullPath:  row.RowCacheOutDiskPath,
			OutputFullPath: cachePathThumb,
			Width:          150,
			Quality:        90,
		}
		res.ResizeFFMPEG()
	}
	var err error
	if row.IsOrginalLocalDiskCacheAvailable {
		//this means we have not thumb in db - so just resize the original to 100px
		resizing()
	} else {
		err = loadOriginalFromStore(row)
	}

	if err == nil {
		if file_disk_cache.IsFileExists(cachePathThumb) {
			http.ServeFile(w, r, cachePathThumb)
		} else {
			resizing()
			http.ServeFile(w, r, cachePathThumb)
		}
	} else {
		http.NotFound(w, r)
	}
}

func loadOriginalFromStore(row *file_common.RowReq) error {
	fmt.Println("loadOriginalFromStore", row)
	cassRow, err := row.FileCategory.GetterOfStore(row.FileDataStoreId)
	if err == nil {
		createRowOutCacheDir(row)
		ioutil.WriteFile(row.RowCacheOutDiskPath, cassRow.Data, os.ModePerm)
		if len(cassRow.DataThumb) > 0 {
			ioutil.WriteFile(row.RowCacheOutDiskPathThumb, cassRow.DataThumb, os.ModePerm)
		}
		return nil
	}
	fmt.Println("err: ", err)
	return err
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
