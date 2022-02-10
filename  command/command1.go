package _command

import "fmt"

type comman interface {
	execute()
}

type BoardMethod1 struct {

}
func (b *BoardMethod1)start(){
	fmt.Println("start!!!")
}
func (b *BoardMethod1) reboot(){
	fmt.Println("roboot!!!")
}

type startCommand struct {
	b *BoardMethod1
}
func (s *startCommand)execute(){
	s.b.start()
}

type rebootCommand struct {
	b *BoardMethod1
}
func (s *rebootCommand)execute(){
	s.b.reboot()
}

type box struct {
	button1 comman
	button2 comman
}