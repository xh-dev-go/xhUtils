package flagUtils

import (
	_ "embed"
	"flag"
	"github.com/xh-dev-go/xhUtils/flagUtils/flagBool"
	"os"
)
var CommandFlag = flag.CommandLine

func Version() *flagBool.BoolParam {
	return flagBool.New("version", "show application version")
}

func TestShowVersion(param flagBool.BoolParam, version string){
	if param.Value() {
		println(version)
		os.Exit(0)
	}
}

