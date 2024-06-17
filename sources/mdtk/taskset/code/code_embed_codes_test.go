package code

import (
	"mdtk/lib"
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_embedCmdsConstraint(t *testing.T) {
	type E struct {
		PN string
		Gtname string
	}

	tests := lib.TestCases[string, E] {
		{Name: "no args", TestArg: "group:task", Expected: E{ lib.Positive, "group:task"}},
		{Name: "has args", TestArg: "group:task -- k1=v1  k2=v2", Expected: E{ PN: lib.Negative}},
		{Name: "head has more than 1", TestArg: "group task -- k1=v1  k2=v2", Expected: E{ PN: lib.Negative }},
		{Name: "no head", TestArg: "-- k1=v1  k2=v2", Expected: E{ PN: lib.Negative }},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			gt, err := embedCmdsConstraint(extractSubCmds(tt.TestArg))
			if tt.Expected.PN == lib.Positive {
				if assert.NoError(t, err) {
					assert.Equal(t, tt.Expected.Gtname, string(gt))
				}
			} else {
				assert.Error(t, err)
			}
		})
	}
}

func Test_CodeEmbedCodes(t *testing.T) {
	t.Run("ApplyEmbedCodes", func(t *testing.T) {
		t.Run("No extra space", func(t *testing.T) {
			code := Code("#TestString\n#embed> aaaaa\n#embed> mdtk\n") 
	
			res, _ := code.ApplyEmbedCodes(TestTaskDataSet1{}, 1)
			assert.Equal(t, "#TestString\n" + TestCode1 + "\n" + TestCode2 + "\n", string(res))
		})

		t.Run("Has extra space", func(t *testing.T) {
			code := Code("#TestString\n  #embed>   	aaaaa   \n#embed>   mdtk   	\n") 
	
			res, _ := code.ApplyEmbedCodes(TestTaskDataSet1{}, 1)
			assert.Equal(t, "#TestString\n" + TestCode1 + "\n" + TestCode2 + "\n", string(res))
		})
	})
}
