package application

import (
	"context"

	me "github.com/octoposprime/op-be-logging/internal/domain/model/entity"
)

// QueryPort is a port for Hexagonal Architecture Pattern.
// It is used to communicate with the application layer.
type QueryPort interface {
	// GetLogsByFilter returns the logs that match the given filter.
	GetLogsByFilter(ctx context.Context, logDataFilter me.LogDataFilter) ([]me.LogData, error)
}
