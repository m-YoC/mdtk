package code

import (
	"mdtk/args"
)

func GetArgsConfigShell() ApplyArgsConfig {
	cfg := ApplyArgsConfig{arg_id_first: 1, arg_id_max: 9, escape: `\`}
	cfg.param_alias_arr = []string{"{$}", "<$>"}
	cfg.set_var_func = func(name string, value string) string {
		return name + "=" + value + "; "
	}
	cfg.id_to_param_func = func(id string) string {
		return "$" + id
	}
	return cfg
}

func (code Code) ApplyArgsShell(args args.Args, enclose_with_quotes bool) (Code, error) {
	return code.applyArgsBase(args, enclose_with_quotes, GetArgsConfigShell())
}