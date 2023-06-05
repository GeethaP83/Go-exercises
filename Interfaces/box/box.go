package box

type Box struct {
	Length float64
	Width  float64
	Height float64
}

func (b Box) Volume() float64 {
	return b.Length * b.Width * b.Height
}
