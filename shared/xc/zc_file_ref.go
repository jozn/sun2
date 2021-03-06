package xc

import (
	"errors"
	"ms/sun/shared/helper"
	"strconv"
	"strings"

	"github.com/gocql/gocql"
)

////////////////////////////////////////// Query seletor updater and deleter /////////////////////////

func (a *FileRef) Exists() bool {
	return a._exists
}

func (a *FileRef) Deleted() bool {
	return a._deleted
}

type __FileRef_Selector struct {
	wheres      []whereClause
	selectCol   []string
	orderBy     []string //" order by id desc //for ints
	limit       int
	allowFilter bool
}

type __FileRef_Updater struct {
	wheres  []whereClause
	updates map[string]interface{}
}

type __FileRef_Deleter struct {
	wheres    []whereClause
	deleteCol []string
}

//////////////////// just Selector
func (u *__FileRef_Selector) Limit(limit int) *__FileRef_Selector {
	u.limit = limit
	return u
}

func (u *__FileRef_Selector) AllowFiltering() *__FileRef_Selector {
	u.allowFilter = true
	return u
}

func NewFileRef_Selector() *__FileRef_Selector {
	u := __FileRef_Selector{}
	return &u
}

func NewFileRef_Updater() *__FileRef_Updater {
	u := __FileRef_Updater{}
	u.updates = make(map[string]interface{})
	return &u
}

func NewFileRef_Deleter() *__FileRef_Deleter {
	u := __FileRef_Deleter{}
	return &u
}

//each select columns

func (u *__FileRef_Selector) Select_Extension() *__FileRef_Selector {
	u.selectCol = append(u.selectCol, "extension")
	return u
}

//each column orders //just ints
func (u *__FileRef_Selector) OrderBy_Extension_Desc() *__FileRef_Selector {
	u.orderBy = append(u.orderBy, " extension DESC")
	return u
}

func (u *__FileRef_Selector) OrderBy_Extension_Asc() *__FileRef_Selector {
	u.orderBy = append(u.orderBy, " extension ASC")
	return u
}

func (u *__FileRef_Selector) Select_FileRefId() *__FileRef_Selector {
	u.selectCol = append(u.selectCol, "file_ref_id")
	return u
}

//each column orders //just ints
func (u *__FileRef_Selector) OrderBy_FileRefId_Desc() *__FileRef_Selector {
	u.orderBy = append(u.orderBy, " file_ref_id DESC")
	return u
}

func (u *__FileRef_Selector) OrderBy_FileRefId_Asc() *__FileRef_Selector {
	u.orderBy = append(u.orderBy, " file_ref_id ASC")
	return u
}

func (u *__FileRef_Selector) Select_FirstFileRefId() *__FileRef_Selector {
	u.selectCol = append(u.selectCol, "first_file_ref_id")
	return u
}

//each column orders //just ints
func (u *__FileRef_Selector) OrderBy_FirstFileRefId_Desc() *__FileRef_Selector {
	u.orderBy = append(u.orderBy, " first_file_ref_id DESC")
	return u
}

func (u *__FileRef_Selector) OrderBy_FirstFileRefId_Asc() *__FileRef_Selector {
	u.orderBy = append(u.orderBy, " first_file_ref_id ASC")
	return u
}

func (u *__FileRef_Selector) Select_Height() *__FileRef_Selector {
	u.selectCol = append(u.selectCol, "height")
	return u
}

//each column orders //just ints
func (u *__FileRef_Selector) OrderBy_Height_Desc() *__FileRef_Selector {
	u.orderBy = append(u.orderBy, " height DESC")
	return u
}

func (u *__FileRef_Selector) OrderBy_Height_Asc() *__FileRef_Selector {
	u.orderBy = append(u.orderBy, " height ASC")
	return u
}

func (u *__FileRef_Selector) Select_Length() *__FileRef_Selector {
	u.selectCol = append(u.selectCol, "length")
	return u
}

//each column orders //just ints
func (u *__FileRef_Selector) OrderBy_Length_Desc() *__FileRef_Selector {
	u.orderBy = append(u.orderBy, " length DESC")
	return u
}

func (u *__FileRef_Selector) OrderBy_Length_Asc() *__FileRef_Selector {
	u.orderBy = append(u.orderBy, " length ASC")
	return u
}

func (u *__FileRef_Selector) Select_Md5Key() *__FileRef_Selector {
	u.selectCol = append(u.selectCol, "md5_key")
	return u
}

//each column orders //just ints
func (u *__FileRef_Selector) OrderBy_Md5Key_Desc() *__FileRef_Selector {
	u.orderBy = append(u.orderBy, " md5_key DESC")
	return u
}

func (u *__FileRef_Selector) OrderBy_Md5Key_Asc() *__FileRef_Selector {
	u.orderBy = append(u.orderBy, " md5_key ASC")
	return u
}

func (u *__FileRef_Selector) Select_MimeType() *__FileRef_Selector {
	u.selectCol = append(u.selectCol, "mime_type")
	return u
}

//each column orders //just ints
func (u *__FileRef_Selector) OrderBy_MimeType_Desc() *__FileRef_Selector {
	u.orderBy = append(u.orderBy, " mime_type DESC")
	return u
}

func (u *__FileRef_Selector) OrderBy_MimeType_Asc() *__FileRef_Selector {
	u.orderBy = append(u.orderBy, " mime_type ASC")
	return u
}

func (u *__FileRef_Selector) Select_Name() *__FileRef_Selector {
	u.selectCol = append(u.selectCol, "name")
	return u
}

//each column orders //just ints
func (u *__FileRef_Selector) OrderBy_Name_Desc() *__FileRef_Selector {
	u.orderBy = append(u.orderBy, " name DESC")
	return u
}

func (u *__FileRef_Selector) OrderBy_Name_Asc() *__FileRef_Selector {
	u.orderBy = append(u.orderBy, " name ASC")
	return u
}

func (u *__FileRef_Selector) Select_StorageUserId() *__FileRef_Selector {
	u.selectCol = append(u.selectCol, "storage_user_id")
	return u
}

//each column orders //just ints
func (u *__FileRef_Selector) OrderBy_StorageUserId_Desc() *__FileRef_Selector {
	u.orderBy = append(u.orderBy, " storage_user_id DESC")
	return u
}

func (u *__FileRef_Selector) OrderBy_StorageUserId_Asc() *__FileRef_Selector {
	u.orderBy = append(u.orderBy, " storage_user_id ASC")
	return u
}

func (u *__FileRef_Selector) Select_StoreType() *__FileRef_Selector {
	u.selectCol = append(u.selectCol, "store_type")
	return u
}

//each column orders //just ints
func (u *__FileRef_Selector) OrderBy_StoreType_Desc() *__FileRef_Selector {
	u.orderBy = append(u.orderBy, " store_type DESC")
	return u
}

func (u *__FileRef_Selector) OrderBy_StoreType_Asc() *__FileRef_Selector {
	u.orderBy = append(u.orderBy, " store_type ASC")
	return u
}

func (u *__FileRef_Selector) Select_UserId() *__FileRef_Selector {
	u.selectCol = append(u.selectCol, "user_id")
	return u
}

//each column orders //just ints
func (u *__FileRef_Selector) OrderBy_UserId_Desc() *__FileRef_Selector {
	u.orderBy = append(u.orderBy, " user_id DESC")
	return u
}

func (u *__FileRef_Selector) OrderBy_UserId_Asc() *__FileRef_Selector {
	u.orderBy = append(u.orderBy, " user_id ASC")
	return u
}

func (u *__FileRef_Selector) Select_Width() *__FileRef_Selector {
	u.selectCol = append(u.selectCol, "width")
	return u
}

//each column orders //just ints
func (u *__FileRef_Selector) OrderBy_Width_Desc() *__FileRef_Selector {
	u.orderBy = append(u.orderBy, " width DESC")
	return u
}

func (u *__FileRef_Selector) OrderBy_Width_Asc() *__FileRef_Selector {
	u.orderBy = append(u.orderBy, " width ASC")
	return u
}

//////////////////// just Deleter
//each column delete

func (u *__FileRef_Deleter) Delete_Extension() *__FileRef_Deleter {
	u.deleteCol = append(u.deleteCol, "extension")
	return u
}

func (u *__FileRef_Deleter) Delete_FileRefId() *__FileRef_Deleter {
	u.deleteCol = append(u.deleteCol, "file_ref_id")
	return u
}

func (u *__FileRef_Deleter) Delete_FirstFileRefId() *__FileRef_Deleter {
	u.deleteCol = append(u.deleteCol, "first_file_ref_id")
	return u
}

func (u *__FileRef_Deleter) Delete_Height() *__FileRef_Deleter {
	u.deleteCol = append(u.deleteCol, "height")
	return u
}

func (u *__FileRef_Deleter) Delete_Length() *__FileRef_Deleter {
	u.deleteCol = append(u.deleteCol, "length")
	return u
}

func (u *__FileRef_Deleter) Delete_Md5Key() *__FileRef_Deleter {
	u.deleteCol = append(u.deleteCol, "md5_key")
	return u
}

func (u *__FileRef_Deleter) Delete_MimeType() *__FileRef_Deleter {
	u.deleteCol = append(u.deleteCol, "mime_type")
	return u
}

func (u *__FileRef_Deleter) Delete_Name() *__FileRef_Deleter {
	u.deleteCol = append(u.deleteCol, "name")
	return u
}

func (u *__FileRef_Deleter) Delete_StorageUserId() *__FileRef_Deleter {
	u.deleteCol = append(u.deleteCol, "storage_user_id")
	return u
}

func (u *__FileRef_Deleter) Delete_StoreType() *__FileRef_Deleter {
	u.deleteCol = append(u.deleteCol, "store_type")
	return u
}

func (u *__FileRef_Deleter) Delete_UserId() *__FileRef_Deleter {
	u.deleteCol = append(u.deleteCol, "user_id")
	return u
}

func (u *__FileRef_Deleter) Delete_Width() *__FileRef_Deleter {
	u.deleteCol = append(u.deleteCol, "width")
	return u
}

//////////////////// End of just Deleter

//////////////////// just Updater
//each column delete

func (u *__FileRef_Updater) Extension(newVal string) *__FileRef_Updater {
	u.updates["extension = ? "] = newVal
	return u
}

func (u *__FileRef_Updater) FileRefId(newVal int) *__FileRef_Updater {
	u.updates["file_ref_id = ? "] = newVal
	return u
}

func (u *__FileRef_Updater) FirstFileRefId(newVal int) *__FileRef_Updater {
	u.updates["first_file_ref_id = ? "] = newVal
	return u
}

func (u *__FileRef_Updater) Height(newVal int) *__FileRef_Updater {
	u.updates["height = ? "] = newVal
	return u
}

func (u *__FileRef_Updater) Length(newVal int) *__FileRef_Updater {
	u.updates["length = ? "] = newVal
	return u
}

func (u *__FileRef_Updater) Md5Key(newVal string) *__FileRef_Updater {
	u.updates["md5_key = ? "] = newVal
	return u
}

func (u *__FileRef_Updater) MimeType(newVal string) *__FileRef_Updater {
	u.updates["mime_type = ? "] = newVal
	return u
}

func (u *__FileRef_Updater) Name(newVal string) *__FileRef_Updater {
	u.updates["name = ? "] = newVal
	return u
}

func (u *__FileRef_Updater) StorageUserId(newVal int) *__FileRef_Updater {
	u.updates["storage_user_id = ? "] = newVal
	return u
}

