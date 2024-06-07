package code

import (
	"strings"
	"mdtk/grtask"
	"mdtk/args"
)



func (code Code) ApplyEmbedCodes(tf TaskDataSetInterface, nestsize int) (Code, error) {
	embeds := code.GetEmbedComment("embed")

	if len(embeds) == 0 {
		return code, nil
	}

	res := string(code)
	for _, embed := range embeds {
		subcode, err := tf.GetTask(grtask.GroupTask(embed[1]), args.Args{}, false, false, nestsize-1)
		if err != nil {
			return "", err
		}
		subcode = subcode.RemoveEmbedDescComment().RemoveEmbedArgsComment()
		res = strings.Replace(res, embed[0], string(subcode), 1)
	}

	return Code(res), nil
}
