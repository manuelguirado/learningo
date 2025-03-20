package main
import (
	"encoding/json"
	"fmt"
	"os"
)
type Persona struct{
	Nombre   string `json:"nombre"`
	Edad     int    `json:"edad"`
	Direccion string `json:"direccion"`

}


func main() {
	personas := []Persona{
		{
			Nombre:   "juan",
			Edad:      30,
			Direccion: "calle cuba 41",
		},
	}
	file, err := os.Create("personas.json")
	if err != nil {
		fmt.Println("error al crear el archivo" , err)
	}
	defer  file.Close()
    encoder := json.NewEncoder(file)
	encoder.SetIndent(" " , " ")
	if err := encoder.Encode(personas); err != nil {
		fmt.Println("error al escribir el archivo" , err)
		return
	}
	fmt.Println("archivo creado con exito")


}