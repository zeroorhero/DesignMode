package main

import "fmt"

type Downloader1 interface {
	Download1(str string)
	download1(str string)
	save()
}


type template1 struct {
	// 这个是一个抽象类 是最后需要被继承的
	// 而在继承中的需要重写这个接口中的方法的
	Downloader1
	name string
}
func Newtemplate1(do Downloader1, name string) *template1{
	return &template1{
		Downloader1: do,
		name:        name,
	}
}
func (t *template1)Download1(str string){
	fmt.Println("准备开始进行下载操作")
	t.download1(str)
	fmt.Println("开始进行下载")
	t.save()
	fmt.Println("下载完成")
}


type HttpDownload1 struct {
	*template1
}
func NewHttpDownload1() Downloader1{
	download := &HttpDownload1{}
	tem := Newtemplate1(download, "lcq")
	download.template1 = tem
	return download
}

func (h *HttpDownload1)download1(str string){
	fmt.Printf("%s start to download", str)
}
func (h *HttpDownload1) save(){
	fmt.Printf("finish saving!")
}