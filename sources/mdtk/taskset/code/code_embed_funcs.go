package code

import (
	"fmt"
	"strings"
	"mdtk/taskset/grtask"
	"mdtk/args"
	"mdtk/parse"
)

func getEmbedCommentFuncsData(embed_comment string) (string, grtask.GroupTask, args.Args, error) {
	res := parse.MustLexArgString(embed_comment)
	head, bottom := parse.SplitArgs(res)

	if len(head) != 2 {
		s := fmt.Sprintln("Parsing error: too many or too few words.")
		s += fmt.Sprintln("Bad embed func comment.")
		return "", "", args.Args{}, fmt.Errorf("%s", s)
	}
 
	fname := head[0]
	return fname, grtask.GroupTask(head[1]), args.ToArgs(bottom...), nil
}

func (code Code) ApplyFuncs(tf TaskDataSetInterface, nestsize int) (Code, error) {
	tasks := code.GetEmbedComment("func")

	if len(tasks) == 0 {
		return code, nil
	}

	indent := strings.Repeat(" ", 2)

	res := string(code)
	for _, task := range tasks {
		fname, gtname, args, err1 := getEmbedCommentFuncsData(task[1])
		use_new_task_stack := true
		head := "function " + fname + "() "

		if err1 != nil {
			return "", err1
		}

		subcode, err2 := tf.GetTask(gtname, args, false, use_new_task_stack, nestsize-1)
		if err2 != nil {
			return "", err2
		}
		subcode = subcode.RemoveEmbedDescComment().RemoveEmbedArgsComment()
		rsubcode := indent + strings.Replace(string(subcode), "\n", "\n" + indent, -1)
		execsubcode := head + "(  # " + string(gtname) + "\n" + rsubcode + "\n)"
		res = strings.Replace(res, task[0], execsubcode, 1)
	}

	return Code(res), nil
}