func (u *__FileRef_Updater) StoreType(newVal string) *__FileRef_Updater {
	u.updates["store_type = ? "] = newVal
	return u
}

func (u *__FileRef_Updater) UserId(newVal int) *__FileRef_Updater {
	u.updates["user_id = ? "] = newVal
	return u
}

func (u *__FileRef_Updater) Width(newVal int) *__FileRef_Updater {
	u.updates["width = ? "] = newVal
	return u
}

//////////////////// End just Updater

//{_Eq_FILTERING  =  Extension_Eq_FILTERING}

func (d *__FileRef_Deleter) Extension_Eq_FILTERING(val string) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " extension = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_FILTERING And = And And_Extension_Eq_FILTERING}

func (d *__FileRef_Deleter) And_Extension_Eq_FILTERING(val string) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And extension = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_FILTERING Or = Or Or_Extension_Eq_FILTERING}

func (d *__FileRef_Deleter) Or_Extension_Eq_FILTERING(val string) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or extension = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq  =  FileRefId_Eq}

func (d *__FileRef_Deleter) FileRefId_Eq(val int) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " file_ref_id = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LT_Filtering  <  FileRefId_LT_Filtering}

func (d *__FileRef_Deleter) FileRefId_LT_Filtering(val int) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " file_ref_id < ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LE_Filtering  <=  FileRefId_LE_Filtering}

func (d *__FileRef_Deleter) FileRefId_LE_Filtering(val int) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " file_ref_id <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GT_Filtering  >  FileRefId_GT_Filtering}

func (d *__FileRef_Deleter) FileRefId_GT_Filtering(val int) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " file_ref_id > ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GE_Filtering  >=  FileRefId_GE_Filtering}

func (d *__FileRef_Deleter) FileRefId_GE_Filtering(val int) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " file_ref_id >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq And = And And_FileRefId_Eq}

func (d *__FileRef_Deleter) And_FileRefId_Eq(val int) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And file_ref_id = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LT_Filtering And < And And_FileRefId_LT_Filtering}

func (d *__FileRef_Deleter) And_FileRefId_LT_Filtering(val int) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And file_ref_id < ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LE_Filtering And <= And And_FileRefId_LE_Filtering}

func (d *__FileRef_Deleter) And_FileRefId_LE_Filtering(val int) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And file_ref_id <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GT_Filtering And > And And_FileRefId_GT_Filtering}

func (d *__FileRef_Deleter) And_FileRefId_GT_Filtering(val int) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And file_ref_id > ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GE_Filtering And >= And And_FileRefId_GE_Filtering}

func (d *__FileRef_Deleter) And_FileRefId_GE_Filtering(val int) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And file_ref_id >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq Or = Or Or_FileRefId_Eq}

func (d *__FileRef_Deleter) Or_FileRefId_Eq(val int) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or file_ref_id = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LT_Filtering Or < Or Or_FileRefId_LT_Filtering}

func (d *__FileRef_Deleter) Or_FileRefId_LT_Filtering(val int) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or file_ref_id < ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LE_Filtering Or <= Or Or_FileRefId_LE_Filtering}

func (d *__FileRef_Deleter) Or_FileRefId_LE_Filtering(val int) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or file_ref_id <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GT_Filtering Or > Or Or_FileRefId_GT_Filtering}

func (d *__FileRef_Deleter) Or_FileRefId_GT_Filtering(val int) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or file_ref_id > ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GE_Filtering Or >= Or Or_FileRefId_GE_Filtering}

func (d *__FileRef_Deleter) Or_FileRefId_GE_Filtering(val int) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or file_ref_id >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_Filtering  =  FirstFileRefId_Eq_Filtering}

func (d *__FileRef_Deleter) FirstFileRefId_Eq_Filtering(val int) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " first_file_ref_id = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LT_Filtering  <  FirstFileRefId_LT_Filtering}

func (d *__FileRef_Deleter) FirstFileRefId_LT_Filtering(val int) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " first_file_ref_id < ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LE_Filtering  <=  FirstFileRefId_LE_Filtering}

func (d *__FileRef_Deleter) FirstFileRefId_LE_Filtering(val int) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " first_file_ref_id <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GT_Filtering  >  FirstFileRefId_GT_Filtering}

func (d *__FileRef_Deleter) FirstFileRefId_GT_Filtering(val int) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " first_file_ref_id > ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GE_Filtering  >=  FirstFileRefId_GE_Filtering}

func (d *__FileRef_Deleter) FirstFileRefId_GE_Filtering(val int) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " first_file_ref_id >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_Filtering And = And And_FirstFileRefId_Eq_Filtering}

func (d *__FileRef_Deleter) And_FirstFileRefId_Eq_Filtering(val int) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And first_file_ref_id = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LT_Filtering And < And And_FirstFileRefId_LT_Filtering}

func (d *__FileRef_Deleter) And_FirstFileRefId_LT_Filtering(val int) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And first_file_ref_id < ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LE_Filtering And <= And And_FirstFileRefId_LE_Filtering}

func (d *__FileRef_Deleter) And_FirstFileRefId_LE_Filtering(val int) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And first_file_ref_id <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GT_Filtering And > And And_FirstFileRefId_GT_Filtering}

func (d *__FileRef_Deleter) And_FirstFileRefId_GT_Filtering(val int) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And first_file_ref_id > ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GE_Filtering And >= And And_FirstFileRefId_GE_Filtering}

func (d *__FileRef_Deleter) And_FirstFileRefId_GE_Filtering(val int) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And first_file_ref_id >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_Filtering Or = Or Or_FirstFileRefId_Eq_Filtering}

func (d *__FileRef_Deleter) Or_FirstFileRefId_Eq_Filtering(val int) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or first_file_ref_id = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LT_Filtering Or < Or Or_FirstFileRefId_LT_Filtering}

func (d *__FileRef_Deleter) Or_FirstFileRefId_LT_Filtering(val int) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or first_file_ref_id < ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LE_Filtering Or <= Or Or_FirstFileRefId_LE_Filtering}

func (d *__FileRef_Deleter) Or_FirstFileRefId_LE_Filtering(val int) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or first_file_ref_id <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GT_Filtering Or > Or Or_FirstFileRefId_GT_Filtering}

func (d *__FileRef_Deleter) Or_FirstFileRefId_GT_Filtering(val int) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or first_file_ref_id > ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GE_Filtering Or >= Or Or_FirstFileRefId_GE_Filtering}

func (d *__FileRef_Deleter) Or_FirstFileRefId_GE_Filtering(val int) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or first_file_ref_id >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_Filtering  =  Height_Eq_Filtering}

func (d *__FileRef_Deleter) Height_Eq_Filtering(val int) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " height = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LT_Filtering  <  Height_LT_Filtering}

func (d *__FileRef_Deleter) Height_LT_Filtering(val int) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " height < ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LE_Filtering  <=  Height_LE_Filtering}

func (d *__FileRef_Deleter) Height_LE_Filtering(val int) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " height <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GT_Filtering  >  Height_GT_Filtering}

func (d *__FileRef_Deleter) Height_GT_Filtering(val int) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " height > ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GE_Filtering  >=  Height_GE_Filtering}

func (d *__FileRef_Deleter) Height_GE_Filtering(val int) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " height >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_Filtering And = And And_Height_Eq_Filtering}

func (d *__FileRef_Deleter) And_Height_Eq_Filtering(val int) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And height = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LT_Filtering And < And And_Height_LT_Filtering}

func (d *__FileRef_Deleter) And_Height_LT_Filtering(val int) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And height < ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LE_Filtering And <= And And_Height_LE_Filtering}

func (d *__FileRef_Deleter) And_Height_LE_Filtering(val int) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And height <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GT_Filtering And > And And_Height_GT_Filtering}

func (d *__FileRef_Deleter) And_Height_GT_Filtering(val int) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And height > ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GE_Filtering And >= And And_Height_GE_Filtering}

func (d *__FileRef_Deleter) And_Height_GE_Filtering(val int) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And height >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_Filtering Or = Or Or_Height_Eq_Filtering}

func (d *__FileRef_Deleter) Or_Height_Eq_Filtering(val int) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or height = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LT_Filtering Or < Or Or_Height_LT_Filtering}

func (d *__FileRef_Deleter) Or_Height_LT_Filtering(val int) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or height < ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LE_Filtering Or <= Or Or_Height_LE_Filtering}

func (d *__FileRef_Deleter) Or_Height_LE_Filtering(val int) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or height <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GT_Filtering Or > Or Or_Height_GT_Filtering}

func (d *__FileRef_Deleter) Or_Height_GT_Filtering(val int) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or height > ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GE_Filtering Or >= Or Or_Height_GE_Filtering}

func (d *__FileRef_Deleter) Or_Height_GE_Filtering(val int) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or height >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_Filtering  =  Length_Eq_Filtering}

func (d *__FileRef_Deleter) Length_Eq_Filtering(val int) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " length = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LT_Filtering  <  Length_LT_Filtering}

func (d *__FileRef_Deleter) Length_LT_Filtering(val int) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " length < ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LE_Filtering  <=  Length_LE_Filtering}

func (d *__FileRef_Deleter) Length_LE_Filtering(val int) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " length <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GT_Filtering  >  Length_GT_Filtering}

func (d *__FileRef_Deleter) Length_GT_Filtering(val int) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " length > ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GE_Filtering  >=  Length_GE_Filtering}

func (d *__FileRef_Deleter) Length_GE_Filtering(val int) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " length >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_Filtering And = And And_Length_Eq_Filtering}

func (d *__FileRef_Deleter) And_Length_Eq_Filtering(val int) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And length = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LT_Filtering And < And And_Length_LT_Filtering}

func (d *__FileRef_Deleter) And_Length_LT_Filtering(val int) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And length < ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LE_Filtering And <= And And_Length_LE_Filtering}

func (d *__FileRef_Deleter) And_Length_LE_Filtering(val int) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And length <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GT_Filtering And > And And_Length_GT_Filtering}

func (d *__FileRef_Deleter) And_Length_GT_Filtering(val int) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And length > ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GE_Filtering And >= And And_Length_GE_Filtering}

func (d *__FileRef_Deleter) And_Length_GE_Filtering(val int) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And length >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_Filtering Or = Or Or_Length_Eq_Filtering}

func (d *__FileRef_Deleter) Or_Length_Eq_Filtering(val int) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or length = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LT_Filtering Or < Or Or_Length_LT_Filtering}

func (d *__FileRef_Deleter) Or_Length_LT_Filtering(val int) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or length < ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LE_Filtering Or <= Or Or_Length_LE_Filtering}

func (d *__FileRef_Deleter) Or_Length_LE_Filtering(val int) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or length <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GT_Filtering Or > Or Or_Length_GT_Filtering}

func (d *__FileRef_Deleter) Or_Length_GT_Filtering(val int) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or length > ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GE_Filtering Or >= Or Or_Length_GE_Filtering}

func (d *__FileRef_Deleter) Or_Length_GE_Filtering(val int) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or length >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_FILTERING  =  Md5Key_Eq_FILTERING}

func (d *__FileRef_Deleter) Md5Key_Eq_FILTERING(val string) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " md5_key = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_FILTERING And = And And_Md5Key_Eq_FILTERING}

func (d *__FileRef_Deleter) And_Md5Key_Eq_FILTERING(val string) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And md5_key = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_FILTERING Or = Or Or_Md5Key_Eq_FILTERING}

func (d *__FileRef_Deleter) Or_Md5Key_Eq_FILTERING(val string) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or md5_key = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_FILTERING  =  MimeType_Eq_FILTERING}

