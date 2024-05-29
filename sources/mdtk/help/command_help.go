package help

import (
	"fmt"
	"mdtk/parse"
	_ "embed"
)

//go:embed cmd_help.txt
var cmdhelp string

func ShowCommandHelp(f parse.Flag, descline int) {
	fmt.Println(cmdhelp)

	fmt.Println("  options")
	fmt.Printf("%s\n", f.GetHelpStr(4, descline))
}


