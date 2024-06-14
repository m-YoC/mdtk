package sub

import (
	"fmt"
	"mdtk/taskset/grtask"
	"mdtk/args"
	"mdtk/taskset"
)

func Validate(gtname grtask.GroupTask, args args.Args, tds taskset.TaskDataSet, has_all_task_flag bool) error {
	if err := gtname.Validate(); err != nil {
		return err
	}
	if err := args.Validate(); err != nil {
		return err
	}

	if td, err := tds.GetTaskData(gtname.Split()); err != nil {
		return err
	} else if td.HasAttr(taskset.AttrHidden) && !has_all_task_flag {
		s := fmt.Sprintln("Private/Hidden group cannot be executed directly.")
		s += fmt.Sprintf("[group: %s | task: %s | path: %s]\n", td.Group, td.Task, td.FilePath)
		return fmt.Errorf("%s", s)
	}

	return nil
}
