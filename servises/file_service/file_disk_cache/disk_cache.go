package file_disk_cache

import (
	"ms/sun/servises/file_service/file_common"
	"os"
)

func SelectDisk(req *file_common.RowReq, config *file_common.FileServingConfig) {
	if len(config.DiskDirs) > 0 {
        diskId := req.FileDataStoreId % len(config.DiskDirs)
        diskPath := config.DiskDirs[diskId]

        isOriginalExists := false
        //check if we have the file in others disks, if we have choose that disk
        for _, parentDir := range config.DiskDirs {
            fullpath := parentDir + req.RowCacheOutRelativePath
            if IsFileExists(fullpath) {
                diskPath = parentDir
                isOriginalExists = true
            }
        }

		req.SelectedCacheDiskPath = diskPath
		req.RowCacheOutDiskPathDir = diskPath + req.RowCacheOutRelativePathDir
		req.RowCacheOutDiskPath = diskPath + req.RowCacheOutRelativePath
		req.RowCacheOutDiskPathSized = diskPath + req.RowCacheOutRelativePathSized
		req.RowCacheOutDiskPathThumb = diskPath + req.RowCacheOutRelativePathThumb

		req.IsLocalDiskCacheAvailable = isOriginalExists
	}
}

//dep
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
