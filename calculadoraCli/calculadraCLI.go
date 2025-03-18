package main

import "fmt"

func runCalculatorCli() {
	var operacion string
	var numero1 int
	var numero2 int
	var resultado int
	fmt.Printf("Ingrese la operacion que desea realizar: ")
	if operacion == "-a" {
		fmt.Printf("Ingrese el primer numero: ")
		fmt.Scanf("%d", &numero1)
		fmt.Printf("Ingrese el segundo numero: ")
		fmt.Scanf("%d", &numero2)
		resultado = numero1 + numero2
		fmt.Printf("El resultado de la suma es: %d", resultado)
	} else if operacion == "-b" {
		fmt.Print("Ingrese el primer numero: ")
		fmt.Scanf("%d", &numero1)
		fmt.Printf("Ingrese el segundo numero: ")
		fmt.Scanf("%d", &numero2)
		resultado = numero1 - numero2
		fmt.Printf("El resultado de la resta es: %d", resultado)

	}

}
func main(){
	runCalculatorCli()
}
