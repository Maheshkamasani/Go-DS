package main  

import (
	"fmt"
)

type Sport interface {
	sportsName() string 

}

type Human struct{

	name string 
    sport string 
}
func (h Human ) sportName() string{

	return h.name + "plays"  +  h.sport + "."
}

func main() {
	human1 := Human {"Mahesh", "chess"}

	fmt.Println(human1.sportName())
	human2 := Human{"Chinna", "Cricket"}
	fmt.Println(human2.sportName())
}