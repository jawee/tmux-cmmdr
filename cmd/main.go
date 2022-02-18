package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

// create type for project-commands.json
//type ProjectCommands struct {
//    Commands []Command `json:"commands"`
//}

type ProjectCommands struct {
    Projects []Project `json:"projects"`
}

type Project struct {
    Name string `json:"name"`
    Windows []Window `json:"windows"`
}

type Window struct {
    Name string `json:"name"`
    Command string `json:"command"`
}

func main() {
    var nameFlag = flag.String("name", "", "name of the project")
    //parse flag
    flag.Parse()

    name := *nameFlag

    fmt.Printf("Name: %s\n", name)
    pwd, _ := os.Getwd()
    jsonFile, err := ioutil.ReadFile(pwd + "/project-commands.json")
    if err != nil {
        fmt.Println(err)
    }   
    var projectCommands ProjectCommands
    err = json.Unmarshal(jsonFile, &projectCommands)
    if err != nil {
        fmt.Println(err)
    }

    for _, project := range projectCommands.Projects {
        if project.Name == name {
            for _, window := range project.Windows {
                fmt.Printf("%s -> %s: %s\n", project.Name,  window.Name, window.Command)
            }
        }
    }
}
