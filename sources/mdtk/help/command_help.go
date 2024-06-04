package help

import (
	"fmt"
	"mdtk/parse"
	"mdtk/config"
	_ "embed"
)

//go:embed cmd_help.txt
var cmdhelp string

func ShowCommandHelp(f parse.Flag, descline uint) {
	s := fmt.Sprintln(cmdhelp)

	s += fmt.Sprintln("  options")
	s += fmt.Sprintf("%s\n", f.GetHelpStr(4, descline))

	PagerOutput(s, config.Config.PagerMinLimit)
}


