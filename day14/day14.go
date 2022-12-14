package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

var verbose bool

const MaxIterations = 30000

func Atoi(s string) int {
	if i, err := strconv.ParseInt(s, 0, 0); err != nil {
		log.Fatalf("could not parse %s as int: %v", s, err)
	} else {
		return int(i)
	}
	return -1
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Format(im *image.Alpha, addIndices bool) string {
	formatted := ""
	for y := im.Rect.Min.Y; y <= im.Rect.Max.Y; y++ {
		if y == im.Rect.Min.Y && addIndices {
			formatted = fmt.Sprintf("%s%-3s", formatted, "")
			for l := 0; l < 3; l++ {
				for x := im.Rect.Min.X + 1; x <= im.Rect.Max.X; x++ {
					formatted = fmt.Sprintf("%s%s", formatted, fmt.Sprint(x)[l:l+1])
				}
				if l != 2 {
					formatted = fmt.Sprintf("%s%s%3s", formatted, "\n", "")
				}
			}

		} else {
			for x := im.Rect.Min.X; x <= im.Rect.Max.X; x++ {
				if x == im.Rect.Min.X && addIndices {
					formatted = fmt.Sprintf("%s%3d", formatted, y)
				} else {
					switch im.AlphaAt(x, y).A {
					case 'X':
						{
							formatted = fmt.Sprintf("%s%s", formatted, "X")
						}
					case 'o':
						{
							formatted = fmt.Sprintf("%s%s", formatted, "o")
						}
					case '+':
						{
							formatted = fmt.Sprintf("%s%s", formatted, "+")
						}
					case '~':
						{
							formatted = fmt.Sprintf("%s%s", formatted, "~")
						}
					case 0:
						{
							formatted = fmt.Sprintf("%s%s", formatted, " ")
						}
					default:
						{
							formatted = fmt.Sprintf("%s%s", formatted, "?")
						}
					}
				}
			}
		}
		formatted = fmt.Sprintf("%s%s", formatted, "\n")
	}
	return formatted
}

func SimulateSand(im *image.Alpha, floor int) bool {
	changed := false
	for y := im.Rect.Min.Y; y <= im.Rect.Max.Y; y++ {
		for x := im.Rect.Min.X; x <= im.Rect.Max.X; x++ {
			if im.AlphaAt(x, y).A == 'o' || im.AlphaAt(x, y).A == '~' {
				if im.AlphaAt(x, y+1).A == 0 && y != floor {
					im.Set(x, y+1, color.Alpha{A: im.AlphaAt(x, y).A})
					if im.AlphaAt(x, y).A != '~' {
						im.Set(x, y, color.Alpha{A: 0})
                        changed = !(y == im.Rect.Max.Y-1 && (x == im.Rect.Max.X || x == im.Rect.Min.X))
					}
				} else if im.AlphaAt(x-1, y+1).A == 0 && y != floor {
					im.Set(x-1, y+1, color.Alpha{A: im.AlphaAt(x, y).A})
					if im.AlphaAt(x, y).A != '~' {
						im.Set(x, y, color.Alpha{A: 0})
                        changed = !(y == im.Rect.Max.Y-1 && (x == im.Rect.Max.X || x == im.Rect.Min.X))
					}
				} else if im.AlphaAt(x+1, y+1).A == 0 && y != floor {
					im.Set(x+1, y+1, color.Alpha{A: im.AlphaAt(x, y).A})
					if im.AlphaAt(x, y).A != '~' {
						im.Set(x, y, color.Alpha{A: 0})
                        changed = !(y == im.Rect.Max.Y-1 && (x == im.Rect.Max.X || x == im.Rect.Min.X))
					}
				}
			}
		}
	}
	return changed
}

func DeepCopyPix(dest *[]uint8, src *image.Alpha) {
	if len(*dest) != len(src.Pix) {
		log.Fatal("trying to deep copy unequal slices")
	}
	for i := 0; i < len(src.Pix); i++ {
		(*dest)[i] = src.Pix[i]
	}
}

func Equals(one, two *[]uint8) bool {
	if len(*one) != len(*two) {
		log.Fatalf("shouldn't be comparing different sizes")
		return false
	}
	for index := len(*two) - 1; index >= 0; index-- {
		if (*one)[index] != (*two)[index] {
			return false
		}
	}
	return true
}

func SimulateCaveIn(im *image.Alpha, floor int) int {
	if floor > 0 {
		for x := im.Bounds().Min.X - 500; x < im.Bounds().Max.X+500; x++ { // add floor
			im.Set(x, floor, color.Alpha{A: 'X'})
		}
	}
	var oldImagePix []uint8
	for i := 0; i < len(im.Pix); i++ { // has to be a deep copy
		oldImagePix = append(oldImagePix, (im.Pix)[i])
	}
	for iteration := 0; iteration < MaxIterations; iteration++ {
//		DeepCopyPix(&oldImagePix, im)
		im.Set(500, 0, color.Alpha{A: 'o'})
		changed := SimulateSand(im, floor)
		if !changed || im.AlphaAt(500, 0).A == 'o' /*|| Equals(&im.Pix, &oldImagePix)*/ {
			if im.AlphaAt(500, 0).A == 'o' {
				iteration++ // saves on the final check
			} else {
				im.Set(500, 0, color.Alpha{A: '~'})
				SimulateSand(im, floor)
			}
			im.SetAlpha(500, 0, color.Alpha{A: '+'})
			return iteration
		}
	}
	log.Fatalf("steps == max iterations (%d) therefore, sand never finished settling", MaxIterations)
	return -1
}

func Load(input []string, minX, maxX, minY, maxY *int, rock *[]image.Point) {
	for _, line := range input {
		vs := strings.Split(line, " -> ")
		for vertIndex := 0; vertIndex < len(vs)-1; vertIndex++ {
			v1 := strings.Split(vs[vertIndex], ",")
			x1 := Atoi(v1[0])
			y1 := Atoi(v1[1])
			v2 := strings.Split(vs[vertIndex+1], ",")
			x2 := Atoi(v2[0])
			y2 := Atoi(v2[1])
			*minX = Min(Min(*minX, x1), Min(*minX, x2))
			*maxX = Max(Max(*maxX, x1), Max(*maxX, x2))
			*minY = Min(Min(*minY, y1), Min(*minY, y2))
			*maxY = Max(Max(*maxY, y1), Max(*maxY, y2))
			for y := Min(y1, y2); y <= Max(y1, y2); y++ {
				for x := Min(x1, x2); x <= Max(x1, x2); x++ {
					*rock = append(*rock, image.Point{X: x, Y: y})
				}
			}
		}
	}
}

func SetRock(im *image.Alpha, rock *[]image.Point) {
	for _, point := range *rock {
		im.SetAlpha(point.X, point.Y, color.Alpha{A: 'X'})
	}
}

func WriteImage(path string, im *image.Alpha) {
	if out, err := os.Create(path); err != nil {
		log.Fatalf("couldn't write image: %v", err)
	} else {
		if err := png.Encode(out, im); err != nil {
			log.Fatalf("couldn't write image: %v", err)
		} else {
			if err := out.Close(); err != nil {
				log.Fatalf("couldn't close image: %v", err)
			}
		}
	}
}

func ShowCave(prefix string, im *image.Alpha) {
	if verbose {
		if prefix != "" {
			fmt.Printf("%s:\n%s", prefix, Format(im, true))
		} else {
			fmt.Printf("%s", Format(im, true))
		}
	}
}

func main() {
	flag.BoolVar(&verbose, "v", false, "Sets the output to verbose")
	flag.Parse()
	fmt.Println("Advent of Code - Day 14\n===============-=======")
	var input []string
	if data, err := os.ReadFile("data.txt"); err != nil {
		log.Fatalf("Failed to open file: %v", err)
	} else {
		input = strings.Split(string(data), "\n")
	}

	minX, maxX, minY, maxY := 500, 500, 0, 0

	var rock []image.Point

	start := time.Now()
	Load(input, &minX, &maxX, &minY, &maxY, &rock)
	elapsed0 := time.Since(start)
	floor := math.MaxInt
	fmt.Println("\nPart 1:")
	im1 := image.NewAlpha(image.Rect(minX-1, minY-1, maxX+1, maxY+1))
	SetRock(im1, &rock)

	ShowCave("Start:", im1)
	start = time.Now()
	steps1 := SimulateCaveIn(im1, -1)
	elapsed1 := time.Since(start)
	ShowCave("End:", im1)

	fmt.Println("\nPart 2:")
	floor = maxY + 2
	// these changes to X were estimated and trial and errored, would like to make them dynamic
	im2 := image.NewAlpha(image.Rect(minX-150, minY-1, maxX+137, maxY+3))
	SetRock(im2, &rock)

	ShowCave("Start:", im2)
	start = time.Now()
	steps2 := SimulateCaveIn(im2, floor)
	elapsed2 := time.Since(start)
	ShowCave("End:", im2)

	// write out pngs for inspection
	WriteImage("result1.png", im1)
	WriteImage("result2.png", im2)

	fmt.Println("Results:")
	fmt.Printf("\tData loading and converting took %s\n", elapsed0)
	fmt.Println("\tIt took", steps1, "units to stabilise part one (should be", 755, "with my example input)") // 755
	fmt.Printf("\tSimulate function for part one took %s\n", elapsed1)
	fmt.Println("\tIt took", steps2, "units to stabilise part two (should be", 29805, "with my example input)") // 29805
	fmt.Printf("\tSimulate function for part one took %s\n", elapsed2)
}
