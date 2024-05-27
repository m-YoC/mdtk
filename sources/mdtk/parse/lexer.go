package parse

import (
	"fmt"
	"os"
)

// 文字列をスペースで分割する. ただし括弧やクォート系で閉じられているスペースなどは除外する

func LexArgString(str string) ([]string, error) {
	runes := []rune(str)
	res := []string{}
	buf := ""

	const (
		stateText = iota
		stateString
	)
	var quote_buf rune

	lexBase := func(i int, buf string) (int, string, error) {
		switch runes[i] {
		case '\\':
			if i+1 == len(runes) || runes[i+1] == ' ' {
				s := fmt.Sprintln("Lexical error: escape sequence not written correctly.")
				s += fmt.Sprintf("%s\033[4;1;35m%s\033[0m%s\n", string(runes[:i]), "\\", string(runes[i+1:]))
        		return i, buf, fmt.Errorf("%s", s)
			}
			buf += string(runes[i:i+2])
			i++

		default: 
			buf += string(runes[i])
		}

		return i, buf, nil
	}

	lexText := func(i int, buf string) (int, string, int, error) {
		next_state := stateText
		switch runes[i] {
		case ' ':
			if len(buf) > 0 {
				res = append(res, buf)
				buf = ""
			}

		case '\'', '"', '`':
			buf += string(runes[i])
			quote_buf = runes[i]
			next_state = stateString

		default: 
			var err error
			if i, buf, err = lexBase(i, buf); err != nil {
				return i, buf, next_state, err
			}
		}

		return i, buf, next_state, nil
	}

	lexString := func(i int, buf string) (int, string, int, error) {
		next_state := stateString
		switch runes[i] {
		case quote_buf:
			buf += string(runes[i])
			next_state = stateText

		default: 
		var err error
		if i, buf, err = lexBase(i, buf); err != nil {
			return i, buf, next_state, err
		}
		}

		return i, buf, next_state, nil
	}

	state := stateText
	for i := 0; i < len(runes); i++ {
		// fmt.Printf("%s %v |", string(runes[i]), blockstack.Size())
		var err error
		switch state {
		case stateString:
			i, buf, state, err = lexString(i, buf)
		default:
			i, buf, state, err = lexText(i, buf)
		}

		if err != nil {
			return []string{}, err
		}
	}

	if len(buf) > 0 {
		res = append(res, buf)
		buf = ""
	}

	if state != stateText {
		s := fmt.Sprintln("Lexical error: string quotes not closed.")
		s += fmt.Sprintln(str)
        return []string{}, fmt.Errorf("%s", s)
	}

	/*/
	for _, v := range res {
		fmt.Println(v)
	}//*/

	return res, nil
}

func MustLexArgString(str string) []string {
	s, err := LexArgString(str)

	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	return s
}
