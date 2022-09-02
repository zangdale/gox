package xsync

import (
	"sync"
	"testing"
	"time"
)

func TestSyncMap(t *testing.T) {
	m := sync.Map{}
	s := time.Now()
	for i := 0; i < 100000; i++ {
		m.Store(i, i)
	}
	for i := 0; i < 100000; i++ {
		_, _ = m.Load(i)
	}
	t.Log(time.Since(s))
}

func BenchmarkSyncMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m := sync.Map{}
		s := time.Now()
		for i := 0; i < 100000; i++ {
			m.Store(i, i)
		}
		for i := 0; i < 100000; i++ {
			_, _ = m.Load(i)
		}

		b.Log(time.Since(s))
	}
}

func BenchmarkMutexMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m := NewMutexMap[int, int]()
		s := time.Now()
		for i := 0; i < 100000; i++ {
			m.Store(i, i)
		}
		for i := 0; i < 100000; i++ {
			_, _ = m.Load(i)
		}
		b.Log(time.Since(s))
	}
}
