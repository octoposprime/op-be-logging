package application

import (
	"context"

	as "github.com/octoposprime/op-be-logging/internal/application/service"
	me "github.com/octoposprime/op-be-logging/internal/domain/model/entity"
)

// QueryAdapter is an adapter for Hexagonal Architecture Pattern.
// It is used to communicate with the application layer.
type QueryAdapter struct {
	Service *as.Service
}

// NewQueryAdapter creates a new *QueryAdapter.
func NewQueryAdapter(s *as.Service) QueryAdapter {
	return QueryAdapter{
		s,
	}
}

// GetLogsByFilter returns the logs that match the given filter.
func (a QueryAdapter) GetLogsByFilter(ctx context.Context, logDataFilter me.LogDataFilter) ([]me.LogData, error) {
	return a.Service.GetLogsByFilter(ctx, logDataFilter)
}
