package shellscript

import (
	"path/filepath"
	"os"
	"errors"
)

func PathInfo(p string) (info os.FileInfo, err error) {
	return os.Stat(p)
}

func PathExist(p string) (exist bool, err error) {
	_, pathErr := PathInfo(p)
	exist = !os.IsNotExist(pathErr)

	if !exist {
		return false, errors.New("path does not exist\n")
	}

	return exist, nil
}

func AbsolutePath(p string) (fullPath string, err error) {
	fullPath, err = filepath.Abs(p)

	if err == nil {
		_, err = PathExist(fullPath)
	} else {
		err = errors.New("can't get the absolute path directory\n")
	}

	return fullPath, err
}

func IsDir(p string) (result bool, err error) {
	pathInfo, err := PathInfo(p)
	isDir := pathInfo.Mode().IsDir()

	if !isDir {
		return false, errors.New("path is not a directory\n")
	}

	return isDir, nil
}
