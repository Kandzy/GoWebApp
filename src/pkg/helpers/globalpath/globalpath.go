package globalpath

import (
	"os"
	"path/filepath"
)

func GetExecutableLocation() string {
	pathExec, err := os.Executable()

	if err != nil {
		return ""
	}

	return filepath.Dir(pathExec)
}