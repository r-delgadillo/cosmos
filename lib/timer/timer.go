package timer

import (
	"fmt"
	"time"
)

type StopWatch struct {
	startTime time.Time
	endTime   time.Time
	duration  time.Duration
}

func (t *StopWatch) Start() time.Time {
	t.startTime = time.Now()
	return t.startTime
}

func (t *StopWatch) Stop() time.Time {
	t.endTime = time.Now()
	return t.endTime
}

func (t *StopWatch) latency() time.Duration {
	t.duration = t.endTime.Sub(t.startTime)
	return t.duration
}

func (t *StopWatch) PrintResults() {
	t.latency()
	fmt.Printf("Duration time: %s\n", t.duration)
}
