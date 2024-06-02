package code

import (
	"fmt"
	"strings"
	"strconv"
	"mdtk/args"
)

func escapeQuoteAndEnclose(s string, q string) string {
	if q == "" {
		return s
	}

	return q + strings.Replace(s, q, `\` + q, -1) + q
}

func (code Code) ApplyArgs(args args.Args, enclose_with_quotes bool) (Code, error) {
	argstr := ""
	q := ""
	if enclose_with_quotes {
		q = "'"
	}

	count := 1

	for _, arg := range args {
		name, value, err := arg.GetData()
		if err != nil {
			return "", err
		}
		if enclose_with_quotes && value == "{$}" {
			if count > 9 {
				return "", fmt.Errorf("you set too many special variable {$} ( > 9).")
			}
			argstr += name + "=$" + strconv.Itoa(count) + "; "
			count++
		} else {
			argstr += name + "=" + escapeQuoteAndEnclose(value, q) + "; "
		}
		
	}

	return Code(argstr + "\n" + string(code)), nil
}

func (code Code) RemoveEmbedArgsComment() Code {
	embeds := code.GetEmbedComment("args")

	if len(embeds) == 0 {
		return code
	}

	res := string(code)
	for _, embed := range embeds {
		res = strings.Replace(res, embed[0] + "\n", "", 1)
	}

	return Code(res)
}