func (d *__FileRef_Deleter) MimeType_Eq_FILTERING(val string) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " mime_type = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_FILTERING And = And And_MimeType_Eq_FILTERING}

func (d *__FileRef_Deleter) And_MimeType_Eq_FILTERING(val string) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And mime_type = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_FILTERING Or = Or Or_MimeType_Eq_FILTERING}

func (d *__FileRef_Deleter) Or_MimeType_Eq_FILTERING(val string) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or mime_type = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_FILTERING  =  Name_Eq_FILTERING}

func (d *__FileRef_Deleter) Name_Eq_FILTERING(val string) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " name = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_FILTERING And = And And_Name_Eq_FILTERING}

func (d *__FileRef_Deleter) And_Name_Eq_FILTERING(val string) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And name = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_FILTERING Or = Or Or_Name_Eq_FILTERING}

func (d *__FileRef_Deleter) Or_Name_Eq_FILTERING(val string) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or name = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_Filtering  =  StorageUserId_Eq_Filtering}

func (d *__FileRef_Deleter) StorageUserId_Eq_Filtering(val int) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " storage_user_id = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LT_Filtering  <  StorageUserId_LT_Filtering}

func (d *__FileRef_Deleter) StorageUserId_LT_Filtering(val int) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " storage_user_id < ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LE_Filtering  <=  StorageUserId_LE_Filtering}

func (d *__FileRef_Deleter) StorageUserId_LE_Filtering(val int) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " storage_user_id <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GT_Filtering  >  StorageUserId_GT_Filtering}

func (d *__FileRef_Deleter) StorageUserId_GT_Filtering(val int) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " storage_user_id > ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GE_Filtering  >=  StorageUserId_GE_Filtering}

func (d *__FileRef_Deleter) StorageUserId_GE_Filtering(val int) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " storage_user_id >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_Filtering And = And And_StorageUserId_Eq_Filtering}

func (d *__FileRef_Deleter) And_StorageUserId_Eq_Filtering(val int) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And storage_user_id = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LT_Filtering And < And And_StorageUserId_LT_Filtering}

func (d *__FileRef_Deleter) And_StorageUserId_LT_Filtering(val int) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And storage_user_id < ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LE_Filtering And <= And And_StorageUserId_LE_Filtering}

func (d *__FileRef_Deleter) And_StorageUserId_LE_Filtering(val int) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And storage_user_id <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GT_Filtering And > And And_StorageUserId_GT_Filtering}

func (d *__FileRef_Deleter) And_StorageUserId_GT_Filtering(val int) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And storage_user_id > ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GE_Filtering And >= And And_StorageUserId_GE_Filtering}

func (d *__FileRef_Deleter) And_StorageUserId_GE_Filtering(val int) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And storage_user_id >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_Filtering Or = Or Or_StorageUserId_Eq_Filtering}

func (d *__FileRef_Deleter) Or_StorageUserId_Eq_Filtering(val int) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or storage_user_id = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LT_Filtering Or < Or Or_StorageUserId_LT_Filtering}

func (d *__FileRef_Deleter) Or_StorageUserId_LT_Filtering(val int) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or storage_user_id < ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LE_Filtering Or <= Or Or_StorageUserId_LE_Filtering}

func (d *__FileRef_Deleter) Or_StorageUserId_LE_Filtering(val int) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or storage_user_id <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GT_Filtering Or > Or Or_StorageUserId_GT_Filtering}

func (d *__FileRef_Deleter) Or_StorageUserId_GT_Filtering(val int) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or storage_user_id > ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GE_Filtering Or >= Or Or_StorageUserId_GE_Filtering}

func (d *__FileRef_Deleter) Or_StorageUserId_GE_Filtering(val int) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or storage_user_id >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_FILTERING  =  StoreType_Eq_FILTERING}

func (d *__FileRef_Deleter) StoreType_Eq_FILTERING(val string) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " store_type = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_FILTERING And = And And_StoreType_Eq_FILTERING}

func (d *__FileRef_Deleter) And_StoreType_Eq_FILTERING(val string) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And store_type = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_FILTERING Or = Or Or_StoreType_Eq_FILTERING}

func (d *__FileRef_Deleter) Or_StoreType_Eq_FILTERING(val string) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or store_type = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_Filtering  =  UserId_Eq_Filtering}

func (d *__FileRef_Deleter) UserId_Eq_Filtering(val int) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " user_id = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LT_Filtering  <  UserId_LT_Filtering}

func (d *__FileRef_Deleter) UserId_LT_Filtering(val int) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " user_id < ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LE_Filtering  <=  UserId_LE_Filtering}

func (d *__FileRef_Deleter) UserId_LE_Filtering(val int) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " user_id <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GT_Filtering  >  UserId_GT_Filtering}

func (d *__FileRef_Deleter) UserId_GT_Filtering(val int) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " user_id > ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GE_Filtering  >=  UserId_GE_Filtering}

func (d *__FileRef_Deleter) UserId_GE_Filtering(val int) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " user_id >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_Filtering And = And And_UserId_Eq_Filtering}

func (d *__FileRef_Deleter) And_UserId_Eq_Filtering(val int) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And user_id = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LT_Filtering And < And And_UserId_LT_Filtering}

func (d *__FileRef_Deleter) And_UserId_LT_Filtering(val int) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And user_id < ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LE_Filtering And <= And And_UserId_LE_Filtering}

func (d *__FileRef_Deleter) And_UserId_LE_Filtering(val int) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And user_id <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GT_Filtering And > And And_UserId_GT_Filtering}

func (d *__FileRef_Deleter) And_UserId_GT_Filtering(val int) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And user_id > ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GE_Filtering And >= And And_UserId_GE_Filtering}

func (d *__FileRef_Deleter) And_UserId_GE_Filtering(val int) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And user_id >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_Filtering Or = Or Or_UserId_Eq_Filtering}

func (d *__FileRef_Deleter) Or_UserId_Eq_Filtering(val int) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or user_id = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LT_Filtering Or < Or Or_UserId_LT_Filtering}

func (d *__FileRef_Deleter) Or_UserId_LT_Filtering(val int) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or user_id < ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LE_Filtering Or <= Or Or_UserId_LE_Filtering}

func (d *__FileRef_Deleter) Or_UserId_LE_Filtering(val int) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or user_id <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GT_Filtering Or > Or Or_UserId_GT_Filtering}

func (d *__FileRef_Deleter) Or_UserId_GT_Filtering(val int) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or user_id > ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GE_Filtering Or >= Or Or_UserId_GE_Filtering}

func (d *__FileRef_Deleter) Or_UserId_GE_Filtering(val int) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or user_id >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_Filtering  =  Width_Eq_Filtering}

func (d *__FileRef_Deleter) Width_Eq_Filtering(val int) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " width = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LT_Filtering  <  Width_LT_Filtering}

func (d *__FileRef_Deleter) Width_LT_Filtering(val int) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " width < ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LE_Filtering  <=  Width_LE_Filtering}

func (d *__FileRef_Deleter) Width_LE_Filtering(val int) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " width <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GT_Filtering  >  Width_GT_Filtering}

func (d *__FileRef_Deleter) Width_GT_Filtering(val int) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " width > ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GE_Filtering  >=  Width_GE_Filtering}

func (d *__FileRef_Deleter) Width_GE_Filtering(val int) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " width >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_Filtering And = And And_Width_Eq_Filtering}

func (d *__FileRef_Deleter) And_Width_Eq_Filtering(val int) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And width = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LT_Filtering And < And And_Width_LT_Filtering}

func (d *__FileRef_Deleter) And_Width_LT_Filtering(val int) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And width < ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LE_Filtering And <= And And_Width_LE_Filtering}

func (d *__FileRef_Deleter) And_Width_LE_Filtering(val int) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And width <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GT_Filtering And > And And_Width_GT_Filtering}

func (d *__FileRef_Deleter) And_Width_GT_Filtering(val int) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And width > ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GE_Filtering And >= And And_Width_GE_Filtering}

func (d *__FileRef_Deleter) And_Width_GE_Filtering(val int) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And width >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_Filtering Or = Or Or_Width_Eq_Filtering}

func (d *__FileRef_Deleter) Or_Width_Eq_Filtering(val int) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or width = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LT_Filtering Or < Or Or_Width_LT_Filtering}

func (d *__FileRef_Deleter) Or_Width_LT_Filtering(val int) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or width < ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LE_Filtering Or <= Or Or_Width_LE_Filtering}

func (d *__FileRef_Deleter) Or_Width_LE_Filtering(val int) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or width <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GT_Filtering Or > Or Or_Width_GT_Filtering}

func (d *__FileRef_Deleter) Or_Width_GT_Filtering(val int) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or width > ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GE_Filtering Or >= Or Or_Width_GE_Filtering}

func (d *__FileRef_Deleter) Or_Width_GE_Filtering(val int) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or width >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_FILTERING  =  Extension_Eq_FILTERING}

func (d *__FileRef_Updater) Extension_Eq_FILTERING(val string) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " extension = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_FILTERING And = And And_Extension_Eq_FILTERING}

func (d *__FileRef_Updater) And_Extension_Eq_FILTERING(val string) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And extension = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_FILTERING Or = Or Or_Extension_Eq_FILTERING}

func (d *__FileRef_Updater) Or_Extension_Eq_FILTERING(val string) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or extension = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq  =  FileRefId_Eq}

func (d *__FileRef_Updater) FileRefId_Eq(val int) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " file_ref_id = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LT_Filtering  <  FileRefId_LT_Filtering}

func (d *__FileRef_Updater) FileRefId_LT_Filtering(val int) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " file_ref_id < ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LE_Filtering  <=  FileRefId_LE_Filtering}

func (d *__FileRef_Updater) FileRefId_LE_Filtering(val int) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " file_ref_id <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GT_Filtering  >  FileRefId_GT_Filtering}

func (d *__FileRef_Updater) FileRefId_GT_Filtering(val int) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " file_ref_id > ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GE_Filtering  >=  FileRefId_GE_Filtering}

func (d *__FileRef_Updater) FileRefId_GE_Filtering(val int) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " file_ref_id >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq And = And And_FileRefId_Eq}

func (d *__FileRef_Updater) And_FileRefId_Eq(val int) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And file_ref_id = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LT_Filtering And < And And_FileRefId_LT_Filtering}

func (d *__FileRef_Updater) And_FileRefId_LT_Filtering(val int) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And file_ref_id < ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LE_Filtering And <= And And_FileRefId_LE_Filtering}

func (d *__FileRef_Updater) And_FileRefId_LE_Filtering(val int) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And file_ref_id <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GT_Filtering And > And And_FileRefId_GT_Filtering}

func (d *__FileRef_Updater) And_FileRefId_GT_Filtering(val int) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And file_ref_id > ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GE_Filtering And >= And And_FileRefId_GE_Filtering}

func (d *__FileRef_Updater) And_FileRefId_GE_Filtering(val int) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And file_ref_id >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq Or = Or Or_FileRefId_Eq}

func (d *__FileRef_Updater) Or_FileRefId_Eq(val int) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or file_ref_id = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LT_Filtering Or < Or Or_FileRefId_LT_Filtering}

func (d *__FileRef_Updater) Or_FileRefId_LT_Filtering(val int) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or file_ref_id < ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LE_Filtering Or <= Or Or_FileRefId_LE_Filtering}

func (d *__FileRef_Updater) Or_FileRefId_LE_Filtering(val int) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or file_ref_id <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GT_Filtering Or > Or Or_FileRefId_GT_Filtering}

