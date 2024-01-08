---
hide:
  - toc
---

# Concept

`glctl` uses the concept and logic of `kubectl` to manage golinks. The namespace is the basic unit of management. The namespace is a logical grouping of golinks. Golinks are stored in a namespace. A namespace can be used to group golinks by team, department, or any other logical grouping.

!!! example "Basic usage"

    === "get"

        This command is used to get the namespace list.

        ``` sh
        $> glctl get namespace
        NAME          STATUS      LINKS
        default       Enabled     1
        myproject     Enabled     4
        ```

    === "add"

        This command is used to add a link to a namespace.

        ``` sh
        $> glctl add link -n myproject demo /demo https://demo.example.com
        ```

    === "delete"

        This command is used to delete a link from a namespace.

        ``` sh
        $> glctl delete link -n myproject demo
        ```
