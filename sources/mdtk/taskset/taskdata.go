package taskset

import (
	"strings"
	"mdtk/config"
	"mdtk/taskset/group"
	"mdtk/taskset/task"
	"mdtk/taskset/code"
	"mdtk/taskset/path"
)

const (
	AttrHidden = "hidden"
	AttrTop    = "t"
	AttrBottom = "b"
)

// Inclusive name of shell languanses
const ShellLangs = "SHELL"

type TaskData struct {
	Lang string
	Group group.Group
	Task task.Task
	Code code.Code
	Description []string
	Attributes []string
	ArgsTexts []string
	FilePath path.Path
}


func (td TaskData) LangIsContainedIn(l []string) bool {
	for _, d := range l {
		if d == td.Lang { return true }
	}
	return false
}

func (td *TaskData) SetLang(str string) {
	td.Lang = str
	if td.Lang == "" || td.LangIsContainedIn(config.Config.LangAlias) {
		td.Lang = ShellLangs
	}
}

func (td *TaskData) SetGroup(str string) {
	td.Group.Set(str)
}

func (td *TaskData) SetTask(str string) {
	td.Task.Set(str)
}

func (td *TaskData) SetDescription(str string) {
	td.Description = []string{str}
}

func (td *TaskData) AppendDescription(strs ...string) {
	td.Description = append(td.Description, strs...)
}


// return: (attrs, desc_that_removed_attrs)
func (td TaskData) getAttributesFromDesc() ([]string, string) {
	// Description size is guaranteed to at least 1.
	str := td.Description[0] 
	runes := []rune(str)
	// Must have at least '[', ']' and one letter attribute.
	// And must begin with '['.
	if len(runes) < 3 || runes[0] != '[' {
		return []string{}, str
	}

	end_idx := strings.Index(str, "]")
	if end_idx == -1 {
		return []string{}, str
	}
	attrs := strings.Fields(str[1:end_idx])

	// return attrs, str
	return attrs, strings.TrimSpace(str[end_idx+1:])
}

func (td *TaskData) GetAttrsAndSet() {
	attrs, desc := td.getAttributesFromDesc()
	if td.Group.IsPrivate() || td.Lang != ShellLangs {
		attrs = append(attrs, AttrHidden)
	}

	td.Description[0] = desc
	td.Attributes = attrs
}

func (td TaskData) HasAttr(attr string) bool {
	for _, v := range td.Attributes {
		if v == attr { return true }
	}
	return false
}

