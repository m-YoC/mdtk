package read

import (
	"fmt"
	"strings"
	"regexp"
	"mdtk/base"
	"mdtk/group"
	"mdtk/task"
	"mdtk/code"
	"mdtk/taskset"
	"mdtk/path"
	"mdtk/config"
)

const block_reg = "(?P<block>`{3,}|~{3,})"
const code_reg = "(?P<code>.*?)"

var task_head_rex *regexp.Regexp

func getTaskHeadReg() string {
	greg := "(?P<group>(?:" + base.NameReg + ")?)"
	treg := "(?P<task>" + base.NameReg + ")"
	dspacer := "(?:-+)"
	dreg := "(?P<description>[^\n]*)"

	res := "task:" + greg + ":" + treg
	res += "(?:[ \t]+" + dspacer + ")?"
	res += "(?:[ \t]+" + dreg    + ")?"

	return res
	// return "task:" + greg + ":" + treg
}

func getProgTypeReg() string {
	progs := config.Config.LangAlias
	prog_reg := "(?:(?:" + strings.Join(progs, "|") + ")[ \t]+)?"
	return prog_reg
}

func GetTaskHeadRex() *regexp.Regexp {
	if task_head_rex == nil {
		task_head_rex = regexp.MustCompile("(?m)^" + block_reg + getProgTypeReg() + getTaskHeadReg() + "$")
	}
	return task_head_rex
}




type Markdown string

func (md Markdown) SimplifyNewline() Markdown {
	return Markdown(strings.Replace(string(md), "\r\n", "\n", -1))
}

func (md Markdown) ExtractCode(begin int, end_block string) (code.Code, error) {
// 見つからなかった場合を考慮する
	idx := strings.Index(string(md[begin:]), end_block)
	if idx == -1 {
		return code.Code(""), fmt.Errorf("Code block not closed.\n")
	}

	return code.Code(strings.Trim(string(md[begin : begin + idx]), "\n")), nil
}

func (md Markdown) GetTaskBlock() ([]taskset.TaskData, error) {
	rex := GetTaskHeadRex()
	heads := rex.FindAllStringSubmatch(string(md), -1)
	indices := rex.FindAllStringIndex(string(md), -1)

	res := []taskset.TaskData{}
	for i, head := range heads {
		block := head[rex.SubexpIndex("block")]
		c, err := md.ExtractCode(indices[i][1], block)
		if err != nil {
			return []taskset.TaskData{}, err
		}

		var task_data taskset.TaskData
		gbuf := head[rex.SubexpIndex("group")]
		if gbuf == "" { gbuf = "_" }
		task_data.Group = group.Group(gbuf)
		task_data.Task = task.Task(head[rex.SubexpIndex("task")])
		task_data.Description = []string{head[rex.SubexpIndex("description")]}
		task_data.Code = c

		res = append(res, task_data)
	}

	return res, nil
}


func ReadTask(filename path.Path) (taskset.TaskDataSet, error) {
	readTaskImpl := func(filename path.Path, is_root_file bool) (taskset.TaskDataSet, error) {
		tds := taskset.TaskDataSet{}
		base_abs_path := filename.GetFileAbsPath()
		base_dir := base_abs_path.Dir()

		md := ReadFile(filename).SimplifyNewline()
		tdarr, err1 := md.GetTaskBlock()
		tfp, err2 := md.GetTaskfileBlockPath()
		if err1 != nil {
			return taskset.TaskDataSet{}, fmt.Errorf("%w%s\n", err1, base_abs_path)
		}
		if err2 != nil {
			return taskset.TaskDataSet{}, fmt.Errorf("%w%s\n", err2, base_abs_path)
		}

		tds.Data = tdarr

		tds.FilePath = map[path.Path]bool{base_abs_path: true}
		for path, _ := range tfp {
			if f := base_dir.GetSubFilePath(path); f != base_abs_path {
				tds.FilePath[f] = false
			}
		}

		for i, _ := range tds.Data {
			tds.Data[i].FilePath = base_abs_path
		}

		if is_root_file {
			tdgo, err := md.GetTaskConfigGroupOrder()
			if err != nil {
				return taskset.TaskDataSet{}, fmt.Errorf("%w%s\n", err, base_abs_path)
			}
			tds.GroupOrder = tdgo
		}

		return tds, nil
	}

	tds, err := readTaskImpl(filename, true)
	if err != nil {
		return taskset.TaskDataSet{}, err
	}

	// Add Taskfile
	for !tds.HasOnlyFilePathsAlreadyRead() {
		for k, v := range tds.FilePath {
			if !v {
				sub_tds, errr := readTaskImpl(k, false)
				if errr != nil {
					return taskset.TaskDataSet{}, errr
				}

				tds.Merge(&sub_tds)
				tds.FilePath[k] = true
			}
		}
	}

	return tds, nil
}

