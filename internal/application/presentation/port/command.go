package application

import (
	"context"

	me "github.com/octoposprime/op-be-logging/internal/domain/model/entity"
)

// CommandPort is a port for Hexagonal Architecture Pattern.
// It is used to communicate with the application layer.
type CommandPort interface {
	// Log sends the given log to the application layer.
	Log(ctx context.Context, logData me.LogData)
}
