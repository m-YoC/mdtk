package sub

import (
	"mdtk/parse"
	"mdtk/config"
)

func GetNestSize(nest_flag parse.FlagData) uint {
	nestsize := config.Config.NestMaxDepth
	if nest_flag.Exist {
		nestsize = uint(nest_flag.ValueUint())
	}

	return nestsize
}
