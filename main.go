package area

import "math"

// Pi é uma proporção numerica definida pela relação entre
// o perimetro de uma circunferencia e seu diametro
const Pi = 3.1416

// Circ é responsavel por calcular a area da circunferencia
func Circ(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}

// Rect é responsavel por calcular a area de um retangulo
func Rect(base, altura float64) float64 {
	return base * altura
}

// Nao é visivel
func _TrianguloEq(base, altura float64) float64 {
	return (base * altura) / 2
}
