package code

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_CodeEmbedCodes(t *testing.T) {
	t.Run("ApplyEmbedCodes", func(t *testing.T) {
		t.Run("No extra space", func(t *testing.T) {
			code := Code("#TestString\n#embed> aaaaa\n#embed> mdtk\n") 
	
			res, _ := code.ApplyEmbedCodes(TestTaskDataSet1{}, 1)
			assert.Equal(t, "#TestString\n" + TestCode1 + "\n" + TestCode2 + "\n", string(res))
		})

		t.Run("Has extra space", func(t *testing.T) {
			code := Code("#TestString\n#embed>   	aaaaa   \n#embed>   mdtk   	\n") 
	
			res, _ := code.ApplyEmbedCodes(TestTaskDataSet1{}, 1)
			assert.Equal(t, "#TestString\n" + TestCode1 + "\n" + TestCode2 + "\n", string(res))
		})
	})
}
