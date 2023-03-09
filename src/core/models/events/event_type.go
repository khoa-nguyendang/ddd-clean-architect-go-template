package events

import "sync"

// EventType implements the EventyTypeRegister interface
type EventType struct {
	sync.RWMutex
}
