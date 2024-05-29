package help

import (
	"fmt"
	_ "embed"
)

//go:embed md_help.txt
var mdhelp string

func ShowMarkdownHelp() {
	fmt.Println(mdhelp)
	return
}
