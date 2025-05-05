package main
import (
	"crypto/sha256"
	"fmt"
)
func main() {
	data := "hola mundo"
	hash  := sha256.Sum256([]byte(data))
	// Print the hash in hexadecimal format
	fmt.Printf("SHA256 hash of '%s': %x\n", data, hash)

}