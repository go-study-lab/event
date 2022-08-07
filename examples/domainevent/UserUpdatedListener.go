package domainevent

import (
	"fmt"
	"github.com/go-study-lab/event"
)

type UserUpdatedListener struct {
}

func (*UserUpdatedListener) EventHandler() event.Handler {
	return func(e *event.Event) {
		fmt.Println("UserUpdatedListener executed!")
	}
}

func (*UserUpdatedListener) AsyncProcess() bool {
	return true
}
