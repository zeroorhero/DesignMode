package facade

import "fmt"

// 考虑创建多个接口的
type API interface {
	Test() string
}

type ModuleA interface {
	TestA() string
}

type ModuleB interface {
	TestB() string
}

type apiImp struct {
	modulea ModuleA
	moduleb ModuleB
}

func (a *apiImp) Test() string{
	return fmt.Sprintf("%s_%s", a.modulea.TestA(), a.moduleb.TestB())
}

func NewapiImp(modulea ModuleA, moduleb ModuleB) API{
	return &apiImp{
		modulea: modulea,
		moduleb: moduleb,
	}
}

type ModuleAImp struct {}
func NewModuleAImp()ModuleA{
	return &ModuleAImp{}
}
func (m *ModuleAImp)TestA() string{
	return "moduleA"
}

type ModuleBImp struct {}
func NewModuleBImp()ModuleB{
	return &ModuleBImp{}
}
func (m *ModuleBImp)TestB() string{
	return "moduleB"
}

