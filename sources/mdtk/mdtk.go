package main

import (
	"fmt"
	"mdtk/base"
	"mdtk/exec"
	"mdtk/parse"
	"mdtk/help"
	"mdtk/taskset"
	"mdtk/taskset/path"
	"mdtk/taskset/read"
	"mdtk/taskset/grtask"
	"mdtk/args"
	"mdtk/taskset/cache"
	"mdtk/config"
	"mdtk/mdtk_sub"
	"os"
	"strconv"
	_ "embed"
)

//go:embed version.txt
var version string

func getVersion() string {
	_, v, _ := args.Arg(version).GetData()
	return v
}

func GetFlag () parse.Flag {
	nestsizestr := strconv.FormatUint(uint64(config.Config.NestMaxDepth), 10)

	flags := parse.Flag{}
	flags.Set("--file", "-f").SetHasValue("")
	flags.Back().SetDescription("Select a task file. Task is run in working directory.")
	flags.Set("--File", "-F").SetHasValue("")
	flags.Back().SetDescription("Select a task file. Task is run in taskfile directory.")

	flags.Set("--nest", "-n").SetHasValue(nestsizestr)
	flags.Back().SetDescription("Set the nest maximum depth of embedded comment (embed/task).\nDefault is " + nestsizestr + ".")
	
	flags.Set("--quiet", "-q")
	flags.Back().SetDescription("Task output is not sent to standard output.")
	
	flags.Set("--all-task", "-a")
	flags.Back().SetDescription("Can select private groups and hidden tasks at the command.\nOr show all tasks that include private groups and hidden tasks at task help.")
	
	flags.Set("--script", "-s").SetDescription("Display script.")
	flags.Set("--no-head-script", "-S").SetDescription("Display script (No shebang, etc.).")

	flags.Set("--path").SetDescription("Get file path of selected task.")
	flags.Set("--dir").SetDescription("Get directory of taskfile in which selected task is written.")
	
	// -------------------------------------------------------------------------

	flags.Set("--make-cache", "-c").SetDescription("Make taskdata cache.")
	flags.Set("--lib", "-l").SetHasValue("")
	flags.Back().SetDescription("Select a library file.\nThis is a special version of --file option.\nNo need to add an extension '.mdtklib'.")
	flags.Set("--make-library").SetHasValue("")
	flags.Back().SetDescription("Make taskdata library.\nValue is library name.")

	// -------------------------------------------------------------------------

	flags.Set("--version", "-v").SetDescription("Show version.")
	flags.Set("--groups", "-g").SetDescription("Show groups.")
	flags.Set("--help", "-h").SetDescription("Show command help.")
	flags.Set("--manual", "-m").SetDescription("Show mdtk manual.")
	flags.Set("--write-configbase").SetDescription("Write config base file to current directory.")
	return flags
}

func CheckConflict(flags parse.Flag) {
	var fclist parse.FlagConflictList
	fclist.Conflict("--file", "--File", "--lib")
	fclist.Conflict("--script", "--no-head-script", "--path", "--dir", "--make-library", "--version", "--groups", "--help", "--manual", "--write-configbase")
	
	if err := fclist.Check(flags); err != nil {
		fmt.Print(err)
		base.MdtkExit(1)
	}
}


func main() {
	gtname_str, flags, task_args_strarr := parse.Parse(os.Args, GetFlag())
	CheckConflict(flags)

	var oflags sub.OtherFlags
	flags, oflags = sub.FixFlags(flags)
	
	sub.ReadConfig(flags.GetData("--file"))

	gtname := grtask.GroupTask(gtname_str)
	task_args := args.ToArgs(task_args_strarr...)
	nestsize := sub.GetNestSize(flags.GetData("--nest"))

	args := ArgsGroupA{flags: flags, oflags: oflags, gtname: gtname, args: task_args, nestsize: nestsize}

	RunGroupA(args)
}

/**
<If: need TaskDataSet>
|y        \n
|          ---- Group A
| 
<If: need Selected TaskData>
|y        \n
|          ---- Group B
|
<If: need Selected Code>
|y        \n
|          ---- Group C
|
Group D
--------------------------
Priority: A > B > C > D
Priorities within each group are in mdtk_sub/action_orders.go
*/

// ----------------------------------------------------------------------------

