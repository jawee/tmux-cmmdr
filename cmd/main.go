package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"github.com/jawee/tmux-cmmdr/internal/cli"
	"github.com/jawee/tmux-cmmdr/internal/configuration"
)


func getJsonFileBytes() ([]byte, error) {
    pwd, _ := os.Getwd()
    jsonFile, err := ioutil.ReadFile(pwd + "/project-commands.json")
    if err != nil {
        log.Println(err)
        return nil, err
    }
    return jsonFile, nil
}
func main() {
    tmux, err := exec.LookPath("tmux")
    if err != nil {
        log.Printf("tmux not found")
        os.Exit(1)
    }
    //
    // cmd := exec.Command(tmux, "new-window", "-t", "tmux-cmmdr",  "-d", "-n", "dostuff")
    // cmd2 := exec.Command(tmux, "send-keys", "-t", "tmux-cmmdr:dostuff", "echo 'hello world'", "Enter")
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
        log.Println(err)
        os.Exit(1)
    }

    name := cliArgs.ProjectName

    log.Printf("Name: %s\n", name)

    jsonFile, err := getJsonFileBytes()
    if err != nil {
        log.Println(err)
        os.Exit(1)
    }

    projectsConfig, err := configuration.New(jsonFile)
    if err != nil {
        log.Println(err)
        os.Exit(1)
    }

    project := projectsConfig.GetProject(name)

    cmd := exec.Command(tmux, "has-session", "-t", name)

    err = cmd.Run()
    if err != nil {
        cmd := exec.Command(tmux, "new-session", "-d", "-s", name)
        err = cmd.Run()
        if err != nil {
            log.Println(err)
            os.Exit(1)
        }
    }

    log.Println("Session exists")


    for _, window := range project.Windows {
        for _, command := range window.Commands {
            log.Printf("Creating window: %s\n", window.Name)
            cmd := exec.Command(tmux, "new-window", "-t", name, "-n", window.Name, "-d")
            err = cmd.Run()
            if err != nil {
                log.Println(err)
                os.Exit(1)
            }

            log.Printf("Running command '%s' in window '%s\n", window.Name, command)
            cmd = exec.Command(tmux, "send-keys", "-t", name + ":" + window.Name, command.Command, "Enter")
            err = cmd.Run()
            if err != nil {
                log.Println(err)
                os.Exit(1)
            }
        }
    }
}
