package main

import (
	"fmt"
)
var x int = 10

func main() {
	if x == 10 {
		fmt.Println("x é  10!")	
	
	}else if x > 100 {
		fmt.Println("x é maior que 100!")	
	
	}else if x < 100 {
		fmt.Println("x é menor que 100!")	
	
	}else {
		fmt.Println("x é menor que 10!")	
	}
}
