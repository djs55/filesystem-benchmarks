# write-small-files

A simple filesystem benchmark which creates a directory and then creates lots
(default 1000) of small (default 1 KiB) files.

## Usage:

```
docker run -v /Users/djs/workspace:/volume djs55/write-small-files
```

The output is a single time in seconds.
