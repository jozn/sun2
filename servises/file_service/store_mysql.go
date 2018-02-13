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
	r.Save(base.DB)
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
