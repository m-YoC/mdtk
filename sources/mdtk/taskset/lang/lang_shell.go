package lang

import (
	"mdtk/config"
	"mdtk/args"
	"mdtk/taskset/grtask"
	"mdtk/taskset/code"
	// "github.com/gookit/color"
)

type LangShell Lang

func (l LangShell) GetCmd(code string) (string, []string) {
	s, ss := splitFirstAndOther(config.Config.Shell)
	execcode := config.Config.ScriptHeadSet + "\n" + code
	return s, append(ss, execcode)
}

func (l LangShell) GetScriptData() (string, string) {
	s, _ := splitFirstAndOther(config.Config.Shell)
	return s, config.Config.ScriptHeadSet
}

func (l LangShell) GetRunnableCode(c code.Code, tf code.TaskDataSetInterface, 
									gtname grtask.GroupTask, 
									args args.Args, 
									args_enclose_with_quotes bool, 
									use_new_task_stack bool, 
									nestsize int) (code.Code, error) {
	return c.GetRunnableShellCode(tf, gtname, args, args_enclose_with_quotes, use_new_task_stack, nestsize)
}


func (l LangShell) GetScriptNameColor() string {
	return ""
}


