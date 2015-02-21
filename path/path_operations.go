package path

import (
	"os/user"
	"path/filepath"
	"strings"
)

const ImgBasePath = "~/Pictures/wallpapers"

func PreProcessPath(path string) (string, error) {
	p := ""
	usr, _ := user.Current()
	dir := usr.HomeDir

	if path[:2] == "~/" {
		path = strings.Replace(path, "~/", "", 1)
		path = filepath.Join(dir, path)
	}

	p, err := filepath.Abs(path)
	return p, err
}
