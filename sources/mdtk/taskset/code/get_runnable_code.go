package code

import (
	"mdtk/args"
	"mdtk/taskset/grtask"
)

type LangInterface interface {
	IsShell() bool
	IsPwSh() bool
}

type CodeWithLang struct {
	code Code
	lang LangInterface
}

func (code Code) WithLang(lang LangInterface) CodeWithLang {
	return CodeWithLang{code: code, lang: lang}
}

func (cl CodeWithLang) GetRunnableCode(tf TaskDataSetInterface, gtname grtask.GroupTask, args args.Args, args_enclose_with_quotes bool, use_new_task_stack bool, nestsize int) (Code, error) {
	switch {
	case cl.lang.IsShell():
		return cl.code.GetRunnableShellCode(tf, gtname, args, args_enclose_with_quotes, use_new_task_stack, nestsize)
	case cl.lang.IsPwSh():
		return cl.code.GetRunnablePwShCode(tf, gtname, args, args_enclose_with_quotes, use_new_task_stack, nestsize)
	default:
		return cl.code.GetRunnableSubCode(tf, gtname, nestsize)
	}
}
