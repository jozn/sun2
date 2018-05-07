package file_disk_cache

import (
	"ms/sun/servises/file_service/file_common"
	"os"
)

func SelectDisk(req *file_common.RowReq, config *file_common.FileServingConfig) {
	if len(config.DiskDirs) > 0 {
		diskId := req.FileDataStoreId % len(config.DiskDirs)
		diskPath := config.DiskDirs[diskId]
		req.SelectedCacheDiskPath = diskPath
	}
}

func IsLocalCacheAvailable(req *file_common.RowReq, config *file_common.FileServingConfig) (string, bool) {
	for _, parentDir := range config.DiskDirs {
		fullpath := parentDir + req.RowCacheOutRelativePath
		if IsFileExists(fullpath) {
			return fullpath, true
		}
	}
	return "", false
}

func IsFileExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}
