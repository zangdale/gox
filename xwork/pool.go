package xwork

import "sync"

/*
	相同类型数据的批量处理任务, 平行任务
*/

const (
	DefaultWorkGoroutines int = 10
)

type DoneErrFunc func() error

// DefultDone 默认执行 task 的方法
var DefultDone = func(fn DoneErrFunc) {
	defer recover()
	_ = fn()
}

// Pool 工作池
type Pool struct {
	work chan WorkerInter
	wg   sync.WaitGroup
	Done func(DoneErrFunc)
}

// NewPool 创建一个新的工作池
// workGoroutines 工作的 gorountine
func NewWorkerPool(workGoroutines int) *Pool {
	if workGoroutines < 1 {
		workGoroutines = DefaultWorkGoroutines
	}

	p := Pool{
		work: make(chan WorkerInter),
		Done: DefultDone,
	}

	p.wg.Add(workGoroutines)
	for i := 0; i < workGoroutines; i++ {
		go func() {
			defer p.wg.Done()
			for w := range p.work {
				p.Done(w.Task)
			}
		}()
	}

	return &p
}

// Add 添加任务
func (p *Pool) Add(w WorkerInter) {
	p.Run(w)
}

// Run 立刻执行
func (p *Pool) Run(w WorkerInter) {
	p.work <- w
}

// Shutdown 关闭工作池并等待所有Worker执行完毕
func (p *Pool) Shutdown() {
	close(p.work)
	p.wg.Wait()
}
