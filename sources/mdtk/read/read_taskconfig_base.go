package read

import (
	"strings"
	"regexp"
)

func getTaskConfigBase[KT comparable, VT any](md Markdown, rex *regexp.Regexp, f func(*map[KT]VT, string)) (map[KT]VT, error) {
	heads := rex.FindAllStringSubmatch(string(md), -1)
	indices := rex.FindAllStringIndex(string(md), -1)

	res := map[KT]VT{}
	for i, head := range heads {
		block := head[rex.SubexpIndex("block")]

		code, err := md.ExtractCode(indices[i][1], block)
		if err != nil {
			return map[KT]VT{}, err
		}

		codearr := strings.Split(string(code), "\n")
	
		for _, v := range codearr {
			f(&res, v)
		}
	}
	
	return res, nil
}
