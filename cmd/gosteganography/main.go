package main

import (
	"fmt"
	"log"

	"github.com/lucasmenendez/gosteganography"
)

func main() {
	msg := []byte("hello world!")

	bits, err := gosteganography.HideMessage("./input.png", "./output.png", msg)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Ok. %d bits writted.\n", bits)
}
