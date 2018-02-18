package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
	"regexp"
	"errors"
	"os"
	"path/filepath"
)

var projectName string
var virtualPath string

var initCmd = &cobra.Command{
	Use: "init (-n|--name <name>) [(-p|--gopath <path>)]",
	DisableFlagsInUseLine: true,
	Short: "create virtual environment",
	RunE: func(cmd *cobra.Command, args []string) error {

		// check name
		if projectName == "" {
			return errors.New("name can't be empty")
		}

		r, _ := regexp.Compile("[a-z0-9]+")
		filtered := r.FindString(projectName)

		if filtered != projectName {
			return errors.New("name can be only alphanumeric\n")
		}

		// check path
		if virtualPath == "" {
			currentDir, wdErr := os.Getwd()

			if wdErr != nil {
				return errors.New("can't get path for virtual GOPATH\n")
			}

			virtualPath = currentDir
		}

		virtualPath = func() string {p,_:=filepath.Abs(virtualPath);return p}()
		if _, err := os.Stat(virtualPath); os.IsNotExist(err) {
			return errors.New("gopath can't be used\n")
		}

		fmt.Println("set virtual GOPATH \"" + filtered + "\" in \"" + virtualPath + "\"")

		return nil
	},
}

func init() {
	initCmd.Flags().StringVarP(&projectName,"name", "n", "", "name of the virtual environment")
	initCmd.Flags().StringVarP(&virtualPath,"gopath", "p", "", "virtual value of GOPATH")

	initCmd.MarkFlagRequired("name")

	rootCmd.AddCommand(initCmd)
}
