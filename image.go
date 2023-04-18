package gosteganography

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"

	_ "image/gif"
	_ "image/jpeg"
)

type pixel struct {
	x, y  int
	color color.Color
}

type Img struct {
	original image.Image
	bounds   image.Rectangle
	imgType  string
	pixels   []*pixel
}

func Open(path string) (*Img, error) {
	reader, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	original, itype, err := image.Decode(reader)
	if err != nil {
		return nil, err
	}
	res := &Img{
		original: original,
		bounds:   original.Bounds(),
		imgType:  itype,
		pixels:   []*pixel{},
	}

	for y := res.bounds.Min.Y; y < res.bounds.Max.Y; y++ {
		for x := res.bounds.Min.X; x < res.bounds.Max.X; x++ {
			res.pixels = append(res.pixels, &pixel{x, y, original.At(x, y)})
		}
	}

	return res, err
}

func (i *Img) BytesAvailable() int {
	return len(i.pixels) * 3 / 8 // 3 color components divided by number of bits in a byte
}

func (i *Img) Hide(msg []byte) (int, error) {
	if len(msg) > i.BytesAvailable() {
		return 0, fmt.Errorf("available bytes limit exceeded")
	}
	bmsg := encodeMessage(msg)
	newPixels := make([]*pixel, len(i.pixels))

	idx := 0
	mgsLen := len(bmsg)
	for n, p := range i.pixels {
		r, g, b, a := p.color.RGBA()

		if idx < mgsLen {
			br := num2bin(uint(r))
			br[len(br)-1] = bmsg[idx]
			r = uint32(bin2num(br))
		}
		if idx+1 < mgsLen {
			bg := num2bin(uint(g))
			bg[len(bg)-1] = bmsg[idx+1]
			g = uint32(bin2num(bg))
		}
		if idx+2 < mgsLen {
			bb := num2bin(uint(b))
			bb[len(bb)-1] = bmsg[idx+2]
			b = uint32(bin2num(bb))
		}

		newPixels[n] = &pixel{
			x: p.x, y: p.y,
			color: color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)},
		}
		idx += 3
	}
	i.pixels = newPixels
	return len(bmsg), nil
}

func (i *Img) Unhide(nbits int) []byte {
	bmsg := []uint{}

	currentBit := 0
	for _, p := range i.pixels {
		r, g, b, _ := p.color.RGBA()

		if currentBit < nbits {
			br := num2bin(uint(r))
			bmsg = append(bmsg, br[len(br)-1])
		}
		if currentBit+1 < nbits {
			bg := num2bin(uint(g))
			bmsg = append(bmsg, bg[len(bg)-1])
		}
		if currentBit+2 < nbits {
			bb := num2bin(uint(b))
			bmsg = append(bmsg, bb[len(bb)-1])
		}
		currentBit += 3
	}

	return decodeMessage(bmsg)
}

func (i *Img) Save(path string) error {
	var newImage = image.NewRGBA(i.bounds)
	for _, pix := range i.pixels {
		newImage.Set(pix.x, pix.y, pix.color)
	}

	output, err := os.Create(path)
	if err != nil {
		return err
	}

	switch i.imgType {
	case "png":
		return png.Encode(output, newImage)
	}

	return nil
}
