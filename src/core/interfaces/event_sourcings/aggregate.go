package eventsourcings

import "github.com/khoa-nguyendang/ddd-clean-architect-go-template/core/models/events"

// BaseAggregate contains the basic info
// that all aggregates should have
type BaseAggregate struct {
	ID      string
	Type    string
	Version int
	Changes []events.Event
}
