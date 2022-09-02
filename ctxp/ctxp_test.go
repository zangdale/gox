package ctxp

import (
	"context"
	"io/ioutil"
	"os"
	"testing"
)

func TestRead(t *testing.T) {
	f, err := os.Open("xxx.log")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	ctx := NewCtxpWithReader(context.TODO(), 55782932, f)
	ctx = ctx.WithAfterFunc(nil)
	go func() {
		for {
			t.Log(ctx.Process())
			p, err := ctx.ProcessPercent()
			if err != nil {
				t.Log(err)
				return
			}
			t.Log(p)
			if p == 0 {
				return
			}
		}
	}()
	body, err := ioutil.ReadAll(ctx)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("---->", len(body))
}
