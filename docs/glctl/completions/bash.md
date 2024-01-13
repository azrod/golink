---
hide:
  - toc
---

# Bash completion

`glctl` supports bash completion for commands and flags.
This script depends on the 'bash-completion' package.
If it is not installed already, you can install it via your OS's package manager.

!!! example "To load completions in your current shell session"

    ``` sh
    source <(glctl completion bash)
    ```

!!! example "To load completions for every new session"

    You will need to start a new shell for this setup to take effect.

    === "Linux"
        ``` sh
        glctl completion bash > /etc/bash_completion.d/glctl
        ```

    === "macOS"
        ``` sh
        glctl completion bash > $(brew --prefix)/etc/bash_completion.d/glctl
        ```
