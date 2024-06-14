package grtask

import (
	"fmt"
	"mdtk/taskset/group"
	"mdtk/taskset/task"
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_GroupTask(t *testing.T) {
	const (
		positive = "positive"
		negative = "negative"
	)

	type ValidateTestType struct {
		gtname string
		expected string
	}

	t.Run("Validate", func(t *testing.T) {
		tests := []ValidateTestType {
			{"_group:task", positive},
			{"group:task", positive},
			{"task", positive},
			{"_:task", positive},
			{"g roup:task", negative},
			{"g~roup:task", negative},
			{"group:t ask", negative},
			{"group:t~ask", negative},
			{":task", negative},
		}

		for _, tt := range tests {
			t.Run(fmt.Sprintf("'%s' is %s", tt.gtname, tt.expected), func(t *testing.T) {
				err := GroupTask(tt.gtname).Validate()
				if tt.expected == positive {
					assert.NoError(t, err)
				} else {
					assert.Error(t, err)
				}
				
			})
		}
	})

	t.Run("ValidatePublic", func(t *testing.T) {
		tests := []ValidateTestType {
			{"_group:task", negative},
			{"group:task", positive},
			{"_:task", positive},
			{"task", positive},
		}

		for _, tt := range tests {
			t.Run(fmt.Sprintf("'%s' is %s", tt.gtname, tt.expected), func(t *testing.T) {
				err := GroupTask(tt.gtname).ValidatePublic()
				if tt.expected == positive {
					assert.NoError(t, err)
				} else {
					assert.Error(t, err)
				}
				
			})
		}
	})


	t.Run("Split", func(t *testing.T) {
		tests := []struct {
			gtname string
			g_expected string
			t_expected string
			expected string
		} {
			{"_group:task", "_group", "task", positive},
			{"group:task", "group", "task", positive},
			{"_:task", "_", "task", positive},
			{"task", "", "task", positive},
		}

		for _, tt := range tests {
			t.Run(fmt.Sprintf("'%s' is %v", tt.gtname, tt.expected), func(t *testing.T) {
				group, task, err := GroupTask(tt.gtname).Split()
				
				if tt.expected == positive {
					if assert.NoError(t, err) {
						assert.Equal(t, tt.g_expected, string(group))
						assert.Equal(t, tt.t_expected, string(task))
					}
				} else {
					assert.Error(t, err)
				}
			})
		}
	})

	t.Run("Create", func(t *testing.T) {
		tests := []struct {
			gname string
			tname string
			gt_expected string
		} {
			{"group", "task", "group:task"},
			{"", "task", "task"},
			{"_", "task", "_:task"},
		}

		for _, tt := range tests {
			t.Run(fmt.Sprintf("group '%s' and task '%s' are to '%s'", tt.gname, tt.tname, tt.gt_expected), func(t *testing.T) {
				gt := Create(group.Group(tt.gname), task.Task(tt.tname))
				
				assert.Equal(t, tt.gt_expected, string(gt))
			})
		}
	})
}
