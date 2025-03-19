package main
import (
	"fmt"
	"os"
	"strings"
)
func main(){
	dir := "./"
	extension := ".go"
	entries,err := os.ReadDir(dir)
	if err != nil {
		fmt.Println("error al leer el archivo" , err)
		return
	}
	fmt.Print("archivos que tienen la extensión %s: \n" , extension,dir)
	for _, entry := range entries{
		if  !entry.IsDir()	&& strings.HasSuffix(entry.Name(), extension){
			fmt.Println(entry.Name())
		}
	}
}