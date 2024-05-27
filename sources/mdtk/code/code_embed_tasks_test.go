package code

import (
	"mdtk/args"
	"testing"
	"regexp"
	"github.com/stretchr/testify/assert"
)

func Test_getEmbedCommentTaskAndArgs(t *testing.T) {
	const (
		positive = iota
		negative
	)

	tests := []struct {
		expected_e int
		name string
		comment string
		expected_b bool
		expected_gt string
		expected_a args.Args
	} {
		{positive, "no args", "group:task", false, "group:task", args.Args{}},
		{positive, "has args", "group:task -- k1=v1  k2=v2", false, "group:task", args.ToArgs("k1=v1", "k2=v2")},
		{positive, "no args (with @)", "@ group:task", true, "group:task", args.Args{}},
		{positive, "has args (with @)", "@ group:task -- k1=v1  k2=v2", true, "group:task", args.ToArgs("k1=v1", "k2=v2")},
		{negative, "head has more than 1", "group task -- k1=v1  k2=v2", false, "", args.Args{}},
		{negative, "head has more than 2 (with @)", "@ group task -- k1=v1  k2=v2", true, "", args.Args{}},
		{negative, "no head", "-- k1=v1  k2=v2", false, "", args.Args{}},
		{negative, "head has only @", "@ -- k1=v1  k2=v2", false, "", args.Args{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, gt, a, err := getEmbedCommentTaskAndArgs(tt.comment)
			if tt.expected_e == positive {
				if assert.NoError(t, err) {
					assert.Equal(t, tt.expected_b, b)
					assert.Equal(t, tt.expected_gt, string(gt))
					assert.Equal(t, tt.expected_a, a)
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
			rex := regexp.MustCompile("(?s)\\(\n.*#TestString1.*\n\\).*\\(\n.*#TestString2.*\n\\)")
			assert.Regexp(t, rex, string(res))
		})

		t.Run("Has extra space", func(t *testing.T) {
			code := Code("#TestString\n#task>   	aaaaa   \n#task>   mdtk   	\n") 
	
			res, _ := code.ApplySubTasks(TestTaskDataSet1{}, 1)
			rex := regexp.MustCompile("(?s)\\(\n.*#TestString1.*\n\\).*\\(\n.*#TestString2.*\n\\)")
			assert.Regexp(t, rex, string(res))
		})
	})
}
