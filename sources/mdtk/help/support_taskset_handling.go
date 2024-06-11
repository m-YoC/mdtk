package help

import (
	"mdtk/taskset"
)

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

		if task.Lang != taskset.ShellLangs {
			tds.Data[i].Description[0] = blue + "<" + task.Lang + "> " + clear + tds.Data[i].Description[0]
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


