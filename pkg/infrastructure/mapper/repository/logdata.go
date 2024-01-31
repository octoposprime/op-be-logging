package infrastructure

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	me "github.com/octoposprime/op-be-logging/internal/domain/model/entity"
	mo "github.com/octoposprime/op-be-logging/internal/domain/model/object"
	tgorm "github.com/octoposprime/op-be-shared/tool/gorm"
)

// LogData is a struct that represents the db mapper of a log.
type LogData struct {
	tgorm.Model

	EventDate   time.Time `json:"event_date" gorm:"not null;default:CURRENT_TIMESTAMP"` // EventDate is the date of the event.
	LogType     int       `json:"log_type" gorm:"not null;default:0"`                   // LogType is the type of the log.
	ServiceName string    `json:"service_name" gorm:"not null;default:''"`              // ServiceName is the name of source micro service.
	Path        string    `json:"path" gorm:"not null;default:''"`                      // Path is the name of the source function of the source micro service.
	UserId      string    `json:"user_id" gorm:"not null;default:''"`                   // UserId represents the id of the user that the log is for.
	Message     string    `json:"message" gorm:"not null;default:''"`                   // Message is the message of the log.
}

// NewLogData creates a new *LogData.
func NewLogData(id uuid.UUID,
	eventDate time.Time,
	logType int,
	serviceName string,
	path string,
	userId string,
	message string) *LogData {
	return &LogData{
		Model:       tgorm.Model{ID: id},
		EventDate:   eventDate,
		LogType:     logType,
		ServiceName: serviceName,
		Path:        path,
		UserId:      userId,
		Message:     message,
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
		s.ID,
		s.EventDate,
		s.LogType,
		s.ServiceName,
		s.Path,
		s.UserId,
		s.Message)
}

// NewLogDataFromEntity creates a new *LogData from entity.
func NewLogDataFromEntity(entity *me.LogData) *LogData {
	return &LogData{
		Model:       tgorm.Model{ID: entity.Id},
		EventDate:   entity.LogHeader.EventDate,
		LogType:     int(entity.LogHeader.LogType),
		ServiceName: entity.LogHeader.ServiceName,
		Path:        entity.LogHeader.Path,
		UserId:      entity.LogHeader.UserId,
		Message:     entity.LogBody.Message,
	}
}

// ToEntity returns a entity representation of the LogData.
func (s *LogData) ToEntity() *me.LogData {
	return &me.LogData{
		Id: s.ID,
		LogHeader: mo.LogHeader{
			EventDate:   s.EventDate,
			LogType:     mo.LogType(s.LogType),
			ServiceName: s.ServiceName,
			Path:        s.Path,
			UserId:      s.UserId,
		},
		LogBody: mo.LogBody{
			Message: s.Message,
		},
		CreatedAt: s.CreatedAt,
		UpdatedAt: s.UpdatedAt,
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

// ToEntities creates a new []me.LogData entity.
func (s LogDatas) ToEntities() []me.LogData {
	logDatas := make([]me.LogData, len(s))
	for i, logData := range s {
		logDatas[i] = *logData.ToEntity()
	}
	return logDatas
}
