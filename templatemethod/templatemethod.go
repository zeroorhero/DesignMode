package main

import "fmt"

// 定义接口
// 这个接口就是一系列小方法的集合的
// 就是用于整合小方法的具体的流程的
type Downloader interface {
	Download(uri string)
}

// 用于定义子类的接口的
type implement interface {
	download()
	save()
}

// 定义模版
// 属于父类
type template struct {
	// 放到这里是要有重新写的方法的
	implement  // 包含子类的方法
	uri string  // 自己的字段
}

func newTemplate(impl implement) *template {
	return &template{
		implement: impl,
	}
}


// 定义了模版方法
func (t *template) Download(uri string) {
	t.uri = uri
	fmt.Print("prepare downloading\n")
	// 父类调用子类的方法
	t.implement.download()
	t.implement.save()
	fmt.Print("finish downloading\n")
}

func (t *template) save() {
	fmt.Print("default save\n")
}

type HTTPDownloader struct {
	*template
}

func NewHTTPDownloader() Downloader {
	downloader := &HTTPDownloader{}
	template := newTemplate(downloader)
	downloader.template = template
	return downloader
}

func (d *HTTPDownloader) download() {
	fmt.Printf("download %s via http\n", d.uri)
}

func (*HTTPDownloader) save() {
	fmt.Printf("http save\n")
}

type FTPDownloader struct {
	*template
}

func NewFTPDownloader() Downloader {
	downloader := &FTPDownloader{}
	template := newTemplate(downloader)
	downloader.template = template
	return downloader
}

func (d *FTPDownloader) download() {
	fmt.Printf("download %s via ftp\n", d.uri)
}
