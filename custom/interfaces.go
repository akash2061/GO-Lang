package custom

import (
	"fmt"
	"log"
)

type Animal interface {
	AngrySound()
	HappySound()
}

type Cat string

func (c Cat) Attack() {
	pl("Cat Attacks")
}

func (c Cat) Name() string {
	return string(c)
}

func (c Cat) AngrySound() {
	pl("Cat says Hissssss...")
	// c.Attack()
}
func (c Cat) HappySound() {
	pl("Cat says Purrrrrr...")
}

func Func_inter() {
	log.Print(great_inter())
	var kitty Animal = Cat("Tom")
	kitty.AngrySound()
	kitty.HappySound()
	var kitty2 Cat = kitty.(Cat)
	kitty2.Attack()
	pl("Cat Name :", kitty2.Name())
}

func great_inter() string {
	fmt.Println("\nInterfaces:")
	return "\n"
}
