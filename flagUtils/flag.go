package flagUtils

import (
	_ "embed"
	"flag"
	"github.com/xh-dev-go/xhUtils/flagUtils/flagString"
)
var CommandFlag = flag.CommandLine

//go:embed version
var version string
func Version() *flagString.StringParam {
	return flagString.New("version", "application version").SetValue(&version)
}
