package code

import (
	"fmt"
	"strings"
	"mdtk/grtask"
	"mdtk/args"
	"mdtk/parse"
)

func getEmbedCommentTaskAndArgs(embed_comment string) (bool, grtask.GroupTask, args.Args, error) {
	res := parse.MustLexArgString(embed_comment)
	head, bottom := parse.SplitArgs(res)

	head_idx := 0
	has_at := false

	if len(head) >= 1 && head[0] == "@" {
		head_idx = 1
		has_at = true
	}

	if len(head) == head_idx + 1 {
		return has_at, grtask.GroupTask(head[head_idx]), args.ToArgs(bottom...), nil
	}

	if len(head) < head_idx + 1 {
		s := fmt.Sprintln("Parsing error: too few words.")
		s += fmt.Sprintln("Bad embed task comment.")
		return has_at, "", args.Args{}, fmt.Errorf("%s", s)
	}

	s := fmt.Sprintln("Parsing error: too many words.")
	s += fmt.Sprintln("Bad embed task comment.")
	return has_at, "", args.Args{}, fmt.Errorf("%s", s)
}

func (code Code) ApplySubTasks(tf TaskDataSetInterface, nestsize int) (Code, error) {
	tasks := code.GetEmbedComment("task")

	if len(tasks) == 0 {
		return code, nil
	}

	indent := strings.Repeat(" ", 2)

	res := string(code)
	for _, task := range tasks {
		use_new_task_stack, grtaskname, args, err1 := getEmbedCommentTaskAndArgs(task[1])
		if err1 != nil {
			return "", err1
		}

		subcode, err2 := tf.GetTask(grtaskname, args, false, use_new_task_stack, nestsize-1)
		if err2 != nil {
			return "", err2
		}
		// subcode = subcode.RemoveEmbedArgsComment()
		rsubcode := indent + strings.Replace(string(subcode), "\n", "\n" + indent, -1)
		execsubcode := "\n# subtask: " + string(grtaskname) + "\n(\n" + rsubcode + "\n) # end: " + string(grtaskname) + "\n"
		res = strings.Replace(res, task[0], execsubcode, 1)
	}

	return Code(res), nil
}
