package flagUtils

import (
	"flag"
)

type Command interface {
	Help() *bool
}

const StrEmpty string = ""

const CMD_HELP_SHORT, CMD_HELP_ABBR = "h", "help"
func SetHelp(command Command, flagset *flag.FlagSet){
	flagset.BoolVar(command.Help(), CMD_HELP_SHORT, false, "print help")
	flagset.BoolVar(command.Help(), CMD_HELP_ABBR, false, "print help")
}

var CommandFlag = flag.CommandLine
