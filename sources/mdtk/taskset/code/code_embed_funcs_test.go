package code

import (
	"regexp"
	"mdtk/lib"
	"mdtk/args"
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_funcCmdsConstraint(t *testing.T) {
	type E struct {
		PN string
		FName string
		Gtname string
		Args args.Args
	}

	tests := lib.TestCases[string, E] {
		{Name: "no args", TestArg: "f group:task", 
		Expected: E{ lib.Positive, "f", "group:task", args.Args{}}},
		{Name: "has args", TestArg: "f group:task -- k1=v1  k2=v2", 
		Expected: E{ lib.Positive, "f", "group:task", args.ToArgs("k1=v1", "k2=v2")}},
		{Name: "head has less than 2", TestArg: "f -- k1=v1  k2=v2", Expected: E{ PN: lib.Negative }},
		{Name: "head has more than 2", TestArg: "f group task -- k1=v1  k2=v2", Expected: E{ PN: lib.Negative }},
		{Name: "no head", TestArg: "-- k1=v1  k2=v2", Expected: E{ PN: lib.Negative }},
		{Name: "has sp args (1)", TestArg: "f group:task -- k1=v1  k2={$}", 
		Expected: E{ lib.Positive, "f", "group:task", args.ToArgs("k1=v1", "k2={$}")}},
		{Name: "has sp args (2)", TestArg: "f group:task -- k1=v1  k2=<$>", 
		Expected: E{ lib.Positive, "f", "group:task", args.ToArgs("k1=v1", "k2=<$>")}},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			b, gt, a, err := funcCmdsConstraint(extractSubCmds(tt.TestArg))
			if tt.Expected.PN == lib.Positive {
				if assert.NoError(t, err) {
					assert.Equal(t, tt.Expected.FName, b)
					assert.Equal(t, tt.Expected.Gtname, string(gt))
					assert.Equal(t, tt.Expected.Args, a)
				}
			} else {
				assert.Error(t, err)
			}
		})
	}
}

func Test_CodeEmbedFuncs(t *testing.T) {
	t.Run("ApplyFuncs", func(t *testing.T) {
		t.Run("No extra space", func(t *testing.T) {
			code := Code("#TestString\n#func> f aaaaa\n#func> g mdtk\n") 
	
			res, _ := code.ApplyFuncs(TestTaskDataSet1{}, ParenTheses, ":", 1)
			rex := regexp.MustCompile("(?s)function f\\(.*\n.*#TestString1.*\n\\)\nfunction g\\(.*\n.*#TestString2.*\n\\)")
			assert.Regexp(t, rex, string(res))
		})

		t.Run("Has extra space", func(t *testing.T) {
			code := Code("#TestString\n#func>   f    	aaaaa   \n#func> g  mdtk   	\n") 
	
			res, _ := code.ApplyFuncs(TestTaskDataSet1{}, ParenTheses, ":", 1)
			rex := regexp.MustCompile("(?s)function f\\(.*\n.*#TestString1.*\n\\)\nfunction g\\(.*\n.*#TestString2.*\n\\)")
			assert.Regexp(t, rex, string(res))
		})
	})
}