func (d *__FileRef_Updater) Or_FileRefId_GT_Filtering(val int) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or file_ref_id > ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GE_Filtering Or >= Or Or_FileRefId_GE_Filtering}

func (d *__FileRef_Updater) Or_FileRefId_GE_Filtering(val int) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or file_ref_id >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_Filtering  =  FirstFileRefId_Eq_Filtering}

func (d *__FileRef_Updater) FirstFileRefId_Eq_Filtering(val int) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " first_file_ref_id = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LT_Filtering  <  FirstFileRefId_LT_Filtering}

func (d *__FileRef_Updater) FirstFileRefId_LT_Filtering(val int) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " first_file_ref_id < ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LE_Filtering  <=  FirstFileRefId_LE_Filtering}

func (d *__FileRef_Updater) FirstFileRefId_LE_Filtering(val int) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " first_file_ref_id <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GT_Filtering  >  FirstFileRefId_GT_Filtering}

func (d *__FileRef_Updater) FirstFileRefId_GT_Filtering(val int) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " first_file_ref_id > ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GE_Filtering  >=  FirstFileRefId_GE_Filtering}

func (d *__FileRef_Updater) FirstFileRefId_GE_Filtering(val int) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " first_file_ref_id >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_Filtering And = And And_FirstFileRefId_Eq_Filtering}

func (d *__FileRef_Updater) And_FirstFileRefId_Eq_Filtering(val int) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And first_file_ref_id = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LT_Filtering And < And And_FirstFileRefId_LT_Filtering}

func (d *__FileRef_Updater) And_FirstFileRefId_LT_Filtering(val int) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And first_file_ref_id < ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LE_Filtering And <= And And_FirstFileRefId_LE_Filtering}

func (d *__FileRef_Updater) And_FirstFileRefId_LE_Filtering(val int) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And first_file_ref_id <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GT_Filtering And > And And_FirstFileRefId_GT_Filtering}

func (d *__FileRef_Updater) And_FirstFileRefId_GT_Filtering(val int) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And first_file_ref_id > ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GE_Filtering And >= And And_FirstFileRefId_GE_Filtering}

func (d *__FileRef_Updater) And_FirstFileRefId_GE_Filtering(val int) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And first_file_ref_id >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_Filtering Or = Or Or_FirstFileRefId_Eq_Filtering}

func (d *__FileRef_Updater) Or_FirstFileRefId_Eq_Filtering(val int) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or first_file_ref_id = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LT_Filtering Or < Or Or_FirstFileRefId_LT_Filtering}

func (d *__FileRef_Updater) Or_FirstFileRefId_LT_Filtering(val int) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or first_file_ref_id < ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LE_Filtering Or <= Or Or_FirstFileRefId_LE_Filtering}

func (d *__FileRef_Updater) Or_FirstFileRefId_LE_Filtering(val int) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or first_file_ref_id <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GT_Filtering Or > Or Or_FirstFileRefId_GT_Filtering}

func (d *__FileRef_Updater) Or_FirstFileRefId_GT_Filtering(val int) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or first_file_ref_id > ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GE_Filtering Or >= Or Or_FirstFileRefId_GE_Filtering}

func (d *__FileRef_Updater) Or_FirstFileRefId_GE_Filtering(val int) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or first_file_ref_id >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_Filtering  =  Height_Eq_Filtering}

func (d *__FileRef_Updater) Height_Eq_Filtering(val int) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " height = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LT_Filtering  <  Height_LT_Filtering}

func (d *__FileRef_Updater) Height_LT_Filtering(val int) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " height < ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LE_Filtering  <=  Height_LE_Filtering}

func (d *__FileRef_Updater) Height_LE_Filtering(val int) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " height <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GT_Filtering  >  Height_GT_Filtering}

func (d *__FileRef_Updater) Height_GT_Filtering(val int) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " height > ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GE_Filtering  >=  Height_GE_Filtering}

func (d *__FileRef_Updater) Height_GE_Filtering(val int) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " height >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_Filtering And = And And_Height_Eq_Filtering}

func (d *__FileRef_Updater) And_Height_Eq_Filtering(val int) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And height = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LT_Filtering And < And And_Height_LT_Filtering}

func (d *__FileRef_Updater) And_Height_LT_Filtering(val int) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And height < ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LE_Filtering And <= And And_Height_LE_Filtering}

func (d *__FileRef_Updater) And_Height_LE_Filtering(val int) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And height <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GT_Filtering And > And And_Height_GT_Filtering}

func (d *__FileRef_Updater) And_Height_GT_Filtering(val int) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And height > ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GE_Filtering And >= And And_Height_GE_Filtering}

func (d *__FileRef_Updater) And_Height_GE_Filtering(val int) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And height >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_Filtering Or = Or Or_Height_Eq_Filtering}

func (d *__FileRef_Updater) Or_Height_Eq_Filtering(val int) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or height = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LT_Filtering Or < Or Or_Height_LT_Filtering}

func (d *__FileRef_Updater) Or_Height_LT_Filtering(val int) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or height < ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LE_Filtering Or <= Or Or_Height_LE_Filtering}

func (d *__FileRef_Updater) Or_Height_LE_Filtering(val int) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or height <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GT_Filtering Or > Or Or_Height_GT_Filtering}

func (d *__FileRef_Updater) Or_Height_GT_Filtering(val int) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or height > ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GE_Filtering Or >= Or Or_Height_GE_Filtering}

func (d *__FileRef_Updater) Or_Height_GE_Filtering(val int) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or height >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_Filtering  =  Length_Eq_Filtering}

func (d *__FileRef_Updater) Length_Eq_Filtering(val int) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " length = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LT_Filtering  <  Length_LT_Filtering}

func (d *__FileRef_Updater) Length_LT_Filtering(val int) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " length < ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LE_Filtering  <=  Length_LE_Filtering}

func (d *__FileRef_Updater) Length_LE_Filtering(val int) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " length <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GT_Filtering  >  Length_GT_Filtering}

func (d *__FileRef_Updater) Length_GT_Filtering(val int) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " length > ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GE_Filtering  >=  Length_GE_Filtering}

func (d *__FileRef_Updater) Length_GE_Filtering(val int) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " length >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_Filtering And = And And_Length_Eq_Filtering}

func (d *__FileRef_Updater) And_Length_Eq_Filtering(val int) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And length = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LT_Filtering And < And And_Length_LT_Filtering}

func (d *__FileRef_Updater) And_Length_LT_Filtering(val int) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And length < ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LE_Filtering And <= And And_Length_LE_Filtering}

func (d *__FileRef_Updater) And_Length_LE_Filtering(val int) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And length <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GT_Filtering And > And And_Length_GT_Filtering}

func (d *__FileRef_Updater) And_Length_GT_Filtering(val int) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And length > ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GE_Filtering And >= And And_Length_GE_Filtering}

func (d *__FileRef_Updater) And_Length_GE_Filtering(val int) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And length >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_Filtering Or = Or Or_Length_Eq_Filtering}

func (d *__FileRef_Updater) Or_Length_Eq_Filtering(val int) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or length = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LT_Filtering Or < Or Or_Length_LT_Filtering}

func (d *__FileRef_Updater) Or_Length_LT_Filtering(val int) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or length < ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LE_Filtering Or <= Or Or_Length_LE_Filtering}

func (d *__FileRef_Updater) Or_Length_LE_Filtering(val int) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or length <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GT_Filtering Or > Or Or_Length_GT_Filtering}

func (d *__FileRef_Updater) Or_Length_GT_Filtering(val int) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or length > ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GE_Filtering Or >= Or Or_Length_GE_Filtering}

func (d *__FileRef_Updater) Or_Length_GE_Filtering(val int) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or length >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_FILTERING  =  Md5Key_Eq_FILTERING}

func (d *__FileRef_Updater) Md5Key_Eq_FILTERING(val string) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " md5_key = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_FILTERING And = And And_Md5Key_Eq_FILTERING}

func (d *__FileRef_Updater) And_Md5Key_Eq_FILTERING(val string) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And md5_key = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_FILTERING Or = Or Or_Md5Key_Eq_FILTERING}

func (d *__FileRef_Updater) Or_Md5Key_Eq_FILTERING(val string) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or md5_key = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_FILTERING  =  MimeType_Eq_FILTERING}

func (d *__FileRef_Updater) MimeType_Eq_FILTERING(val string) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " mime_type = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_FILTERING And = And And_MimeType_Eq_FILTERING}

func (d *__FileRef_Updater) And_MimeType_Eq_FILTERING(val string) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And mime_type = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_FILTERING Or = Or Or_MimeType_Eq_FILTERING}

func (d *__FileRef_Updater) Or_MimeType_Eq_FILTERING(val string) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or mime_type = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_FILTERING  =  Name_Eq_FILTERING}

func (d *__FileRef_Updater) Name_Eq_FILTERING(val string) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " name = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_FILTERING And = And And_Name_Eq_FILTERING}

func (d *__FileRef_Updater) And_Name_Eq_FILTERING(val string) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And name = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_FILTERING Or = Or Or_Name_Eq_FILTERING}

func (d *__FileRef_Updater) Or_Name_Eq_FILTERING(val string) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or name = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_Filtering  =  StorageUserId_Eq_Filtering}

func (d *__FileRef_Updater) StorageUserId_Eq_Filtering(val int) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " storage_user_id = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LT_Filtering  <  StorageUserId_LT_Filtering}

func (d *__FileRef_Updater) StorageUserId_LT_Filtering(val int) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " storage_user_id < ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LE_Filtering  <=  StorageUserId_LE_Filtering}

func (d *__FileRef_Updater) StorageUserId_LE_Filtering(val int) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " storage_user_id <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GT_Filtering  >  StorageUserId_GT_Filtering}

func (d *__FileRef_Updater) StorageUserId_GT_Filtering(val int) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " storage_user_id > ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GE_Filtering  >=  StorageUserId_GE_Filtering}

func (d *__FileRef_Updater) StorageUserId_GE_Filtering(val int) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " storage_user_id >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_Filtering And = And And_StorageUserId_Eq_Filtering}

func (d *__FileRef_Updater) And_StorageUserId_Eq_Filtering(val int) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And storage_user_id = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LT_Filtering And < And And_StorageUserId_LT_Filtering}

func (d *__FileRef_Updater) And_StorageUserId_LT_Filtering(val int) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And storage_user_id < ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LE_Filtering And <= And And_StorageUserId_LE_Filtering}

func (d *__FileRef_Updater) And_StorageUserId_LE_Filtering(val int) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And storage_user_id <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GT_Filtering And > And And_StorageUserId_GT_Filtering}

func (d *__FileRef_Updater) And_StorageUserId_GT_Filtering(val int) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And storage_user_id > ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GE_Filtering And >= And And_StorageUserId_GE_Filtering}

func (d *__FileRef_Updater) And_StorageUserId_GE_Filtering(val int) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And storage_user_id >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_Filtering Or = Or Or_StorageUserId_Eq_Filtering}

func (d *__FileRef_Updater) Or_StorageUserId_Eq_Filtering(val int) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or storage_user_id = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LT_Filtering Or < Or Or_StorageUserId_LT_Filtering}

func (d *__FileRef_Updater) Or_StorageUserId_LT_Filtering(val int) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or storage_user_id < ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LE_Filtering Or <= Or Or_StorageUserId_LE_Filtering}

func (d *__FileRef_Updater) Or_StorageUserId_LE_Filtering(val int) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or storage_user_id <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GT_Filtering Or > Or Or_StorageUserId_GT_Filtering}

func (d *__FileRef_Updater) Or_StorageUserId_GT_Filtering(val int) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or storage_user_id > ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GE_Filtering Or >= Or Or_StorageUserId_GE_Filtering}

