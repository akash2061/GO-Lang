package functions

import (
	"fmt"
	"time"
)

func Gene() {
	fmt.Print(gene())
	for i := 0; i < 3; i++ {
		time.Sleep(500 * time.Millisecond)
		fmt.Print(".")
	}
	fmt.Println("\nHii")
}

func gene() string {
	return "Calling Generics"
}
