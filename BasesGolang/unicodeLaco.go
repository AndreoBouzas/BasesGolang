package main

import (
	"fmt"
)

func main() {
	// início do laço 
	
	for x := 33; x <= 122; x++{
	
		//toda vez que o laço e rodado,é mostrado o valor de x transformando em string e unicode
		
		fmt.Printf("%d - %v\n",x, string(x))
	
	}
	
}
