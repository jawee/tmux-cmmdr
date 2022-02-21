package main

import (
	"commandizizer/internal/cli"
	"commandizizer/internal/configuration"
	"io/ioutil"
	"log"
	"os"
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

    for _, window := range project.Windows {
        for _, command := range window.Commands {
            log.Printf("%s\n", command)
        }
    }
}
