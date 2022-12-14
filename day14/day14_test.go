package main

import (
	//    "github.com/stretchr/testify/assert"
	"image"
	"image/color"
	"testing"
)

var sharedIm *image.Alpha

func BenchmarkSimulateCaveIn(b *testing.B) {
	sharedIm = image.NewAlpha(image.Rect(450, 450, 550, 550))
	for x := 450; x < 550; x++ { // add floor
		sharedIm.SetAlpha(x, 530, color.Alpha{A: 'X'})
	}
	var oldImagePix []uint8
	for i := 0; i < len(sharedIm.Pix); i++ { // has to be a deep copy
		oldImagePix = append(oldImagePix, (sharedIm.Pix)[i])
	}
	for i := 0; i < b.N; i++ {
		SimulateCaveIn(sharedIm, 530)
	}
}

func BenchmarkSimulateSand(b *testing.B) {
	sharedIm = image.NewAlpha(image.Rect(0, 250, 0, 250))
	for i := 0; i < b.N; i++ {
		SimulateSand(sharedIm, 230)
	}
}
