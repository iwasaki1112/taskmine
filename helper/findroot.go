package helper

import (
	"errors"
	"os"
	"path/filepath"
)

func FindRoot() (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	root := cwd
	for {
		_, err := os.Stat(filepath.Join(root, "go.mod"))
		if os.IsNotExist(err) {
			root = filepath.Dir(root)
		} else if err != nil {
			return "", err
		} else {
			return root, nil
		}

		if root == "/" {
			return "", errors.New("could not find project root")
		}
	}
}
