package lang

import (
	"mdtk/config"
	"mdtk/lib"
	"mdtk/args"
	"mdtk/taskset/grtask"
	"mdtk/taskset/code"
)

// Inclusive name of shell languanses
const ShellLangs = "SHELL"
const PwShLangs = "pwsh"

const (
	LANG_SHELL = iota
	LANG_PWSH
	LANG_SUB
	LANG_SIZE
)


type Lang string

func (l *Lang) Set(str string) {
	switch {
	case str == "":
		(*l) = ShellLangs
	case lib.Var(str).IsContainedIn(config.Config.LangAlias):
		(*l) = ShellLangs
	case lib.Var(str).IsContainedIn(config.Config.LangAliasPwSh):
		(*l) = PwShLangs
	default:
		(*l) = Lang(str)
	}
}

func (l Lang) String() string {
	return string(l)
}

func splitFirstAndOther(strs []string) (string, []string) {
	switch len(strs) {
	case 0:
		return "echo", []string{"Bad Exec Command"}
	case 1:
		return strs[0], []string{}
	default:
		return strs[0], strs[1:]
	}
}

// ----------------------------------------------------------------

// The part where the behavior changes with the language is written here.

type LangXInterface interface {
	GetCmd(string) (string, []string)
	GetScriptData() (string, string)

	GetRunnableCode(code.Code, code.TaskDataSetInterface, grtask.GroupTask, args.Args, bool, bool, int) (code.Code, error)
	
	GetScriptNameColor() string
}

func (l Lang) IsShell() bool {
	return l == ShellLangs
}

func (l Lang) IsPwSh() bool {
	return l == PwShLangs
}

func (l Lang) IsSub() bool {
	return !(l.IsShell() || l.IsPwSh())
}

func (l Lang) LangX() LangXInterface {
	switch {
	case l.IsShell():
		return LangShell(l)
	case l.IsPwSh():
		return LangPwSh(l)
	default:
		return LangSub(l)
	}
}


