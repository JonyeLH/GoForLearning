package main

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

//StructFunc的测试方法
func TestStructFunc(t *testing.T) {

	//声明结构函数输入的变量
	var (
		idIn   string
		nameIn string
	)

	//给结构函数输入的变量赋值
	idIn = "inputId"
	nameIn = "inputName"

	//得到结构
	ts := &TestStruct{}

	//调用结构函数1
	idOut, nameOut := ts.StructFunc(idIn, nameIn)

	if idOut == idIn && nameOut == nameIn {
		t.Log("测试通过！")
	} else {
		t.Error("函数执行错误")
	}

}

/*
goconvey使用基本的方法
*/
func TestAdd(t *testing.T) {
	Convey("测试add方法", t, func() {
		So(Add(2, 3), ShouldEqual, 5)
	})
}

//Json Decoder 使用方法
func TestDecoder(t *testing.T) {
	Decoder()
	fmt.Println("完成调用")
}

//find file path
func TestFind(t *testing.T) {
	//FindFile()
	FindAllFile()
	fmt.Println("查找的文件")
}
