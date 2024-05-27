package parse

import (
	"fmt"
	"os"
)

// 最初の -- で分割する
func SplitArgs(strs []string) ([]string, []string) {
	for i := 0; i < len(strs); i++ {
		if strs[i] == "--" {
			return strs[:i], strs[i+1:]
		}
	}

	return strs, []string{}
}

func Parse(args []string, flags Flag) (string, Flag, []string) {
	commands, task_args := SplitArgs(args)

	taskname := make([]string, 0, 10)

	for i := 1; i < len(commands); i++ {
		if commands[i][0:1] != "-" {
			taskname = append(taskname, commands[i])
			continue
		}

		for j, fd := range flags {
			if !fd.MatchName(commands[i]) {
				continue
			}

			if fd.HasValue {
				if i+1 >= len(commands) {
					fmt.Printf("Parsing error: option %s is not set correctly.\n", fd.Name)
					os.Exit(1)
				}

				flags[j].Exist = true
				flags[j].Value = commands[i+1]
				i++
			} else {
				flags[j].Exist = true
			}

			break
		}

	}

	res_taskname := ""
	switch len(taskname) {
	case 0: // no taskname
		res_taskname = "default"
	case 1: // only taskname or groupname:taskname
		res_taskname = taskname[0]
	case 2: // groupname and taskname
		res_taskname = taskname[0] + ":" + taskname[1]
	default:
		fmt.Println("Parsing error: too many words (excluding options).")
		fmt.Println(taskname)
        os.Exit(1)
	}

	return res_taskname, flags, task_args
}
