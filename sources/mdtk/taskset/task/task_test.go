package task

import (
	"fmt"
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_Task(t *testing.T) {
	
	const (
		positive = "positive"
		negative = "negative"
	)

	type ValidateTestType struct {
		tname string
		expected string
	}

	t.Run("Validate", func(t *testing.T) {
		tests := []ValidateTestType {
			{"task", positive},
			{"t ask", negative},
			{"t~ask", negative},
		}

		for _, tt := range tests {
			t.Run(fmt.Sprintf("'%s' is %s", tt.tname, tt.expected), func(t *testing.T) {
				err := Task(tt.tname).Validate()
				if tt.expected == positive {
					assert.NoError(t, err)
				} else {
					assert.Error(t, err)
				}
				
			})
		}
	})
}