package help

import (
	_ "embed"
)

//go:embed mdtk_manual.txt
var mdtkman string

func ShowManual(pager_min_row uint) {
	// fmt.Println(mdhelp)
	PagerOutput(mdtkman, pager_min_row)
}
