// Copyright (c) 2023, Lucas Menendez <hi@lucasmenendez.me>
// See LICENSE for licensing information

package gosteganography

import (
	"fmt"
	"image"
	"image/png"
	"os"
)

type Image struct {
	original image.Image
	imgType  string
	pixels   pixels
}

func Open(path string) (*Image, error) {
	reader, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer reader.Close()
	original, itype, err := image.Decode(reader)
	if err != nil {
		return nil, err
	}
	switch itype {
	case "png":
		break
	// TODO: add support to other formats
	default:
		return nil, fmt.Errorf("no supported image format '%s'", itype)
	}
	newImage := &Image{
		original: original,
		imgType:  itype,
	}
	newImage.pixels = new(pixels).from(newImage)
	return newImage, nil
}

func (i *Image) BytesAvailable() int {
	// number of pixels multiplied by 3 color components (RGB) and divided by
	// the number of bits in a byte
	return len(i.pixels) * 3 / bitsPerByte
}

func (i *Image) Hide(msg []byte) (int, error) {
	if len(msg) > i.BytesAvailable() {
		return 0, fmt.Errorf("available bytes limit exceeded")
	}
	bmsg := encodeMessage(msg)
	i.pixels = i.pixels.writebin(bmsg)
	return len(bmsg), nil
}

func (i *Image) Unhide(nbits int) []byte {
	return decodeMessage(i.pixels.readbin(nbits))
}

func (i *Image) Save(path string) error {
	output, err := os.Create(path)
	if err != nil {
		return err
	}
	defer output.Close()
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
