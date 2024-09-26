# fschecker-cli

### fschecker-cli is a beginner-friendly Go CLI project that watches a directory for file changes and automatically backs up data when files are created, modified, or deleted. This project leverages Go's powerful concurrency and filesystem capabilities, making it a perfect introduction to building efficient CLI tools in Go.
Features

Watch a specified directory for file changes:
  1. File creation
  2. File modification
  3. File deletion
     
Automatically trigger a backup process when changes occur
Built using Go and fsnotify

## Installation

Ensure you have Go 1.23 or later installed.

Clone the repository:

    git clone https://github.com/mehulambastha/fschecker-cli.git

Build the CLI:

    cd fschecker-cli
    go build -o fschecker .

Usage

To start watching a directory:

bash

    ./fschecker /path/to/directory

The CLI will actively monitor the directory and report file events in real-time.
Under Development
Implementing backup functionality, such as automatic Git commits and pushes, according to user-defined rules
Future plans for integrating with cloud storage and other backup solutions
Enhancements for error handling and reporting

Dependencies

  fsnotify: A Go library for filesystem notifications that allows the CLI to track changes in real-time.

Future Plans

Some of the potential future enhancements:

  Git integration for tracking file changes, with automatic commits and pushes
    Support for watching multiple directories simultaneously
    Configurable backup paths and advanced backup scheduling options
