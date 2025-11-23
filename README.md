# conf-tree 🌳

> ⚠️ **Work in Progress** - This project is in early development. Features are being actively developed and the API may change.

A terminal-based user interface (TUI) for navigating large JSON configuration files through an intuitive tree view.

![License](https://img.shields.io/badge/license-MIT-blue.svg)
![Go Version](https://img.shields.io/badge/go-1.24.5-blue.svg)

## Overview

**conf-tree** is a command-line tool for viewing large JSON files in a more intuitive way. Instead of manually scrolling through thousands of lines of JSON, conf-tree presents your data in a navigable tree structure with a clean, two-panel interface.

## Current Features

- 📊 **Tree View Navigation** - Browse JSON structures in a tree format
- 🎨 **Dual-Panel Interface** - Left panel shows the tree structure, right panel displays children of the selected node
- ⌨️ **Basic Keyboard Navigation** - Navigate using arrow keys and vim-style keybindings
- 🚀 **Built with Bubble Tea** - Smooth, responsive terminal UI

## Installation

### Prerequisites

- Go 1.24.5 or higher

### Build from Source

```bash
git clone https://github.com/TompaSkitfet/conf-tree.git
cd conf-tree
go build -o ctree ./cmd/ctree
```

### Run

```bash
./ctree
```

> **Note:** Currently, the application looks for a file named `test.json` in the current directory. Support for custom file paths is planned.

## Usage

1. Place your JSON file as `test.json` in the directory where you run the application
2. Launch the application: `./ctree`
3. Navigate through your JSON structure using the keyboard controls

### Key Bindings

| Key | Action |
|-----|--------|
| `↑` / `k` | Move up in the tree |
| `↓` / `j` | Move down in the tree |
| `q` | Quit the application |

## Project Structure

```
conf-tree/
├── cmd/
│   └── ctree/          # Main application entry point
├── internal/
│   ├── config/         # JSON loading and parsing
│   ├── domain/         # Core data structures (Node types)
│   └── ui/             # Bubble Tea UI components
│       ├── components/
│       │   └── tree/   # Tree view component
│       ├── model.go    # Main UI model
│       ├── keys.go     # Keyboard bindings
│       ├── styles.go   # Visual styling
│       └── layout.go   # Layout components
└── README.md
```

## Technologies

- **[Bubble Tea](https://github.com/charmbracelet/bubbletea)** - Terminal UI framework
- **[Bubbles](https://github.com/charmbracelet/bubbles)** - TUI components for Bubble Tea
- **[Lipgloss](https://github.com/charmbracelet/lipgloss)** - Style definitions for terminal layouts

## Roadmap

- [ ] Add file path argument support
- [ ] Implement editing capabilities
- [ ] Add search/filter functionality
- [ ] Support for multiple file formats (YAML, TOML, etc.)
- [ ] Expand/collapse tree nodes
- [ ] Save edited configurations
- [ ] Copy/paste functionality
- [ ] Syntax highlighting for values

## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests.

## License

This project is open source and available under the MIT License.

## Author

Created by [@TompaSkitfet](https://github.com/TompaSkitfet)

