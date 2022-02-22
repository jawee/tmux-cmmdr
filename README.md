# tmux-cmmdr

Creates tmux windows and sends commands to windows on the provided project tmux session.

## Usage

```bash
tmux-cmmdr -project-name <projectname>
```


Example
```bash
tmux-cmmdr - project-name tmux-cmmdr
```

Example configuration
```json
{
    "projects": [
        {
            "name": "tmux-cmmdr",
            "windows": [
                {
                    "name": "echo",
                    "commands": [
                        {"command": "echo $PATH"}
                    ]
                }, 
                {
                    "name": "run",
                    "commands": [
                        { "command": "go run cmd/main.go" }
                    ]
                }
            ]
        },
        {
            "name": "some-other-project",
            "windows": [
                {
                    "name": "run",
                    "commands": [
                        { "command": "go run cmd/main.go" }
                    ]
                }
            ]
        }
    ]
}
```
