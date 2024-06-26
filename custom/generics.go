package custom

import (
	"fmt"
	"time"
)

type Constraint interface {
	int | float64
}

func Sum[T Constraint](x T, y T) T {
	return x + y
}

func Gene() {
	fmt.Print(gen_great())
	for i := 0; i < 3; i++ {
		time.Sleep(500 * time.Millisecond)
		fmt.Print(".")
	}
	time.Sleep(500 * time.Millisecond)
	pl("\nHii Generics...!")
	pl()
	pl("Adding int: ", Sum(5, 4))
	pl("Adding float64: ", Sum(5.912, 4.726493))
}

func gen_great() string {
	return "Calling Generics"
}
