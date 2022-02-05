package proxy

import (
	"fmt"
	"testing"
)

func TestProxy_Do(t *testing.T) {
	// var s Subject
	var p Subject
	// s = &Real{}
	p = &Proxy{}

	fmt.Println(p.Do())
}
