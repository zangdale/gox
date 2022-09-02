package xsync

import (
	"sync"
	"testing"
)

func TestOnce_Do(t *testing.T) {
	o := NewMore(2)
	f := func() {
		t.Log("once done")
	}
	t.Log(o.Do(f) == true)
	t.Log(o.Do(f) == true)
	t.Log(o.Do(f) == false)
}

/*
	使用 sync 的数组 进行控制执行多次
*/

func TestOnce_SomeDo(t *testing.T) {
	os := make([]*sync.Once, 10)
	for i := 0; i < 10; i++ {
		os[i] = &sync.Once{}
	}

	f := func() {
			t.Log( "true")
	}
	for i := 0; i < 12; i++ {
		for _, o := range os {
			o.Do(f)
		}
	}
}


func TestOnce_Dos(t *testing.T) {
	o := NewMore(10)
	f := func() {
		t.Log("once done")
	}
	for i := 0; i < 15; i++ {
		t.Log(o.Do(f))
	}
}