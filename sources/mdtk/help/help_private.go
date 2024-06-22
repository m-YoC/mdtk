package help

import (
	"fmt"
	"sort"
	"math"
	"mdtk/taskset/group"
	"mdtk/taskset/grtask"
	"mdtk/taskset"
)


func doNotExistExplicitDefaultTask(tds taskset.TaskDataSet) bool {
	for _, task := range tds.Data {
		if string(task.Group) == "_" && string(task.Task) == "default" {
			return false
		}
	}

	return true
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
	g, _, _ := gtname.Split()

	const (
		TOP = iota
		MID
		BTM
	)
	data := [3][]taskset.TaskData{{}, {}, {}}

	for _, t := range tds.Data {
		if !show_private && t.HasAttr(taskset.AttrHidden) {
			continue
		}

		if g != "" && g != t.Group {
			continue
		}

		idx := MID
		if  t.HasAttr(taskset.AttrTop) && !t.HasAttr(taskset.AttrBottom) { idx = TOP }
		if !t.HasAttr(taskset.AttrTop) &&  t.HasAttr(taskset.AttrBottom) { idx = BTM }

		data[idx] = append(data[idx], t)
	}
	

	group_map := map[string][]taskset.TaskData{}
	
	for _, d := range data {
		for _, t := range d {
			if _, ok := group_map[string(t.Group)]; ok {
				group_map[string(t.Group)] = append(group_map[string(t.Group)], t)
			} else {
				group_map[string(t.Group)] = []taskset.TaskData{ t }
			}
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


type GroupTaskIntData map[string]int

func gtKey(td taskset.TaskData) string {
	return string(td.Group) + ":" + string(td.Task)
}

func gtpKey(td taskset.TaskData) string {
	return fmt.Sprintf("%s:%s:%d", string(td.Group), string(td.Task), td.GetPriority()) 
}

func getGroupTaskMaxPriority(tds taskset.TaskDataSet) GroupTaskIntData {
	res := GroupTaskIntData{}

	for _, t := range tds.Data {
		key := gtKey(t)
		p := t.GetPriority()
		if v, ok := res[key]; ok {
			if p > v {
				res[key] = p
			}
		} else {
			res[key] = p
		}
	}

	return res
}

func countGroupTaskName(tds taskset.TaskDataSet) GroupTaskIntData {
	max_priorities := getGroupTaskMaxPriority(tds)
	res := GroupTaskIntData{}

	for _, t := range tds.Data {
		p := t.GetPriority()
		maxp := max_priorities[gtKey(t)]
		if p == maxp {
			key := gtpKey(t)
			if v, ok := res[key]; ok {
				res[key] = v + 1
			} else {
				res[key] = 1
			}
		}
		
	}

	return res
}



