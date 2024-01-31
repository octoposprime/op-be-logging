package domain

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	mo "github.com/octoposprime/op-be-logging/internal/domain/model/object"
)

// LogDataFilter is a struct that represents the filter of a log.
type LogDataFilter struct {
	Id            uuid.UUID  `json:"id"`                                      // Id is the id of the log.
	EventDateFrom time.Time  `json:"event_date_from"`                         // EventDate of the log is in the between of EventDateFrom and EventDateTo.
	EventDateTo   time.Time  `json:"event_date_to"`                           // EventDate of the log is in the between of EventDateFrom and EventDateTo.
	LogType       mo.LogType `json:"log_type"`                                // LogType is the type of the log.
	ServiceName   string     `json:"service_name" gorm:"not null;default:''"` // ServiceName is the name of source micro service.
	Path          string     `json:"path" gorm:"not null;default:''"`         // Path is the name of the source function of the source micro service.
	UserId        string     `json:"user_id" gorm:"not null;default:''"`      // UserId represents the id of the user that the log is for.

	CreatedAtFrom time.Time `json:"created_at_from"` // CreatedAt is in the between of CreatedAtFrom and CreatedAtTo.
	CreatedAtTo   time.Time `json:"created_at_to"`   // CreatedAt is in the between of CreatedAtFrom and CreatedAtTo.
	UpdatedAtFrom time.Time `json:"updated_at_from"` // UpdatedAt is in the between of UpdatedAtFrom and UpdatedAtTo.
	UpdatedAtTo   time.Time `json:"updated_at_to"`   // UpdatedAt is in the between of UpdatedAtFrom and UpdatedAtTo.

	SearchText string          `json:"search_text"` // SearchText is the full-text search value.
	SortType   string          `json:"sort_type"`   // SortType is the sorting type (ASC,DESC).
	SortField  mo.LogSortField `json:"sort_field"`  // SortField is the sorting field of the logData.

	Limit  int `json:"limit"`  // Limit provides to limitation row size.
	Offset int `json:"offset"` // Offset provides a starting row number of the limitation.
}

// NewLogDataFilter creates a new *LogDataFilter.
func NewLogDataFilter(id uuid.UUID,
	eventDateFrom time.Time,
	eventDateTo time.Time,
	logType mo.LogType,
	serviceName string,
	path string,
	userId string,
	createdAtFrom time.Time,
	createdAtTo time.Time,
	updatedAtFrom time.Time,
	updatedAtTo time.Time,
	searchText string,
	sortType string,
	sortField mo.LogSortField,
	limit int,
	offset int) *LogDataFilter {
	return &LogDataFilter{
		Id:            id,
		EventDateFrom: eventDateFrom,
		EventDateTo:   eventDateTo,
		LogType:       logType,
		ServiceName:   serviceName,
		Path:          path,
		UserId:        userId,
		CreatedAtFrom: createdAtFrom,
		CreatedAtTo:   createdAtTo,
		UpdatedAtFrom: updatedAtFrom,
		UpdatedAtTo:   updatedAtTo,
		SearchText:    searchText,
		SortType:      sortType,
		SortField:     sortField,
		Limit:         limit,
		Offset:        offset,
	}
}

// NewEmptyLogDataFilter creates a new *LogDataFilter with empty values.
func NewEmptyLogDataFilter() *LogDataFilter {
	return &LogDataFilter{
		Id:            uuid.UUID{},
		EventDateFrom: time.Time{},
		EventDateTo:   time.Time{},
		LogType:       mo.LogTypeNONE,
		ServiceName:   "",
		Path:          "",
		UserId:        "",
		CreatedAtFrom: time.Time{},
		CreatedAtTo:   time.Time{},
		UpdatedAtFrom: time.Time{},
		UpdatedAtTo:   time.Time{},
		SearchText:    "",
		SortType:      "",
		SortField:     mo.LogSortFieldNONE,
		Limit:         0,
		Offset:        0,
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
		s.Id,
		s.EventDateFrom,
		s.EventDateTo,
		s.LogType,
		s.ServiceName,
		s.Path,
		s.UserId,
		s.CreatedAtFrom,
		s.CreatedAtTo,
		s.UpdatedAtFrom,
		s.UpdatedAtTo,
		s.SearchText,
		s.SortType,
		s.SortField,
		s.Limit,
		s.Offset)
}
