package usecases

import (
	"log/slog"

	"git.n-hub.ru/neosy/npulse-watcher/port/persistence"
)

// Usecases represents the business layer of the application.
type Usecases struct {
	//Link *link.LinkUsecase
}

type DepRepositories struct {
	PulseState persistence.PulseStateRepository
}

// Dependencies contains external dependencies required by the Usecases.
type Dependencies struct {
	Repositories DepRepositories
}

// New returns a new instance of Usecases.
func New(
	logger *slog.Logger,
	deps *Dependencies,
) *Usecases {

	return &Usecases{
		//Link: linkUsecase,
	}
}
