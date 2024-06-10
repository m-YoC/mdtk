package sub

import (
	"fmt"
	"mdtk/parse"
	"mdtk/path"
	"mdtk/config"
)

func ReadConfig(flags parse.Flag) {
	if fd := flags.GetData("--file"); fd.Exist {
		if err := config.ReadConfig(string(path.Path(fd.Value).Dir())); err != nil {
			fmt.Print(err)
			MdtkExit(1)
		}
	} else {
		if err := config.ReadConfig(""); err != nil {
			fmt.Print(err)
			MdtkExit(1)
		}
	}
}
