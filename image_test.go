// Copyright (c) 2023, Lucas Menendez <hi@lucasmenendez.me>
// See LICENSE for licensing information

package gosteganography

import (
	"os"
	"reflect"
	"testing"
)

func TestEnd2End(t *testing.T) {
	image, err := OpenFile("./input.png")
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
	defer os.Remove("./temp_output.png")
	if err := image.WriteFile("./temp_output.png"); err != nil {
		t.Errorf("expected nil, got %v", err)
	}

}
