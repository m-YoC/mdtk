package read

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

const md_sample = `
~~~task::task
echo test
~~~

~~~task:group:task
echo test2
~~~

~~~taskfile
go/test/file.md
  /markdown/read/data.md  
~~~
`

func Test_Markdown(t *testing.T) {
	t.Run("SimplifyNewline", func(t *testing.T) {
		const text1 = "Hello\nGolang\nTest\n"
		const text2 = "Hello\r\nGolang\r\nTest\r\n"

		md1 := Markdown(text1)
		md2 := Markdown(text2)

		assert.Equal(t, text1, string(md1.SimplifyNewline()))
		assert.Equal(t, text1, string(md2.SimplifyNewline()))
	})

	t.Run("ExtractCode", func(t *testing.T) {
		t.Run("positive", func(t *testing.T) {
			const begin = 14
			const block = "~~~"

			code, err := Markdown(md_sample).ExtractCode(begin, block)

			if assert.NoError(t, err) {
				assert.Equal(t, "echo test", string(code))
			}
		})

		t.Run("negative (cannnot find the end '~~~~' of code block)", func(t *testing.T) {
			const begin = 14
			const block = "~~~~"

			_, err := Markdown(md_sample).ExtractCode(begin, block)

			assert.Error(t, err)
		})
	})

	t.Run("GetTaskBlock", func(t *testing.T) {
		t.Run("positive", func(t *testing.T) {
			tdarr, err := Markdown(md_sample).GetTaskBlock()

			if assert.NoError(t, err) {
				for i, s := range [][]string{{"_", "task", "echo test"}, {"group", "task", "echo test2"}} {
					assert.Equal(t, s[0], string(tdarr[i].Group))
					assert.Equal(t, s[1], string(tdarr[i].Task))
					assert.Equal(t, s[2], string(tdarr[i].Code))
				}
				
			}
		})

		t.Run("negative (error handling of Markdown.ExtractCode)", func(t *testing.T) {})
	})
}
