package taskset

import (
	"mdtk/taskset/path"
	"mdtk/taskset/task"
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_TaskDataSet(t *testing.T) {
	createDataSet := func() TaskDataSet {
		tds := TaskDataSet{}
		tds.Data = append(tds.Data, TaskData{Group: "group", Task: "task1", Code: "echo test"})
		tds.Data = append(tds.Data, TaskData{Group: "group", Task: "task2", Code: "echo test"})
		tds.Data = append(tds.Data, TaskData{Group: "group", Task: "task3", Code: "echo test"})
		tds.FilePath = map[path.Path]bool{"path1": true, "path2": false, "path3": true}
		return tds
	}

	t.Run("HasOnlyFilePathsAlreadyRead", func(t *testing.T) {
		tests := []struct {
			name string
			filepath map[path.Path]bool
			expected bool
		} {
			{"values of path are only true", map[path.Path]bool{"path1": true, "path2": true, "path3": true}, true},
			{"values of path contain false", map[path.Path]bool{"path1": true, "path2": false, "path3": true}, false},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				tds := createDataSet()
				tds.FilePath = tt.filepath
				assert.Equal(t, tt.expected, tds.HasOnlyFilePathsAlreadyRead())
			})
		}
	})

	t.Run("Merge", func(t *testing.T) {
		tests := []struct {
			name string
			other_filepath map[path.Path]bool
			expected_datasize int
			expected_pathsize int
		} {
			{"merge with only different path", map[path.Path]bool{"path4": true, "path5": true, "path6": true}, 6, 6},
			{"merge with path contained same", map[path.Path]bool{"path1": true, "path4": false, "path3": true}, 6, 4},
			{"path values take precedence over true (Boolean OR operation)", map[path.Path]bool{"path1": true, "path2": false, "path3": false}, 6, 3},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				tds1 := createDataSet()
				tds2 := createDataSet()
				tds2.FilePath = tt.other_filepath
				tds1.Merge(&tds2)
				assert.Equal(t, tt.expected_datasize, len(tds1.Data))
				assert.Equal(t, tt.expected_pathsize, len(tds1.FilePath))

				tds1b := createDataSet()
				for _, p := range []path.Path{"path1", "path2", "path3"} {
					assert.Equal(t, tds1b.FilePath[p] || tds2.FilePath[p], tds1.FilePath[p])
				}
			})
		}
	})

	t.Run("GetCode", func(t *testing.T) {
		const (
			positive = iota
			negative
		)
		tests := []struct {
			name string
			taskname string
			rename_task2 string // if empty, skip set
			expected int
		} {
			{"get task1 code", "task1", "", positive},
			{"cannot get task00 code", "task00", "", negative},
			{"has 2 task1 data", "task1", "task1", negative},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				tds := createDataSet()
				if tt.rename_task2 != "" {
					tds.Data[1].Task = task.Task(tt.rename_task2)
				}

				c, _, err := tds.GetCode("group", task.Task(tt.taskname), nil)
				if tt.expected == positive {
					if assert.NoError(t, err) {
						assert.Equal(t, "echo test", string(c))
					}
				} else {
					assert.Error(t, err)
				}
				
			})
		}
	})

}
