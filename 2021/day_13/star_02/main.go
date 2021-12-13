package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/derWhity/AdventOfCode/lib/input"
	"github.com/disintegration/imaging"
	"github.com/spf13/cast"
)

var (
	foldRegex = regexp.MustCompile(`^fold along ([xy])=([0-9]+)$`)
)

func failOnError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	items, err := input.ReadString(filepath.Join("..", "input.txt"), true)
	failOnError(err)
	// Determine the size of the canvas
	var width, height int
	for _, line := range items {
		if line == "" {
			// Folding section begins
			break
		}
		coords := strings.Split(line, ",")
		x := cast.ToInt(coords[0])
		if x > width {
			width = x
		}
		y := cast.ToInt(coords[1])
		if y > height {
			height = y
		}
	}
	if width%2 != 0 {
		width++
	}
	if height%2 != 0 {
		height++
	}
	canvas := image.NewNRGBA(image.Rect(0, 0, width+1, height+1))
	fmt.Printf("Base image size: %dx%d\n", canvas.Rect.Dx(), canvas.Rect.Dy())
	red := color.NRGBA{R: 255, G: 0, B: 0, A: 255}
	folding := false
	for _, line := range items {
		if !folding {
			// Coordinate plotting
			if line == "" {
				// Skip to the folding section
				folding = true
				continue
			}
			coords := strings.Split(line, ",")
			x := cast.ToInt(coords[0])
			y := cast.ToInt(coords[1])
			canvas.Set(x, y, red)
		} else if line != "" {
			fmt.Println(line)
			// Folding section
			var newCanvas *image.NRGBA
			var foldedPart *image.NRGBA
			match := foldRegex.FindStringSubmatch(line)
			pos := cast.ToInt(match[2])
			if match[1] == "x" {
				// Vertical split
				newCanvas = imaging.Crop(imaging.Clone(canvas), image.Rect(0, 0, pos, canvas.Rect.Dy()))
				foldedPart = imaging.Crop(imaging.Clone(canvas), image.Rect(pos+1, 0, canvas.Rect.Dx(), canvas.Rect.Dy()))
				foldedPart = imaging.FlipH(foldedPart)
			} else {
				// Horizontal split
				newCanvas = imaging.Crop(imaging.Clone(canvas), image.Rect(0, 0, canvas.Rect.Dx(), pos))
				foldedPart = imaging.Crop(imaging.Clone(canvas), image.Rect(0, pos+1, canvas.Rect.Dx(), canvas.Rect.Dy()))
				foldedPart = imaging.FlipV(foldedPart)
			}
			fmt.Printf("New: %dx%d | Fold: %dx%d | ", newCanvas.Rect.Dx(), newCanvas.Rect.Dy(), foldedPart.Rect.Dx(), foldedPart.Rect.Dy())
			canvas = imaging.Overlay(imaging.Clone(newCanvas), imaging.Clone(foldedPart), image.Pt(0, 0), 1.0)
			fmt.Printf("New image size: %dx%d\n", canvas.Rect.Dx(), canvas.Rect.Dy())
		}
	}
	// Output the image
	w, err := os.Create("out.png")
	failOnError(err)
	defer w.Close()
	png.Encode(w, canvas)
}
