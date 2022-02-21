package configuration

import (
	"encoding/json"
	"log"
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

func New(jsonFile []byte) (*ProjectsConfig, error) {
    return getProjectsConfig(jsonFile)
}

func (projectsConfig *ProjectsConfig) GetProject(name string) *Project {
    for _, project := range projectsConfig.Projects {
        if project.Name == name {
            return &project
        }
    }
    return nil
}

func getProjectsConfig(jsonFile []byte) (*ProjectsConfig, error) {
    var projectsConfig ProjectsConfig
    err := json.Unmarshal(jsonFile, &projectsConfig)
    if err != nil {
        log.Println("Could not unmarshal json file", err)
        return nil, err
    }
    return &projectsConfig, nil
}
