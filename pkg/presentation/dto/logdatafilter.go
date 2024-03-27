package presentation

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	me "github.com/octoposprime/op-be-logging/internal/domain/model/entity"
	mo "github.com/octoposprime/op-be-logging/internal/domain/model/object"
	pb "github.com/octoposprime/op-be-shared/pkg/proto/pb/logging"
	tuuid "github.com/octoposprime/op-be-shared/tool/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// LogDataFilter is a struct that represents the filter dto of a log.
type LogDataFilter struct {
	//This block has deleted because this is converted for decorator
	/*
		Id              int       `json:"id"`               // Id is the id of the log.
		EventDateFrom   time.Time `json:"event_date_from"`  // EventDate of the log is in the between of EventDateFrom and EventDateTo.
		EventDateTo     time.Time `json:"event_date_to"`    // EventDate of the log is in the between of EventDateFrom and EventDateTo.
		LogType         int       `json:"log_type"`         // LogType is the type of the log.
		SourceType      int       `json:"source_type"`      // SourceType is the type of the source.
		SourceId        int       `json:"source_id"`        // SourceId represents the id of the source that the log is for.
		DestinationType int       `json:"destination_type"` // DestinationType is the type of the destination.
		DestinationId   int       `json:"destination_id"`   // DestinationId represents the id of the destination that the log is for.
	*/
	proto *pb.LogDataFilter
}

// NewLogDataFilter creates a new *LogDataFilter.
func NewLogDataFilter(pb *pb.LogDataFilter) *LogDataFilter {
	return &LogDataFilter{
		proto: pb,
	}
}

// String returns a string representation of the LogDataFilter.
func (s *LogDataFilter) String() string {

	return fmt.Sprintf("Id: %v, "+
		"EventDateFrom: %v, "+
		"EventDateTo: %v, "+
		"LogType: %v, "+
		"ServiceName: %v, "+
		"Path: %v, "+
		"UserId: %v, "+
		"CreatedAtFrom: %v, "+
		"CreatedAtTo: %v, "+
		"UpdatedAtFrom: %v, "+
		"UpdatedAtTo: %v, "+
		"SearchText: %v, "+
		"SortType: %v, "+
		"SortField: %v, "+
		"Limit: %v, "+
		"Offset: %v",
		s.proto.Id,
		s.proto.EventDateFrom,
		s.proto.EventDateTo,
		s.proto.LogType,
		s.proto.ServiceName,
		s.proto.Path,
		s.proto.UserId,
		s.proto.CreatedAtFrom,
		s.proto.CreatedAtTo,
		s.proto.UpdatedAtFrom,
		s.proto.UpdatedAtTo,
		s.proto.SearchText,
		s.proto.SortType,
		s.proto.SortField,
		s.proto.Limit,
		s.proto.Offset)
}

// NewLogDataFilterFromPb creates a new *LogDataFilter from protobuf.
//
// Deprecated: This struct is converted for decorator
// Use NewLogDataFilter instead.
func NewLogDataFilterFromPb(pb *pb.LogDataFilter) *LogDataFilter {
	return &LogDataFilter{
		pb,
	}
}

