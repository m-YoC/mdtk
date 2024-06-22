package code

import (
	"fmt"
	"strings"
	"mdtk/base"
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

func (code Code) ApplyFuncs(tf TaskDataSetInterface, brackets []string, nothing_cmd string, nestsize int) (Code, error) {
	tasks := code.GetEmbedComment("func")

	if len(tasks) == 0 {
		return code, nil
	}

	indent := strings.Repeat(" ", 2)

	res := string(code)
	for i, task := range tasks {
		fname, gtname, args, err := funcCmdsConstraint(extractSubCmds(task[1]))
		if err != nil {
			return "", fmt.Errorf("%s-> %s\n", err, task[0])
		}
		use_new_task_stack := true
		head := "function " + fname + "() "

		base.DebugLogGreen(nestsize-1, fmt.Sprintf("#func(%d)> %s\n", i+1, fname))
		subcode, err := tf.GetTask(gtname, args, false, use_new_task_stack, nestsize-1)
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
