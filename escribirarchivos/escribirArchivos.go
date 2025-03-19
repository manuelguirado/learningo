package main

import (
	"fmt"
	"os"
)

func main(){
 file, err := os.Create("hola.txt")
  if err != nil {
	 fmt.Print(err)
		return
	}
	l, err := file.WriteString("hello wolrd")
	if err != nil {
		fmt.Println(err)
		file.Close()
		return
	}
	fmt.Println(l, "el archivo ha sido escrito exitósamente")
	err = file.Close()
    if err != nil {
		fmt.Println(err)
		return
	}

}