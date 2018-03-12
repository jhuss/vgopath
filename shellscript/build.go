package shellscript

import (
	"fmt"
	"errors"
	"os"
)

var options InitOpts

func Create(venvName string, virtualPath string) (result bool, err error) {

	// if path is not provided, set current dir
	if virtualPath == "" {
		currentDir, err := os.Getwd()

		if err != nil {
			return false, errors.New("can't get path for virtual GOPATH\n")
		}

		virtualPath = currentDir
	}

	// absolute path
	virtualPath, err = AbsolutePath(virtualPath)
	if err != nil {
		return false, err
	}

	_, err = IsDir(virtualPath)
	if err != nil {
		return false, err
	}

	// all ok, set values
	options.VenvName = venvName
	options.Gopath = virtualPath

	fmt.Println("set virtual GOPATH \"" + venvName + "\" in \"" + virtualPath + "\"")

	return true, err
}
