package xwork

import (
	"errors"
	"time"

	"github.com/zangdale/gox/xtime"
)

/*
	指定时间执行任务
*/

var (
	ErrCronTimeStaleDated = errors.New("input time stale dated")
	ErrDoneFuncIsNull     = errors.New("input done func is null")
)

const (
	DefaultCronTimerTime time.Duration = xtime.Second
)

// Cron 某个时间指定某个函数
type Cron struct {
	t  time.Time
	fn JobFunc
	*Timer
}

func NewCron(t time.Time, fn JobFunc) (*Cron, error) {
	if fn == nil {
		return nil, ErrDoneFuncIsNull
	}

	if t.Before(time.Now()) {
		return nil, ErrCronTimeStaleDated
	}

	c := &Cron{t: t, fn: fn}

	tmer := NewTimer(DefaultCronTimerTime, func() {
		if t.Before(time.Now()) {
			fn()
		}
	})

	c.Timer = tmer

	return c, nil
}
