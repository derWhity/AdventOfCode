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

type point struct {
	x int64
	y int64
}

type ship struct {
	position point
	waypoint point
}

func (s *ship) North(dist int64) {
	s.MoveWaypoint(north, dist)
}

func (s *ship) South(dist int64) {
	s.MoveWaypoint(south, dist)
}

func (s *ship) East(dist int64) {
	s.MoveWaypoint(east, dist)
}

func (s *ship) West(dist int64) {
	s.MoveWaypoint(west, dist)
}

func (s *ship) MoveWaypoint(head *heading, dist int64) {
	s.waypoint.x += head.x * dist
	s.waypoint.y += head.y * dist
	//fmt.Printf("Moving %s by %d - new position: [%d, %d]\n", head.txt, dist, s.x, s.y)
}

func (s *ship) TurnLeft(degrees int64) {
	times := degrees / 90
	for i := 0; i < int(times); i++ {
		newPt := point{
			x: -1 * s.waypoint.y,
			y: s.waypoint.x,
		}
		s.waypoint = newPt
	}
}

func (s *ship) TurnRight(degrees int64) {
	times := degrees / 90
	for i := 0; i < int(times); i++ {
		newPt := point{
			x: s.waypoint.y,
			y: -1 * s.waypoint.x,
		}
		s.waypoint = newPt
	}
}

func (s *ship) Forward(dist int64) {
	s.position.x += s.waypoint.x * dist
	s.position.y += s.waypoint.y * dist
	fmt.Printf("New position: [%d, %d]\n", s.position.x, s.position.y)
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
		waypoint: point{x: 10, y: 1},
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
	fmt.Printf("Manhattan distance: %d\n", int64(math.Abs(float64(ferry.position.x)))+int64(math.Abs(float64(ferry.position.y))))

}
