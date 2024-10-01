# GWTree

## Build instructions

### MacOS

`go build && mv gwtree /usr/local/bin`

## To enable worktree switching, add this to your shell rc file(e.g. `~/.bashrc`)

```
function gwt {
    export GWT_NEW_DIR_FILE=~/.gwt/newdir

    gwtree "$@"

    if [ -f $GWT_NEW_DIR_FILE ]; then
      cd "$(cat $GWT_NEW_DIR_FILE)"
      rm -f $GWT_NEW_DIR_FILE > /dev/null
    fi
}
```

### References

- [LazyGit](https://github.com/jesseduffield/lazygit)

### Thirdparty licenses

- [LazyGit](https://github.com/jesseduffield/lazygit/blob/d11e11d179ec7df1b14a536a3965b254430b0504/LICENSE)
