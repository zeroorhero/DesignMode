package abstractfactory

import "fmt"

//抽象工厂模式
//
//抽象工厂模式用于生成产品族的工厂，所生成的对象是有关联的。
//
//如果抽象工厂退化成生成的对象无关联则成为工厂函数模式。
//
//比如本例子中使用RDB和XML存储订单信息，抽象工厂分别能生成相关
//的主订单信息和订单详情信息。 如果业务逻辑中需要替换使用的时候只需
//要改动工厂函数相关的类就能替换使用不同的存储方式了。

// 首先定义接口 这个是产品相关的
type MainDao interface {
	SaveMainDao()
}

type DetialDao interface {
	SaveDetailDao()
}

// 定义工厂接口
// 用于生产一个相关的产品组
type DaoFactory interface {
	CreatMainDao()MainDao
	CreatDetailDao() DetialDao
}

// 这个是产品1
type RGBMainDao struct {}
func (r *RGBMainDao) SaveMainDao(){
	fmt.Println("rgb Main save")
}

// 这个是产品2
type RGBDetailDao struct {}
func (r *RGBDetailDao) SaveDetailDao(){
	fmt.Println("rgb detail save")
}

// 实现工厂
type RGBFactory struct {}
func (r *RGBFactory)CreatMainDao()MainDao{
	return &RGBMainDao{}
}
func (r *RGBFactory)CreatDetailDao() DetialDao{
	return &RGBDetailDao{}
}