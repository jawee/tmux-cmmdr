package cli

import "testing"

func TestFlags(t *testing.T) {
    var tests = []struct {
        args []string
        cliArgs CliArguments
    }{
        {[]string{"-project-name","someproject"},
        CliArguments{ProjectName:"someproject"}},
    }

    for _, test := range tests {
        cliArgs, err := GetArguments("prog", test.args)
        if err != nil {
            t.Errorf("Error parsing arguments: %s", err)
        }
        if *cliArgs != test.cliArgs {
            t.Errorf("Expected %v, got %v", test.cliArgs, cliArgs)
        }
    }
}
