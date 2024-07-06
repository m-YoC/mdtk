package taskset

import (
	"mdtk/lib"
	"mdtk/taskset/path"
	"mdtk/taskset/task"
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_TaskDataSet(t *testing.T) {
	createDataSet := func() TaskDataSet {
		tds := TaskDataSet{}
		tds.Data = append(tds.Data, TaskData{Group: "group", Task: "task1", Code: "echo test1", Lang: "go"})
		tds.Data = append(tds.Data, TaskData{Group: "group", Task: "task2", Code: "echo test2", Lang: "fo"})
		tds.Data = append(tds.Data, TaskData{Group: "group", Task: "task3", Code: "echo test3", Lang: "ho"})
		tds.FilePath = map[path.Path]bool{"path1": true, "path2": false, "path3": true}
		return tds
	}

	t.Run("HasOnlyFilePathsAlreadyRead", func(t *testing.T) {
		tests := lib.TestCases[map[path.Path]bool, bool] {
			{Name: "values of path are only true", 
			TestArg: map[path.Path]bool{"path1": true, "path2": true, "path3": true}, 
			Expected: true},
			{Name: "values of path contain false", 
			TestArg: map[path.Path]bool{"path1": true, "path2": false, "path3": true}, 
			Expected: false},
		}

		tests.Run(t, func(t *testing.T, i int) {
			tds := createDataSet()
			tds.FilePath = tests[i].TestArg
			assert.Equal(t, tests[i].Expected, tds.HasOnlyFilePathsAlreadyRead())
		})
	})

	t.Run("Merge", func(t *testing.T) {
		type E struct {
			datasize int
			pathsize int
		}
		tests := lib.TestCases[map[path.Path]bool, E] {
			{Name: "merge with only different path", 
			TestArg: map[path.Path]bool{"path4": true, "path5": true, "path6": true}, 
			Expected: E{datasize: 6, pathsize: 6}},
			{Name: "merge with path contained same", 
			TestArg: map[path.Path]bool{"path1": true, "path4": false, "path3": true}, 
			Expected: E{datasize: 6, pathsize: 4}},
			{Name: "path values take precedence over true (Boolean OR operation)", 
			TestArg: map[path.Path]bool{"path1": true, "path2": false, "path3": false}, 
			Expected: E{datasize: 6, pathsize: 3}},
		}

		tests.Run(t, func(t *testing.T, i int) {
			tt := tests.Get(i)
			tds1 := createDataSet()
			tds2 := createDataSet()
			tds2.FilePath = tt.TestArg
			tds1.Merge(&tds2)
			assert.Equal(t, tt.Expected.datasize, len(tds1.Data))
			assert.Equal(t, tt.Expected.pathsize, len(tds1.FilePath))

			tds1b := createDataSet()
			for _, p := range []path.Path{"path1", "path2", "path3"} {
				assert.Equal(t, tds1b.FilePath[p] || tds2.FilePath[p], tds1.FilePath[p])
			}
		})
	})

	t.Run("RemovePathData", func(t *testing.T) {
		tds := createDataSet()
		tds.Data[0].FilePath = "aaaaa"
		tds.Data[1].FilePath = "bbbbb"
		tds.Data[2].FilePath = "ccccc"

		tds.RemovePathData("removed")

		for _, v := range tds.Data {
			assert.Equal(t, "removed", string(v.FilePath))
		}
		assert.Equal(t, 0, len(tds.FilePath))
		
	})

	t.Run("GetTaskData", func(t *testing.T) {
		type A struct {
			tname string
			re_tname2 string
			new_td TaskData
		} 

		tests := lib.TestCases[A, TaskData] {
			{Name: "get task1 code", 
			TestArg: A{tname: "task1"}, Expected: createDataSet().Data[0]},
			{Name: "Tasks with same name but different priorities", 
			TestArg: A{tname: "task1", new_td: TaskData{
				Group: "group", 
				Task: "task1", 
				Code: "echo newdata", 
				Lang: "go",
				Attributes: []string{"priority:2"},
			}},
			Expected: TaskData{
				Group: "group", 
				Task: "task1", 
				Code: "echo newdata", 
				Lang: "go",
				Attributes: []string{"priority:2"},
			}},
		}

		tests.Run(t, func(t *testing.T, i int) {
			tt := tests.Get(i)
			tds := createDataSet()

			if tt.TestArg.new_td.Task != "" { tds.Data = append(tds.Data, tt.TestArg.new_td) }

			td, err := tds.GetTaskData("group", task.Task(tt.TestArg.tname), nil)
			assert.NoError(t, err)
			assert.Equal(t, tt.Expected, td)
		})

		t.Run("Negative Cases", func(t *testing.T) {
			tests := lib.TestCases[A, TaskData] {
				{Name: "cannot get task00 code", 
				TestArg: A{tname: "task00"}},
				{Name: "has 2 task1 data", 
				TestArg: A{tname: "task1", re_tname2: "task1"}},
			}
	
			tests.Run(t, func(t *testing.T, i int) {
				tt := tests.Get(i)
				tds := createDataSet()

				if tt.TestArg.re_tname2 != "" {
					tds.Data[1].Task = task.Task(tt.TestArg.re_tname2)
				}
	
				_, err := tds.GetTaskData("group", task.Task(tt.TestArg.tname), nil)
				assert.Error(t, err)
			})
		})
	})


	t.Run("GetCode", func(t *testing.T) {
		type A struct {
			tname string
			re_tname2 string
		} 

		type E struct {
			c string
			l string
		}

		tests := lib.TestCases[A, E] {
			{Name: "get task1 code", 
			TestArg: A{tname: "task1"}, Expected: E{c: "echo test1", l: "go"}},
		}

		tests.Run(t, func(t *testing.T, i int) {
			tt := tests.Get(i)
			tds := createDataSet()

			c, l, err := tds.GetCode("group", task.Task(tt.TestArg.tname), nil)
			assert.NoError(t, err)
			assert.Equal(t, tt.Expected.c, string(c))
			assert.Equal(t, tt.Expected.l, string(l))
		})

		t.Run("Negative Cases", func(t *testing.T) {
			tests := lib.TestCases[A, E] {
				{Name: "cannot get task00 code", 
				TestArg: A{tname: "task00"}},
				{Name: "has 2 task1 data", 
				TestArg: A{tname: "task1", re_tname2: "task1"}},
			}
	
			tests.Run(t, func(t *testing.T, i int) {
				tt := tests.Get(i)
				tds := createDataSet()

				if tt.TestArg.re_tname2 != "" {
					tds.Data[1].Task = task.Task(tt.TestArg.re_tname2)
				}
	
				_, _, err := tds.GetCode("group", task.Task(tt.TestArg.tname), nil)
				assert.Error(t, err)
			})
		})

	})

}
