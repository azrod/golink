---
hide:
  - toc
---

# Powershell completion

`glctl` supports powershell completion for commands and flags.

!!! example "To load completions in your current shell session"

    ``` sh
    glctl completion powershell | Out-String | Invoke-Expression
    ```

!!! info

    To load completions for every new session, add the output of the above command to your powershell profile.
