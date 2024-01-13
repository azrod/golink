---
hide:
  - toc
---

# Fish completion

`glctl` supports fish completion for commands and flags.

!!! example "To load completions in your current shell session"

    ``` sh
    glctl completion fish | source
    ```

!!! example "To load completions for every new session"

    You will need to start a new shell for this setup to take effect.

    ``` sh
    glctl completion fish > ~/.config/fish/completions/glctl.fish
    ```
