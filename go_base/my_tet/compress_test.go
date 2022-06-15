package main

import (
	"fmt"
	"testing"
)

const (
	originFilePath = "D:\\go_test\\go_test\\go_base\\my_tet\\"
	targetFilePath = "D:\\go_test\\go_test\\go_base\\my_tet\\"
)

func TestCompress(t *testing.T) {
	if err := CompressZip(originFilePath, targetFilePath); err != nil {
		fmt.Println("压缩异常")
	}
}
