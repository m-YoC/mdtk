package code

import (
	"mdtk/args"
	"testing"
	"github.com/stretchr/testify/assert"
)


func Test_CodeEmbedArgsShell(t *testing.T) {
	t.Run("ApplyArgsShell", func(t *testing.T) {
		t.Run("positive, with quotes", func(t *testing.T) {
			code := Code("#TestString\n") 
	
			res, err := code.ApplyArgsShell(args.Args{"k1=v1", "k2:v2", "k3=v'3"}, true)
			if assert.NoError(t, err) {
				assert.Equal(t, "k1='v1'; k2='v2'; k3='v\\'3'; \n#TestString\n", string(res))
			}
			
		})

		t.Run("positive, with quotes, with {$}", func(t *testing.T) {
			code := Code("#TestString\n") 
	
			res, err := code.ApplyArgsShell(args.Args{"k1=v1", "k2:{$}", "k3={$}"}, true)
			if assert.NoError(t, err) {
				assert.Equal(t, "k1='v1'; k2=$1; k3=$2; \n#TestString\n", string(res))
			}
			
		})

		t.Run("positive, no quotes", func(t *testing.T) {
			code := Code("#TestString\n") 
	
			res, err := code.ApplyArgsShell(args.Args{"k1=v1", "k2:v2", "k3=v'3"}, false)
			if assert.NoError(t, err) {
				assert.Equal(t, "k1=v1; k2=v2; k3=v'3; \n#TestString\n", string(res))
			}
			
		})

		t.Run("negative", func(t *testing.T) {
			code := Code("#TestString\n") 
	
			_, err := code.ApplyArgsShell(args.Args{"k1=v1", "k2; v2", "k3=v3"}, true)
			assert.Error(t, err)
		})
	})

}
