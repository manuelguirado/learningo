package main

import (
	"log"
	"os"
	"path/filepath"
	"time"
	"github.com/fsnotify/fsnotify"
)
func main(){
archivo := "hola.txt"
	
	watcher, err  := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
defer 	watcher.Close()
addFileToWatcher := func(file string) {
	err := watcher.Add(archivo)
	if err != nil {
		log.Println("error al agregar el archivo al watcher" , err)

	}else{
		log.Println("archivo agregado al watcher " , archivo)
	}

}

addFileToWatcher(archivo)
log.Println("cambios echo en el archivo" , archivo)
done := make(chan bool)
	go func() {
		for {
			select{
			case event, ok := <-watcher.Events:
				if !ok{
					return	
				}
				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Print("el archivo  fue modificado" , event.Name)

				}
				if event.Op&fsnotify.Rename == fsnotify.Rename || event.Op&fsnotify.Remove == fsnotify.Remove {
					log.Print("el archivo fue renombrado o eliminado" , event.Name)
					 newName  := findNewFileName(filepath.Dir(event.Name), filepath.Base(event.Name))
					 if newName != " " {
						archivo = newName
						log.Print("El archivo ha sido nombrado a " , archivo)
						time.Sleep(500 * time.Millisecond)
						addFileToWatcher(archivo)
					 }
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
					log.Println("error" , err)
			}
		}

	}()
	<-done 
	log.Println("programa ha terminado")


}
func findNewFileName(dir,oldName string) string{
	entries,err :=  os.ReadDir(dir)
	if err != nil{
		log.Println("error al leer el directorio", err)
		return ""
	}
	for _, entry := range entries {
		if !entry.IsDir() && entry.Name() != oldName {
			return filepath.Join(dir,entry.Name())
		}
	}
	return ""
}