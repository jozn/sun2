package file_service

import (
	"errors"
	"net/http"
)

func Run() {
	for _, cat := range httpCategories {
		http.Handle("/"+cat.urlPath+"/", cat)
	}
	//http.Handle("/getFile/", &fileCategory{})
	//http.Handle("/", &fileCategory{})
}

type fileFormat int

const (
	IMAGE = fileFormat(1)
	VIDEO = fileFormat(2)
	AUDIO = fileFormat(3)
	GIF   = fileFormat(4)
)

var errBadReq = errors.New("bad request")
var errWrongFileName = errors.New("wrong file name")
var errFileNameTooShort = errors.New("file name too short")
