package help

import (
	"strconv"
	"mdtk/lib"
	"mdtk/taskset"
	"mdtk/config"
	"github.com/gookit/color"
)

func ShowGroups(filename string, tds taskset.TaskDataSet, show_private bool) {
	group_map := getGroupMap(tds, "", show_private)
	group_arr := getGroupArrAndSort(group_map, tds.GroupOrder)

	maxlen := 0
	for _, kk := range group_arr {
		if len(kk.Name) > maxlen { maxlen = len(kk.Name) }
	}
	gwidth := max(10, maxlen + 2)
	unit := [2]string{" task", " tasks"}

	s := color.Style{color.FgGray, color.OpBold}.Sprintf("[%s groups]\n", filename)

	for _, kk := range group_arr {
		k := kk.Name
		tasknum := len(group_map[k])
		s += color.Cyan.Sprintf("%-" + strconv.Itoa(gwidth) + "s", k)
		s += color.Gray.Sprintf(": " + strconv.Itoa(tasknum) + unit[lib.Btoi[int](tasknum > 1)]) + "\n"
	}

	PagerOutput(s, config.Config.PagerMinLimit)
}
