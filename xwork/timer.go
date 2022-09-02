package xwork

import "time"

/*
 Timer 周期执行任务
*/

// Timer 周期执行任务
type Timer struct {
	*time.Timer
	fn JobFunc
}

func NewTimer(d time.Duration, fn JobFunc) *Timer {
	if fn == nil {
		return nil
	}

	tmer := &Timer{
		Timer: time.NewTimer(d),
		fn:    fn,
	}

	go func(t *Timer) {
		defer recover()
		for range t.Timer.C {
			t.fn()
		}
	}(tmer)

	return tmer
}

func (t *Timer) Stop() bool {
	return t.Timer.Stop()
}
