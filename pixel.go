// Copyright (c) 2023, Lucas Menendez <hi@lucasmenendez.me>
// See LICENSE for licensing information

package gosteganography

import "image/color"

// pixel struct contains the coordenates of the pixel and its color
// (color.Color) reference
type pixel struct {
	x, y  int
	color color.Color
}

// pixels is the type abstraction for a list of pointers to pixels
type pixels []*pixel

// from function fills the current pixels with the pixels of the provided Image.
// It also returns the filled pixels.
func (ps pixels) from(img *Image) pixels {
	// get original image bounds from Image provided
	bounds := img.original.Bounds()
	// iterates over every original image pixels storing their colors and
	// coordenates
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			ps = append(ps, &pixel{x, y, img.original.At(x, y)})
		}
	}
	return ps
}

// writebin function writes the provided binary slice  into current pixels. It
// iterates over the pixels writing the bits three by three, one per pixel color
// component. To write a bit, it takes the binary representation of each color
// component and replaces the less important bit with each bit.
func (ps pixels) writebin(bin []uint) pixels {
	// create some variables to keep track of the bits written, to avoid writing
	// unnecessary pixels or components, only the required number of bits.
	bit, maxbits := 0, len(bin)
	// iterate over the pixels, changing the same number of pixel components
	// as the number of bits of the binary slice. The rest of pixels remain
	// the same.
	for n, p := range ps {
		r, g, b, a := p.color.RGBA()
		if bit < maxbits {
			br := num2bin(uint(r))
			br[len(br)-1] = bin[bit]
			r = uint32(bin2num(br))
		}
		if bit+1 < maxbits {
			bg := num2bin(uint(g))
			bg[len(bg)-1] = bin[bit+1]
			g = uint32(bin2num(bg))
		}
		if bit+2 < maxbits {
			bb := num2bin(uint(b))
			bb[len(bb)-1] = bin[bit+2]
			b = uint32(bin2num(bb))
		}
		ps[n] = &pixel{
			x: p.x, y: p.y,
			color: color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)},
		}
		bit += 3
	}
	return ps
}

// readbin functions returns the binary representation of a hidden message into
// the current pixels. It returns the less important bits of the n first pixel
// colors components.
func (ps pixels) readbin(nbits int) []uint {
	bin := []uint{}
	currentBit := 0
	for _, p := range ps {
		r, g, b, _ := p.color.RGBA()
		if currentBit < nbits {
			br := num2bin(uint(r))
			bin = append(bin, br[len(br)-1])
		}
		if currentBit+1 < nbits {
			bg := num2bin(uint(g))
			bin = append(bin, bg[len(bg)-1])
		}
		if currentBit+2 < nbits {
			bb := num2bin(uint(b))
			bin = append(bin, bb[len(bb)-1])
		}
		currentBit += 3
	}
	return bin
}
