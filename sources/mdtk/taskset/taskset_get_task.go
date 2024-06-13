package taskset

import (
	"fmt"
	"mdtk/taskset/grtask"
	"mdtk/taskset/code"
	"mdtk/args"
)

func (tds TaskDataSet) GetTask(gtname grtask.GroupTask, args args.Args, args_enclose_with_quotes bool, use_new_task_stack bool, nestsize int) (code.Code, error) {
	if nestsize <= 0 {
		return "", fmt.Errorf("Nest of embed/task comments is too deep.\n")
	}

	c, l, err := tds.GetCode(gtname.Split())
	if err != nil {
		return "", err
	}

	switch l {
	case ShellLangs:
		return c.GetRunnableShellCode(tds, gtname, args, args_enclose_with_quotes, use_new_task_stack, nestsize)
	default:
		return c.GetRunnableSubCode(tds, gtname, nestsize)
	}
}


func (tds TaskDataSet) GetTaskStart(gtname grtask.GroupTask, args args.Args, nestsize int) (code.Code, error) {
	s, err := tds.GetTask(gtname, args, true, false, nestsize)
	if err != nil {
		return "", err
	}
	return s, nil
}

