package main

import (
	"flag"
	"fmt"
	"image"
	"image/png"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

var verbose bool

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

func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func Atoi(s string) int {
	if i, err := strconv.ParseInt(s, 0, 0); err != nil {
		log.Fatalf("could not parse %s as int: %v", s, err)
	} else {
		return int(i)
	}
	return -1
}

type Sensor struct {
	sensor        image.Point
	nearestBeacon image.Point
	clearRange    int
}

func Load(input []string, minX, maxX, minY, maxY *int, sensors *[]Sensor) {
	for _, line := range input {
		beacon := strings.Split(line, ": closest beacon is at ")[1]
		xBeacon := Atoi(strings.Split(beacon, ", ")[0][2:])
		yBeacon := Atoi(strings.Split(beacon, ", ")[1][2:])
		newBeacon := image.Point{X: xBeacon, Y: yBeacon}
		sensorInput := strings.Split(line, ": closest beacon is at ")[0]
		xSensor := Atoi(strings.Split(sensorInput[len("Sensor at "):], ", ")[0][2:])
		ySensor := Atoi(strings.Split(sensorInput[len("Sensor at "):], ", ")[1][2:])
		newSensor := image.Point{X: xSensor, Y: ySensor}
		dist := ManhattanDist(newBeacon, newSensor)
		sensor := Sensor{
			sensor:        newSensor,
			nearestBeacon: newBeacon,
			clearRange:    dist,
		}

		*minX = Min(Min(*minX, xBeacon), Min(*minX, xSensor-dist))
		*maxX = Max(Max(*maxX, xBeacon), Max(*maxX, xSensor+dist))
		*minY = Min(Min(*minY, yBeacon), Min(*minY, ySensor-dist))
		*maxY = Max(Max(*maxY, yBeacon), Max(*maxY, ySensor+dist))
		*sensors = append(*sensors, sensor)

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

func ManhattanDist(dest, src image.Point) int {
	return Abs(dest.X-src.X) + Abs(dest.Y-src.Y)
}

func (s *Sensor) InRangeOf(y int) bool {
	return s.sensor.Y == y || s.nearestBeacon.Y == y || (s.sensor.Y+s.clearRange >= y && s.sensor.Y-s.clearRange <= y)
}

func CountInvalidSpots(im *image.Alpha, minX, maxX, yLine int) int {
	count := 0
	for x := minX; x < maxX; x++ {
		if im.AlphaAt(x, yLine).A != 0 && im.AlphaAt(x, yLine).A != 'B' {
			count++
		}
	}
	return count
}

func PrintImage(im *image.Alpha) {
	for y := im.Rect.Min.Y; y < im.Rect.Max.Y; y++ {
		fmt.Printf("%3d ", y)
		for x := im.Rect.Min.X; x < im.Rect.Max.X; x++ {
			if im.AlphaAt(x, y).A != 0 {
				fmt.Printf("%c ", im.AlphaAt(x, y).A)
			} else {
				fmt.Printf("%c ", '.')
			}
		}
		fmt.Println()
	}
}

func CheckRow(sensors *[]Sensor, minX, y int, pix *[]uint8) int {
	count := 0
	for _, sensor := range *sensors {
		if sensor.InRangeOf(y) {
			if sensor.sensor.Y == y {
				(*pix)[sensor.sensor.X-minX] = 'S'
				count++
			} else if sensor.nearestBeacon.Y == y {
				(*pix)[sensor.nearestBeacon.X-minX] = 'B'
			}
			index := 0
			for x := sensor.sensor.X - sensor.clearRange; x < sensor.sensor.X+sensor.clearRange; x++ {
				index = x - minX
				pixel := &(*pix)[index]
				if *pixel != 'X' && *pixel != 'S' && *pixel != 'B' && ManhattanDist(image.Point{X: x, Y: y}, sensor.sensor) <= sensor.clearRange {
					*pixel = 'X'
					count++
				}
			}
		}
	}
	return count

}

func CountRows(y1, y2, minX, maxX int, counts *[]int, sensors *[]Sensor, wg *sync.WaitGroup) {
	for y := y1; y < y2; y++ {
		rect := image.Rectangle{Min: image.Point{X: minX, Y: y}, Max: image.Point{X: maxX + 1, Y: y + 1}}
		im := image.NewAlpha(rect)
		(*counts)[y] = CheckRow(sensors, minX, y, &im.Pix)
	}
	wg.Done()
}

func main() {
	flag.BoolVar(&verbose, "v", false, "Sets the output to verbose")
	flag.Parse()
	fmt.Println("Advent of Code - Day 15\n===============-=======")
	var input []string
	if data, err := os.ReadFile("data.txt"); err != nil {
		log.Fatalf("Failed to open file: %v", err)
	} else {
		input = strings.Split(string(data), "\n")
	}
	minX, maxX, minY, maxY := math.MaxInt, math.MinInt, math.MaxInt, math.MinInt
	var sensors []Sensor
	Load(input, &minX, &maxX, &minY, &maxY, &sensors)
	counts := make([]int, 4000000)

	startY := 2000000 - 500
	endY := 2000000 + 500
	divisor := 50
	//	for divisor := 1; divisor<= 100; divisor+=1 {
	start := time.Now()
	step := Max((endY-startY)/divisor, 1)
	var wg sync.WaitGroup
	for y := startY; y < endY; y += step {
		go CountRows(y, Min(y+step, endY), minX, maxX, &counts, &sensors, &wg)
		wg.Add(1)
	}
	wg.Wait()
	fmt.Printf("Total execution time with divisor of %d and step of %d:   \t%v\n", divisor, step, time.Since(start))
	fmt.Printf("Average execution time with divisor of %d and step of %d: \t%.2fms\n", divisor, step, float64(time.Since(start).Milliseconds())/float64(endY-startY))
	//	}

	fmt.Printf("y: %d (should be this with my input: %d)\n", counts[2000000], 4737567)
}
