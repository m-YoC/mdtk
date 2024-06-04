package cache

import (
	"mdtk/grtask"
	"mdtk/taskset"
	"mdtk/args"
	"mdtk/path"
	"path/filepath"
)

func expandPublicGroupTask(tds taskset.TaskDataSet, nestsize int) taskset.TaskDataSet {

	for i, data := range tds.Data {
		if data.Group.IsPrivate() {
			continue
		}
		code := tds.GetTaskStart(grtask.Create(data.Group, data.Task), args.Args{}, nestsize)
		tds.Data[i].Code = code
	}

	return tds
}

func removePrivateGroupTask(tds taskset.TaskDataSet) taskset.TaskDataSet {
	d := []taskset.TaskData{}
	for _, data := range tds.Data {
		if !data.Group.IsPrivate() {
			d = append(d, data)
		}
	}
	tds.Data = d
	return tds
}

func cleanFilePath(tds taskset.TaskDataSet, s string) taskset.TaskDataSet {
	tds.FilePath = map[path.Path]bool{path.Path(s): true}
	for i, _ := range tds.Data {
		tds.Data[i].FilePath = path.Path(s)
	}
	return tds
}


func WriteLib(tds taskset.TaskDataSet, dir path.Path, output_namebase string, nestsize int) error {
	tdsb := expandPublicGroupTask(tds, nestsize)
	tdsb = removePrivateGroupTask(tdsb)
	tdsb = cleanFilePath(tdsb, output_namebase)
	return writeBase(tdsb, filepath.Join(string(dir), output_namebase + ".mdtklib"))
}

func ReadLib(filename path.Path) (taskset.TaskDataSet, error) {
	return readBase(string(filename))
}
