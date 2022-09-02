package xtools

import (
	"testing"
)

func TestIF(t *testing.T) {
	t.Log("Hello BuGuai !!! ")

	t.Log(IF(true, "true", "false").(string))

	t.Log(IF(false, "true", "false").(string))

	t.Log(IF(1 > 2, 1, 2).(int))

	t.Log(IFFunc(func(a, b string) func() bool {
		return func() bool {
			return len(a) > len(b)
		}
	}("aaa", "bb"), true, false).(bool))
}
