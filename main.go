package main

import (
	"fmt"

	"github.com/zangdale/gox/xpath"
)

func main() {
	ss, s, _ := xpath.ExecPath()
	fmt.Println(ss, s)
	fmt.Println(xpath.GetThisDirectory())
}
