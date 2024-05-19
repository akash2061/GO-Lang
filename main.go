package main

import (
	"bufio"   //! Read-Input
	"fmt"     //! Print-Format
	"log"     //! For Errors
	"os"      //! Input-Output
	"reflect" //! TypeOf
	"strconv" //! To Deal with String Manipulation
	"unsafe"  //! SizeOf
)

var pl = fmt.Println               //? Print - Alias
var text = "This is a global text" //? Global var

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
	fmt.Println(age, "\t", reflect.TypeOf(age), "\tSize = ", unsafe.Sizeof(age))
	fmt.Println(num, "\t", reflect.TypeOf(num), "\tSize = ", unsafe.Sizeof(num))
	fmt.Println(name2, "\t", reflect.TypeOf(name2), "\tSize = ", unsafe.Sizeof(name2))

	fmt.Printf("My name is %v and age is %v\n", name, age)
	fmt.Printf("My name is %s and age is %d\n", name, age)
	fmt.Printf("name type = %T\nage type = %T\nnum type = %T\n", name, age, num)

	str := fmt.Sprintf("My name is %v and age is %v\n", name, age)
	fmt.Println("\nString is: ", str)

	pl("Hello... Alias... :)")

	pl("\nWhat is your Name?")
	reader := bufio.NewReader(os.Stdin)
	usr_name, err := reader.ReadString('\n')
	if err == nil {
		pl("Hello", usr_name)
	} else {
		log.Fatal(err)
	}

	var v1, v2 = 10, 1.5
	pl(float64(v1) + v2)
	v1 = int(v2)
	pl("Floor of v2:", v1)

	cv := "50"
	cv1, err := strconv.Atoi(cv)
	pl("String: cv", cv, "Atoi: cv1", cv1, "ERROR:", err, "Type1: cv", reflect.TypeOf(cv), "Type2: cv1", reflect.TypeOf(cv1))
	cv2 := "3.14"
	if cv3, err := strconv.ParseFloat(cv2, 32); err == nil {
		pl(cv3)
	}
}
