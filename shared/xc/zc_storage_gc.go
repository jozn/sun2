package xc

import (
	"errors"
	"ms/sun/shared/helper"
	"strconv"
	"strings"

	"github.com/gocql/gocql"
)

////////////////////////////////////////// Query seletor updater and deleter /////////////////////////

func (a *StorageGc) Exists() bool {
	return a._exists
}

func (a *StorageGc) Deleted() bool {
	return a._deleted
}

type __StorageGc_Selector struct {
	wheres      []whereClause
	selectCol   []string
	orderBy     []string //" order by id desc //for ints
	limit       int
	allowFilter bool
}

type __StorageGc_Updater struct {
	wheres  []whereClause
	updates map[string]interface{}
}

type __StorageGc_Deleter struct {
	wheres    []whereClause
	deleteCol []string
}

//////////////////// just Selector
func (u *__StorageGc_Selector) Limit(limit int) *__StorageGc_Selector {
	u.limit = limit
	return u
}

func (u *__StorageGc_Selector) AllowFiltering() *__StorageGc_Selector {
	u.allowFilter = true
	return u
}

func NewStorageGc_Selector() *__StorageGc_Selector {
	u := __StorageGc_Selector{}
	return &u
}

func NewStorageGc_Updater() *__StorageGc_Updater {
	u := __StorageGc_Updater{}
	u.updates = make(map[string]interface{})
	return &u
}

func NewStorageGc_Deleter() *__StorageGc_Deleter {
	u := __StorageGc_Deleter{}
	return &u
}

//each select columns

func (u *__StorageGc_Selector) Select_FileRefId() *__StorageGc_Selector {
	u.selectCol = append(u.selectCol, "file_ref_id")
	return u
}

//each column orders //just ints
func (u *__StorageGc_Selector) OrderBy_FileRefId_Desc() *__StorageGc_Selector {
	u.orderBy = append(u.orderBy, " file_ref_id DESC")
	return u
}

func (u *__StorageGc_Selector) OrderBy_FileRefId_Asc() *__StorageGc_Selector {
	u.orderBy = append(u.orderBy, " file_ref_id ASC")
	return u
}

func (u *__StorageGc_Selector) Select_Md5Key() *__StorageGc_Selector {
	u.selectCol = append(u.selectCol, "md5_key")
	return u
}

//each column orders //just ints
func (u *__StorageGc_Selector) OrderBy_Md5Key_Desc() *__StorageGc_Selector {
	u.orderBy = append(u.orderBy, " md5_key DESC")
	return u
}

func (u *__StorageGc_Selector) OrderBy_Md5Key_Asc() *__StorageGc_Selector {
	u.orderBy = append(u.orderBy, " md5_key ASC")
	return u
}

func (u *__StorageGc_Selector) Select_Shard() *__StorageGc_Selector {
	u.selectCol = append(u.selectCol, "shard")
	return u
}

//each column orders //just ints
func (u *__StorageGc_Selector) OrderBy_Shard_Desc() *__StorageGc_Selector {
	u.orderBy = append(u.orderBy, " shard DESC")
	return u
}

func (u *__StorageGc_Selector) OrderBy_Shard_Asc() *__StorageGc_Selector {
	u.orderBy = append(u.orderBy, " shard ASC")
	return u
}

func (u *__StorageGc_Selector) Select_StoreType() *__StorageGc_Selector {
	u.selectCol = append(u.selectCol, "store_type")
	return u
}

//each column orders //just ints
func (u *__StorageGc_Selector) OrderBy_StoreType_Desc() *__StorageGc_Selector {
	u.orderBy = append(u.orderBy, " store_type DESC")
	return u
}

func (u *__StorageGc_Selector) OrderBy_StoreType_Asc() *__StorageGc_Selector {
	u.orderBy = append(u.orderBy, " store_type ASC")
	return u
}

//////////////////// just Deleter
//each column delete