func (d *__FileRef_Updater) Or_StorageUserId_GE_Filtering(val int) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or storage_user_id >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_FILTERING  =  StoreType_Eq_FILTERING}

func (d *__FileRef_Updater) StoreType_Eq_FILTERING(val string) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " store_type = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_FILTERING And = And And_StoreType_Eq_FILTERING}

func (d *__FileRef_Updater) And_StoreType_Eq_FILTERING(val string) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And store_type = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_FILTERING Or = Or Or_StoreType_Eq_FILTERING}

func (d *__FileRef_Updater) Or_StoreType_Eq_FILTERING(val string) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or store_type = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_Filtering  =  UserId_Eq_Filtering}

func (d *__FileRef_Updater) UserId_Eq_Filtering(val int) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " user_id = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LT_Filtering  <  UserId_LT_Filtering}

func (d *__FileRef_Updater) UserId_LT_Filtering(val int) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " user_id < ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LE_Filtering  <=  UserId_LE_Filtering}

func (d *__FileRef_Updater) UserId_LE_Filtering(val int) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " user_id <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GT_Filtering  >  UserId_GT_Filtering}

func (d *__FileRef_Updater) UserId_GT_Filtering(val int) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " user_id > ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GE_Filtering  >=  UserId_GE_Filtering}

func (d *__FileRef_Updater) UserId_GE_Filtering(val int) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " user_id >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_Filtering And = And And_UserId_Eq_Filtering}

func (d *__FileRef_Updater) And_UserId_Eq_Filtering(val int) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And user_id = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LT_Filtering And < And And_UserId_LT_Filtering}

func (d *__FileRef_Updater) And_UserId_LT_Filtering(val int) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And user_id < ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LE_Filtering And <= And And_UserId_LE_Filtering}

func (d *__FileRef_Updater) And_UserId_LE_Filtering(val int) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And user_id <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GT_Filtering And > And And_UserId_GT_Filtering}

func (d *__FileRef_Updater) And_UserId_GT_Filtering(val int) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And user_id > ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GE_Filtering And >= And And_UserId_GE_Filtering}

func (d *__FileRef_Updater) And_UserId_GE_Filtering(val int) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And user_id >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_Filtering Or = Or Or_UserId_Eq_Filtering}

func (d *__FileRef_Updater) Or_UserId_Eq_Filtering(val int) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or user_id = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LT_Filtering Or < Or Or_UserId_LT_Filtering}

func (d *__FileRef_Updater) Or_UserId_LT_Filtering(val int) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or user_id < ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LE_Filtering Or <= Or Or_UserId_LE_Filtering}

func (d *__FileRef_Updater) Or_UserId_LE_Filtering(val int) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or user_id <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GT_Filtering Or > Or Or_UserId_GT_Filtering}

func (d *__FileRef_Updater) Or_UserId_GT_Filtering(val int) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or user_id > ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GE_Filtering Or >= Or Or_UserId_GE_Filtering}

func (d *__FileRef_Updater) Or_UserId_GE_Filtering(val int) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or user_id >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_Filtering  =  Width_Eq_Filtering}

func (d *__FileRef_Updater) Width_Eq_Filtering(val int) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " width = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LT_Filtering  <  Width_LT_Filtering}

func (d *__FileRef_Updater) Width_LT_Filtering(val int) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " width < ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LE_Filtering  <=  Width_LE_Filtering}

func (d *__FileRef_Updater) Width_LE_Filtering(val int) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " width <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GT_Filtering  >  Width_GT_Filtering}

func (d *__FileRef_Updater) Width_GT_Filtering(val int) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " width > ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GE_Filtering  >=  Width_GE_Filtering}

func (d *__FileRef_Updater) Width_GE_Filtering(val int) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " width >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_Filtering And = And And_Width_Eq_Filtering}

func (d *__FileRef_Updater) And_Width_Eq_Filtering(val int) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And width = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LT_Filtering And < And And_Width_LT_Filtering}

func (d *__FileRef_Updater) And_Width_LT_Filtering(val int) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And width < ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LE_Filtering And <= And And_Width_LE_Filtering}

func (d *__FileRef_Updater) And_Width_LE_Filtering(val int) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And width <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GT_Filtering And > And And_Width_GT_Filtering}

func (d *__FileRef_Updater) And_Width_GT_Filtering(val int) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And width > ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GE_Filtering And >= And And_Width_GE_Filtering}

func (d *__FileRef_Updater) And_Width_GE_Filtering(val int) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And width >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_Filtering Or = Or Or_Width_Eq_Filtering}

func (d *__FileRef_Updater) Or_Width_Eq_Filtering(val int) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or width = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LT_Filtering Or < Or Or_Width_LT_Filtering}

func (d *__FileRef_Updater) Or_Width_LT_Filtering(val int) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or width < ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LE_Filtering Or <= Or Or_Width_LE_Filtering}

func (d *__FileRef_Updater) Or_Width_LE_Filtering(val int) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or width <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GT_Filtering Or > Or Or_Width_GT_Filtering}

func (d *__FileRef_Updater) Or_Width_GT_Filtering(val int) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or width > ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GE_Filtering Or >= Or Or_Width_GE_Filtering}

func (d *__FileRef_Updater) Or_Width_GE_Filtering(val int) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or width >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_FILTERING  =  Extension_Eq_FILTERING}

func (d *__FileRef_Selector) Extension_Eq_FILTERING(val string) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " extension = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_FILTERING And = And And_Extension_Eq_FILTERING}

func (d *__FileRef_Selector) And_Extension_Eq_FILTERING(val string) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And extension = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_FILTERING Or = Or Or_Extension_Eq_FILTERING}

func (d *__FileRef_Selector) Or_Extension_Eq_FILTERING(val string) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or extension = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq  =  FileRefId_Eq}

func (d *__FileRef_Selector) FileRefId_Eq(val int) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " file_ref_id = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LT_Filtering  <  FileRefId_LT_Filtering}

func (d *__FileRef_Selector) FileRefId_LT_Filtering(val int) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " file_ref_id < ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LE_Filtering  <=  FileRefId_LE_Filtering}

func (d *__FileRef_Selector) FileRefId_LE_Filtering(val int) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " file_ref_id <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GT_Filtering  >  FileRefId_GT_Filtering}

func (d *__FileRef_Selector) FileRefId_GT_Filtering(val int) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " file_ref_id > ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GE_Filtering  >=  FileRefId_GE_Filtering}

func (d *__FileRef_Selector) FileRefId_GE_Filtering(val int) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " file_ref_id >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq And = And And_FileRefId_Eq}

func (d *__FileRef_Selector) And_FileRefId_Eq(val int) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And file_ref_id = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LT_Filtering And < And And_FileRefId_LT_Filtering}

func (d *__FileRef_Selector) And_FileRefId_LT_Filtering(val int) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And file_ref_id < ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LE_Filtering And <= And And_FileRefId_LE_Filtering}

func (d *__FileRef_Selector) And_FileRefId_LE_Filtering(val int) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And file_ref_id <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GT_Filtering And > And And_FileRefId_GT_Filtering}

func (d *__FileRef_Selector) And_FileRefId_GT_Filtering(val int) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And file_ref_id > ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GE_Filtering And >= And And_FileRefId_GE_Filtering}

func (d *__FileRef_Selector) And_FileRefId_GE_Filtering(val int) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And file_ref_id >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq Or = Or Or_FileRefId_Eq}

func (d *__FileRef_Selector) Or_FileRefId_Eq(val int) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or file_ref_id = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LT_Filtering Or < Or Or_FileRefId_LT_Filtering}

func (d *__FileRef_Selector) Or_FileRefId_LT_Filtering(val int) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or file_ref_id < ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LE_Filtering Or <= Or Or_FileRefId_LE_Filtering}

func (d *__FileRef_Selector) Or_FileRefId_LE_Filtering(val int) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or file_ref_id <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GT_Filtering Or > Or Or_FileRefId_GT_Filtering}

func (d *__FileRef_Selector) Or_FileRefId_GT_Filtering(val int) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or file_ref_id > ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GE_Filtering Or >= Or Or_FileRefId_GE_Filtering}

func (d *__FileRef_Selector) Or_FileRefId_GE_Filtering(val int) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or file_ref_id >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_Filtering  =  FirstFileRefId_Eq_Filtering}

func (d *__FileRef_Selector) FirstFileRefId_Eq_Filtering(val int) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " first_file_ref_id = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LT_Filtering  <  FirstFileRefId_LT_Filtering}

func (d *__FileRef_Selector) FirstFileRefId_LT_Filtering(val int) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " first_file_ref_id < ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LE_Filtering  <=  FirstFileRefId_LE_Filtering}

func (d *__FileRef_Selector) FirstFileRefId_LE_Filtering(val int) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " first_file_ref_id <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GT_Filtering  >  FirstFileRefId_GT_Filtering}

func (d *__FileRef_Selector) FirstFileRefId_GT_Filtering(val int) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " first_file_ref_id > ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GE_Filtering  >=  FirstFileRefId_GE_Filtering}

func (d *__FileRef_Selector) FirstFileRefId_GE_Filtering(val int) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " first_file_ref_id >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_Filtering And = And And_FirstFileRefId_Eq_Filtering}

func (d *__FileRef_Selector) And_FirstFileRefId_Eq_Filtering(val int) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And first_file_ref_id = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LT_Filtering And < And And_FirstFileRefId_LT_Filtering}

func (d *__FileRef_Selector) And_FirstFileRefId_LT_Filtering(val int) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And first_file_ref_id < ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LE_Filtering And <= And And_FirstFileRefId_LE_Filtering}

func (d *__FileRef_Selector) And_FirstFileRefId_LE_Filtering(val int) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And first_file_ref_id <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GT_Filtering And > And And_FirstFileRefId_GT_Filtering}

func (d *__FileRef_Selector) And_FirstFileRefId_GT_Filtering(val int) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And first_file_ref_id > ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GE_Filtering And >= And And_FirstFileRefId_GE_Filtering}

func (d *__FileRef_Selector) And_FirstFileRefId_GE_Filtering(val int) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And first_file_ref_id >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_Filtering Or = Or Or_FirstFileRefId_Eq_Filtering}

func (d *__FileRef_Selector) Or_FirstFileRefId_Eq_Filtering(val int) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or first_file_ref_id = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LT_Filtering Or < Or Or_FirstFileRefId_LT_Filtering}

func (d *__FileRef_Selector) Or_FirstFileRefId_LT_Filtering(val int) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or first_file_ref_id < ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LE_Filtering Or <= Or Or_FirstFileRefId_LE_Filtering}

func (d *__FileRef_Selector) Or_FirstFileRefId_LE_Filtering(val int) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or first_file_ref_id <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GT_Filtering Or > Or Or_FirstFileRefId_GT_Filtering}

func (d *__FileRef_Selector) Or_FirstFileRefId_GT_Filtering(val int) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or first_file_ref_id > ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GE_Filtering Or >= Or Or_FirstFileRefId_GE_Filtering}

func (d *__FileRef_Selector) Or_FirstFileRefId_GE_Filtering(val int) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or first_file_ref_id >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_Filtering  =  Height_Eq_Filtering}

func (d *__FileRef_Selector) Height_Eq_Filtering(val int) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " height = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LT_Filtering  <  Height_LT_Filtering}

func (d *__FileRef_Selector) Height_LT_Filtering(val int) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " height < ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LE_Filtering  <=  Height_LE_Filtering}

