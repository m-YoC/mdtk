package read

import (
	"strings"
	"strconv"
	"regexp"
	"mdtk/taskset/group"
	"mdtk/base"
)

var taskconfig_go_head_rex = regexp.MustCompile("(?m)^" + block_reg + "taskconfig:group-order[ \t]*$")
var taskconfig_go_kv_rex = regexp.MustCompile("^[ \t]*(?P<key>" + base.NameReg + ")[ \t]*:[ \t]*(?P<value>[^ \t]*)")

func (md Markdown) GetTaskConfigGroupOrder() (map[group.Group]int64, error) {
	return getTaskConfigBase[group.Group, int64](md, taskconfig_go_head_rex, func(res *map[group.Group]int64, v string) {
		if tv := strings.TrimSpace(v); taskconfig_go_kv_rex.MatchString(tv) {
			kv := taskconfig_go_kv_rex.FindStringSubmatch(tv)
			key := kv[taskconfig_go_kv_rex.SubexpIndex("key")]
			value := kv[taskconfig_go_kv_rex.SubexpIndex("value")]
			if vv, err := strconv.ParseInt(value, 10, 64); err == nil {
				(*res)[group.Group(key)] = vv
			}
		}
	})
}