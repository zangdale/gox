package xwork

/*
	按顺序执行某些任务, 线性任务
*/
import (
	"fmt"
	"sync"
)

type JobFunc func()

type JobRunType int

const (
	OrderRunJob           JobRunType = 0
	GoroutionWithWGRunJob JobRunType = 1
)

type jobBox struct {
	name       string
	wg         *sync.WaitGroup
	jobRunType JobRunType
	jobs       []JobFunc
}

func NewJobBox(jobRunType JobRunType, name string) *jobBox {
	return &jobBox{
		name:       name,
		jobRunType: jobRunType,
		wg:         &sync.WaitGroup{},
	}
}

func (box *jobBox) Add(fs ...JobFunc) {
	box.jobs = append(box.jobs, fs...)
}

func (box *jobBox) Doing() {
	l := len(box.jobs)
	if l == 0 {
		return
	}
	box.wg.Add(l)

	switch box.jobRunType {
	case OrderRunJob:
		DefaultOrderRunJob(box.wg, box.jobs)
	case GoroutionWithWGRunJob:
		DefaultGoroutionWithWGRunJob(box.wg, box.jobs)
	default:
		DefaultRunJob(box.wg, box.jobs)
	}
	box.wg.Wait()
	return
}

var DefaultOrderRunJob = func(wg *sync.WaitGroup, jobs []JobFunc) {
	for _, v := range jobs {
		defer wg.Done()
		v()
	}
}

var DefaultGoroutionWithWGRunJob = func(wg *sync.WaitGroup, jobs []JobFunc) {
	for _, v := range jobs {
		go func(f JobFunc, w *sync.WaitGroup) {
			defer w.Done()
			f()
		}(v, wg)
	}
}

var DefaultRunJob = func(wg *sync.WaitGroup, jobs []JobFunc) {
	DefaultOrderRunJob(wg, jobs)
}

// -------------------------- 一些任务例子 -------------------------

var Test1_JobFunc = func(name string) JobFunc {
	return func() {
		fmt.Printf("start do %s jobs \n", name)
	}
}

var Test2_JobFunc = func(name string, i int) JobFunc {
	return func() {
		fmt.Printf("start do %s jobs %d \n", name, i)
	}
}
