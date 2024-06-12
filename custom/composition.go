package functions

import (
	"fmt"
	"time"
)

type contact struct {
	Name   string
	number string
}
type business struct {
	name string
	addr string
	contact
}

func (b business) info() {
	fmt.Printf("Contact: %s at %s", b.contact.Name, b.name)
}
func Compos() {
	pl(cmp_great())
	c1 := contact{
		"V",
		"555-1212",
	}
	b1 := business{
		"Arasaka-Corporation",
		"City Center, NC",
		c1,
	}
	b1.info()
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
