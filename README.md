# inotify-proxy

This tools helps to detect changed files in Docker containers.
If a file is changed from hostsystem a file watcher inside the container detects the change
and triggers a inotify event.

## Purpose

Enables file watcher in a Docker container with a NFS mounted filesystem.

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
      - dir: /tmp/watch1
      - dir: /tmp/watch2
    profile: magento2
    
The profile setting is optional.
The config loading can be skiped by adding the option `-no-config`.    

## Supported Profiles

| Profile        | Allowed file extensions                         |
|----------------|-------------------------------------------------|
| default        | All extensions are allowed                      |
| javascript     | .js .ts                                         |
| magento2       | .css .html .less .sass .js .php .phtml .ts .xml |
| magento2-theme | .css .hs .less .sass .ts                        |
| vue-storefront | .css .js .sass .ts                              |
