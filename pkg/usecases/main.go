package usecases

import (
	"log/slog"

	"github.com/joaovfsousa/gwtree/pkg/git_commands"
)

type UseCases struct {
	logger *slog.Logger
	gc     *git_commands.GitCommander
}

func CreateUseCases(
	logger *slog.Logger,
	gc *git_commands.GitCommander,
) *UseCases {
	return &UseCases{
		logger, gc,
	}
}
