package code

import (
	"mdtk/args"
)

func GetArgsConfigPwSh() ApplyArgsConfig {
	cfg := ApplyArgsConfig{arg_id_first: 0, arg_id_max: 99, escape: `'`}
	cfg.set_var_func = func(name string, value string) string {
		return "$" + name + " = " + value + "; "
	}

	cfg.param_alias_arr = []string{"{$}", "<$>"}
	cfg.id_to_param_func = func(id string) string {
		return "$Args[" + id + "]"
	}

	cfg.op_param_alias_arr = []string{"{?}", "<?>"}
	cfg.id_to_op_param_func = func(id string) string {
		return "try{$Args[" + id + "]}catch{$null}"
	}

	return cfg
}

func (code Code) ApplyArgsPwSh(args args.Args, enclose_with_quotes bool) (Code, error) {
	return code.applyArgsBase(args, enclose_with_quotes, GetArgsConfigPwSh())
}