// NewLogDataFilterFromEntity creates a new *LogDataFilter from entity.
func NewLogDataFilterFromEntity(entity me.LogDataFilter) *LogDataFilter {
	id := entity.Id.String()
	eventDateFrom := timestamppb.New(entity.EventDateFrom)
	eventDateTo := timestamppb.New(entity.EventDateTo)
	logType := pb.LogType(entity.LogType)
	servieName := entity.ServiceName
	path := entity.Path
	userId := entity.UserId
	createdAtFrom := timestamppb.New(entity.CreatedAtFrom)
	createdAtTo := timestamppb.New(entity.CreatedAtTo)
	updatedAtFrom := timestamppb.New(entity.UpdatedAtFrom)
	updatedAtTo := timestamppb.New(entity.UpdatedAtTo)
	searchText := entity.SearchText
	sortType := entity.SortType
	sortField := pb.LogSortField(entity.SortField)
	limit := int32(entity.Limit)
	offset := int32(entity.Offset)
	return &LogDataFilter{
		&pb.LogDataFilter{
			Id:            &id,
			EventDateFrom: eventDateFrom,
			EventDateTo:   eventDateTo,
			LogType:       &logType,
			ServiceName:   &servieName,
			Path:          &path,
			UserId:        &userId,
			CreatedAtFrom: createdAtFrom,
			CreatedAtTo:   createdAtTo,
			UpdatedAtFrom: updatedAtFrom,
			UpdatedAtTo:   updatedAtTo,
			SearchText:    &searchText,
			SortType:      &sortType,
			SortField:     &sortField,
			Limit:         &limit,
			Offset:        &offset,
		},
	}
}

// ToPb returns a protobuf representation of the LogDataFilter.
func (s *LogDataFilter) ToPb() *pb.LogDataFilter {
	return s.proto
}

// ToEntity returns a entity representation of the LogDataFilter.
func (s *LogDataFilter) ToEntity() *me.LogDataFilter {
	id := uuid.UUID{}
	if s.proto.Id != nil {
		id = tuuid.FromString(*s.proto.Id)
	}
	eventDateFrom := time.Time{}
	if s.proto.EventDateFrom != nil {
		eventDateFrom = s.proto.EventDateFrom.AsTime()
	}
	eventDateTo := time.Time{}
	if s.proto.EventDateTo != nil {
		eventDateTo = s.proto.EventDateTo.AsTime()
	}
	logType := 0
	if s.proto.LogType != nil {
		logType = int(*s.proto.LogType)
	}
	serviceName := ""
	if s.proto.ServiceName != nil {
		serviceName = string(*s.proto.ServiceName)
	}
	path := ""
	if s.proto.Path != nil {
		path = string(*s.proto.Path)
	}
	userId := ""
	if s.proto.UserId != nil {
		userId = *s.proto.UserId
	}
	createdAtFrom := time.Time{}
	if s.proto.CreatedAtFrom != nil {
		createdAtFrom = s.proto.CreatedAtFrom.AsTime()
	}
	createdAtTo := time.Time{}
	if s.proto.CreatedAtTo != nil {
		createdAtTo = s.proto.CreatedAtTo.AsTime()
	}
	updatedAtFrom := time.Time{}
	if s.proto.UpdatedAtFrom != nil {
		updatedAtFrom = s.proto.UpdatedAtFrom.AsTime()
	}
	updatedAtTo := time.Time{}
	if s.proto.UpdatedAtTo != nil {
		updatedAtTo = s.proto.UpdatedAtTo.AsTime()
	}
	searchText := ""
	if s.proto.SearchText != nil {
		searchText = string(*s.proto.SearchText)
	}
	sortType := ""
	if s.proto.SortType != nil {
		sortType = string(*s.proto.SortType)
	}
	sortField := 0
	if s.proto.SortField != nil {
		sortField = int(*s.proto.SortField)
	}
	limit := 0
	if s.proto.Limit != nil {
		limit = int(*s.proto.Limit)
	}
	offset := 0
	if s.proto.Offset != nil {
		offset = int(*s.proto.Offset)
	}
	return &me.LogDataFilter{
		Id:            id,
		EventDateFrom: eventDateFrom,
		EventDateTo:   eventDateTo,
		LogType:       mo.LogType(logType),
		ServiceName:   serviceName,
		Path:          path,
		UserId:        userId,
		CreatedAtFrom: createdAtFrom,
		CreatedAtTo:   createdAtTo,
		UpdatedAtFrom: updatedAtFrom,
		UpdatedAtTo:   updatedAtTo,
		SearchText:    searchText,
		SortType:      sortType,
		SortField:     mo.LogSortField(sortField),
		Limit:         limit,
		Offset:        offset,
	}
}
