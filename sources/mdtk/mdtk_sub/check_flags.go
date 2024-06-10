package sub

import (
	"mdtk/parse"
)

func LibToFile(flags parse.Flag) parse.Flag {
	fi := flags.GetIndex("--file")
	li := flags.GetIndex("--lib")
	if !flags[fi].Exist && flags[li].Exist {
		flags[fi].Exist = true
		flags[fi].Value = flags[li].Value + ".mdtklib"
	}
	return flags
}
