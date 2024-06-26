package config

import (
	"fmt"
	"os"
	"runtime"
	"os/user"
	"strings"
	"strconv"
	"mdtk/args"
	"mdtk/parse"
	"path/filepath"
	_ "embed"
)

const configName = ".mdtkconfig"


//go:embed default_config.txt
var dflt string

type cfg struct {
	Shell []string
	ScriptHeadSet string
	LangAlias []string

	PowerShell []string
	PwShHeadSet string
	LangAliasPwSh []string

	LangAliasSub []string

	NestMaxDepth uint
	Pager []string
	PagerMinLimit uint
}

var Config cfg

func init() {
	setConfig(strings.Split(dflt, "\n"))

	if os.Getenv("SHELL") != "" {
		Config.Shell = []string{os.Getenv("SHELL"), "-c"}
	}

	if runtime.GOOS == "windows" {
		Config.PowerShell = []string{"powershell", "-c"}
	}

	if os.Getenv("PAGER") != "" {
		Config.Pager = []string{os.Getenv("PAGER")}
	}
	
}

func GetMergedLangAlias() []string {
	return append(append(Config.LangAlias, Config.LangAliasPwSh...), Config.LangAliasSub...)
}

// ------------------------------------------------------------------------------

//dir: (-f path -> pwd) -> home -> init
func getConfigPath(dir string) string {
	if _, err := os.Stat(filepath.Join(dir, configName)); err == nil {
		return filepath.Join(dir, configName)
	}

	usr, _ := user.Current()
	// fmt.Println(usr.HomeDir)
	if _, err := os.Stat(filepath.Join(usr.HomeDir, configName)); err == nil {
		return filepath.Join(usr.HomeDir, configName)
	}

	return ""
}

func setConfig(data []string) error {
	args := args.ToArgs(data...)
	for _, a := range args {
		// If it is a comment line, ignore it.
		/* if a[0:1] == "#" {
			continue
		}*/

		k, v, err := a.GetData()
		if err != nil {
			continue
		}
		if strings.TrimSpace(v) == "" {
			return fmt.Errorf("%s: [%s] has empty value.\n", configName, k)
		}

		switch k {
		case "shell":
			if s, err := parse.LexArgString(v); err != nil {
				return err
			} else {
				Config.Shell = s
			}
		case "script_head_set":
			Config.ScriptHeadSet = strings.TrimSpace(v)
		case "acceptable_langs":
			if s, err := parse.LexArgString(v); err != nil {
				return err
			} else {
				Config.LangAlias = s
			}
		case "powershell":
			if s, err := parse.LexArgString(v); err != nil {
				return err
			} else {
				Config.PowerShell = s
			}
		case "powershell_head_set":
			Config.PwShHeadSet = strings.TrimSpace(v)
		case "powershell_langs":
			if s, err := parse.LexArgString(v); err != nil {
				return err
			} else {
				Config.LangAliasPwSh = s
			}
		case "acceptable_sub_langs":
			if s, err := parse.LexArgString(v); err != nil {
				return err
			} else {
				Config.LangAliasSub = s
			}
		case "nest_max_depth":
			if vv, err := strconv.ParseUint(strings.TrimSpace(v), 10, 64); err != nil {
				return err
			} else {
				Config.NestMaxDepth = uint(vv)
			}
		case "pager":
			if s, err := parse.LexArgString(v); err != nil {
				return err
			} else {
				Config.Pager = s
			}
		case "pager_min_limit":
			if vv, err := strconv.ParseUint(strings.TrimSpace(v), 10, 64); err != nil {
				return err
			} else {
				Config.PagerMinLimit = uint(vv)
			}
		default:
			return fmt.Errorf("%s: [%s] is invalid parameter.\n", configName, k)
		}
	}

	return nil
}


func ReadConfig(dir string) error {
	p := getConfigPath(dir)
	if p == "" {
		return nil
	}

	b, _ := os.ReadFile(p)
	strs := strings.Split(strings.Replace(string(b), "\r\n", "\n", -1), "\n")

	return setConfig(strs)
}


func WriteDefaultConfig() {
	if _, err := os.Stat(filepath.Join(".", configName)); err == nil {
		return
	}

	os.WriteFile(configName, []byte(dflt), 0666)
}