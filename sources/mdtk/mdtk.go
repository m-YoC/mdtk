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
	"os"
	_ "embed"
)

func GetFlag () parse.Flag {
	flags := parse.Flag{}
	flags.Set("--file", []string{"-f"}).SetHasValue("").SetDescription("Select a task file.")
	flags.Set("--nest", []string{"-n"}).SetHasValue("20").SetDescription("Set the nest maximum times of embedded comment (embed/task).\nDefault is 20.")
	flags.Set("--quiet", []string{"-q"}).SetDescription("Task output is not sent to standard output.")
	flags.Set("--all-task", []string{"-a"}).SetDescription("Can select private groups at the command.\nOr show all tasks that include private groups at task help.")
	flags.Set("--script", []string{"-s"}).SetDescription("Display run-script.\n(= shebang + '" + exec.GetShHead() + "' + expanded-script)\nIf --debug option is not set, do not run.")
	flags.Set("--debug", []string{"-d"}).SetDescription("Display expanded-script and run.\nIf --script option is set, display run-script.")
	
	flags.Set("--make-cache", []string{"-c"}).SetDescription("Make taskdata cache.")
	flags.Set("--lib", []string{"-l"}).SetHasValue("").SetDescription("Select a library file.\nThis is a special version of --file option.\nNo need to add an extension '.mdtklib'.")
	flags.Set("--make-library", []string{}).SetHasValue("").SetDescription("Make taskdata library.\nValue is library name.")

	flags.Set("--version", []string{"-v"}).SetDescription("Show version.")
	flags.Set("--help", []string{"-h"}).SetDescription("Show command help.")
	flags.Set("--manual", []string{"-m"}).SetDescription("Show mdtk manual.")
	flags.Set("--write-configbase", []string{}).SetDescription("Write config base file to current directory.")
	return flags
}

func LibToFile(flags parse.Flag) parse.Flag {
	fi := flags.GetIndex("--file")
	li := flags.GetIndex("--lib")
	if !flags[fi].Exist && flags[li].Exist {
		flags[fi].Exist = true
		flags[fi].Value = flags[li].Value + ".mdtklib"
	}
	return flags
}

func main() {
	gtname_str, flags, task_args_strarr := parse.Parse(os.Args, GetFlag())
	flags = LibToFile(flags)
	gtname := grtask.GroupTask(gtname_str)
	task_args := args.ToArgs(task_args_strarr...)

	if fd := flags.GetData("--file"); fd.Exist {
		config.ReadConfig(string(path.Path(fd.Value).Dir()))
	} else {
		config.ReadConfig("")
	}

	// show command/md help
	if checkOptionsThatDoNotRequireTaskFile(flags) {
		return
	}

	// validation
	if err := gtname.Validate(); err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	if err := task_args.Validate(); err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	
	// get Taskfile
	filename := path.Path("")
	if fd := flags.GetData("-f"); fd.Exist {
		filename = path.Path(fd.Value)
	} else {
		filename = read.SearchTaskfile()
	}

	// read Taskfile
	tds := readTaskDataSet(filename, flags)

	// make lib
	if fd := flags.GetData("--make-library"); fd.Exist {
		cache.WriteLib(tds, filename.Dir(), fd.Value, int(flags.GetData("--nest").ValueUint()))		
		return
	}
	
	// show task help
	if help.ShouldShowHelp(gtname, tds) {
		help.ShowHelp(filename, gtname, tds, flags.GetData("--all-task").Exist)
		return
	}

	// check PrivateGroup
	if err := gtname.ValidatePublic(); err != nil && !flags.GetData("--all-task").Exist {
		fmt.Print(err)
		os.Exit(1)
	}
	
	code := tds.GetTaskStart(gtname, task_args, int(flags.GetData("--nest").ValueUint()))

	if checkOptionsAndWriteScriptToStdout(code, flags) {
		return
	}

	exec.Run(string(code), flags.GetData("--quiet").Exist)
	
	
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

func readTaskDataSet(filename path.Path, flags parse.Flag) taskset.TaskDataSet {
	// check filename -> *.md / *.mdtklib
	ext := filename.Ext()
	switch ext {
	case ".md":
		return readTaskDataSetMd(filename, flags)
	case ".mdtklib":
		return readTaskDataSetLib(filename)
	default:
		fmt.Printf("Extension of [%s] is not '.md' or '.mdtklib'.\n", filename)
		os.Exit(1)
	}

	return taskset.TaskDataSet{}
}

func readTaskDataSetMd(filename path.Path, flags parse.Flag) taskset.TaskDataSet {
	make_cache_flag := flags.GetData("--make-cache").Exist

	if cache.ExistCacheFile(filename) {
		tds, err := cache.ReadCache(filename)
		// fmt.Println("from cache")
		if err != nil {
			fmt.Print(err)
			os.Exit(1)
		}

		if cache.IsLatestCache(tds, filename) {
			return tds
		} else {
			make_cache_flag = true
		}
	}

	tds, err := read.ReadTask(filename)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	if make_cache_flag {
		cache.WriteCache(tds, filename)
		fmt.Printf("mdtk: Made %s.cache.\n", filename)
	}

	return tds
}

func readTaskDataSetLib(filename path.Path) taskset.TaskDataSet {
	tds, err := cache.ReadLib(filename)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return tds
}



func checkOptionsAndWriteScriptToStdout(codedata code.Code, flags parse.Flag) bool {
	sb := flags.GetData("--script").Exist
	db := flags.GetData("--debug").Exist

	if !(sb || db) {
		return false
	} 

	c := codedata.RemoveEmbedArgsComment()
	if sb {
		h := fmt.Sprintln("#!" + exec.Shname())
		h += fmt.Sprintln(exec.GetShHead())
		h += fmt.Sprintln("")
		c = code.Code(h) + c
	}

	if db {
		cc := fmt.Sprintln("--script--")
		cc += string(c)
		cc += fmt.Sprintln("\n----------")
		c = code.Code(cc)
	}

	fmt.Println(c)
	return !db
}


