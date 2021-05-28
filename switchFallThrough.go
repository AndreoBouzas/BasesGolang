package main

import (
	"fmt"
)

func main() {

	//Switch= avalia as diferentes expressões e executa a verdadeira, analisando caso a caso. 
	
	//Criando uma variável 
	
	comidafavorita := "Massa Carbonara"
	
	
	//Declarando o verificador Switch
	
	switch comidafavorita {
	
		//para determinar as verificarções usamos o Case
		
		//verifica se o valor da variável  é igual = Pizza
		case "Pizza":
		
			//caso verdadeiro da um print em tela 
			fmt.Println("minha comida favorita é Pizza!")									
		
		//verifica se o valor da variável  é igual = Massa Carbonara
		case "Massa Carbonara":
			//caso verdadeiro da um print em tela 
			fmt.Println("minha comida favorita é Massa Carbonara!")
			
			
			// A condição fallthrough dentro de um case torna o próximo case verdadeiro!
			fallthrough
			
			
			
		//verifica se o valor da variável  é igual = Lasanha
		case "Lasanha":
			//caso verdadeiro da um print em tela 
			fmt.Println("minha comida favorita é Lasanha!")
		
		//verifica se o valor da variável  é igual = Bife a parmegiana
		case "Bife a parmegiana":
			//caso verdadeiro da um print em tela 
			fmt.Println("minha comida favorita é Bife a parmegiana:")
	
	}
	
	
}
