## tmux-cmmdr

Creates tmux windows and sends commands to windows on the provided project tmux session.

### TODO

- [x] Basic tmux functionality
- [ ] Initialize configuration file
- [ ] Add to configuration file CLI
- [ ] Remove from configuration file CLI

### Building From Source
```
git clone git@github.com:jawee/tmux-cmmdr.git
cd tmux-cmmdr

# Install it where you want
go build -o ~/.local/bin/tmux-cmmdr ./cmd
```

### Usage

```bash
tmux-cmmdr -project-name <projectname>
```


Example
```bash
tmux-cmmdr -project-name tmux-cmmdr
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
