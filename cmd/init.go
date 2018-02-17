package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
	"regexp"
	"errors"
)

var projectName string

var initCmd = &cobra.Command{
	Use: "init -n|--name <name>",
	DisableFlagsInUseLine: true,
	Short: "create virtual environment",
	RunE: func(cmd *cobra.Command, args []string) error {

		if projectName == "" {
			return errors.New("name can't be empty")
		}

		r, _ := regexp.Compile("[a-z0-9]+")
		filtered := r.FindString(projectName)

		if filtered != projectName {
			return errors.New("name can be only alphanumeric\n")
		}

		fmt.Println("create script " + filtered)

		return nil
	},
}

func init() {
	initCmd.Flags().StringVarP(&projectName,"name", "n", "", "name of the virtual environment")
	initCmd.MarkFlagRequired("name")

	rootCmd.AddCommand(initCmd)
}
