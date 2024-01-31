package domain

// LogSortField is a type that represents the sort fields of a log data.
type LogSortField int8

const (
	LogSortFieldNONE LogSortField = iota
	LogSortFieldId
	LogSortFieldServiceName
	LogSortFieldEventDate
	LogSortFieldCreatedAt
	LogSortFieldUpdatedAt
)
