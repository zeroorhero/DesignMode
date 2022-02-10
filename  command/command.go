package _command

import "fmt"

// 定义的命令
type Command interface {
	Execute()
}

// 命令的接受者
type BoardMethod struct {}
func (b *BoardMethod)start(){
	fmt.Println("start !")
}

func (b *BoardMethod) reboot(){
	fmt.Println("reboot")
}

type StartCommand struct {
	b *BoardMethod
}
func NewStartCommand(b *BoardMethod) *StartCommand{
	return &StartCommand{b: b}
}
func (s *StartCommand)Execute(){
	s.b.start()
}

type RebootCommand struct {
	b *BoardMethod
}
func NewRebootCommand(b *BoardMethod) *RebootCommand{
	return &RebootCommand{b: b}
}
func (r *RebootCommand)Execute(){
	r.b.reboot()
}

// 命令的发起者
type Box struct {
	button1 Command
	button2 Command
}

func NewBox(button1, button2 Command)*Box{
	return &Box{
		button1: button1,
		button2: button2,
	}
}

func (b *Box) PressButton1(){
	b.button1.Execute()
}
func (b *Box) PressButton2(){
	b.button2.Execute()
}