package domainevent

import "github.com/go-study-lab/event"

func init() {
	eventDispatcher := event.Dispatcher()
	eventDispatcher.Subscribe(UserCreatedEvent(new(UserCreated)), new(UserCreatedListener))
	eventDispatcher.Subscribe(UserUpdatedEvent(new(UserUpdated)), new(UserUpdatedListener))
}