func (u *__StorageGc_Deleter) Delete_FileRefId() *__StorageGc_Deleter {
	u.deleteCol = append(u.deleteCol, "file_ref_id")
	return u
}

func (u *__StorageGc_Deleter) Delete_Md5Key() *__StorageGc_Deleter {
	u.deleteCol = append(u.deleteCol, "md5_key")
	return u
}

func (u *__StorageGc_Deleter) Delete_Shard() *__StorageGc_Deleter {
	u.deleteCol = append(u.deleteCol, "shard")
	return u
}

func (u *__StorageGc_Deleter) Delete_StoreType() *__StorageGc_Deleter {
	u.deleteCol = append(u.deleteCol, "store_type")
	return u
}

//////////////////// End of just Deleter

//////////////////// just Updater
//each column delete

func (u *__StorageGc_Updater) FileRefId(newVal int) *__StorageGc_Updater {
	u.updates["file_ref_id = ? "] = newVal
	return u
}

func (u *__StorageGc_Updater) Md5Key(newVal string) *__StorageGc_Updater {
	u.updates["md5_key = ? "] = newVal
	return u
}

func (u *__StorageGc_Updater) Shard(newVal string) *__StorageGc_Updater {
	u.updates["shard = ? "] = newVal
	return u
}

func (u *__StorageGc_Updater) StoreType(newVal string) *__StorageGc_Updater {
	u.updates["store_type = ? "] = newVal
	return u
}

//////////////////// End just Updater

//{_Eq  =  FileRefId_Eq}

func (d *__StorageGc_Deleter) FileRefId_Eq(val int) *__StorageGc_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " file_ref_id = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LT  <  FileRefId_LT}

func (d *__StorageGc_Deleter) FileRefId_LT(val int) *__StorageGc_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " file_ref_id < ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LE  <=  FileRefId_LE}

func (d *__StorageGc_Deleter) FileRefId_LE(val int) *__StorageGc_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " file_ref_id <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GT  >  FileRefId_GT}

func (d *__StorageGc_Deleter) FileRefId_GT(val int) *__StorageGc_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " file_ref_id > ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GE  >=  FileRefId_GE}

func (d *__StorageGc_Deleter) FileRefId_GE(val int) *__StorageGc_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " file_ref_id >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq And = And And_FileRefId_Eq}

func (d *__StorageGc_Deleter) And_FileRefId_Eq(val int) *__StorageGc_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And file_ref_id = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LT And < And And_FileRefId_LT}

func (d *__StorageGc_Deleter) And_FileRefId_LT(val int) *__StorageGc_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And file_ref_id < ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LE And <= And And_FileRefId_LE}

func (d *__StorageGc_Deleter) And_FileRefId_LE(val int) *__StorageGc_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And file_ref_id <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GT And > And And_FileRefId_GT}

func (d *__StorageGc_Deleter) And_FileRefId_GT(val int) *__StorageGc_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And file_ref_id > ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GE And >= And And_FileRefId_GE}

func (d *__StorageGc_Deleter) And_FileRefId_GE(val int) *__StorageGc_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And file_ref_id >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq Or = Or Or_FileRefId_Eq}

func (d *__StorageGc_Deleter) Or_FileRefId_Eq(val int) *__StorageGc_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or file_ref_id = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LT Or < Or Or_FileRefId_LT}

func (d *__StorageGc_Deleter) Or_FileRefId_LT(val int) *__StorageGc_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or file_ref_id < ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LE Or <= Or Or_FileRefId_LE}

func (d *__StorageGc_Deleter) Or_FileRefId_LE(val int) *__StorageGc_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or file_ref_id <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GT Or > Or Or_FileRefId_GT}

func (d *__StorageGc_Deleter) Or_FileRefId_GT(val int) *__StorageGc_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or file_ref_id > ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GE Or >= Or Or_FileRefId_GE}

func (d *__StorageGc_Deleter) Or_FileRefId_GE(val int) *__StorageGc_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or file_ref_id >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq  =  Md5Key_Eq}

