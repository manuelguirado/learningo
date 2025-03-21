package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

type Producto struct {
	ID          int     `json:"id"`
	Nombre      string  `json:"nombre"`
	Descripcion string  `json:"descripcion"`
	Precio      float64 `json:"precio"`
	Categoria   string  `json:"categoria"`
	Disponible  bool    `json:"disponible"`
	Stock       int     `json:"stock"`
}
type usuarios struct {
	ID int `json:"id"`
	Nombre string `json:"nombre"`
	Telefono string `json:"telefono"`
	Email string `json:"email"`
	Fecha_registro string `json:"fecha_registro"`
	Activo bool `json:"activo"`
}
type users struct {
	Usuarios []usuarios `json:"usuarios"`

}
type Productos struct {
	Productos []Producto `json:"productos"`
}

func main() {
	file, err := os.Open("productos.json")
	if err != nil {
		fmt.Printf("Error al abrir el archivo: %v\n", err)
		return
	}
	defer file.Close()
	userFile, err := os.Open("usuarios.json")
	if err != nil {
		fmt.Printf("error al abrir el archivo %v\n", err)
	}
	defer userFile.Close()
	var users users 
	var productos Productos
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&productos)
	if err != nil {
		fmt.Printf("Error al leer el archivo: %v\n", err)
		return
	}
    decoder = json.NewDecoder(userFile)
	err = decoder.Decode(&users)
	if err != nil {
		fmt.Printf("error al leer el archivo %v\n", err)
		return
	}

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hola Mundo",
		})
	})
	r.GET("/productos", func(c *gin.Context) {
		c.JSON(200, productos)
	})
	r.GET("/usuarios" , func (c *gin.Context)  {
		c.JSON(200,users)
		
	})
	r.Run(":8080")
}
