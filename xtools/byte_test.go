package xtools

import "testing"

// Test .
func TestBuytenull(t *testing.T) {
	a := "" // request.get("key")
	t.Log([]byte(a) == nil == false)
	t.Log(StringToBytes(a) == nil == true)
}
