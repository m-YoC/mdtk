package help

import (
	"mdtk/config"
	_ "embed"
)

//go:embed mdtk_manual.txt
var mdtkman string

func ShowManual() {
	// fmt.Println(mdhelp)
	PagerOutput(mdtkman, config.Config.PagerMinLimit)
}
