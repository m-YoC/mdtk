package lang

import (
	// "mdtk/config"
	"mdtk/args"
	"mdtk/taskset/grtask"
	"mdtk/taskset/code"
	"github.com/gookit/color"
)

type LangSub Lang

func (l LangSub) GetCmd(code string) (string, []string) {
	return "echo", []string{"Bad Exec Command"}
}

func (l LangSub) GetScriptData() (string, string) {
	return "nothing", "Bad Script Head"
}

func (l LangSub) GetRunnableCode(c code.Code, tf code.TaskDataSetInterface, 
									gtname grtask.GroupTask, 
									args args.Args, 
									args_enclose_with_quotes bool, 
									use_new_task_stack bool, 
									nestsize int) (code.Code, error) {
	return c.GetRunnableSubCode(tf, gtname, nestsize)
}

func (l LangSub) GetScriptNameColor() string {
	return color.Blue.Sprint( "<" + Lang(l).String() + "> ")
}


