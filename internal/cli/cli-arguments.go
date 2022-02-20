package cli

import (
	"bytes"
	"flag"
	"fmt"
)

type CliArguments struct {
    ProjectName string
}


func New(program string, args []string) (*CliArguments, error) {
    return getArguments(program, args)
}

func getArguments(progname string, args []string) (*CliArguments, error) {

    if len(args) == 0 {
        return nil, fmt.Errorf("No arguments provided")
    }

    flags := flag.NewFlagSet(progname, flag.ContinueOnError)
    var buf bytes.Buffer
    flags.SetOutput(&buf)

    var projectName string
    flags.StringVar(&projectName, "project-name", "", "Project name")

    err := flags.Parse(args)
    if err != nil {
        return nil, err
    }

    if projectName == "" {
        return nil, fmt.Errorf("No project name provided")
    }

    return &CliArguments{ProjectName: projectName}, nil
}
