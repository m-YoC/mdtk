package read

import (
	"strings"
	"regexp"
	"mdtk/path"
)

var taskfile_head_rex = regexp.MustCompile("(?m)^" + block_reg + "taskfile[ \t]*$")


func (md Markdown) GetTaskfileBlockPath() ([]path.Path, error) {
	heads := taskfile_head_rex.FindAllStringSubmatch(string(md), -1)
	indices := taskfile_head_rex.FindAllStringIndex(string(md), -1)

	res := []path.Path{}
	for i, head := range heads {
		block := head[taskfile_head_rex.SubexpIndex("block")]

		code, err := md.ExtractCode(indices[i][1], block)
		if err != nil {
			return []path.Path{}, err
		}

		codearr := strings.Split(string(code), "\n")
	
		for _, v := range codearr {
			if tv := strings.TrimSpace(v); tv != "" {
				res = append(res, path.Path(tv))
			}
		}
	}
	
	return res, nil
}
