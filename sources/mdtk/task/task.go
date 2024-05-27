package task

import (
	"fmt"
	"regexp"
	"mdtk/base"
)

var validate_taskname_rex = regexp.MustCompile("^" + base.NameReg + "$")

type Task string

func (t Task) Validate() error {
	if !validate_taskname_rex.MatchString(string(t)) {
		return fmt.Errorf("Validation error: task name. => %s\n", t)
	}
	return nil
}

func (t Task) Match(task Task) bool {
	return task == t
}
