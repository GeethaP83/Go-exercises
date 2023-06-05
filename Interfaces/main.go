package main

import (
	"pratice/box"
	"pratice/cube"
	"pratice/shape"
	"pratice/sphere"
)

func main() {

	c := cube.Cube{
		Length: 7,
	}

	b := box.Box{
		Length: 5.5,
		Width:  5.5,
		Height: 7.7,
	}

	s := sphere.Sphere{
		Radius: 7.14,
	}

	shape.CalculateVolume(c, "Cube")
	shape.CalculateVolume(b, "Box")
	shape.CalculateVolume(s, "Sphere")
}
