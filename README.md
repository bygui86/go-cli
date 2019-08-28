
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

Source the `autocomplete-scripts/*_autocomplete` file in your `.bashrc | .zshrc` file while setting the `PROG` variable to the name of your program.

`example`
```
go build examples/bash-completion.go
PROG=bash-completion source autocomplete-scripts/zsh_autocomplete
./bash-completion
# now play with tab
```

`go-cli`
```
go build .
source <(./go-cli autocompletion zsh)
./go-cli
# now play with tab

# OR

go build .
PROG=go-cli source autocomplete-scripts/zsh_autocomplete
./go-cli
# now play with tab
```

---

## Links

* https://github.com/urfave/cli

### Config

* https://medium.com/@felipedutratine/manage-config-in-golang-to-get-variables-from-file-and-env-variables-33d876887152
* https://github.com/Tkanos/gonfig
* https://github.com/spf13/viper

### Collaterals

* https://flaviocopes.com/go-list-files/
