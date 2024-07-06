package code

import (
	"fmt"
	"strings"
	"strconv"
	"mdtk/lib"
	"mdtk/args"
)

func EscapeQuoteAndEnclose(s string, q string, esc string) string {
	if q == "" {
		return s
	}

	return q + strings.Replace(s, q, esc + q, -1) + q
}

type ApplyArgsConfig struct {
	arg_id_first int
	arg_id_max int
	escape string
	set_var_func func(string, string) string

	param_alias_arr []string
	id_to_param_func func(string) string

	op_param_alias_arr []string
	id_to_op_param_func func(string) string
}

func (code Code) applyArgsBase(args args.Args, quotes bool, cfg ApplyArgsConfig) (Code, error) {
	argstr := ""
	q := ""
	if quotes {
		q = "'"
	}

	count := cfg.arg_id_first

	for _, arg := range args {
		name, value, err := arg.GetData()
		if err != nil {
			return "", err
		}
		if lib.Var(value).IsContainedIn(cfg.param_alias_arr) {
			if count > cfg.arg_id_max {
				return "", fmt.Errorf("you set too many special variable ( > %d).", cfg.arg_id_max)
			}
			argstr += cfg.set_var_func(name, cfg.id_to_param_func(strconv.Itoa(count)))
			count++
		} else if lib.Var(value).IsContainedIn(cfg.op_param_alias_arr) { 
			if count > cfg.arg_id_max {
				return "", fmt.Errorf("you set too many special variable ( > %d).", cfg.arg_id_max)
			}
			argstr += cfg.set_var_func(name, cfg.id_to_op_param_func(strconv.Itoa(count)))
			count++
		} else {
			argstr += cfg.set_var_func(name, EscapeQuoteAndEnclose(value, q, cfg.escape))
		}
		
	}

	return Code(argstr + "\n" + string(code)), nil
}


func (code Code) RemoveEmbedArgsComment() Code {
	return code.RemoveEmbedComment("args")
}

func (code Code) GetEmbedArgsText() []string {
	return code.GetEmbedCommentText("args")
}

