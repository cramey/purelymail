## Purelymail Command Line Tool
This tool uses the Purelymail API to allow you to control your Purelymail
account. Currently it supports management of users, routing, and domains.

## Configuration
Please see `purelymail.sample.yaml` for configuration examples. The
configuration file is searched for in `$HOME/.config/purelymail.yaml` on Linux,
`$HOME/Library/Application Support/purelymail.yaml` on MacOS, and 
`%APPDATA%/purelymail.yaml` on Windows.

## Development
Building is as simple as `go build` in the project directory.
