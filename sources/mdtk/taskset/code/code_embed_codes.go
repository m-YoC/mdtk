package code

import (
	"fmt"
	"strings"
	"mdtk/taskset/grtask"
	"mdtk/args"
)

func embedCmdsConstraint(cmds []string, args args.Args, errstr string, err error) (grtask.GroupTask, error) {
	if err != nil {
		return  grtask.GroupTask(""), err
	}
	if len(cmds) != 1 || len(args) != 0 {
		return grtask.GroupTask(""), fmt.Errorf("%s", errstr)
	}
	return grtask.GroupTask(cmds[0]), nil
}


func (code Code) ApplyEmbedCodes(tf TaskDataSetInterface, nestsize int) (Code, error) {
	embeds := code.GetEmbedComment("embed")

	if len(embeds) == 0 {
		return code, nil
	}

	res := string(code)
	for _, embed := range embeds {
		gtname, err := embedCmdsConstraint(extractSubCmds(embed[1]))
		if err != nil {
			return "", fmt.Errorf("%s-> %s\n", err, embed[0])
		}

		subcode, err := tf.GetTask(grtask.GroupTask(gtname), args.Args{}, false, false, nestsize-1)
		if err != nil {
			return "", err
		}
		subcode = subcode.RemoveEmbedDescComment().RemoveEmbedArgsComment()
		res = strings.Replace(res, embed[0], string(subcode), 1)
	}

	return Code(res), nil
}
