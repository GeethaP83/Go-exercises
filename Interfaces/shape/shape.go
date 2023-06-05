package shape

import (
	"fmt"
)

type OfStructure interface {
	Volume() float64
}

func CalculateVolume(kind OfStructure, called string) {
	fmt.Printf("The Volume calculated for our %s is: %f \n", called, kind.Volume())
}
