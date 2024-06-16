package parse

import (
	"fmt"
	"mdtk/base"
)

// 最初の -- で分割する
func SplitArgs(strs []string) ([]string, []string) {
	for i := 0; i < len(strs); i++ {
		if strs[i] == "--" {
			if i+1 < len(strs) {
				return strs[:i], strs[i+1:]
			} else {
				return strs[:i], []string{}
			}
			
		}
	}

	return strs, []string{}
}

const (
	notOp = iota
	singleOp
	multiOps
)

func GetOpType(cmd string) int {
	cmdrunes := []rune(cmd)

	if cmdrunes[0] != '-' {
		return notOp
	}

	if len(cmdrunes) > 2 && cmdrunes[1] != '-' {
		return multiOps
	}

	return singleOp
}

func MatchOp(cmd string, flags Flag) ([]int, error) {
	cmdrunes := []rune(cmd)

	switch GetOpType(cmd) {
	case notOp:
		return nil, fmt.Errorf("[%s] is not option type.\n", cmd)

	case singleOp:
		idx := flags.GetIndex(cmd)
		if idx < 0 {
			return nil, fmt.Errorf("[%s] is invalid option.\n", cmd)
		}
		return []int{idx}, nil

	case multiOps:
		res := []int{}
		for _, v := range cmdrunes[1:] {
			idx := flags.GetIndex("-" + string(v))
			if idx < 0 {
				return nil, fmt.Errorf("[%s] is invalid option.\n", "-" + string(v))
			}
			res = append(res, idx)
		}
		return res, nil
	}

	return nil, fmt.Errorf("Do not come here.\n")
}


func Parse(args []string, flags Flag) (string, Flag, []string) {
	commands, task_args := SplitArgs(args)

	taskname := make([]string, 0, 10)

	for i := 1; i < len(commands); i++ {
		if commands[i][0:1] != "-" {
			taskname = append(taskname, commands[i])
			continue
		}

		fis, err := MatchOp(commands[i], flags)
		if err != nil {
			fmt.Printf("Parsing error: %v", err)
			base.MdtkExit(1)
		}

		c := 0
		for _, j := range fis {
			if flags[j].HasValue {
				if c >= 1 {
					fmt.Printf("Parsing error: %s includes 2 or more options with a value\n", commands[i-1])
					base.MdtkExit(1)
				}

				if i+1 >= len(commands) || GetOpType(commands[i+1]) != notOp {
					fmt.Printf("Parsing error: option %s is not set correctly.\n", flags[j].Name)
					base.MdtkExit(1)
				}

				flags[j].Exist = true
				flags[j].Value = commands[i+1]
				i++
				c++
			} else {
				flags[j].Exist = true
			}
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
        base.MdtkExit(1)
	}

	return res_taskname, flags, task_args
}
