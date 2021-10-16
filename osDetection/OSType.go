package osDetection

import "runtime"

type OSType int

const (
	OS_WIN OSType = iota
	OS_LINUX
)

func getOs() OSType {
	if runtime.GOOS == "windows" {
		return OS_WIN
	} else {
		return OS_LINUX
	}
}

var CurOS = getOs()
var CurOSName = runtime.GOOS
