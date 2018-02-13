package file_service

func SavePostFile(row Row) {
	mysqlSavePost(row)
}

func SaveMsgFile(row Row) {
	mysqlSaveMsg(row)
}
