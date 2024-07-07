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

func (code Code) ApplySubTasks(tf TaskDataSetInterface, args_enclose_with_quotes bool, head string, brackets []string, nothing_cmd string, nestsize int) (Code, error) {
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

		base.DebugLogGreen(nestsize-1, fmt.Sprintf("#task(%d)>\n", i+1))
		subcode, err := tf.GetTask(gtname, args, args_enclose_with_quotes, use_new_task_stack, nestsize-1)
		if err != nil {
			return "", err
		}
		subcode = subcode.RemoveEmbedDescComment().RemoveEmbedArgsComment()
		rsubcode := indent + strings.Replace(string(subcode), "\n", "\n" + indent, -1)
		execsubcode := head + brackets[0] + "\n" // + "  # " + string(gtname) + "\n"
		execsubcode += indent + nothing_cmd + " '" + string(gtname) + "'\n"
		execsubcode += rsubcode + "\n" + brackets[1]
		res = strings.Replace(res, task[0], execsubcode, 1)
	}

	return Code(res), nil
}

func (code Code) ApplySubTasksShell(tf TaskDataSetInterface, nestsize int) (Code, error) {
	return code.ApplySubTasks(tf, false, "", ParenTheses, ":", nestsize)
}

func (code Code) ApplySubTasksPwSh(tf TaskDataSetInterface, nestsize int) (Code, error) {
	return code.ApplySubTasks(tf, true, "& ", CurlyBrackets, "? .", nestsize)
}
