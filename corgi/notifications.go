package corgi

import (
	"github.com/wailsapp/wails"
	"time"
)

func PushError(rt *wails.Runtime, title, message string) {
	if rt == nil {
		waitAndPush := func() {
			for rt == nil {
				time.Sleep(time.Millisecond * 50)
			}
			rt.Events.Emit("errorEvent", title, message)
		}
		go waitAndPush()
	} else {
		rt.Events.Emit("errorEvent", title, message)
	}
}
