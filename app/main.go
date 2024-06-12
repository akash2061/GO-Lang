package main

import (
	utils "example/project/utils"
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Hello", utils.Name)
	intArr := []int{2, 3, 5, 7, 11}
	strArr := utils.IntArrToStrArr(intArr)
	fmt.Println(strArr)
	fmt.Println(reflect.TypeOf((strArr)))
}
