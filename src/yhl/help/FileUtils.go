package help

import (
	"os"
)

func PathExist(pathName string) bool {
	_, err := os.Stat(pathName)

	return !os.IsNotExist(err)
}

func MkDirPath(pathRoute string) error {
	if !PathExist(pathRoute) {
		return os.MkdirAll(pathRoute, os.ModePerm)
	}

	return nil
}
