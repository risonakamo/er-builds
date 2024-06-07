# Development Requirements
- golang
    - last tested with 1.22.3
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
- nica-downloader.exe: Downloader program for Nica builds. Built with go.
- lang-downloader: Downloader program to refresh lang file. Should be used by developer only.

# Development Modes
## Web ui development
Runs web ui build in watch mode and runs server in tmux

1. Be in tmux session
2. `bash run-all-dev.sh`

## Build and run single go program
How to build and run single program. After build, the program will stay in the top dir.

- `bash erbuilds-dev.sh run`
- `bash downloader-dev.sh run`

If exclude `run`, it will only build without running

## Chars yml file
The `chars.yml` is used to specify what to download. It can be regenerated with all the current chars/weapons by using `run-char-yml-gen.sh`. Make sure to comment out all the lines after it is generated.

# Generating Release
[how to create release](../release/readme.md)

# Langfile
The langfile is a resource committed to the repo as it is needed to fill in info during some processes such as nica downloading. This file probably needs to be kept up to date, but since it needs a eternal return api key to do it, has to be done by developer.

To update the langfile:

1. Make copy of `config/dev-config-ex.yml`. Remove the `-ex`.
2. Edit the config and paste in your api key
3. `bash scripts/langfile-download-dev.sh run`

# Characters Yml File
This file can be refreshed with all the characters using the char-yaml-gen binary.

Go to `bin/char-yaml-gen` and `go run` the file. This should be committed to the repo so users don't need to ever run this.