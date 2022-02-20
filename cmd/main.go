package main

import (
	"commandizizer/internal/cli"
	"commandizizer/internal/configuration"
	"fmt"
	"os"
)


func main() {
    // tmux, err := exec.LookPath("tmux")
    // if err != nil {
    //     fmt.Printf("tmux not found")
    //     os.Exit(1)
    // }
    //
    // cmd := exec.Command(tmux, "new-window", "-t", "commandizizer",  "-d", "-n", "dostuff")
    // cmd2 := exec.Command(tmux, "send-keys", "-t", "commandizizer:dostuff", "echo 'hello world'", "Enter")
    // err = cmd.Run()
    // if err != nil {
    //     fmt.Println(err)
    // }
    // err = cmd2.Run()
    // if err != nil {
    //     fmt.Println(err)
    // }

    cliArgs, err := cli.New(os.Args[0], os.Args[1:])
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    name := cliArgs.ProjectName

    fmt.Printf("Name: %s\n", name)

    projectCommands := configuration.GetProjectsConfig()

    for _, project := range projectCommands.Projects {
        if project.Name == name {
            for _, window := range project.Windows {
                for _, command := range window.Commands {
                    fmt.Printf("%s\n", command)
                }
            }
        }
    }
}