func (d *__FileRef_Selector) Height_LE_Filtering(val int) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " height <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GT_Filtering  >  Height_GT_Filtering}

func (d *__FileRef_Selector) Height_GT_Filtering(val int) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " height > ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GE_Filtering  >=  Height_GE_Filtering}

func (d *__FileRef_Selector) Height_GE_Filtering(val int) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " height >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_Filtering And = And And_Height_Eq_Filtering}

func (d *__FileRef_Selector) And_Height_Eq_Filtering(val int) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And height = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LT_Filtering And < And And_Height_LT_Filtering}

func (d *__FileRef_Selector) And_Height_LT_Filtering(val int) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And height < ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LE_Filtering And <= And And_Height_LE_Filtering}

func (d *__FileRef_Selector) And_Height_LE_Filtering(val int) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And height <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GT_Filtering And > And And_Height_GT_Filtering}

func (d *__FileRef_Selector) And_Height_GT_Filtering(val int) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And height > ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GE_Filtering And >= And And_Height_GE_Filtering}

func (d *__FileRef_Selector) And_Height_GE_Filtering(val int) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And height >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_Filtering Or = Or Or_Height_Eq_Filtering}

func (d *__FileRef_Selector) Or_Height_Eq_Filtering(val int) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or height = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LT_Filtering Or < Or Or_Height_LT_Filtering}

func (d *__FileRef_Selector) Or_Height_LT_Filtering(val int) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or height < ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LE_Filtering Or <= Or Or_Height_LE_Filtering}

func (d *__FileRef_Selector) Or_Height_LE_Filtering(val int) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or height <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GT_Filtering Or > Or Or_Height_GT_Filtering}

func (d *__FileRef_Selector) Or_Height_GT_Filtering(val int) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or height > ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GE_Filtering Or >= Or Or_Height_GE_Filtering}

func (d *__FileRef_Selector) Or_Height_GE_Filtering(val int) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or height >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_Filtering  =  Length_Eq_Filtering}

func (d *__FileRef_Selector) Length_Eq_Filtering(val int) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " length = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LT_Filtering  <  Length_LT_Filtering}

func (d *__FileRef_Selector) Length_LT_Filtering(val int) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " length < ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LE_Filtering  <=  Length_LE_Filtering}

func (d *__FileRef_Selector) Length_LE_Filtering(val int) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " length <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GT_Filtering  >  Length_GT_Filtering}

func (d *__FileRef_Selector) Length_GT_Filtering(val int) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " length > ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GE_Filtering  >=  Length_GE_Filtering}

func (d *__FileRef_Selector) Length_GE_Filtering(val int) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " length >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_Filtering And = And And_Length_Eq_Filtering}

func (d *__FileRef_Selector) And_Length_Eq_Filtering(val int) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And length = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LT_Filtering And < And And_Length_LT_Filtering}

func (d *__FileRef_Selector) And_Length_LT_Filtering(val int) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And length < ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LE_Filtering And <= And And_Length_LE_Filtering}

func (d *__FileRef_Selector) And_Length_LE_Filtering(val int) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And length <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GT_Filtering And > And And_Length_GT_Filtering}

func (d *__FileRef_Selector) And_Length_GT_Filtering(val int) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And length > ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GE_Filtering And >= And And_Length_GE_Filtering}

func (d *__FileRef_Selector) And_Length_GE_Filtering(val int) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And length >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_Filtering Or = Or Or_Length_Eq_Filtering}

func (d *__FileRef_Selector) Or_Length_Eq_Filtering(val int) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or length = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LT_Filtering Or < Or Or_Length_LT_Filtering}

func (d *__FileRef_Selector) Or_Length_LT_Filtering(val int) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or length < ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LE_Filtering Or <= Or Or_Length_LE_Filtering}

func (d *__FileRef_Selector) Or_Length_LE_Filtering(val int) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or length <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GT_Filtering Or > Or Or_Length_GT_Filtering}

func (d *__FileRef_Selector) Or_Length_GT_Filtering(val int) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or length > ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GE_Filtering Or >= Or Or_Length_GE_Filtering}

func (d *__FileRef_Selector) Or_Length_GE_Filtering(val int) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or length >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_FILTERING  =  Md5Key_Eq_FILTERING}

func (d *__FileRef_Selector) Md5Key_Eq_FILTERING(val string) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " md5_key = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_FILTERING And = And And_Md5Key_Eq_FILTERING}

func (d *__FileRef_Selector) And_Md5Key_Eq_FILTERING(val string) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And md5_key = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_FILTERING Or = Or Or_Md5Key_Eq_FILTERING}

func (d *__FileRef_Selector) Or_Md5Key_Eq_FILTERING(val string) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or md5_key = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_FILTERING  =  MimeType_Eq_FILTERING}

func (d *__FileRef_Selector) MimeType_Eq_FILTERING(val string) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " mime_type = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_FILTERING And = And And_MimeType_Eq_FILTERING}

func (d *__FileRef_Selector) And_MimeType_Eq_FILTERING(val string) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And mime_type = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_FILTERING Or = Or Or_MimeType_Eq_FILTERING}

func (d *__FileRef_Selector) Or_MimeType_Eq_FILTERING(val string) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or mime_type = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_FILTERING  =  Name_Eq_FILTERING}

func (d *__FileRef_Selector) Name_Eq_FILTERING(val string) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " name = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_FILTERING And = And And_Name_Eq_FILTERING}

func (d *__FileRef_Selector) And_Name_Eq_FILTERING(val string) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And name = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_FILTERING Or = Or Or_Name_Eq_FILTERING}

func (d *__FileRef_Selector) Or_Name_Eq_FILTERING(val string) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or name = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_Filtering  =  StorageUserId_Eq_Filtering}

func (d *__FileRef_Selector) StorageUserId_Eq_Filtering(val int) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " storage_user_id = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LT_Filtering  <  StorageUserId_LT_Filtering}

func (d *__FileRef_Selector) StorageUserId_LT_Filtering(val int) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " storage_user_id < ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LE_Filtering  <=  StorageUserId_LE_Filtering}

func (d *__FileRef_Selector) StorageUserId_LE_Filtering(val int) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " storage_user_id <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GT_Filtering  >  StorageUserId_GT_Filtering}

func (d *__FileRef_Selector) StorageUserId_GT_Filtering(val int) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " storage_user_id > ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GE_Filtering  >=  StorageUserId_GE_Filtering}

func (d *__FileRef_Selector) StorageUserId_GE_Filtering(val int) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " storage_user_id >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_Filtering And = And And_StorageUserId_Eq_Filtering}

func (d *__FileRef_Selector) And_StorageUserId_Eq_Filtering(val int) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And storage_user_id = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LT_Filtering And < And And_StorageUserId_LT_Filtering}

func (d *__FileRef_Selector) And_StorageUserId_LT_Filtering(val int) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And storage_user_id < ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LE_Filtering And <= And And_StorageUserId_LE_Filtering}

func (d *__FileRef_Selector) And_StorageUserId_LE_Filtering(val int) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And storage_user_id <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GT_Filtering And > And And_StorageUserId_GT_Filtering}

func (d *__FileRef_Selector) And_StorageUserId_GT_Filtering(val int) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And storage_user_id > ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GE_Filtering And >= And And_StorageUserId_GE_Filtering}

func (d *__FileRef_Selector) And_StorageUserId_GE_Filtering(val int) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And storage_user_id >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_Filtering Or = Or Or_StorageUserId_Eq_Filtering}

func (d *__FileRef_Selector) Or_StorageUserId_Eq_Filtering(val int) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or storage_user_id = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LT_Filtering Or < Or Or_StorageUserId_LT_Filtering}

func (d *__FileRef_Selector) Or_StorageUserId_LT_Filtering(val int) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or storage_user_id < ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LE_Filtering Or <= Or Or_StorageUserId_LE_Filtering}

func (d *__FileRef_Selector) Or_StorageUserId_LE_Filtering(val int) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or storage_user_id <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GT_Filtering Or > Or Or_StorageUserId_GT_Filtering}

func (d *__FileRef_Selector) Or_StorageUserId_GT_Filtering(val int) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or storage_user_id > ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GE_Filtering Or >= Or Or_StorageUserId_GE_Filtering}

func (d *__FileRef_Selector) Or_StorageUserId_GE_Filtering(val int) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or storage_user_id >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_FILTERING  =  StoreType_Eq_FILTERING}

func (d *__FileRef_Selector) StoreType_Eq_FILTERING(val string) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " store_type = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_FILTERING And = And And_StoreType_Eq_FILTERING}

func (d *__FileRef_Selector) And_StoreType_Eq_FILTERING(val string) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And store_type = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_FILTERING Or = Or Or_StoreType_Eq_FILTERING}

func (d *__FileRef_Selector) Or_StoreType_Eq_FILTERING(val string) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or store_type = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_Filtering  =  UserId_Eq_Filtering}

func (d *__FileRef_Selector) UserId_Eq_Filtering(val int) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " user_id = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LT_Filtering  <  UserId_LT_Filtering}

func (d *__FileRef_Selector) UserId_LT_Filtering(val int) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " user_id < ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LE_Filtering  <=  UserId_LE_Filtering}

func (d *__FileRef_Selector) UserId_LE_Filtering(val int) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " user_id <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GT_Filtering  >  UserId_GT_Filtering}

func (d *__FileRef_Selector) UserId_GT_Filtering(val int) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " user_id > ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GE_Filtering  >=  UserId_GE_Filtering}

func (d *__FileRef_Selector) UserId_GE_Filtering(val int) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " user_id >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_Filtering And = And And_UserId_Eq_Filtering}

func (d *__FileRef_Selector) And_UserId_Eq_Filtering(val int) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And user_id = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LT_Filtering And < And And_UserId_LT_Filtering}

func (d *__FileRef_Selector) And_UserId_LT_Filtering(val int) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And user_id < ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LE_Filtering And <= And And_UserId_LE_Filtering}

func (d *__FileRef_Selector) And_UserId_LE_Filtering(val int) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And user_id <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GT_Filtering And > And And_UserId_GT_Filtering}

func (d *__FileRef_Selector) And_UserId_GT_Filtering(val int) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And user_id > ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GE_Filtering And >= And And_UserId_GE_Filtering}

func (d *__FileRef_Selector) And_UserId_GE_Filtering(val int) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And user_id >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_Filtering Or = Or Or_UserId_Eq_Filtering}

func (d *__FileRef_Selector) Or_UserId_Eq_Filtering(val int) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or user_id = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LT_Filtering Or < Or Or_UserId_LT_Filtering}

func (d *__FileRef_Selector) Or_UserId_LT_Filtering(val int) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or user_id < ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LE_Filtering Or <= Or Or_UserId_LE_Filtering}

func (d *__FileRef_Selector) Or_UserId_LE_Filtering(val int) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or user_id <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GT_Filtering Or > Or Or_UserId_GT_Filtering}

func (d *__FileRef_Selector) Or_UserId_GT_Filtering(val int) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or user_id > ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GE_Filtering Or >= Or Or_UserId_GE_Filtering}

func (d *__FileRef_Selector) Or_UserId_GE_Filtering(val int) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or user_id >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_Filtering  =  Width_Eq_Filtering}

func (d *__FileRef_Selector) Width_Eq_Filtering(val int) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " width = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LT_Filtering  <  Width_LT_Filtering}

func (d *__FileRef_Selector) Width_LT_Filtering(val int) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " width < ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LE_Filtering  <=  Width_LE_Filtering}

func (d *__FileRef_Selector) Width_LE_Filtering(val int) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " width <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GT_Filtering  >  Width_GT_Filtering}

