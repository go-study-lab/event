package domainevent

import "github.com/go-study-lab/event"

type UserUpdated struct {
	UserId   int
	UserName string
}

func (*UserUpdated) EventName() string {
	return "UserUpdated"
}

func UserUpdatedEvent(e *UserUpdated) *event.Event {
	return event.NewEvent(e)
}
