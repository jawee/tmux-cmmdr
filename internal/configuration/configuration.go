package configuration

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type ProjectsConfig struct {
    Projects []Project `json:"projects"`
}

type Project struct {
    Name string `json:"name"`
    Windows []Window `json:"windows"`
}

type Window struct {
    Name string `json:"name"`
    Commands []WindowCommand `json:"commands"`
}

type WindowCommand struct {
    Name string `json:"name"`
    Command string `json:"command"`
}

func GetProjectsConfig() *ProjectsConfig {
    pwd, _ := os.Getwd()
    jsonFile, err := ioutil.ReadFile(pwd + "/project-commands.json")
    if err != nil {
        fmt.Println(err)
    }   
    var projectsConfig ProjectsConfig
    err = json.Unmarshal(jsonFile, &projectsConfig)
    if err != nil {
        fmt.Println(err)
    }
    return &projectsConfig
}
