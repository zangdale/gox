package xwork

import "testing"

func TestDefaultWork(t *testing.T) {
	p := NewWorkerPool(10)

	for i := 0; i < 100; i++ {
		p.Add(&DefaultWork{Count: i})
	}

	p.Shutdown()
}
