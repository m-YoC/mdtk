package sub

import (
	"mdtk/parse"
)

type OtherFlags struct {
	RunInTaskFileDir bool
}

func FixFlags(flags parse.Flag) (parse.Flag, OtherFlags) {
	var of OtherFlags
	flags = libToFile(flags)
	flags, of.RunInTaskFileDir = f2f(flags)

	return flags, of
} 

func f2f(flags parse.Flag) (parse.Flag, bool) {
	fil := flags.GetIndex("--file")
	fiu := flags.GetIndex("--File")
	if !flags[fil].Exist && flags[fiu].Exist {
		flags[fil].Exist = true
		flags[fil].Value = flags[fiu].Value
	}
	return flags, flags[fiu].Exist
}

func libToFile(flags parse.Flag) parse.Flag {
	fi := flags.GetIndex("--file")
	li := flags.GetIndex("--lib")
	if !flags[fi].Exist && flags[li].Exist {
		flags[fi].Exist = true
		flags[fi].Value = flags[li].Value + ".mdtklib"
	}
	return flags
}


