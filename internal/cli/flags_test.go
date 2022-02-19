package cli

import (
    "testing"
)

type TestError struct {
}

func (e *TestError) Error() string {
    return "test error"
}

func TestNoProjectNameFlag(t *testing.T) {
    _, err := GetArguments("prog", []string{})

    if err == nil {
        t.Errorf("Expected error, got nil")
    }
}

func TestNoProjectName(t *testing.T) {
    cliArgs, err := GetArguments("prog", []string{"-project-name"})

    if cliArgs != nil {
        t.Errorf("Expected nil, got %v", cliArgs)
    }

    if err == nil {
        t.Errorf("Expected error, got nil")
    }
}

func TestEmptyProjectName(t *testing.T) {
    cliArgs, err := GetArguments("prog", []string{"-project-name", ""})

    if cliArgs != nil {
        t.Errorf("Expected nil, got %v", cliArgs)
    }

    if err == nil {
        t.Errorf("Expected error, got nil")
    }
}

func TestProjectNameRequired(t *testing.T) {
    var tests = []struct {
        args []string
        cliArgs CliArguments
        err error
    }{
        {
            []string{"-project-name","someproject"},
            CliArguments{ProjectName:"someproject"},
            nil,
        },
        //        {
        //            []string{},
        //            CliArguments{},
        //            new(TestError),
        //        },
        //        {
        //            []string{"-project-name"},
        //            CliArguments{},
        //            new(TestError),
        //        },
    }

    for _, test := range tests {
        cliArgs, err := GetArguments("prog", test.args)
        if test.err == nil {
            if err != nil {
                t.Errorf("Error parsing arguments: %s", err)
            }

            if *cliArgs != test.cliArgs {
                t.Errorf("Expected %v, got %v", test.cliArgs, cliArgs)
            }
        } else {
            if test.err != nil && err == nil {
                t.Errorf("Expected error %s, got nil", test.err)
            }
        }
    }
}
