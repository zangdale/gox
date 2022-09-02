package ctxp

import (
	"context"
	"errors"
	"io"
	"sync/atomic"
)

var (
	ErrNoAllProcess = errors.New("no all process")
)

type ContextProcess struct {
	context.Context
	process    int64
	allProcess float64
	afterFunc  func(process int64, allProcess float64)
}

func NewContextProcess(ctx context.Context, allProcess float64) *ContextProcess {
	return &ContextProcess{Context: ctx, allProcess: allProcess}
}

func (ctx *ContextProcess) WithAfterFunc(f func(process int64, allProcess float64)) *ContextProcess {
	ctx.afterFunc = f
	return ctx
}

func (ctx *ContextProcess) Process() int64 {
	return atomic.LoadInt64(&ctx.process)
}

func (ctx *ContextProcess) ProcessPercent() (float64, error) {
	if ctx.allProcess < 0 {
		return 0, ErrNoAllProcess
	}
	p := float64(atomic.LoadInt64(&ctx.process))
	return p / ctx.allProcess, nil
}

func (ctx *ContextProcess) ZeroProcess() {
	atomic.SwapInt64(&ctx.process, 0)
}

func (ctx *ContextProcess) AddInt64(p int64) {
	newP := atomic.AddInt64(&ctx.process, p)
	if ctx.afterFunc != nil {
		ctx.afterFunc(newP, ctx.allProcess)
	}
}

type contextWithReaderProcess struct {
	*ContextProcess
	reader io.Reader
}

func NewCtxpWithReader(ctx context.Context, allProcess float64, reader io.Reader) *contextWithReaderProcess {
	return &contextWithReaderProcess{ContextProcess: NewContextProcess(ctx, allProcess), reader: reader}
}

func (ctx *contextWithReaderProcess) Read(p []byte) (n int, err error) {
	select {
	case <-ctx.Done():
		return 0, ctx.Err()
	default:
	}
	n, err = ctx.reader.Read(p)
	if err != nil {
		return 0, err
	}
	if n != 0 {
		ctx.AddInt64(+int64(n))
	}
	return n, nil
}

type contextWithWriterProcess struct {
	*ContextProcess
	writer io.Writer
}

func NewCtxpWithWriter(ctx context.Context, allProcess float64, writer io.Writer) *contextWithWriterProcess {
	return &contextWithWriterProcess{ContextProcess: NewContextProcess(ctx, allProcess), writer: writer}
}

func (ctx *contextWithWriterProcess) Write(p []byte) (n int, err error) {
	select {
	case <-ctx.Done():
		return 0, ctx.Err()
	default:
	}
	n, err = ctx.writer.Write(p)
	if err != nil {
		return 0, err
	}
	if n != 0 {
		ctx.AddInt64(+int64(n))
	}

	return n, nil
}
