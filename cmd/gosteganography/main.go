// Copyright (c) 2023, Lucas Menendez <hi@lucasmenendez.me>
// See LICENSE for licensing information

package main

import (
	"fmt"
	"os"
	"strconv"
)

var helpMessage = `
GoSteganography CLI helps to you to hide a message in a PNG image and unhide it 
from the output.

Usage:

	gosteganography <command> [arguments]

The commands are:

	hide	Hides the content of the secret file in a new copy of input image.
	unhide	Recovers the content of the secret from the input image.
`

var hideMessage = `
usage: gosteganography hide [input image] [secret file] [output image]

Hide command hides the provided message in the current Image. Before modify the 
image pixels, it checks if the provided message exceeds the maximun number of 
bytes that the Image can hide safely. If the limit is not exceeded, it encodes 
the message in its binary representation and hides it in the Image pixels.
`

var unhideMessage = `
usage: gosteganography unhide [input image] [secret file] [number of bits]

Unhide function writes the hidden message from the image provided using the 
number of bits provided. If the number of bits in the hidden message is not 
correct, the result will be wrong (truncated or badly formatted).
`

func main() {
	if len(os.Args) < 2 {
		fmt.Println("no command provided")
		fmt.Println(helpMessage)
		return
	}

	switch cmd, args := os.Args[1], os.Args[2:]; cmd {
	case "help":
		fmt.Println(helpMessage)
	case "hide":
		if err := hide(args...); err != nil {
			fmt.Println(err)
			fmt.Println(hideMessage)
		}
	case "unhide":
		if err := unhide(args...); err != nil {
			fmt.Println(err)
			fmt.Println(unhideMessage)
		}
	default:
		fmt.Println("unknown command")
		fmt.Println(helpMessage)
	}
}

func hide(args ...string) error {
	if len(args) != 3 {
		return fmt.Errorf("bad arguments")
	}

	input, message, output := args[0], args[1], args[2]
	fmt.Println(input, message, output)
	return nil
}

func unhide(args ...string) error {
	if len(args) != 3 {
		return fmt.Errorf("bad arguments")
	}

	input, message := args[0], args[1]
	nbits, err := strconv.Atoi(args[2])
	if err != nil {
		return fmt.Errorf("requires a valid number of bits")
	}
	fmt.Println(input, message, nbits)
	return nil
}
