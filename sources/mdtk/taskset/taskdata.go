package taskset

import (
	"strings"
	"mdtk/group"
	"mdtk/task"
	"mdtk/code"
	"mdtk/path"
)

type TaskData struct {
	Group group.Group
	Task task.Task
	Code code.Code
	Description []string
	Attributes []string
	ArgsTexts []string
	FilePath path.Path
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
				attrs_end_idx = i
				break
			}
		} else {
			buf = append(buf, v)
		}		
	}

	if attrs_end_idx < 0 {
		return []string{}, str
	}

	return attrs, str
	// return attrs, strings.TrimSpace(string(runes[attrs_end_idx:]))
}

func (td *TaskData) GetAttrsAndSet() {
	attrs, _ := td.getAttributesFromDesc()
	if td.Group.IsPrivate() {
		attrs = append(attrs, "hidden")
	}
	td.Attributes = attrs
}

func (td TaskData) HasAttr(attr string) bool {
	for _, v := range td.Attributes {
		if v == attr { return true }
	}
	return false
}

