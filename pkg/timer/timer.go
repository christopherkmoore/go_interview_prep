package timer

import (
	"fmt"
	"time"
)

type Timer struct {
	Start time.Time

	TimeLogger
}

type TimeLogger interface {
	LogTimeMS() time.Time
}

func New() Timer {
	t := Timer{
		Start: time.Now(),
	}

	return t
}

func (t Timer) LogTimeMS() {
	now := time.Now()

	difference := now.UnixMicro() - t.Start.UnixMicro()

	fmt.Printf("\n Finished execution: %v MS\n", difference)
}
