/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize .envrc file",
	Long:  "Initialize .envrc with glenv export. The file will be created in current directory.",
	Run: func(cmd *cobra.Command, args []string) {
		if err := initialize(); err != nil {
			errMsg := fmt.Sprintf("error: %v\n", err)
			os.Stderr.WriteString(errMsg)
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}

// this function is executed when you call `glenv init`
// and create `.envrc` file for you in current dir.
func initialize() error {
	body := []byte("eval \"$(glenv export)\"\n")
	err := ioutil.WriteFile(".envrc", body, 0644)
	if err != nil {
		return err
	}
	return nil
}
