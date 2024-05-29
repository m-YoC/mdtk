package help

import (
	"fmt"
	"mdtk/parse"
	_ "embed"
)

//go:embed cmd_help.txt
var cmdhelp string

func ShowCommandHelp(f parse.Flag, descline int) {
	s := fmt.Sprintln(cmdhelp)

	s += fmt.Sprintln("  options")
	s += fmt.Sprintf("%s\n", f.GetHelpStr(4, descline))

	PagerOutput(s, 40)
}


