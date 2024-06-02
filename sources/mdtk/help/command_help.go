package help

import (
	"fmt"
	"mdtk/parse"
	_ "embed"
)

//go:embed cmd_help.txt
var cmdhelp string

func ShowCommandHelp(f parse.Flag, descline uint, pager_min_row uint) {
	s := fmt.Sprintln(cmdhelp)

	s += fmt.Sprintln("  options")
	s += fmt.Sprintf("%s\n", f.GetHelpStr(4, descline))

	PagerOutput(s, pager_min_row)
}


