package pkg

import (
	"os/user"
	"path"
)

func GetHomeDir() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}

	return usr.HomeDir, nil
}

func GetDefaultDBPath() (string, error) {
	home, err := GetHomeDir()
	if err != nil {
		return "", err
	}

	return path.Join(home, ".todo-clist.data"), nil
}
