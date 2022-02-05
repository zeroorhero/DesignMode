package singleton

import "sync"

type Singleton interface {
	foo()
}

type sington struct {}
func (s *sington)foo(){

}

var (
	instance *sington
	once sync.Once
)

func getInstance() Singleton{
	once.Do(func() {
		instance = &sington{}
	})
	return instance
}