package main

import "fmt"

// 原型对象需要实现的接口
type Cloneable interface {
	Clone()Cloneable
}

// 原型对象的类
type ProtoManager struct {
	prototypes map[string]Cloneable
}

func NewProtoManager() *ProtoManager{
	return &ProtoManager{make(map[string]Cloneable)}
}
// 设置
func (p *ProtoManager) Set(name string, prototype Cloneable){
	p.prototypes[name] = prototype
}
// 获取
func (p *ProtoManager) Get(name string)Cloneable{
	return p.prototypes[name]
}

type Type1 struct {
	name string
}
func (t *Type1)Clone()Cloneable{
	tc := *t
	return &tc
}

type Type2 struct {
	name string
}
func (t *Type2)Clone()Cloneable{
	tc := *t
	return &tc
}

func main(){
	mgr := NewProtoManager()
	t1 := &Type1{name: "lcq"}
	mgr.Set("t1", t1)
	t11 := mgr.Get("t1")
	t22 := t11.Clone()
	if t11 == t22{
		fmt.Println("浅复制")
	}else {
		fmt.Println("深复制")
	}
}