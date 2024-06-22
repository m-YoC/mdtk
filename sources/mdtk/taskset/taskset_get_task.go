package taskset

import (
	"fmt"
	"mdtk/base"
	"mdtk/taskset/grtask"
	"mdtk/taskset/code"
	"mdtk/args"
)

func (tds TaskDataSet) GetTask(gtname grtask.GroupTask, args args.Args, args_enclose_with_quotes bool, use_new_task_stack bool, nestsize int) (code.Code, error) {
	if nestsize <= 0 {
		return "", fmt.Errorf("Nest of embed/task comments is too deep.\n")
	}

	base.DebugLog(nestsize, fmt.Sprintf("%s / args: %v\n", gtname, args))

	td, err := tds.GetTaskData(gtname.Split())
	if err != nil {
		return "", err
	}

	base.DebugLogGray(nestsize, fmt.Sprintf("lang: %s / path: %s\n",  td.Lang, td.FilePath))

	return td.Lang.LangX().GetRunnableCode(td.Code, tds, gtname, args, args_enclose_with_quotes, use_new_task_stack, nestsize)
}


func (tds TaskDataSet) GetTaskStart(gtname grtask.GroupTask, args args.Args, nestsize int) (code.Code, error) {
	s, err := tds.GetTask(gtname, args, true, false, nestsize)
	if err != nil {
		return "", err
	}
	return s, nil
}

