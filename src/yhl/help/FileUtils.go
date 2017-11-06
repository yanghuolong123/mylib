package help

import (
	"os"
)

func PathExist(pathName string) bool {
	_, err := os.Stat(pathName)

	return !os.IsNotExist(err)
}
