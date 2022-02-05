package proxy

import "fmt"

// 抽象出来的接口
// 目标类和代理类都要对其进行实现的
// 代理类实现了对其功能的增强
type Subject interface {
	Do() string
}


type Real struct {}
func (r *Real) Do() string{
	return "real"
}

type Proxy struct {
	real Real
}
func (r *Proxy)Do() string{
	var res string
	res += "before"
	// 会对real 进行创建
	// 结构体的零值是每一个属性都是零值
	res += r.real.Do()
	fmt.Println(r.real)

	res += "after"
	return res
}
