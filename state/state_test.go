package state

import "testing"

func TestContext_NexDay(t *testing.T) {
	context := NewContext()
	funcs := func() {
		context.Today()
		context.NexDay()
	}
	for i:=0; i<2; i++{
		funcs()
	}
}
