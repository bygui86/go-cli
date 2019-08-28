
# Go CLI sample project

## Info

### Precedence

The precedence for flag value sources is as follows (highest to lowest):
1. Command line flag value from user
2. Environment variable from user (if specified)
3. Configuration file (if specified) `[FilePath]`
4. Environment variable (if specified) `[EnvVar]`
5. Default defined on the flag `[Value]`

---

## Autocompletion

Source the `autocomplete/*_autocomplete` file in your `.bashrc | .zshrc` file while setting the `PROG` variable to the name of your program.

`example`
```
go build examples/bash-completion.go
PROG=bash-completion source autocomplete-scripts/zsh_autocomplete
./bash-completion
```

`go-cli`
```
go build .
source <(./go-cli autocompletion zsh)
./go-cli

# OR

go build .
PROG=go-cli source autocomplete-scripts/zsh_autocomplete
./go-cli
```

---

## Links

* https://github.com/urfave/cli

### Collaterals

* https://flaviocopes.com/go-list-files/
