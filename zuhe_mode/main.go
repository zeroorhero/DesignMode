package main

import "fmt"

type Component interface {
	Parent()Component
	SetParent(Component)
	Name()string
	SetName(string)
	AddChild(Component)
	Print(string)
}

const(
	LeafLoad = iota
	CompositeNode
)

type component struct {
	parent Component
	name string
}

func (c *component)Parent()Component{
	return c.parent
}
func (c *component)SetParent(parent Component){
	c.parent=parent
}
func (c *component)Name()string{
	return c.name
}
func (c *component)SetName(name string){
	c.name = name
}
func (c *component)AddChild(Component){

}
func (c *component)Print(string){

}


type Leaf struct {
	component
}

func NewLeaf() *Leaf{
	return &Leaf{}
}

type Composite struct {
	component
	child []Component  // 叶子集合
}

// 创建一个组合结构体
func NewComposite() *Composite{
	return &Composite{
		child: make([]Component, 0),
	}
}

func (c *Composite)AddChild(child Component){
	child.SetName(c)
	c.child=append(c.child, child) // 加入孩子节点
}
func (c *Composite)Print(pre string){
	fmt.Println(pre, c.name)
	pre += " "
	for _, com := range c.child{
		com.Print(pre)
	}

}