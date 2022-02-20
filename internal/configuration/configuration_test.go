package configuration

import (
	"bytes"
	"testing"
)

func getTestJson() []byte {
    buffer := bytes.NewBuffer([]byte{})
    buffer.WriteString("{\r\n    \"projects\": [\r\n        {\r\n            \"name\": \"commandizizer\",\r\n            \"windows\": [\r\n                {\r\n                    \"name\": \"zsh\",\r\n                    \"commands\": [\r\n                        {\"command\": \"zsh\"}\r\n                    ]\r\n                },\r\n                {\r\n                    \"name\": \"run\",\r\n                    \"commands\": [\r\n                        { \"command\": \"go run cmd/main.go\" }\r\n                    ]\r\n                }\r\n            ]\r\n        },\r\n        {\r\n            \"name\": \"commandizizer-cli\",\r\n            \"windows\": [\r\n                {\r\n                    \"name\": \"run\",\r\n                    \"commands\": [\r\n                        { \"command\": \"go run cmd/main.go\" }\r\n                    ]\r\n                }\r\n            ]\r\n        }\r\n    ]\r\n}\r\n")

    return buffer.Bytes()
}

func TestGetConfiguration(t *testing.T) {
    bytes := getTestJson()
    config, err := GetProjectsConfig(bytes)

    if err != nil {
        t.Errorf("Error: %s", err)
    }
    if config == nil {
        t.Error("config is nil")
    }

    if config.Projects[0].Name != "commandizizer" {
        t.Errorf("Expected: commandizizer, got: %s", config.Projects[0].Name)
    }
}

