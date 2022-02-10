package flyweight

import "fmt"

// 这个是要进行共享的物体
type ImageFlyWeight struct {
	data string
}
func NewImageFlyWeight(data string) *ImageFlyWeight{
	return &ImageFlyWeight{
		data: data,
	}
}
func (i *ImageFlyWeight) Data() string{
	return i.data
}

// 这个是工厂
// 通过map将数据进行保存
var imageFlyWeightFactory *ImageFlyWeightFactory
type ImageFlyWeightFactory struct {
	factoty map[string]*ImageFlyWeight
}
// 单例模式  保证只有一个map
func GetImageFlyWeightFactory() *ImageFlyWeightFactory{
	if imageFlyWeightFactory == nil{
		imageFlyWeightFactory = &ImageFlyWeightFactory{
			factoty: make(map[string]*ImageFlyWeight),
		}
	}
	return imageFlyWeightFactory
}
func (f *ImageFlyWeightFactory)Get(name string)*ImageFlyWeight{
	data := f.factoty[name]
	if data == nil{
		data = &ImageFlyWeight{
			data: name,
		}
		f.factoty[name] = data
	}
	return data
}


type ImageViewer struct {
	*ImageFlyWeight
}
// 生成构造函数
func NewImageViewer(name string) *ImageViewer{
	image := GetImageFlyWeightFactory().factoty[name]
	return &ImageViewer{
		ImageFlyWeight : image,
	}
}
func (i *ImageViewer) display(){
	fmt.Println(i.data)
}