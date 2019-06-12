package main

import (
	"fmt"
	"math/rand"
)

func main() {
	r := rand.New(rand.NewSource(99))
	fmt.Println("Spaceline\t Days\t Trip\t type\t Price")
	fmt.Println("=====================================================")
	s := []string{"Virgin Galactic", "SpaceX", "Space Adventures"}
	for _, v := range s {
		fmt.Printf("%v\t %d\t Round-trip\t $96\n", v, r.Int)
		fmt.Printf("%v\t %d\t Round-trip\t $37\n", v, r.Int)
	}
}
