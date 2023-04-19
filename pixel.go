// Copyright (c) 2023, Lucas Menendez <hi@lucasmenendez.me>
// See LICENSE for licensing information

package gosteganography

import "image/color"

type pixel struct {
	x, y  int
	color color.Color
}

type pixels []*pixel

func (ps pixels) from(img *Image) pixels {
	bounds := img.original.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			ps = append(ps, &pixel{x, y, img.original.At(x, y)})
		}
	}
	return ps
}

func (ps pixels) writebin(bin []uint) pixels {
	bit, maxbits := 0, len(bin)
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
