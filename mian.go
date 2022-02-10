package main

import "fmt"

type Manager interface {
	HaveRight(money int) bool
	HandleFeeRequest(name string, money int) bool
}

// 通过嵌入一个接口  并且添加属性
// 实现类似java中的抽象类
type RequestChain struct {
	Manager                    // 相当于继承一个接口了
	successor *RequestChain    // 责任链的下一个  继承者
}

func add(i, j int)int{
	return i + j
}

// 定义的是一个类型
type cal func(int, int)int

func main(){
	var c cal
	c = add
	fmt.Println(c(1, 2))
}
