package file_common

import "errors"

type fileFormat int

const (
	IMAGE = fileFormat(1)
	VIDEO = fileFormat(2)
	AUDIO = fileFormat(3)
	GIF   = fileFormat(4)
)


var ErrBadReq = errors.New("bad request")
var ErrWrongFileName = errors.New("wrong file name")
var ErrFileNameTooShort = errors.New("file name too short")
var ErrResizeNotAllowed = errors.New("image resize not allowed")
