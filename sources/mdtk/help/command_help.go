package help

import (
	"fmt"
	"mdtk/parse"
)

func ShowCommandHelp(f parse.Flag, descline int) {
	fmt.Println("[mdtk command help]")
	fmt.Println("")
	fmt.Println("mdtk is a markdown task-runner using codeblock.")
	fmt.Println("")
	fmt.Println("  command:  mdtk group-task")
	fmt.Println("            mdtk group-task [--options] -- args...")
	fmt.Println("")
	fmt.Println("  How to write each command")
	fmt.Println("    group-task -> group task, group:task, task")
	fmt.Println("                    Can write without group.")
	fmt.Println("                    In this case, all groups will be searched.")
	fmt.Println("                    Special group-task names are as follows.")
	fmt.Println("                    - '_' group : Searchs only empty-name group.")
	fmt.Println("                    - 'help'    : Show task help.")
	fmt.Println("                    - empty     : If 'default' task is defined, run it.")
	fmt.Println("                                  Otherwise, run 'help'.")
	fmt.Println("    args       -> arg_name=arg_value")
	fmt.Println("                    Write after 2 underbars.")
	fmt.Println("")

	fmt.Println("  options")
	fmt.Printf("%s\n", f.GetHelpStr(4, descline))
}


