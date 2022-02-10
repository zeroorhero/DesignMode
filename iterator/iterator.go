package iterator

import "fmt"

// 聚合类的接口
// 该方法返回迭代器
type Aggregate interface {
	Iterator() Iterator
}

// 迭代器的接口
type Iterator interface {
	First()
	IsDone() bool
	Next() interface{}
}


// 具体的聚合类
type Numbers struct {
	start, end int
}

func NewNumbers(start, end int) *Numbers {
	return &Numbers{
		start: start,
		end:   end,
	}
}

func (n *Numbers) Iterator() Iterator {
	return &NumbersIterator{
		numbers: n,
		next:    n.start,
	}
}


// 具体的迭代器类
type NumbersIterator struct {
	// 迭代器类应该包含聚合类的
	numbers *Numbers
	next    int
}

func (i *NumbersIterator) First() {
	i.next = i.numbers.start
}

func (i *NumbersIterator) IsDone() bool {
	return i.next > i.numbers.end
}

func (i *NumbersIterator) Next() interface{} {
	if !i.IsDone() {
		next := i.next
		i.next++
		return next
	}
	return nil
}

func IteratorPrint(i Iterator) {
	for i.First(); !i.IsDone(); {
		c := i.Next()
		fmt.Printf("%#v\n", c)
	}
}
