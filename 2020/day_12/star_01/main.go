package main

import (
	"fmt"
	"math"
	"path/filepath"
	"strconv"

	"github.com/derWhity/AdventOfCode/lib/input"
)

type heading struct {
	x   int64
	y   int64
	txt string
	l   *heading // Left heading
	r   *heading // Right heading
	b   *heading // Backwards heading
}

var (
	north = &heading{x: 0, y: 1, txt: "north"}
	east  = &heading{x: 1, y: 0, txt: "east"}
	south = &heading{x: 0, y: -1, txt: "south"}
	west  = &heading{x: -1, y: 0, txt: "west"}
)

func init() {
	north.l = west
	north.r = east
	north.b = south
	east.l = north
	east.r = south
	east.b = west
	south.l = east
	south.r = west
	south.b = north
	west.l = south
	west.r = north
	west.b = east
}

type ship struct {
	x       int64
	y       int64
	heading *heading
}

func (s *ship) North(dist int64) {
	s.Move(north, dist)
}

func (s *ship) South(dist int64) {
	s.Move(south, dist)
}

func (s *ship) East(dist int64) {
	s.Move(east, dist)
}

func (s *ship) West(dist int64) {
	s.Move(west, dist)
}

func (s *ship) Move(head *heading, dist int64) {
	s.x += head.x * dist
	s.y += head.y * dist
	//fmt.Printf("Moving %s by %d - new position: [%d, %d]\n", head.txt, dist, s.x, s.y)
}

func (s *ship) TurnLeft(degrees int64) {
	fmt.Printf("Turning left by %d° - ", degrees)
	switch degrees {
	// Only increments of 90°
	case 90:
		s.heading = s.heading.l
	case -90:
		s.heading = s.heading.r
	case 180:
		s.heading = s.heading.b
	case 270:
		s.heading = s.heading.r
	}
	// Everything else leads to keeping the heading
	fmt.Printf("Ship is now facing %s\n", s.heading.txt)
}

func (s *ship) TurnRight(degrees int64) {
	fmt.Printf("Turning right by %d° - ", degrees)
	switch degrees {
	// Only increments of 90°
	case 90:
		s.heading = s.heading.r
	case -90:
		s.heading = s.heading.l
	case 180:
		s.heading = s.heading.b
	case 270:
		s.heading = s.heading.l
	}
	// Everything else leads to keeping the heading
	fmt.Printf("Ship is now facing %s\n", s.heading.txt)
}

func (s *ship) Forward(dist int64) {
	s.Move(s.heading, dist)
}

func failOnError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	lines, err := input.ReadString(filepath.Join("..", "input.txt"), true)
	failOnError(err)
	ferry := ship{
		heading: east,
	}
	for _, instruction := range lines {
		command := instruction[0:1]
		distance, err := strconv.ParseInt(instruction[1:], 10, 64)
		failOnError(err)
		switch command {
		case "N":
			ferry.North(distance)
		case "E":
			ferry.East(distance)
		case "S":
			ferry.South(distance)
		case "W":
			ferry.West(distance)
		case "L":
			ferry.TurnLeft(distance)
		case "R":
			ferry.TurnRight(distance)
		case "F":
			ferry.Forward(distance)
		default:
			panic("Unknown command")
		}
	}
	fmt.Printf("Manhattan distance: %d\n", int64(math.Abs(float64(ferry.x)))+int64(math.Abs(float64(ferry.y))))

}
