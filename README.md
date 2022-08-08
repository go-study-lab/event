# event
A lightweight event dispatcher for go project to implement event driven programming 


## installation
```shell
go get -u github.com/go-study-lab/event
```

## Dispatcher Methods

### Load event dispatcher singleton
```
event.Dispacthcer()
```
It will return a singleton of `*event.eventDispatcher`, which is a core component and all functionality would supply through it by.

### Bind event listener
```shell
func (dispatcher *eventDispatcher) Subscribe(e *Event, l IListener)
```
`Subscribe` method can specify an event listener for a specific event.
An event listener must implement `IListener` interface, we will explain this interface in next section. 

### Publish event
```shell
func (dispatcher *eventDispatcher) Dispatch(e *Event)
```
Dispatcher's `Dispatch` method would let us push an event, then the dispatcher would execute event's bound listeners. 
### Verify event has listeners
```shell
func (dispatcher *eventDispatcher) HasEventListener(e *Event) bool
```

### Remove event listener
```shell
func (dispatcher *eventDispatcher) RemoveEventListener(e *Event, l IListener)
```

## Event
Type `Event` is an event wrapper for concrete event data
```go
type Event struct {
    eventName     string
    eventTime     time.Time
    concreteEvent IEvent
    eventType     string
}
```
A concrete event must implements interface `IEvent`
```go
type IEvent interface {
    EventName() string
}
```

## Event Methods

### NewEvent

```go
func NewEvent(concreteEvent IEvent) *Event
```
This is a factory method of Event instance.

### OccurredOn
```go
func (e *Event) OccurredOn() time.Time
```
Return event occurred time.

### EventData
```go
func (e *Event) EventData() IEvent
```
Return concrete event data.

### EventName

```go
func (e *Event) EventName() string
```
Return name of event.

## IListener

## Usage Examples
