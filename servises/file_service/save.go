package file_service

type Row struct {
	Id        int
	FileType  int
	Extension string
	DataThumb []byte
	Data      []byte
}

func SaveAvatarFile(row Row) {
	mysqlSavePost(row)
}

func SavePostFile(row Row) {
    mysqlSavePost(row)
}

func SaveMsgFile(row Row) {
	mysqlSaveMsg(row)
}
