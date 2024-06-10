package sub

import (
	"fmt"
	"mdtk/parse"
	"mdtk/path"
	"mdtk/read"
	"mdtk/taskset"
	"mdtk/cache"
	_ "embed"
	"sync"
)


func ReadTaskDataSet(filename path.Path, flags parse.Flag) taskset.TaskDataSet {
	// check filename -> *.md / *.mdtklib
	ext := filename.Ext()
	switch ext {
	case ".md":
		return readTaskDataSetMdAsync(filename, flags)
	case ".mdtklib":
		return readTaskDataSetLib(filename)
	default:
		fmt.Printf("Extension of [%s] is not '.md' or '.mdtklib'.\n", filename)
		MdtkExit(1)
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
			MdtkExit(1)
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
		MdtkExit(1)
	}

	if make_cache_flag {
		cache.WriteCache(tds, filename)
		fmt.Printf("mdtk: Made %s.cache.\n", filename)
	}

	return tds
}



func readTaskDataSetMdAsync(filename path.Path, flags parse.Flag) taskset.TaskDataSet {
	make_cache_flag := flags.GetData("--make-cache").Exist
	var res taskset.TaskDataSet

	if cache.ExistCacheFile(filename) {
		ch := make(chan taskset.TaskDataSet)
		cherr := make(chan error)
		
		go func() {
			defer close(ch)
			defer close(cherr)
			tds, err := read.ReadTask(filename)
			ch <- tds
			cherr <- err
		}()

		tds, err := cache.ReadCache(filename)
		if err != nil {
			fmt.Print(err)
			MdtkExit(1)
		}

		if cache.IsLatestCache(tds, filename) {
			return tds
		}
		make_cache_flag = true

		res = <-ch
		err = <-cherr
		if err != nil {
			fmt.Print(err)
			MdtkExit(1)
		}
	} else {
		var err error
		res, err = read.ReadTask(filename)
		if err != nil {
			fmt.Print(err)
			MdtkExit(1)
		}
	}

	if make_cache_flag {
		var wg sync.WaitGroup
		wg.Add(1)
		AddFinalize(func() { wg.Wait() })

		go func() {
			cache.WriteCache(res, filename)
			wg.Done()
		}()
	}

	return res
}

func readTaskDataSetLib(filename path.Path) taskset.TaskDataSet {
	tds, err := cache.ReadLib(filename)

	if err != nil {
		fmt.Println(err)
		MdtkExit(1)
	}

	return tds
}
