package file_service

var allowedResizingExceptionsMap = map[string]bool{
	".jpeg": true,
	".jpg":  true,
	".png":  true,
	".gif":  true,
	".webp": true,
	".mp4":  true,
}

var allowedThumbExceptionsMap = map[string]bool{
	".jpeg": true,
	".jpg":  true,
	".png":  true,
	".gif":  true,
	".webp": true,
}
