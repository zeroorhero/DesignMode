package state
import "fmt"

// 定义状态的接口
type Week interface {
	Today()
	NextDay(ctx *Context)
}

// 定义的环境类
// 定义客户端访问的接口
type Context struct {
	today Week
}

func NewContext() *Context{
	return &Context{today: &Sunday{}}
}
func (c *Context)Today(){
	c.today.Today()
}
func (c *Context)NexDay(){
	c.today.NextDay(c)
}

// 具体状态的实现
type Sunday struct {

}
func (s *Sunday)Today(){
	fmt.Println("today is sunday")
}
func (s *Sunday)NextDay(ctx *Context){
	ctx.today = &Monday{}
}

type Monday struct {

}
func (s *Monday)Today(){
	fmt.Println("today is monday")
}
func (s *Monday)NextDay(ctx *Context){
	ctx.today = &Sunday{}
}