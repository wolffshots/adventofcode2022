package main

import (
	"flag"
	"fmt"
	"image"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
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
		dist := ManhattanDist(&newBeacon, &newSensor)
		sensor := Sensor{
			sensor:        newSensor,
			nearestBeacon: newBeacon,
			clearRange:    dist,
		}

		*minX = Min(Min(*minX, xBeacon), Min(*minX, xSensor-dist))
		*maxX = Max(Max(*maxX, xBeacon), Max(*maxX, xSensor+dist))
		*minY = Min(Min(*minY, yBeacon), Min(*minY, ySensor))
		*maxY = Max(Max(*maxY, yBeacon), Max(*maxY, ySensor))
		*sensors = append(*sensors, sensor)

	}
}

func ManhattanDist(dest, src *image.Point) int {
	return Abs(dest.X-src.X) + Abs(dest.Y-src.Y)
}

func IsNotPresent(c*image.Point, s*Sensor)bool{
    if ManhattanDist(c, &s.sensor) <= s.clearRange{
        return true
    }
    return false
}

func IsInRangeOf(c*image.Point, s*Sensor)bool{
    if ManhattanDist(c, &s.sensor) <= s.clearRange && !(c.X == s.nearestBeacon.X && c.Y == s.nearestBeacon.Y ) && !(c.X == s.sensor.X && c.Y == s.sensor.Y ) {
        return true
    }
    return false
}

func PointsOutsideOfSensor(sensor*Sensor) []image.Point{
    var points []image.Point

    points = append(points, image.Point{X: sensor.sensor.X, Y: sensor.sensor.Y-sensor.clearRange-1})
    points = append(points, image.Point{X: sensor.sensor.X, Y: sensor.sensor.Y+sensor.clearRange+1})
    for y := sensor.sensor.Y-sensor.clearRange; y <= sensor.sensor.Y+sensor.clearRange; y++ {
        diff := sensor.clearRange - Abs(sensor.sensor.Y - y) + 1
        points = append(points, image.Point{X: sensor.sensor.X-diff, Y: y})
        points = append(points, image.Point{X: sensor.sensor.X+diff, Y: y})
    }
//    fmt.Println(sensor.sensor.X, sensor.sensor.Y, sensor.clearRange, points)
    return points
}

func FindOpenPoint(points*[][]image.Point, sensors *[]Sensor, size int)image.Point{
    for _, pointForSensor := range *points {
        for _, point := range pointForSensor {
            if point.X <= size &&point.X >= 0 && point.Y <= size &&point.Y >= 0 {
                inRangeOfAtLeastOne := false
                for _, sensor := range *sensors {
                    if ManhattanDist(&point, &sensor.sensor) <= sensor.clearRange{
                        inRangeOfAtLeastOne = true
                    }
                }
                if !inRangeOfAtLeastOne{
                    fmt.Println(point)
                    return point
                }
            }
        }
    }
    return image.Point{X:-1, Y:-1}
}

func main() {
	flag.BoolVar(&verbose, "v", false, "Sets the output to verbose")
	flag.Parse()
	fmt.Println("Advent of Code - Day 15\n<=============>-<=====>")
	var input []string
	if data, err := os.ReadFile("data.txt"); err != nil {
		log.Fatalf("Failed to open file: %v", err)
	} else {
		input = strings.Split(string(data), "\n")
	}
	minX, maxX, minY, maxY := math.MaxInt, math.MinInt, math.MaxInt, math.MinInt
	var sensors []Sensor
	Load(input, &minX, &maxX, &minY, &maxY, &sensors)
    const size = 4000000
    const lineNo = size/2
    count:=0
    start:=time.Now()
        count =0
        for x := minX; x < maxX; x++ {
            inRange := false
            for _, sensor := range sensors {
                if IsInRangeOf(&image.Point{X: x, Y: lineNo}, &sensor){
                    inRange = true
                }
            }
            if inRange{
                count++
            }
        }
        fmt.Println("count:",count)
    fmt.Println("that took about", time.Since(start))
    var pointsToCheck [][]image.Point
    start=time.Now()
    for _, sensor := range sensors {
        pointsToCheck = append(pointsToCheck, PointsOutsideOfSensor(&sensor))
    }
    fmt.Println("PointsOutsideOfSensor's took about", time.Since(start))
    start=time.Now()
    origin := FindOpenPoint(&pointsToCheck, &sensors, size)
    fmt.Println("FindOpenPoint took about", time.Since(start))
    tuningFreq:= origin.X*4000000+origin.Y
    fmt.Println("Tuning freq:",tuningFreq)
}
