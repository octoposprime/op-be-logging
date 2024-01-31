package domain

import (
	"fmt"
	"time"

	smodel "github.com/octoposprime/op-be-shared/pkg/model"
)

// LogHeader is a struct that represents the header of a log.
type LogHeader struct {
	EventDate   time.Time `json:"event_date"`                              // EventDate is the date of the event.
	LogType     LogType   `json:"log_type"`                                // LogType is the type of the log.
	ServiceName string    `json:"service_name" gorm:"not null;default:''"` // ServiceName is the name of source micro service.
	Path        string    `json:"path" gorm:"not null;default:''"`         // Path is the name of the source function of the source micro service.
	UserId      string    `json:"user_id" gorm:"not null;default:''"`      // UserId represents the id of the user that the log is for.
}

// NewLogHeader creates a new *LogHeader.
func NewLogHeader(eventDate time.Time, logType LogType, serviceName string, path string, userId string) *LogHeader {
	return &LogHeader{
		EventDate:   eventDate,
		LogType:     logType,
		ServiceName: serviceName,
		Path:        path,
		UserId:      userId,
	}
}

// NewEmptyLogHeader creates a new *LogHeader with empty values.
func NewEmptyLogHeader() *LogHeader {
	return &LogHeader{
		EventDate:   time.Time{},
		LogType:     LogTypeNONE,
		ServiceName: smodel.ServiceNone,
		Path:        "",
		UserId:      "",
	}
}

// String returns a string representation of the LogHeader.
func (s *LogHeader) String() string {
	return fmt.Sprintf("EventDate: %v, "+
		"LogType: %v, "+
		"ServiceName: %v, "+
		"Path: %v, "+
		"UserId: %v",
		s.EventDate,
		s.LogType,
		s.ServiceName,
		s.Path,
		s.UserId)
}

// Equals returns true if the LogHeader is equal to the other LogHeader.
func (s *LogHeader) Equals(other *LogHeader) bool {
	if s.EventDate != other.EventDate {
		return false
	}
	if s.LogType != other.LogType {
		return false
	}
	if s.ServiceName != other.ServiceName {
		return false
	}
	if s.Path != other.Path {
		return false
	}
	if s.UserId != other.UserId {
		return false
	}
	return true
}

// Clone returns a clone of the LogHeader.
func (s *LogHeader) Clone() *LogHeader {
	return &LogHeader{
		EventDate:   s.EventDate,
		LogType:     s.LogType,
		ServiceName: s.ServiceName,
		Path:        s.Path,
		UserId:      s.UserId,
	}
}

// IsEmpty returns true if the LogHeader is empty.
func (s *LogHeader) IsEmpty() bool {
	if s.EventDate.IsZero() {
		return false
	}
	if s.LogType != LogTypeNONE {
		return false
	}
	if s.ServiceName != "" {
		return false
	}
	if s.Path != "" {
		return false
	}
	if s.UserId != "" {
		return false
	}
	return true
}

// IsNotEmpty returns true if the LogHeader is not empty.
func (s *LogHeader) IsNotEmpty() bool {
	return !s.IsEmpty()
}

// Clear clears the LogHeader.
func (s *LogHeader) Clear() {
	s.EventDate = time.Time{}
	s.LogType = LogTypeNONE
	s.ServiceName = ""
	s.Path = ""
	s.UserId = ""
}

// Validate validates the LogHeader.
func (s *LogHeader) Validate() error {
	if s.IsEmpty() {
		return ErrorLogHeaderIsEmpty
	}
	return nil
}
