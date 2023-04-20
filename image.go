// Copyright (c) 2023, Lucas Menendez <hi@lucasmenendez.me>
// See LICENSE for licensing information

// gosteganography is a simple implementation of the LSB steganography algorithm
// in go, which uses the least significant bit (LSB) of each colour component
// (RGB) of each pixel of an image to hide a given message.
package gosteganography

import (
	"image"
	"image/png"
	"io"
)

// Image struct abstracts the needed parameters to hide a message into it. It
// contains the original image.Image instance, the image type and the list of
// pixels.
type Image struct {
	original image.Image
	imgType  string
	pixels   pixels
}

// Read function instances a new Image struct from the io.Reader provided. It
// reads the original image and decodes it. It then initialises the Image
// pixels from the original image.Image.
func Read(input io.Reader) (*Image, error) {
	original, itype, err := image.Decode(input)
	if err != nil {
		return nil, wrap(ErrDecodingImage, "input reader err: %w", err)
	}
	switch itype {
	case "png":
		break
	// TODO: add support to other formats
	default:
		return nil, wrap(ErrFormatNotSupported, "format provided: '%s'", itype)
	}
	newImage := &Image{
		original: original,
		imgType:  itype,
	}
	newImage.pixels = new(pixels).from(newImage)
	return newImage, nil
}

// Write function stores the current Image in io.Writer provided. It creates a
// new image.Image with the same dimensions of the original image and writes
// every Image pixel in the new one. Then write the new image.Image in the
// io.Writer provided.
func (i *Image) Write(output io.Writer) error {
	var newImage = image.NewRGBA(i.original.Bounds())
	for _, pix := range i.pixels {
		newImage.Set(pix.x, pix.y, pix.color)
	}
	switch i.imgType {
	// TODO: add support to other formats
	case "png":
		return png.Encode(output, newImage)
	}
	return nil
}

// BytesAvailable returns the number of bytes the current image can safely hold.
// Actually, the maximun number of bytes that an Image can hold safely is
// calculated by multiplying the number of pixels by the number of colour
// components (three) and dividing by the number of bits in a byte.
func (i *Image) BytesAvailable() int {
	return len(i.pixels) * 3 / bitsPerByte
}

// Hide function hides the provided message in the current Image. Before modify
// the Image pixels, it checks if the provided message exceeds the maximun
// number of bytes that the Image can hide safely. If the limit is not exceeded,
// it encodes the message in its binary representation and hides it in the Image
// pixels. If the limit is exceeded, it returns an ErrBytesLimitExceeded error.
func (i *Image) Hide(msg []byte) (int, error) {
	if len(msg) > i.BytesAvailable() {
		return 0, ErrBytesLimitExceeded
	}
	bmsg := encodeMessage(msg)
	i.pixels = i.pixels.writebin(bmsg)
	return len(bmsg), nil
}

// Unhide function returns the hide message in the current Image with the length
// given as the `nbits` argument. If the number of bits in the hidden message is
// not correct, the result will be wrong (truncated or badly formatted).
func (i *Image) Unhide(nbits int) []byte {
	return decodeMessage(i.pixels.readbin(nbits))
}
