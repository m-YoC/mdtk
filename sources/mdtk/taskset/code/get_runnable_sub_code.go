package code

import (
	"mdtk/taskset/grtask"
)

func (c Code) GetRunnableSubCode(tf TaskDataSetInterface, gtname grtask.GroupTask, nestsize int) (Code, error) {
	var f bool
	if c, f = c.ApplyConfigOnce(gtname); f {
		return c, nil
	}
	var err error
	c, err = c.ApplyEmbedCodes(tf, nestsize)
	if err != nil {
		return "", err
	}
	return c, nil
}

