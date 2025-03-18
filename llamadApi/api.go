package main

import (
	"fmt"
	"io"
	"net/http"
)
func main(){
	api := "https://pokeapi.co/api/v2/pokemon"
	response, err := http.Get(api)
	if err != nil{
		fmt.Print("error al hacer la petición", err)
		return 
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Print("error leyendo la respuesta", err)
	}
	fmt.Println(string(body))

}