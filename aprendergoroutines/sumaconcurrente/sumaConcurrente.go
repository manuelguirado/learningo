package main
import (
	"fmt"
	"time" 
)

func sumar (numbers[]int, ch chan int) {
	sum := 0
	for _, v := range numbers {
		sum += v
	}
	ch <- sum // mandar la informacion de sum a channel previamente definido como ch 
}
func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	ch := make(chan int)
	go sumar(numbers[:len(numbers)/2] , ch) //divide el array en dos partes y los suma de manera concurrente 
	go sumar(numbers[len(numbers)/2:] , ch) //esta es la 2 mitad del array
	sum1, sum2 := <-ch, <- ch //recibe la informacion de los dos canales
	fmt.Println("suma total" , sum1 +  	sum2) 
	time.Sleep(1 * time.Second)

}