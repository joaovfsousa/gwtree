package main

import (
	"os"

	cmd "github.com/joaovfsousa/gwtree/cmd/gwtree"
	logger "github.com/joaovfsousa/gwtree/internal"
)

func main() {
	err := cmd.RootExecute()
	if err != nil {
		logger := logger.GetLogger()

		logger.Error(err.Error())
	}
	os.Exit(0)
}