func (d *__FileRef_Selector) Width_GT_Filtering(val int) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " width > ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GE_Filtering  >=  Width_GE_Filtering}

func (d *__FileRef_Selector) Width_GE_Filtering(val int) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " width >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_Filtering And = And And_Width_Eq_Filtering}

func (d *__FileRef_Selector) And_Width_Eq_Filtering(val int) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And width = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LT_Filtering And < And And_Width_LT_Filtering}

func (d *__FileRef_Selector) And_Width_LT_Filtering(val int) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And width < ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LE_Filtering And <= And And_Width_LE_Filtering}

func (d *__FileRef_Selector) And_Width_LE_Filtering(val int) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And width <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GT_Filtering And > And And_Width_GT_Filtering}

func (d *__FileRef_Selector) And_Width_GT_Filtering(val int) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And width > ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GE_Filtering And >= And And_Width_GE_Filtering}

func (d *__FileRef_Selector) And_Width_GE_Filtering(val int) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And width >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_Filtering Or = Or Or_Width_Eq_Filtering}

func (d *__FileRef_Selector) Or_Width_Eq_Filtering(val int) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or width = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LT_Filtering Or < Or Or_Width_LT_Filtering}

func (d *__FileRef_Selector) Or_Width_LT_Filtering(val int) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or width < ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LE_Filtering Or <= Or Or_Width_LE_Filtering}

func (d *__FileRef_Selector) Or_Width_LE_Filtering(val int) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or width <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GT_Filtering Or > Or Or_Width_GT_Filtering}

func (d *__FileRef_Selector) Or_Width_GT_Filtering(val int) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or width > ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GE_Filtering Or >= Or Or_Width_GE_Filtering}

func (d *__FileRef_Selector) Or_Width_GE_Filtering(val int) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or width >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

///////////////////////////////////////// ins for all //////////////////

