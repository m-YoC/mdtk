package cache

import (
	"os"
	"mdtk/lib"
	"mdtk/taskset"
	"mdtk/taskset/path"
)


func toCacheName(filename path.Path) string {
	return filename.String() + ".cache"
}

func ExistCacheFile(filename path.Path) bool {
	_, err := os.Stat(toCacheName(filename))
	return err == nil
}

func IsLatestCache(tds taskset.TaskDataSet, filename path.Path) bool {
	status, err := os.Stat(toCacheName(filename))
	if err != nil {
		return false
	}

	// path.Path
	for k, _ := range tds.FilePath {
		substatus, err := os.Stat(string(k))
		if err != nil || status.ModTime().Before(substatus.ModTime()) {
			return false
		}
	}

	return true
}

func WriteCache(tds taskset.TaskDataSet, filename path.Path) error {
	return lib.WriteStruct[taskset.TaskDataSet](tds, toCacheName(filename))
}

func ReadCache(filename path.Path) (taskset.TaskDataSet, error) {
	return lib.ReadStruct[taskset.TaskDataSet](toCacheName(filename))
}