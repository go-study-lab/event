package main

import (
	"fmt"
	"github.com/go-study-lab/event"
	"github.com/go-study-lab/event/examples/domainevent"
)

type User struct {
	UserId   int
	UserName string
}

func createUser() {
	user := &User{
		UserId:   1,
		UserName: "KK",
	}
	tx := &domainevent.Tx{DSN: "localhost:3306/test_db", Tid: 1}
	fmt.Printf("Tx in envet trigger: %p \n", tx)

	// assume use tx to create user, then publish a userCreated event
	event.Dispatcher().Dispatch(event.NewEvent(&domainevent.UserCreated{
		UserId:   user.UserId,
		UserName: user.UserName,
		Tx:       tx, // in event handler, we can use this tx to execute other db operation
		// in the same transaction , this method will ensure event trigger and listener as
		// an atomically processed unit in system.
	}))
}

func HasEventListener() {
	e := new(domainevent.UserUpdated)
	ok := event.Dispatcher().HasEventListener(domainevent.UserUpdatedEvent(e))
	fmt.Printf("Event: %s have listener? %t \n", e.EventName(), ok)

	if ok {
		event.Dispatcher().Dispatch(event.NewEvent(new(domainevent.UserUpdated)))
	}
}

func RemoveEventListener() {
	listener := new(domainevent.UserUpdatedListener)

	event.Dispatcher().Subscribe(
		domainevent.UserUpdatedEvent(new(domainevent.UserUpdated)),
		listener,
	)

	event.Dispatcher().RemoveEventListener(
		domainevent.UserUpdatedEvent(new(domainevent.UserUpdated)),
		listener,
	)

}

func main() {
	createUser()
	HasEventListener()
	RemoveEventListener()
}
