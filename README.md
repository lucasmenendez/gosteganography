# gosteganography
Simple implementation of the LSB steganography algorithm in go, which uses the least significant bit (LSB) of each colour component (RGB) of each pixel of an image to hide a given message. 

It can be used in three ways:
 1. **As a library:** Check the documentation [here]().
 2. **As CLI:**  -> *coming soon*
 3. **As web app:** (using WASM) -> *coming soon*

### What is LSB Steganography?

Steganography is the practice of hiding secret information within an innocuous carrier medium, such as an image or a sound file, in a way that it is undetectable by human senses or analysis tools.

One common method of steganography is to use the least significant bit (LSB) technique, which involves replacing the least significant bit of each pixel in an image or the least significant sample in an audio file with a bit from the secret message. The LSBs are the bits that have the least impact on the overall value of the pixel or sample and therefore changing them slightly will not affect the quality of the media file.

For example, if the value of a pixel is `10101110` in binary, the least significant bit would be 0. If we want to hide a binary message "1101" within the pixel, we can replace the last 4 bits of the pixel with the message bits, resulting in a new pixel value of `10101111`.

By repeating this process for all pixels in the image, we can encode the entire secret message. To extract the message, the LSBs of each pixel are simply read and assembled together to reconstruct the original binary message.

However, it is worth noting that LSB steganography is a relatively simple and easily detectable method, and there are more advanced steganography techniques available that offer better security and reliability.

## Use & example 

### Installation
To include `gosteganography` as package dependency run:

```sh
go get github.com/lucasmenendez/gosteganography
```

### Use

#### Hide a message

1. Open image file
```go
    // open the input image
    image, err := gosteganography.OpenFile("./input.png")
	if err != nil {
		log.Fatal(err)
	}
```
2. Hide a message into the `gosteganograpy.Image`
```go
    // hide a message, it returns the number of bits writen
	secret := []byte("secret number: 1234")
	nbits, err := image.Hide(secret)
	if err != nil {
		log.Fatal(err)
	}
```
3. Write the `gosteganograpy.Image` with the hidden message into a file
```go
    // store the output
	if err := image.WriteFile("./output.png"); err != nil {
		log.Fatal(err)
	}
```
4. Share the image to someone with the number of bits writtem

#### Unhide the message

1. Open image file
```go
    // open the input image
    image, err := gosteganography.OpenFile("./input.png")
	if err != nil {
		log.Fatal(err)
	}
```

2. Unhide the message that the `gosteganograpy.Image` contains with the number of bits written
```go
   // get hided message using the number of bits
	secret := image.Unhide(nbits)
	fmt.Println(string(secret))
```

#### Full example
```go
    package main

    import (
        "fmt"

        "github.com/lucasmenendez/gosteganography"
    )

    func main() {
        // open the input image
        image, err := gosteganography.OpenFile("./input.png")
        if err != nil {
            log.Fatal(err)
        }
        // hide a message, it returns the number of bits writen
        secret := []byte("secret number: 1234")
        nbits, err := image.Hide(secret)
        if err != nil {
            log.Fatal(err)
        }
        // get hided message using the number of bits
        recovered := image.Unhide(nbits)
        fmt.Println(string(recovered))
        // store the output
        if err := image.WriteFile("./output.png"); err != nil {
            log.Fatal(err)
        }
    }
```