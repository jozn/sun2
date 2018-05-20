package xc

import "strings"

type FileRef struct {
	Extension      string // extension  regular
	FileRefId      int    // file_ref_id  partition_key
	FirstFileRefId int    // first_file_ref_id  regular
	Height         int    // height  regular
	Length         int    // length  regular
	Md5Key         string // md5_key  regular
	MimeType       string // mime_type  regular
	Name           string // name  regular
	StorageUserId  int    // storage_user_id  regular
	StoreType      string // store_type  regular
	UserId         int    // user_id  regular
	Width          int    // width  regular

	_exists, _deleted bool
}

/*
:= &xc.FileRef {
	Extension: "",
	FileRefId: 0,
	FirstFileRefId: 0,
	Height: 0,
	Length: 0,
	Md5Key: "",
	MimeType: "",
	Name: "",
	StorageUserId: 0,
	StoreType: "",
	UserId: 0,
	Width: 0,
*/

type FileStorage struct {
	Extension      string // extension  regular
	FirstFileRefId int    // first_file_ref_id  regular
	Height         int    // height  regular
	Length         int    // length  regular
	Md5Key         string // md5_key  partition_key
	MimeType       string // mime_type  regular
	UserId         int    // user_id  regular
	Width          int    // width  regular
	Zdata          []byte // zdata  regular
	ZdataThumb     []byte // zdata_thumb  regular

	_exists, _deleted bool
}

/*
:= &xc.FileStorage {
	Extension: "",
	FirstFileRefId: 0,
	Height: 0,
	Length: 0,
	Md5Key: "",
	MimeType: "",
	UserId: 0,
	Width: 0,
	Zdata: []byte{},
	ZdataThumb: []byte{},
*/

type StorageGc struct {
	FileRefId int    // file_ref_id  clustering
	Md5Key    string // md5_key  partition_key
	Shard     string // shard  regular
	StoreType string // store_type  regular

	_exists, _deleted bool
}

/*
:= &xc.StorageGc {
	FileRefId: 0,
	Md5Key: "",
	Shard: "",
	StoreType: "",
*/

type StorageShadow struct {
	MimeType      string // mime_type  regular
	ShadowMd5Key  string // shadow_md5_key  partition_key
	StorageMd5Key string // storage_md5_key  regular

	_exists, _deleted bool
}

/*
:= &xc.StorageShadow {
	MimeType: "",
	ShadowMd5Key: "",
	StorageMd5Key: "",
*/

type StorageUsage struct {
	FileRefId int    // file_ref_id  clustering
	Md5Key    string // md5_key  partition_key
	StoreType string // store_type  regular

	_exists, _deleted bool
}

/*
:= &xc.StorageUsage {
	FileRefId: 0,
	Md5Key: "",
	StoreType: "",
*/

// logs tables
type LogTableCql struct {
	FileRef       bool
	FileStorage   bool
	StorageGc     bool
	StorageShadow bool
	StorageUsage  bool
}

var LogTableCqlReq = LogTableCql{
	FileRef:       true,
	FileStorage:   true,
	StorageGc:     true,
	StorageShadow: true,
	StorageUsage:  true,
}

//////////////// Constants //////////////////

type whereClause struct {
	condition string
	args      []interface{}
}

func whereClusesToSql(wheres []whereClause, whereSep string) (string, []interface{}) {
	var wheresArr []string
	for _, w := range wheres {
		wheresArr = append(wheresArr, w.condition)
	}
	wheresStr := strings.Join(wheresArr, whereSep)

	var args []interface{}
	for _, w := range wheres {
		args = append(args, w.args...)
	}
	return wheresStr, args
}

//////////////// End of Constants ///////////////
