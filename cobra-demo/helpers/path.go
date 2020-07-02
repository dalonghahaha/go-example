package helpers

import (
	"os"
	"path/filepath"
)

func GetPath() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}
	return dir
}
