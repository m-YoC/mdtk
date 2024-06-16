package code

import (
	"fmt"
	"strings"
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

	// TODO: check args. <$>, {$} are bad.
	return has_at, grtask.GroupTask(cmds[head_idx]), args, nil
}

func (code Code) ApplySubTasks(tf TaskDataSetInterface, nestsize int) (Code, error) {
	tasks := code.GetEmbedComment("task")

	if len(tasks) == 0 {
		return code, nil
	}

	indent := strings.Repeat(" ", 2)

	res := string(code)
	for _, task := range tasks {
		use_same_stack, gtname, args, err := taskCmdsConstraint(extractSubCmds(task[1]))
		if err != nil {
			return "", fmt.Errorf("%s-> %s\n", err, task[0])
		}
		head := ""

		subcode, err := tf.GetTask(gtname, args, false, !use_same_stack, nestsize-1)
		if err != nil {
			return "", err
		}
		subcode = subcode.RemoveEmbedDescComment().RemoveEmbedArgsComment()
		rsubcode := indent + strings.Replace(string(subcode), "\n", "\n" + indent, -1)
		execsubcode := head + "(  # " + string(gtname) + "\n" + rsubcode + "\n)"
		res = strings.Replace(res, task[0], execsubcode, 1)
	}

	return Code(res), nil
}
