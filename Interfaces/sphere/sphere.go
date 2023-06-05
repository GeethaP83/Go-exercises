package sphere

import (
	"math"
)

type Sphere struct {
	Radius float64
}

func (s Sphere) Volume() float64 {
	return (4 * math.Pi * math.Pow(s.Radius, float64(3))) / 3

}
