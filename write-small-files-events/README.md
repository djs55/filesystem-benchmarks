# write-small-files with concurrent events

A simple filesystem benchmark which creates a directory and then creates lots
(default 1000) of small (default 1 KiB) files, while simultaneously changing
different files on the host, generating cache invalidations and inotify injections.

## Usage:

```
go run main.go
```

The output is a single time in seconds.
