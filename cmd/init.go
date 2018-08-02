package cmd

import (
	"github.com/spf13/cobra"
	"regexp"
	"errors"
	"vgopath/shellscript"
)

var projectName string
var virtualPath string
var outputPath string

var initCmd = &cobra.Command{
	Use: "init (-n|--name <name>) [(-p|--gopath <path>)]",
	Short: "create virtual environment",
	DisableFlagsInUseLine: true,
	RunE: func(cmd *cobra.Command, args []string) error {

		if projectName == "" {
			return errors.New("name can't be empty")
		}

		r, _ := regexp.Compile("[a-z0-9]+")
		venvName := r.FindString(projectName)

		if venvName != projectName {
			return errors.New("name can be only alphanumeric\n")
		}

		_, err := shellscript.Create(venvName, virtualPath, outputPath)
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	initCmd.Flags().StringVarP(&projectName,"name", "n", "", "name of the virtual environment")
	initCmd.Flags().StringVarP(&virtualPath,"gopath", "p", "", "virtual value of GOPATH")
	initCmd.Flags().StringVarP(&outputPath,"output", "o", "", "output path for script")

	initCmd.MarkFlagRequired("name")

	rootCmd.AddCommand(initCmd)
}
