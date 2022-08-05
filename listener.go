package event

// Handler of Event in Listeners
type Handler func(e *Event)

type IListener interface {
	EventHandler() Handler
	// AsyncProcess
	// return false: Handler func will execute in same goroutine with event trigger func,
	// means event trigger will be blocked until all his synchronized listener's handler are finished
	// return true: Handler will execute in other goroutine, event trigger will not
	// be blocked to wait listener's completion.
	AsyncProcess() bool
}
