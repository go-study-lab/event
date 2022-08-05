package domainevent

import (
	"fmt"
	"github.com/go-study-lab/event"
)

type UserCreatedListener struct {
}

func (*UserCreatedListener) EventHandler() event.Handler {
	return func(e *event.Event) {
		var eData *UserCreated
		var ok bool

		if eData, ok = e.EventData().(*UserCreated); !ok {
			fmt.Println("can not convert event data to type *UserCreated")
			return
		}

		fmt.Printf("event data, userId:%d, userName:%s, Tx:%p \n", eData.UserId, eData.UserName, eData.Tx)
	}
}

// When we want to use a same
// DB transaction in event trigger and listener to  get ACID assurance, then
// AsyncProcess must return false.
func (*UserCreatedListener) AsyncProcess() bool {
	return false
}
