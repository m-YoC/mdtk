package help

import (
	_ "embed"
)

//go:embed md_help.txt
var mdhelp string

func ShowMarkdownHelp() {
	// fmt.Println(mdhelp)
	PagerOutput(mdhelp)
	return
}
