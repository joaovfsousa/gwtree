package main

import (
	"os"

	cmd "github.com/joaovfsousa/gwtree/cmd/gwtree"
)

func main() {
	err := cmd.RootExecute()
	if err != nil {
		panic(err)
	}
	os.Exit(0)
}
