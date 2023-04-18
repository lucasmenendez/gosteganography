// Copyright (c) 2023, Lucas Menendez <hi@lucasmenendez.me>
// See LICENSE for licensing information

package gosteganography

import (
	"fmt"
	"log"
)

func Example() {
	expected := []byte("secret number: 1234")

	// open the input image
	image, err := Open("./input.png")
	if err != nil {
		log.Fatal(err)
	}
	// hide a message, it returns the number of bits writen
	nbits, err := image.Hide(expected)
	if err != nil {
		log.Fatal(err)
	}
	// get hided message using the number of bits
	got := image.Unhide(nbits)
	// [Optional] store the output
	// if err := image.Save("./output.png"); err != nil {
	// 	log.Fatal(err)
	// }
	fmt.Println(string(got))
	// Output: secret number: 1234
}