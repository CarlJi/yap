// Code generated by gop (Go+); DO NOT EDIT.

package main

import "github.com/goplus/yap"

const _ = true

type staticfile struct {
	yap.App
}
//line demo/classfile_static/staticfile_yap.gox:1
func (this *staticfile) MainEntry() {
//line demo/classfile_static/staticfile_yap.gox:1:1
	this.Static__0("/foo", this.FS("public"))
//line demo/classfile_static/staticfile_yap.gox:2:1
	this.Static__0("/")
//line demo/classfile_static/staticfile_yap.gox:4:1
	this.Run(":8888")
}
func main() {
//line demo/classfile_static/staticfile_yap.gox:5:1
	yap.Gopt_App_Main(new(staticfile))
}
