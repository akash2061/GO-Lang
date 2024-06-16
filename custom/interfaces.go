package custom

import (
	"fmt"
	"log"
)

func Func_inter() {
	log.Print(great_inter())
}

func great_inter() string {
	fmt.Println("\nInterfaces:")
	return "\n"
}
