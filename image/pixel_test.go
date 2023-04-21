// Copyright (c) 2023, Lucas Menendez <hi@lucasmenendez.me>
// See LICENSE for licensing information

package image

import (
	"os"
	"reflect"
	"testing"
)

const testImage = "./input.png"

func TestPixels_from(t *testing.T) {
	input, err := os.Open(testImage)
	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}
	img, err := Read(input)
	if err != nil {
		t.Errorf("expected nil, got %s", err)
	}

	ps := new(pixels).from(img)
	xnpixels := img.original.Bounds().Max.X - img.original.Bounds().Min.X
	ynpixels := img.original.Bounds().Max.Y - img.original.Bounds().Min.Y
	expectedPs := xnpixels * ynpixels
	if len(ps) != expectedPs {
		t.Errorf("expected %d, got %d", expectedPs, len((ps)))
	}

	expectedColor := img.original.At(ps[1432].x, ps[1432].y)
	if !reflect.DeepEqual(expectedColor, ps[1432].color) {
		t.Errorf("expected %v, got %v", expectedColor, ps[1432].color)
	}
}

func TestPixels_writebin(t *testing.T) {
	input, err := os.Open(testImage)
	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}
	img, err := Read(input)
	if err != nil {
		t.Errorf("expected nil, got %s", err)
	}

	bin := []uint{1, 0, 1, 1, 0}
	ogPs := new(pixels).from(img)

	img.pixels.writebin(bin)
	r0, g0, b0, _ := img.pixels[0].color.RGBA()
	br0, bg0, bb0 := num2bin(uint(r0)), num2bin(uint(g0)), num2bin(uint(b0))
	r1, g1, b1, _ := img.pixels[1].color.RGBA()
	br1, bg1, bb1 := num2bin(uint(r1)), num2bin(uint(g1)), num2bin(uint(b1))
	_, _, ogb1, _ := ogPs[1].color.RGBA()
	bogb1 := num2bin(uint(ogb1))

	if expected, got := bin[0], br0[len(br0)-1]; expected != got {
		t.Errorf("expected %d, got %d", expected, got)
	}
	if expected, got := bin[1], bg0[len(bg0)-1]; expected != got {
		t.Errorf("expected %d, got %d", expected, got)
	}
	if expected, got := bin[2], bb0[len(bb0)-1]; expected != got {
		t.Errorf("expected %d, got %d", expected, got)
	}
	if expected, got := bin[3], br1[len(br1)-1]; expected != got {
		t.Errorf("expected %d, got %d", expected, got)
	}
	if expected, got := bin[4], bg1[len(bg1)-1]; expected != got {
		t.Errorf("expected %d, got %d", expected, got)
	}
	if expected, got := bogb1[len(bogb1)-1], bb1[len(bb1)-1]; expected != got {
		t.Errorf("expected %d, got %d", expected, got)
	}
}

func TestPixels_readbin(t *testing.T) {
	input, err := os.Open(testImage)
	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}
	img, err := Read(input)
	if err != nil {
		t.Errorf("expected nil, got %s", err)
	}

	expected := []uint{1, 0, 1, 1, 0}
	img.pixels.writebin(expected)
	got := img.pixels.readbin(len(expected))
	if !reflect.DeepEqual(expected, got) {
		t.Errorf("expected %v, got %v", expected, got)
	}
}
