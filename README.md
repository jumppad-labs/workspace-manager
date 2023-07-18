# Workspace Manager

VSCode extension that lets you manage browser and terminal tabs inside your VSCode workspace.

## Example config

```json
{
  "tabs": [
    {"uri": "file:///c%3A/Users/jacks/code/src/github.com/shipyard-run/save-restore-editors/src/extension.ts", "viewColumn": 1, "focus": true},
    {"uri": "https://www.shipyard.run", "title": "Shipyard"},
    {"uri": "https://www.shipyard.run/docs/install", "title": "Docs"},
    {"uri": "http://localhost:8080", "title": "Consul Thing"}
  ],
  "terminals": [
    {
      "name": "Terminal 1", 
      "command": "docker ps", 
      "viewColumn": 1, 
      "location": "editor", 
      "env": {"HOME": "/root"},
      "cwd": "/root"
    },
    {"name": "Terminal 2", "command": "ls -la", "viewColumn": 1, "location": "panel", "focus": true}
  ],
  "closeTabsOnStart": true,
  "closeTerminalsOnStart": true
}
```