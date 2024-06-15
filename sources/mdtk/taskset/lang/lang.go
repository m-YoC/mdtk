package lang

import (
	"mdtk/config"
	"mdtk/lib"
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

func (l Lang) IsShell() bool {
	return l == ShellLangs
}

func (l Lang) IsPwSh() bool {
	return l == PwShLangs
}

func (l Lang) IsSub() bool {
	return !(l.IsShell() || l.IsPwSh())
}

func (l Lang) GetId() int {
	switch {
	case l.IsShell():
		return LANG_SHELL
	case l.IsPwSh():
		return LANG_PWSH
	default:
		return LANG_SUB
	}
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

func (l Lang) GetCmd(code string) (string, []string) {
	switch {
	case l.IsShell():
		s, ss := splitFirstAndOther(config.Config.Shell)
		execcode := config.Config.ScriptHeadSet + "\n" + code
		return s, append(ss, execcode)
	case l.IsPwSh():
		s, ss := splitFirstAndOther(config.Config.PowerShell)
		execcode := config.Config.PwShHeadSet + "\n" + code
		return s, append(ss, execcode)
	default:
		return "echo", []string{"Bad Exec Command"}
	}
}

func (l Lang) GetScriptData() (string, string) {
	switch {
	case l.IsShell():
		s, _ := splitFirstAndOther(config.Config.Shell)
		return s, config.Config.ScriptHeadSet
	case l.IsPwSh():
		s, _ := splitFirstAndOther(config.Config.PowerShell)
		return s, config.Config.PwShHeadSet
	default:
		return "nothing", "Bad Script Head"
	}
}


