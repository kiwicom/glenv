package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/kiwicom/glenv/internal/glenv"
)

func main() {
	var err error

	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "init":
			err = initialize()
		default:
			printHelp()
		}
	} else {
		err = export()
	}

	exit(err)
}

// main functionality of glenv is here
//
// retrieve all GitLab variables for project and print the vars
// as exports into console std output. The output can be used as entry into
// bash eval:
//
//     eval "$(glenv)"
//
func export() error {
	token := os.Getenv("GITLAB_TOKEN")
	if token == "" {

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

// this function is executed when you call `glenv init`
// and create `.envrc` file for you in current dir.
func initialize() error {
	body := []byte("eval \"$(glenv)\"\n")
	err := ioutil.WriteFile(".envrc", body, 0644)
	if err != nil {
		return err
	}
	return nil
}

// print some help to std. output
//
// Because the tool is very simple, I don't want to mess it with some
// heavy CLI framework (e.g. cobra) right now. Maybe later
func printHelp() {
	fmt.Println("glenv - export the env. variables from GitLab for current repository")
	fmt.Println("")
	fmt.Println("   subcommands:")
	fmt.Println("       init - create '.envrc' file in current directory")
	fmt.Println("")
}

func exit(err error) {
	if err != nil {
		errMsg := fmt.Sprintf("error: %v\n", err)
		os.Stderr.WriteString(errMsg)
		os.Exit(1)
	} else {
		os.Exit(0)
	}
}
