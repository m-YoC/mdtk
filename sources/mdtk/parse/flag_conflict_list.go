package parse

import (
	"fmt"
)

type FlagConflictList struct {
	List [][]string
}

func (l *FlagConflictList) Conflict(fstrs ...string) {
	l.List = append(l.List, fstrs)
}

func (l FlagConflictList) Check(flags Flag) error {
	s := ""
	b := false
	for _, conflictset := range l.List {
		exists := []string{}
		for _, flagname := range conflictset {
			if flags.GetData(flagname).Exist {
				exists = append(exists, flagname)
			}
		}
		if len(exists) > 1 {
			s += fmt.Sprintf("Option Conflict!: %v cannot be used at the same time.\n", exists)
			s += fmt.Sprintln("List of options that cause conflicts:")
			s += fmt.Sprintf(" %v\n", conflictset)
			b = true
		}
		
	}

	if b {
		return fmt.Errorf("%s", s)
	}
	return nil
}
