package help

import (
	"sort"
	"math"
	"mdtk/group"
	"mdtk/grtask"
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

func getEmbedDescTexts(tds taskset.TaskDataSet) taskset.TaskDataSet {
	res := tds
	for i, task := range tds.Data {
		desc := task.Code.GetEmbedDescText()
		if len(tds.Data[i].Description) == 1 && tds.Data[i].Description[0] == "" {
			if len(desc) != 0 {
				tds.Data[i].Description = desc
			} else {
				tds.Data[i].Description = []string{"----"}
			}
		} else {
			tds.Data[i].Description = append(tds.Data[i].Description, desc...)
		}
	}
	
	return res
}

func getEmbedArgsTexts(tds taskset.TaskDataSet) taskset.TaskDataSet {
	res := tds
	for i, task := range tds.Data {
		tds.Data[i].ArgsTexts = task.Code.GetEmbedArgsText()
	}
	
	return res
}

func getTaskNameMaxLength(tds taskset.TaskDataSet) int {
	buf_len := 0
	for _, task := range tds.Data {
		if len(task.Task) > buf_len {
			buf_len = len(task.Task)
		}
	}
	
	return buf_len
}

func getGroupMap(tds taskset.TaskDataSet, gtname grtask.GroupTask, show_private bool) map[string][]taskset.TaskData {
	group_map := map[string][]taskset.TaskData{}
	g, _, _ := gtname.Split()
	
	for _, t := range tds.Data {
		if !show_private && t.Group.IsPrivate() {
			continue
		}

		if g != "" && g != t.Group {
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

type groupArr struct {
	Name string
	Order int64
}
// Ensure that private groups are at the bottom
func getGroupArrAndSort(group_map map[string][]taskset.TaskData, group_order map[group.Group]int64) []groupArr {
	garr_pub := []groupArr{}
	garr_prv := []groupArr{}

	for k, _ := range group_map {
		o := int64(0)
		if k == "_" { o = math.MaxInt64 }
		if vv, ok := group_order[group.Group(k)]; ok { o = vv }
		
		if group.Group(k).IsPrivate() {
			garr_prv = append(garr_prv, groupArr{Name: k, Order: o})
		} else {
			garr_pub = append(garr_pub, groupArr{Name: k, Order: o})
		}
		
	}
	sort.Slice(garr_pub, func(i, j int) bool {
		if garr_pub[i].Order == garr_pub[j].Order {
			return garr_pub[i].Name < garr_pub[j].Name
		} else {
			return garr_pub[i].Order > garr_pub[j].Order
		}
	})
	sort.Slice(garr_prv, func(i, j int) bool {
		if garr_prv[i].Order == garr_prv[j].Order {
			return garr_prv[i].Name < garr_prv[j].Name
		} else {
			return garr_prv[i].Order > garr_prv[j].Order
		}
	})

	return append(garr_pub, garr_prv...)
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

