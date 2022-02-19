package cli

import (
	"bytes"
	"flag"
	"fmt"
)

type CliArguments struct {
    ProjectName string
}

func GetArguments(progname string, args []string) (*CliArguments, error) {
    for _, arg := range args {
        fmt.Println(arg)
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

    return &CliArguments{ProjectName: projectName}, nil
}
