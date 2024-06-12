package sub

import (
	"fmt"
	"mdtk/base"
	"mdtk/taskset/path"
	"mdtk/taskset/read"
	"mdtk/taskset"
	"mdtk/cache"
	_ "embed"
	"sync"
)


func ReadTaskDataSet(filename path.Path, has_make_cache_flag bool) taskset.TaskDataSet {
	// check filename -> *.md / *.mdtklib
	ext := filename.Ext()
	switch ext {
	case ".md":
		return readTaskDataSetMdAsync(filename, has_make_cache_flag)
	case ".mdtklib":
		return readTaskDataSetLib(filename)
	default:
		fmt.Printf("Extension of [%s] is not '.md' or '.mdtklib'.\n", filename)
		base.MdtkExit(1)
	}

	return taskset.TaskDataSet{}
}

func readTaskDataSetMd(filename path.Path, make_cache_flag bool) taskset.TaskDataSet {
	if cache.ExistCacheFile(filename) {
		tds, err := cache.ReadCache(filename)
		// fmt.Println("from cache")
		if err != nil {
			fmt.Print(err)
			base.MdtkExit(1)
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
		base.MdtkExit(1)
	}

	if make_cache_flag {
		cache.WriteCache(tds, filename)
		fmt.Printf("mdtk: Made %s.cache.\n", filename)
	}

	return tds
}


/**
<<has cache>>
main goroutine
[read cache] -- <If: latest> _y_ {return: cache} 
                            \_n_ [wait: read task] -- [get task] ---- {return: task}
sub goroutine                                       /             \
[read task ] ---------------------------------------               -- [write cache]

<<no cache>>
main goroutine
[read task] ---- {return: task}
sub goroutine  \
                --<If: has make_flag> _y_ [write cache]
*/
func readTaskDataSetMdAsync(filename path.Path, make_cache_flag bool) taskset.TaskDataSet {
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
			base.MdtkExit(1)
		}

		if cache.IsLatestCache(tds, filename) {
			return tds
		}
		make_cache_flag = true

		res = <-ch
		err = <-cherr
		if err != nil {
			fmt.Print(err)
			base.MdtkExit(1)
		}
	} else {
		var err error
		res, err = read.ReadTask(filename)
		if err != nil {
			fmt.Print(err)
			base.MdtkExit(1)
		}
	}

	if make_cache_flag {
		var wg sync.WaitGroup
		wg.Add(1)
		base.AddFinalize(func() { wg.Wait() })

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
		base.MdtkExit(1)
	}

	return tds
}
