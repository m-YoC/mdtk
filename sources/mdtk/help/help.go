package help

import (
	"fmt"
	"strconv"
	"mdtk/lib"
	"mdtk/taskset/grtask"
	"mdtk/taskset"
	"mdtk/config"
	"github.com/gookit/color"
)

func ShouldShowHelp(gtname grtask.GroupTask, tds taskset.TaskDataSet) bool {
	_, t, _ := gtname.Split()
	return string(t) == "help" || (string(gtname) == "default" && doNotExistExplicitDefaultTask(tds))
}

func ShowHelp(filename string, gtname grtask.GroupTask, tds taskset.TaskDataSet, show_private bool) {
	tds = getEmbedDescTexts(tds)
	tds = getEmbedArgsTexts(tds)

	group_map := getGroupMap(tds, gtname, show_private)
	group_arr := getGroupArrAndSort(group_map, tds.GroupOrder)

	counts := countGroupTaskName(tds)

	taskNameMaxLength := getTaskNameMaxLength(tds)
	taskname_width := max(20, taskNameMaxLength + 7)

	getTaskNameFormatStr := func(head_str string) string {
		s := color.Gray.Sprint(head_str)
		s += color.Cyan.Sprint("%-" + strconv.Itoa(taskname_width - len(head_str)) + "s")
		return s
	}

	// For displaying validation error
	getTaskNameFormatStr2 := func(head_str string) string {
		s := color.Gray.Sprint(head_str)
		s += color.Style{color.FgMagenta, color.OpBold}.Sprint("! %-" + strconv.Itoa(taskname_width - len(head_str) - 2) + "s")
		return s
	}

	getDescFormatStr :=func(cstyle color.Color, head string) string {
		return cstyle.Sprint(head + "%s")
	}

	group_format_str := getDescFormatStr(color.Gray, "") + "\n"
	plane_task_format_str := getTaskNameFormatStr("") + getDescFormatStr(color.Normal, "") + "\n"
	group_task_format_str := getTaskNameFormatStr("\\_ ") + getDescFormatStr(color.Normal, "") + "\n"
	task_format_str := [2]string{plane_task_format_str, group_task_format_str}

	// For displaying validation error
	plane_task_format_str2 := getTaskNameFormatStr2("") + getDescFormatStr(color.Magenta, "path: ") + "\n"
	group_task_format_str2 := getTaskNameFormatStr2("\\_ ") + getDescFormatStr(color.Magenta, "path: ") + "\n"
	task_format_str2 := [2]string{plane_task_format_str2, group_task_format_str2}

	desc_format_str1 := getTaskNameFormatStr("") + getDescFormatStr(color.Normal, "") + "\n"
	desc_format_str2 := getTaskNameFormatStr("|") + getDescFormatStr(color.Normal, "") + "\n"
	desc_format_str := [2]string{desc_format_str1, desc_format_str2}

	args_format_str1 := getTaskNameFormatStr("") + getDescFormatStr(color.Gray, "args: ") + "\n"
	args_format_str2 := getTaskNameFormatStr("|") + getDescFormatStr(color.Gray, "args: ") + "\n"
	args_format_str := [2]string{args_format_str1, args_format_str2}

	// ------ Print -------------------------------------

	s := color.Style{color.FgGray, color.OpBold}.Sprintf("[%s help]\n", filename)

	for _, gg := range group_arr {
		g := gg.Name
		n := 0
		if g != "_" {
			n = 1
			s += color.Gray.Sprintf(group_format_str, g)
		}
		
		for i, t := range group_map[g] {
			if counts[gtpKey(t)] > 1 {
				s += fmt.Sprintf(task_format_str2[n], t.Task, t.FilePath)
				continue
			}

			s += fmt.Sprintf(task_format_str[n], t.Task, t.Description[0])

			idx := lib.Btoi[int](n == 1 && (i + 1) < len(group_map[g]))

			if len(t.Description) > 1 {
				for _, d := range t.Description[1:] {
					s += fmt.Sprintf(desc_format_str[idx], "", d)
				}
			}
			for _, a := range t.ArgsTexts {
				s += fmt.Sprintf(args_format_str[idx], "", a)
			}
			
		}
		
		back := group_map[g][len(group_map[g])-1]
		if len(back.Description) < 2 && len(back.ArgsTexts) == 0 {
			s += fmt.Sprintln("")
		} 
	}

	PagerOutput(s, config.Config.PagerMinLimit)
}
