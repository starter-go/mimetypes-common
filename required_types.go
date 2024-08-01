package mimetypescommon

import (
	"sort"
	"strings"

	"github.com/starter-go/mimetypes"
)

// GetRequiredTypeNameList 取类型名称列表
func GetRequiredTypeNameList() []mimetypes.Type {
	loader := &myRequiredListLoader{}
	return loader.loadTypeNameList()
}

// GetRequiredFileNameSuffixList 取文件名后缀列表
func GetRequiredFileNameSuffixList() []string {
	loader := &myRequiredListLoader{}
	return loader.loadSuffixList()
}

////////////////////////////////////////////////////////////////////////////////

type myRequiredListLoader struct {
}

func (inst *myRequiredListLoader) loadTypeNameList() []mimetypes.Type {
	src := inst.loadTypeTable()
	mid := make([]string, 0)
	dst := make([]mimetypes.Type, 0)
	for _, name := range src {
		name = strings.TrimSpace(name)
		name = strings.ToLower(name)
		mid = append(mid, name)
	}
	sort.Strings(mid)
	for _, name := range mid {
		dst = append(dst, mimetypes.Type(name))
	}
	return dst
}

func (inst *myRequiredListLoader) loadSuffixList() []string {
	src := inst.loadTypeTable()
	dst := make([]string, 0)
	for suffix := range src {
		suffix = strings.TrimSpace(suffix)
		suffix = strings.ToLower(suffix)
		if !strings.HasPrefix(suffix, ".") {
			suffix = "." + suffix
		}
		dst = append(dst, suffix)
	}
	sort.Strings(dst)
	return dst
}

func (inst *myRequiredListLoader) loadTypeTable() map[string]string {

	tt := map[string]string{

		// application
		".unknown": "application/octet-stream",
		".json":    "application/json",
		".gz":      "application/gzip",
		".pdf":     "application/pdf",
		".zip":     "application/zip",

		// audio
		".aac":  "audio/aac",
		".mid":  "audio/midi",
		".midi": "audio/midi",
		".mp3":  "audio/mpeg",
		".oga":  "audio/ogg",
		".wav":  "audio/wav",
		".weba": "audio/webm",

		// image
		".png":  "image/png",
		".jpg":  "image/jpeg",
		".jpeg": "image/jpeg",
		".gif":  "image/gif",
		".ico":  "image/vnd.microsoft.icon",
		".svg":  "image/svg+xml",
		".bmp":  "image/bmp",
		".tif":  "image/tiff",
		".tiff": "image/tiff",
		".webp": "image/webp",

		// text
		".htm":  "text/html",
		".html": "text/html",
		".txt":  "text/plain",
		".css":  "text/css",
		".csv":  "text/csv",
		".js":   "text/javascript",
		".md":   "text/markdown",
		".xml":  "text/xml",

		// video
		".avi":  "video/x-msvideo", // avi
		".mp4":  "video/mp4",
		".mpeg": "video/mpeg",
		".ogv":  "video/ogg",
		".ts":   "video/mp2t", // ts
		".webm": "video/webm",
		".3gp":  "video/3gpp",

		// other
		".x-fs-file":   "fs/file",
		".x-fs-folder": "fs/folder",
		".x-fs-link":   "fs/link",
	}
	return tt
}
