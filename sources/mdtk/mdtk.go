package main

import (
	"fmt"
	"mdtk/base"
	"mdtk/exec"
	"mdtk/parse"
	"mdtk/help"
	"mdtk/taskset/path"
	"mdtk/taskset/read"
	"mdtk/taskset/grtask"
	"mdtk/taskset"
	"mdtk/args"
	"mdtk/cache"
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
	flags.Set("--file", []string{"-f"}).SetHasValue("")
	flags.Back().SetDescription("Select a task file.")

	flags.Set("--nest", []string{"-n"}).SetHasValue(nestsizestr)
	flags.Back().SetDescription("Set the nest maximum depth of embedded comment (embed/task).\nDefault is " + nestsizestr + ".")
	
	flags.Set("--run-in-filedir", []string{"--rfd"})
	flags.Back().SetDescription("Not run task in working directory, but run it in taskfile directory.")
	
	flags.Set("--quiet", []string{"-q"})
	flags.Back().SetDescription("Task output is not sent to standard output.")
	
	flags.Set("--all-task", []string{"-a"})
	flags.Back().SetDescription("Can select private groups and hidden tasks at the command.\nOr show all tasks that include private groups and hidden tasks at task help.")
	
	flags.Set("--script", []string{"-s"}).SetDescription("Display script.")
	flags.Set("--no-head-script", []string{"-S"}).SetDescription("Display script (No shebang, etc.).")

	flags.Set("--path", []string{}).SetDescription("Get file path of selected task.")
	flags.Set("--dir", []string{}).SetDescription("Get directory of taskfile in which selected task is written.")
	
	// -------------------------------------------------------------------------

	flags.Set("--make-cache", []string{"-c"}).SetDescription("Make taskdata cache.")
	flags.Set("--lib", []string{"-l"}).SetHasValue("")
	flags.Back().SetDescription("Select a library file.\nThis is a special version of --file option.\nNo need to add an extension '.mdtklib'.")
	flags.Set("--make-library", []string{}).SetHasValue("")
	flags.Back().SetDescription("Make taskdata library.\nValue is library name.")

	// -------------------------------------------------------------------------

	flags.Set("--version", []string{"-v"}).SetDescription("Show version.")
	flags.Set("--groups", []string{"-g"}).SetDescription("Show groups.")
	flags.Set("--help", []string{"-h"}).SetDescription("Show command help.")
	flags.Set("--manual", []string{"-m"}).SetDescription("Show mdtk manual.")
	flags.Set("--write-configbase", []string{}).SetDescription("Write config base file to current directory.")
	return flags
}


func main() {
	gtname_str, flags, task_args_strarr := parse.Parse(os.Args, GetFlag())
	flags = sub.LibToFile(flags)
	gtname := grtask.GroupTask(gtname_str)
	task_args := args.ToArgs(task_args_strarr...)

	sub.ReadConfig(flags.GetData("--file"))
	nestsize := sub.GetNestSize(flags.GetData("--nest"))

	if flags.GetData("--version").Exist {
		fmt.Println("mdtk version", getVersion())
		base.MdtkExit(0)
	}

	// show command/md help
	if flags.GetData("--help").Exist {
		help.ShowCommandHelp(flags, 26)
		base.MdtkExit(0)
	}

	if flags.GetData("--manual").Exist {
		help.ShowManual()
		base.MdtkExit(0)
	}

	if flags.GetData("--write-configbase").Exist {
		config.WriteDefaultConfig()
		base.MdtkExit(0)
	}

	// get Taskfile
	filename := path.Path("")
	if fd := flags.GetData("--file"); fd.Exist {
		filename = path.Path(fd.Value)
	} else {
		filename = read.SearchTaskfile()
	}

	// get directory of root Taskfile
	dir := filename.Dir()

	// read Taskfile
	tds := sub.ReadTaskDataSet(filename, flags.GetData("--make-cache").Exist)

	// make lib
	if fd := flags.GetData("--make-library"); fd.Exist {
		cache.WriteLib(tds, dir, fd.Value, int(nestsize))		
		base.MdtkExit(0)
	}

	all_task_flag := flags.GetData("--all-task").Exist

	// show groups
	if flags.GetData("--groups").Exist {
		help.ShowGroups(filename, tds, all_task_flag)
		base.MdtkExit(0)
	}
	
	// show task help
	if help.ShouldShowHelp(gtname, tds) {
		help.ShowHelp(filename, gtname, tds, all_task_flag)
		base.MdtkExit(0)
	}

	// validation
	if err := sub.Validate(gtname, task_args, tds, all_task_flag); err != nil {
		fmt.Print(err)
		base.MdtkExit(1)
	}

	if flags.GetData("--path").Exist {
		if td, err := tds.GetTaskData(gtname.Split()); err != nil {
			fmt.Print(err)
			base.MdtkExit(1)
		} else {
			fmt.Println(string(td.FilePath))
			base.MdtkExit(0)
		}
	}

	if flags.GetData("--dir").Exist {
		if td, err := tds.GetTaskData(gtname.Split()); err != nil {
			fmt.Print(err)
			base.MdtkExit(1)
		} else {
			fmt.Println(string(td.FilePath.Dir()))
			base.MdtkExit(0)
		}
	}
	
	code, err := tds.GetTaskStart(gtname, task_args, int(nestsize))
	if err != nil {
		fmt.Print(err)
		base.MdtkExit(1)
	}

	// td, err := tds.GetTaskData(gtname.Split())
	// -> From the previous steps, we know there is no error, so remove it.
	is_not_shell_langs := base.PairFirst(tds.GetTaskData(gtname.Split())).Lang != taskset.ShellLangs
	if flags.GetData("--no-head-script").Exist || is_not_shell_langs {
		fmt.Println(code.GetRawScript())
		base.MdtkExit(0)
	}

	if flags.GetData("--script").Exist {
		fmt.Println(code.GetRunnableScript())
		base.MdtkExit(0)
	}

	if err := exec.Run(string(code), string(dir), 
	                   flags.GetData("--quiet").Exist,
					   flags.GetData("--run-in-filedir").Exist); err != nil {
		fmt.Print(err)
		base.MdtkExit(1)
	}
	base.MdtkExit(0)
	
}



