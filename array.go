package main

import (
	"fmt"
)

func main() {

	//Switch= avalia as diferentes expressões e executa a verdadeira, analisando caso a caso. 
	
	//Criando uma variável 
	
	destinoDeFerias := "Acacou o dinheiro"
	
	
	//Declarando o verificador Switch
	
	switch destinoDeFerias {
	
		//para determinar as verificarções usamos o Case
		
		//verifica se o valor da variável  é igual = Nova York
		case "Nova York":
			//caso verdadeiro da um print em tela 
			fmt.Println("Pretendo viajar para Nova York esse mês!")									
		
		//verifica se o valor da variável  é igual = São Francisco
		case "São Francisco":
			//caso verdadeiro da um print em tela 
			fmt.Println("Pretendo viajar para São Francisco esse mês!")
			
		//verifica se o valor da variável  é igual = Canada
		case "Canada":
			//caso verdadeiro da um print em tela 
			fmt.Println("Pretendo viajar para o Canada esse mês!")
		
		//verifica se o valor da variável  é igual = Londres
		case "Londres":
			//caso verdadeiro da um print em tela 
			fmt.Println("Pretendo viajar para Londres esse mês!")
			
			
			
		//Defalut marca um padrão para o caso em que todas as verificações são negativas !
		default:
			fmt.Println("Não vou viajar porque to sem dinheiro!")
	
	}
	
	
}