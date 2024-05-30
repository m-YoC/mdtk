package help

import (
	_ "embed"
)

//go:embed mdtk_manual.txt
var mdtkman string

func ShowManual() {
	// fmt.Println(mdhelp)
	PagerOutput(mdtkman, 40)
}
