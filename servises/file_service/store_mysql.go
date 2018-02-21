package file_service

import (
	"ms/sun/base"
	"ms/sun2/shared/x"
)

func mysqlSaveMsg(row Row) {
	r := &x.FileMsg{
		Id:        row.Id,
		FileType:  row.FileType,
		Extension: row.Extension,
		DataThumb: row.DataThumb,
		Data:      row.Data,
	}
    if r.DataThumb == nil {
        r.DataThumb = []byte{}
    }
	r.Save(base.DB)
}

func mysqlSavePost(row Row) {
	r := &x.FilePost{
		Id:        row.Id,
		FileType:  row.FileType,
		Extension: row.Extension,
		DataThumb: row.DataThumb,
		Data:      row.Data,
	}
	if r.DataThumb == nil {
		r.DataThumb = []byte{}
	}
	err := r.Save(base.DB)
	if err != nil {
		panic(err)
	}
}

func mysqlGetmsgFile(id int) (*Row, error) {
	r, err := x.FileMsgById(base.DB, id)
	if err != nil {
		return nil, err
	}
	row := &Row{
		Id:        r.Id,
		FileType:  r.FileType,
		Extension: r.Extension,
		DataThumb: r.DataThumb,
		Data:      r.Data,
	}
	if r.DataThumb == nil {
		r.DataThumb = []byte{}
	}
	return row, nil
}

func mysqlGetPostFile(id int) (*Row, error) {
	r, err := x.FilePostById(base.DB, id)
	if err != nil {
		return nil, err
	}
	row := &Row{
		Id:        r.Id,
		FileType:  r.FileType,
		Extension: r.Extension,
		DataThumb: r.DataThumb,
		Data:      r.Data,
	}
	return row, nil
}
