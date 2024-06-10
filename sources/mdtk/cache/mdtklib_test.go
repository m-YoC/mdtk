package cache

import (
	"strings"
	"mdtk/code"
	"mdtk/taskset"
	"mdtk/path"
	"testing"
	"github.com/stretchr/testify/assert"
)

var code1 = strings.Trim(`
echo hello
#embed> task2
`, "\n")

var code2 = strings.Trim(`
echo world
`, "\n")

var code12 = strings.Trim(`
echo hello
echo world
`, "\n")

func Test_expandPublicGroupTask(t *testing.T) {
	tds := taskset.TaskDataSet{Data: []taskset.TaskData{}}
	tds.Data = append(tds.Data, taskset.TaskData{Group: "_", Task: "task1", Code: code.Code(code1)})
	tds.Data = append(tds.Data, taskset.TaskData{Group: "_", Task: "task2", Code: code.Code(code2)})
	tds.Data = append(tds.Data, taskset.TaskData{Group: "_", Task: "task3", Code: code.Code(code1), Attributes: []string{taskset.ATTR_HIDDEN}})
	
	assert.False(t, tds.Data[0].HasAttr(taskset.ATTR_HIDDEN))
	assert.False(t, tds.Data[1].HasAttr(taskset.ATTR_HIDDEN))
	assert.True(t, tds.Data[2].HasAttr(taskset.ATTR_HIDDEN))
	assert.Equal(t, code12, string(expandPublicGroupTask(tds, 10).Data[0].Code))
	assert.Equal(t, code1, string(expandPublicGroupTask(tds, 10).Data[2].Code))
}

func Test_removePrivateGroupTask(t *testing.T) {
	tds := taskset.TaskDataSet{Data: []taskset.TaskData{}}
	tds.Data = append(tds.Data, taskset.TaskData{Group: "_", Task: "task1", Code: code.Code(code1)})
	tds.Data = append(tds.Data, taskset.TaskData{Group: "_private", Task: "task2", Code: code.Code(code2), Description: []string{""}})
	tds.Data[1].GetAttrsAndSet()

	assert.False(t, tds.Data[0].HasAttr(taskset.ATTR_HIDDEN))
	assert.True(t, tds.Data[1].HasAttr(taskset.ATTR_HIDDEN))
	assert.Equal(t, 1, len(removePrivateGroupTask(tds).Data))
}

func Test_cleanFilePath(t *testing.T) {
	tds := taskset.TaskDataSet{Data: []taskset.TaskData{}, FilePath: map[path.Path]bool{"Taskfile.md": true}}
	tds.Data = append(tds.Data, taskset.TaskData{Group: "_", Task: "task1", Code: code.Code(code1), FilePath: "Taskfile.md"})
	tds = cleanFilePath(tds, "testlib")
	assert.Equal(t, map[path.Path]bool{"testlib": true}, tds.FilePath)
	assert.Equal(t, "testlib", string(tds.Data[0].FilePath))
}
