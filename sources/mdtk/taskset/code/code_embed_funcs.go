package code

import (
	"fmt"
	"strings"
	"mdtk/base"
	"mdtk/taskset/grtask"
	"mdtk/args"
)


func funcCmdsConstraint(cmds []string, args args.Args, errstr string, err error) (string, grtask.GroupTask, args.Args, error) {
	if err != nil {
		return "", grtask.GroupTask(""), args, err
	}
	if len(cmds) != 2 {
		return "", grtask.GroupTask(""), args, fmt.Errorf("%s", errstr)
	}
	return cmds[0], grtask.GroupTask(cmds[1]), args, nil
}

func (code Code) ApplyFuncs(tf TaskDataSetInterface, args_enclose_with_quotes bool, brackets []string, nothing_cmd string, nestsize int) (Code, error) {
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

func (code Code) ApplyFuncsShell(tf TaskDataSetInterface, nestsize int) (Code, error) {
	return code.ApplyFuncs(tf, false, ParenTheses, ":", nestsize)
}

func (code Code) ApplyFuncsPwSh(tf TaskDataSetInterface, nestsize int) (Code, error) {
	return code.ApplyFuncs(tf,  true, CurlyBrackets, "? .", nestsize)
}
