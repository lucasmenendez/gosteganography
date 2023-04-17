package main

import "github.com/lucasmenendez/gosteganography"

func main() {
	msg := []byte("hello world!")
	gosteganography.HideMessage("./input.png", "./output.png", msg)
}
