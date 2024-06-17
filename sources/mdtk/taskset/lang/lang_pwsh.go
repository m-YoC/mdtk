package lang

import (
	"mdtk/config"
	"mdtk/args"
	"mdtk/taskset/grtask"
	"mdtk/taskset/code"
	"github.com/gookit/color"
)

type LangPwSh Lang

func (l LangPwSh) Iam() int {
	return LANG_PWSH
}

func (l LangPwSh) GetCmd(basecode string, use_tmpfile bool) (string, []string, func()) {
	if use_tmpfile {
		return l.GetCmdUsingTmp(basecode)
	} else {
		return l.GetCmdDirect(basecode)
	}
}

func (l LangPwSh) GetCmdDirect(basecode string) (string, []string, func()) {
	s, ss := splitFirstAndOther(config.Config.PowerShell)
	execcode := config.Config.PwShHeadSet + "\n" + basecode
	return s, append(ss, execcode), func(){}
}

func (l LangPwSh) GetCmdUsingTmp(basecode string) (string, []string, func()) {
	s, ss := splitFirstAndOther(config.Config.PowerShell)
	execcode := config.Config.PwShHeadSet + "\n" + basecode
	fname, rmf := writeTmpFileAndGetName(execcode, ".ps1")
	return s, append(removeOpC(ss), fname), rmf
}

func (l LangPwSh) GetScriptData() (string, string) {
	s, _ := splitFirstAndOther(config.Config.PowerShell)
	return s, config.Config.PwShHeadSet
}

func (l LangPwSh) GetRunnableCode(c code.Code, tf code.TaskDataSetInterface, 
									gtname grtask.GroupTask, 
									args args.Args, 
									args_enclose_with_quotes bool, 
									use_new_task_stack bool, 
									nestsize int) (code.Code, error) {
	return c.GetRunnablePwShCode(tf, gtname, args, args_enclose_with_quotes, use_new_task_stack, nestsize)
}

func (l LangPwSh) GetScriptNameColor() string {
	return color.Green.Sprint( "<" + Lang(l).String() + "> ")
}



