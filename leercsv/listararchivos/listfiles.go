package main
import (
	"fmt"
	"os"
)
func main(){ 
	dir := "../"
	entries, err := os.ReadDir(dir)
	 if err != nil {
		fmt.Println("error al listar los archivos")
	}
	fmt.Println("listando los archivos de :" ,dir , "\n")
	for _, entry  := range entries{
		if !entry.IsDir()  {
			fmt.Print(entry.Name())
		}
	}
}