package main

import (
	"github.com/stretchr/testify/assert"
	"image"
	"testing"
)

func TestManhattanDist(t *testing.T) {
	dests := []image.Point{
		{2, 18},
	}
	srcs := []image.Point{
		{-2, 15},
	}
	wants := []int{
		7,
	}
	for testIndex, want := range wants {
		result := ManhattanDist(dests[testIndex], srcs[testIndex])
		assert.Equal(t, want, result)
		result = ManhattanDist(srcs[testIndex], dests[testIndex])
		assert.Equal(t, want, result)
	}
}
