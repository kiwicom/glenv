package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/kiwicom/glenv/internal/glenv"
	"github.com/spf13/cobra"
)

// exportCmd represents the export command
var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export variables from GitLab",
	Long:  "Export variables from GitLab. This command is called e.g. by direnv",
	Run: func(cmd *cobra.Command, args []string) {
		if err := export(); err != nil {
			errMsg := fmt.Sprintf("error: %v\n", err)
			os.Stderr.WriteString(errMsg)
		}
	},
}

func init() {
	rootCmd.AddCommand(exportCmd)
}

// main functionality of glenv is here
//
// retrieve all GitLab variables for project and print the vars
// as exports into console std output. The output can be used as entry into
// bash eval:
//
//     eval "$(glenv export)"
//
func export() error {
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
	for key, val := range vars {
		fmt.Printf("export %s='%s'\n", key, val)
	}

	return nil
}
