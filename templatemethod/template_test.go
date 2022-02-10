package main

import "testing"

func TestNewHttpDownload1(t *testing.T) {
	var d Downloader1 = NewHttpDownload1()
	d.Download1("hhh")
}
