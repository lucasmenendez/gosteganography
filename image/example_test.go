// Copyright (c) 2023, Lucas Menendez <hi@lucasmenendez.me>
// See LICENSE for licensing information

package image

import (
	"fmt"
	"os"
)

func Example() {
	input, err := os.Open("./input.png")
	if err != nil {
		fmt.Println(err)
		return
	}
	// open the input image
	image, err := Read(input)
	if err != nil {
		fmt.Println(err)
		return
	}
	// hide a message, it returns the number of bits written
	expected := []byte("secret number: 1234")
	nbits, err := image.Hide(expected)
	if err != nil {
		fmt.Println(err)
		return
	}
	// get hided message using the number of bits
	got := image.Unhide(nbits)
	// store the output
	output, err := os.Create("./output.png")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer output.Close()

	if err := image.Write(output); err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(got))
	// Output: secret number: 1234
}
