# inotify-proxy

This tools helps to detect changed files in Docker containers.
If a file is changed from hostsystem a file watcher inside the container detects the change
and triggers a inotify event.

## Purpose

Enables file watcher in a Docker container with a NFS mounted filesystem.

## Usage Example

    # Magento 2 Profile
    ./inotify-proxy -profile magento2 path/to/your/project
    
    # Change frequence of file checks to 5s (default 2s)
    ./inotify-proxy -sleep 5 path/to/your/project

## Supported Profiles

| Profile        | Allowed file extensions                         |
|----------------|-------------------------------------------------|
| default        | All extensions are allowed                      |
| magento2       | .css .html .less .sass .js .php .phtml .ts .xml |
| magento2-theme | .css .hs .less .sass .ts                        |
| vue-storefront | .css .js .sass ts                               |
