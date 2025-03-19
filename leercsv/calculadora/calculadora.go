package main
import "fmt"
func sumar( a int , b int) int {
	return a + b
}
func restar(a int, b int) int{
	return a - b
}
func multiplicar(a int , b int) int{
	return a * b
}
func dividir ( a int , b int) int{
	return a / b
}
func resto( a int, b int) int {
	return a % b
}
func runCalculator(){
	var operacion string
	var numero1 int 
	var numero2  int
	var resultado int
	fmt.Print("Ingrese la operacion que desea realizar: ")
	fmt.Scanf("%s", &operacion)
	fmt.Print("Ingrese el primer numero: ")
	fmt.Scanf("%d", &numero1)
	fmt.Print("Ingrese el segundo numero: ")
	fmt.Scanf("%d ", &numero2)
	switch operacion{
	case "+" :
		resultado = sumar(numero1,numero2)
		fmt.Print("El resultado de la suma es: ", resultado)
	case "-" :
		resultado = restar(numero1,numero2)
		fmt.Printf("El resultado de la resta es: %d" , resultado)
	case "*":
		resultado = multiplicar(numero1,numero2)
		fmt.Printf("El resultado de la multiplicacion es: %d", resultado)
	case "/" :
		resultado = dividir(numero1, numero2)
		fmt.Printf("El resultado de la division es: %d", resultado)
	case "%":
		resultado = resto(numero1,numero2)
		fmt.Printf("El resultado del resto es: %d", resultado)
	default:
		fmt.Print("Operacion no valida")
	}
}
func main(){
	runCalculator()
}
