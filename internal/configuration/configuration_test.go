package configuration

import (
	"bytes"
	"testing"
)

func getTestJson() []byte {
    buffer := bytes.NewBuffer([]byte{})
    buffer.WriteString("{\r\n    \"projects\": [\r\n        {\r\n            \"name\": \"tmux-cmmdr\",\r\n            \"windows\": [\r\n                {\r\n                    \"name\": \"zsh\",\r\n                    \"commands\": [\r\n                        {\"command\": \"zsh\"}\r\n                    ]\r\n                },\r\n                {\r\n                    \"name\": \"run\",\r\n                    \"commands\": [\r\n                        { \"command\": \"go run cmd/main.go\" }\r\n                    ]\r\n                }\r\n            ]\r\n        },\r\n        {\r\n            \"name\": \"tmux-cmmdr-cli\",\r\n            \"windows\": [\r\n                {\r\n                    \"name\": \"run\",\r\n                    \"commands\": [\r\n                        { \"command\": \"go run cmd/main.go\" }\r\n                    ]\r\n                }\r\n            ]\r\n        }\r\n    ]\r\n}\r\n")

    return buffer.Bytes()
}

func TestNewConfiguration(t *testing.T) {
    bytes := getTestJson()
    config, err := New(bytes)

    if err != nil {
        t.Errorf("Error: %s", err)
    }
    if config == nil {
        t.Error("config is nil")
    }

    if config.Projects[0].Name != "tmux-cmmdr" {
        t.Errorf("Expected: tmux-cmmdr, got: %s", config.Projects[0].Name)
    }
}

func TestGetProject(t *testing.T) {
    bytes := getTestJson()
    config, err := New(bytes)

    if err != nil {
        t.Errorf("Error: %s", err)
    }
    if config == nil {
        t.Error("config is nil")
    }

    project := config.GetProject("tmux-cmmdr")

    if project == nil {
        t.Error("project is nil")
    }
    if project.Name != "tmux-cmmdr" {
        t.Errorf("Expected: tmux-cmmdr, got: %s", project.Name)
    }
}

func TestGetProjectDoesNotExist(t *testing.T) {
    bytes := getTestJson()
    config, err := New(bytes)

    if err != nil {
        t.Errorf("Error: %s", err)
    }
    if config == nil {
        t.Error("config is nil")
    }

    project := config.GetProject("does-not-exist")

    if project != nil {
        t.Error("project is not nil")
    }
}
