package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"reflect"
	"testing"
)

func TestTime(t *testing.T) {
	TimeStamp()
}

func TestType(t *testing.T) {
	var x int32 = 20
	fmt.Println("type:", reflect.TypeOf(x))
}

const (
	originFilePath = "D:\\go_test\\go_test\\go_base\\my_tet\\"
	targetFilePath = "D:\\go_test\\go_test\\go_base\\my_tet\\"
)

func CompressTest(t *testing.T) {
	if err := CompressZip(originFilePath, targetFilePath); err != nil {
		fmt.Println("压缩异常")
	}
}

func TestGin(t *testing.T) {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

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
//func TestAdd(t *testing.T) {
//	Convey("测试add方法", t, func() {
//		So(Add(2, 3), ShouldEqual, 5)
//	})
//}

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
