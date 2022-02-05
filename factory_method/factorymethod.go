package factory_method
//工厂方法模式
//
//工厂方法模式使用子类的方式延迟生成对象到子类中实现。
//
//Go中不存在继承 所以使用匿名组合来实现
// 和简单工厂都是比较相似的  只不过是工厂函数由子类进行实现的
type Operator interface {
	SetA(string)
	SetB(string)
	Result()string
}

type OperatorFactory interface {
	Create()Operator
}

type OperatorBase struct {
	A string
	B string
}
func (o *OperatorBase)SetA(a string){
	o.A = a
}
func (o *OperatorBase)SetB(b string){
	o.B = b
}

type PlusOperator struct {
	*OperatorBase
}
func (p *PlusOperator)Result()string{
	return p.A + p.B
}

type PlusOperatorFactory struct {}
func (p *PlusOperatorFactory) Create()Operator{
	return &PlusOperator{&OperatorBase{}}
}