package simple_factory

import (
	"fmt"
	"testing"
)

func TestType1(t *testing.T){
	api := NewAPI("1")
	fmt.Println(api.Say())
}
