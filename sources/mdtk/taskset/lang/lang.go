package lang

import (
	"strings"
	"mdtk/config"
	"mdtk/lib"
	"mdtk/args"
	"mdtk/taskset/grtask"
	"mdtk/taskset/code"
	"time"
	"strconv"
	"path/filepath"
	"os"
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

/*
1.  If get empty slice, do nothing.
2. -c (-Command) option must be at the end of all options. 
3. If -c or -Command exists the end of slice as it is, it will return all the options without the last value. 
4. If -c used such as -cx, only c is removed and return slice.
*/
func removeOpC(strs []string) []string {
	if len(strs) == 0 {
		return []string{}
	}

	back_id := len(strs) - 1
	back := strs[back_id]

	if back == "-c" || back == "-Command" {
		strs = strs[0:back_id]
	} else if lib.Var('c').IsContainedIn([]rune(back)) {
		back := strings.Replace(back, "c", "", 1)
		strs[back_id] = back
	}
	return strs
}

func writeTmpFileAndGetName(code string, ext string) (string, func()) {
	r := int(time.Now().UnixNano())
	fpath := "./mdtk_exec_tmp_" + strconv.Itoa(r) + ext

	abs_fpath, _ := filepath.Abs(fpath)
	abs_fpath = filepath.ToSlash(abs_fpath)

	os.WriteFile(fpath, []byte(code), 0666)
	// In WSL Bash environment on Windows PwSh Terminal, 
	// absolute path of the working directory changes temporarily at runtime, 
	// so relative path must be returned.
	// ex: ( D:\dir\task => /mnt/d/dir/task )
	return fpath, func() { os.Remove(abs_fpath) }
}

// ----------------------------------------------------------------

// The part where the behavior changes with the language is written here.

type LangXInterface interface {
	Iam() int
	GetCmd(string, bool) (string, []string, func())
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


