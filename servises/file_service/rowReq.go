package file_service

import (
	"fmt"
	"io/ioutil"
	"ms/sun/helper"
	"net/http"
	"net/url"
	"os"
	"strings"
)

//const GALAXY_CACHE_PARENT_DIR = `C:\Go\_gopath\src\ms\sun2\files\`
const GALAXY_CACHE_PARENT_DIR = `E:\sun\`

type rowReq struct {
	fileCategory              fileCategory
	fileName                  string
	fileDataStoreId           int //64bit
	fileExtensionWithoutDot   string
	fileFormat                fileFormat
	requestedImageSize        int
	rowCacheOutFullPathDir    string
	rowCacheOutFullPath       string
	rowCacheOutFullPathThumb  string
	hasError                  bool
	isLocalDiskCacheAvailable bool
	cacheFullModuleDirectory  string // "/web/galxy/chat/"
}

func newRowReq(category fileCategory, url *url.URL) (row *rowReq, err error) {
	urlPath := strings.Trim(url.Path, "/")
	sep := strings.Split(urlPath, "/")
	fileName := sep[len(sep)-1]
	if len(fileName) < 6 {
		err = errFileNameTooShort
		return
	}
	id, size, ext, err := nameToParts(fileName)
	if err != nil {
		return
	}
	row = &rowReq{
		fileCategory:             category,
		fileName:                 fileName,
		fileDataStoreId:          id,
		fileExtensionWithoutDot:  ext,
		requestedImageSize:       size,
		cacheFullModuleDirectory: GALAXY_CACHE_PARENT_DIR + category.cachePath + "/", //"cache/",
	}
	err = row.setOutputChaseFullPath()
	if err != nil {
		return
	}
	return
}

//1254_24.jpg -> 1254 24 "jpg" nil  ------ 1254.jpg -> 1254 0 "jpg"
//erors : 125 125_   156
//todo add size support
func nameToParts(name string) (id int, size int, ext string, err error) {
	sep := strings.Split(name, "/")
	fileName := sep[len(sep)-1]
	sep2 := strings.Split(fileName, ".")
	//fmt.Println("6666666", name, "9999999")
	if len(sep2) == 2 && len(sep2[0]) > 0 && len(sep2[1]) > 0 {
		id = helper.StrToInt(sep2[0], 0)
		ext = sep2[1] //it dosent have any dot
		if id > 0 {
			return
		}
	}
	err = errWrongFileName
	return
}

func (r *rowReq) setOutputChaseFullPath() (err error) {
	ids := r.fileName
	//fmt.Println(ids)
	r.rowCacheOutFullPathDir = fmt.Sprintf("%s%s/%s/%s/%s/", r.cacheFullModuleDirectory,
		ids[0:2], ids[2:4], ids[4:6], ids[6:8])
	r.rowCacheOutFullPath = fmt.Sprintf("%s%d.%s", r.rowCacheOutFullPathDir, r.fileDataStoreId, r.fileExtensionWithoutDot)
	r.rowCacheOutFullPathThumb = fmt.Sprintf("%s%d_thumb.%s", r.rowCacheOutFullPathDir, r.fileDataStoreId, r.fileExtensionWithoutDot)

	return
}

func (r *rowReq) createRowOutCacheDir() {
	os.MkdirAll(r.rowCacheOutFullPathDir, os.ModeDir)
}

func (row *rowReq) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if row == nil {
		http.NotFound(w, r)
		return
	}
	if row.isLocalCacheAvailable() {
		http.ServeFile(w, r, row.rowCacheOutFullPath)
		return
	}
	//cassRow, ok := getFromCassandraById(row.fileDataStoreId)
	cassRow, err := row.fileCategory.getterOfStore(row.fileDataStoreId)
	if err == nil {
		row.createRowOutCacheDir()
		//fmt.Println("mysql", cassRow)
		ioutil.WriteFile(row.rowCacheOutFullPath, cassRow.Data, os.ModePerm)
		http.ServeFile(w, r, row.rowCacheOutFullPath)
	} else {
		http.NotFound(w, r)
	}
}

func (row *rowReq) isLocalCacheAvailable() bool {
	if _, err := os.Stat(row.rowCacheOutFullPath); os.IsNotExist(err) {
		return false
	}
	return true
}
