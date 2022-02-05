package factory_method

import (
	"fmt"
	"testing"
)

func TestOperator(t *testing.T){
	var fac OperatorFactory
	fac = &PlusOperatorFactory{}
	plus := fac.Create()
	plus.SetA("90")
	plus.SetB("1000")
	fmt.Println(plus.Result())
}
