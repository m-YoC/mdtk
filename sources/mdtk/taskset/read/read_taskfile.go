package read

import (
	"strings"
	"regexp"
	"mdtk/taskset/path"
)

var taskfile_head_rex = regexp.MustCompile("(?m)^" + block_reg + "taskfile[ \t]*$")


func (md Markdown) GetTaskfileBlockPath(base_abs_path path.Path) (map[path.Path]bool, error) {
	bdir := base_abs_path.Dir()

	path, err := getTaskConfigBase[path.Path, bool](md, taskfile_head_rex, func(res *map[path.Path]bool, v string ) {
		if tv := strings.TrimSpace(v); tv != "" {
			abspath := bdir.GetSubFilePath(path.Path(tv).ToSlash())
			(*res)[abspath] = false
		}
	})
	if err != nil {
		return path, err
	}
	path[base_abs_path] = true
	return path, nil
}
