---
hide:
  - toc
---

# How to use

```bash
glctl is a CLI for golink. It allows you to manage golink from the command line.

Usage:
  glctl [command]

GoLink Commands
  add         Add commands
  delete      Delete commands
  get         Get commands

Other Commands
  version     Returns the version of the application

Additional Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command

Flags:
      --config string      config file (default is $HOME/.golink/config.yaml)
      --debug              debug mode
  -h, --help               help for glctl
      --host string        golink host (default "http://localhost:8081")
  -n, --namespace string   namespace (default "default")
  -o, --output string      output format (default "short")
      --timeout int        timeout in seconds (default 10)

Use "glctl [command] --help" for more information about a command.
```

## How to configure

`glctl` uses a configuration file looking for `$HOME/.golink/config.yaml` by default.
You can specify a different configuration file with the `--config` flag.

If you don't have a configuration file, it will be created automatically when you run the `glctl` command for the first time.

**Default configuration file:**

```{ .yaml .copy }
debug: false
host: http://localhost:8081
namespace: default
timeout: 10
```
