package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/kiwicom/glenv/internal/glenv"
	"github.com/spf13/cobra"
)

func newExportCmd() *cobra.Command {

	var cmd = &cobra.Command{
		Use:   "export",
		Short: "Export variables from GitLab",
		Long:  "Export variables from GitLab. This command is called e.g. by direnv",
		Run: func(cmd *cobra.Command, args []string) {
			ignoreList, _ := cmd.Flags().GetStringArray("ignore")
			if err := export(ignoreList); err != nil {
				errMsg := fmt.Sprintf("error: %v\n", err)
				os.Stderr.WriteString(errMsg)
			}
		},
	}

	cmd.Flags().StringArrayP("ignore", "i", []string{}, "sometimes you want to ignore some variable")

	return cmd
}

func init() {
	rootCmd.AddCommand(newExportCmd())
}

// main functionality of glenv is here
//
// retrieve all GitLab variables for project and print the vars
// as exports into console std output. The output can be used as entry into
// bash eval:
//
//     eval "$(glenv export)"
//
// example with ignore
//
//     eval "$(glenv export -i SOME_VAR)
//
func export(ignoreList []string) error {
	token := os.Getenv("GITLAB_TOKEN")
	if token == "" {
		return errors.New("Missing GITLAB_TOKEN. Please ensure GITLAB_TOKEN env. variable is present")
	}

	host, project, err := glenv.GetHostAndProject(".")
	if err != nil {
		return err
	}

	vars, err := glenv.GetAllProjectVariables(token, host, project)
	if err != nil {
		return err
	}

	// print exports to output
	// if env. variable is in ignore list, it's ignored
	for key, val := range vars {
		if isNotIgnored(key, ignoreList) {
			fmt.Printf("export %s='%s'\n", key, val)
		}
	}

	return nil
}

func isNotIgnored(value string, ignoreList []string) bool {
	for _, ignored := range ignoreList {
		if value == ignored {
			return false
		}
	}
	return true
}
