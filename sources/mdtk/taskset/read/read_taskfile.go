package read

import (
	"strings"
	"regexp"
	"mdtk/path"
)

var taskfile_head_rex = regexp.MustCompile("(?m)^" + block_reg + "taskfile[ \t]*$")


func (md Markdown) GetTaskfileBlockPath() (map[path.Path]bool, error) {
	return getTaskConfigBase[path.Path, bool](md, taskfile_head_rex, func(res *map[path.Path]bool, v string ) {
		if tv := strings.TrimSpace(v); tv != "" {
			(*res)[path.Path(tv)] = false
		}
	})
}