func (d *__FileRef_Deleter) Extension_In_FILTERING(val ...string) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = " extension IN (" + helper.DbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__FileRef_Deleter) And_Extension_In_FILTERING(val ...string) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = "And extension IN (" + helper.DbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__FileRef_Deleter) Or_Extension_In_FILTERING(val ...string) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = "Or extension IN (" + helper.DbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__FileRef_Deleter) FileRefId_In(val ...int) *__FileRef_Deleter {
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

func (d *__FileRef_Deleter) And_FileRefId_In(val ...int) *__FileRef_Deleter {
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

func (d *__FileRef_Deleter) Or_FileRefId_In(val ...int) *__FileRef_Deleter {
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

func (d *__FileRef_Deleter) FirstFileRefId_In_FILTERING(val ...int) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = " first_file_ref_id IN (" + helper.DbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__FileRef_Deleter) And_FirstFileRefId_In_FILTERING(val ...int) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = "And first_file_ref_id IN (" + helper.DbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__FileRef_Deleter) Or_FirstFileRefId_In_FILTERING(val ...int) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = "Or first_file_ref_id IN (" + helper.DbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__FileRef_Deleter) Height_In_FILTERING(val ...int) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = " height IN (" + helper.DbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__FileRef_Deleter) And_Height_In_FILTERING(val ...int) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = "And height IN (" + helper.DbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__FileRef_Deleter) Or_Height_In_FILTERING(val ...int) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = "Or height IN (" + helper.DbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__FileRef_Deleter) Length_In_FILTERING(val ...int) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = " length IN (" + helper.DbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__FileRef_Deleter) And_Length_In_FILTERING(val ...int) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = "And length IN (" + helper.DbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__FileRef_Deleter) Or_Length_In_FILTERING(val ...int) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = "Or length IN (" + helper.DbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__FileRef_Deleter) Md5Key_In_FILTERING(val ...string) *__FileRef_Deleter {
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

func (d *__FileRef_Deleter) And_Md5Key_In_FILTERING(val ...string) *__FileRef_Deleter {
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

func (d *__FileRef_Deleter) Or_Md5Key_In_FILTERING(val ...string) *__FileRef_Deleter {
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

func (d *__FileRef_Deleter) MimeType_In_FILTERING(val ...string) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = " mime_type IN (" + helper.DbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__FileRef_Deleter) And_MimeType_In_FILTERING(val ...string) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = "And mime_type IN (" + helper.DbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__FileRef_Deleter) Or_MimeType_In_FILTERING(val ...string) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = "Or mime_type IN (" + helper.DbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__FileRef_Deleter) Name_In_FILTERING(val ...string) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = " name IN (" + helper.DbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__FileRef_Deleter) And_Name_In_FILTERING(val ...string) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = "And name IN (" + helper.DbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__FileRef_Deleter) Or_Name_In_FILTERING(val ...string) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = "Or name IN (" + helper.DbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__FileRef_Deleter) StorageUserId_In_FILTERING(val ...int) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = " storage_user_id IN (" + helper.DbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__FileRef_Deleter) And_StorageUserId_In_FILTERING(val ...int) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = "And storage_user_id IN (" + helper.DbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__FileRef_Deleter) Or_StorageUserId_In_FILTERING(val ...int) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = "Or storage_user_id IN (" + helper.DbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__FileRef_Deleter) StoreType_In_FILTERING(val ...string) *__FileRef_Deleter {
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

func (d *__FileRef_Deleter) And_StoreType_In_FILTERING(val ...string) *__FileRef_Deleter {
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

func (d *__FileRef_Deleter) Or_StoreType_In_FILTERING(val ...string) *__FileRef_Deleter {
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

func (d *__FileRef_Deleter) UserId_In_FILTERING(val ...int) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = " user_id IN (" + helper.DbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__FileRef_Deleter) And_UserId_In_FILTERING(val ...int) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = "And user_id IN (" + helper.DbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__FileRef_Deleter) Or_UserId_In_FILTERING(val ...int) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = "Or user_id IN (" + helper.DbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__FileRef_Deleter) Width_In_FILTERING(val ...int) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = " width IN (" + helper.DbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__FileRef_Deleter) And_Width_In_FILTERING(val ...int) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = "And width IN (" + helper.DbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__FileRef_Deleter) Or_Width_In_FILTERING(val ...int) *__FileRef_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = "Or width IN (" + helper.DbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__FileRef_Updater) Extension_In_FILTERING(val ...string) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = " extension IN (" + helper.DbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__FileRef_Updater) And_Extension_In_FILTERING(val ...string) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = "And extension IN (" + helper.DbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__FileRef_Updater) Or_Extension_In_FILTERING(val ...string) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = "Or extension IN (" + helper.DbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__FileRef_Updater) FileRefId_In(val ...int) *__FileRef_Updater {
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

func (d *__FileRef_Updater) And_FileRefId_In(val ...int) *__FileRef_Updater {
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

func (d *__FileRef_Updater) Or_FileRefId_In(val ...int) *__FileRef_Updater {
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

func (d *__FileRef_Updater) FirstFileRefId_In_FILTERING(val ...int) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = " first_file_ref_id IN (" + helper.DbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__FileRef_Updater) And_FirstFileRefId_In_FILTERING(val ...int) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = "And first_file_ref_id IN (" + helper.DbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__FileRef_Updater) Or_FirstFileRefId_In_FILTERING(val ...int) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = "Or first_file_ref_id IN (" + helper.DbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__FileRef_Updater) Height_In_FILTERING(val ...int) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = " height IN (" + helper.DbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__FileRef_Updater) And_Height_In_FILTERING(val ...int) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = "And height IN (" + helper.DbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__FileRef_Updater) Or_Height_In_FILTERING(val ...int) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = "Or height IN (" + helper.DbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__FileRef_Updater) Length_In_FILTERING(val ...int) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = " length IN (" + helper.DbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__FileRef_Updater) And_Length_In_FILTERING(val ...int) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = "And length IN (" + helper.DbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__FileRef_Updater) Or_Length_In_FILTERING(val ...int) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = "Or length IN (" + helper.DbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__FileRef_Updater) Md5Key_In_FILTERING(val ...string) *__FileRef_Updater {
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

func (d *__FileRef_Updater) And_Md5Key_In_FILTERING(val ...string) *__FileRef_Updater {
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

func (d *__FileRef_Updater) Or_Md5Key_In_FILTERING(val ...string) *__FileRef_Updater {
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

func (d *__FileRef_Updater) MimeType_In_FILTERING(val ...string) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = " mime_type IN (" + helper.DbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__FileRef_Updater) And_MimeType_In_FILTERING(val ...string) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = "And mime_type IN (" + helper.DbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__FileRef_Updater) Or_MimeType_In_FILTERING(val ...string) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = "Or mime_type IN (" + helper.DbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__FileRef_Updater) Name_In_FILTERING(val ...string) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = " name IN (" + helper.DbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__FileRef_Updater) And_Name_In_FILTERING(val ...string) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = "And name IN (" + helper.DbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__FileRef_Updater) Or_Name_In_FILTERING(val ...string) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = "Or name IN (" + helper.DbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__FileRef_Updater) StorageUserId_In_FILTERING(val ...int) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = " storage_user_id IN (" + helper.DbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__FileRef_Updater) And_StorageUserId_In_FILTERING(val ...int) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = "And storage_user_id IN (" + helper.DbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__FileRef_Updater) Or_StorageUserId_In_FILTERING(val ...int) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = "Or storage_user_id IN (" + helper.DbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__FileRef_Updater) StoreType_In_FILTERING(val ...string) *__FileRef_Updater {
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

func (d *__FileRef_Updater) And_StoreType_In_FILTERING(val ...string) *__FileRef_Updater {
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

func (d *__FileRef_Updater) Or_StoreType_In_FILTERING(val ...string) *__FileRef_Updater {
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

func (d *__FileRef_Updater) UserId_In_FILTERING(val ...int) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = " user_id IN (" + helper.DbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__FileRef_Updater) And_UserId_In_FILTERING(val ...int) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = "And user_id IN (" + helper.DbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__FileRef_Updater) Or_UserId_In_FILTERING(val ...int) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = "Or user_id IN (" + helper.DbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__FileRef_Updater) Width_In_FILTERING(val ...int) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = " width IN (" + helper.DbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__FileRef_Updater) And_Width_In_FILTERING(val ...int) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = "And width IN (" + helper.DbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__FileRef_Updater) Or_Width_In_FILTERING(val ...int) *__FileRef_Updater {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = "Or width IN (" + helper.DbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__FileRef_Selector) Extension_In_FILTERING(val ...string) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = " extension IN (" + helper.DbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__FileRef_Selector) And_Extension_In_FILTERING(val ...string) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = "And extension IN (" + helper.DbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__FileRef_Selector) Or_Extension_In_FILTERING(val ...string) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = "Or extension IN (" + helper.DbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__FileRef_Selector) FileRefId_In(val ...int) *__FileRef_Selector {
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

func (d *__FileRef_Selector) And_FileRefId_In(val ...int) *__FileRef_Selector {
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

func (d *__FileRef_Selector) Or_FileRefId_In(val ...int) *__FileRef_Selector {
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

func (d *__FileRef_Selector) FirstFileRefId_In_FILTERING(val ...int) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = " first_file_ref_id IN (" + helper.DbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__FileRef_Selector) And_FirstFileRefId_In_FILTERING(val ...int) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = "And first_file_ref_id IN (" + helper.DbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__FileRef_Selector) Or_FirstFileRefId_In_FILTERING(val ...int) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = "Or first_file_ref_id IN (" + helper.DbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__FileRef_Selector) Height_In_FILTERING(val ...int) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = " height IN (" + helper.DbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__FileRef_Selector) And_Height_In_FILTERING(val ...int) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = "And height IN (" + helper.DbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__FileRef_Selector) Or_Height_In_FILTERING(val ...int) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = "Or height IN (" + helper.DbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__FileRef_Selector) Length_In_FILTERING(val ...int) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = " length IN (" + helper.DbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__FileRef_Selector) And_Length_In_FILTERING(val ...int) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = "And length IN (" + helper.DbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__FileRef_Selector) Or_Length_In_FILTERING(val ...int) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = "Or length IN (" + helper.DbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__FileRef_Selector) Md5Key_In_FILTERING(val ...string) *__FileRef_Selector {
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

func (d *__FileRef_Selector) And_Md5Key_In_FILTERING(val ...string) *__FileRef_Selector {
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

func (d *__FileRef_Selector) Or_Md5Key_In_FILTERING(val ...string) *__FileRef_Selector {
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

func (d *__FileRef_Selector) MimeType_In_FILTERING(val ...string) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = " mime_type IN (" + helper.DbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__FileRef_Selector) And_MimeType_In_FILTERING(val ...string) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = "And mime_type IN (" + helper.DbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__FileRef_Selector) Or_MimeType_In_FILTERING(val ...string) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = "Or mime_type IN (" + helper.DbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__FileRef_Selector) Name_In_FILTERING(val ...string) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = " name IN (" + helper.DbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__FileRef_Selector) And_Name_In_FILTERING(val ...string) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = "And name IN (" + helper.DbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__FileRef_Selector) Or_Name_In_FILTERING(val ...string) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = "Or name IN (" + helper.DbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__FileRef_Selector) StorageUserId_In_FILTERING(val ...int) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = " storage_user_id IN (" + helper.DbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__FileRef_Selector) And_StorageUserId_In_FILTERING(val ...int) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = "And storage_user_id IN (" + helper.DbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__FileRef_Selector) Or_StorageUserId_In_FILTERING(val ...int) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = "Or storage_user_id IN (" + helper.DbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__FileRef_Selector) StoreType_In_FILTERING(val ...string) *__FileRef_Selector {
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

func (d *__FileRef_Selector) And_StoreType_In_FILTERING(val ...string) *__FileRef_Selector {
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

func (d *__FileRef_Selector) Or_StoreType_In_FILTERING(val ...string) *__FileRef_Selector {
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

func (d *__FileRef_Selector) UserId_In_FILTERING(val ...int) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = " user_id IN (" + helper.DbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__FileRef_Selector) And_UserId_In_FILTERING(val ...int) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = "And user_id IN (" + helper.DbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__FileRef_Selector) Or_UserId_In_FILTERING(val ...int) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = "Or user_id IN (" + helper.DbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__FileRef_Selector) Width_In_FILTERING(val ...int) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = " width IN (" + helper.DbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__FileRef_Selector) And_Width_In_FILTERING(val ...int) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = "And width IN (" + helper.DbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__FileRef_Selector) Or_Width_In_FILTERING(val ...int) *__FileRef_Selector {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = "Or width IN (" + helper.DbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

/////////////////////////////////////// End of Ins //////////////////////
///////////////////////////// start of where cluases

/////////////////////////////////////// Start of select //////////////////
func (u *__FileRef_Selector) _toSql() (string, []interface{}) {

	sqlWheres, whereArgs := whereClusesToSql(u.wheres, "")
	selectCols := "*"
	if len(u.selectCol) > 0 {
		selectCols = strings.Join(u.selectCol, ", ")
	}
	sqlstr := "SELECT " + selectCols + " FROM sunc_file.file_ref"

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

func (u *__FileRef_Selector) GetRow(session *gocql.Session) (*FileRef, error) {
	var err error

	u.limit = 1
	sqlstr, whereArgs := u._toSql()

	if LogTableCqlReq.FileRef {
		helper.XCLog(sqlstr, whereArgs)
	}

	query := session.Query(sqlstr, whereArgs...)
	var row *FileRef
	//by Sqlx
	// err = gocqlx.Get(row ,query)
	rows, err := FileRef_Iter(query.Iter(), 1)
	if err != nil {
		if LogTableCqlReq.FileRef {
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

	//OnFileRef_LoadOne(row)

	return row, nil
}

func (u *__FileRef_Selector) GetRows(session *gocql.Session) ([]*FileRef, error) {
	var err error

	sqlstr, whereArgs := u._toSql()

	if LogTableCqlReq.FileRef {
		helper.XCLog(sqlstr, whereArgs)
	}

	query := session.Query(sqlstr, whereArgs...)

	rows, err := FileRef_Iter(query.Iter(), -1)
	if err != nil {
		if LogTableCqlReq.FileRef {
			helper.XCLogErr(err)
		}
		return rows, err
	}

	for i := 0; i < len(rows); i++ {
		rows[i]._exists = true
	}

	// OnFileRef_LoadMany(rows)

	return rows, nil
}

func (u *__FileRef_Updater) Update(session *gocql.Session) error {
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

	sqlstr := `UPDATE sunc_file.file_ref SET ` + sqlUpdate

	if len(strings.Trim(sqlWheres, " ")) > 0 {
		sqlstr += " WHERE " + sqlWheres
	}
	if LogTableCqlReq.FileRef {
		helper.XCLog(sqlstr, allArgs)
	}
	err = session.Query(sqlstr, allArgs...).Exec()
	if err != nil {
		helper.XCLogErr(err)
		return err
	}

	return nil
}

func (d *__FileRef_Deleter) Delete(session *gocql.Session) error {
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

	sqlstr := "DELETE" + delCols + " FROM sunc_file.file_ref WHERE " + wheresStr

	// run query
	if LogTableCqlReq.FileRef {
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
func MassInsert_FileRef(rows []*FileRef, session *gocql.Session) error {
    if len(rows) == 0 {
        return errors.New("rows slice should not be empty - inserted nothing")
    }
    var err error
    ln := len(rows)
    insVals := helper.SqlManyDollars( 12 ,len(rows),true)

    sqlstr := "INSERT INTO sunc_file.file_ref (" +
       " extension,file_ref_id,first_file_ref_id,height,length,md5_key,mime_type,name,storage_user_id,store_type,user_id,width " +
        ") VALUES " + insVals

    // run query
    vals := make([]interface{}, 0, ln*5) //5 fields

    for _, row := range rows {
    		vals = append(vals, row.Extension)
    		vals = append(vals, row.FileRefId)
    		vals = append(vals, row.FirstFileRefId)
    		vals = append(vals, row.Height)
    		vals = append(vals, row.Length)
    		vals = append(vals, row.Md5Key)
    		vals = append(vals, row.MimeType)
    		vals = append(vals, row.Name)
    		vals = append(vals, row.StorageUserId)
    		vals = append(vals, row.StoreType)
    		vals = append(vals, row.UserId)
    		vals = append(vals, row.Width)
    }

    if LogTableCqlReq.FileRef {
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

func (r *FileRef) Save(session *gocql.Session) error {
	var cols []string
	var q []string
	var vals []interface{}

	if r.Extension != "" {
		cols = append(cols, "extension")
		q = append(q, "?")
		vals = append(vals, r.Extension)
	}

	if r.FileRefId != 0 {
		cols = append(cols, "file_ref_id")
		q = append(q, "?")
		vals = append(vals, r.FileRefId)
	}

	if r.FirstFileRefId != 0 {
		cols = append(cols, "first_file_ref_id")
		q = append(q, "?")
		vals = append(vals, r.FirstFileRefId)
	}

	if r.Height != 0 {
		cols = append(cols, "height")
		q = append(q, "?")
		vals = append(vals, r.Height)
	}

	if r.Length != 0 {
		cols = append(cols, "length")
		q = append(q, "?")
		vals = append(vals, r.Length)
	}

	if r.Md5Key != "" {
		cols = append(cols, "md5_key")
		q = append(q, "?")
		vals = append(vals, r.Md5Key)
	}

	if r.MimeType != "" {
		cols = append(cols, "mime_type")
		q = append(q, "?")
		vals = append(vals, r.MimeType)
	}

	if r.Name != "" {
		cols = append(cols, "name")
		q = append(q, "?")
		vals = append(vals, r.Name)
	}

	if r.StorageUserId != 0 {
		cols = append(cols, "storage_user_id")
		q = append(q, "?")
		vals = append(vals, r.StorageUserId)
	}

	if r.StoreType != "" {
		cols = append(cols, "store_type")
		q = append(q, "?")
		vals = append(vals, r.StoreType)
	}

	if r.UserId != 0 {
		cols = append(cols, "user_id")
		q = append(q, "?")
		vals = append(vals, r.UserId)
	}

	if r.Width != 0 {
		cols = append(cols, "width")
		q = append(q, "?")
		vals = append(vals, r.Width)
	}

	if len(cols) == 0 {
		return errors.New("can not insert empty row.")
	}

	colOut := strings.Join(cols, ",")
	qOut := strings.Join(q, ",")
	cql := "insert into sunc_file.file_ref (" + colOut + ") values (" + qOut + ") "

	if LogTableCqlReq.FileRef {
		helper.XCLog(cql, vals)
	}
	err := session.Query(cql, vals...).Exec()
	if err != nil {
		if LogTableCqlReq.FileRef {
			helper.XCLogErr(err)
		}
	}
	r._exists = true
	return err
}

func (r *FileRef) Delete(session *gocql.Session) error {
	var err error
	del := NewFileRef_Deleter()

	del.FileRefId_Eq(r.FileRefId)

	err = del.Delete(session)
	if err != nil {
		return err
	}
	r._exists = false
	return nil
}

func FileRef_Iter(iter *gocql.Iter, limit int) ([]*FileRef, error) {
	var rows []*FileRef
	if limit < 1 {
		limit = 1e6
	}

	for i := 0; i < limit; i++ {
		m := make(map[string]interface{}, 10)
		row := &FileRef{}
		if iter.MapScan(m) {

			if val, ok := m["extension"]; ok {
				row.Extension = string(val.(string))
				//row.Extension = val.(string)
			}

			if val, ok := m["file_ref_id"]; ok {
				row.FileRefId = int(val.(int64))
				//row.FileRefId = val.(int)
			}

			if val, ok := m["first_file_ref_id"]; ok {
				row.FirstFileRefId = int(val.(int64))
				//row.FirstFileRefId = val.(int)
			}

			if val, ok := m["height"]; ok {
				row.Height = int(val.(int))
				//row.Height = val.(int)
			}

			if val, ok := m["length"]; ok {
				row.Length = int(val.(int))
				//row.Length = val.(int)
			}

			if val, ok := m["md5_key"]; ok {
				row.Md5Key = string(val.(string))
				//row.Md5Key = val.(string)
			}

			if val, ok := m["mime_type"]; ok {
				row.MimeType = string(val.(string))
				//row.MimeType = val.(string)
			}

			if val, ok := m["name"]; ok {
				row.Name = string(val.(string))
				//row.Name = val.(string)
			}

			if val, ok := m["storage_user_id"]; ok {
				row.StorageUserId = int(val.(int))
				//row.StorageUserId = val.(int)
			}

			if val, ok := m["store_type"]; ok {
				row.StoreType = string(val.(string))
				//row.StoreType = val.(string)
			}

			if val, ok := m["user_id"]; ok {
				row.UserId = int(val.(int))
				//row.UserId = val.(int)
			}

			if val, ok := m["width"]; ok {
				row.Width = int(val.(int))
				//row.Width = val.(int)
			}

			rows = append(rows, row)
		} else {
			break
		}
	}
	err := iter.Close()

	return rows, err
}
