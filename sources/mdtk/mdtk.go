package main

import (
	"fmt"
	"mdtk/exec"
	"mdtk/parse"
	"mdtk/help"
	"mdtk/path"
	"mdtk/read"
	"mdtk/grtask"
	"mdtk/taskset"
	"mdtk/code"
	"mdtk/args"
	"mdtk/cache"
	"mdtk/config"
	"mdtk/mdtk_sub"
	"os"
	"strconv"
	_ "embed"
)

func GetFlag () parse.Flag {
	nestsizestr := strconv.FormatUint(uint64(config.Config.NestMaxDepth), 10)

	flags := parse.Flag{}
	flags.Set("--file", []string{"-f"}).SetHasValue("").SetDescription("Select a task file.")
	flags.Set("--nest", []string{"-n"}).SetHasValue(nestsizestr).SetDescription("Set the nest maximum depth of embedded comment (embed/task).\nDefault is " + nestsizestr + ".")
	flags.Set("--quiet", []string{"-q"}).SetDescription("Task output is not sent to standard output.")
	flags.Set("--all-task", []string{"-a"}).SetDescription("Can select private groups and hidden tasks at the command.\nOr show all tasks that include private groups and hidden tasks at task help.")
	flags.Set("--script", []string{"-s"}).SetDescription("Display script.")
	
	flags.Set("--make-cache", []string{"-c"}).SetDescription("Make taskdata cache.")
	flags.Set("--lib", []string{"-l"}).SetHasValue("").SetDescription("Select a library file.\nThis is a special version of --file option.\nNo need to add an extension '.mdtklib'.")
	flags.Set("--make-library", []string{}).SetHasValue("").SetDescription("Make taskdata library.\nValue is library name.")

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

	sub.ReadConfig(flags)

	nestsize := config.Config.NestMaxDepth
	if fd := flags.GetData("--nest"); fd.Exist {
		nestsize = uint(fd.ValueUint())
	}

	// show command/md help
	if checkOptionsThatDoNotRequireTaskFile(flags) {
		sub.MdtkExit(0)
	}

	// validation
	if err := gtname.Validate(); err != nil {
		fmt.Print(err)
		sub.MdtkExit(1)
	}
	if err := task_args.Validate(); err != nil {
		fmt.Print(err)
		sub.MdtkExit(1)
	}
	
	// get Taskfile
	filename := path.Path("")
	if fd := flags.GetData("--file"); fd.Exist {
		filename = path.Path(fd.Value)
	} else {
		filename = read.SearchTaskfile()
	}

	// read Taskfile
	tds := sub.ReadTaskDataSet(filename, flags)

	// make lib
	if fd := flags.GetData("--make-library"); fd.Exist {
		cache.WriteLib(tds, filename.Dir(), fd.Value, int(nestsize))		
		sub.MdtkExit(0)
	}

	// show groups
	if flags.GetData("--groups").Exist {
		help.ShowGroups(filename, tds, flags.GetData("--all-task").Exist)
		sub.MdtkExit(0)
	}
	
	// show task help
	if help.ShouldShowHelp(gtname, tds) {
		help.ShowHelp(filename, gtname, tds, flags.GetData("--all-task").Exist)
		sub.MdtkExit(0)
	}

	// check Private/Hidden Group
	if td, err := tds.GetTaskData(gtname.Split()); err != nil {
		fmt.Print(err)
		sub.MdtkExit(1)
	} else if td.HasAttr(taskset.ATTR_HIDDEN) && !flags.GetData("--all-task").Exist {
		fmt.Println("Private/Hidden group cannot be executed directly.")
		fmt.Printf("[group: %s | task: %s | path: %s]\n", td.Group, td.Task, td.FilePath)
		sub.MdtkExit(1)
	}
	
	code := tds.GetTaskStart(gtname, task_args, int(nestsize))

	if checkOptionsAndWriteScriptToStdout(code, flags) {
		sub.MdtkExit(0)
	}

	exec.Run(string(code), flags.GetData("--quiet").Exist)
	sub.MdtkExit(0)
	
}

// ---------------------------------------------------------------------------------

//go:embed version.txt
var version string

func getVersion() string {
	_, v, _ := args.Arg(version).GetData()
	return v
}

func checkOptionsThatDoNotRequireTaskFile(flags parse.Flag) bool {
	// show command/md help
	if flags.GetData("--version").Exist {
		fmt.Println("mdtk version", getVersion())
		return true
	}

	// show command/md help
	if flags.GetData("--help").Exist {
		help.ShowCommandHelp(flags, 26)
		return true
	}

	if flags.GetData("--manual").Exist {
		help.ShowManual()
		return true
	}

	if flags.GetData("--write-configbase").Exist {
		config.WriteDefaultConfig()
		return true
	}
	

	return false
}


func checkOptionsAndWriteScriptToStdout(codedata code.Code, flags parse.Flag) bool {
	sb := flags.GetData("--script").Exist

	if !(sb) {
		return false
	} 

	c := codedata.RemoveEmbedDescComment().RemoveEmbedArgsComment()
	h := fmt.Sprintln("#!" + exec.GetShell()) //shebang
	h += fmt.Sprintln(exec.GetShHead())
	h += fmt.Sprintln("")
	c = code.Code(h) + c

	fmt.Println(c)
	return true
}


