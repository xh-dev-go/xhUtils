package flagUtils

import (
	"embed"
	"flag"
	"github.com/xh-dev-go/xhUtils/flagUtils/flagBool"
)
var CommandFlag = flag.CommandLine

func Version(fs embed.FS) *flagBool.BoolParam {
	return flagBool.New("version", "application version")
}
