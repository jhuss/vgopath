package shellscript

import (
	"errors"
	"os"
	"fmt"
	"io/ioutil"
)

func Create(venvName string, virtualPath string, outputPath string) (result bool, err error) {
	if virtualPath == "" {
		currentDir, err := os.Getwd()

		if err != nil {
			return false, errors.New("can't get path for virtual GOPATH\n")
		}

		virtualPath = currentDir
	}

	virtualPath, err = AbsolutePath(virtualPath)
	if err != nil {
		return false, err
	}

	_, err = IsDir(virtualPath)
	if err != nil {
		return false, err
	}

	if outputPath == "" {
		currentDir, err := os.Getwd()

		if err != nil {
			return false, errors.New("can't get output path\n")
		}

		outputPath = currentDir
	}

	outputPath, err = AbsolutePath(outputPath)
	if err != nil {
		return false, err
	}

	_, err = IsDir(outputPath)
	if err != nil {
		return false, err
	}

	var scriptContent = createScript(InitOpts{VenvName:venvName, Gopath:virtualPath})()

	os.MkdirAll(fmt.Sprintf("%s/.vgopath", outputPath), 0755)
	err = ioutil.WriteFile(fmt.Sprintf("%s/.vgopath/%s", outputPath, "activate.sh"), []byte(scriptContent), 0644)
	if err != nil {
		return false, err
	}

	return true, err
}