func (d *__StorageGc_Deleter) Md5Key_Eq(val string) *__StorageGc_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " md5_key = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq And = And And_Md5Key_Eq}

func (d *__StorageGc_Deleter) And_Md5Key_Eq(val string) *__StorageGc_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And md5_key = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq Or = Or Or_Md5Key_Eq}

func (d *__StorageGc_Deleter) Or_Md5Key_Eq(val string) *__StorageGc_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or md5_key = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_FILTERING  =  Shard_Eq_FILTERING}

func (d *__StorageGc_Deleter) Shard_Eq_FILTERING(val string) *__StorageGc_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " shard = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_FILTERING And = And And_Shard_Eq_FILTERING}

func (d *__StorageGc_Deleter) And_Shard_Eq_FILTERING(val string) *__StorageGc_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And shard = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_FILTERING Or = Or Or_Shard_Eq_FILTERING}

func (d *__StorageGc_Deleter) Or_Shard_Eq_FILTERING(val string) *__StorageGc_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or shard = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_FILTERING  =  StoreType_Eq_FILTERING}

func (d *__StorageGc_Deleter) StoreType_Eq_FILTERING(val string) *__StorageGc_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " store_type = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_FILTERING And = And And_StoreType_Eq_FILTERING}

func (d *__StorageGc_Deleter) And_StoreType_Eq_FILTERING(val string) *__StorageGc_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And store_type = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_FILTERING Or = Or Or_StoreType_Eq_FILTERING}

func (d *__StorageGc_Deleter) Or_StoreType_Eq_FILTERING(val string) *__StorageGc_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or store_type = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq  =  FileRefId_Eq}

func (d *__StorageGc_Updater) FileRefId_Eq(val int) *__StorageGc_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " file_ref_id = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LT  <  FileRefId_LT}

func (d *__StorageGc_Updater) FileRefId_LT(val int) *__StorageGc_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " file_ref_id < ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LE  <=  FileRefId_LE}

func (d *__StorageGc_Updater) FileRefId_LE(val int) *__StorageGc_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " file_ref_id <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GT  >  FileRefId_GT}

func (d *__StorageGc_Updater) FileRefId_GT(val int) *__StorageGc_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " file_ref_id > ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GE  >=  FileRefId_GE}

func (d *__StorageGc_Updater) FileRefId_GE(val int) *__StorageGc_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " file_ref_id >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq And = And And_FileRefId_Eq}

func (d *__StorageGc_Updater) And_FileRefId_Eq(val int) *__StorageGc_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And file_ref_id = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LT And < And And_FileRefId_LT}

func (d *__StorageGc_Updater) And_FileRefId_LT(val int) *__StorageGc_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And file_ref_id < ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LE And <= And And_FileRefId_LE}

func (d *__StorageGc_Updater) And_FileRefId_LE(val int) *__StorageGc_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And file_ref_id <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GT And > And And_FileRefId_GT}

func (d *__StorageGc_Updater) And_FileRefId_GT(val int) *__StorageGc_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And file_ref_id > ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GE And >= And And_FileRefId_GE}

func (d *__StorageGc_Updater) And_FileRefId_GE(val int) *__StorageGc_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And file_ref_id >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq Or = Or Or_FileRefId_Eq}

func (d *__StorageGc_Updater) Or_FileRefId_Eq(val int) *__StorageGc_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or file_ref_id = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LT Or < Or Or_FileRefId_LT}

func (d *__StorageGc_Updater) Or_FileRefId_LT(val int) *__StorageGc_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or file_ref_id < ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LE Or <= Or Or_FileRefId_LE}

func (d *__StorageGc_Updater) Or_FileRefId_LE(val int) *__StorageGc_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or file_ref_id <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GT Or > Or Or_FileRefId_GT}

