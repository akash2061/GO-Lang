package functions

import (
	"fmt"
	"time"
)

func Compos() {
	cmp_great()
}

func cmp_great() string {
	fmt.Print("\nLoading Composition")
	for i := 0; i < 3; i++ {
		time.Sleep(500 * time.Millisecond)
		fmt.Print(".")
	}
	time.Sleep(500 * time.Millisecond)
	return "\n"
}
