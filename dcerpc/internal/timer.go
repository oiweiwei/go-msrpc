package internal

import "time"

type Timer struct {
	timer *time.Timer
}

func NewTimer() *Timer {
	return &Timer{}
}

func (t *Timer) After(dur time.Duration) <-chan time.Time {
	if t.timer == nil {
		t.timer = time.NewTimer(dur)
	} else {
		if !t.timer.Stop() {
			<-t.timer.C
		}
		t.timer.Reset(dur)
	}
	return t.timer.C
}

func (t *Timer) Clear() {
	t.timer.Stop()
	t.timer = nil
}

func (t *Timer) Stop() {
	if t.timer != nil && !t.timer.Stop() {
		<-t.timer.C
	}
}
