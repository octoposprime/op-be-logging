package presentation

import (
	"context"

	dto "github.com/octoposprime/op-be-logging/pkg/presentation/dto"
	pb_logging "github.com/octoposprime/op-be-shared/pkg/proto/pb/logging"
)

// Log sends the given log to the application layer.
func (a *Grpc) Log(ctx context.Context, logData *pb_logging.LogData) (*pb_logging.LoggingResult, error) {
	go a.commandHandler.Log(ctx, *dto.NewLogData(logData).ToEntity())
	return &pb_logging.LoggingResult{}, nil
}

// GetLogsByFilter returns the logs that match the given filter.
func (a *Grpc) GetLogsByFilter(ctx context.Context, filter *pb_logging.LogDataFilter) (*pb_logging.LogDatas, error) {
	logDatas, err := a.queryHandler.GetLogsByFilter(ctx, *dto.NewLogDataFilterFromPb(filter).ToEntity())
	return dto.NewLogDataFromEntities(logDatas).ToPbs(), err
}
