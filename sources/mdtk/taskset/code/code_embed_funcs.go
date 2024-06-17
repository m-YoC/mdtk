package code

import (
	"fmt"
	"strings"
	"mdtk/taskset/grtask"
	"mdtk/args"
)

var ParenTheses = []string{"(", ")"}
var CurlyBrackets = []string{"{", "}"}

func funcCmdsConstraint(cmds []string, args args.Args, errstr string, err error) (string, grtask.GroupTask, args.Args, error) {
	if err != nil {
		return "", grtask.GroupTask(""), args, err
	}
	if len(cmds) != 2 {
		return "", grtask.GroupTask(""), args, fmt.Errorf("%s", errstr)
	}
	return cmds[0], grtask.GroupTask(cmds[1]), args, nil
}

func (code Code) ApplyFuncs(tf TaskDataSetInterface, brackets []string, nestsize int) (Code, error) {
	tasks := code.GetEmbedComment("func")

	if len(tasks) == 0 {
		return code, nil
	}

	indent := strings.Repeat(" ", 2)

	res := string(code)
	for _, task := range tasks {
		fname, gtname, args, err := funcCmdsConstraint(extractSubCmds(task[1]))
		if err != nil {
			return "", fmt.Errorf("%s-> %s\n", err, task[0])
		}
		use_new_task_stack := true
		head := "function " + fname + "() "

		subcode, err := tf.GetTask(gtname, args, false, use_new_task_stack, nestsize-1)
		if err != nil {
			return "", err
		}
		subcode = subcode.RemoveEmbedDescComment().RemoveEmbedArgsComment()
		rsubcode := indent + strings.Replace(string(subcode), "\n", "\n" + indent, -1)
		execsubcode := head + brackets[0] + "  # " + string(gtname) + "\n" + rsubcode + "\n" + brackets[1]
		res = strings.Replace(res, task[0], execsubcode, 1)
	}

	return Code(res), nil
}