type ArgsGroupA struct {
	flags parse.Flag
	oflags sub.OtherFlags
	gtname grtask.GroupTask
	args args.Args
	nestsize uint
}

func RunGroupA(a ArgsGroupA) {
	FlagHas := func(str string) bool { return a.flags.GetData(str).Exist }

	switch sub.EnumGroupA(FlagHas("-v"), FlagHas("-h"), FlagHas("-m"), FlagHas("--write-configbase")) {
	case sub.ACT_VERSION:
		fmt.Println("mdtk version", getVersion())
	case sub.ACT_CMD_HELP:
		help.ShowCommandHelp(a.flags, 26)
	case sub.ACT_MANUAL:
		help.ShowManual()
	case sub.ACT_WRITE_CONFIG:
		config.WriteDefaultConfig()
	default:
		RunGroupB(a)
	}

	base.MdtkExit(0)
}

// ----------------------------------------------------------------------------

type ArgsGroupB struct {
	filename path.Path
	tds taskset.TaskDataSet
}

func RunGroupB(a ArgsGroupA) {
	FlagHas := func(str string) bool { return a.flags.GetData(str).Exist }

	// get Taskfile
	filename := path.Path("")
	if fd := a.flags.GetData("--file"); fd.Exist {
		filename = path.Path(fd.Value)
	} else {
		filename = read.SearchTaskfile()
	}

	// read Taskfile
	tds := sub.ReadTaskDataSet(filename, FlagHas("--make-cache"))

	// make lib
	fdml := a.flags.GetData("--make-library")
	switch sub.EnumGroupB(FlagHas("--groups"), help.ShouldShowHelp(a.gtname, tds), fdml.Exist) {
	case sub.ACT_TASK_HELP:
		help.ShowHelp(filename.String(), a.gtname, tds, FlagHas("--all-task"))
	case sub.ACT_GROUPS:
		help.ShowGroups(filename.String(), tds, FlagHas("--all-task"))
	case sub.ACT_MAKE_LIB:
		cache.WriteLib(tds, filename.Dir(), fdml.Value, int(a.nestsize))
	default:
		RunGroupC(a, ArgsGroupB{filename: filename, tds: tds})
	}

	base.MdtkExit(0)
}

// ----------------------------------------------------------------------------

type ArgsGroupC struct {
	td taskset.TaskData
}

func RunGroupC(a ArgsGroupA, b ArgsGroupB) {
	FlagHas := func(str string) bool { return a.flags.GetData(str).Exist }

	// validation, also check hidden attr, etc.
	base.Exit1_IfHasError(sub.Validate(a.gtname, a.args, b.tds, FlagHas("--all-task")))

	td, err := b.tds.GetTaskData(a.gtname.Split())
	base.Exit1_IfHasError(err)

	switch sub.EnumGroupC_WritePath(FlagHas("--path"), FlagHas("--dir")) {
	case sub.ACT_PATH:
		fmt.Println(string(td.FilePath))
	case sub.ACT_DIR:
		fmt.Println(string(td.FilePath.Dir()))
	default:
		RunGroupD(a, b, ArgsGroupC{td: td})
	}

	base.MdtkExit(0)
}

// ----------------------------------------------------------------------------

func RunGroupD(a ArgsGroupA, b ArgsGroupB, c ArgsGroupC) {
	FlagHas := func(str string) bool { return a.flags.GetData(str).Exist }

	code, err := b.tds.GetTaskStart(a.gtname, a.args, int(a.nestsize))
	base.Exit1_IfHasError(err)

	// td, err := tds.GetTaskData(gtname.Split())
	// -> From the previous steps, we know there is no error, so remove it.
	switch sub.EnumGroupD_RunOrWriteScript(c.td.Lang.IsSub(), FlagHas("--script"), FlagHas("--no-head-script")) {
	case sub.ACT_RUN:
		err := exec.Run(c.td.Lang, string(code), FlagHas("--quiet"), a.oflags.RunInTaskFileDir, string(b.filename.Dir()))
		base.Exit1_IfHasError(err)
	case sub.ACT_SCRIPT:
		fmt.Println(code.GetRunnableScript(c.td.Lang.GetScriptData()))
	case sub.ACT_RAW_SCRIPT:
		fmt.Println(code.GetRawScript())
	}
	
	base.MdtkExit(0)	
}
