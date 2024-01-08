---
hide:
  - toc
---

# Namespace

A namespace is a logical grouping of golinks. Golinks are stored in a namespace. A namespace can be used to group golinks by team, department, or any other logical grouping.

## Usages

### Get namespace list

``` sh
$> glctl get namespace
NAME          STATUS      LINKS
default       Enabled     1
myproject     Enabled     4
```

### Add namespace

``` sh
$> glctl add namespace myproject
```

### Delete namespace

``` sh
$> glctl delete namespace myproject
```

!!! warning "Delete namespace does not delete links"
    Deleting a namespace is impossible if there are links in the namespace.
    Use `--force` to delete the namespace and all links in the namespace.

## Aliases

The `ns` alias is available for `namespace`.

``` sh
$> glctl get ns
NAME          STATUS      LINKS
default       Enabled     1
myproject     Enabled     4
```

## Help

The `--help` flag is available for `namespace`.

``` sh
$> glctl get|add|delete namespace -h
```
