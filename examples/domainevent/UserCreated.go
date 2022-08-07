package domainevent

import "github.com/go-study-lab/event"

type UserCreated struct {
	UserId   int
	UserName string
	Tx       *Tx
}

type Tx struct {
	// Mock as a DB transaction
	DSN string
	Tid int64
}

func (*UserCreated) EventName() string {
	return "UserCreated"
}

func UserCreatedEvent(e *UserCreated) *event.Event {
	return event.NewEvent(e)
}
