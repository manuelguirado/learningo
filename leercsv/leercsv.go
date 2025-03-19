package main

import (
	
	"encoding/csv"
	"fmt"
	"os"
)
func main(){
	file, err := os.Open("./fichero.csv")
	if err != nil {
		fmt.Println("error al abrir el archivo" , err)
		return
	}
	defer file.Close()
	reader := csv.NewReader(file)
	reader.Comma = ';'
	reader.FieldsPerRecord = -1
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("error al leer el archivo", err)
	}
	for i , record := range records {
		fmt.Print(i, record)
	}
}