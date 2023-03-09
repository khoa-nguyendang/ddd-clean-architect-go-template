package eventsourcings

// Register defines generic methods to create a registry
type Register interface {
	Set(source interface{})
	Get(name string) (interface{}, error)
	Count() int
}

// EventTypeRegister defines the register for all the events that are Data field child of event struct
type EventTypeRegister interface {
	Register
	Events() []string
}
