package config

import (
	"fmt"
	"os"
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
	Shell string
	ScriptHeadSet string
	LangAlias []string
	Pager []string
	PagerMinLimit uint
}

var Config cfg

func init() {
	setConfig(strings.Split(dflt, "\n"))

	if os.Getenv("SHELL") != "" {
		Config.Shell = os.Getenv("SHELL")
	}

	if os.Getenv("PAGER") != "" {
		Config.Pager = []string{os.Getenv("PAGER")}
	}
	
}

// ------------------------------------------------------------------------------

//dir: (-f path -> pwd) -> home -> init
func getConfigPath(dir string) string {
	if dir != "" {
		if _, err := os.Stat(filepath.Join(dir, configName)); err == nil {
			return filepath.Join(dir, configName)
		}
	}

	if _, err := os.Stat(filepath.Join("./", configName)); err == nil {
		return filepath.Join("./", configName)
	}

	usr, _ := user.Current()
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

		switch k {
		case "shell":
			Config.Shell = strings.TrimSpace(v)
		case "script_head_set":
			Config.ScriptHeadSet = strings.TrimSpace(v)
		case "acceptable_langs":
			if s, err := parse.LexArgString(v); err != nil {
				return err
			} else {
				Config.LangAlias = s
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
	if _, err := os.Stat(filepath.Join("./", configName)); err == nil {
		return
	}

	os.WriteFile(configName, []byte(dflt), 0666)
}