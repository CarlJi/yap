// Code generated by gop (Go+); DO NOT EDIT.

package main

import "github.com/goplus/yap"

const _ = true

type AppV2 struct {
	yap.AppV2
}
//line demo/classfile2_static/main.yap:1
func (this *AppV2) MainEntry() {
//line demo/classfile2_static/main.yap:1:1
	this.Static__0("/foo", this.FS("public"))
//line demo/classfile2_static/main.yap:2:1
	this.Static__0("/")
//line demo/classfile2_static/main.yap:4:1
	this.Run(":8080")
}
func main() {
	yap.Gopt_AppV2_Main(new(AppV2))
}
