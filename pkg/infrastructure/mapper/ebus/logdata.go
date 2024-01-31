package infrastructure

import (
	"fmt"

	me "github.com/octoposprime/op-be-logging/internal/domain/model/entity"
	mo "github.com/octoposprime/op-be-logging/internal/domain/model/object"
	pb "github.com/octoposprime/op-be-shared/pkg/proto/pb/logging"
	tuuid "github.com/octoposprime/op-be-shared/tool/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// LogData is a struct that represents the ebus mapper of a log.
type LogData struct {
	//This block has deleted because this is converted for decorator
	/*
		Id              int       `json:"id"`               // Id is the id of the log.
		LogType         int       `json:"log_type"`         // LogType is the type of the log.
		SourceType      int       `json:"source_type"`      // SourceType is the type of the source.
		SourceId        int       `json:"source_id"`        // SourceId represents the id of the source that the log is for.
		DestinationType int       `json:"destination_type"` // DestinationType is the type of the destination.
		DestinationId   int       `json:"destination_id"`   // DestinationId represents the id of the destination that the log is for.
		EventDate       time.Time `json:"event_date"`       // EventDate is the date of the event.
		Message         string    `json:"message"`          // Message is the message of the log.
	*/
	proto *pb.LogData
}

// NewLogData creates a new *LogData.
func NewLogData(pb *pb.LogData) *LogData {
	return &LogData{
		proto: pb,
	}
}

// String returns a string representation of the LogData.
func (s *LogData) String() string {
	return fmt.Sprintf("Id: %v, "+
		"EventDate: %v, "+
		"LogType: %v, "+
		"ServiceName: %v, "+
		"Path: %v, "+
		"UserId: %v, "+
		"Message: %v",
		s.proto.Id,
		s.proto.Header.EventDate,
		s.proto.Header.LogType,
		s.proto.Header.ServiceName,
		s.proto.Header.Path,
		s.proto.Header.UserId,
		s.proto.Body.Message)
}

// NewLogDataFromPb creates a new *LogData from protobuf.
//
// Deprecated: This struct is converted for decorator
// Use NewLogData instead.
func NewLogDataFromPb(pb *pb.LogData) *LogData {
	return &LogData{
		pb,
	}
}

// NewLogDataFromEntity creates a new *LogData from entity.
func NewLogDataFromEntity(entity *me.LogData) *LogData {
	return &LogData{
		&pb.LogData{
			Id: entity.Id.String(),
			Header: &pb.LogHeader{
				EventDate:   timestamppb.New(entity.EventDate),
				LogType:     pb.LogType(entity.LogType),
				ServiceName: entity.ServiceName,
				Path:        entity.Path,
				UserId:      entity.UserId,
			},
			Body: &pb.LogBody{
				Message: entity.Message,
			},
		},
	}
}

// ToPb returns a protobuf representation of the LogData.
func (s *LogData) ToPb() *pb.LogData {
	return s.proto
}

// ToEntity returns a entity representation of the LogData.
func (s *LogData) ToEntity() *me.LogData {
	return &me.LogData{
		Id: tuuid.FromString(s.proto.Id),
		LogHeader: mo.LogHeader{
			EventDate:   s.proto.Header.EventDate.AsTime(),
			LogType:     mo.LogType(s.proto.Header.LogType),
			ServiceName: s.proto.Header.ServiceName,
			Path:        s.proto.Header.Path,
			UserId:      s.proto.Header.UserId,
		},
		LogBody: mo.LogBody{
			Message: s.proto.Body.Message,
		},
	}
}

type LogDatas []*LogData

// NewLogDatasFromEntities creates a new []*LogData from entities.
func NewLogDataFromEntities(entities []me.LogData) LogDatas {
	logDatas := make([]*LogData, len(entities))
	for i, entity := range entities {
		logDatas[i] = NewLogDataFromEntity(&entity)
	}
	return logDatas
}

// ToPbs returns a protobuf representation of the LogDatas.
func (s LogDatas) ToPbs() *pb.LogDatas {
	logDatas := make([]*pb.LogData, len(s))
	for i, logData := range s {
		logDatas[i] = logData.ToPb()
	}
	return &pb.LogDatas{
		LogDatas: logDatas,
	}
}
