package composite
// 继承？？？？？？？？？？？
import "fmt"

type Composite interface {
	Parent() Composite
	SetParent(parent Composite)
	Name() string
	SetName(name string)
	Print(pre string)
	AddChild(child Composite)
}

type composite struct {
	parent Composite
	name string
}
func (c *composite)Parent() Composite{
	return c.parent
}
func (c *composite)SetParent(parent Composite){
	c.parent = parent
}
func (c *composite)Name() string{
	return c.name
}
func (c *composite)SetName(name string){
	c.name = name
}
func (c *composite)Print(pre string){
	// 空实现
}
func (c *composite)AddChild(child Composite){

}

type Leaf struct {
	composite
}
// 这个构造函数一定要返回特定的类型
// 应该返回该类型的
func NewLeaf()*Leaf{
	return &Leaf{}
}
func (l *Leaf)Print(pre string){
	fmt.Printf("%s-%s\n", pre, l.Name())
}


type Component struct {
	composite
	child []Composite
}
func NewComponent() *Component{
	return &Component{
		child: make([]Composite,8),
	}
}
func (c *Component)AddChild(child Composite){
	child.SetParent(c)
	c.child = append(c.child, child)
}
func (c *Component)Print(pre string){

}