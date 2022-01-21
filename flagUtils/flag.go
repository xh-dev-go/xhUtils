package flagUtils

import (
	_ "embed"
	"flag"
	"github.com/xh-dev-go/xhUtils/flagUtils/flagString"
)
var CommandFlag = flag.CommandLine

func Version(version *string) *flagString.StringParam {
	return flagString.New("version", "application version").SetValue(version)
}
