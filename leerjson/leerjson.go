package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"log"
)
type Persona struct {
	Nombre string `json:"nombre"`
	email string `json:"email"`
	Cargo  string `json:"cargo"`
}
//mapping del JSON
type usuarios struct {
	Usuarios []Persona `json:"usuarios"`
}

func main() {
	
	file, err := os.Open("usuarios.json")
	if err != nil {
		log.Fatalf("error al leer el archivo %v", err)
	}
	defer file.Close()
	filemanager, err := io.ReadAll(file)
	if err != nil {
		log.Fatal("error al leer el archivo", err)
	}

		
		var data usuarios
	err = json.Unmarshal(filemanager, &data)
	if err != nil {
		log.Fatal("error al leer el archivo")
	}
	for _, v := range data.Usuarios {
		fmt.Println("nombre" , v.Nombre)
		fmt.Println("gmail", v.email)
		fmt.Println("cargo", v.Cargo)

	}

	}
	