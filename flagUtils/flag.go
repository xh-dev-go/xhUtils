package flagUtils

import (
	_ "embed"
	"github.com/xh-dev-go/xhUtils/flagUtils/flagBool"
	"os"
)

func Version() *flagBool.BoolParam {
	return flagBool.New("version", "show application version")
}

func TestShowVersion(param flagBool.BoolParam, version string){
	if param.Value() {
		println(version)
		os.Exit(0)
	}
}

