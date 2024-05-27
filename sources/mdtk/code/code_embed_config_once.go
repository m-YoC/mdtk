package code

import (
	"strings"
	"mdtk/lib"
	"mdtk/grtask"
)

type TaskStackData struct {
	AlreadyRead map[grtask.GroupTask]bool
}

func CreateTaskStackData() TaskStackData {
	return TaskStackData{AlreadyRead: map[grtask.GroupTask]bool{} }
}

func (t TaskStackData) HasData(gtname grtask.GroupTask) bool {
	_, ok := t.AlreadyRead[gtname]
	return ok
}

func (t *TaskStackData) Set(gtname grtask.GroupTask) {
	t.AlreadyRead[gtname] = true
}

var TaskStack lib.Stack[TaskStackData]

var CurrentTaskStackData TaskStackData

func init() {
	CurrentTaskStackData = CreateTaskStackData()
}

func WithNewTaskStackData(code Code, impl_fn func() (Code, error)) (Code, error) {
	TaskStack.Push(CurrentTaskStackData)
	CurrentTaskStackData = CreateTaskStackData()

	res, err := impl_fn()

	CurrentTaskStackData, _ = TaskStack.Pop()

	if err != nil {
		return "", err
	}

	return res, nil
}

// -----------------------------------------------------------------------

func (code Code) CheckAndRemoveConfigOnce() (Code, bool) {
	configs := code.GetEmbedComment("config")

	res := string(code)
	once_flag := false
	for _, config := range configs {
		if config[1] == "once" {
			once_flag = true

			res = strings.Replace(res, config[0] + "\n", "", -1)
		}
	}
	
	return Code(res), once_flag
}

// return: code, isReplacedByComment
func (code Code) ApplyConfigOnce(gtname grtask.GroupTask) (Code, bool) {
	res, f := code.CheckAndRemoveConfigOnce()
	if f {
		if CurrentTaskStackData.HasData(gtname) {
			res = Code("# task: " + string(gtname) + " is already embedded.")
			return res, true
		} else {
			CurrentTaskStackData.Set(gtname)
			return res, false
		}
	}

	return res, false
}

