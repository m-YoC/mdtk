package code

import (
	"fmt"
	"strings"
	"strconv"
	"mdtk/lib"
	"mdtk/args"
)

func escapeQuoteAndEnclose(s string, q string, esc string) string {
	if q == "" {
		return s
	}

	return q + strings.Replace(s, q, esc + q, -1) + q
}

type applyArgsConfig struct {
	count_first int
	count_max int
	escape string
	param_alias_arr []string
	set_var_func func(string, string) string
	to_param_func func(string) string
}

func (code Code) applyArgsBase(args args.Args, quotes bool, cfg applyArgsConfig) (Code, error) {
	argstr := ""
	q := ""
	if quotes {
		q = "'"
	}

	count := cfg.count_first

	for _, arg := range args {
		name, value, err := arg.GetData()
		if err != nil {
			return "", err
		}
		if lib.Var(value).IsContainedIn(cfg.param_alias_arr) {
			if count > cfg.count_max {
				return "", fmt.Errorf("you set too many special variable ( > %d).", cfg.count_max)
			}
			argstr += cfg.set_var_func(name, cfg.to_param_func(strconv.Itoa(count)))
			count++
		} else {
			argstr += cfg.set_var_func(name, escapeQuoteAndEnclose(value, q, cfg.escape))
		}
		
	}

	return Code(argstr + "\n" + string(code)), nil
}

func (code Code) ApplyArgsShell(args args.Args, enclose_with_quotes bool) (Code, error) {
	cfg := applyArgsConfig{count_first: 1, count_max: 9, escape: `\`}
	cfg.param_alias_arr = []string{"{$}", "<$>"}
	cfg.set_var_func = func(name string, value string) string {
		return name + "=" + value + "; "
	}
	cfg.to_param_func = func(id string) string {
		return "$" + id
	}

	return code.applyArgsBase(args, enclose_with_quotes, cfg)
}

func (code Code) ApplyArgsPwSh(args args.Args, enclose_with_quotes bool) (Code, error) {
	cfg := applyArgsConfig{count_first: 0, count_max: 99, escape: `'`}
	cfg.param_alias_arr = []string{"<$>"}
	cfg.set_var_func = func(name string, value string) string {
		return "$" + name + " = " + value + "; "
	}
	cfg.to_param_func = func(id string) string {
		return "$Args[" + id + "]"
	}

	return code.applyArgsBase(args, enclose_with_quotes, cfg)
}

func (code Code) RemoveEmbedArgsComment() Code {
	return code.RemoveEmbedComment("args")
}

func (code Code) GetEmbedArgsText() []string {
	return code.GetEmbedCommentText("args")
}

