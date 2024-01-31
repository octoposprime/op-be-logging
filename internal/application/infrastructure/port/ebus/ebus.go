package application

import (
	"context"

	me "github.com/octoposprime/op-be-logging/internal/domain/model/entity"
)

// EBusPort is a port for Hexagonal Architecture Pattern.
// It is used to communicate with the event bus.
type EBusPort interface {
	// Listen listens to the event bus and calls the given callBack function for each received log.
	Listen(ctx context.Context, channelName string, callBack func(channelName string, logData me.LogData))
}
