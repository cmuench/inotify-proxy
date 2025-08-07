# inotify-proxy

![Maintenance Badge](https://img.shields.io/maintenance/yes/2025.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/cmuench/inotify-proxy)](https://goreportcard.com/report/github.com/cmuench/inotify-proxy)
[![Go Github Action Workflow](https://github.com/cmuench/inotify-proxy/workflows/Go/badge.svg)](https://github.com/cmuench/inotify-proxy/actions?query=workflow%3AGo)

This tools helps to detect changed files in Docker Containers or in Virtual Machines (e.g. Vagrant).
If a file is changed from host system a file watcher inside the container detects the change
and triggers an inotify event.

The tool is designed to run over a longer period of time. It comes with a garbage collector to cleanup old watched files in memory.

Installation is simple -> It's one binary which can be downloaded and executed.

## Purpose

Enables file watcher in a Docker Container/Virtual Machine with a NFS mounted filesystem.

## Installation

- Download compiled application on release page: https://github.com/cmuench/inotify-proxy/releases
- Extract the zip/tarball.
- Run `./inotify-proxy` binary.

See Wiki page for more informations.
https://github.com/cmuench/inotify-proxy/wiki/Installation

## Usage

    Usage of ./inotify-proxy:
      -no-config
            Do not load config.
      -profile string
            Defines a special profile with extensions to look for. This speeds up the process. Available profiles are 'magento2-theme' (default "default")
      -sleep int
            Cycle time in seconds. Defines time to sleep after each filesystem walk. Default 2s (default 2)

### Examples

    # Magento 2 Profile
    ./inotify-proxy -profile magento2 path/to/your/project
    
    # Change frequence of file checks to 5s (default 2s)
    ./inotify-proxy -sleep 5 path/to/your/project
    
    # Multiple pathes to watch ...
    ./inotify-proxy project/path1 project/path2 

### Config

If the file `inotify-proxy.yaml` exist in the current working directory, it will be applied.

Example config:

    ---
    watch:
    - directory: /tmp/watch1
      profile: magento2

    - directory: /tmp/watch2
      profile: sass

    - directory: /tmp/watch3
      extensions: [.css, .html]

The profile setting is optional.
The config loading can be skipped by adding the option `-no-config`.    

## Supported Profiles

| Profile        | Allowed file extensions                         |
|----------------|-------------------------------------------------|
| default        | All extensions are allowed                      |
| javascript     | .js .ts                                         |
| less           | .less                                           |
| magento2       | .css .html .less .sass .js .php .phtml .ts .xml |
| magento2-theme | .css .hs .less .sass .ts                        |
| sass           | .sass .scss                                     |
| vue-storefront | .css .js .sass .ts                              |

## Comparison with other tools

There are other tools available that solve a similar problem. Here's a comparison to help you choose the right tool for your needs.

### notify-forwarder

[notify-forwarder](https://github.com/mhallin/notify-forwarder) is a tool that forwards file system events from a host to a client. It requires a component running on both the host and the client.

### fs_eventbridge

[fs_eventbridge](https://github.com/TomFrost/fs_eventbridge) is another tool that bridges file system events from a host to a guest VM. It also requires a component running on both the host and the client.

### inotify-proxy (this tool)

`inotify-proxy` is different from the other two tools in that it **does not require a host component**. It's a single binary that runs inside the Docker Container or VM and polls the filesystem for changes. This makes it a simpler solution for use cases where you only need to detect file changes within a container or VM.
