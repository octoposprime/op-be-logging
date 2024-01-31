package domain

import me "github.com/octoposprime/op-be-logging/internal/domain/model/entity"

// This is the service layer of the domain layer.
type Service struct {
}

// NewService creates a new *Service.
func NewService() *Service {
	return &Service{}
}

// ValidateLogData validates the logData.
func (s *Service) ValidateLogData(logData *me.LogData) error {
	return logData.Validate()
}
