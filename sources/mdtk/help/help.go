package help

import (
	"fmt"
	"strconv"
	"mdtk/lib"
	"mdtk/grtask"
	"mdtk/taskset"
	"mdtk/path"
	"mdtk/config"
)

func ShouldShowHelp(gtname grtask.GroupTask, tds taskset.TaskDataSet) bool {
	_, t, _ := gtname.Split()
	return string(t) == "help" || (string(gtname) == "default" && doNotExistExplicitDefaultTask(tds))
}

func ShowHelp(filename path.Path, gtname grtask.GroupTask, tds taskset.TaskDataSet, show_private bool) {
	tds = getEmbedDescTexts(tds)
	tds = getEmbedArgsTexts(tds)

	group_map := getGroupMap(tds, gtname, show_private)
	group_arr := getGroupArrAndSort(group_map, tds.GroupOrder)

	counts := countGroupTaskName(tds)

	taskNameMaxLength := getTaskNameMaxLength(tds)
	taskname_width := max(20, taskNameMaxLength + 7)

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
	task_format_str := [2]string{plane_task_format_str, group_task_format_str}

	// For displaying validation error
	plane_task_format_str2 := getTaskNameFormatStr2("") + getDescFormatStr(magenta, "path: ") + "\n"
	group_task_format_str2 := getTaskNameFormatStr2("\\_ ") + getDescFormatStr(magenta, "path: ") + "\n"
	task_format_str2 := [2]string{plane_task_format_str2, group_task_format_str2}

	desc_format_str1 := getTaskNameFormatStr("") + getDescFormatStr(clear, "") + "\n"
	desc_format_str2 := getTaskNameFormatStr("|") + getDescFormatStr(clear, "") + "\n"
	desc_format_str := [2]string{desc_format_str1, desc_format_str2}

	args_format_str1 := getTaskNameFormatStr("") + getDescFormatStr(gray, "args: ") + "\n"
	args_format_str2 := getTaskNameFormatStr("|") + getDescFormatStr(gray, "args: ") + "\n"
	args_format_str := [2]string{args_format_str1, args_format_str2}

	// ------ Print -------------------------------------

	s := fmt.Sprintf(bgray + "[%s help]" + clear + "\n", filename)

	for _, kk := range group_arr {
		k := kk.Name
		n := 0
		if k != "_" {
			n = 1
			s += fmt.Sprintf(group_format_str, k)
		}
		
		for i, t := range group_map[k] {
			if counts[k + ":" + string(t.Task)] > 1 {
				s += fmt.Sprintf(task_format_str2[n], t.Task, t.FilePath)
				continue
			}

			s += fmt.Sprintf(task_format_str[n], t.Task, t.Description[0])

			idx := lib.Btoi[int](n == 1 && (i + 1) < len(group_map[k]))

			if len(t.Description) > 1 {
				for _, d := range t.Description[1:] {
					s += fmt.Sprintf(desc_format_str[idx], "", d)
				}
			}
			for _, a := range t.ArgsTexts {
				s += fmt.Sprintf(args_format_str[idx], "", a)
			}
			
		}
		
		back := group_map[k][len(group_map[k])-1]
		if len(back.Description) < 2 && len(back.ArgsTexts) == 0 {
			s += fmt.Sprintln("")
		} 
	}

	PagerOutput(s, config.Config.PagerMinLimit)
}
