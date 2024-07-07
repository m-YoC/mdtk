package code

import (
	"fmt"
	"strings"
	"mdtk/base"
	"mdtk/taskset/grtask"
	"mdtk/args"
)

func taskCmdsConstraint(cmds []string, args args.Args, errstr string, err error) (grtask.GroupTask, args.Args, error) {
	if err != nil {
		return grtask.GroupTask(""), args, err
	}
	if len(cmds) != 1 {
		return grtask.GroupTask(""), args, fmt.Errorf("%s", errstr)
	}
	if args.HasValue("{$}", "<$>") {
		return grtask.GroupTask(""), args, fmt.Errorf("%s\n", "Args of '#task>' contains a value that cannot be used ({$}, <$>).")
	}

	return grtask.GroupTask(cmds[0]), args, nil
}

func (code Code) ApplySubTasks(tf TaskDataSetInterface, nestsize int) (Code, error) {
	tasks := code.GetEmbedComment("task")

	if len(tasks) == 0 {
		return code, nil
	}

	indent := strings.Repeat(" ", 2)

	res := string(code)
	for i, task := range tasks {
		gtname, args, err := taskCmdsConstraint(extractSubCmds(task[1]))
		if err != nil {
			return "", fmt.Errorf("%s-> %s\n", err, task[0])
		}
		use_new_task_stack := true
		head := ""

		base.DebugLogGreen(nestsize-1, fmt.Sprintf("#task(%d)>\n", i+1))
		subcode, err := tf.GetTask(gtname, args, false, use_new_task_stack, nestsize-1)
		if err != nil {
			return "", err
		}
		subcode = subcode.RemoveEmbedDescComment().RemoveEmbedArgsComment()
		rsubcode := indent + strings.Replace(string(subcode), "\n", "\n" + indent, -1)
		execsubcode := head + "(\n"// + string(gtname) + "\n"
		execsubcode += indent + ":" + " '" + string(gtname) + "'\n"
		execsubcode += rsubcode + "\n)"
		res = strings.Replace(res, task[0], execsubcode, 1)
	}

	return Code(res), nil
}
