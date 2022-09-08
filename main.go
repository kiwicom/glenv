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
package main

import (
	"github.com/kiwicom/glenv/cmd"
	"github.com/kiwicom/glenv/internal/glenv"
	"github.com/spf13/cobra"
)

// version is set by goreleaser, via -ldflags="-X 'main.version=...'".
var version = "development"

func main() {

	cobra.OnInitialize(func() {})

	// rootCmd represents the base command when called without any subcommands
	var rootCmd = &cobra.Command{
		Use:     "glenv",
		Short:   "GitLab environment variables in your shell",
		Version: version,
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

	rootCmd.PersistentFlags().BoolP("debug", "d", false, "Print debug logs")
	rootCmd.SetHelpCommand(&cobra.Command{Hidden: true})
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.AddCommand(cmd.InitCmd())
	rootCmd.AddCommand(cmd.ExportCmd())

	cobra.CheckErr(rootCmd.Execute())
}

func preRun(cmd *cobra.Command, args []string) error {
	isDebug, _ := cmd.Flags().GetBool("debug")
	if isDebug {
		glenv.EnableDebug()
	}
	return nil
}
