package Obsever

import "fmt"


// 主题类
// 内部含有一个观察者的列表
// 含有添加 删除 和 通知的方法
type Subject struct {
	obsevers []Obsever
	context string
}

func NewSubject(context string) *Subject{
	return &Subject{
		obsevers: make([]Obsever,  0),
		context:  context,
	}
}

func (s *Subject) Add(obsever Obsever){
	s.obsevers = append(s.obsevers, obsever)
}

func (s *Subject) Notify(){
	for _, o := range s.obsevers{
		o.Update(s)
	}
}

func (s *Subject) UpdateContext(context string){
	s.context = context
	s.Notify()
}

type Obsever interface {
	Update(subject *Subject)
}

type Reader struct {
	name string
}

func NewReader(name string) *Reader{
	return &Reader{
		name: name,
	}
}

func (r *Reader)Update(subject *Subject){
	fmt.Printf("reader %s, %s", r.name, subject.context)
}