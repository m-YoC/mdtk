package main

import (
	"strings"
	"mdtk/lib"
	"testing"
	// "github.com/stretchr/testify/assert"
)

func Test_Command(t *testing.T) {
	file := " -f ../SampleTaskfiles/first-sample.taskrun.md "

	type A struct {
		cmd string
		args string
	}
	
	// first-sample.taskrun.md
	tests := lib.TestCases[A, string] {
		{Name: "run simple task", TestArg: A{cmd: "mdtk hello"}, 
		Expected: "hello world"},
		{Name: "with group name", TestArg: A{cmd: "mdtk group-test group-test"}, 
		Expected: "group test"},
		{Name: "task conflict", TestArg: A{cmd: "mdtk group-test conflict-test"}, 
		Expected: "Task of max priority cannot be identified."},
		{Name: "exists description test", TestArg: A{cmd: "mdtk desc desc-test"}, 
		Expected: "desc test"},
		{Name: "description is written into taskhelp", TestArg: A{cmd: "mdtk desc help"}, 
		Expected: "description test"},
		{Name: "cannot run hidden attribite test without -a", TestArg: A{cmd: "mdtk attr hidden-test"}, 
		Expected: "Private/Hidden group cannot be executed directly."},
		{Name: "exists hidden attribite test", TestArg: A{cmd: "mdtk attr hidden-test -a"}, 
		Expected: "hidden test"},
		{Name: "exists bottom attribite test", TestArg: A{cmd: "mdtk attr bottom-test"}, 
		Expected: "bottom layout test"},
		{Name: "exists top attribite test", TestArg: A{cmd: "mdtk attr top-test"}, 
		Expected: "top layout test"},
		{Name: "select task that has priority attribite +5", TestArg: A{cmd: "mdtk attr priority-test"}, 
		Expected: "priority test (+5)"},
		{Name: "select task that does not have weak attr", TestArg: A{cmd: "mdtk attr weak-test"}, 
		Expected: "weak test (not weak)"},
	}

	tests.Run(t, func(t *testing.T, i int) {
		tt := tests.Get(i)
		lib.AssertStringContains(t, tt.Expected, lib.RemoveANSIColor(lib.CmdTest(tt.TestArg.cmd + file)))
	})

	t.Run("Test2: position of 'top attr' task & 'bottom attr' one", func(t *testing.T) {
		res := lib.RemoveANSIColor(lib.CmdTest("mdtk attr help" + file))
		sres := strings.Split(res, "\n")
		ssres := sres[2:len(sres)-2]
		lib.AssertStringContains(t, "top-test", ssres[0])
		lib.AssertStringContains(t, "bottom-test", ssres[len(ssres)-1])
	})


	// embed-sample.task.md
	tests2 := lib.TestCases[A, string] {
		{Name: "test #embed>", TestArg: A{cmd: "mdtk eb embed-test"}, 
		Expected: "hello"},
		{Name: "test #task>", TestArg: A{cmd: "mdtk eb task-test"}, 
		Expected: "hello"},
		{Name: "test #task> with args", TestArg: A{cmd: "mdtk eb taskarg-test"}, 
		Expected: "hello / arg: task-arg"},
		{Name: "test #func>", TestArg: A{cmd: "mdtk eb func-test"}, 
		Expected: "hello"},
		{Name: "test #func> with args", TestArg: A{cmd: "mdtk eb funcarg-test"}, 
		Expected: "hello / arg: func-arg"},
		{Name: "test #func> with args2: required type positional parameter is not set", TestArg: A{cmd: "mdtk eb funcarg2-test"}, 
		Expected: "$1: unbound variable"},
		{Name: "test #func> with args3: optional type positional parameter is not set", TestArg: A{cmd: "mdtk eb funcarg3-test"}, 
		Expected: "hello / arg: default"},
		{Name: "test #config> once (embed)", TestArg: A{cmd: "mdtk eb once1-test"}, 
		Expected: "* first\nhello\n* second\nend\n"},
		{Name: "test #config> once (task)", TestArg: A{cmd: "mdtk eb once2-test"}, 
		Expected: "* first\nhello\n* second\nhello\nend\n"},
		{Name: "test #config> once (func)", TestArg: A{cmd: "mdtk eb once3-test"}, 
		Expected: "* first\nhello\n* second\nhello\nend\n"},
		{Name: "test #desc>", TestArg: A{cmd: "mdtk eb help"}, 
		Expected: "This is a description of '#desc>'."},
		{Name: "test #args>", TestArg: A{cmd: "mdtk eb help"}, 
		Expected: "Text of '#args>' is simply a type of task help description."},
	}

	tests2.Run(t, func(t *testing.T, i int) {
		tt := tests2.Get(i)
		lib.AssertStringContains(t, tt.Expected, lib.RemoveANSIColor(lib.CmdTest(tt.TestArg.cmd + file)))
	})


	// embed-replace-sample.task.md
	tests3 := lib.TestCases[A, string] {
		{Name: "replacable", TestArg: A{cmd: "mdtk ebrp embed-replace-test"}, 
		Expected: "Already replaced."},
	}

	tests3.Run(t, func(t *testing.T, i int) {
		tt := tests3.Get(i)
		lib.AssertStringContains(t, tt.Expected, lib.RemoveANSIColor(lib.CmdTest(tt.TestArg.cmd + file)))
	})
	
}


func Test_CommandPwSh(t *testing.T) {
	file := " -f ../SampleTaskfiles/pwsh-sample.task.md "

	type A struct {
		cmd string
		args string
	}

	// embed-sample.task.md
	tests := lib.TestCases[A, string] {
		{Name: "test #embed>", TestArg: A{cmd: "mdtk pseb embed-test"}, 
		Expected: "hello"},
		{Name: "test #func>", TestArg: A{cmd: "mdtk pseb func-test"}, 
		Expected: "hello"},
		{Name: "test #func> with args", TestArg: A{cmd: "mdtk pseb funcarg-test"}, 
		Expected: "hello / arg: func-arg"},
		{Name: "test #func> with args2: required type positional parameter is not set", TestArg: A{cmd: "mdtk pseb funcarg2-test"}, 
		Expected: "OperationStopped:"},
		{Name: "test #func> with args3: optional type positional parameter is not set", TestArg: A{cmd: "mdtk pseb funcarg3-test"}, 
		Expected: "hello / arg: default\n"},
		{Name: "test #config> once (embed)", TestArg: A{cmd: "mdtk pseb once1-test"}, 
		Expected: "* first\nhello\n* second\nend\n"},
		{Name: "test #config> once (func)", TestArg: A{cmd: "mdtk pseb once2-test"}, 
		Expected: "* first\nhello\n* second\nhello\nend\n"},
		{Name: "test #desc>", TestArg: A{cmd: "mdtk pseb help"}, 
		Expected: "This is a description of '#desc>'."},
		{Name: "test #args>", TestArg: A{cmd: "mdtk pseb help"}, 
		Expected: "Text of '#args>' is simply a type of task help description."},
	}

	tests.Run(t, func(t *testing.T, i int) {
		tt := tests.Get(i)
		lib.AssertStringContains(t, tt.Expected, lib.RemoveANSIColor(lib.CmdTest(tt.TestArg.cmd + file)))
	})
}
