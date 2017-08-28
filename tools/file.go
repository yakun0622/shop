package tools

import "os"

func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func MkDir(path string, perm os.FileMode) (err error) {
	if !PathExists(path) {
		err = os.Mkdir(path, perm)
	}
	return
}