package domain

import (
	"fmt"
)

// LogBody is a struct that represents the body of a log.
type LogBody struct {
	Message string `json:"message"` // Message is the message of the log.
}

// NewLogBody creates a new *LogBody.
func NewLogBody(message string) *LogBody {
	return &LogBody{
		Message: message,
	}
}

// NewEmptyLogBody creates a new *LogBody with empty values.
func NewEmptyLogBody() *LogBody {
	return &LogBody{
		Message: "",
	}
}

// String returns a string representation of the LogBody.
func (s *LogBody) String() string {
	return fmt.Sprintf("Message: %v",
		s.Message)
}

// Equals returns true if the LogBody is equal to the other LogBody.
func (s *LogBody) Equals(other *LogBody) bool {
	if s.Message != other.Message {
		return false
	}
	return true
}

// Clone returns a clone of the LogBody.
func (s *LogBody) Clone() *LogBody {
	return &LogBody{
		Message: s.Message,
	}
}

// IsEmpty returns true if the LogBody is empty.
func (s *LogBody) IsEmpty() bool {
	if s.Message != "" {
		return false
	}
	return true
}

// IsNotEmpty returns true if the LogBody is not empty.
func (s *LogBody) IsNotEmpty() bool {
	return !s.IsEmpty()
}

// Clear clears the LogBody.
func (s *LogBody) Clear() {
	s.Message = ""
}

// Validate validates the LogBody.
func (s *LogBody) Validate() error {
	if s.IsEmpty() {
		return ErrorLogBodyIsEmpty
	}
	return nil
}
