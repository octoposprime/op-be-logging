package presentation

import (
	"context"

	mo "github.com/octoposprime/op-be-logging/internal/domain/model/object"
	pb_error "github.com/octoposprime/op-be-shared/pkg/proto/pb/error"
)

// GetErrors returns the embedded errors.
func (a *Grpc) GetErrors(ctx context.Context, req *pb_error.ErrorRequest) (*pb_error.Errors, error) {
	var results pb_error.Errors
	for _, err := range mo.ERRORS {
		if err != nil {
			results.Errors = append(results.Errors, &pb_error.Error{
				Error: err.Error(),
			})
		}
	}
	return &results, nil
}
