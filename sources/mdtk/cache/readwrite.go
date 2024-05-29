package cache

import (
	"os"
	"encoding/gob"
	"mdtk/path"
	"mdtk/taskset"
)

func toCacheName(filename path.Path) string {
	return string(filename) + ".cache"
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

	for k, _ := range tds.FilePath {
		substatus, err := os.Stat(string(k))
		if err != nil || status.ModTime().Before(substatus.ModTime()) {
			return false
		}
	}

	return true
}

func WriteCache(tds taskset.TaskDataSet, filename path.Path) error {
	file, err := os.Create(toCacheName(filename))
	defer file.Close()
	if err != nil {
		return err
	}

	gob.NewEncoder(file).Encode(tds)
	return nil
}

func ReadCache(filename path.Path) (taskset.TaskDataSet, error) {
	file, err := os.Open(toCacheName(filename))
	defer file.Close()
	if err != nil {
		return taskset.TaskDataSet{}, err
	}

	tds := taskset.TaskDataSet{}
	err = gob.NewDecoder(file).Decode(&tds)
	if err != nil {
		return taskset.TaskDataSet{}, err
	}

	return tds, nil
}
