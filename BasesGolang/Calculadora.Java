import java.util.Scanner;

public class Main {
    public static void main(String[] args) {
        boolean continuar = true;
        char opcao;
        Scanner entrada = new Scanner(System.in);
        Scanner in = new Scanner(System.in);
        int num1, num2, resultado, operador;
        
        while(continuar){
            
            System.out.println("Calculadora Iniciada!");
            System.out.print("deseja realizar um novo Calculo? S ou N ?");
            opcao = entrada.nextLine().charAt(0);
           
            if(opcao=='N'){
                continuar=false;
                System.out.println("Calculadora encerrada");
            }
            else{
                System.out.println("Novo calculo");
                System.out.println("Insira um número: ");
                num1 = in.nextInt();
                System.out.println("Insira outro número: ");
                num2 = in.nextInt();
                System.out.println("Qual Operação voccê deseja realizar:");
                System.out.println("digite 1 para Somar!");
                System.out.println("digite 2 para Subtrair!");
                System.out.println("digite 3 para Multiplicar!");
                System.out.println("digite 4 para Dividir!");
                operador = in.nextInt();
                
                switch(operador){
                    case 1:
                        resultado = num1 + num2;
                        System.out.println("O resultado da soma é: " + resultado);
                        break;
                    case 2:
                        resultado = num1 - num2;
                        System.out.println("O resultado da subtração é: "+ resultado);
                        break;
                    case 3:
                        resultado = num1 * num2;
                        System.out.println("O resultado da multiplicação é: "+ resultado);
                        break;
                    case 4:
                        resultado = num1 / num2;
                        System.out.println("O resultado da divisão é: "+ resultado);
                        break;
                    default:
                        System.out.println("Favor informar uma opção válida para o calculo!");
                }
            };
        }
    }
}