package path

import (
	"strings"
    "os/user"
    "path/filepath"
)

type Path string

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
