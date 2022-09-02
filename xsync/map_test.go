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

// BenchmarkSyncMap
// BenchmarkSyncMap-8   	      19	  54256131 ns/op	13683535 B/op	  503489 allocs/op

// BenchmarkMutexMap
// BenchmarkMutexMap-8   	      68	  18442004 ns/op	 5760979 B/op	    4002 allocs/op

func BenchmarkSyncMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m := sync.Map{}
		// s := time.Now()
		for i := 0; i < 100000; i++ {
			m.Store(i, i)
		}
		for i := 0; i < 100000; i++ {
			_, _ = m.Load(i)
		}

		// b.Log(time.Since(s))
	}
}

func BenchmarkMutexMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m := NewMutexMap[int, int]()
		// s := time.Now()
		for i := 0; i < 100000; i++ {
			m.Store(i, i)
		}
		for i := 0; i < 100000; i++ {
			_, _ = m.Load(i)
		}
		// b.Log(time.Since(s))
	}
}

/*
BenchmarkSyncMapR
BenchmarkSyncMapR-8   	     150	   7800146 ns/op	   91350 B/op	    3357 allocs/op

BenchmarkMutexMapR
BenchmarkMutexMapR-8   	     278	   4213131 ns/op	   20756 B/op	      14 allocs/op
*/
func BenchmarkSyncMapR(b *testing.B) {
	m := sync.Map{}

	for i := 0; i < 100000; i++ {
		m.Store(i, i)
	}
	for i := 0; i < b.N; i++ {
		// s := time.Now()
		for i := 0; i < 100000; i++ {
			_, _ = m.Load(i)
		}

		// b.Log(time.Since(s))
	}
}

func BenchmarkMutexMapR(b *testing.B) {
	m := NewMutexMap[int, int]()
	for i := 0; i < 100000; i++ {
		m.Store(i, i)
	}
	for i := 0; i < b.N; i++ {
		// s := time.Now()
		for i := 0; i < 100000; i++ {
			_, _ = m.Load(i)
		}
		// b.Log(time.Since(s))
	}
}
