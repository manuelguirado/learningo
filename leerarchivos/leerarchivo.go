package main

import (
	"io"
	"os"
	"log"
)

func main() {
	// Abrir el archivo de entrada
	inputFile, err := os.Open("hola.txt")
	if err != nil {
		log.Fatalf("Error al abrir el archivo de entrada: %v", err)
	}
	defer closeFile(inputFile)

	// Crear el archivo de salida
	outputFile, err := os.Create("output.txt")
	if err != nil {
		log.Fatalf("Error al crear el archivo de salida: %v", err)
	}
	defer closeFile(outputFile)

	// Copiar el contenido del archivo de entrada al archivo de salida
	if err := copyFileContents(inputFile, outputFile); err != nil {
		log.Fatalf("Error al copiar el contenido del archivo: %v", err)
	}
}

// closeFile cierra el archivo y maneja cualquier error
func closeFile(f *os.File) {
	if err := f.Close(); err != nil {
		log.Fatalf("Error al cerrar el archivo: %v", err)
	}
}

// copyFileContents copia el contenido de un archivo a otro
func copyFileContents(src, dst *os.File) error {
	buf := make([]byte, 1024)
	for {
		n, err := src.Read(buf)
		if err != nil && err != io.EOF {
			return err
		}
		if n == 0 {
			break
		}
		if _, err := dst.Write(buf[:n]); err != nil {
			return err
		}
	}
	return nil
}
