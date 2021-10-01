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
	"github.com/spf13/cobra"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "glenv",
	Short: "GitLab environment variables in your shell",
	Long: `
Export env. variables from GitLab.

Jump into your favourite repository folder and GitLab env. variables 
are automatically loaded into your shell. This tool in combination 
with direnv will export your project's env. variables into current shell 
automatically.

This tool doesn't have any configuration.. Only GITLAB_TOKEN variable 
need to be present in your environment.

For more details, visit https://github.com/kiwicom/glenv
`,

	PersistentPreRunE: preRun,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(func() {})

	rootCmd.PersistentFlags().BoolP("debug", "d", false, "Print debug logs")

	rootCmd.SetHelpCommand(&cobra.Command{Hidden: true})
	rootCmd.CompletionOptions.DisableDefaultCmd = true
}

func preRun(cmd *cobra.Command, args []string) error {
	return nil
}
