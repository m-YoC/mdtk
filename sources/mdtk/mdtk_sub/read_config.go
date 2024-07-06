package sub

import (
	"fmt"
	"os"
	"mdtk/base"
	"mdtk/parse"
	"path/filepath"
	"mdtk/config"
)

func ReadConfig(file_flag parse.FlagData) {
	if file_flag.Exist {
		dir := filepath.Dir(file_flag.Value)
		if err := config.ReadConfig(dir); err != nil {
			fmt.Fprint(os.Stderr, err)
			base.MdtkExit(1)
		}
	} else {
		wd, err := base.GetWorkingDir()
		if err != nil {
			fmt.Fprint(os.Stderr, err)
			base.MdtkExit(1)
		}
		if err := config.ReadConfig(wd); err != nil {
			fmt.Fprint(os.Stderr, err)
			base.MdtkExit(1)
		}
	}
}
