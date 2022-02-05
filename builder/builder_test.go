package builder

import (
	"fmt"
	"testing"
)

func TestOpe(t *testing.T){
	bu := &type1{
		res:"hhh",
	}
	director := NewDirector(bu)
	director.construct()
	fmt.Println(bu.GetRes())
}
