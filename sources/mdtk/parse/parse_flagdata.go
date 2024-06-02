package parse

import (
	"fmt"
	"strings"
	"strconv"
)

type FlagData struct {
	Name string
	Alias []string
	HasValue bool
	Value string
	DefaultValue string
	Description string
	Exist bool
}

func (fd FlagData) MatchName(flagname string) bool {
	if flagname == fd.Name {
		return true
	} else {
		for _, v := range fd.Alias {
			if flagname == v {
				return true
			}
		}
	}
	
	return false
}

func (fd FlagData) HasValueInt() bool {
	_, err := strconv.ParseInt(fd.Value, 10, 64)
	return err == nil
}

func (fd FlagData) HasValueUint() bool {
	_, err := strconv.ParseUint(fd.Value, 10, 64)
	return err == nil
}

func (fd FlagData) ValueInt() int64 {
	if v, err := strconv.ParseInt(fd.Value, 10, 64); err == nil {
		return v
	}
	v, _ := strconv.ParseInt(fd.DefaultValue, 10, 64)
	return v
}

func (fd FlagData) ValueUint() uint64 {
	if v, err := strconv.ParseUint(fd.Value, 10, 64); err == nil {
		return v
	}
	v, _ := strconv.ParseUint(fd.DefaultValue, 10, 64)
	return v
}

func (fd FlagData) GetHelpStr(head_space_num uint, descline uint) string {
	h := strings.Repeat(" ", int(head_space_num))
	ss := fmt.Sprintf("%s", fd.Name)
	for _, v := range fd.Alias {
		ss += fmt.Sprintf(", %s", v)
	}
	if fd.HasValue {
		ss += "  [+value]"
	}

	s := ""
	descs := strings.Split(fd.Description, "\n")
	for i, v := range descs {
		s += fmt.Sprintf("%s%-" + strconv.Itoa(int(descline)) + "s%s", h, ss, v)
		if i != len(descs)-1 { s += "\n" }
		ss = ""
	}
	// if len(descs) == 1 { s += "\n" }
	
	return s
}




type Flag []FlagData

func (f *Flag) Set(name string, alias []string) *FlagData {
	fd := FlagData{Name: name, Alias: alias, HasValue: false, Exist: false}
	*f = append(*f, fd)
	return &(*f)[len(*f)-1]
}

func (fd *FlagData) SetHasValue(default_value string) *FlagData {
	fd.HasValue = true
	fd.Value = default_value
	fd.DefaultValue = default_value
	return fd
}

func (fd *FlagData) SetDescription(desc string) *FlagData {
	fd.Description = desc
	return fd
}

func (f Flag) GetIndex(flagname string) int {
	for i, fd := range f {
		if fd.MatchName(flagname) {
			return i
		}
	}
	return -1
}

func (f Flag) GetData(flagname string) FlagData {
	i := f.GetIndex(flagname)

	if i < 0 {
		fmt.Printf("Flag.GetData(\"%s\") cannot find data.\n", flagname)
		return FlagData{Alias: []string{}}
	}

	return f[i]
}


func (f Flag) GetHelpStr(head_space_num uint, descline uint) string {
	s := ""
	for _, v := range f {
		s += v.GetHelpStr(head_space_num, descline) + "\n"
	}
	return s
}

