package simple_factory

import "fmt"

type coffee interface {
	addSuger()
	getname()string
}

type BaseCoffee struct {}
func (b *BaseCoffee) addSuger(){
	fmt.Println("添加糖")
}
func (b *BaseCoffee) getname()string{
	return " "
}

