package domain

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	mo "github.com/octoposprime/op-be-logging/internal/domain/model/object"
)

// LogData is a struct that represents the entity of a log.
type LogData struct {
	Id           uuid.UUID `json:"id"` // Id is the id of the log.
	mo.LogHeader           // LogHeader is the header of the log.
	mo.LogBody             // LogBody is the body of the log.

	// Only for view
	CreatedAt time.Time `json:"created_at"` // CreatedAt is the create time.
	UpdatedAt time.Time `json:"updated_at"` // UpdatedAt is the update time.
}

// NewLogData creates a new *LogData.
func NewLogData(id uuid.UUID,
	logHeader mo.LogHeader,
	logBody mo.LogBody) *LogData {
	return &LogData{
		Id:        id,
		LogHeader: logHeader,
		LogBody:   logBody,
	}
}

// NewEmptyLogData creates a new *LogData with empty values.
func NewEmptyLogData() *LogData {
	return &LogData{
		Id:        uuid.UUID{},
		LogHeader: *mo.NewEmptyLogHeader(),
		LogBody:   *mo.NewEmptyLogBody(),
	}
}

// String returns a string representation of the LogData.
func (s *LogData) String() string {
	return fmt.Sprintf("Id: %v, "+
		"LogHeader: %v, "+
		"LogBody: %v",
		s.Id,
		s.LogHeader,
		s.LogBody)
}

// Equals returns true if the LogData is equal to the other LogData.
func (s *LogData) Equals(other *LogData) bool {
	if s.Id != other.Id {
		return false
	}
	if !s.LogHeader.Equals(&other.LogHeader) {
		return false
	}
	if !s.LogBody.Equals(&other.LogBody) {
		return false
	}
	return true
}

// Clone returns a clone of the LogData.
func (s *LogData) Clone() *LogData {
	return &LogData{
		Id:        s.Id,
		LogHeader: *s.LogHeader.Clone(),
		LogBody:   *s.LogBody.Clone(),
	}
}

// IsEmpty returns true if the LogData is empty.
func (s *LogData) IsEmpty() bool {
	if s.Id.String() != "" && s.Id != (uuid.UUID{}) {
		return false
	}
	if !s.LogHeader.IsEmpty() {
		return false
	}
	if !s.LogBody.IsEmpty() {
		return false
	}
	return true
}

// IsNotEmpty returns true if the LogData is not empty.
func (s *LogData) IsNotEmpty() bool {
	return !s.IsEmpty()
}

// Clear clears the LogData.
func (s *LogData) Clear() {
	s.Id = uuid.UUID{}
	s.LogHeader.Clear()
	s.LogBody.Clear()
}

// Validate validates the LogData.
func (s *LogData) Validate() error {
	if s.IsEmpty() {
		return mo.ErrorLogDataIsEmpty
	}
	return nil
}
