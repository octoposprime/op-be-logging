package application

import (
	"context"

	as "github.com/octoposprime/op-be-logging/internal/application/service"
	me "github.com/octoposprime/op-be-logging/internal/domain/model/entity"
)

// CommandAdapter is an adapter for Hexagonal Architecture Pattern.
// It is used to communicate with the application layer.
type CommandAdapter struct {
	Service *as.Service
}

// NewCommandAdapter creates a new *CommandAdapter.
func NewCommandAdapter(s *as.Service) CommandAdapter {
	return CommandAdapter{
		s,
	}
}

// Log sends the given log to the application layer.
func (a CommandAdapter) Log(ctx context.Context, logData me.LogData) {
	a.Service.Log(ctx, logData)
}
