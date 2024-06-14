package read

import (
	"fmt"
	"strings"
	"regexp"
	"mdtk/base"
	"mdtk/taskset/code"
	"mdtk/taskset"
	"mdtk/taskset/path"
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
	progs := append(config.Config.LangAlias, config.Config.LangAliasSub...)
	prog_reg := "(?:(?P<lang>" + strings.Join(progs, "|") + ")[ \t]+)?"
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
	idx := strings.Index(string(md[begin:]), end_block)
	if idx == -1 {
		return code.Code(""), fmt.Errorf("Code block not closed.\n")
	}

	return code.Code(strings.Trim(string(md[begin : begin + idx]), "\n")), nil
}

func (md Markdown) GetTaskBlock(filepath path.Path) ([]taskset.TaskData, error) {
	rex := GetTaskHeadRex()
	heads := rex.FindAllStringSubmatch(string(md), -1)
	indices := rex.FindAllStringIndex(string(md), -1)

	res := []taskset.TaskData{}
	for i, head := range heads {
		// get fence (```, ~~~, ````, ~~~~, ...)
		block := head[rex.SubexpIndex("block")]
		// get code area
		c, err := md.ExtractCode(indices[i][1], block)
		if err != nil {
			return []taskset.TaskData{}, err
		}

		// create TaskData
		var task_data taskset.TaskData
		task_data.SetLang(head[rex.SubexpIndex("lang")])
		task_data.SetGroup(head[rex.SubexpIndex("group")])
		task_data.SetTask(head[rex.SubexpIndex("task")])
		task_data.SetDescription(head[rex.SubexpIndex("description")])
		task_data.Code = c
		task_data.FilePath = filepath

		task_data.GetAttrsAndSet()

		res = append(res, task_data)
	}

	return res, nil
}


func ReadTask(filename path.Path) (taskset.TaskDataSet, error) {
	nilset := taskset.TaskDataSet{}

	readTaskImpl := func(filename path.Path, is_root_file bool) (taskset.TaskDataSet, error) {
		md := ReadFile(filename).SimplifyNewline()
		base_abs_path := filename.GetFileAbsPath()

		tdarr, err := md.GetTaskBlock(base_abs_path)
		if err != nil {
			return nilset, fmt.Errorf("%w%s\n", err, base_abs_path)
		}

		tfp, err := md.GetTaskfileBlockPath(base_abs_path)
		if err != nil {
			return nilset, fmt.Errorf("%w%s\n", err, base_abs_path)
		}

		tds := taskset.TaskDataSet{Data: tdarr, FilePath: tfp}

		if is_root_file {
			tdgo, err := md.GetTaskConfigGroupOrder()
			if err != nil {
				return nilset, fmt.Errorf("%w%s\n", err, base_abs_path)
			}
			tds.GroupOrder = tdgo
		}

		return tds, nil
	}

	tds, err := readTaskImpl(filename, true)
	if err != nil {
		return nilset, err
	}

	// Add Taskfile
	for !tds.HasOnlyFilePathsAlreadyRead() {
		for k, v := range tds.FilePath {
			if v { continue }

			sub_tds, err := readTaskImpl(k, false)
			if err != nil {
				return nilset, err
			}

			tds.Merge(&sub_tds)
			tds.FilePath[k] = true
		}
	}

	return tds, nil
}

