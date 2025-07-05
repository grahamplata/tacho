# Tacho

A gauge cluster project built with [ebiten](https://ebitengine.org/) 🦐  

## Getting Started

### Build & Run

```sh
go build -o tacho
./tacho
```

#### Optional Flags

- `-config <path>`: Specify a custom config file path
- `-config-type <type>`: Specify config type (default: json)

### Project Structure

```bash
# tree
.
├── main.go                # Entry point
├── pkg/
│   ├── config/            # Config loading/saving
│   ├── scene/             # Scene controller
│   └── state/             # Game state management
├── scenes/                # Example scenes (loading, home, etc.)
└── assets/                # Asset loading (fonts, images, etc.)
```
