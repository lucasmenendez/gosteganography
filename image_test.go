// Copyright (c) 2023, Lucas Menendez <hi@lucasmenendez.me>
// See LICENSE for licensing information

package gosteganography

import (
	"os"
	"reflect"
	"testing"
)

func TestEnd2End(t *testing.T) {
	input, err := os.Open("./input.png")
	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}
	image, err := Read(input)
	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}

	expected := []byte("secret number: 1234")
	nbits, err := image.Hide(expected)
	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}
	got := image.Unhide(nbits)
	if !reflect.DeepEqual(expected, got) {
		t.Errorf("expected %v, got %v", expected, got)
	}
	output, err := os.Create("./temp_output.png")
	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}
	defer output.Close()

	defer os.Remove("./temp_output.png")
	if err := image.Write(output); err != nil {
		t.Errorf("expected nil, got %v", err)
	}
}
