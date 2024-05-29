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
	"mdtk/args"
	"mdtk/cache"
	"os"
	_ "embed"
)

//go:embed version.txt
var version string

func getVersion() string {
	_, v, _ := args.Arg(version).GetData()
	return v
}

func GetFlag () parse.Flag {
	flags := parse.Flag{}
	flags.Set("--file", []string{"-f"}).SetHasValue("").SetDescription("Specify a task file.")
	flags.Set("--nest", []string{"-n"}).SetHasValue("20").SetDescription("Set the nest maximum times of embedded comment (embed/task).\nDefault is 20.")
	flags.Set("--make-cache", []string{"-c"}).SetDescription("Make taskdata cache.")
	flags.Set("--debug", []string{}).SetDescription("Show run-script.")
	flags.Set("--version", []string{"-v"}).SetDescription("Show version.")
	flags.Set("--help", []string{"-h"}).SetDescription("Show command help.")
	flags.Set("--md-help", []string{}).SetDescription("Show Markdown taskfile help.")
	flags.Set("--task-help-all", []string{}).SetDescription("Show all tasks that include private groups at task help.")
	return flags
}

func main() {
	gtname_str, flags, task_args_strarr := parse.Parse(os.Args, GetFlag())
	gtname := grtask.GroupTask(gtname_str)
	task_args := args.ToArgs(task_args_strarr...)

	// show command/md help
	if flags.GetData("--help").Exist {
		help.ShowCommandHelp(flags, 26)
		return
	}

	if flags.GetData("--md-help").Exist {
		help.ShowMarkdownHelp()
		return
	}

	// show command/md help
	if flags.GetData("--version").Exist {
		fmt.Println("mdtk version", getVersion())
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
	var tds taskset.TaskDataSet 
	if cache.ExistCacheFile(filename) {
		var err error
		tds, err = cache.ReadCache(filename)
		// fmt.Println("from cache")
		if err != nil {
			fmt.Print(err)
			os.Exit(1)
		}
		if !cache.IsLatestCache(tds, filename) {
			var err error
			tds, err = read.ReadTask(filename)
			if err != nil {
				fmt.Print(err)
				os.Exit(1)
			}
			fmt.Println("update cache")
			cache.WriteCache(tds, filename)
		}

	} else {
		var err error
		tds, err = read.ReadTask(filename)
		if err != nil {
			fmt.Print(err)
			os.Exit(1)
		}

		if flags.GetData("--make-cache").Exist {
			cache.WriteCache(tds, filename)
		}
	}

	
	// show task help
	if help.ShouldShowHelp(gtname, tds) {
		// help.ShowHelp(filename, tds, flags.GetData("--task-help-all").Exist)
		return
	}

	// fmt.Println(md)

	// check PrivateGroup
	if err := gtname.ValidatePublic(); err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	
	code := tds.GetTaskStart(gtname, task_args, int(flags.GetData("--nest").ValueUint()))

	if flags.GetData("--debug").Exist {
		printcode := code.RemoveEmbedArgsComment()
		fmt.Println("--script--")
		fmt.Println(printcode)
		fmt.Println("----------")
		fmt.Println("")
	}

	exec.Run(string(code))
	
	
}

