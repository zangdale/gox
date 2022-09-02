package xwork

import (
	"fmt"
	"time"
)

// WorkerInter 指定任务的最小单位
type WorkerInter interface {
	Task() error
}

type DefaultWork struct {
	Count int
}

func (w *DefaultWork) Task() error {
	fmt.Printf("DefaultWork Task %d \n", w.Count)
	time.Sleep(time.Second)
	return nil
}
