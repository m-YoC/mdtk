package lib

import (
	"os"
	"encoding/gob"
)

func zero[T any]() T {
	var zero T
	return zero
}

func WriteStruct[T any](data T, filename string) error {
	file, err := os.Create(filename)
	defer file.Close()
	if err != nil {
		return err
	}

	gob.NewEncoder(file).Encode(data)
	return nil
}

func ReadStruct[T any](filename string) (T, error) {
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		return zero[T](), err
	}

	var data T
	err = gob.NewDecoder(file).Decode(&data)
	if err != nil {
		return zero[T](), err
	}

	return data, nil
}

