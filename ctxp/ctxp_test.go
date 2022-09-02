package ctxp

import (
	"context"
	"io/ioutil"
	"os"
	"testing"
)

func TestRead(t *testing.T) {
	f, err := os.Open("../ximage/fonts/ttfs/CEF-CJK.ttf")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	ctx := NewCtxpWithReader(context.TODO(), 23469120, f)
	go func() {
		for {
			t.Log(ctx.Process())
			p, err := ctx.ProcessPercent()
			if err != nil {
				t.Log(err)
				return
			}
			t.Logf("%.4f", p)
			if p == 1 {
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
