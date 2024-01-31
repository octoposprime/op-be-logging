package domain

// LogType is a type that represents the type of a log.
type LogType int8

const (
	LogTypeNONE LogType = iota
	LogTypeINFO
	LogTypeWARNING
	LogTypeERROR
	LogTypeDEBUG
)
