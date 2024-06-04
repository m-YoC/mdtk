package grtask

import (
	"fmt"
	"regexp"
	"mdtk/base"
	"mdtk/group"
	"mdtk/task"
)

const group_and_task_name_reg = "(?:(?P<group>" + base.NameReg + "):)?(?P<task>" + base.NameReg + ")"
var group_and_task_name_rex = regexp.MustCompile("^" + group_and_task_name_reg + "$")

type GroupTask string

func (gt GroupTask) Validate() error {
	if !group_and_task_name_rex.MatchString(string(gt)) {
		return fmt.Errorf("Validation error: group/task name. => '%s'\n", gt)
	}
	return nil
}

func (gt GroupTask) ValidatePublic() error {
	// empty group is safe
	if g, _, err := gt.Split(); err != nil {
		return err
	} else {
		if err = g.ValidatePublicEmptyIsSafe(); err != nil {
			return err
		}
	}
	return nil
}

func (gt GroupTask) Split() (group.Group, task.Task, error) {
	if err := gt.Validate(); err != nil {
		return "", "", err
	}

	res := group_and_task_name_rex.FindStringSubmatch(string(gt))
	g := res[group_and_task_name_rex.SubexpIndex("group")]
	t := res[group_and_task_name_rex.SubexpIndex("task")]

	return group.Group(g), task.Task(t), nil
}

func Create(g group.Group, t task.Task) GroupTask {
	gt := GroupTask(t)
	if g != "" {
		gt = GroupTask(string(g) + ":" + string(t))
	}
	return gt
}

