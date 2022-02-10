package adaptee

// 需要进行适配的目标接口
// 这个是adapter需要进行继承操作的
type target interface {
	Request() string
}


// 这个是现在已经写好的接口的
type adaptee interface {
	SpecialRequest() string
}

type adapteeImp struct {}
func NewadapteeImp()adaptee{
	return &adapteeImp{}
}
func (a *adapteeImp)SpecialRequest() string{
	return "specialrequest"
}


// 这个是转换器
type adapter struct {
	adaptee
}
// 注意一定要写一个构造函数
func Newadapter(ad adaptee)target{
	return &adapter{adaptee:ad}
}
// 实现了目标接口但是还是使用原先的接口的
func (a *adapter)Request() string{
	return a.adaptee.SpecialRequest()
}