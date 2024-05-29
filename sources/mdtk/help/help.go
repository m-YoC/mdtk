package help

import (
	"fmt"
	"strconv"
	"mdtk/grtask"
	"mdtk/taskset"
	"mdtk/path"
)

func ShouldShowHelp(gtname grtask.GroupTask, tds taskset.TaskDataSet) bool {
	return string(gtname) == "help" || (string(gtname) == "default" && doNotExistExplicitDefaultTask(tds))
}

func ShowHelp(filename path.Path, tds taskset.TaskDataSet, show_private bool) {
	tds = getEmbedArgsTexts(tds)
	tds = getTaskNameMaxLength(tds)

	group_map := getGroupMap(tds, show_private)
	group_arr := getGroupArrAndSort(group_map)

	counts := countGroupTaskName(tds)

	taskname_width := max(20, tds.TaskNameMaxLength + 7)

	getTaskNameFormatStr := func(head_str string) string {
		return gray + head_str + cyan + "%-" + strconv.Itoa(taskname_width - len(head_str)) + "s" + clear
	}

	// For displaying validation error
	getTaskNameFormatStr2 := func(head_str string) string {
		return gray + head_str + bmagenta + "! %-" + strconv.Itoa(taskname_width - len(head_str) - 2) + "s" + clear
	}

	getDescFormatStr :=func(color string, head string) string {
		return color + head + "%s" + clear
	}

	group_format_str := getDescFormatStr(gray, "") + "\n"
	plane_task_format_str := getTaskNameFormatStr("") + getDescFormatStr(clear, "") + "\n"
	group_task_format_str := getTaskNameFormatStr("\\_ ") + getDescFormatStr(clear, "") + "\n"

	// For displaying validation error
	plane_task_format_str2 := getTaskNameFormatStr2("") + getDescFormatStr(magenta, "path: ") + "\n"
	group_task_format_str2 := getTaskNameFormatStr2("\\_ ") + getDescFormatStr(magenta, "path: ") + "\n"

	args_format_str1 := getTaskNameFormatStr("") + getDescFormatStr(gray, "args: ") + "\n"
	args_format_str2 := getTaskNameFormatStr("|") + getDescFormatStr(gray, "args: ") + "\n"

	args_format_str := [2]string{args_format_str1, args_format_str2}

	// ------ Print -------------------------------------

	s := fmt.Sprintf(bgray + "[%s help]" + clear + "\n", filename)

	for _, k := range group_arr {
		if k == "_" {
			for _, t := range group_map["_"] {
				if counts[k + ":" + string(t.Task)] > 1 {
					s += fmt.Sprintf(plane_task_format_str2, t.Task, t.FilePath)
					continue
				}

				s += fmt.Sprintf(plane_task_format_str, t.Task, t.Description)

				for _, a := range t.ArgsTexts {
					s += fmt.Sprintf(args_format_str[0], "", a)
				}
			}

		} else {
			s += fmt.Sprintf(group_format_str, k)

			for i, t := range group_map[k] {
				if counts[k + ":" + string(t.Task)] > 1 {
					s += fmt.Sprintf(group_task_format_str2, t.Task, t.FilePath)
					continue
				}

				s += fmt.Sprintf(group_task_format_str, t.Task, t.Description)

				idx := -(i + 1) / len(group_map[k]) + 1
				for _, a := range t.ArgsTexts {
					s += fmt.Sprintf(args_format_str[idx], "", a)
				}
				
			}
			
		}
		
		if len(group_map[k][len(group_map[k])-1].ArgsTexts) == 0 {
			s += fmt.Sprintln("")
		} 
	}

	PagerOutput(s, 40)
}
