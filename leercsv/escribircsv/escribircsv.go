package main
import (
	"encoding/csv"
	"fmt"
	"os"
)
func main(){
	file,err := os.Create("fichero.csv")
	if err != nil{
		fmt.Print("error al crear el archivo ", err)
		return 
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()
	recors :=  [][] string{
		{"nombre","apellido", "edad"},
		{"manuel", "guirado","20"},
		{"paco" , "perez" , "30"},
		{"luis", "gomez", "40"},
	}	
	for _, record := range recors{
		err := writer.Write(record)
		if err != nil {
			fmt.Println("error al escribir el archivo" , err)
			return 
		}
	}
}