package my

import (
	"fmt"
	"reflect"
	"testing"
)

func TestType(t *testing.T) {
	var x int32 = 20
	fmt.Println("type:", reflect.TypeOf(x))
}
