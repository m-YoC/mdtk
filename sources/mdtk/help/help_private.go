package help

import (
	"sort"
	"mdtk/group"
	"mdtk/taskset"
)

const (
	gray = "\033[30m"
	bgray = "\033[1;30m"
	cyan = "\033[36m"
	magenta = "\033[35m"
	bmagenta = "\033[1;35m"
	clear = "\033[0m"
)

func doNotExistExplicitDefaultTask(tds taskset.TaskDataSet) bool {
	for _, task := range tds.Data {
		if string(task.Group) == "_" && string(task.Task) == "default" {
			return false
		}
	}

	return true
}


func getEmbedArgsTexts(tds taskset.TaskDataSet) taskset.TaskDataSet {
	res := tds
	for i, task := range tds.Data {
		tds.Data[i].ArgsTexts = []string{}

		embeds := task.Code.GetEmbedComment("args")

		if len(embeds) == 0 {
			continue
		}
	
		for _, embed := range embeds {
			tds.Data[i].ArgsTexts = append(tds.Data[i].ArgsTexts, embed[1])
		}
	}
	
	return res
}

func getTaskNameMaxLength(tds taskset.TaskDataSet) taskset.TaskDataSet {
	buf_len := 0
	for _, task := range tds.Data {
		if len(task.Task) > buf_len {
			buf_len = len(task.Task)
		}
	}

	tds.TaskNameMaxLength = buf_len
	
	return tds
}

func getGroupMap(tds taskset.TaskDataSet, show_private bool) map[string][]taskset.TaskData {
	group_map := map[string][]taskset.TaskData{}
	
	for _, t := range tds.Data {
		if !show_private && t.Group.IsPrivate() {
			continue
		}

		if _, ok := group_map[string(t.Group)]; ok {
			group_map[string(t.Group)] = append(group_map[string(t.Group)], t)
		} else {
			group_map[string(t.Group)] = []taskset.TaskData{ t }
		}
	}

	return group_map
}

// Ensure that private groups are at the bottom
func getGroupArrAndSort(group_map map[string][]taskset.TaskData) []string {
	group_arr_public := []string{}
	group_arr_private := []string{}

	for k, _ := range group_map {
		if group.Group(k).IsPrivate() {
			group_arr_private = append(group_arr_private, k)
		} else {
			group_arr_public = append(group_arr_public, k)
		}
		
	}
	sort.Strings(group_arr_public)
	sort.Strings(group_arr_private)

	return append(group_arr_public, group_arr_private...)
}

func countGroupTaskName(tds taskset.TaskDataSet) map[string]int {
	res := map[string]int{}

	for _, t := range tds.Data {
		key := string(t.Group) + ":" + string(t.Task)
		if v, ok := res[key]; ok {
			res[key] = v + 1
		} else {
			res[key] = 1
		}
	}

	return res
}

