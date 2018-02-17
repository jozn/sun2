package file_service

import (
	"fmt"
	"io/ioutil"
	"ms/sun/helper"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"
)

//const GALAXY_CACHE_PARENT_DIR = `C:\Go\_gopath\src\ms\sun2\files\`
const GALAXY_CACHE_PARENT_DIR = `E:\sun\`

type rowReq struct {
	fullPath                  string
	fileCategory              fileCategory
	fileName                  string
	fileDataStoreId           int //64bit
	fileExtensionWithoutDot   string
	fileFormat                fileFormat
	requestedImageSize        int
	isThumb                   bool
	rowCacheOutFullPathDir    string
	rowCacheOutFullPath       string
	rowCacheOutFullPathSized  string
	rowCacheOutFullPathThumb  string
	isLocalDiskCacheAvailable bool
	cacheFullModuleDirectory  string // "/web/galxy/chat/"
	err                       error
}

func NewRowReq(category fileCategory, url *url.URL) (row *rowReq) {
	//fmt.Println(url.Path)
	row = &rowReq{
		fullPath:                 url.Path,
		fileCategory:             category,
		cacheFullModuleDirectory: GALAXY_CACHE_PARENT_DIR + category.cachePath + "/", //"cache/",
	}
	row.extractParams()
	row.setOutputChaseFullPath()

	if row.fileDataStoreId == 0 || row.fileExtensionWithoutDot == "" {
		row.err = errBadReq
	}
	return
}

func (r *rowReq) setOutputChaseFullPath() (err error) {
	ids := r.fileName
	//fmt.Println(ids)
	if len(ids) < 10 {
		r.err = errFileNameTooShort
		return errFileNameTooShort
	}
	r.rowCacheOutFullPathDir = fmt.Sprintf("%s%s/%s/%s/%s/", r.cacheFullModuleDirectory,
		ids[0:2], ids[2:4], ids[4:6], ids[6:8])
	r.rowCacheOutFullPath = fmt.Sprintf("%s%d.%s", r.rowCacheOutFullPathDir, r.fileDataStoreId, r.fileExtensionWithoutDot)
	r.rowCacheOutFullPathThumb = fmt.Sprintf("%s%d_thumb.%s", r.rowCacheOutFullPathDir, r.fileDataStoreId, r.fileExtensionWithoutDot)
	if r.requestedImageSize > 0 {
		r.rowCacheOutFullPathSized = fmt.Sprintf("%s%d_%d.%s", r.rowCacheOutFullPathDir, r.fileDataStoreId, r.requestedImageSize, r.fileExtensionWithoutDot)
	}
	return
}

// matches : [[/1518506476136010007_thumb.jpg 1518506476136010007_thumb.jpg _thumb]]
//var fileRegex = regexp.MustCompile(`/(\w+(_[[:alpha:]]+)\..+)`)
var fileRegex = regexp.MustCompile(`/([0-9]+(_[[:alnum:]]+)?\.\w+)`)

func (r *rowReq) extractParams() {
	parts := fileRegex.FindStringSubmatch(r.fullPath)
	//fmt.Println(parts)
	if len(parts) < 2 {
		r.err = errBadReq
	}
	if len(parts) >= 2 {
		r.fileName = parts[1]
		r.requestedImageSize = 0
		r.isThumb = false
	}
	if len(parts) == 3 {
		if parts[2] == "_thumb" {
			r.isThumb = true
		} else {
			if len(parts[2]) > 1 {
				r.requestedImageSize = helper.StrToInt(parts[2][1:], 0)
			}
		}
	}
	sep := strings.Split(r.fileName, ".")
	if len(sep) == 2 {
		r.fileDataStoreId = helper.StrToInt(strings.Split(sep[0], "_")[0], 0)
		r.fileExtensionWithoutDot = sep[1]
	} else {
		r.err = errBadReq
		return
	}

	return
}

func (r *rowReq) createRowOutCacheDir() {
	os.MkdirAll(r.rowCacheOutFullPathDir, os.ModeDir)
}

func (row *rowReq) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	helper.PertyPrint(row)
	if row == nil {
		http.NotFound(w, r)
		return
	}
	if row.err != nil {
		http.Error(w, row.err.Error(), http.StatusBadRequest)
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

////////////////////////////////////////////////
func newRowReq_bk(category fileCategory, url *url.URL) (row *rowReq, err error) {
	urlPath := strings.Trim(url.Path, "/")
	sep := strings.Split(urlPath, "/")
	fileName := sep[len(sep)-1]
	if len(fileName) < 10 {
		err = errFileNameTooShort
		return
	}
	id, size, ext, err := nameToParts(fileName)
	if err != nil {
		return
	}
	fmt.Println(url.RequestURI())

	row = &rowReq{
		fullPath:                 url.RequestURI(),
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
