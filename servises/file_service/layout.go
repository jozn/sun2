package file_service

type Layout interface {
	saveToDataStore()
	getFromDataStore()
	cacheDirectoryName()
}

type Row struct {
	Id        int
	FileType  int
	Extension string
	DataThumb []byte
	Data      []byte
}

type storeLayer interface {
}
