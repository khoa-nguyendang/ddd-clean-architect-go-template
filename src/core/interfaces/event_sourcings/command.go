package eventsourcings

// Command contains the methods to retreive basic info about it
type Command interface {
	GetType() string
	GetAggregateID() string
	GetAggregateType() string
	IsValid() bool
	GetVersion() int
}
