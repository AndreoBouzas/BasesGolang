package main

import (
	"fmt"
)

func main() {
for h := 0; h <= 12; h++{
	fmt.Println("h:",h)
	for m := 0; m <= 60; m++{
		fmt.Println("m:",m)
		for s := 0; s <= 60; s++ {
			fmt.Print("s:",s)
		}
		fmt.Print("\n")
	}
	fmt.Print("\n")
}

}
