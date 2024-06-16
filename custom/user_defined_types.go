package custom

import "fmt"

type Tsp float64
type TBs float64
type ML float64

// Convert with functions (Bad Way)
// func tspToML(tsp Tsp) ML {
// 	return ML(tsp * 4.92)
// }

// func TBToML(tbs TBs) ML {
// 	return ML(tbs * 14.79)
// }

// Associate method with types
func (tsp Tsp) ToMLs() ML {
	return ML(tsp * 4.92)
}
func (tbs TBs) ToMLs() ML {
	return ML(tbs * 14.79)
}

func Udt() {
	pl("\nUser Defined Datatypes...\n")
	ml1 := Tsp.ToMLs(3)
	fmt.Printf("3 tps = %.2f mL\n", ml1)
	ml1 = TBs.ToMLs(2)
	fmt.Printf("2 tbs = %.2f mL\n", ml1)
}
