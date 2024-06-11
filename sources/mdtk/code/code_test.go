package code

import (
	"fmt"
	"testing"
	"mdtk/grtask"
	"mdtk/args"
	"github.com/stretchr/testify/assert"
)

const TestCode1 = `#TestString1
#test> Hello Golang Test!!
echo Happy Test
#test> HogeFuga  
echo hahahahaha
`

const TestCode2 = `#TestString2
echo markdown taskrunner
`

const TestCode3 = `#TestString3
RUN #embed> Hello Golang Test!!
//test> HogeFuga  
`

type TestTaskDataSet1 struct {}

func (ttds TestTaskDataSet1) GetTask(a grtask.GroupTask, b args.Args, c bool, d bool, e int) (Code, error) {
	if string(a) == "mdtk" {
		return Code(TestCode2), nil
	}
	return Code(TestCode1), nil	
}

type TestTaskDataSet2 struct {}

func (ttds TestTaskDataSet2) GetTask(a grtask.GroupTask, b args.Args, c bool, d bool, e int) (Code, error) {
	return Code(""), fmt.Errorf("")
}


func Test_Code(t *testing.T) {
	t.Run("GetEmbedComment", func(t *testing.T) {
		code := Code(TestCode1) 

		res := code.GetEmbedComment("test")
		assert.Equal(t, 2, len(res))
		assert.Equal(t, "Hello Golang Test!!", res[0][1])
		assert.Equal(t, "HogeFuga", res[1][1])
	})

	t.Run("GetEmbedCommentText", func(t *testing.T) {
		code := Code(TestCode1) 

		res := code.GetEmbedCommentText("test")
		assert.Equal(t, 2, len(res))
		assert.Equal(t, "Hello Golang Test!!", res[0])
		assert.Equal(t, "HogeFuga", res[1])
	})

	t.Run("RemoveEmbedComment", func(t *testing.T) {
		code := Code(TestCode1) 

		res := code.RemoveEmbedComment("test")
		assert.Equal(t, "#TestString1\necho Happy Test\necho hahahahaha\n", string(res))
	})
}




