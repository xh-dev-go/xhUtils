package flagUtils

import (
	_ "embed"
	"flag"
	"github.com/xh-dev-go/xhUtils/flagUtils/flagBool"
)
var CommandFlag = flag.CommandLine

func Version() *flagBool.BoolParam {
	return flagBool.New("version", "show application version")
}
