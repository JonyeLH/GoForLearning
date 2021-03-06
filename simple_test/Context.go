package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Context struct {
	W http.ResponseWriter //ResponseWriter是接口所以直接使用
	R *http.Request       //Request是一个结构体所以使用指针
}

func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		W: w,
		R: r,
		// 一般路径参数都是一个，所以容量1就可以了
		// PathParams: make(map[string]string, 1),
	}
}

func (c *Context) ReadJson(data interface{}) error { //这里使用interface作为输入参数，表示接受任何格式的输入
	body, err := io.ReadAll(c.R.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, data)
	if err != nil {
		return err
	}
	fmt.Println("读取Json内容", data)
	return nil
}

func (c *Context) WriteJson(status int, data interface{}) error {
	c.W.WriteHeader(status)
	bs, err := json.Marshal(data)
	if err != nil {
		return err
	}
	_, err = c.W.Write(bs)
	if err != nil {
		return err
	}
	return err
}

func (c *Context) OkJson(data interface{}) error {
	// http 库里面提前定义好了各种响应码
	fmt.Println("成功相应", 200)
	return c.WriteJson(http.StatusOK, data)
}

func (c *Context) SystemErrJson(data interface{}) error {
	return c.WriteJson(http.StatusInternalServerError, data)
}

func (c *Context) BadRequestJson(data interface{}) error {
	return c.WriteJson(http.StatusBadRequest, data)
}
