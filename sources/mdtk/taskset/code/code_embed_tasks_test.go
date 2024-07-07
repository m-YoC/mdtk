package code

import (
	"regexp"
	"mdtk/lib"
	"mdtk/args"
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_taskCmdsConstraint(t *testing.T) {
	type E struct {
		PN string
		Gtname string
		Args args.Args
	}

	tests := lib.TestCases[string, E] {
		{Name: "no args", TestArg: "group:task", 
		Expected: E{ lib.Positive, "group:task", args.Args{}}},
		{Name: "has args", TestArg: "group:task -- k1=v1  k2=v2", 
		Expected: E{ lib.Positive, "group:task", args.ToArgs("k1=v1", "k2=v2")}},
		{Name: "head has more than 1", TestArg: "group task -- k1=v1  k2=v2", Expected: E{ PN: lib.Negative }},
		{Name: "no head", TestArg: "-- k1=v1  k2=v2", Expected: E{ PN: lib.Negative }},
		{Name: "has bad args (1)", TestArg: "group:task -- k1=v1  k2={$}", Expected: E{ PN: lib.Negative }},
		{Name: "has bad args (2)", TestArg: "group:task -- k1=v1  k2=<$>", Expected: E{ PN: lib.Negative }},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			gt, a, err := taskCmdsConstraint(extractSubCmds(tt.TestArg))
			if tt.Expected.PN == lib.Positive {
				if assert.NoError(t, err) {
					assert.Equal(t, tt.Expected.Gtname, string(gt))
					assert.Equal(t, tt.Expected.Args, a)
				}
			} else {
				assert.Error(t, err)
			}
		})
	}
}

func Test_CodeEmbedTasks(t *testing.T) {
	t.Run("ApplySubTasks", func(t *testing.T) {
		t.Run("No extra space", func(t *testing.T) {
			code := Code("#TestString\n#task> aaaaa\n#task> mdtk\n") 
	
			res, _ := code.ApplySubTasks(TestTaskDataSet1{}, 1)
			rex := regexp.MustCompile("(?s)\\(.*\n.*#TestString1.*\n\\)\n\\(.*\n.*#TestString2.*\n\\)")
			assert.Regexp(t, rex, string(res))
		})

		t.Run("Has extra space", func(t *testing.T) {
			code := Code("#TestString\n#task>   	aaaaa   \n#task>   mdtk   	\n") 
	
			res, _ := code.ApplySubTasks(TestTaskDataSet1{}, 1)
			rex := regexp.MustCompile("(?s)\\(.*\n.*#TestString1.*\n\\)\n\\(.*\n.*#TestString2.*\n\\)")
			assert.Regexp(t, rex, string(res))
		})
	})
}
