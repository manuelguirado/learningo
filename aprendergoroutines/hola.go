package main
import (
	"fmt"
	"time"
)
func sayhi() {
	fmt.Println("Hola")

}
func sayhello() {
	fmt.Println("hello")
}
func main() {
	go sayhi()
	go sayhello()
	time.Sleep(1 * time.Second)
}