package sub

import (
	"fmt"
	"mdtk/base"
	"mdtk/parse"
	"mdtk/path"
	"mdtk/config"
)

func ReadConfig(file_flag parse.FlagData) {
	if file_flag.Exist {
		dir := path.Path(file_flag.Value).Dir()
		if err := config.ReadConfig(string(dir)); err != nil {
			fmt.Print(err)
			base.MdtkExit(1)
		}
	} else {
		wd, err := path.GetWorkingDir[string]()
		if err != nil {
			fmt.Print(err)
			base.MdtkExit(1)
		}
		if err := config.ReadConfig(wd); err != nil {
			fmt.Print(err)
			base.MdtkExit(1)
		}
	}
}
