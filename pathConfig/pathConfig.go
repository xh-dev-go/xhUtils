package pathConfig

import (
	"log"
	"os"
	"github.com/xh-dev-go/xhUtils/osDetection"
)

func Info(){
	println(ConfigDir)
	println(HomeDir)
	println(CacheDir)
	println(string(SysSeparator))
}

func GetConfig() string {
	dir, err := os.UserConfigDir()
	if err != nil {
		log.Fatal("Fail to retrieve config directory", err)
	}
	return dir
}

func GetHome() string {
	dir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("Fail to retrieve config directory", err)
	}
	return dir
}

func GetCache() string {
	dir, err := os.UserCacheDir()
	if err != nil {
		log.Fatal("Fail to retrieve config directory", err)
	}
	return dir
}

var ConfigDir = GetConfig()
var CacheDir = GetCache()
var HomeDir = GetHome()


func getSysSeparator() rune {
	if osDetection.CurOS == osDetection.OS_WIN{
		return '\\'
	} else {
		return '/'
	}
}
var SysSeparator = getSysSeparator()
