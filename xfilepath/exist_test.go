package xfilepath

import (
	"testing"
)

func TestIfNotExists(t *testing.T) {
	t.Log("Hello BuGuai !!! ")

	showN := func(fPath string) error {
		t.Log("no show:", fPath)
		return nil
	}
	show := func(fPath string) error {
		t.Log("show:", fPath)
		return nil
	}

	err := IfNotExists("./log/log.go", showN, show)
	if err != nil {
		t.Fatal(err)
	}

	err = IfNotExists("./log/logx.go", showN, show)
	if err != nil {
		t.Fatal(err)
	}

	err = IfNotExists("log/log.go", showN, show)
	if err != nil {
		t.Fatal(err)
	}

	err = IfNotExists("", show, nil)
	if err != nil {
		t.Fatal(err)
	}
}
