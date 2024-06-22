package help

import (
	"mdtk/taskset"
	"github.com/gookit/color"
)

func getEmbedDescTexts(tds taskset.TaskDataSet) taskset.TaskDataSet {
	res := tds
	for i, task := range tds.Data {
		desc := task.Code.GetEmbedDescText()
		if len(tds.Data[i].Description) == 1 && tds.Data[i].Description[0] == "" {
			if len(desc) != 0 {
				tds.Data[i].Description = desc
			} else {
				tds.Data[i].Description = []string{""}
			}
		} else {
			tds.Data[i].Description = append(tds.Data[i].Description, desc...)
		}

		p := task.GetPriority()
		if p != 0 {
			s := color.Gray.Sprintf(" (%+2d)", p)
			tds.Data[i].Description[0] = tds.Data[i].Description[0] + s
		}

		s := task.Lang.LangX().GetScriptNameColor()
		tds.Data[i].Description[0] = s + tds.Data[i].Description[0]

		if len(tds.Data[i].Description) == 1 && tds.Data[i].Description[0] == "" {
			tds.Data[i].Description = []string{"----"}
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


