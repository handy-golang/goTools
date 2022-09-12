package mFile

func ContentToExtName(lType string) string {
	ext := "avi"
	switch lType {

	case "image/bmp":
		ext = "bmp"

	case "image/gif":
		ext = "gif"

	case "image/jpeg":
		ext = "jpeg"
	case "image/png":
		ext = "png"

	case "text/html":
		ext = "html"

	case "text/plain":
		ext = "txt"

	case "application/vnd.visio":
		ext = "vsd"

	case "application/vnd.ms-powerpoint":
		ext = "pptx"

	case "application/msword":
		ext = "docx"

	case "application/msexcel":
		ext = "xlsx"

	case "application/csv":
		ext = "csv"

	case "text/xml":
		ext = "xml"

	case "video/mp4":
		ext = "mp4"

	case "video/x-msvideo":
		ext = "avi"

	case "video/quicktime":
		ext = "mov"

	case "video/mpeg":
		ext = "mpeg"

	case "video/x-ms-wmv":
		ext = "wm"

	case "video/x-flv":
		ext = "flv"

	case "video/x-matroska":
		ext = "mkv"

	}

	return ext
}
