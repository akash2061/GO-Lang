package map_func

import (
	"fmt"
)

var pl = fmt.Println

func Maps() string {

	//! var MyMap map [keyType]valueType

	// var values map[int8]string
	// values = make(map[int8]string)
	// Values := make(map[int8]string)

	cmd := make(map[string]string)
	cmd["h"] = "help"
	cmd["v"] = "version"

	values := map[int]string{1: "One", 2: "Two", 3: "Three"}

	fmt.Printf("Values: %v\n", values)
	pl("Out of range access :", values[4])
	_, ok := values[4]
	pl("value ok:", ok)

	return "Hello Maps"
}
