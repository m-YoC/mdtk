package path

import (
    "os"
	"strings"
    "os/user"
    "path/filepath"
)

type Path string

func GetWorkingDir[T Path | string]() (T, error) {
	p, err := os.Getwd()
	if err != nil {
		return T(""), err
	}
	return T(p), nil
}

func (path Path) homeDirToAbs() Path {
    if path[0:2] == "~/" {
        usr, _ := user.Current()
        return Path(strings.Replace(string(path), "~", usr.HomeDir, 1))
    }
    return path
}

func (path Path) GetFileAbsPath() Path {
    x, _ := filepath.Abs(string(path.homeDirToAbs()))
    return Path(x)
}

func (path Path) Dir() Path {
	return Path(filepath.Dir(string(path)))
}


func (base_dir Path) GetSubFilePath(path Path) Path {
    if p := path.homeDirToAbs(); filepath.IsAbs(string(p)) {
        return p
    }
    x, _ := filepath.Abs(filepath.Clean(filepath.Join(string(base_dir), string(path))))
    return Path(x)
}


func (path Path) Ext() string {
    return filepath.Ext(string(path))
}
