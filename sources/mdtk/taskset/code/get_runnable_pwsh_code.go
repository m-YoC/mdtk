package code

import (
	"mdtk/base"
	"mdtk/args"
	"mdtk/taskset/grtask"
)

func (c Code) GetRunnablePwShCode(tf TaskDataSetInterface, gtname grtask.GroupTask, args args.Args, args_enclose_with_quotes bool, use_new_task_stack bool, nestsize int) (Code, error) {
	if err := args.Validate(); err != nil {
		return "", err
	}
	
	if f:= false; !use_new_task_stack {
		if c, f = c.ApplyConfigOnce(gtname); f {
			base.DebugLogGray(nestsize, "[config once] Passed because already loaded.\n")
			return c, nil
		}
	} else {
		c, f = c.CheckAndRemoveConfigOnce()
	}

	var err error
	if len(args) != 0 {
		c, err = c.ApplyArgsPwSh(args, args_enclose_with_quotes)
		if err != nil {
			return "", err
		}
	}

	impl_func := func() (Code, error) {
		CurrentTaskStackData.Set(gtname)

		res, err := c.ApplyEmbedCodes(tf, nestsize)
		if err != nil {
			return "", err
		}

		//*
		res, err = res.ApplySubTasksPwSh(tf, nestsize)
		if err != nil {
			return "", err
		}//*/
		
		// ? is alias of Where-Object
		return res.ApplyFuncsPwSh(tf, nestsize)
	}

	if use_new_task_stack {
		c, err = WithNewTaskStackData(c, impl_func)
	} else {
		c, err = impl_func()
	}

	if err != nil {
		return "", err
	}
	return c, nil
}

