---
hide:
  - toc
---

# link

Link refers to the mapping between a golink and a URL. The golink is the key and the URL is the value.

## Usages

### Get link list

!!! info "Namespace"
    If `--namespace | -n` is not specified, the `default` namespace is used.

``` sh
$> glctl get link
NAMESPACE     NAME        PATH         TARGET URL                      STATUS
default       grafana     /grafana     https://grafana.example.com     Enabled
default       prometheus  /prometheus  https://prometheus.example.com  Enabled
```

### Add link

!!! warning "Unique path"
    The path must be unique in the namespace.

``` sh
$> glctl add link demo /demo https://demo.example.com
```

### Delete link

``` sh
$> glctl delete link demo
```

## Aliases

The `li` alias is available for `link`.

``` sh
$> glctl get li
NAMESPACE     NAME        PATH         TARGET URL                      STATUS
default       grafana     /grafana     https://grafana.example.com     Enabled
default       prometheus  /prometheus  https://prometheus.example.com  Enabled
```

## Help

The `--help` flag is available for `link`.

``` sh
$> glctl get|add|delete link -h
```
