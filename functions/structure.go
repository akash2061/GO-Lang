package functions

import (
	"fmt"
	"time"
)

type custom struct {
	name string
	addr string
	bal  float64
}

func Str() {
	pl(str_great())
	var c1 custom
	c1.name = "Vincent-V"
	c1.addr = "8 Megabuilding H10, NC"
	c1.bal = 15324.609

	cInfo(c1)
	pl()
	uAddr(&c1, "Motel Kabuki Market, NC")
	pl(c1.name+" completing gigs at Address :", c1.addr)
}

func cInfo(c custom) {
	fmt.Printf("%s owns : %.2f Eddied.\n", c.name, c.bal)
	fmt.Printf("%s lives at %s\n", c.name, c.addr)
}
func uAddr(c *custom, addrs string) {
	c.addr = addrs
}

func str_great() string {
	fmt.Print("\nLoating Structures")
	for i := 0; i < 3; i++ {
		time.Sleep(500 * time.Millisecond)
		fmt.Print(".")
	}
	time.Sleep(500 * time.Millisecond)
	return "\n"
}
