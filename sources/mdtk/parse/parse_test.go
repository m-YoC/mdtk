package parse

import (
	"fmt"
	"mdtk/base"
	"mdtk/lib"
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

func Test_Parse(t *testing.T) {
	
	GetFlag := func() Flag {
		var flags Flag
		flags.Set("--flag1", "-x").SetHasValue("314")
		flags.Set("--flag2", "-f")
		flags.Set("--flag3", "-g").SetHasValue("314")
		flags.Set("--flag4", "-h")
		return flags
	}

	type FE =  map[string]bool
	type ExpectedT struct {
		gtname string
		flags_exist FE
		args_size int
	}

	tests := lib.TestCases[[]string, ExpectedT] {
		{Name: "Array has nothing (except arr[0])", TestArg: []string{}, 
		Expected: ExpectedT{"default", FE{"--flag1": false, "--flag2": false, "--flag3": false, "--flag4": false}, 0}},
		{Name: "Only gtname", TestArg: []string{"group:hello"}, 
		Expected: ExpectedT{"group:hello", FE{"--flag1": false, "--flag2": false, "--flag3": false, "--flag4": false}, 0}},
		{Name: "Group and task are split", TestArg: []string{"group", "hello"}, 
		Expected: ExpectedT{"group:hello", FE{"--flag1": false, "--flag2": false, "--flag3": false, "--flag4": false}, 0}},
		{Name: "Has one option", TestArg: []string{"group", "hello", "-f"}, 
		Expected: ExpectedT{"group:hello", FE{"--flag1": false, "--flag2": true, "--flag3": false, "--flag4": false}, 0}},
		{Name: "Has two options", TestArg: []string{"group", "hello", "-f", "-h"}, 
		Expected: ExpectedT{"group:hello", FE{"--flag1": false, "--flag2": true, "--flag3": false, "--flag4": true}, 0}},
		{Name: "Has two options (multi)", TestArg: []string{"group", "hello", "-fh"}, 
		Expected: ExpectedT{"group:hello", FE{"--flag1": false, "--flag2": true, "--flag3": false, "--flag4": true}, 0}},
		{Name: "Has one option and value", TestArg: []string{"group", "hello", "-g", "777"}, 
		Expected: ExpectedT{"group:hello", FE{"--flag1": false, "--flag2": false, "--flag3": true, "--flag4": false}, 0}},
		{Name: "Has Args", TestArg: []string{"group:hello", "--", "arg1=value1", "arg2=value2"}, 
		Expected: ExpectedT{"group:hello", FE{"--flag1": false, "--flag2": false, "--flag3": false, "--flag4": false}, 2}},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			gtname, flags, args := Parse(append([]string{"mdtk"}, tt.TestArg...), GetFlag())
	
			assert.Equal(t, tt.Expected.gtname, string(gtname))
			for k, v := range tt.Expected.flags_exist {
				assert.Equal(t, v, flags.GetData(k).Exist)
			}
			assert.Equal(t, tt.Expected.args_size, len(args))
		})
	}

	t.Run("Negative Cases", func(t *testing.T){
		tests := lib.TestCases[[]string, bool] {
			{Name: "Too many words", TestArg: []string{"group", "hello", "world"}},
			{Name: "Invalid option", TestArg: []string{"group", "hello", "-p"}},
			{Name: "Multi type option includes over 2 options that need value", TestArg: []string{"group", "hello", "-xg", "777", "777"}},
			{Name: "Option do not have value", TestArg: []string{"group", "hello", "-g"}},
			{Name: "Option do not have value (2)", TestArg: []string{"group", "hello", "-g", "-h"}},
		}

		for _, tt := range tests {
			t.Run(tt.Name, func(t *testing.T) {
				test_status := -1
				defer base.NewExit(func(status int) { 
					test_status = status
					panic("")
				})()
				assert.Panics(t, func(){ Parse(append([]string{"mdtk"}, tt.TestArg...), GetFlag()) })
				
				assert.Equal(t, 1, test_status)
			})
		}
	})
	
	
	
}
