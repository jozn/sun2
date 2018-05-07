package file_common

import (
	"fmt"
	"ms/sun/shared/helper"
	"net/url"
	"os"
	"regexp"
	"strings"
)

//const GALAXY_CACHE_PARENT_DIR = `C:\Go\_gopath\src\ms\sun\files\`
const CACHE_DIR_VERSION = `v1`

type RowReq struct {
	FullUrlPath                  string
	FileCategory                 FileCategory
	FileName                         string
	FileDataStoreId                  int //64bit
	FileExtensionWithoutDot          string
	FileFormat                       fileFormat
	RequestedImageSize               int
	IsThumb                          bool
	RowCacheOutRelativePathDir       string
	RowCacheOutRelativePath          string
	RowCacheOutRelativePathSized     string
	RowCacheOutRelativePathThumb     string
	RowCacheOutDiskPathDir           string
	RowCacheOutDiskPath              string
	RowCacheOutDiskPathSized         string
	RowCacheOutDiskPathThumb         string
	IsOrginalLocalDiskCacheAvailable bool
	CacheFullModuleDirectory         string // "/web/galxy/chat/"
	SelectedCacheDiskPath            string //web/file/a disk
	Err                              error
}

func NewRowReq(category FileCategory, url *url.URL) (row *RowReq) {
	//fmt.Println(url.Path)
	row = &RowReq{
		FullUrlPath:              url.Path,
		FileCategory:             category,
		CacheFullModuleDirectory: CACHE_DIR_VERSION + "/" + category.cachePath + "/",
	}
	row.extractParams()
	row.CacheFullModuleDirectory = row.CacheFullModuleDirectory + row.FileExtensionWithoutDot + "/"
	//row.FileExtensionWithoutDot = "media"

	row.setOutputCacheFullPath()

	if row.FileDataStoreId == 0 || row.FileExtensionWithoutDot == "" {
		row.Err = ErrBadReq
	}
	return
}

func (r *RowReq) setOutputCacheFullPath() (err error) {
	ids := r.FileName
	//fmt.Println(ids)
	if len(ids) < 10 {
		r.Err = ErrFileNameTooShort
		return ErrFileNameTooShort
	}
	r.RowCacheOutRelativePathDir = fmt.Sprintf("%s%s/%s/%s/%s/", r.CacheFullModuleDirectory,
		ids[0:2], ids[2:4], ids[4:6], ids[6:8])
	r.RowCacheOutRelativePath = fmt.Sprintf("%s%d.%s", r.RowCacheOutRelativePathDir, r.FileDataStoreId, r.FileExtensionWithoutDot)
	r.RowCacheOutRelativePathThumb = fmt.Sprintf("%s%d_thumb.%s", r.RowCacheOutRelativePathDir, r.FileDataStoreId, r.FileExtensionWithoutDot)
	if r.RequestedImageSize > 0 {
		r.RowCacheOutRelativePathSized = fmt.Sprintf("%s%d_%d.%s", r.RowCacheOutRelativePathDir, r.FileDataStoreId, r.RequestedImageSize, r.FileExtensionWithoutDot)
	}
	return
}

// matches : [[/1518506476136010007_thumb.jpg 1518506476136010007_thumb.jpg _thumb]]
//var fileRegex = regexp.MustCompile(`/(\w+(_[[:alpha:]]+)\..+)`)
var fileRegex = regexp.MustCompile(`/([0-9]+(_[[:alnum:]]+)?\.\w+)`)

func (r *RowReq) extractParams() {
	parts := fileRegex.FindStringSubmatch(r.FullUrlPath)
	//fmt.Println(parts)
	if len(parts) < 2 {
		r.Err = ErrBadReq
	}
	if len(parts) >= 2 {
		r.FileName = parts[1]
		r.RequestedImageSize = 0
		r.IsThumb = false
	}
	if len(parts) == 3 {
		if parts[2] == "_thumb" {
			r.IsThumb = true
		} else {
			if len(parts[2]) > 1 {
				r.RequestedImageSize = helper.StrToInt(parts[2][1:], 0)
			}
		}
	}
	sep := strings.Split(r.FileName, ".")
	if len(sep) == 2 {
		r.FileDataStoreId = helper.StrToInt(strings.Split(sep[0], "_")[0], 0)
		r.FileExtensionWithoutDot = sep[1]
	} else {
		r.Err = ErrBadReq
		return
	}

	return
}

func (r *RowReq) createRowOutCacheDir() {
	os.MkdirAll(r.RowCacheOutRelativePathDir, os.ModeDir)
}

/*
&file_service_old.RowReq{
    FullUrlPath:                  "/post_file/1518506476136010007_180.jpg",
    FileCategory:              file_service_old.FileCategory{UrlPath:"post_file", cachePath:"post_file", setterToStore:func(file_service_old.Row) {...}, getterOfStore:func(int) (*file_service_old.Row, error) {...}},
    FileName:                  "1518506476136010007_180.jpg",
    FileDataStoreId:           1518506476136010007,
    FileExtensionWithoutDot:   "jpg",
    FileFormat:                0,
    RequestedImageSize:        180,
    IsThumb:                   false,
    RowCacheOutRelativePathDir:    "D:\\sun\\/v1/post_file/15/18/50/64/",
    RowCacheOutRelativePath:       "D:\\sun\\/v1/post_file/15/18/50/64/1518506476136010007.jpg",
    RowCacheOutRelativePathSized:  "D:\\sun\\/v1/post_file/15/18/50/64/1518506476136010007_180.jpg",
    RowCacheOutRelativePathThumb:  "D:\\sun\\/v1/post_file/15/18/50/64/1518506476136010007_thumb.jpg",
    IsOrginalLocalDiskCacheAvailable: false,
    CacheFullModuleDirectory:  "D:\\sun\\/v1/post_file/",
    Err:                       nil,
}
*/
