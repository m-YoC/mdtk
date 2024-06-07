package help

import (
	"fmt"
	"strconv"
	"mdtk/lib"
	"mdtk/taskset"
	"mdtk/path"
	"mdtk/config"
)

func ShowGroups(filename path.Path, tds taskset.TaskDataSet, show_private bool) {
	group_map := getGroupMap(tds, "", show_private)
	group_arr := getGroupArrAndSort(group_map, tds.GroupOrder)

	maxlen := 0
	for _, kk := range group_arr {
		if len(kk.Name) > maxlen { maxlen = len(kk.Name) }
	}
	gwidth := max(10, maxlen + 2)
	unit := [2]string{" task", " tasks"}

	s := fmt.Sprintf(bgray + "[%s groups]" + clear + "\n", filename)

	for _, kk := range group_arr {
		k := kk.Name
		tasknum := len(group_map[k])
		s += fmt.Sprintf(cyan + "%-" + strconv.Itoa(gwidth) + "s", k)
		s += fmt.Sprintf(gray + ": " + strconv.Itoa(tasknum) + unit[lib.Btoi[int](tasknum > 1)] + clear + "\n")
	}

	PagerOutput(s, config.Config.PagerMinLimit)
}
