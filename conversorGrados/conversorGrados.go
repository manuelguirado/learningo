package main
import "fmt"
func convertToFahreint(grados int ) int{
  return (grados * 9/5) + 32
}
func conversorgrados(){
	var grados int
	fmt.Printf("dame grados para pasarlos a fahrenheint ")
	fmt.Scanf("%d" , &grados)
	fmt.Print("el resultado es:" ,  convertToFahreint(grados))
}
func main(){
	conversorgrados()
}