package main
import (
	"crypto/md5"
	"fmt"
)
func main() {
	word := "hello world"
	// Create a new hash object
	h := md5.New()
	// Write the data to the hash
	h.Write([]byte(word))
	// Calculate the hash
	hash := h.Sum(nil)
	// Print the hash in hexadecimal format
	fmt.Printf("MD5 hash of '%s': %x\n", word, hash)

}