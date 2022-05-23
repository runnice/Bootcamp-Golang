package circ

type Circulo struct {
	raio float64
}

func (c *Circulo) SetarRaio(raio float64) {
	c.raio = raio
}

func (c Circulo) GetRaio() float64 {
	return c.raio
}
