package parse

import (
	"fmt"
	"testing"
	"github.com/stretchr/testify/assert"
)


func Test_SplitArgs(t *testing.T) {
	tests := []struct {
		name string
		cmds []string
		expected_front []string
		expected_args []string
	} {
		{"has both", []string{"echo", "hello", "world", "--", "a=1", "b=2"}, 
		[]string{"echo", "hello", "world"}, []string{"a=1", "b=2"}},
		{"no args", []string{"echo", "hello", "world"}, 
		[]string{"echo", "hello", "world"}, []string{}},
		{"has '--' but does not have args", []string{"echo", "hello", "world", "--"}, 
		[]string{"echo", "hello", "world"}, []string{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			
			f, a := SplitArgs(tt.cmds)
			assert.Equal(t, tt.expected_front, f)
			assert.Equal(t, tt.expected_args, a)
		})
	}
	
}

func Test_GetOpType(t *testing.T) {
	tests := []struct {
		name string
		cmd string
		expected int
	} {
		{"is not op", "hello", notOp},
		{"is single option", "--flag", singleOp},
		{"is single option", "-f", singleOp},
		{"is multi options", "-fgh", multiOps},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s %s", tt.cmd, tt.name), func(t *testing.T) {
			optype := GetOpType(tt.cmd)
			assert.Equal(t, tt.expected, optype)
		})
	}
}

func Test_MatchOp(t *testing.T) {
	const (
		positive = iota
		negative
	)

	f := Flag{}
	f.Set("--flag1", "-f").SetHasValue("default1")
	f.Set("--flag2", "-g")
	f.Set("--flag3", "-h").SetHasValue("default3")

	tests := []struct {
		name string
		cmd string
		expected []int
		expected_e int 
	} {
		{"valid option '--xxx'", "--flag1", []int{0}, positive},
		{"valid option '-x'", "-g", []int{1}, positive},
		{"valid multi options '-xy'", "-fh", []int{0, 2}, positive},
		{"not option 'xxx'", "hello", []int{}, negative},
		{"invalid option", "-a", []int{}, negative},
		{"multi options includes invalid option", "-ga", []int{}, negative},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x, err := MatchOp(tt.cmd, f)
			if tt.expected_e == positive {
				if assert.NoError(t, err) {
					assert.Equal(t, tt.expected, x)
				}
			} else {
				assert.Error(t, err)
			}
		})
	}
	
}
