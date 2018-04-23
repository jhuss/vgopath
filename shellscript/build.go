package shellscript

import (
	"errors"
	"os"
	"fmt"
)

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
	var scriptContent = createScript(InitOpts{VenvName:venvName,Gopath:virtualPath})()
	fmt.Println(scriptContent)

	return true, err
}
