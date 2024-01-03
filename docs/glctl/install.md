---
hide:
  - toc
---

# Installation

## Automatic installation

Follow the instructions below to install glctl (GoLink Command Line Tool).

```{ .sh .copy }
curl -sSfL https://raw.githubusercontent.com/azrod/golink/main/scripts/install.sh | sudo sh
```

!!! info "Note"
    On Windows, you can run the above commands with Git Bash, which comes with [Git for Windows](https://git-scm.com/download/win).

This will install glctl to `/usr/local/bin`. If you want to install to a different location use `-b` flag.

```bash
curl -sSfL https://raw.githubusercontent.com/azrod/golink/main/scripts/install.sh | sudo sh -b /usr/bin
```

## Manual installation

Download the latest release from [GitHub Releases](https://github.com/azrod/golink/releases/latest) and extract the binary to a directory in your `PATH`.

## Install from source

Go 1.21 or later is required to install `glctl` from source.

```bash
go install github.com/azrod/golink/cmd/glctl@latest
```
