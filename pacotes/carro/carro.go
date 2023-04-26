package carro

type Carro struct {
	marca  string
	modelo string
	km     float64
	Test   string
}

func (c *Carro) SetMarca(marca string) {
	c.marca = marca
}

func (c *Carro) SetModelo(modelo string) {
	c.modelo = modelo
}
func (c *Carro) SetKm(km float64) {
	c.km = km
}
func (c *Carro) GetMarca() string {
	return c.marca
}
func (c *Carro) GetModelo() string {
	return c.modelo
}
func (c *Carro) GetKm() float64 {
	return c.km
}
