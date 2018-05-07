package file_store

type Row struct {
	Id        int
	FileType  int
	Extension string
	DataThumb []byte
	Data      []byte
}

func SaveAvatarFile(row Row) {
	MysqlSavePost(row)
}

func SavePostFile(row Row) {
    MysqlSavePost(row)
}

func SaveMsgFile(row Row) {
	MysqlSaveMsg(row)
}
