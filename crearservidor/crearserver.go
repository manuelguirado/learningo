package main
import (
	"fmt"
	"net/http"
)
func manejador(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hola este es el servidor ", r.URL.Path)
}
func main(){
	http.HandleFunc("/health",manejador	)
	fmt.Print("el servidor se encuentra en ejecución  en :" , "http:/localhost:3000")
	http.ListenAndServe(":8080", nil)
}