func (d *__StorageGc_Updater) Or_FileRefId_GT(val int) *__StorageGc_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or file_ref_id > ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GE Or >= Or Or_FileRefId_GE}

func (d *__StorageGc_Updater) Or_FileRefId_GE(val int) *__StorageGc_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or file_ref_id >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq  =  Md5Key_Eq}

func (d *__StorageGc_Updater) Md5Key_Eq(val string) *__StorageGc_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " md5_key = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq And = And And_Md5Key_Eq}

func (d *__StorageGc_Updater) And_Md5Key_Eq(val string) *__StorageGc_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And md5_key = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq Or = Or Or_Md5Key_Eq}

func (d *__StorageGc_Updater) Or_Md5Key_Eq(val string) *__StorageGc_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or md5_key = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_FILTERING  =  Shard_Eq_FILTERING}

func (d *__StorageGc_Updater) Shard_Eq_FILTERING(val string) *__StorageGc_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " shard = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_FILTERING And = And And_Shard_Eq_FILTERING}

func (d *__StorageGc_Updater) And_Shard_Eq_FILTERING(val string) *__StorageGc_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And shard = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_FILTERING Or = Or Or_Shard_Eq_FILTERING}

func (d *__StorageGc_Updater) Or_Shard_Eq_FILTERING(val string) *__StorageGc_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or shard = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_FILTERING  =  StoreType_Eq_FILTERING}

func (d *__StorageGc_Updater) StoreType_Eq_FILTERING(val string) *__StorageGc_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " store_type = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_FILTERING And = And And_StoreType_Eq_FILTERING}

func (d *__StorageGc_Updater) And_StoreType_Eq_FILTERING(val string) *__StorageGc_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And store_type = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_FILTERING Or = Or Or_StoreType_Eq_FILTERING}

func (d *__StorageGc_Updater) Or_StoreType_Eq_FILTERING(val string) *__StorageGc_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or store_type = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq  =  FileRefId_Eq}

func (d *__StorageGc_Selector) FileRefId_Eq(val int) *__StorageGc_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " file_ref_id = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LT  <  FileRefId_LT}

func (d *__StorageGc_Selector) FileRefId_LT(val int) *__StorageGc_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " file_ref_id < ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LE  <=  FileRefId_LE}

func (d *__StorageGc_Selector) FileRefId_LE(val int) *__StorageGc_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " file_ref_id <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GT  >  FileRefId_GT}

func (d *__StorageGc_Selector) FileRefId_GT(val int) *__StorageGc_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " file_ref_id > ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GE  >=  FileRefId_GE}

func (d *__StorageGc_Selector) FileRefId_GE(val int) *__StorageGc_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " file_ref_id >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq And = And And_FileRefId_Eq}

func (d *__StorageGc_Selector) And_FileRefId_Eq(val int) *__StorageGc_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And file_ref_id = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LT And < And And_FileRefId_LT}

func (d *__StorageGc_Selector) And_FileRefId_LT(val int) *__StorageGc_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And file_ref_id < ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LE And <= And And_FileRefId_LE}

func (d *__StorageGc_Selector) And_FileRefId_LE(val int) *__StorageGc_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And file_ref_id <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GT And > And And_FileRefId_GT}

func (d *__StorageGc_Selector) And_FileRefId_GT(val int) *__StorageGc_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And file_ref_id > ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GE And >= And And_FileRefId_GE}

func (d *__StorageGc_Selector) And_FileRefId_GE(val int) *__StorageGc_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And file_ref_id >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq Or = Or Or_FileRefId_Eq}

func (d *__StorageGc_Selector) Or_FileRefId_Eq(val int) *__StorageGc_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or file_ref_id = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LT Or < Or Or_FileRefId_LT}

func (d *__StorageGc_Selector) Or_FileRefId_LT(val int) *__StorageGc_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or file_ref_id < ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LE Or <= Or Or_FileRefId_LE}

func (d *__StorageGc_Selector) Or_FileRefId_LE(val int) *__StorageGc_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or file_ref_id <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GT Or > Or Or_FileRefId_GT}

