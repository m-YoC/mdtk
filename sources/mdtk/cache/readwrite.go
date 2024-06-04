package cache

import (
	"os"
	"encoding/gob"
	"mdtk/taskset"
)

func writeBase(tds taskset.TaskDataSet, filename string) error {
	file, err := os.Create(filename)
	defer file.Close()
	if err != nil {
		return err
	}

	gob.NewEncoder(file).Encode(tds)
	return nil
}

func readBase(filename string) (taskset.TaskDataSet, error) {
	file, err := os.Open(filename)
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



