package custom

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Fmt_val() {
	great_inter()
	x := []int{10, 20, 30}
	fmt.Printf("%#v\n", x)
	fmt.Printf("%+v\n", x)
	fmt.Printf("%T\n", x)
	fmt.Printf("%b\n", x)
	fmt.Printf("%X\n", x)
	fmt.Printf("%p\n", &x)
	x = []int{1, 2, 3}
	m, n := fmt.Printf("%o\n", x)
	fmt.Println("M =", m, "N =", n)

	reader := bufio.NewReader(os.Stdin)
	s, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	s = strings.TrimSpace(s)
	value, err := strconv.ParseInt(s, 16, 16)
	if err != nil {
		log.Fatal(err)
	}
	pl(value)
}
