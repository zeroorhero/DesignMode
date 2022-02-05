package simple_factory

//简单工厂模式
//
//go 语言没有构造函数一说，所以一般会定义NewXXX函数来初始化相关类。
//NewXXX 函数返回接口时就是简单工厂模式，也就是说Golang的一般推荐做法就是简单工厂。
//
//在这个simplefactory包中只有API 接口和NewAPI函数为包外可见，封装了实现细节。

type API interface {
	Say()string
}


// 类似于构造函数
func NewAPI(str string) API{
	if str == "1"{
		return &type1{}
	}else if str == "2"{
		return &type2{}
	}else{
		return nil
	}
}

type type1 struct {

}
func(t *type1)Say()string{
	return "type1"
}

type type2 struct {

}
func(t *type2)Say()string{
	return "type2"
}
