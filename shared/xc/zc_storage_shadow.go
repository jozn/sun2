package xc

import (
	"errors"
	"ms/sun/shared/helper"
	"strconv"
	"strings"

	"github.com/gocql/gocql"
)

////////////////////////////////////////// Query seletor updater and deleter /////////////////////////

func (a *StorageShadow) Exists() bool {
	return a._exists
}

func (a *StorageShadow) Deleted() bool {
	return a._deleted
}

type __StorageShadow_Selector struct {
	wheres      []whereClause
	selectCol   []string
	orderBy     []string //" order by id desc //for ints
	limit       int
	allowFilter bool
}

type __StorageShadow_Updater struct {
	wheres  []whereClause
	updates map[string]interface{}
}

type __StorageShadow_Deleter struct {
	wheres    []whereClause
	deleteCol []string
}

//////////////////// just Selector
func (u *__StorageShadow_Selector) Limit(limit int) *__StorageShadow_Selector {
	u.limit = limit
	return u
}

func (u *__StorageShadow_Selector) AllowFiltering() *__StorageShadow_Selector {
	u.allowFilter = true
	return u
}

func NewStorageShadow_Selector() *__StorageShadow_Selector {
	u := __StorageShadow_Selector{}
	return &u
}

func NewStorageShadow_Updater() *__StorageShadow_Updater {
	u := __StorageShadow_Updater{}
	u.updates = make(map[string]interface{})
	return &u
}

func NewStorageShadow_Deleter() *__StorageShadow_Deleter {
	u := __StorageShadow_Deleter{}
	return &u
}

//each select columns

func (u *__StorageShadow_Selector) Select_MimeType() *__StorageShadow_Selector {
	u.selectCol = append(u.selectCol, "mime_type")
	return u
}

//each column orders //just ints
func (u *__StorageShadow_Selector) OrderBy_MimeType_Desc() *__StorageShadow_Selector {
	u.orderBy = append(u.orderBy, " mime_type DESC")
	return u
}

func (u *__StorageShadow_Selector) OrderBy_MimeType_Asc() *__StorageShadow_Selector {
	u.orderBy = append(u.orderBy, " mime_type ASC")
	return u
}

func (u *__StorageShadow_Selector) Select_ShadowMd5Key() *__StorageShadow_Selector {
	u.selectCol = append(u.selectCol, "shadow_md5_key")
	return u
}

//each column orders //just ints
func (u *__StorageShadow_Selector) OrderBy_ShadowMd5Key_Desc() *__StorageShadow_Selector {
	u.orderBy = append(u.orderBy, " shadow_md5_key DESC")
	return u
}

func (u *__StorageShadow_Selector) OrderBy_ShadowMd5Key_Asc() *__StorageShadow_Selector {
	u.orderBy = append(u.orderBy, " shadow_md5_key ASC")
	return u
}

func (u *__StorageShadow_Selector) Select_StorageMd5Key() *__StorageShadow_Selector {
	u.selectCol = append(u.selectCol, "storage_md5_key")
	return u
}

//each column orders //just ints
func (u *__StorageShadow_Selector) OrderBy_StorageMd5Key_Desc() *__StorageShadow_Selector {
	u.orderBy = append(u.orderBy, " storage_md5_key DESC")
	return u
}

func (u *__StorageShadow_Selector) OrderBy_StorageMd5Key_Asc() *__StorageShadow_Selector {
	u.orderBy = append(u.orderBy, " storage_md5_key ASC")
	return u
}

//////////////////// just Deleter
//each column delete

func (u *__StorageShadow_Deleter) Delete_MimeType() *__StorageShadow_Deleter {
	u.deleteCol = append(u.deleteCol, "mime_type")
	return u
}

func (u *__StorageShadow_Deleter) Delete_ShadowMd5Key() *__StorageShadow_Deleter {
	u.deleteCol = append(u.deleteCol, "shadow_md5_key")
	return u
}

func (u *__StorageShadow_Deleter) Delete_StorageMd5Key() *__StorageShadow_Deleter {
	u.deleteCol = append(u.deleteCol, "storage_md5_key")
	return u
}

//////////////////// End of just Deleter

//////////////////// just Updater
//each column delete

