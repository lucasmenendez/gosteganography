package gosteganography

import (
	"reflect"
	"testing"
)

func Test_num2bin(t *testing.T) {
	t.Run("num2bin test 1: 234", func(t *testing.T) {
		input := uint(234)
		expected := []uint{1, 1, 1, 0, 1, 0, 1, 0}
		if got := num2bin(input); !reflect.DeepEqual(expected, got) {
			t.Errorf("expected %v, got %v", expected, got)
		}
	})

	t.Run("num2bin test 1: 12", func(t *testing.T) {
		input := uint(12)
		expected := []uint{1, 1, 0, 0}
		if got := num2bin(input); !reflect.DeepEqual(expected, got) {
			t.Errorf("expected %v, got %v", expected, got)
		}
	})
}

func Test_bin2num(t *testing.T) {
	t.Run("bin2num test 1: 234", func(t *testing.T) {
		input := []uint{1, 1, 1, 0, 1, 0, 1, 0}
		expected := uint(234)
		if got := bin2num(input); expected != got {
			t.Errorf("expected %v, got %v", expected, got)
		}
	})

	t.Run("bin2num test 2: 12", func(t *testing.T) {
		input := []uint{1, 1, 0, 0}
		expected := uint(12)
		if got := bin2num(input); expected != got {
			t.Errorf("expected %v, got %v", expected, got)
		}
	})

	t.Run("bin2num test 3: 234 padded", func(t *testing.T) {
		input := []uint{0, 0, 1, 1, 1, 0, 1, 0, 1, 0}
		expected := uint(234)
		if got := bin2num(input); expected != got {
			t.Errorf("expected %v, got %v", expected, got)
		}
	})

	t.Run("bin2num test 4: 12 padded", func(t *testing.T) {
		input := []uint{0, 0, 1, 1, 0, 0}
		expected := uint(12)
		if got := bin2num(input); expected != got {
			t.Errorf("expected %v, got %v", expected, got)
		}
	})
}

func Test_resize(t *testing.T) {
	t.Run("padding", func(t *testing.T) {
		input := []uint{1, 1, 0, 0}
		expected := []uint{0, 0, 1, 1, 0, 0}
		if got := resize(input, 6); !reflect.DeepEqual(expected, got) {
			t.Errorf("expected %v, got %v", expected, got)
		}
	})
	t.Run("truncate", func(t *testing.T) {
		input := []uint{1, 1, 0, 0}
		expected := []uint{1, 1}
		if got := resize(input, 2); !reflect.DeepEqual(expected, got) {
			t.Errorf("expected %v, got %v", expected, got)
		}
	})
	t.Run("same", func(t *testing.T) {
		input := []uint{1, 1, 0, 0}
		expected := []uint{1, 1, 0, 0}
		if got := resize(input, 4); !reflect.DeepEqual(expected, got) {
			t.Errorf("expected %v, got %v", expected, got)
		}
	})
}

func Test_encodeAndDecode(t *testing.T) {
	expected := []byte("hello world!")
	got := decodeMessage(encodeMessage(expected))
	if !reflect.DeepEqual(expected, got) {
		t.Errorf("expected %v, got %v", expected, got)
	}
}