func (d *__StorageGc_Selector) Or_FileRefId_GT(val int) *__StorageGc_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or file_ref_id > ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GE Or >= Or Or_FileRefId_GE}

func (d *__StorageGc_Selector) Or_FileRefId_GE(val int) *__StorageGc_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or file_ref_id >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq  =  Md5Key_Eq}

func (d *__StorageGc_Selector) Md5Key_Eq(val string) *__StorageGc_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " md5_key = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq And = And And_Md5Key_Eq}

func (d *__StorageGc_Selector) And_Md5Key_Eq(val string) *__StorageGc_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And md5_key = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq Or = Or Or_Md5Key_Eq}

func (d *__StorageGc_Selector) Or_Md5Key_Eq(val string) *__StorageGc_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or md5_key = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_FILTERING  =  Shard_Eq_FILTERING}

func (d *__StorageGc_Selector) Shard_Eq_FILTERING(val string) *__StorageGc_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " shard = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_FILTERING And = And And_Shard_Eq_FILTERING}

func (d *__StorageGc_Selector) And_Shard_Eq_FILTERING(val string) *__StorageGc_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And shard = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_FILTERING Or = Or Or_Shard_Eq_FILTERING}

func (d *__StorageGc_Selector) Or_Shard_Eq_FILTERING(val string) *__StorageGc_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or shard = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_FILTERING  =  StoreType_Eq_FILTERING}

func (d *__StorageGc_Selector) StoreType_Eq_FILTERING(val string) *__StorageGc_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " store_type = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_FILTERING And = And And_StoreType_Eq_FILTERING}

func (d *__StorageGc_Selector) And_StoreType_Eq_FILTERING(val string) *__StorageGc_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And store_type = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_FILTERING Or = Or Or_StoreType_Eq_FILTERING}

func (d *__StorageGc_Selector) Or_StoreType_Eq_FILTERING(val string) *__StorageGc_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or store_type = ? "
	d.wheres = append(d.wheres, w)

	return d
}

///////////////////////////////////////// ins for all //////////////////

