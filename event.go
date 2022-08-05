package event

import (
	"reflect"
	"time"
)

type IEvent interface {
	EventName() string
}

// Event : Type of domainevent
type Event struct {
	eventName     string
	eventTime     time.Time
	concreteEvent interface{}
	eventType     string
}

// NewEvent : Factory of Event obj
func NewEvent(concreteEvent IEvent) *Event {
	event := &Event{
		eventName:     concreteEvent.EventName(),
		eventTime:     time.Now(),
		concreteEvent: concreteEvent,
		// not use eventName to avoid key collision in event holder
		eventType: reflect.TypeOf(concreteEvent).String(),
	}

	return event
}

func (e *Event) EventName() string {
	return e.eventName
}

func (e *Event) OccurredOn() time.Time {
	return e.eventTime
}

func (e *Event) EventData() interface{} {
	return e.concreteEvent
}

var _Dispatcher *eventDispatcher

func init() {
	_Dispatcher = NewDispatcher()
}

func Dispatcher() *eventDispatcher {
	return _Dispatcher
}
