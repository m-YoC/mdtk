package lang

import (
	"fmt"
	"strings"
	"mdtk/base"
	"mdtk/config"
	"mdtk/args"
	"mdtk/taskset/grtask"
	"mdtk/taskset/code"
	"os/exec"
	// "github.com/gookit/color"
)

type LangShell Lang

func (l LangShell) Iam() int {
	return LANG_SHELL
}

func (l LangShell) GetCmd(basecode string, use_tmpfile bool) (string, []string, func()) {
	if use_tmpfile {
		return l.GetCmdUsingTmp(basecode)
	} else if CanRunSh() {
		return l.GetCmdDirect(basecode)
	} else {
		// In powershell environment, the following replaces are required.
		basecode = strings.Replace(basecode, "$", "\\$", -1)
		return l.GetCmdDirect(basecode)
	}
}

func (l LangShell) GetCmdDirect(basecode string) (string, []string, func()) {
	s, ss := splitFirstAndOther(config.Config.Shell)
	execcode := config.Config.ScriptHeadSet + "\n" + basecode
	return s, append(ss, execcode), func(){}
}

func (l LangShell) GetCmdUsingTmp(basecode string) (string, []string, func()) {
	s, ss := splitFirstAndOther(config.Config.Shell)
	execcode := config.Config.ScriptHeadSet + "\n" + basecode
	fname, rmf := writeTmpFileAndGetName(execcode, ".sh")
	return s, append(removeOpC(ss), fname), rmf
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

// ------------------------------------------------------------------------------

func CanRunSh() bool {
	s, _ := splitFirstAndOther(config.Config.Shell)
	res, err := exec.Command(s, "-c", "x=ok; echo $x").Output()
	if err != nil {
		fmt.Println("Could not run test shell script.")
		base.MdtkExit(1)
	}
	if string(res) == "ok\n" {
		return true
	}
	return false
}



