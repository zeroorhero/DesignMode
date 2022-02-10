package decorator

import (
	"fmt"
	"testing"
)

func TestAddComponent_Cal(t *testing.T) {
	var c Component = &ConcretComponent{}
	c = WrapAddComponent(c, 10)
	c = WrapMulComponent(c, 8)
	fmt.Println(c.Cal())
}
