# GWTree

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

- [LazyGit](https://github.dev/jesseduffield/lazygit)
