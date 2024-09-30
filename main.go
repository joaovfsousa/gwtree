package main

import (
	"os"

	cmd "github.com/joaovfsousa/gwtree/cmd/gwtree"
	"github.com/joaovfsousa/gwtree/internal/logger"
)

func main() {
	err := cmd.RootExecute()
	if err != nil {
		logger := logger.GetLogger()

		logger.Error(err.Error())
	}
	os.Exit(0)
}
