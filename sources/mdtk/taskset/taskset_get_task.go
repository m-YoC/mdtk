package taskset

import (
	"fmt"
	"mdtk/grtask"
	"mdtk/code"
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

	if l == ShellLangs {
		return tds.GetTaskShell(c, gtname, args, args_enclose_with_quotes, use_new_task_stack, nestsize)
	} else {
		return tds.GetTaskSubLangs(c, gtname, nestsize)
	}
}

func (tds TaskDataSet) GetTaskShell(c code.Code, gtname grtask.GroupTask, args args.Args, args_enclose_with_quotes bool, use_new_task_stack bool, nestsize int) (code.Code, error) {
	if err := args.Validate(); err != nil {
		return "", err
	}

	if f:= false; !use_new_task_stack {
		if c, f = c.ApplyConfigOnce(gtname); f {
			return c, nil
		}
	} else {
		c, f = c.CheckAndRemoveConfigOnce()
	}
	
	var err error
	if len(args) != 0 {
		c, err = c.ApplyArgs(args, args_enclose_with_quotes)
		if err != nil {
			return "", err
		}
	}

	impl_func := func() (code.Code, error) {
		code.CurrentTaskStackData.Set(gtname)

		res, err := c.ApplyEmbedCodes(tds, nestsize)
		if err != nil {
			return "", err
		}

		res, err = res.ApplySubTasks(tds, nestsize)
		if err != nil {
			return "", err
		}
		
		return res.ApplyFuncs(tds, nestsize)
	}

	if use_new_task_stack {
		c, err = code.WithNewTaskStackData(c, impl_func)
	} else {
		c, err = impl_func()
	}

	if err != nil {
		return "", err
	}

	return c, nil
}

func (tds TaskDataSet) GetTaskSubLangs(c code.Code, gtname grtask.GroupTask, nestsize int) (code.Code, error) {
	var f bool
	if c, f = c.ApplyConfigOnce(gtname); f {
		return c, nil
	}
	var err error
	c, err = c.ApplyEmbedCodes(tds, nestsize)
	if err != nil {
		return "", err
	}
	return c, nil
}

func (tds TaskDataSet) GetTaskStart(gtname grtask.GroupTask, args args.Args, nestsize int) (code.Code, error) {
	s, err := tds.GetTask(gtname, args, true, false, nestsize)
	if err != nil {
		return "", err
	}
	return s, nil
}

