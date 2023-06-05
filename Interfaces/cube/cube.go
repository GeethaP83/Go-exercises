package cube

type Cube struct {
	Length float64
}

func (c Cube) Volume() float64 {
	return c.Length * c.Length * c.Length
}