func (u *__StorageShadow_Updater) MimeType(newVal string) *__StorageShadow_Updater {
	u.updates["mime_type = ? "] = newVal
	return u
}

func (u *__StorageShadow_Updater) ShadowMd5Key(newVal string) *__StorageShadow_Updater {
	u.updates["shadow_md5_key = ? "] = newVal
	return u
}

func (u *__StorageShadow_Updater) StorageMd5Key(newVal string) *__StorageShadow_Updater {
	u.updates["storage_md5_key = ? "] = newVal
	return u
}

//////////////////// End just Updater

//{_Eq_FILTERING  =  MimeType_Eq_FILTERING}

func (d *__StorageShadow_Deleter) MimeType_Eq_FILTERING(val string) *__StorageShadow_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " mime_type = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_FILTERING And = And And_MimeType_Eq_FILTERING}

func (d *__StorageShadow_Deleter) And_MimeType_Eq_FILTERING(val string) *__StorageShadow_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And mime_type = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_FILTERING Or = Or Or_MimeType_Eq_FILTERING}

func (d *__StorageShadow_Deleter) Or_MimeType_Eq_FILTERING(val string) *__StorageShadow_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or mime_type = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq  =  ShadowMd5Key_Eq}

func (d *__StorageShadow_Deleter) ShadowMd5Key_Eq(val string) *__StorageShadow_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " shadow_md5_key = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq And = And And_ShadowMd5Key_Eq}

func (d *__StorageShadow_Deleter) And_ShadowMd5Key_Eq(val string) *__StorageShadow_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And shadow_md5_key = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq Or = Or Or_ShadowMd5Key_Eq}

func (d *__StorageShadow_Deleter) Or_ShadowMd5Key_Eq(val string) *__StorageShadow_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or shadow_md5_key = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_FILTERING  =  StorageMd5Key_Eq_FILTERING}

func (d *__StorageShadow_Deleter) StorageMd5Key_Eq_FILTERING(val string) *__StorageShadow_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " storage_md5_key = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_FILTERING And = And And_StorageMd5Key_Eq_FILTERING}

func (d *__StorageShadow_Deleter) And_StorageMd5Key_Eq_FILTERING(val string) *__StorageShadow_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And storage_md5_key = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_FILTERING Or = Or Or_StorageMd5Key_Eq_FILTERING}

func (d *__StorageShadow_Deleter) Or_StorageMd5Key_Eq_FILTERING(val string) *__StorageShadow_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or storage_md5_key = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_FILTERING  =  MimeType_Eq_FILTERING}

func (d *__StorageShadow_Updater) MimeType_Eq_FILTERING(val string) *__StorageShadow_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " mime_type = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_FILTERING And = And And_MimeType_Eq_FILTERING}

func (d *__StorageShadow_Updater) And_MimeType_Eq_FILTERING(val string) *__StorageShadow_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And mime_type = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_FILTERING Or = Or Or_MimeType_Eq_FILTERING}

func (d *__StorageShadow_Updater) Or_MimeType_Eq_FILTERING(val string) *__StorageShadow_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or mime_type = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq  =  ShadowMd5Key_Eq}

func (d *__StorageShadow_Updater) ShadowMd5Key_Eq(val string) *__StorageShadow_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " shadow_md5_key = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq And = And And_ShadowMd5Key_Eq}

func (d *__StorageShadow_Updater) And_ShadowMd5Key_Eq(val string) *__StorageShadow_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And shadow_md5_key = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq Or = Or Or_ShadowMd5Key_Eq}

func (d *__StorageShadow_Updater) Or_ShadowMd5Key_Eq(val string) *__StorageShadow_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or shadow_md5_key = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_FILTERING  =  StorageMd5Key_Eq_FILTERING}

func (d *__StorageShadow_Updater) StorageMd5Key_Eq_FILTERING(val string) *__StorageShadow_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " storage_md5_key = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_FILTERING And = And And_StorageMd5Key_Eq_FILTERING}

func (d *__StorageShadow_Updater) And_StorageMd5Key_Eq_FILTERING(val string) *__StorageShadow_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And storage_md5_key = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_FILTERING Or = Or Or_StorageMd5Key_Eq_FILTERING}

