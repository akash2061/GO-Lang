package main

import (
	packet "example/project/package"
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Hello", packet.Name)
	intArr := []int{2, 3, 5, 7, 11}
	strArr := packet.IntArrToStrArr(intArr)
	fmt.Println(strArr)
	fmt.Println(reflect.TypeOf((strArr)))
}
