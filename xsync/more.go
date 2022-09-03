package xsync

/*
 对 sync. Do 的扩充,
- 支持返回是否 本次执行
- 支持设置总共执行的次数

方法1: TestOnce_Do
方法2: TestOnce_SomeDo

用途:
	1. 初始化次数限制
	2. 限制数据更新次数
	3. 限制每次运行使用的次数

*/

import (
	"sync"
	"sync/atomic"
)

type More struct {
	doneCount uint32
	count     uint32
	sync.Mutex
}

// NewMore 执行几次 , count 为 0 和 count =1 相同
func NewMore(count uint32) *More {
	if count == 0 {
		count = 1
	}
	return &More{
		count: count,
	}
}

func (o *More) Do(f func()) bool {

	if atomic.LoadUint32(&o.doneCount) < atomic.LoadUint32(&o.count) {
		o.doSlow(f)
		return true
	}

	return false
}

func (o *More) doSlow(f func()) {
	o.Lock()
	defer o.Unlock()
	if o.doneCount < o.count {
		defer atomic.StoreUint32(&o.doneCount, o.doneCount+1)
		f()
	}
}
