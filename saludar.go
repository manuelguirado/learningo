package main

import "fmt"
func hello() {
	var nombre string
	fmt.Print("Ingresa tu nombre: ")
	fmt.Scanf("%s", &nombre)
	fmt.Print("hola", nombre)
}
