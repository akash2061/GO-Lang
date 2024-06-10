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

	values := make(map[int8]string)

	values[1] = "one"

	return "Hello Maps"
}
