# Development Requirements
- golang
    - last tested with 1.22.0
- nodejs with pnpm
    - `npm i -g pnpm`
    - last tested with 20.10.0

# Initial setup
1. Submodule update to get all submods
2. In `er-builds-web`, `pnpm i`

# List of components
The software consists of these components:

- Web ui: built with node/pnpm. Needs `pnpm i`
- erbuilds.exe: Main backend. Built with go.
- builds-downloader.exe: Downloader program. Built with go.

# Development Modes
## Web ui development
Runs web ui build in watch mode and runs server in tmux

1. Be in tmux session
2. `bash run-all-dev.sh`

## Build and run single go program
How to build and run single program. After build, the program will stay in the top dir.

- `bash run-erbuilds-dev.sh`
- `bash run-downloader-dev.sh`

# Generating Release
[how to create release](../release/readme.md)