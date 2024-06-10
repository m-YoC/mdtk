package taskset

import (
	"fmt"
	"mdtk/group"
	"mdtk/task"
	"mdtk/code"
	"mdtk/path"
)

type TaskDataSet struct {
	Data []TaskData
	FilePath map[path.Path]bool
	GroupOrder map[group.Group]int64
}


func (tds TaskDataSet) HasOnlyFilePathsAlreadyRead() bool {
	f := true
	for _, v := range tds.FilePath {
		f = f && v
	}

	return f
}

func (tds *TaskDataSet) Merge(tds2 *TaskDataSet) {
	tds.Data = append(tds.Data, tds2.Data...)

	for k, v := range tds2.FilePath {
		if vv, ok := tds.FilePath[k]; !ok || !vv {
			tds.FilePath[k] = v
		}
	}
}

func (tds TaskDataSet) RemovePathData(set_str string) TaskDataSet {
	tds.FilePath = map[path.Path]bool{}
	
	for i, _ := range tds.Data {
		tds.Data[i].FilePath = path.Path(set_str)
	}
	return tds
}

func (tds TaskDataSet) GetTaskData(gname group.Group, tname task.Task, err error) (TaskData, error) {
	if err != nil {
		return TaskData{}, err
	}

	found := []TaskData{}

	for _, t := range tds.Data {
		if t.Group.Match(gname) && t.Task.Match(tname) {
			found = append(found, t)
		}
	}

	if len(found) == 0 {
		s := fmt.Sprintln("Do not find task.")
		s += fmt.Sprintln("wanted =>", "group:", gname, "/ task:", tname)
		return TaskData{}, fmt.Errorf("%s", s)
	}

	if len(found) > 1 {
		s := fmt.Sprintln("Task cannot be identified.")
		s += fmt.Sprintln("wanted =>", "group:", gname, "/ task:", tname)
		s += fmt.Sprintf("  %-10s  %-10s  %s\n", "[group]", "[task]", "[filepath]")
		for _, v := range found {
			s += fmt.Sprintf("- %-10s  %-10s  %s\n", v.Group, v.Task, v.FilePath)
		}
		
		return TaskData{}, fmt.Errorf("%s", s)
	}

	return found[0], nil
}


func (tds TaskDataSet) GetCode(gname group.Group, tname task.Task, err error) (code.Code, error) {
	td, err := tds.GetTaskData(gname, tname, err)
	if err != nil {
		return "", err
	}

	return td.Code, nil
}

