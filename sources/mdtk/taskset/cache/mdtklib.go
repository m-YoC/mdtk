package cache

import (
	"fmt"
	"mdtk/base"
	"mdtk/lib"
	"mdtk/taskset"
	"mdtk/taskset/path"
	"mdtk/taskset/grtask"
	"mdtk/args"
	"path/filepath"
)

func expandPublicGroupTask(tds taskset.TaskDataSet, nestsize int) taskset.TaskDataSet {

	for i, data := range tds.Data {
		/*if data.Group.IsPrivate() {
			continue
		}*/
		if data.HasAttr(taskset.AttrHidden) {
			continue
		}
		code, err := tds.GetTaskStart(grtask.Create(data.Group, data.Task), args.Args{}, nestsize)
		if err != nil {
			fmt.Print(err)
			base.MdtkExit(1)
		}
		tds.Data[i].Code = code
	}

	return tds
}

func removePrivateGroupTask(tds taskset.TaskDataSet) taskset.TaskDataSet {
	d := []taskset.TaskData{}
	for _, data := range tds.Data {
		/*if !data.Group.IsPrivate() {
			d = append(d, data)
		}*/
		if !data.HasAttr(taskset.AttrHidden) {
			d = append(d, data)
		}
	}
	tds.Data = d
	return tds
}

func WriteLib(tds taskset.TaskDataSet, dir path.Path, output_namebase string, nestsize int) error {
	tdsb := expandPublicGroupTask(tds, nestsize)
	tdsb = removePrivateGroupTask(tdsb)
	tdsb.RemovePathData(output_namebase)
	return lib.WriteStruct[taskset.TaskDataSet](tdsb, filepath.Join(dir.String(), output_namebase + ".mdtklib"))
}

func ReadLib(filename path.Path) (taskset.TaskDataSet, error) {
	return lib.ReadStruct[taskset.TaskDataSet](filename.String())
}
