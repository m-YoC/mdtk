package code

import (
	"mdtk/args"
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_escapeQuoteAndEnclose(t *testing.T) {
	tests := []struct {
		name string
		quote string
		str string
		expected string
	} {
		{"no quote", "", `Hello Go's Test`, `Hello Go's Test`},
		{"single quote, str has no quote", `'`, `Hello Test`, `'Hello Test'`},
		{"single quote, str has single quote", `'`, `Hello Go's Test`, `'Hello Go\'s Test'`},
		{"single quote, str has double quote", `'`, `Hello Go"s Test`, `'Hello Go"s Test'`},
		{"double quote, str has no quote", `"`, `Hello Test`, `"Hello Test"`},
		{"double quote, str has single quote", `"`, `Hello Go's Test`, `"Hello Go's Test"`},
		{"double quote, str has double quote", `"`, `Hello Go"s Test`, `"Hello Go\"s Test"`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := escapeQuoteAndEnclose(tt.str, tt.quote)
			assert.Equal(t, tt.expected, s)
		})
	}

}

func Test_CodeEmbedArgs(t *testing.T) {
	t.Run("ApplyArgs", func(t *testing.T) {
		t.Run("positive, with quotes", func(t *testing.T) {
			code := Code("#TestString\n") 
	
			res, err := code.ApplyArgs(args.Args{"k1=v1", "k2:v2", "k3=v'3"}, true)
			if assert.NoError(t, err) {
				assert.Equal(t, "k1='v1'; k2='v2'; k3='v\\'3'; \n#TestString\n", string(res))
			}
			
		})

		t.Run("positive, no quotes", func(t *testing.T) {
			code := Code("#TestString\n") 
	
			res, err := code.ApplyArgs(args.Args{"k1=v1", "k2:v2", "k3=v'3"}, false)
			if assert.NoError(t, err) {
				assert.Equal(t, "k1=v1; k2=v2; k3=v'3; \n#TestString\n", string(res))
			}
			
		})

		t.Run("negative", func(t *testing.T) {
			code := Code("#TestString\n") 
	
			_, err := code.ApplyArgs(args.Args{"k1=v1", "k2; v2", "k3=v3"}, true)
			assert.Error(t, err)
		})
	})

	t.Run("RemoveEmbedArgsComment", func(t *testing.T) {
		t.Run("positive1", func(t *testing.T) {
			code := Code("#TestString\n#args> hoge hoge hoge hoge\n") 
	
			res := code.RemoveEmbedArgsComment()
			assert.Equal(t, "#TestString\n", string(res))
			
		})

		t.Run("positive2", func(t *testing.T) {
			code := Code("#TestString\n") 
	
			res := code.RemoveEmbedArgsComment()
			assert.Equal(t, "#TestString\n", string(res))
			
		})
	})
}