func (d *__StorageGc_Deleter) FileRefId_In(val ...int) *__StorageGc_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = " file_ref_id IN (" + helper.DbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__StorageGc_Deleter) And_FileRefId_In(val ...int) *__StorageGc_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = "And file_ref_id IN (" + helper.DbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__StorageGc_Deleter) Or_FileRefId_In(val ...int) *__StorageGc_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = "Or file_ref_id IN (" + helper.DbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__StorageGc_Deleter) Md5Key_In(val ...string) *__StorageGc_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = " md5_key IN (" + helper.DbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__StorageGc_Deleter) And_Md5Key_In(val ...string) *__StorageGc_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = "And md5_key IN (" + helper.DbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__StorageGc_Deleter) Or_Md5Key_In(val ...string) *__StorageGc_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = "Or md5_key IN (" + helper.DbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__StorageGc_Deleter) Shard_In_FILTERING(val ...string) *__StorageGc_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = " shard IN (" + helper.DbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__StorageGc_Deleter) And_Shard_In_FILTERING(val ...string) *__StorageGc_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = "And shard IN (" + helper.DbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__StorageGc_Deleter) Or_Shard_In_FILTERING(val ...string) *__StorageGc_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = "Or shard IN (" + helper.DbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__StorageGc_Deleter) StoreType_In_FILTERING(val ...string) *__StorageGc_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = " store_type IN (" + helper.DbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__StorageGc_Deleter) And_StoreType_In_FILTERING(val ...string) *__StorageGc_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = "And store_type IN (" + helper.DbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__StorageGc_Deleter) Or_StoreType_In_FILTERING(val ...string) *__StorageGc_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = "Or store_type IN (" + helper.DbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__StorageGc_Updater) FileRefId_In(val ...int) *__StorageGc_Updater {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = " file_ref_id IN (" + helper.DbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__StorageGc_Updater) And_FileRefId_In(val ...int) *__StorageGc_Updater {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = "And file_ref_id IN (" + helper.DbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__StorageGc_Updater) Or_FileRefId_In(val ...int) *__StorageGc_Updater {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = "Or file_ref_id IN (" + helper.DbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__StorageGc_Updater) Md5Key_In(val ...string) *__StorageGc_Updater {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = " md5_key IN (" + helper.DbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__StorageGc_Updater) And_Md5Key_In(val ...string) *__StorageGc_Updater {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = "And md5_key IN (" + helper.DbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__StorageGc_Updater) Or_Md5Key_In(val ...string) *__StorageGc_Updater {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = "Or md5_key IN (" + helper.DbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__StorageGc_Updater) Shard_In_FILTERING(val ...string) *__StorageGc_Updater {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = " shard IN (" + helper.DbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__StorageGc_Updater) And_Shard_In_FILTERING(val ...string) *__StorageGc_Updater {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = "And shard IN (" + helper.DbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__StorageGc_Updater) Or_Shard_In_FILTERING(val ...string) *__StorageGc_Updater {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = "Or shard IN (" + helper.DbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__StorageGc_Updater) StoreType_In_FILTERING(val ...string) *__StorageGc_Updater {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = " store_type IN (" + helper.DbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__StorageGc_Updater) And_StoreType_In_FILTERING(val ...string) *__StorageGc_Updater {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = "And store_type IN (" + helper.DbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__StorageGc_Updater) Or_StoreType_In_FILTERING(val ...string) *__StorageGc_Updater {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = "Or store_type IN (" + helper.DbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__StorageGc_Selector) FileRefId_In(val ...int) *__StorageGc_Selector {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = " file_ref_id IN (" + helper.DbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__StorageGc_Selector) And_FileRefId_In(val ...int) *__StorageGc_Selector {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = "And file_ref_id IN (" + helper.DbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__StorageGc_Selector) Or_FileRefId_In(val ...int) *__StorageGc_Selector {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = "Or file_ref_id IN (" + helper.DbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__StorageGc_Selector) Md5Key_In(val ...string) *__StorageGc_Selector {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = " md5_key IN (" + helper.DbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__StorageGc_Selector) And_Md5Key_In(val ...string) *__StorageGc_Selector {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = "And md5_key IN (" + helper.DbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__StorageGc_Selector) Or_Md5Key_In(val ...string) *__StorageGc_Selector {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = "Or md5_key IN (" + helper.DbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__StorageGc_Selector) Shard_In_FILTERING(val ...string) *__StorageGc_Selector {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = " shard IN (" + helper.DbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__StorageGc_Selector) And_Shard_In_FILTERING(val ...string) *__StorageGc_Selector {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = "And shard IN (" + helper.DbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__StorageGc_Selector) Or_Shard_In_FILTERING(val ...string) *__StorageGc_Selector {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = "Or shard IN (" + helper.DbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__StorageGc_Selector) StoreType_In_FILTERING(val ...string) *__StorageGc_Selector {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = " store_type IN (" + helper.DbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__StorageGc_Selector) And_StoreType_In_FILTERING(val ...string) *__StorageGc_Selector {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = "And store_type IN (" + helper.DbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__StorageGc_Selector) Or_StoreType_In_FILTERING(val ...string) *__StorageGc_Selector {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = "Or store_type IN (" + helper.DbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

/////////////////////////////////////// End of Ins //////////////////////
///////////////////////////// start of where cluases

/////////////////////////////////////// Start of select //////////////////
func (u *__StorageGc_Selector) _toSql() (string, []interface{}) {

	sqlWheres, whereArgs := whereClusesToSql(u.wheres, "")
	selectCols := "*"
	if len(u.selectCol) > 0 {
		selectCols = strings.Join(u.selectCol, ", ")
	}
	sqlstr := "SELECT " + selectCols + " FROM sunc_file.storage_gc"

	if len(strings.Trim(sqlWheres, " ")) > 0 { //2 for safty
		sqlstr += " WHERE " + sqlWheres
	}

	if len(u.orderBy) > 0 {
		orders := strings.Join(u.orderBy, ", ")
		sqlstr += " ORDER BY " + orders
	}

	if u.limit != 0 {
		sqlstr += " LIMIT " + strconv.Itoa(u.limit)
	}
	if u.allowFilter {
		sqlstr += "  ALLOW FILTERING"
	}

	return sqlstr, whereArgs
}

func (u *__StorageGc_Selector) GetRow(session *gocql.Session) (*StorageGc, error) {
	var err error

	u.limit = 1
	sqlstr, whereArgs := u._toSql()

	if LogTableCqlReq.StorageGc {
		helper.XCLog(sqlstr, whereArgs)
	}

	query := session.Query(sqlstr, whereArgs...)
	var row *StorageGc
	//by Sqlx
	// err = gocqlx.Get(row ,query)
	rows, err := StorageGc_Iter(query.Iter(), 1)
	if err != nil {
		if LogTableCqlReq.StorageGc {
			helper.XCLogErr(err)
		}
		return nil, err
	}
	if len(rows) == 0 {
		return nil, errors.New("empty rows")
	} else {
		row = rows[0]
	}

	row._exists = true

	//OnStorageGc_LoadOne(row)

	return row, nil
}

func (u *__StorageGc_Selector) GetRows(session *gocql.Session) ([]*StorageGc, error) {
	var err error

	sqlstr, whereArgs := u._toSql()

	if LogTableCqlReq.StorageGc {
		helper.XCLog(sqlstr, whereArgs)
	}

	query := session.Query(sqlstr, whereArgs...)

	rows, err := StorageGc_Iter(query.Iter(), -1)
	if err != nil {
		if LogTableCqlReq.StorageGc {
			helper.XCLogErr(err)
		}
		return rows, err
	}

	for i := 0; i < len(rows); i++ {
		rows[i]._exists = true
	}

	// OnStorageGc_LoadMany(rows)

	return rows, nil
}

func (u *__StorageGc_Updater) Update(session *gocql.Session) error {
	var err error

	var updateArgs []interface{}
	var sqlUpdateArr []string
	for up, newVal := range u.updates {
		sqlUpdateArr = append(sqlUpdateArr, up)
		updateArgs = append(updateArgs, newVal)
	}
	sqlUpdate := strings.Join(sqlUpdateArr, ",")

	sqlWheres, whereArgs := whereClusesToSql(u.wheres, "")

	var allArgs []interface{}
	allArgs = append(allArgs, updateArgs...)
	allArgs = append(allArgs, whereArgs...)

	sqlstr := `UPDATE sunc_file.storage_gc SET ` + sqlUpdate

	if len(strings.Trim(sqlWheres, " ")) > 0 {
		sqlstr += " WHERE " + sqlWheres
	}
	if LogTableCqlReq.StorageGc {
		helper.XCLog(sqlstr, allArgs)
	}
	err = session.Query(sqlstr, allArgs...).Exec()
	if err != nil {
		helper.XCLogErr(err)
		return err
	}

	return nil
}

func (d *__StorageGc_Deleter) Delete(session *gocql.Session) error {
	var err error

	var wheresArr []string
	var args []interface{}

	var delCols string
	if len(d.deleteCol) > 0 {
		delCols = strings.Join(d.deleteCol, ",")
	}

	for _, w := range d.wheres {
		wheresArr = append(wheresArr, w.condition)
		args = append(args, w.args...)
	}
	wheresStr := strings.Join(wheresArr, "")

	sqlstr := "DELETE" + delCols + " FROM sunc_file.storage_gc WHERE " + wheresStr

	// run query
	if LogTableCqlReq.StorageGc {
		helper.XCLog(sqlstr, args)
	}
	err = session.Query(sqlstr, args...).Exec()
	if err != nil {
		helper.XCLogErr(err)
		return err
	}

	return nil
}

/*
func MassInsert_StorageGc(rows []*StorageGc, session *gocql.Session) error {
    if len(rows) == 0 {
        return errors.New("rows slice should not be empty - inserted nothing")
    }
    var err error
    ln := len(rows)
    insVals := helper.SqlManyDollars( 4 ,len(rows),true)

    sqlstr := "INSERT INTO sunc_file.storage_gc (" +
       " file_ref_id,md5_key,shard,store_type " +
        ") VALUES " + insVals

    // run query
    vals := make([]interface{}, 0, ln*5) //5 fields

    for _, row := range rows {
    		vals = append(vals, row.FileRefId)
    		vals = append(vals, row.Md5Key)
    		vals = append(vals, row.Shard)
    		vals = append(vals, row.StoreType)
    }

    if LogTableCqlReq.StorageGc {
        helper.XCLog(" MassInsert len = ", ln, sqlstr ,vals)
    }
    err = session.Query(sqlstr, vals...).Exec()
    if err != nil {
        helper.XCLogErr(err)
        return err
    }

    return nil
}
*/
///////

func (r *StorageGc) Save(session *gocql.Session) error {
	var cols []string
	var q []string
	var vals []interface{}

	if r.FileRefId != 0 {
		cols = append(cols, "file_ref_id")
		q = append(q, "?")
		vals = append(vals, r.FileRefId)
	}

	if r.Md5Key != "" {
		cols = append(cols, "md5_key")
		q = append(q, "?")
		vals = append(vals, r.Md5Key)
	}

	if r.Shard != "" {
		cols = append(cols, "shard")
		q = append(q, "?")
		vals = append(vals, r.Shard)
	}

	if r.StoreType != "" {
		cols = append(cols, "store_type")
		q = append(q, "?")
		vals = append(vals, r.StoreType)
	}

	if len(cols) == 0 {
		return errors.New("can not insert empty row.")
	}

	colOut := strings.Join(cols, ",")
	qOut := strings.Join(q, ",")
	cql := "insert into sunc_file.storage_gc (" + colOut + ") values (" + qOut + ") "

	if LogTableCqlReq.StorageGc {
		helper.XCLog(cql, vals)
	}
	err := session.Query(cql, vals...).Exec()
	if err != nil {
		if LogTableCqlReq.StorageGc {
			helper.XCLogErr(err)
		}
	}
	r._exists = true
	return err
}

func (r *StorageGc) Delete(session *gocql.Session) error {
	var err error
	del := NewStorageGc_Deleter()

	del.Md5Key_Eq(r.Md5Key)

	del.And_FileRefId_Eq(r.FileRefId)

	err = del.Delete(session)
	if err != nil {
		return err
	}
	r._exists = false
	return nil
}

func StorageGc_Iter(iter *gocql.Iter, limit int) ([]*StorageGc, error) {
	var rows []*StorageGc
	if limit < 1 {
		limit = 1e6
	}

	for i := 0; i < limit; i++ {
		m := make(map[string]interface{}, 10)
		row := &StorageGc{}
		if iter.MapScan(m) {

			if val, ok := m["file_ref_id"]; ok {
				row.FileRefId = int(val.(int64))
				//row.FileRefId = val.(int)
			}

			if val, ok := m["md5_key"]; ok {
				row.Md5Key = string(val.(string))
				//row.Md5Key = val.(string)
			}

			if val, ok := m["shard"]; ok {
				row.Shard = string(val.(string))
				//row.Shard = val.(string)
			}

			if val, ok := m["store_type"]; ok {
				row.StoreType = string(val.(string))
				//row.StoreType = val.(string)
			}

			rows = append(rows, row)
		} else {
			break
		}
	}
	err := iter.Close()

	return rows, err
}
