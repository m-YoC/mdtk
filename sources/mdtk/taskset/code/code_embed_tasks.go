package code

import (
	"fmt"
	"strings"
	"mdtk/base"
	"mdtk/taskset/grtask"
	"mdtk/args"
	"mdtk/lib"
)

func taskCmdsConstraint(cmds []string, args args.Args, errstr string, err error) (bool, grtask.GroupTask, args.Args, error) {
	if err != nil {
		return false, grtask.GroupTask(""), args, err
	}

	has_at := len(cmds) >= 1 && cmds[0] == "@"
	head_idx := lib.Btoi[int](has_at)

	if len(cmds) != head_idx + 1 {
		return has_at, grtask.GroupTask(""), args, fmt.Errorf("%s", errstr)
	}

	if args.HasValue("{$}", "<$>") {
		return has_at, grtask.GroupTask(""), args, fmt.Errorf("%s\n", "Args of '#task>' contains a value that cannot be used ({$}, <$>).")
	}

	return has_at, grtask.GroupTask(cmds[head_idx]), args, nil
}

func (code Code) ApplySubTasks(tf TaskDataSetInterface, nestsize int) (Code, error) {
	tasks := code.GetEmbedComment("task")

	if len(tasks) == 0 {
		return code, nil
	}

	indent := strings.Repeat(" ", 2)

	res := string(code)
	for i, task := range tasks {
		use_same_stack, gtname, args, err := taskCmdsConstraint(extractSubCmds(task[1]))
		if err != nil {
			return "", fmt.Errorf("%s-> %s\n", err, task[0])
		}
		head := ""

		base.DebugLogGreen(nestsize-1, fmt.Sprintf("#task(%d)>\n", i+1))
		subcode, err := tf.GetTask(gtname, args, false, !use_same_stack, nestsize-1)
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
