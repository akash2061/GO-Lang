package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

var text = "This is a global text"

func main() {
	fmt.Println("Hello... Go...!")

	var name string = "Morningstar_2061"
	var name1 = "akash2061"
	var name2 string
	name3 := "Hello"
	name2 = "Foo..."
	fmt.Println(name, "\n", name1, "\n", name2, "\n", name3, "\n", text)

	var age uint8 = 20
	num := 2.01
	fmt.Println(age, "\t", reflect.TypeOf(age), "\t", unsafe.Sizeof(age))
	fmt.Println(num, "\t", reflect.TypeOf(num), "\t", unsafe.Sizeof(num))
}
