package decorator

type Component interface {
	Cal() int
}

type ConcretComponent struct {}
func (c *ConcretComponent)Cal() int{
	return 0
}


type MulComponent struct {
	Component
	num int
}
func WrapMulComponent(c Component, num int) Component{
	return &MulComponent{
		Component: c,
		num:       num,
	}
}
func (m *MulComponent)Cal() int{
	return m.Component.Cal()*m.num
}

type AddComponent struct {
	Component
	num int
}
func WrapAddComponent(c Component, num int) Component{
	return &AddComponent{
		Component: c,
		num:       num,
	}
}
func (m *AddComponent)Cal() int{
	return m.Component.Cal()+m.num
}