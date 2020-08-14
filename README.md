# inotify-proxy

This tools helps to detect changed files in Docker containers.
If a file is changed from hostsystem a file watcher inside the container detects the change
and triggers a inotify event.

## Purpose

Enables file watcher in a Docker container with a NFS mounted filesystem.
