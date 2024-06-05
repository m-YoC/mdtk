package args

import (
	"fmt"
	"regexp"
)

// ARG=value の形式
const arg_reg = "(?P<name>\\w+)[:=](?P<value>.+)"
var arg_rex = regexp.MustCompile("^" + arg_reg + "$")
var validate_arg_rex = arg_rex


type Arg string
type Args []Arg

func (arg Arg) Validate() error {
	if !validate_arg_rex.MatchString(string(arg)) {
		return fmt.Errorf("Validation error: arg notation. Correct is [name=value]. => %s\n", arg)
	}
	return nil
}

func (arg Arg) GetData() (string, string, error) {
	if err := arg.Validate(); err != nil {
		return "", "", err
	}

	res := arg_rex.FindStringSubmatch(string(arg))
	return res[arg_rex.SubexpIndex("name")], res[arg_rex.SubexpIndex("value")], nil
}


func (args Args) Validate() error {
	for _, arg := range args {
		if err := arg.Validate(); err != nil {
			return err
		}
	}
	return nil
}

func ToArgs(strs ...string) Args {
	args := make(Args, len(strs))
	for i, _ := range strs {
		args[i] = Arg(strs[i])
	} 
	return args
}



