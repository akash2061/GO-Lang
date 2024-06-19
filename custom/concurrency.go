package custom

import (
	"fmt"
	"time"
)

func printTo3() {
	for i := 1; i <= 3; i++ {
		pl("Fun 1:", i)
	}
}
func printTo2() {
	for i := 1; i <= 2; i++ {
		pl("Fun 2:", i)
	}
}

func Conc() {
	fmt.Println("\nConcurrency:")
	go printTo2()
	go printTo3()
	time.Sleep(2 * time.Second)
}
