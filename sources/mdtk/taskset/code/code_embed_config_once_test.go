package code

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_TaskStackData(t *testing.T) {
	t.Run("Set", func(t *testing.T) {
		CurrentTaskStackData = CreateTaskStackData()
		CurrentTaskStackData.Set("group:task")

		_, ok := CurrentTaskStackData.AlreadyRead["group:task"]
		assert.True(t, ok)

	})

	t.Run("HasData", func(t *testing.T) {
		CurrentTaskStackData = CreateTaskStackData()
		CurrentTaskStackData.Set("group:task")

		assert.True(t, CurrentTaskStackData.HasData("group:task"))
		assert.False(t, CurrentTaskStackData.HasData("group2:task"))
	})

	t.Run("WithNewTaskStackData", func(t *testing.T) {
		CurrentTaskStackData = CreateTaskStackData()
		CurrentTaskStackData.Set("group:task")
		
		assert.True(t, CurrentTaskStackData.HasData("group:task"))

		WithNewTaskStackData("", func() (Code, error) { 
			assert.False(t, CurrentTaskStackData.HasData("group:task"))
			return "", nil 
		})

		assert.True(t, CurrentTaskStackData.HasData("group:task"))
	})
}

func Test_CodeEmbedConfigOnce(t *testing.T) {
	t.Run("CheckAndRemoveConfigOnce", func(t *testing.T) {
		t.Run("has config once", func(t *testing.T) {
			code := Code("#TestString\n#config> once  \n")
	
			res, b := code.CheckAndRemoveConfigOnce()
			assert.Equal(t, "#TestString\n", string(res))
			assert.True(t, b)
		})

		t.Run("does not have config once", func(t *testing.T) {
			code := Code("#TestString\n#config> not once  \n")
	
			res, b := code.CheckAndRemoveConfigOnce()
			assert.Equal(t, "#TestString\n#config> not once  \n", string(res))
			assert.False(t, b)
		})
	})

	t.Run("ApplyConfigOnce", func(t *testing.T) {
		t.Run("has config once and already used", func(t *testing.T) {
			CurrentTaskStackData = CreateTaskStackData()
			CurrentTaskStackData.Set("group:task")

			code := Code("#TestString\n#config> once  \n")
	
			res, b := code.ApplyConfigOnce("group:task")
			assert.Equal(t, "# task: group:task is already embedded.", string(res))
			assert.True(t, b)
		})

		t.Run("has config once and does not use yet", func(t *testing.T) {
			CurrentTaskStackData = CreateTaskStackData()

			code := Code("#TestString\n#config> once  \n")
	
			res, b := code.ApplyConfigOnce("group:task")
			assert.Equal(t, "#TestString\n", string(res))
			assert.False(t, b)
		})

		t.Run("does not have config once", func(t *testing.T) {
			CurrentTaskStackData = CreateTaskStackData()

			code := Code("#TestString\n#config> not once  \n")
	
			res, b := code.ApplyConfigOnce("group:task")
			assert.Equal(t, string(code), string(res))
			assert.False(t, b)
		})
	})
}

