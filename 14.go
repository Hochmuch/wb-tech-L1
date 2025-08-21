package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x interface{}

	x = false

	switch x.(type) {
	case int:
		fmt.Println("x целое число")
	case string:
		fmt.Println("x это строка")
	case bool:
		fmt.Println("x это булево значение")
	}
	isChan := reflect.TypeOf(x)
	if isChan.Kind() == reflect.Chan {
		fmt.Println("x это канал")
	}
}
