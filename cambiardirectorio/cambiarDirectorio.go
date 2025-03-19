package main
import (
	"log"
	"os"
	"path/filepath"
)
func main(){
	 file, err := os.Create("hola.txt")
	if err != nil{
		log.Fatal("error al crear el archivo")
	}
	defer closeFile(file)
	dir := "/home/manudev/Escritorio/informatica/learningo/calculadora"
	fileInfo,err := os.Stat(dir )
	if err != nil {
		log.Fatal("error al verificar si es un archivo", err)
		return
	}
	if !fileInfo.IsDir(){
		log.Fatal("no es un directorio", dir)
	}
	targetPath := filepath.Join(dir,file.Name())
	err = os.Rename(file.Name(),targetPath)
	if err != nil {
		log.Fatal("error al cambiar de nombre")
	}
	if err != nil {
		log.Fatal("error al cambiar de directorio")
	}
	log.Print("archivo cambiado exitosamente",dir)

}
func closeFile(f *os.File){
	if err := f.Close(); err != nil{
		log.Fatal("error al cerar el archivo")
	}
}