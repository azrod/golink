---
hide:
  - toc
---

# Zsh completion

`glctl` supports zsh completion for commands and flags.
If shell completion is not already enabled in your environment you will need to enable it.

You can execute the following once:

``` sh
echo "autoload -U compinit; compinit" >> ~/.zshrc
```

!!! example "To load completions in your current shell session"

    ``` sh
    source <(glctl completion zsh)
    ```

!!! example "To load completions for every new session"

    You will need to start a new shell for this setup to take effect.

    === "Linux"
        ``` sh
        glctl completion zsh > "${fpath[1]}/_glctl"
        ```

    === "macOS"
        ``` sh
        glctl completion zsh > $(brew --prefix)/share/zsh/site-functions/_glctl
        ```
