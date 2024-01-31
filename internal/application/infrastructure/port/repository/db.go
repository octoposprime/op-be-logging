package application

import (
	"context"

	me "github.com/octoposprime/op-be-logging/internal/domain/model/entity"
)

// DbPort is a port for Hexagonal Architecture Pattern.
// It is used to communicate with the database.
type DbPort interface {
	// GetLogsByFilter returns the logs that match the given filter.
	GetLogsByFilter(ctx context.Context, logDataFilter me.LogDataFilter) ([]me.LogData, error)

	// Log inserts a new log into the database.
	Log(ctx context.Context, logData me.LogData)
}
