// Copyright (c) 2023, Lucas Menendez <hi@lucasmenendez.me>
// See LICENSE for licensing information

package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/lucasmenendez/gosteganography"
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
	// parse args
	if len(args) != 3 {
		return fmt.Errorf("bad arguments")
	}
	input, message, output := args[0], args[1], args[2]
	// open input image
	inputFd, err := os.Open(input)
	if err != nil {
		return fmt.Errorf("error opening input image: %w", err)
	}
	defer inputFd.Close()
	// reading message file content
	secretMessage, err := os.ReadFile(message)
	if err != nil {
		return fmt.Errorf("error reading message file: %w", err)
	}
	// reading image from imput file
	img, err := gosteganography.Read(inputFd)
	if err != nil {
		return err
	}
	// hide message conntent and gets the number of bits written
	nbits, err := img.Hide(secretMessage)
	if err != nil {
		return err
	}
	// create output image file
	outputFd, err := os.Create(output)
	if err != nil {
		return fmt.Errorf("error creating output image file: %w", err)
	}
	defer outputFd.Close()
	// wirte the output image in the created file
	if err := img.Write(outputFd); err != nil {
		return fmt.Errorf("error writting output image file: %w", err)
	}
	// print the number of bits written
	fmt.Println("bits written: ", nbits)
	return nil
}

func unhide(args ...string) error {
	// parse args
	if len(args) != 3 {
		return fmt.Errorf("bad arguments")
	}
	input, message := args[0], args[1]
	nbits, err := strconv.Atoi(args[2])
	if err != nil {
		return fmt.Errorf("requires a valid number of bits")
	}
	// open input image
	inputFd, err := os.Open(input)
	if err != nil {
		return fmt.Errorf("error opening input image: %w", err)
	}
	defer inputFd.Close()
	// reading image from imput file
	img, err := gosteganography.Read(inputFd)
	if err != nil {
		return err
	}
	// retrieve the message content from the image
	secretMessage := img.Unhide(nbits)
	// write in the message file
	if err := os.WriteFile(message, secretMessage, os.ModePerm); err != nil {
		return fmt.Errorf("error writting message file: %w", err)
	}
	return nil
}
