package taskset

import (
	"strings"
	"mdtk/group"
	"mdtk/task"
	"mdtk/code"
	"mdtk/path"
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

// return: (attrs, desc_that_removed_attrs)
func (td TaskData) getAttributesFromDesc() ([]string, string) {
	// Description size is guaranteed to at least 1.
	str := td.Description[0] 
	runes := []rune(str)
	// Must have at least '[', ']' and one letter attribute.
	if len(runes) < 3 {
		return []string{}, str
	}

	// Must begin with '['.
	if runes[0] != '[' {
		return []string{}, str
	}

	attrs := []string{}
	buf := []rune{}
	attrs_end_idx := -1

	for i, v := range runes[1:] {
		// If ']' is found, attrs exists. 
		if v == ']' || v == ' ' {
			if len(buf) != 0 {
				attrs = append(attrs, strings.ToLower(string(buf)))
				buf = []rune{}
			}
			if v == ']' {
				attrs_end_idx = i + 1
				break
			}
		} else {
			buf = append(buf, v)
		}		
	}

	if attrs_end_idx < 0 {
		return []string{}, str
	}

	// return attrs, str
	return attrs, strings.TrimSpace(string(runes[attrs_end_idx+1:]))
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

