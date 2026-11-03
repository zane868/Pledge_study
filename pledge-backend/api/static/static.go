package static

import (
	"embed"
	"path"
	"runtime"
)

func GetCurrentAbPathByCaller() string {
	var abPath string
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		abPath = path.Dir(filename)
	}
	return abPath
}

//go:embed img
var Files embed.FS