func (d *__StorageShadow_Updater) Or_StorageMd5Key_Eq_FILTERING(val string) *__StorageShadow_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or storage_md5_key = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_FILTERING  =  MimeType_Eq_FILTERING}

func (d *__StorageShadow_Selector) MimeType_Eq_FILTERING(val string) *__StorageShadow_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " mime_type = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_FILTERING And = And And_MimeType_Eq_FILTERING}

func (d *__StorageShadow_Selector) And_MimeType_Eq_FILTERING(val string) *__StorageShadow_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And mime_type = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_FILTERING Or = Or Or_MimeType_Eq_FILTERING}

func (d *__StorageShadow_Selector) Or_MimeType_Eq_FILTERING(val string) *__StorageShadow_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or mime_type = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq  =  ShadowMd5Key_Eq}

func (d *__StorageShadow_Selector) ShadowMd5Key_Eq(val string) *__StorageShadow_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " shadow_md5_key = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq And = And And_ShadowMd5Key_Eq}

func (d *__StorageShadow_Selector) And_ShadowMd5Key_Eq(val string) *__StorageShadow_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And shadow_md5_key = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq Or = Or Or_ShadowMd5Key_Eq}

func (d *__StorageShadow_Selector) Or_ShadowMd5Key_Eq(val string) *__StorageShadow_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or shadow_md5_key = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_FILTERING  =  StorageMd5Key_Eq_FILTERING}

func (d *__StorageShadow_Selector) StorageMd5Key_Eq_FILTERING(val string) *__StorageShadow_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " storage_md5_key = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_FILTERING And = And And_StorageMd5Key_Eq_FILTERING}

func (d *__StorageShadow_Selector) And_StorageMd5Key_Eq_FILTERING(val string) *__StorageShadow_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And storage_md5_key = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_FILTERING Or = Or Or_StorageMd5Key_Eq_FILTERING}

func (d *__StorageShadow_Selector) Or_StorageMd5Key_Eq_FILTERING(val string) *__StorageShadow_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or storage_md5_key = ? "
	d.wheres = append(d.wheres, w)

	return d
}

///////////////////////////////////////// ins for all //////////////////

