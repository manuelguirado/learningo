package main
import (
	"fmt"
)

func contarVocales(cadena string) int {
	fmt.Print("ingrese una palabra")
	fmt.Scanf("%s", &cadena)
	var contar int 
	
		for _, char := range cadena {
			switch char {
			case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
				contar++
			}
		}
		return contar
	}

func main() {
	var cadena string
	fmt.Print("dime una cadena ")
	fmt.Scanf("%s", cadena)
	fmt.Print("hay ", contarVocales(cadena), "vocales en la palabra")
}
