// Package xpath provides ...
package xpath

import (
	"fmt"
	"testing"
)

// Test .
func TestGetCurrentDirectory(t *testing.T) {
	ss, s, _ := ExecPath()
	fmt.Println(ss, s)

	fmt.Println(GetThisDirectory())
}
