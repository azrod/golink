---
hide:
    - toc
---

# Configuration

All configuration are available in the config file, environment variables, and command line flags.

Is possible to use all of them together, the configuration precedence is the following:

1. Command line flags
2. Environment variables
3. Config file

## Command line flags

All configuration can be set via command line flags.

!!! example "Command line flags"

    The following example shows the **default** command line flags.

    ``` bash
    golink \
        --app.address="localhost" \
        --app.port="8081" \
        --health.address="localhost" \
        --health.port="8082" \
        --storage.type="local" \
        --storage.local.path="./"
    ```

## Environment variables

All configuration can be set via environment variables.

!!! example "Environment variables"

    The following example shows the **default** environment variables.

    ``` bash
    export GOLINK_APP_ADDRESS="localhost"
    export GOLINK_APP_PORT="8081"
    export GOLINK_HEALTH_ADDRESS="localhost"
    export GOLINK_HEALTH_PORT="8082"
    export GOLINK_STORAGE_TYPE="local"
    export GOLINK_STORAGE_LOCAL_PATH="./"
    ```

## Config file

The config file are located in the following paths:

* `/etc/golink/config.[yaml,json,hcl,toml,ini]`
* `./config.[yaml,json,hcl,toml,ini]`

A lot of formats are supported *(JSON, TOML, YAML, HCL, INI)*.

!!! example "config.yaml"

    The following example shows the **default** config file.

    ``` yaml
    app:
      address: localhost
      port: 8081
    health:
      address: localhost
      port: 8082
    storage:
      type: local
      local:
        path: ./
    ```

## Storage configuration

See the [Storage Backend](/storage) section for more advanced configuration.
