package application

import (
	"context"

	ip_ebus "github.com/octoposprime/op-be-logging/internal/application/infrastructure/port/ebus"
	ip_repo "github.com/octoposprime/op-be-logging/internal/application/infrastructure/port/repository"
	me "github.com/octoposprime/op-be-logging/internal/domain/model/entity"
	ds "github.com/octoposprime/op-be-logging/internal/domain/service"
	smodel "github.com/octoposprime/op-be-shared/pkg/model"
)

// Service is an application service.
// It manages the business logic of the application.
type Service struct {
	*ds.Service
	ip_repo.DbPort
	ip_ebus.EBusPort
}

// NewService creates a new *Service.
func NewService(domainService *ds.Service, dbRepository ip_repo.DbPort, eBus ip_ebus.EBusPort) *Service {
	service := &Service{
		domainService,
		dbRepository,
		eBus,
	}
	service.EventListen()
	return service
}

// This is the event listener handler of the application layer.
func (a *Service) EventListen() *Service {
	go a.Listen(context.Background(), smodel.ChannelLogging, a.EventListenerCallBack)
	return a
}

// This is a call-back function of the event listener handler of the application layer.
func (a *Service) EventListenerCallBack(channelName string, logData me.LogData) {
	a.Log(context.Background(), logData)
}

// Log sends the given log to the repository of the infrastructure layer.
func (a *Service) Log(ctx context.Context, logData me.LogData) {
	a.DbPort.Log(ctx, logData)
}

// GetLogsByFilter returns the logs that match the given filter.
func (a *Service) GetLogsByFilter(ctx context.Context, logDataFilter me.LogDataFilter) ([]me.LogData, error) {
	return a.DbPort.GetLogsByFilter(ctx, logDataFilter)
}