func (d *__StorageShadow_Deleter) MimeType_In_FILTERING(val ...string) *__StorageShadow_Deleter {
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

func (d *__StorageShadow_Deleter) And_MimeType_In_FILTERING(val ...string) *__StorageShadow_Deleter {
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

func (d *__StorageShadow_Deleter) Or_MimeType_In_FILTERING(val ...string) *__StorageShadow_Deleter {
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

func (d *__StorageShadow_Deleter) ShadowMd5Key_In(val ...string) *__StorageShadow_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = " shadow_md5_key IN (" + helper.DbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__StorageShadow_Deleter) And_ShadowMd5Key_In(val ...string) *__StorageShadow_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = "And shadow_md5_key IN (" + helper.DbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__StorageShadow_Deleter) Or_ShadowMd5Key_In(val ...string) *__StorageShadow_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = "Or shadow_md5_key IN (" + helper.DbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__StorageShadow_Deleter) StorageMd5Key_In_FILTERING(val ...string) *__StorageShadow_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = " storage_md5_key IN (" + helper.DbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__StorageShadow_Deleter) And_StorageMd5Key_In_FILTERING(val ...string) *__StorageShadow_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = "And storage_md5_key IN (" + helper.DbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__StorageShadow_Deleter) Or_StorageMd5Key_In_FILTERING(val ...string) *__StorageShadow_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = "Or storage_md5_key IN (" + helper.DbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__StorageShadow_Updater) MimeType_In_FILTERING(val ...string) *__StorageShadow_Updater {
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

func (d *__StorageShadow_Updater) And_MimeType_In_FILTERING(val ...string) *__StorageShadow_Updater {
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

func (d *__StorageShadow_Updater) Or_MimeType_In_FILTERING(val ...string) *__StorageShadow_Updater {
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

func (d *__StorageShadow_Updater) ShadowMd5Key_In(val ...string) *__StorageShadow_Updater {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = " shadow_md5_key IN (" + helper.DbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__StorageShadow_Updater) And_ShadowMd5Key_In(val ...string) *__StorageShadow_Updater {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = "And shadow_md5_key IN (" + helper.DbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__StorageShadow_Updater) Or_ShadowMd5Key_In(val ...string) *__StorageShadow_Updater {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = "Or shadow_md5_key IN (" + helper.DbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__StorageShadow_Updater) StorageMd5Key_In_FILTERING(val ...string) *__StorageShadow_Updater {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = " storage_md5_key IN (" + helper.DbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__StorageShadow_Updater) And_StorageMd5Key_In_FILTERING(val ...string) *__StorageShadow_Updater {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = "And storage_md5_key IN (" + helper.DbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__StorageShadow_Updater) Or_StorageMd5Key_In_FILTERING(val ...string) *__StorageShadow_Updater {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = "Or storage_md5_key IN (" + helper.DbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__StorageShadow_Selector) MimeType_In_FILTERING(val ...string) *__StorageShadow_Selector {
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

func (d *__StorageShadow_Selector) And_MimeType_In_FILTERING(val ...string) *__StorageShadow_Selector {
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

func (d *__StorageShadow_Selector) Or_MimeType_In_FILTERING(val ...string) *__StorageShadow_Selector {
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

func (d *__StorageShadow_Selector) ShadowMd5Key_In(val ...string) *__StorageShadow_Selector {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = " shadow_md5_key IN (" + helper.DbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__StorageShadow_Selector) And_ShadowMd5Key_In(val ...string) *__StorageShadow_Selector {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = "And shadow_md5_key IN (" + helper.DbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__StorageShadow_Selector) Or_ShadowMd5Key_In(val ...string) *__StorageShadow_Selector {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = "Or shadow_md5_key IN (" + helper.DbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__StorageShadow_Selector) StorageMd5Key_In_FILTERING(val ...string) *__StorageShadow_Selector {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = " storage_md5_key IN (" + helper.DbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__StorageShadow_Selector) And_StorageMd5Key_In_FILTERING(val ...string) *__StorageShadow_Selector {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = "And storage_md5_key IN (" + helper.DbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__StorageShadow_Selector) Or_StorageMd5Key_In_FILTERING(val ...string) *__StorageShadow_Selector {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = "Or storage_md5_key IN (" + helper.DbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

/////////////////////////////////////// End of Ins //////////////////////
///////////////////////////// start of where cluases

/////////////////////////////////////// Start of select //////////////////
func (u *__StorageShadow_Selector) _toSql() (string, []interface{}) {

	sqlWheres, whereArgs := whereClusesToSql(u.wheres, "")
	selectCols := "*"
	if len(u.selectCol) > 0 {
		selectCols = strings.Join(u.selectCol, ", ")
	}
	sqlstr := "SELECT " + selectCols + " FROM sunc_file.storage_shadow"

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

func (u *__StorageShadow_Selector) GetRow(session *gocql.Session) (*StorageShadow, error) {
	var err error

	u.limit = 1
	sqlstr, whereArgs := u._toSql()

	if LogTableCqlReq.StorageShadow {
		helper.XCLog(sqlstr, whereArgs)
	}

	query := session.Query(sqlstr, whereArgs...)
	var row *StorageShadow
	//by Sqlx
	// err = gocqlx.Get(row ,query)
	rows, err := StorageShadow_Iter(query.Iter(), 1)
	if err != nil {
		if LogTableCqlReq.StorageShadow {
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

	//OnStorageShadow_LoadOne(row)

	return row, nil
}

func (u *__StorageShadow_Selector) GetRows(session *gocql.Session) ([]*StorageShadow, error) {
	var err error

	sqlstr, whereArgs := u._toSql()

	if LogTableCqlReq.StorageShadow {
		helper.XCLog(sqlstr, whereArgs)
	}

	query := session.Query(sqlstr, whereArgs...)

	rows, err := StorageShadow_Iter(query.Iter(), -1)
	if err != nil {
		if LogTableCqlReq.StorageShadow {
			helper.XCLogErr(err)
		}
		return rows, err
	}

	for i := 0; i < len(rows); i++ {
		rows[i]._exists = true
	}

	// OnStorageShadow_LoadMany(rows)

	return rows, nil
}

func (u *__StorageShadow_Updater) Update(session *gocql.Session) error {
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

	sqlstr := `UPDATE sunc_file.storage_shadow SET ` + sqlUpdate

	if len(strings.Trim(sqlWheres, " ")) > 0 {
		sqlstr += " WHERE " + sqlWheres
	}
	if LogTableCqlReq.StorageShadow {
		helper.XCLog(sqlstr, allArgs)
	}
	err = session.Query(sqlstr, allArgs...).Exec()
	if err != nil {
		helper.XCLogErr(err)
		return err
	}

	return nil
}

func (d *__StorageShadow_Deleter) Delete(session *gocql.Session) error {
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

	sqlstr := "DELETE" + delCols + " FROM sunc_file.storage_shadow WHERE " + wheresStr

	// run query
	if LogTableCqlReq.StorageShadow {
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
func MassInsert_StorageShadow(rows []*StorageShadow, session *gocql.Session) error {
    if len(rows) == 0 {
        return errors.New("rows slice should not be empty - inserted nothing")
    }
    var err error
    ln := len(rows)
    insVals := helper.SqlManyDollars( 3 ,len(rows),true)

    sqlstr := "INSERT INTO sunc_file.storage_shadow (" +
       " mime_type,shadow_md5_key,storage_md5_key " +
        ") VALUES " + insVals

    // run query
    vals := make([]interface{}, 0, ln*5) //5 fields

    for _, row := range rows {
    		vals = append(vals, row.MimeType)
    		vals = append(vals, row.ShadowMd5Key)
    		vals = append(vals, row.StorageMd5Key)
    }

    if LogTableCqlReq.StorageShadow {
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

func (r *StorageShadow) Save(session *gocql.Session) error {
	var cols []string
	var q []string
	var vals []interface{}

	if r.MimeType != "" {
		cols = append(cols, "mime_type")
		q = append(q, "?")
		vals = append(vals, r.MimeType)
	}

	if r.ShadowMd5Key != "" {
		cols = append(cols, "shadow_md5_key")
		q = append(q, "?")
		vals = append(vals, r.ShadowMd5Key)
	}

	if r.StorageMd5Key != "" {
		cols = append(cols, "storage_md5_key")
		q = append(q, "?")
		vals = append(vals, r.StorageMd5Key)
	}

	if len(cols) == 0 {
		return errors.New("can not insert empty row.")
	}

	colOut := strings.Join(cols, ",")
	qOut := strings.Join(q, ",")
	cql := "insert into sunc_file.storage_shadow (" + colOut + ") values (" + qOut + ") "

	if LogTableCqlReq.StorageShadow {
		helper.XCLog(cql, vals)
	}
	err := session.Query(cql, vals...).Exec()
	if err != nil {
		if LogTableCqlReq.StorageShadow {
			helper.XCLogErr(err)
		}
	}
	r._exists = true
	return err
}

func (r *StorageShadow) Delete(session *gocql.Session) error {
	var err error
	del := NewStorageShadow_Deleter()

	del.ShadowMd5Key_Eq(r.ShadowMd5Key)

	err = del.Delete(session)
	if err != nil {
		return err
	}
	r._exists = false
	return nil
}

func StorageShadow_Iter(iter *gocql.Iter, limit int) ([]*StorageShadow, error) {
	var rows []*StorageShadow
	if limit < 1 {
		limit = 1e6
	}

	for i := 0; i < limit; i++ {
		m := make(map[string]interface{}, 10)
		row := &StorageShadow{}
		if iter.MapScan(m) {

			if val, ok := m["mime_type"]; ok {
				row.MimeType = string(val.(string))
				//row.MimeType = val.(string)
			}

			if val, ok := m["shadow_md5_key"]; ok {
				row.ShadowMd5Key = string(val.(string))
				//row.ShadowMd5Key = val.(string)
			}

			if val, ok := m["storage_md5_key"]; ok {
				row.StorageMd5Key = string(val.(string))
				//row.StorageMd5Key = val.(string)
			}

			rows = append(rows, row)
		} else {
			break
		}
	}
	err := iter.Close()

	return rows, err
}
